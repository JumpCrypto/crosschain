package aptos

import (
	"errors"

	xc "github.com/jumpcrypto/crosschain"
)

// Tx for Template
type Tx struct {
}

// Hash returns the tx hash or id
func (tx Tx) Hash() xc.TxHash {
	return xc.TxHash("not implemented")
}

// Sighash returns the tx payload to sign, aka sighash
func (tx Tx) Sighash() (xc.TxDataToSign, error) {
	return xc.TxDataToSign{}, errors.New("not implemented")
}

// AddSignature adds a signature to Tx
func (tx Tx) AddSignature(xc.TxSignature) error {
	return errors.New("not implemented")
}
