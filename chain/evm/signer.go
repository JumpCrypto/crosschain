package evm

import (
	"encoding/hex"

	"github.com/ethereum/go-ethereum/crypto"
	xc "github.com/jumpcrypto/crosschain"
)

// Signer for EVM
type Signer struct {
}

// NewSigner creates a new EVM Signer
func NewSigner(asset xc.ITask) (xc.Signer, error) {
	return Signer{}, nil
}

var _ xc.Signer = &Signer{}

// ImportPrivateKey imports an EVM private key
func (signer Signer) ImportPrivateKey(privateKey string) (xc.PrivateKey, error) {
	bytesPri, err := hex.DecodeString(privateKey)
	return xc.PrivateKey(bytesPri), err
}

// Sign an EVM tx
func (signer Signer) Sign(privateKey xc.PrivateKey, data xc.TxDataToSign) (xc.TxSignature, error) {
	ecdsaKey, err := crypto.HexToECDSA(hex.EncodeToString(privateKey))
	if err != nil {
		return []byte{}, err
	}
	signatureRaw, err := crypto.Sign([]byte(data), ecdsaKey)
	return xc.TxSignature(signatureRaw), err
}
