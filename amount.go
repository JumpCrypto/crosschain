package crosschain

import (
	"math/big"

	"github.com/shopspring/decimal"
)

// AmountBlockchain is a big integer amount as blockchain expects it for tx.
type AmountBlockchain big.Int

// AmountHumanReadable is a decimal amount as a human expects it for readability.
type AmountHumanReadable decimal.Decimal

func (amount AmountBlockchain) String() string {
	bigInt := big.Int(amount)
	return bigInt.String()
}

// Uint64 converts an AmountBlockchain into uint64
func (amount AmountBlockchain) Uint64() uint64 {
	bigInt := big.Int(amount)
	return bigInt.Uint64()
}

// Cmp compares amount with other, via the underlying big.Int.Cmp()
func (amount *AmountBlockchain) Cmp(other AmountBlockchain) int {
	return (*big.Int)(amount).Cmp((*big.Int)(&other))
}

// NewAmountBlockchainFromUint64 creates a new AmountBlockchain from a uint64
func NewAmountBlockchainFromUint64(u64 uint64) AmountBlockchain {
	bigInt := new(big.Int).SetUint64(u64)
	return AmountBlockchain(*bigInt)
}

// NewAmountBlockchainFromStr creates a new AmountBlockchain from a string
func NewAmountBlockchainFromStr(str string) AmountBlockchain {
	bigInt, _ := new(big.Int).SetString(str, 10)
	return AmountBlockchain(*bigInt)
}

func (amount AmountHumanReadable) String() string {
	return decimal.Decimal(amount).String()
}
