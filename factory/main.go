package factory

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"strings"
	"sync"

	"github.com/jinzhu/copier"
	"github.com/shopspring/decimal"
	"gopkg.in/yaml.v2"

	. "github.com/jumpcrypto/crosschain"
	"github.com/jumpcrypto/crosschain/chain/aptos"
	"github.com/jumpcrypto/crosschain/chain/bitcoin"
	"github.com/jumpcrypto/crosschain/chain/cosmos"
	"github.com/jumpcrypto/crosschain/chain/evm"
	"github.com/jumpcrypto/crosschain/chain/solana"
	"github.com/jumpcrypto/crosschain/config"
)

// FactoryContext is the main Factory interface
type FactoryContext interface {
	NewClient(asset ITask) (Client, error)
	NewTxBuilder(asset ITask) (TxBuilder, error)
	NewSigner(asset ITask) (Signer, error)
	NewAddressBuilder(asset ITask) (AddressBuilder, error)

	MarshalTxInput(input TxInput) ([]byte, error)
	UnmarshalTxInput(data []byte) (TxInput, error)

	GetAddressFromPublicKey(asset ITask, publicKey []byte) (Address, error)
	GetAllPossibleAddressesFromPublicKey(asset ITask, publicKey []byte) ([]PossibleAddress, error)

	MustAmountBlockchain(asset ITask, humanAmountStr string) AmountBlockchain
	MustAddress(asset ITask, addressStr string) Address
	MustPrivateKey(asset ITask, privateKey string) PrivateKey

	ConvertAmountToHuman(asset ITask, blockchainAmount AmountBlockchain) (AmountHumanReadable, error)
	ConvertAmountToBlockchain(asset ITask, humanAmount AmountHumanReadable) (AmountBlockchain, error)
	ConvertAmountStrToBlockchain(asset ITask, humanAmountStr string) (AmountBlockchain, error)

	EnrichAssetConfig(partialCfg *TokenAssetConfig) (*TokenAssetConfig, error)
	EnrichDestinations(asset ITask, txInfo TxInfo) (TxInfo, error)

	GetAssetConfig(asset string, nativeAsset string) (ITask, error)
	GetAssetConfigByContract(contract string, nativeAsset string) (ITask, error)
	PutAssetConfig(config ITask) (ITask, error)
	Config() interface{}

	GetTaskConfig(taskName string, assetID AssetID) (ITask, error)
	GetMultiAssetConfig(srcAssetID AssetID, dstAssetID AssetID) ([]ITask, error)
}

// Factory is the main Factory implementation, holding the config
type Factory struct {
	AllAssets    sync.Map
	AllTasks     []*TaskConfig
	AllPipelines []*PipelineConfig
}

var _ FactoryContext = &Factory{}

func (f *Factory) GetAllAssets() []ITask {
	tasks := []ITask{}
	f.AllAssets.Range(func(key, value any) bool {
		asset := value.(ITask)
		task, _ := f.cfgFromAsset(asset.ID())
		tasks = append(tasks, task)
		return true
	})
	return tasks
}

func (f *Factory) cfgFromAsset(assetID AssetID) (ITask, error) {
	cfgI, found := f.AllAssets.Load(assetID)
	if !found {
		return &NativeAssetConfig{}, fmt.Errorf("invalid asset: '%s'", assetID)
	}
	if cfg, ok := cfgI.(*NativeAssetConfig); ok {
		// native asset
		cfg.Type = AssetTypeNative
		cfg.Chain = cfg.Asset
		cfg.NativeAsset = NativeAsset(cfg.Asset)
		return cfg, nil
	}
	if cfg, ok := cfgI.(*TokenAssetConfig); ok {
		// token
		copier.CopyWithOption(&cfg.AssetConfig, &cfg, copier.Option{IgnoreEmpty: false, DeepCopy: false})
		cfg, _ = f.cfgEnrichAssetConfig(cfg)
		return cfg, nil
	}
	return &NativeAssetConfig{}, fmt.Errorf("invalid asset: '%s'", assetID)
}

func (f *Factory) enrichTask(task *TaskConfig, srcAssetID AssetID, dstAssetID AssetID) (*TaskConfig, error) {
	dstAsset, err := f.cfgFromAsset(dstAssetID)
	if err != nil {
		return task, fmt.Errorf("task '%s' has invalid dst asset: '%s'", task.ID(), dstAssetID)
	}

	newTask := *task
	newTask.SrcAsset, _ = f.cfgFromAsset(srcAssetID)
	newTask.DstAsset = dstAsset
	return &newTask, nil
}

