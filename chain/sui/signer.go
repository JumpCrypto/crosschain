package sui

import (
	"crypto/ed25519"
	"encoding/hex"
	"fmt"

	"github.com/coming-chat/go-sui/account"
	xc "github.com/jumpcrypto/crosschain"
	"golang.org/x/crypto/sha3"
)

// Signer for Sui
type Signer struct {
	PrivateKey ed25519.PrivateKey
	PublicKey  ed25519.PublicKey
	Address    string
}

var _ xc.Signer = &Signer{}

// NewSigner creates a new Aptos Signer
func NewSigner(asset xc.ITask) (xc.Signer, error) {
	return &Signer{}, nil
}

// ImportPrivateKey imports an Sui private key
// Private key may be hex (32 bytes / 64 characters) or a mnemonic.
func (signer Signer) ImportPrivateKey(privateKeyString string) (xc.PrivateKey, error) {
	var seed []byte
	var err error
	if len(privateKeyString) == 64 {
		seed, err = hex.DecodeString(privateKeyString)
		if err != nil {
			return []byte{}, err
		}
	} else {
		// mnemonic
		acc, err := account.NewAccountWithMnemonic(privateKeyString)
		if err != nil {
			return nil, err
		}
		seed = acc.KeyPair.PrivateKey()
	}
	privateKey := ed25519.NewKeyFromSeed(seed)
	publicKey := privateKey.Public().(ed25519.PublicKey)

	// indicate ed25519 sig scheme
	tmp := []byte{0}
	tmp = append(tmp, publicKey...)
	addrBytes := sha3.Sum256(tmp)
	// address length is 40
	address := "0x" + hex.EncodeToString(addrBytes[:])[:40]
	fmt.Println("Sui address = ", address)

	return seed, nil
}

// Sign an Aptos tx
func (signer Signer) Sign(privateKeyBz xc.PrivateKey, data xc.TxDataToSign) (xc.TxSignature, error) {
	privateKey := ed25519.NewKeyFromSeed(privateKeyBz)
	return ed25519.Sign(privateKey, data), nil
}
