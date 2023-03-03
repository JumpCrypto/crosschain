package testutil

import (
	"context"

	xc "github.com/jumpcrypto/crosschain"
	"github.com/stretchr/testify/mock"
)

// MockedClient returns a new mock for Client
type MockedClient struct {
	mock.Mock
}

var _ xc.ClientBalance = &MockedClient{}

// FetchTxInput fetches tx input, mocked
func (m *MockedClient) FetchTxInput(ctx context.Context, from xc.Address, to xc.Address) (xc.TxInput, error) {
	args := m.Called(ctx, from, to)
	return args.Get(0).(xc.TxInput), args.Error(1)
}

// FetchTxInfo fetches tx info, mocked
func (m *MockedClient) FetchTxInfo(ctx context.Context, txHash xc.TxHash) (xc.TxInfo, error) {
	args := m.Called(ctx, txHash)
	return args.Get(0).(xc.TxInfo), args.Error(1)
}

// SubmitTx submits a tx, mocked
func (m *MockedClient) SubmitTx(ctx context.Context, tx xc.Tx) error {
	args := m.Called(ctx, tx)
	return args.Error(0)
}

// FetchBalance fetches balance, mocked
func (m *MockedClient) FetchBalance(ctx context.Context, address xc.Address) (xc.AmountBlockchain, error) {
	args := m.Called(ctx, address)
	return args.Get(0).(xc.AmountBlockchain), args.Error(1)
}

// FetchNativeBalance fetches native asset balance, mocked
func (m *MockedClient) FetchNativeBalance(ctx context.Context, address xc.Address) (xc.AmountBlockchain, error) {
	args := m.Called(ctx, address)
	return args.Get(0).(xc.AmountBlockchain), args.Error(1)
}

func (m *MockedClient) FetchBalanceForAsset(ctx context.Context, address xc.Address, assetCfg xc.AssetConfig) (xc.AmountBlockchain, error) {
	args := m.Called(ctx, address)
	return args.Get(0).(xc.AmountBlockchain), args.Error(1)
}

func (m *MockedClient) UpdateAsset(assetCfg xc.AssetConfig) error {
	args := m.Called(assetCfg)
	return args.Error(1)
}