func (f *Factory) cfgFromTask(taskName string, assetID AssetID) (ITask, error) {
	IsAllowedFunc := func(task *TaskConfig, assetID AssetID) (*TaskConfig, error) {
		allowed := false
		dstAssetID := AssetID("")
		for _, entry := range task.AllowList {
			if entry.Src == assetID {
				allowed = true
				dstAssetID = entry.Dst
				fmt.Println(dstAssetID)
				break
			}
		}
		if !allowed {
			return task, fmt.Errorf("task '%s' not allowed: '%s'", taskName, assetID)
		}
		return f.enrichTask(task, assetID, dstAssetID)
	}

	assetCfg, err := f.cfgFromAsset(assetID)
	if taskName == "" {
		return assetCfg, err
	}

	task, err := f.findTask(taskName)
	if err != nil {
		return &TaskConfig{}, fmt.Errorf("invalid task: '%s'", taskName)
	}

	res, err := IsAllowedFunc(task, assetID)
	return res, err
}

func (f *Factory) findTask(taskName string) (*TaskConfig, error) {
	// TODO: switch to map
	for _, task := range f.AllTasks {
		if string(task.ID()) == taskName {
			return task, nil
		}
	}
	return &TaskConfig{}, fmt.Errorf("invalid task: '%s'", taskName)
}

func (f *Factory) cfgFromMultiAsset(srcAssetID AssetID, dstAssetID AssetID) ([]ITask, error) {
	srcAsset, err := f.cfgFromAsset(srcAssetID)
	if err != nil {
		return []ITask{}, fmt.Errorf("invalid src asset in: '%s -> %s'", srcAssetID, dstAssetID)
	}
	if dstAssetID == "" {
		return []ITask{srcAsset}, err
	}
	_, err = f.cfgFromAsset(dstAssetID)
	if err != nil {
		return []ITask{}, fmt.Errorf("invalid dst asset in: '%s -> %s'", srcAssetID, dstAssetID)
	}

	for _, task := range f.AllTasks {
		for _, entry := range task.AllowList {
			if entry.Src == srcAssetID && entry.Dst == dstAssetID {
				newTask, err := f.enrichTask(task, srcAssetID, dstAssetID)
				return []ITask{newTask}, err
			}
		}
	}

	for _, pipeline := range f.AllPipelines {
		for _, entry := range pipeline.AllowList {
			if entry.Src == srcAssetID && entry.Dst == dstAssetID {
				result := []ITask{}
				for _, taskName := range pipeline.Tasks {
					task, err := f.findTask(taskName)
					if err != nil {
						return []ITask{}, fmt.Errorf("pipeline '%s' has invalid task: '%s'", pipeline.ID, taskName)
					}
					newTask, err := f.enrichTask(task, srcAssetID, dstAssetID)
					if err != nil {
						return []ITask{}, fmt.Errorf("pipeline '%s' can't enrich task: '%s' %s -> %s", pipeline.ID, taskName, srcAssetID, dstAssetID)
					}
					result = append(result, newTask)
				}
				return result, err
			}
		}
	}

	return []ITask{}, fmt.Errorf("invalid path: '%s -> %s'", srcAssetID, dstAssetID)
}

func (f *Factory) cfgEnrichAssetConfig(partialCfg *TokenAssetConfig) (*TokenAssetConfig, error) {
	cfg := partialCfg
	if cfg.Chain != "" {
		// token
		cfg.Type = AssetTypeToken
		nativeAsset := cfg.Chain
		cfg.NativeAsset = NativeAsset(nativeAsset)

		chainI, found := f.AllAssets.Load(AssetID(nativeAsset))
		if !found {
			return cfg, fmt.Errorf("unsupported native asset: %s", nativeAsset)
		}
		chain := chainI.(*NativeAssetConfig)
		cfg.NativeAssetConfig = chain

		cfg.Driver = chain.Driver
		cfg.Net = chain.Net
		cfg.URL = chain.URL
		cfg.FcdURL = chain.FcdURL
		cfg.Auth = chain.Auth
		cfg.AuthSecret = chain.AuthSecret
		cfg.Provider = chain.Provider
		cfg.ChainID = chain.ChainID
		cfg.ChainIDStr = chain.ChainIDStr
		cfg.ChainGasMultiplier = chain.ChainGasMultiplier
		cfg.ExplorerURL = chain.ExplorerURL
	} else {
		return cfg, fmt.Errorf("unsupported native asset: (empty)")
	}
	return cfg, nil
}

