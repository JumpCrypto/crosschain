package evm

import (
	"errors"

	xc "github.com/jumpcrypto/crosschain"
)

// Signer for EVM
type Signer struct {
}

// NewSigner creates a new EVM Signer
func NewSigner(asset xc.AssetConfig) (xc.Signer, error) {
	return Signer{}, errors.New("not implemented")
}

// Sign an EVM tx
func (signer Signer) Sign(privateKey xc.PrivateKey, data xc.TxDataToSign) (xc.TxSignature, error) {
	return xc.TxSignature([]byte{}), errors.New("not implemented")
}
