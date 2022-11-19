package newchain

import (
	"context"
	"testing"

	xc "github.com/jumpcrypto/crosschain"
	"github.com/stretchr/testify/suite"
)

type CrosschainTestSuite struct {
	suite.Suite
	Ctx context.Context
}

func (s *CrosschainTestSuite) SetupTest() {
	s.Ctx = context.Background()
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(CrosschainTestSuite))
}

// Address

func (s *CrosschainTestSuite) TestNewAddressBuilder() {
	require := s.Require()
	builder, err := NewAddressBuilder(xc.AssetConfig{})
	require.NotNil(builder)
	require.EqualError(err, "not implemented")
}

func (s *CrosschainTestSuite) TestGetAddressFromPublicKey() {
	require := s.Require()
	builder, _ := NewAddressBuilder(xc.AssetConfig{})
	address, err := builder.GetAddressFromPublicKey([]byte{})
	require.Equal(xc.Address(""), address)
	require.EqualError(err, "not implemented")
}

func (s *CrosschainTestSuite) TestGetAllPossibleAddressesFromPublicKey() {
	require := s.Require()
	builder, _ := NewAddressBuilder(xc.AssetConfig{})
	addresses, err := builder.GetAllPossibleAddressesFromPublicKey([]byte{})
	require.NotNil(addresses)
	require.EqualError(err, "not implemented")
}

// TxBuilder

func (s *CrosschainTestSuite) TestNewTxBuilder() {
	require := s.Require()
	builder, err := NewTxBuilder(xc.AssetConfig{})
	require.NotNil(builder)
	require.EqualError(err, "not implemented")
}

func (s *CrosschainTestSuite) TestNewNativeTransfer() {
	require := s.Require()
	builder, _ := NewTxBuilder(xc.AssetConfig{})
	from := xc.Address("from")
	to := xc.Address("to")
	amount := xc.AmountBlockchain{}
	input := TxInput{}
	tf, err := builder.(xc.TxTokenBuilder).NewNativeTransfer(from, to, amount, input)
	require.Nil(tf)
	require.EqualError(err, "not implemented")
}

func (s *CrosschainTestSuite) TestNewTokenTransfer() {
	require := s.Require()
	builder, _ := NewTxBuilder(xc.AssetConfig{})
	from := xc.Address("from")
	to := xc.Address("to")
	amount := xc.AmountBlockchain{}
	input := TxInput{}
	tf, err := builder.(xc.TxTokenBuilder).NewTokenTransfer(from, to, amount, input)
	require.Nil(tf)
	require.EqualError(err, "not implemented")
}

func (s *CrosschainTestSuite) TestNewTransfer() {
	require := s.Require()
	builder, _ := NewTxBuilder(xc.AssetConfig{})
	from := xc.Address("from")
	to := xc.Address("to")
	amount := xc.AmountBlockchain{}
	input := TxInput{}
	tf, err := builder.NewTransfer(from, to, amount, input)
	require.Nil(tf)
	require.EqualError(err, "not implemented")
}

func (s *CrosschainTestSuite) TestNewTransfer_Token() {
	require := s.Require()
	builder, _ := NewTxBuilder(xc.AssetConfig{Type: xc.AssetTypeToken})
	from := xc.Address("from")
	to := xc.Address("to")
	amount := xc.AmountBlockchain{}
	input := TxInput{}
	tf, err := builder.NewTransfer(from, to, amount, input)
	require.Nil(tf)
	require.EqualError(err, "not implemented")
}

// Client

func (s *CrosschainTestSuite) TestNewClient() {
	require := s.Require()
	client, err := NewClient(xc.AssetConfig{})
	require.NotNil(client)
	require.EqualError(err, "not implemented")
}

func (s *CrosschainTestSuite) TestFetchTxInput() {
	require := s.Require()
	client, _ := NewClient(xc.AssetConfig{})
	from := xc.Address("from")
	input, err := client.FetchTxInput(s.Ctx, from)
	require.NotNil(input)
	require.EqualError(err, "not implemented")
}

func (s *CrosschainTestSuite) TestSubmitTx() {
	require := s.Require()
	client, _ := NewClient(xc.AssetConfig{})
	err := client.SubmitTx(s.Ctx, Tx{})
	require.EqualError(err, "not implemented")
}

func (s *CrosschainTestSuite) TestFetchTxInfo() {
	require := s.Require()
	client, _ := NewClient(xc.AssetConfig{})
	info, err := client.FetchTxInfo(s.Ctx, xc.TxHash("hash"))
	require.NotNil(info)
	require.EqualError(err, "not implemented")
}

// Signer

func (s *CrosschainTestSuite) TestNewSigner() {
	require := s.Require()
	signer, err := NewSigner(xc.AssetConfig{})
	require.NotNil(signer)
	require.EqualError(err, "not implemented")
}

func (s *CrosschainTestSuite) TestImportPrivateKey() {
	require := s.Require()
	signer, _ := NewSigner(xc.AssetConfig{})
	key, err := signer.ImportPrivateKey("key")
	require.NotNil(key)
	require.EqualError(err, "not implemented")
}

func (s *CrosschainTestSuite) TestSign() {
	require := s.Require()
	signer, _ := NewSigner(xc.AssetConfig{})
	sig, err := signer.Sign(xc.PrivateKey{}, xc.TxDataToSign{})
	require.NotNil(sig)
	require.EqualError(err, "not implemented")
}

// Tx

func (s *CrosschainTestSuite) TestTxHash() {
	require := s.Require()
	tx := Tx{}
	require.Equal(xc.TxHash("not implemented"), tx.Hash())
}

func (s *CrosschainTestSuite) TestTxSighash() {
	require := s.Require()
	tx := Tx{}
	sighash, err := tx.Sighash()
	require.NotNil(sighash)
	require.EqualError(err, "not implemented")
}

func (s *CrosschainTestSuite) TestTxAddSignature() {
	require := s.Require()
	tx := Tx{}
	err := tx.AddSignature(xc.TxSignature{})
	require.EqualError(err, "not implemented")
}
