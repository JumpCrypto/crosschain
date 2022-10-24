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

func (s *CrosschainTestSuite) TestAmountHumanReadable() {
	require := s.Require()
	amountDec, _ := decimal.NewFromString("10.3")
	amount := AmountHumanReadable(amountDec)
	require.NotNil(amount)
	require.Equal(amount.String(), "10.3")
}
