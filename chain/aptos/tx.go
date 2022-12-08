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

// Sighashes returns the tx payload to sign, aka sighash
func (tx Tx) Sighashes() ([]xc.TxDataToSign, error) {
	return []xc.TxDataToSign{}, errors.New("not implemented")
}

// AddSignatures adds a signature to Tx
func (tx Tx) AddSignatures(...xc.TxSignature) error {
	return errors.New("not implemented")
}
