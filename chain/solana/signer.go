package solana

import (
	"crypto/ed25519"

	xc "github.com/jumpcrypto/crosschain"
)

// Signer for Solana
type Signer struct {
}

// NewSigner creates a new Solana Signer
func NewSigner(asset xc.AssetConfig) (xc.Signer, error) {
	return Signer{}, nil
}

// Sign a Solana tx
func (signer Signer) Sign(privateKey xc.PrivateKey, data xc.TxDataToSign) (xc.TxSignature, error) {
	signatureRaw := ed25519.Sign(ed25519.PrivateKey(privateKey), []byte(data))
	return xc.TxSignature(signatureRaw), nil
}
