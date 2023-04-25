package sui

import (
	"encoding/hex"
	"errors"

	xc "github.com/jumpcrypto/crosschain"
	"golang.org/x/crypto/sha3"
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

func (ab AddressBuilder) GetAllPossibleAddressesFromPublicKey(publicKeyBytes []byte) ([]xc.PossibleAddress, error) {
	address, err := ab.GetAddressFromPublicKey(publicKeyBytes)
	return []xc.PossibleAddress{
		{
			Address: address,
			Type:    xc.AddressTypeDefault,
		},
	}, err
}

// Import account with mnemonic
// acc, err := account.NewAccountWithMnemonic(mnemonic)

// // Import account with private key
// privateKey, err := hex.DecodeString("4ec5a9eefc0bb86027a6f3ba718793c813505acc25ed09447caf6a069accdd4b")
// acc, err := account.NewAccount(privateKey)

// // Get private key, public key, address
// fmt.Printf("privateKey = %x\n", acc.PrivateKey[:32])
// fmt.Printf(" publicKey = %x\n", acc.PublicKey)
// fmt.Printf("   address = %v\n", acc.Address)

// // Sign data
// signedData := acc.Sign(data)
