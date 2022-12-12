package crosschain

// PrivateKey is a private key or reference to private key
type PrivateKey []byte

// PublicKey is a public key
type PublicKey []byte

// Signer is signer that can sign tx
type Signer interface {
	ImportPrivateKey(privateKey string) (PrivateKey, error)
	Sign(privateKey PrivateKey, data TxDataToSign) (TxSignature, error)
}