func (f *Factory) cfgEnrichDestinations(activity ITask, txInfo TxInfo) (TxInfo, error) {
	asset := activity.GetAssetConfig()
	result := txInfo
	nativeAssetCfg := activity.GetNativeAsset()
	for i, dst := range txInfo.Destinations {
		dst.NativeAsset = asset.NativeAsset
		if dst.ContractAddress != "" {
			assetCfgI, err := f.cfgFromAssetByContract(string(dst.ContractAddress), string(dst.NativeAsset))
			if err != nil {
				// we shouldn't set the amount, if we don't know the contract
				continue
			}
			assetCfg := assetCfgI.GetAssetConfig()
			dst.Asset = Asset(assetCfg.Asset)
			dst.ContractAddress = ContractAddress(assetCfg.Contract)
			dst.AssetConfig = assetCfg
		} else {
			dst.AssetConfig = nativeAssetCfg
		}
		result.Destinations[i] = dst
	}
	return result, nil
}

func (f *Factory) cfgFromAssetByContract(contract string, nativeAsset string) (ITask, error) {
	var res ITask
	nativeAsset = strings.ToUpper(nativeAsset)
	contract = NormalizeAddressString(contract, nativeAsset)
	f.AllAssets.Range(func(key, value interface{}) bool {
		cfg := value.(ITask).GetAssetConfig()
		if cfg.Chain == nativeAsset {
			cfgContract := NormalizeAddressString(cfg.Contract, nativeAsset)
			if cfgContract == contract {
				res = value.(ITask)
				return false
			}
		} else if cfg.Asset == nativeAsset && cfg.ChainCoin == contract {
			res = value.(ITask)
			return false
		}
		return true
	})
	if res != nil {
		return f.cfgFromAsset(res.ID())
	}
	return &TokenAssetConfig{}, fmt.Errorf("invalid contract: '%s'", contract)
}

// NewClient creates a new Client
func (f *Factory) NewClient(cfg ITask) (Client, error) {
	return newClient(cfg)
}

// NewTxBuilder creates a new TxBuilder
func (f *Factory) NewTxBuilder(cfg ITask) (TxBuilder, error) {
	return newTxBuilder(cfg)
}

// NewSigner creates a new Signer
func (f *Factory) NewSigner(cfg ITask) (Signer, error) {
	return newSigner(cfg)
}

// NewAddressBuilder creates a new AddressBuilder
func (f *Factory) NewAddressBuilder(cfg ITask) (AddressBuilder, error) {
	return newAddressBuilder(cfg)
}

// MarshalTxInput marshalls a TxInput struct
func (f *Factory) MarshalTxInput(input TxInput) ([]byte, error) {
	return MarshalTxInput(input)
}

// UnmarshalTxInput unmarshalls data into a TxInput struct
func (f *Factory) UnmarshalTxInput(data []byte) (TxInput, error) {
	return UnmarshalTxInput(data)
}

// GetAddressFromPublicKey returns an Address given a public key
func (f *Factory) GetAddressFromPublicKey(cfg ITask, publicKey []byte) (Address, error) {
	return getAddressFromPublicKey(cfg, publicKey)
}

// GetAllPossibleAddressesFromPublicKey returns all PossibleAddress(es) given a public key
func (f *Factory) GetAllPossibleAddressesFromPublicKey(cfg ITask, publicKey []byte) ([]PossibleAddress, error) {
	builder, err := newAddressBuilder(cfg)
	if err != nil {
		return []PossibleAddress{}, err
	}
	return builder.GetAllPossibleAddressesFromPublicKey(publicKey)
}

// ConvertAmountToHuman converts an AmountBlockchain into AmountHumanReadable, dividing by the appropriate number of decimals
func (f *Factory) ConvertAmountToHuman(cfg ITask, blockchainAmount AmountBlockchain) (AmountHumanReadable, error) {
	return convertAmountToHuman(cfg, blockchainAmount)
}

// ConvertAmountToBlockchain converts an AmountHumanReadable into AmountBlockchain, multiplying by the appropriate number of decimals
func (f *Factory) ConvertAmountToBlockchain(cfg ITask, humanAmount AmountHumanReadable) (AmountBlockchain, error) {
	return convertAmountToBlockchain(cfg, humanAmount)
}

