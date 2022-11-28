package factory

import (
	"errors"
	"fmt"
	"math/big"
	"strings"
	"sync"

	"github.com/shopspring/decimal"
	"gopkg.in/yaml.v2"

	. "github.com/jumpcrypto/crosschain"
	"github.com/jumpcrypto/crosschain/chain/aptos"
	"github.com/jumpcrypto/crosschain/chain/cosmos"
	"github.com/jumpcrypto/crosschain/chain/evm"
	"github.com/jumpcrypto/crosschain/chain/solana"
	"github.com/jumpcrypto/crosschain/config"
)

// FactoryContext is the main Factory interface
type FactoryContext interface {
	NewClient(asset AssetConfig) (Client, error)
	NewTxBuilder(asset AssetConfig) (TxBuilder, error)
	NewSigner(asset AssetConfig) (Signer, error)
	NewAddressBuilder(asset AssetConfig) (AddressBuilder, error)

	GetAddressFromPublicKey(asset AssetConfig, publicKey []byte) (Address, error)
	GetAllPossibleAddressesFromPublicKey(asset AssetConfig, publicKey []byte) ([]PossibleAddress, error)

	MustAmountBlockchain(asset AssetConfig, humanAmountStr string) AmountBlockchain
	MustAddress(asset AssetConfig, addressStr string) Address
	MustPrivateKey(asset AssetConfig, privateKey string) PrivateKey

	ConvertAmountToHuman(asset AssetConfig, blockchainAmount AmountBlockchain) (AmountHumanReadable, error)
	ConvertAmountToBlockchain(asset AssetConfig, humanAmount AmountHumanReadable) (AmountBlockchain, error)
	ConvertAmountStrToBlockchain(asset AssetConfig, humanAmountStr string) (AmountBlockchain, error)

	EnrichAssetConfig(partialCfg AssetConfig) (AssetConfig, error)

	GetAssetID(asset string, nativeAsset string) AssetID
	GetAssetConfig(asset string, nativeAsset string) (AssetConfig, error)
	GetAssetConfigByContract(contract string, nativeAsset string) (AssetConfig, error)
	PutAssetConfig(config AssetConfig) (AssetConfig, error)
	Config() interface{}
}

// Factory is the main Factory implementation, holding the config
type Factory struct {
	AllAssets *sync.Map
}

var _ FactoryContext = &Factory{}

func (f *Factory) cfgFromAsset(assetID AssetID) (AssetConfig, error) {
	cfgI, found := f.AllAssets.Load(assetID)
	if !found {
		return AssetConfig{}, fmt.Errorf("invalid asset: '%s'", assetID)
	}
	cfg := cfgI.(AssetConfig)
	if cfg.Chain != "" {
		// token
		cfg.Type = AssetTypeToken
		nativeAsset := cfg.Chain
		cfg.NativeAsset = NativeAsset(nativeAsset)

		chainI, _ := f.AllAssets.Load(AssetID(nativeAsset))
		chain := chainI.(AssetConfig)
		cfg.Net = chain.Net
		cfg.URL = chain.URL
		cfg.Auth = chain.Auth
		cfg.AuthSecret = chain.AuthSecret
		cfg.Provider = chain.Provider
	} else {
		// native asset
		cfg.Type = AssetTypeNative
		cfg.Chain = cfg.Asset
		cfg.NativeAsset = NativeAsset(cfg.Asset)
	}
	return cfg, nil
}

func (f *Factory) cfgEnrichAssetConfig(partialCfg AssetConfig) (AssetConfig, error) {
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
		chain := chainI.(AssetConfig)
		cfg.Net = chain.Net
		cfg.URL = chain.URL
		cfg.Auth = chain.Auth
		cfg.AuthSecret = chain.AuthSecret
		cfg.Provider = chain.Provider
	} else {
		return cfg, fmt.Errorf("unsupported native asset: (empty)")
	}
	return cfg, nil
}

func (f *Factory) cfgFromAssetByContract(contract string, nativeAsset string) (AssetConfig, error) {
	var res *AssetConfig
	nativeAsset = strings.ToUpper(nativeAsset)
	contract = NormalizeAddressString(contract, nativeAsset)
	f.AllAssets.Range(func(key, value interface{}) bool {
		cfg := value.(AssetConfig)
		if cfg.Chain == nativeAsset {
			cfgContract := NormalizeAddressString(cfg.Contract, nativeAsset)
			if cfgContract == contract {
				res = &cfg
				return false
			}
		}
		return true
	})
	if res != nil {
		return f.cfgFromAsset(res.ID)
	}
	return AssetConfig{}, fmt.Errorf("invalid contract: '%s'", contract)
}

// NewClient creates a new Client
func (f *Factory) NewClient(cfg AssetConfig) (Client, error) {
	return newClient(cfg)
}

// NewTxBuilder creates a new TxBuilder
func (f *Factory) NewTxBuilder(cfg AssetConfig) (TxBuilder, error) {
	return newTxBuilder(cfg)
}

