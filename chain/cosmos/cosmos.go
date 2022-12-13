package cosmos

import (
	"crypto/sha256"

	xc "github.com/jumpcrypto/crosschain"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/evmos/ethermint/crypto/ethsecp256k1"
)

func isNativeAsset(asset xc.AssetConfig) bool {
	return asset.Type == xc.AssetTypeNative || len(asset.Contract) < 8
}

func isEVMOS(asset xc.AssetConfig) bool {
	return xc.Driver(asset.Driver) == xc.DriverCosmosEvmos
}

func getPublicKey(asset xc.AssetConfig, publicKeyBytes xc.PublicKey) cryptotypes.PubKey {
	if isEVMOS(asset) {
		return &ethsecp256k1.PubKey{Key: publicKeyBytes}
	}
	return &secp256k1.PubKey{Key: publicKeyBytes}
}

func getSighash(asset xc.AssetConfig, sigData []byte) []byte {
	if isEVMOS(asset) {
		return crypto.Keccak256(sigData)
	}
	sighash := sha256.Sum256(sigData)
	return sighash[:]
}
