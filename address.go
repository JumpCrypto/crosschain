package crosschain

// Address is an address on the blockchain, either sender or recipient
type Address string

// ContractAddress is a smart contract address
type ContractAddress Address

// AddressBuilder is the interface for building addresses
type AddressBuilder interface {
	GetAddressFromPublicKey(publicKeyBytes []byte) (Address, error)
	GetAllPossibleAddressesFromPublicKey(publicKeyBytes []byte) ([]PossibleAddress, error)
}

// AddressType represents the type of an address, for discovery purposes
type AddressType string

// List of known AddressType
const (
	AddressTypeSegwit    AddressType = AddressType("Segwit")
	AddressTypeP2SH      AddressType = AddressType("P2SH")
	AddressTypeP2PKH     AddressType = AddressType("P2PKH")
	AddressTypeP2WPKH    AddressType = AddressType("P2WPKH")
	AddressTypeP2TR      AddressType = AddressType("P2TR")
	AddressTypeETHKeccak AddressType = AddressType("ETHKeccak")
	AddressTypeDefault   AddressType = AddressType("Default")
)

// PossibleAddress is a pair of (Address, AddressType) used to derive all possible addresses from a public key
type PossibleAddress struct {
	Address Address
	Type    AddressType
}