// ConvertAmountStrToBlockchain converts a string representing an AmountHumanReadable into AmountBlockchain, multiplying by the appropriate number of decimals
func (f *Factory) ConvertAmountStrToBlockchain(cfg ITask, humanAmountStr string) (AmountBlockchain, error) {
	return convertAmountStrToBlockchain(cfg, humanAmountStr)
}

// EnrichAssetConfig augments a partial AssetConfig, for example if some info is stored in a db and other in a config file
func (f *Factory) EnrichAssetConfig(partialCfg *TokenAssetConfig) (*TokenAssetConfig, error) {
	return f.cfgEnrichAssetConfig(partialCfg)
}

// EnrichDestinations augments a TxInfo by resolving assets and amounts in TxInfo.Destinations
func (f *Factory) EnrichDestinations(activity ITask, txInfo TxInfo) (TxInfo, error) {
	return f.cfgEnrichDestinations(activity, txInfo)
}

// GetAssetConfig returns an AssetConfig by asset and native asset (chain)
func (f *Factory) GetAssetConfig(asset string, nativeAsset string) (ITask, error) {
	assetID := GetAssetIDFromAsset(asset, nativeAsset)
	return f.cfgFromAsset(assetID)
}

// GetTaskConfig returns an AssetConfig by task name and assetID
func (f *Factory) GetTaskConfig(taskName string, assetID AssetID) (ITask, error) {
	return f.cfgFromTask(taskName, assetID)
}

// GetMultiAssetConfig returns an AssetConfig by source and destination assetIDs
func (f *Factory) GetMultiAssetConfig(srcAssetID AssetID, dstAssetID AssetID) ([]ITask, error) {
	return f.cfgFromMultiAsset(srcAssetID, dstAssetID)
}

// GetAssetConfigByContract returns an AssetConfig by contract and native asset (chain)
func (f *Factory) GetAssetConfigByContract(contract string, nativeAsset string) (ITask, error) {
	return f.cfgFromAssetByContract(contract, nativeAsset)
}

// PutAssetConfig adds an AssetConfig to the current Config cache
func (f *Factory) PutAssetConfig(cfgI ITask) (ITask, error) {
	f.AllAssets.Store(cfgI.ID(), cfgI)
	return f.cfgFromAsset(cfgI.ID())
}

// Config returns the Config
func (f *Factory) Config() interface{} {
	return f.AllAssets
}

// MustAddress coverts a string to Address, panic if error
func (f *Factory) MustAddress(cfg ITask, addressStr string) Address {
	return Address(addressStr)
}

// MustAmountBlockchain coverts a string into AmountBlockchain, panic if error
func (f *Factory) MustAmountBlockchain(cfg ITask, humanAmountStr string) AmountBlockchain {
	res, err := f.ConvertAmountStrToBlockchain(cfg, humanAmountStr)
	if err != nil {
		panic(err)
	}
	return res
}

// MustPrivateKey coverts a string into PrivateKey, panic if error
func (f *Factory) MustPrivateKey(cfg ITask, privateKeyStr string) PrivateKey {
	signer, err := f.NewSigner(cfg)
	if err != nil {
		panic(err)
	}
	privateKey, err := signer.ImportPrivateKey(privateKeyStr)
	if err != nil {
		panic(err)
	}
	return privateKey
}

func assetsFromConfig(configMap map[string]interface{}) []ITask {
	yamlStr, _ := yaml.Marshal(configMap)
	var mainConfig Config
	yaml.Unmarshal(yamlStr, &mainConfig)

	var allAssets []ITask
	for _, c := range mainConfig.Chains {
		allAssets = append(allAssets, c)
	}

	for _, t := range mainConfig.Tokens {
		copier.CopyWithOption(&t.AssetConfig, &t, copier.Option{IgnoreEmpty: false, DeepCopy: false})
		allAssets = append(allAssets, t)
	}

	return allAssets
}

func parseAllowList(allowList []string) []*AllowEntry {
	result := []*AllowEntry{}
	for _, allow := range allowList {
		var entry AllowEntry
		values := strings.Split(allow, "->")
		if len(values) == 1 {
			value := AssetID(strings.TrimSpace(values[0]))
			entry = AllowEntry{
				Src: value,
				Dst: value,
			}
		}
		if len(values) == 2 {
			src := AssetID(strings.TrimSpace(values[0]))
			dst := AssetID(strings.TrimSpace(values[1]))
			entry = AllowEntry{
				Src: src,
				Dst: dst,
			}
		}
		result = append(result, &entry)
	}
	return result
}

