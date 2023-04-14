package aptos

import (
	"encoding/hex"
	"errors"

	xc "github.com/jumpcrypto/crosschain"
	"golang.org/x/crypto/sha3"
)

// AddressBuilder for Template
type AddressBuilder struct {
}

// NewAddressBuilder creates a new Template AddressBuilder
func NewAddressBuilder(asset xc.ITask) (xc.AddressBuilder, error) {
	return AddressBuilder{}, nil
}

// GetAddressFromPublicKey returns an Address given a public key
func (ab AddressBuilder) GetAddressFromPublicKey(publicKeyBytes []byte) (xc.Address, error) {
	if len(publicKeyBytes) == 32 {
		publicKeyBytes = append(publicKeyBytes, 0x00)
	}
	if len(publicKeyBytes) != 33 {
		return xc.Address(""), errors.New("invalid length for ed25519 public key")
	}
	// we only support ed25519 signature scheme (ends in 0)
	if publicKeyBytes[32] != 0 {
		return xc.Address(""), errors.New("invalid format for ed25519 public key")
	}
	authKey := sha3.Sum256(publicKeyBytes)
	address := "0x" + hex.EncodeToString(authKey[:])
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
