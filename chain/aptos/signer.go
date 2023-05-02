package aptos

import (
	"crypto/ed25519"
	"encoding/hex"

	"github.com/coming-chat/go-aptos/aptosaccount"
	xc "github.com/jumpcrypto/crosschain"
)

// Signer for Aptos
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

// ImportPrivateKey imports an Aptos private key
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
		acc, err := aptosaccount.NewAccountWithMnemonic(privateKeyString)
		if err != nil {
			return nil, err
		}
		seed = acc.PrivateKey
	}

	// To generate the address:
	// privateKey := ed25519.NewKeyFromSeed(seed)
	// publicKey := privateKey.Public().(ed25519.PublicKey)
	// ed255Pubkey := append(publicKey, 0)
	// addrBytes := sha3.Sum256(ed255Pubkey)
	// address := "0x" + hex.EncodeToString(addrBytes[:][:transactionbuilder.ADDRESS_LENGTH])

	return seed, nil
}

// Sign an Aptos tx
func (signer Signer) Sign(privateKeyBz xc.PrivateKey, data xc.TxDataToSign) (xc.TxSignature, error) {
	privateKey := ed25519.NewKeyFromSeed(privateKeyBz)
	return ed25519.Sign(privateKey, data), nil
}