func tasksFromConfig(configMap map[string]interface{}) []*TaskConfig {
	yamlStr, _ := yaml.Marshal(configMap)
	var mainConfig Config
	yaml.Unmarshal(yamlStr, &mainConfig)
	for _, task := range mainConfig.AllTasks {
		task.AllowList = parseAllowList(task.Allow)
	}
	return mainConfig.AllTasks
}

func pipelinesFromConfig(configMap map[string]interface{}) []*PipelineConfig {
	yamlStr, _ := yaml.Marshal(configMap)
	var mainConfig Config
	yaml.Unmarshal(yamlStr, &mainConfig)
	for _, pipeline := range mainConfig.AllPipelines {
		pipeline.AllowList = parseAllowList(pipeline.Allow)
	}
	return mainConfig.AllPipelines
}

// NewDefaultFactory creates a new Factory
func NewDefaultFactory() *Factory {
	// Use our config file loader
	cfg := config.RequireConfig("crosschain")
	return NewDefaultFactoryWithConfig(cfg)
}

// NewDefaultFactoryWithConfig creates a new Factory given a config map
func NewDefaultFactoryWithConfig(cfg map[string]interface{}) *Factory {
	assetsList := assetsFromConfig(cfg)
	assetsMap := AssetsToMap(assetsList)

	tasksList := tasksFromConfig(cfg)
	pipelinesList := pipelinesFromConfig(cfg)

	return &Factory{
		AllAssets:    assetsMap,
		AllTasks:     tasksList,
		AllPipelines: pipelinesList,
	}
}

// AssetsToMap loads chains config without config file
func AssetsToMap(assetsList []ITask) sync.Map {
	assetsMap := sync.Map{}
	for _, cfgI := range assetsList {
		cfg := cfgI.GetAssetConfig()
		if cfg.Auth != "" {
			var err error
			cfgI.(*NativeAssetConfig).AuthSecret, err = config.GetSecret(cfg.Auth)
			if err != nil {
				// ignore error
			}
		}
		assetsMap.Store(cfgI.ID(), cfgI)
	}
	return assetsMap
}

func newClient(cfg ITask) (Client, error) {
	switch Driver(cfg.GetDriver()) {
	case DriverEVM:
		return evm.NewClient(cfg)
	case DriverEVMLegacy:
		return evm.NewLegacyClient(cfg)
	case DriverCosmos, DriverCosmosEvmos:
		return cosmos.NewClient(cfg)
	case DriverSolana:
		return solana.NewClient(cfg)
	case DriverAptos:
		return aptos.NewClient(cfg)
	case DriverBitcoin:
		return bitcoin.NewClient(cfg)
	}
	return nil, errors.New("unsupported asset")
}

func newTxBuilder(cfg ITask) (TxBuilder, error) {
	switch Driver(cfg.GetDriver()) {
	case DriverEVM:
		return evm.NewTxBuilder(cfg)
	case DriverEVMLegacy:
		return evm.NewLegacyTxBuilder(cfg)
	case DriverCosmos, DriverCosmosEvmos:
		return cosmos.NewTxBuilder(cfg)
	case DriverSolana:
		return solana.NewTxBuilder(cfg)
	case DriverAptos:
		return aptos.NewTxBuilder(cfg)
	case DriverBitcoin:
		return bitcoin.NewTxBuilder(cfg)
	}
	return nil, errors.New("unsupported asset")
}

func newSigner(cfg ITask) (Signer, error) {
	switch Driver(cfg.GetDriver()) {
	case DriverEVM, DriverEVMLegacy:
		return evm.NewSigner(cfg)
	case DriverCosmos, DriverCosmosEvmos:
		return cosmos.NewSigner(cfg)
	case DriverSolana:
		return solana.NewSigner(cfg)
	case DriverAptos:
		return aptos.NewSigner(cfg)
	case DriverBitcoin:
		return bitcoin.NewSigner(cfg)
	}
	return nil, errors.New("unsupported asset")
}

func newAddressBuilder(cfg ITask) (AddressBuilder, error) {
	switch Driver(cfg.GetDriver()) {
	case DriverEVM, DriverEVMLegacy:
		return evm.NewAddressBuilder(cfg)
	case DriverCosmos, DriverCosmosEvmos:
		return cosmos.NewAddressBuilder(cfg)
	case DriverSolana:
		return solana.NewAddressBuilder(cfg)
	case DriverAptos:
		return aptos.NewAddressBuilder(cfg)
	case DriverBitcoin:
		return bitcoin.NewAddressBuilder(cfg)
	}
	return nil, errors.New("unsupported asset")
}

