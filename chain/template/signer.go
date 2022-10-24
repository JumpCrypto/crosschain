package newchain

import (
	"errors"

	xc "github.com/jumpcrypto/crosschain"
)

// Signer for Template
type Signer struct {
}

// NewSigner creates a new Template Signer
func NewSigner(asset xc.AssetConfig) (xc.Signer, error) {
	return Signer{}, errors.New("not implemented")
}

// Sign a Template tx
func (signer Signer) Sign(privateKey xc.PrivateKey, data xc.TxDataToSign) (xc.TxSignature, error) {
	return xc.TxSignature([]byte{}), errors.New("not implemented")
}
