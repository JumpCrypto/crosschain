package crosschain

// PrivateKey is a private key or reference to private key
type PrivateKey []byte

// Signer is signer that can sign tx
type Signer interface {
	Sign(privateKey PrivateKey, data TxDataToSign) (TxSignature, error)
}