func MarshalTxInput(txInput TxInput) ([]byte, error) {
	return json.Marshal(txInput)
}

func UnmarshalTxInput(data []byte) (TxInput, error) {
	var env TxInputEnvelope
	buf := []byte(data)
	err := json.Unmarshal(buf, &env)
	if err != nil {
		return nil, err
	}
	switch env.Type {
	case DriverAptos:
		var txInput aptos.TxInput
		err := json.Unmarshal(buf, &txInput)
		return &txInput, err
	case DriverCosmos, DriverCosmosEvmos:
		var txInput cosmos.TxInput
		err := json.Unmarshal(buf, &txInput)
		return &txInput, err
	case DriverEVM, DriverEVMLegacy:
		var txInput evm.TxInput
		err := json.Unmarshal(buf, &txInput)
		return &txInput, err
	case DriverSolana:
		var txInput solana.TxInput
		err := json.Unmarshal(buf, &txInput)
		return &txInput, err
	case DriverBitcoin:
		var txInput bitcoin.TxInput
		err := json.Unmarshal(buf, &txInput)
		return &txInput, err
	default:
		return nil, fmt.Errorf("invalid TxInput type: %s", env.Type)
	}
}

func getAddressFromPublicKey(cfg ITask, publicKey []byte) (Address, error) {
	builder, err := newAddressBuilder(cfg)
	if err != nil {
		return "", err
	}
	return builder.GetAddressFromPublicKey(publicKey)
}

// Amount converter

func convertAmountExponent(cfgI ITask) (int32, error) {
	cfg := cfgI.GetAssetConfig()
	if cfg.Decimals > 0 {
		return cfg.Decimals, nil
	}
	return 0, errors.New("unsupported asset")
}

func convertAmountToHuman(cfg ITask, blockchainAmount AmountBlockchain) (AmountHumanReadable, error) {
	exponent, err := convertAmountExponent(cfg)
	if err != nil {
		return AmountHumanReadable(decimal.NewFromInt(0)), err
	}
	blockchainAmountInt := big.Int(blockchainAmount)
	result := decimal.NewFromBigInt(&blockchainAmountInt, -exponent)
	return AmountHumanReadable(result), nil
}

func convertAmountToBlockchain(cfg ITask, humanAmount AmountHumanReadable) (AmountBlockchain, error) {
	exponent, err := convertAmountExponent(cfg)
	if err != nil {
		return AmountBlockchain(*big.NewInt(0)), err
	}
	result := decimal.Decimal(humanAmount).Shift(exponent).BigInt()
	return AmountBlockchain(*result), nil
}

func convertAmountStrToBlockchain(cfg ITask, humanAmountStr string) (AmountBlockchain, error) {
	humanAmount, err := decimal.NewFromString(humanAmountStr)
	if err != nil {
		return AmountBlockchain(*big.NewInt(0)), err
	}

	return convertAmountToBlockchain(cfg, AmountHumanReadable(humanAmount))
}

// NormalizeAddressString normalizes an address, e.g. returns lowercase when possible
func NormalizeAddressString(address string, nativeAsset string) string {
	if nativeAsset == "" {
		nativeAsset = "ETH"
	}

	address = strings.TrimSpace(address)
	switch NativeAsset(nativeAsset) {
	// hex formatted addresses
	case ETH,
		AVAX, ArbETH, CELO, MATIC, OptETH,
		ETC, FTM, BNB, ROSE, ACA, KAR, KLAY, AurETH, CHZ,
		APTOS:
		if strings.HasPrefix(address, "0x") {
			return strings.ToLower(address)
		}
	case XDC:
		if strings.HasPrefix(address, "0x") || strings.HasPrefix(address, "xdc") {
			return strings.ToLower(address)
		}
	case BCH:
		// remove bitcoincash: prefix
		if strings.Contains(address, ":") {
			return strings.Split(address, ":")[1]
		}
	default:
	}
	return address
}

func CheckError(cfg *AssetConfig, err error) ClientError {
	switch Driver(cfg.Driver) {
	case DriverEVM, DriverEVMLegacy:
		return evm.CheckError(err)
	case DriverCosmos, DriverCosmosEvmos:
		return cosmos.CheckError(err)
	case DriverSolana:
		return solana.CheckError(err)
	case DriverAptos:
		return aptos.CheckError(err)
	case DriverBitcoin:
		return bitcoin.CheckError(err)
	}
	return UnknownError
}
