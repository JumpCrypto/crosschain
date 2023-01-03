package cosmos

import (
	"crypto/sha256"

	xc "github.com/jumpcrypto/crosschain"
	terraApp "github.com/terra-money/core/app"
	"github.com/terra-money/core/app/params"

	injethsecp256k1 "github.com/InjectiveLabs/sdk-go/chain/crypto/ethsecp256k1"
	injectivecodec "github.com/InjectiveLabs/sdk-go/chain/types"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/evmos/ethermint/crypto/ethsecp256k1"
	ethermintcodec "github.com/evmos/ethermint/encoding/codec"
)

const LEN_NATIVE_ASSET = 8

func MakeCosmosConfig() params.EncodingConfig {
	cosmosCfg := terraApp.MakeEncodingConfig()
	ethermintcodec.RegisterInterfaces(cosmosCfg.InterfaceRegistry)
	injectivecodec.RegisterInterfaces(cosmosCfg.InterfaceRegistry)
	cosmosCfg.InterfaceRegistry.RegisterImplementations((*cryptotypes.PubKey)(nil), &injethsecp256k1.PubKey{})
	return cosmosCfg
}

func isNativeAsset(asset xc.AssetConfig) bool {
	return asset.Type == xc.AssetTypeNative || len(asset.Contract) < LEN_NATIVE_ASSET
}

func isEVMOS(asset xc.AssetConfig) bool {
	return xc.Driver(asset.Driver) == xc.DriverCosmosEvmos
}

func isINJ(asset xc.AssetConfig) bool {
	return asset.NativeAsset == xc.NativeAsset("INJ")
}

func getPublicKey(asset xc.AssetConfig, publicKeyBytes xc.PublicKey) cryptotypes.PubKey {
	if isINJ(asset) {
		return &injethsecp256k1.PubKey{Key: publicKeyBytes}
	}
	if isEVMOS(asset) {
		return &ethsecp256k1.PubKey{Key: publicKeyBytes}
	}
	return &secp256k1.PubKey{Key: publicKeyBytes}
}

func getSighash(asset xc.AssetConfig, sigData []byte) []byte {
	if isEVMOS(asset) || isINJ(asset) {
		return crypto.Keccak256(sigData)
	}
	sighash := sha256.Sum256(sigData)
	return sighash[:]
}
