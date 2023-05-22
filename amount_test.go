package crosschain

import (
	"github.com/shopspring/decimal"
)

func (s *CrosschainTestSuite) TestNewAmountBlockchainFromUint64() {
	require := s.Require()
	amount := NewAmountBlockchainFromUint64(123)
	require.NotNil(amount)
	require.Equal(amount.Uint64(), uint64(123))
	require.Equal(amount.String(), "123")
}

func (s *CrosschainTestSuite) TestNewAmountBlockchainFromFloat64() {
	require := s.Require()
	amount := NewAmountBlockchainToMaskFloat64(1.23)
	require.NotNil(amount)
	require.Equal(amount.Uint64(), uint64(1230000))
	require.Equal(amount.String(), "1230000")

	amountFloat := amount.UnmaskFloat64()
	require.Equal(amountFloat, 1.23)
}

func (s *CrosschainTestSuite) TestAmountHumanReadable() {
	require := s.Require()
	amountDec, _ := decimal.NewFromString("10.3")
	amount := AmountHumanReadable(amountDec)
	require.NotNil(amount)
	require.Equal(amount.String(), "10.3")
}

func (s *CrosschainTestSuite) TestNewAmountHumanReadableFromStr() {
	require := s.Require()
	amount := NewAmountHumanReadableFromStr("10.3")
	require.NotNil(amount)
	require.Equal(amount.String(), "10.3")

	amount = NewAmountHumanReadableFromStr("0")
	require.NotNil(amount)
	require.Equal(amount.String(), "0")

	amount = NewAmountHumanReadableFromStr("")
	require.NotNil(amount)
	require.Equal(amount.String(), "0")

	amount = NewAmountHumanReadableFromStr("invalid")
	require.NotNil(amount)
	require.Equal(amount.String(), "0")
}
