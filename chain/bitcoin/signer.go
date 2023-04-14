package bitcoin

import (
	"encoding/hex"

	"github.com/btcsuite/btcutil"
	"github.com/ethereum/go-ethereum/crypto"
	xc "github.com/jumpcrypto/crosschain"
)

// Signer for Bitcoin
type Signer struct {
}

var _ xc.Signer = &Signer{}

// NewSigner creates a new Bitcoin Signer
func NewSigner(asset xc.ITask) (xc.Signer, error) {
	return &Signer{}, nil
}

// ImportPrivateKey imports a Bitcoin private key
// This private key should be either:
//   - a 32 byte hex string representing a k256 private key
//   - a WIF string representing a k256 private key
func (signer *Signer) ImportPrivateKey(privateKey string) (xc.PrivateKey, error) {
	bz, err := hex.DecodeString(privateKey)
	if err != nil || len(bz) != 32 {
		// try wif
		wif, err := btcutil.DecodeWIF(privateKey)
		if err != nil {
			return xc.PrivateKey{}, err
		}
		bz = wif.PrivKey.Serialize()
	}
	return bz, nil
}

// Sign a Bitcoin tx
func (signer *Signer) Sign(privateKeyBytes xc.PrivateKey, data xc.TxDataToSign) (xc.TxSignature, error) {
	ecdsaKey, err := crypto.HexToECDSA(hex.EncodeToString(privateKeyBytes))
	if err != nil {
		return []byte{}, err
	}
	signatureRaw, err := crypto.Sign([]byte(data), ecdsaKey)
	return xc.TxSignature(signatureRaw), err
}
