package sui

import (
	"encoding/hex"
	"errors"
	"fmt"

	xc "github.com/jumpcrypto/crosschain"
	"golang.org/x/crypto/blake2b"
)

const (
	ADDRESS_LENGTH             = 64
	SIGNATURE_SCHEME_ED25519   = 0
	SIGNATURE_SCHEME_SECP256k1 = 1
)

type AddressBuilder struct {
}

var _ xc.AddressBuilder = &AddressBuilder{}

// NewAddressBuilder creates a new Template AddressBuilder
func NewAddressBuilder(asset xc.ITask) (xc.AddressBuilder, error) {
	return AddressBuilder{}, nil
}

func (ab AddressBuilder) GetAddressFromPublicKey(publicKeyBytes []byte) (xc.Address, error) {
	if len(publicKeyBytes) == 32 {
		publicKeyBytes = append([]byte{SIGNATURE_SCHEME_ED25519}, publicKeyBytes...)
	}
	if len(publicKeyBytes) != 33 {
		return xc.Address(""), errors.New("invalid length for ed25519 sui public key")
	}
	// we only support ed25519 signature scheme (starts with 0)
	if publicKeyBytes[0] != SIGNATURE_SCHEME_ED25519 {
		fmt.Println("key: ", hex.EncodeToString(publicKeyBytes))
		return xc.Address(""), errors.New("invalid format for ed25519 sui public key")
	}

	addrBytes := blake2b.Sum256(publicKeyBytes)

	address := "0x" + hex.EncodeToString(addrBytes[:])[:ADDRESS_LENGTH]
	return xc.Address(address), nil
}

func (ab AddressBuilder) GetAllPossibleAddressesFromPublicKey(publicKeyBytes []byte) ([]xc.PossibleAddress, error) {
	address, err := ab.GetAddressFromPublicKey(publicKeyBytes)
	return []xc.PossibleAddress{
		{
			Address: address,
			Type:    xc.AddressTypeDefault,
		},
	}, err
}