// NewSigner creates a new Signer
func (f *Factory) NewSigner(cfg AssetConfig) (Signer, error) {
	return newSigner(cfg)
}

// NewAddressBuilder creates a new AddressBuilder
func (f *Factory) NewAddressBuilder(cfg AssetConfig) (AddressBuilder, error) {
	return newAddressBuilder(cfg)
}

// GetAddressFromPublicKey returns an Address given a public key
func (f *Factory) GetAddressFromPublicKey(cfg AssetConfig, publicKey []byte) (Address, error) {
	return getAddressFromPublicKey(cfg, publicKey)
}

// GetAllPossibleAddressesFromPublicKey returns all PossibleAddress(es) given a public key
func (f *Factory) GetAllPossibleAddressesFromPublicKey(cfg AssetConfig, publicKey []byte) ([]PossibleAddress, error) {
	builder, err := newAddressBuilder(cfg)
	if err != nil {
		return []PossibleAddress{}, err
	}
	return builder.GetAllPossibleAddressesFromPublicKey(publicKey)
}

// ConvertAmountToHuman converts an AmountBlockchain into AmountHumanReadable, dividing by the appropriate number of decimals
func (f *Factory) ConvertAmountToHuman(cfg AssetConfig, blockchainAmount AmountBlockchain) (AmountHumanReadable, error) {
	return convertAmountToHuman(cfg, blockchainAmount)
}

// ConvertAmountToBlockchain converts an AmountHumanReadable into AmountBlockchain, multiplying by the appropriate number of decimals
func (f *Factory) ConvertAmountToBlockchain(cfg AssetConfig, humanAmount AmountHumanReadable) (AmountBlockchain, error) {
	return convertAmountToBlockchain(cfg, humanAmount)
}

// ConvertAmountStrToBlockchain converts a string representing an AmountHumanReadable into AmountBlockchain, multiplying by the appropriate number of decimals
func (f *Factory) ConvertAmountStrToBlockchain(cfg AssetConfig, humanAmountStr string) (AmountBlockchain, error) {
	return convertAmountStrToBlockchain(cfg, humanAmountStr)
}

// GetAssetID returns a canonical AssetID
func (f *Factory) GetAssetID(asset string, nativeAsset string) AssetID {
	return GetAssetIDFromAsset(asset, nativeAsset)
}

// EnrichAssetConfig augments a partial AssetConfig, for example if some info is stored in a db and other in a config file
func (f *Factory) EnrichAssetConfig(partialCfg AssetConfig) (AssetConfig, error) {
	return f.cfgEnrichAssetConfig(partialCfg)
}

// GetAssetConfig returns an AssetConfig by asset and native asset (chain)
func (f *Factory) GetAssetConfig(asset string, nativeAsset string) (AssetConfig, error) {
	assetID := f.GetAssetID(asset, nativeAsset)
	return f.cfgFromAsset(assetID)
}

// GetAssetConfigByContract returns an AssetConfig by contract and native asset (chain)
func (f *Factory) GetAssetConfigByContract(contract string, nativeAsset string) (AssetConfig, error) {
	return f.cfgFromAssetByContract(contract, nativeAsset)
}

// PutAssetConfig adds an AssetConfig to the current Config cache
func (f *Factory) PutAssetConfig(cfg AssetConfig) (AssetConfig, error) {
	cfg.ID = f.GetAssetID(cfg.Asset, cfg.Chain)
	f.AllAssets.Store(cfg.ID, cfg)
	return f.cfgFromAsset(cfg.ID)
}

// Config returns the Config
func (f *Factory) Config() interface{} {
	return f.AllAssets
}

// MustAddress coverts a string to Address, panic if error
func (f *Factory) MustAddress(cfg AssetConfig, addressStr string) Address {
	return Address(addressStr)
}

// MustAmountBlockchain coverts a string into AmountBlockchain, panic if error
func (f *Factory) MustAmountBlockchain(cfg AssetConfig, humanAmountStr string) AmountBlockchain {
	res, err := f.ConvertAmountStrToBlockchain(cfg, humanAmountStr)
	if err != nil {
		panic(err)
	}
	return res
}

