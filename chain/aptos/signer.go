package aptos

import (
	"errors"

	xc "github.com/jumpcrypto/crosschain"
)

// Signer for Aptos
type Signer struct {
}

// NewSigner creates a new Aptos Signer
func NewSigner(asset xc.AssetConfig) (xc.Signer, error) {
	return Signer{}, errors.New("not implemented")
}

// ImportPrivateKey imports an Aptos private key
func (signer Signer) ImportPrivateKey(privateKey string) (xc.PrivateKey, error) {
	return xc.PrivateKey([]byte{}), errors.New("not implemented")
}

// Sign an Aptos tx
func (signer Signer) Sign(privateKey xc.PrivateKey, data xc.TxDataToSign) (xc.TxSignature, error) {
	return xc.TxSignature([]byte{}), errors.New("not implemented")
}
