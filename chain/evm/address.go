package evm

import (
	"crypto/ecdsa"
	"errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	xc "github.com/jumpcrypto/crosschain"
)

// AddressBuilder for EVM
type AddressBuilder struct {
}

// NewAddressBuilder creates a new EVM AddressBuilder
func NewAddressBuilder(asset xc.ITask) (xc.AddressBuilder, error) {
	return AddressBuilder{}, nil
}

// GetAddressFromPublicKey returns an Address given a public key
func (ab AddressBuilder) GetAddressFromPublicKey(publicKeyBytes []byte) (xc.Address, error) {
	var publicKey *ecdsa.PublicKey
	var err error
	if len(publicKeyBytes) == 33 {
		publicKey, err = crypto.DecompressPubkey(publicKeyBytes)
		if err != nil {
			return xc.Address(""), errors.New("invalid k256 public key")
		}
	} else {
		publicKey, err = crypto.UnmarshalPubkey(publicKeyBytes)
		if err != nil {
			return xc.Address(""), err
		}
	}

	address := crypto.PubkeyToAddress(*publicKey).Hex()
	return xc.Address(address), nil
}

// GetAllPossibleAddressesFromPublicKey returns all PossubleAddress(es) given a public key
func (ab AddressBuilder) GetAllPossibleAddressesFromPublicKey(publicKeyBytes []byte) ([]xc.PossibleAddress, error) {
	address, err := ab.GetAddressFromPublicKey(publicKeyBytes)
	return []xc.PossibleAddress{
		{
			Address: address,
			Type:    xc.AddressTypeDefault,
		},
	}, err
}

// HexToAddress returns a go-ethereum Address decoded Crosschain address (hex string).
func HexToAddress(address xc.Address) (common.Address, error) {
	str := TrimPrefixes(string(address))

	// HexToAddress from go-ethereum doesn't handle any error case
	// We wrap it just in case we need to handle some errors in the future
	return common.HexToAddress(str), nil
}
