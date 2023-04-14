package newchain

import (
	"errors"

	xc "github.com/jumpcrypto/crosschain"
)

// Tx for Template
type Tx struct {
}

var _ xc.Tx = &Tx{}

// Hash returns the tx hash or id
func (tx Tx) Hash() xc.TxHash {
	return xc.TxHash("not implemented")
}

// Sighashes returns the tx payload to sign, aka sighash
func (tx Tx) Sighashes() ([]xc.TxDataToSign, error) {
	return []xc.TxDataToSign{}, errors.New("not implemented")

}

// AddSignatures adds a signature to Tx
func (tx *Tx) AddSignatures(...xc.TxSignature) error {
	return errors.New("not implemented")
}

// Serialize returns the serialized tx
func (tx Tx) Serialize() ([]byte, error) {
	return []byte{}, errors.New("not implemented")
}

// ParseTransfer parses a tx and extracts higher-level transfer information
func (tx *Tx) ParseTransfer() {
}

// From is the sender of a transfer
func (tx Tx) From() xc.Address {
	return xc.Address("")
}

// To is the account receiving a transfer
func (tx Tx) To() xc.Address {
	return xc.Address("")
}

// Amount returns the tx amount
func (tx Tx) Amount() xc.AmountBlockchain {
	return xc.NewAmountBlockchainFromUint64(0)
}

// ContractAddress returns the contract address for a token transfer
func (tx Tx) ContractAddress() xc.ContractAddress {
	return xc.ContractAddress("")
}

// Sources returns the sources of a Tx
func (tx Tx) Sources() []*xc.TxInfoEndpoint {
	sources := []*xc.TxInfoEndpoint{}
	return sources
}

// Destinations returns the destinations of a Tx
func (tx Tx) Destinations() []*xc.TxInfoEndpoint {
	destinations := []*xc.TxInfoEndpoint{}
	return destinations
}
