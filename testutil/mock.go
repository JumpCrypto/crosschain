package testutil

import (
	xc "github.com/jumpcrypto/crosschain"
	"github.com/stretchr/testify/mock"
)

// MockedSigner returns a new mock for Signer
type MockedSigner struct {
	mock.Mock
}

// Sign a tx, mock
func (m *MockedSigner) Sign(data xc.TxDataToSign) (xc.TxSignature, error) {
	args := m.Called(data)
	return args.Get(0).(xc.TxSignature), args.Error(1)
}
