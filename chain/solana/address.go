package solana

import (
	"fmt"

	xc "github.com/jumpcrypto/crosschain"
	"github.com/mr-tron/base58"
)

// AddressBuilder for Solana
type AddressBuilder struct {
}

// NewAddressBuilder creates a new Solana AddressBuilder
func NewAddressBuilder(asset xc.AssetConfig) (xc.AddressBuilder, error) {
	return AddressBuilder{}, nil
}

// GetAddressFromPublicKey returns an Address given a public key
func (ab AddressBuilder) GetAddressFromPublicKey(publicKeyBytes []byte) (xc.Address, error) {
	if len(publicKeyBytes) != 32 {
		return xc.Address(""), fmt.Errorf("expected address length 32, got address length %v", len(publicKeyBytes))
	}
	return xc.Address(base58.Encode(publicKeyBytes)), nil
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
