package cosmos

import (
	"crypto/sha256"

	xc "github.com/jumpcrypto/crosschain"

	injethsecp256k1 "github.com/InjectiveLabs/sdk-go/chain/crypto/ethsecp256k1"
	injectivecodec "github.com/InjectiveLabs/sdk-go/chain/types"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	authVestingTypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/evmos/ethermint/crypto/ethsecp256k1"
	ethermintcodec "github.com/evmos/ethermint/encoding/codec"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	paramsclient "github.com/cosmos/cosmos-sdk/x/params/client"

	wasmclient "github.com/CosmWasm/wasmd/x/wasm/client"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/legacy"
	"github.com/cosmos/cosmos-sdk/std"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth"
	cosmTx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting"
	authzmodule "github.com/cosmos/cosmos-sdk/x/authz/module"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/capability"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	distr "github.com/cosmos/cosmos-sdk/x/distribution"
	distrclient "github.com/cosmos/cosmos-sdk/x/distribution/client"
	"github.com/cosmos/cosmos-sdk/x/evidence"
	feegrantmodule "github.com/cosmos/cosmos-sdk/x/feegrant/module"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	"github.com/cosmos/cosmos-sdk/x/gov"
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
	"github.com/cosmos/cosmos-sdk/x/mint"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	"github.com/cosmos/cosmos-sdk/x/staking"
	"github.com/cosmos/cosmos-sdk/x/upgrade"
	upgradeclient "github.com/cosmos/cosmos-sdk/x/upgrade/client"
	ica "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts"
	transfer "github.com/cosmos/ibc-go/v3/modules/apps/transfer"
	ibc "github.com/cosmos/ibc-go/v3/modules/core"
	ibcclientclient "github.com/cosmos/ibc-go/v3/modules/core/02-client/client"
	"github.com/strangelove-ventures/packet-forward-middleware/v2/router"

	"github.com/CosmWasm/wasmd/x/wasm"
)

const LEN_NATIVE_ASSET = 8

var legacyCodecRegistered = false

// EncodingConfig specifies the concrete encoding types to use for a given app.
// This is provided for compatibility between protobuf and amino implementations.
type EncodingConfig struct {
	InterfaceRegistry codectypes.InterfaceRegistry
	Marshaler         codec.Codec
	TxConfig          client.TxConfig
	Amino             *codec.LegacyAmino
}

func getGovProposalHandlers() []govclient.ProposalHandler {
	var govProposalHandlers []govclient.ProposalHandler
	// this line is used by starport scaffolding # stargate/app/govProposalHandlers
	govProposalHandlers = wasmclient.ProposalHandlers

	govProposalHandlers = append(govProposalHandlers,
		paramsclient.ProposalHandler,
		distrclient.ProposalHandler,
		upgradeclient.ProposalHandler,
		upgradeclient.CancelProposalHandler,
		ibcclientclient.UpdateClientProposalHandler,
		ibcclientclient.UpgradeProposalHandler,
		// this line is used by starport scaffolding # stargate/app/govProposalHandler
	)

	return govProposalHandlers
}

// ModuleBasics defines the module BasicManager is in charge of setting up basic,
// non-dependant module elements, such as codec registration
// and genesis verification.
var ModuleBasics = module.NewBasicManager(
	auth.AppModuleBasic{},
	genutil.AppModuleBasic{},
	bank.AppModuleBasic{},
	capability.AppModuleBasic{},
	staking.AppModuleBasic{},
	mint.AppModuleBasic{},
	distr.AppModuleBasic{},
	gov.NewAppModuleBasic(getGovProposalHandlers()...),
	params.AppModuleBasic{},
	crisis.AppModuleBasic{},
	slashing.AppModuleBasic{},
	ibc.AppModuleBasic{},
	feegrantmodule.AppModuleBasic{},
	upgrade.AppModuleBasic{},
	evidence.AppModuleBasic{},
	transfer.AppModuleBasic{},
	vesting.AppModuleBasic{},
	ica.AppModuleBasic{},
	router.AppModuleBasic{},
	authzmodule.AppModuleBasic{},
	// this line is used by starport scaffolding # stargate/app/moduleBasic
	wasm.AppModuleBasic{},
)

func NewEncodingConfig() EncodingConfig {
	amino := codec.NewLegacyAmino()
	interfaceRegistry := codectypes.NewInterfaceRegistry()
	marshaler := codec.NewProtoCodec(interfaceRegistry)
	txCfg := cosmTx.NewTxConfig(marshaler, cosmTx.DefaultSignModes)

	return EncodingConfig{
		InterfaceRegistry: interfaceRegistry,
		Marshaler:         marshaler,
		TxConfig:          txCfg,
		Amino:             amino,
	}
}

// MakeEncodingConfig creates an EncodingConfig for testing
func MakeEncodingConfig() EncodingConfig {
	encodingConfig := NewEncodingConfig()
	std.RegisterLegacyAminoCodec(encodingConfig.Amino)
	std.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	ModuleBasics.RegisterLegacyAminoCodec(encodingConfig.Amino)
	ModuleBasics.RegisterInterfaces(encodingConfig.InterfaceRegistry)

	if !legacyCodecRegistered {
		// authz module use this codec to get signbytes.
		// authz MsgExec can execute all message types,
		// so legacy.Cdc need to register all amino messages to get proper signature
		ModuleBasics.RegisterLegacyAminoCodec(legacy.Cdc)
		legacyCodecRegistered = true
	}

	return encodingConfig
}

func MakeCosmosConfig() EncodingConfig {
	cosmosCfg := MakeEncodingConfig()
	ethermintcodec.RegisterInterfaces(cosmosCfg.InterfaceRegistry)
	injectivecodec.RegisterInterfaces(cosmosCfg.InterfaceRegistry)
	authVestingTypes.RegisterInterfaces(cosmosCfg.InterfaceRegistry)
	cosmosCfg.InterfaceRegistry.RegisterImplementations((*cryptotypes.PubKey)(nil), &injethsecp256k1.PubKey{})
	return cosmosCfg
}

func isNativeAsset(asset *xc.AssetConfig) bool {
	return asset.Type == xc.AssetTypeNative || len(asset.Contract) < LEN_NATIVE_ASSET
}

func isEVMOS(asset xc.NativeAssetConfig) bool {
	return xc.Driver(asset.Driver) == xc.DriverCosmosEvmos
}

func isINJ(asset xc.NativeAssetConfig) bool {
	return asset.NativeAsset == xc.NativeAsset("INJ")
}

func getPublicKey(asset xc.NativeAssetConfig, publicKeyBytes xc.PublicKey) cryptotypes.PubKey {
	if isINJ(asset) {
		return &injethsecp256k1.PubKey{Key: publicKeyBytes}
	}
	if isEVMOS(asset) {
		return &ethsecp256k1.PubKey{Key: publicKeyBytes}
	}
	return &secp256k1.PubKey{Key: publicKeyBytes}
}

func getSighash(asset xc.NativeAssetConfig, sigData []byte) []byte {
	if isEVMOS(asset) || isINJ(asset) {
		return crypto.Keccak256(sigData)
	}
	sighash := sha256.Sum256(sigData)
	return sighash[:]
}
