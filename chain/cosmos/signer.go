package cosmos

import (
	"encoding/hex"
	"strings"

	xc "github.com/jumpcrypto/crosschain"

	"github.com/btcsuite/btcd/btcec"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
)

// Signer for Cosmos
type Signer struct {
	Asset xc.AssetConfig
}

// NewSigner creates a new Cosmos Signer
func NewSigner(asset xc.AssetConfig) (xc.Signer, error) {
	return Signer{
		Asset: asset,
	}, nil
}

// ImportPrivateKey imports a Cosmos private key
func (signer Signer) ImportPrivateKey(privateKeyOrMnemonic string) (xc.PrivateKey, error) {
	keyHex := privateKeyOrMnemonic
	if strings.Contains(privateKeyOrMnemonic, " ") {
		hdPath := hd.CreateHDPath(signer.Asset.ChainCoinHDPath, 0, 0).String()
		kb := keyring.NewUnsafe(keyring.NewInMemory())
		_, err := kb.NewAccount("key", privateKeyOrMnemonic, keyring.DefaultBIP39Passphrase, hdPath, hd.Secp256k1)
		if err != nil {
			return nil, err
		}
		keyHex, _ = kb.UnsafeExportPrivKeyHex("key")
	}
	key, err := hex.DecodeString(keyHex)
	if err != nil {
		return nil, err
	}
	return xc.PrivateKey(key), nil
}

// Sign a Cosmos tx
func (signer Signer) Sign(privateKeyBytes xc.PrivateKey, data xc.TxDataToSign) (xc.TxSignature, error) {
	privateKey, _ := btcec.PrivKeyFromBytes(btcec.S256(), []byte(privateKeyBytes))
	signature, err := privateKey.Sign([]byte(data))
	if err != nil {
		return nil, err
	}
	return xc.TxSignature(serializeSig(signature)), nil
}
