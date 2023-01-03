package crosschain

// TxInput is input data to a tx. Depending on the blockchain it can include nonce, recent block hash, account id, ...
type TxInput interface {
}

// TxInputWithPublicKey is input data to a tx for chains that need to explicitly set the public key, e.g. Cosmos
type TxInputWithPublicKey interface {
	TxInput
	SetPublicKey(PublicKey) error
	SetPublicKeyFromStr(string) error
}

type TxInputEnvelope struct {
	Type Driver `json:"type"`
}

func NewTxInputEnvelope(envType Driver) *TxInputEnvelope {
	return &TxInputEnvelope{
		Type: envType,
	}
}

// TxStatus is the status of a tx on chain, currently success or failure.
type TxStatus uint8

// TxStatus values
const (
	TxStatusSuccess TxStatus = 0
	TxStatusFailure TxStatus = 1
)

// TxInfoEndpoint is a unified view of an endpoint (source or destination) in a TxInfo.
type TxInfoEndpoint struct {
	Address         Address
	ContractAddress ContractAddress
	Amount          AmountBlockchain
	AmountHuman     AmountHumanReadable
	NativeAsset     NativeAsset
	Asset           Asset
	AssetConfig     *AssetConfig
}

// TxInfo is a unified view of common tx info across multiple blockchains. Use it as an example to build your own.
type TxInfo struct {
	TxID            string
	ExplorerURL     string
	From            Address
	To              Address
	ToAlt           Address
	ContractAddress ContractAddress
	Amount          AmountBlockchain
	Fee             AmountBlockchain
	BlockIndex      int64
	BlockTime       int64
	Confirmations   int64
	Status          TxStatus
	Sources         []TxInfoEndpoint
	Destinations    []TxInfoEndpoint
}

// TxHash is a tx hash or id
type TxHash string

// TxDataToSign is the payload that Signer needs to sign, when "signing a tx". It's sometimes called a sighash.
type TxDataToSign []byte

// TxSignature is a tx signature
type TxSignature []byte

// NewTxSignatures creates a new array of TxSignature, useful to cast [][]byte into []TxSignature
func NewTxSignatures(data [][]byte) []TxSignature {
	ret := make([]TxSignature, len(data))
	for i, sig := range data {
		ret[i] = TxSignature(sig)
	}
	return ret
}

// Tx is a transaction
type Tx interface {
	Hash() TxHash
	Sighashes() ([]TxDataToSign, error)
	AddSignatures(...TxSignature) error
}