// MustPrivateKey coverts a string into PrivateKey, panic if error
func (f *Factory) MustPrivateKey(cfg AssetConfig, privateKeyStr string) PrivateKey {
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

func assetsFromConfig(configMap map[string]interface{}) []AssetConfig {
	yamlStr, _ := yaml.Marshal(configMap)
	var mainConfig Config
	yaml.Unmarshal(yamlStr, &mainConfig)
	return mainConfig.AllAssets
}

// NewDefaultFactory creates a new Factory
func NewDefaultFactory() Factory {
	// Use our config file loader
	cfg := config.RequireConfig("crosschain")
	assetsList := assetsFromConfig(cfg)
	assetsMap := AssetsToMap(assetsList)
	return Factory{
		AllAssets: &assetsMap,
	}
}

// AssetsToMap loads chains config without config file
func AssetsToMap(assetsList []AssetConfig) sync.Map {
	assetsMap := sync.Map{}
	for _, cfg := range assetsList {
		if cfg.Auth != "" {
			var err error
			cfg.AuthSecret, err = config.GetSecret(cfg.Auth)
			if err != nil {
				// ignore error
			}
		}
		cfg.ID = GetAssetIDFromAsset(cfg.Asset, cfg.Chain)
		assetsMap.Store(cfg.ID, cfg)
	}
	return assetsMap
}

func configToEthereumURL(cfg AssetConfig) string {
	if cfg.Provider == "infura" {
		return cfg.URL + "/" + cfg.AuthSecret
	}
	return cfg.URL
}

func newClient(cfg AssetConfig) (Client, error) {
	switch cfg.NativeAsset {
	case ETH:
		return evm.NewClient(cfg)
	case AVAX, ArbETH, CELO, MATIC, OptETH:
		return evm.NewClient(cfg)
	case ETC, FTM, BNB, ROSE, ACA, KAR, KLAY, AurETH:
		return evm.NewLegacyClient(cfg)
	case ATOM, LUNA, XPLA:
		return cosmos.NewClient(cfg)
	case SOL:
		return solana.NewClient(cfg)
	case APTOS:
		return aptos.NewClient(cfg)
	}
	return nil, errors.New("unsupported asset")
}

func newTxBuilder(cfg AssetConfig) (TxBuilder, error) {
	switch cfg.NativeAsset {
	case ETH:
		return evm.NewTxBuilder(cfg)
	case AVAX, ArbETH, CELO, MATIC, OptETH:
		return evm.NewTxBuilder(cfg)
	case ETC, FTM, BNB, ROSE, ACA, KAR, KLAY, AurETH:
		return evm.NewTxBuilder(cfg)
	case ATOM, LUNA, XPLA:
		return cosmos.NewTxBuilder(cfg)
	case SOL:
		return solana.NewTxBuilder(cfg)
	case APTOS:
		return aptos.NewTxBuilder(cfg)
	}
	return nil, errors.New("unsupported asset")
}

func newSigner(cfg AssetConfig) (Signer, error) {
	switch cfg.NativeAsset {
	case ETH:
		return evm.NewSigner(cfg)
	case AVAX, ArbETH, CELO, MATIC, OptETH:
		return evm.NewSigner(cfg)
	case ETC, FTM, BNB, ROSE, ACA, KAR, KLAY, AurETH:
		return evm.NewSigner(cfg)
	case ATOM, LUNA, XPLA:
		return cosmos.NewSigner(cfg)
	case SOL:
		return solana.NewSigner(cfg)
	case APTOS:
		return aptos.NewSigner(cfg)
	}
	return nil, errors.New("unsupported asset")
}

func newAddressBuilder(cfg AssetConfig) (AddressBuilder, error) {
	switch cfg.NativeAsset {
	case ETH:
		return evm.NewAddressBuilder(cfg)
	case AVAX, ArbETH, CELO, MATIC, OptETH:
		return evm.NewAddressBuilder(cfg)
	case ETC, FTM, BNB, ROSE, ACA, KAR, KLAY, AurETH:
		return evm.NewAddressBuilder(cfg)
	case ATOM, LUNA, XPLA:
		return cosmos.NewAddressBuilder(cfg)
	case SOL:
		return solana.NewAddressBuilder(cfg)
	case APTOS:
		return aptos.NewAddressBuilder(cfg)
	}
	return nil, errors.New("unsupported asset")
}

func getAddressFromPublicKey(cfg AssetConfig, publicKey []byte) (Address, error) {
	builder, err := newAddressBuilder(cfg)
	if err != nil {
		return "", err
	}
	return builder.GetAddressFromPublicKey(publicKey)
}

// Amount converter

func convertAmountExponent(cfg AssetConfig) (int32, error) {
	if cfg.Decimals > 0 {
		return cfg.Decimals, nil
	}
	return 0, errors.New("unsupported asset")
}

func convertAmountToHuman(cfg AssetConfig, blockchainAmount AmountBlockchain) (AmountHumanReadable, error) {
	exponent, err := convertAmountExponent(cfg)
	if err != nil {
		return AmountHumanReadable(decimal.NewFromInt(0)), err
	}
	blockchainAmountInt := big.Int(blockchainAmount)
	result := decimal.NewFromBigInt(&blockchainAmountInt, -exponent)
	return AmountHumanReadable(result), nil
}

func convertAmountToBlockchain(cfg AssetConfig, humanAmount AmountHumanReadable) (AmountBlockchain, error) {
	exponent, err := convertAmountExponent(cfg)
	if err != nil {
		return AmountBlockchain(*big.NewInt(0)), err
	}
	result := decimal.Decimal(humanAmount).Shift(exponent).BigInt()
	return AmountBlockchain(*result), nil
}

func convertAmountStrToBlockchain(cfg AssetConfig, humanAmountStr string) (AmountBlockchain, error) {
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
		ETC, FTM, BNB, ROSE, ACA, KAR, KLAY, AurETH,
		APTOS:
		if strings.HasPrefix(address, "0x") {
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
