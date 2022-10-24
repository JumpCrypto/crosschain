package cosmos

import (
	"errors"

	xc "github.com/jumpcrypto/crosschain"
)

// Signer for Cosmos
type Signer struct {
}

// NewSigner creates a new Cosmos Signer
func NewSigner(asset xc.AssetConfig) (xc.Signer, error) {
	return Signer{}, errors.New("not implemented")
}

// Sign a Cosmos tx
func (signer Signer) Sign(privateKey xc.PrivateKey, data xc.TxDataToSign) (xc.TxSignature, error) {
	return xc.TxSignature([]byte{}), errors.New("not implemented")
}
