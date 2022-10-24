package solana

import (
	xc "github.com/jumpcrypto/crosschain"
)

func (s *CrosschainTestSuite) TestNewTxBuilder() {
	require := s.Require()
	builder, err := NewTxBuilder(xc.AssetConfig{Asset: "USDC"})
	require.Nil(err)
	require.NotNil(builder)
	require.Equal("USDC", builder.(TxBuilder).Asset.Asset)
}

func (s *CrosschainTestSuite) TestNewNativeTransfer() {
	require := s.Require()
	builder, _ := NewTxBuilder(xc.AssetConfig{})
	from := xc.Address("Hzn3n914JaSpnxo5mBbmuCDmGL6mxWN9Ac2HzEXFSGtb")
	to := xc.Address("BWbmXj5ckAaWCAtzMZ97qnJhBAKegoXtgNrv9BUpAB11")
	amount := xc.NewAmountBlockchainFromUint64(1200000) // 1.2 SOL
	input := TxInput{}
	tx, err := builder.(xc.TxTokenBuilder).NewNativeTransfer(from, to, amount, input)
	require.Nil(err)
	require.NotNil(tx)
	solTx := tx.(*Tx).SolTx
	require.Equal(0, len(solTx.Signatures))
	require.Equal(1, len(solTx.Message.Instructions))
	require.Equal(uint16(0x2), solTx.Message.Instructions[0].ProgramIDIndex) // system tx
}

func (s *CrosschainTestSuite) TestNewNativeTransferErr() {
	require := s.Require()
	builder, _ := NewTxBuilder(xc.AssetConfig{})

	from := xc.Address("from") // fails on parsing from
	to := xc.Address("to")
	amount := xc.AmountBlockchain{}
	input := TxInput{}
	tx, err := builder.(xc.TxTokenBuilder).NewNativeTransfer(from, to, amount, input)
	require.Nil(tx)
	require.EqualError(err, "invalid length, expected 32, got 3")

	from = xc.Address("Hzn3n914JaSpnxo5mBbmuCDmGL6mxWN9Ac2HzEXFSGtb")
	// fails on parsing to
	tx, err = builder.(xc.TxTokenBuilder).NewNativeTransfer(from, to, amount, input)
	require.Nil(tx)
	require.EqualError(err, "invalid length, expected 32, got 2")
}

func (s *CrosschainTestSuite) TestNewTokenTransfer() {
	require := s.Require()
	builder, _ := NewTxBuilder(xc.AssetConfig{
		Type:     xc.AssetTypeToken,
		Contract: "4zMMC9srt5Ri5X14GAgXhaHii3GnPAEERYPJgZJDncDU",
		Decimals: 6,
	})
	from := xc.Address("Hzn3n914JaSpnxo5mBbmuCDmGL6mxWN9Ac2HzEXFSGtb")
	to := xc.Address("BWbmXj5ckAaWCAtzMZ97qnJhBAKegoXtgNrv9BUpAB11")
	amount := xc.NewAmountBlockchainFromUint64(1200000) // 1.2 SOL
	input := TxInput{}
	tx, err := builder.(xc.TxTokenBuilder).NewTokenTransfer(from, to, amount, input)
	require.Nil(err)
	require.NotNil(tx)
	solTx := tx.(*Tx).SolTx
	require.Equal(0, len(solTx.Signatures))
	require.Equal(1, len(solTx.Message.Instructions))
	require.Equal(uint16(0x4), solTx.Message.Instructions[0].ProgramIDIndex) // token tx
}

func (s *CrosschainTestSuite) TestNewTokenTransferErr() {
	require := s.Require()

	// invalid asset
	builder, _ := NewTxBuilder(xc.AssetConfig{})
	from := xc.Address("from")
	to := xc.Address("to")
	amount := xc.AmountBlockchain{}
	input := TxInput{}
	tx, err := builder.(xc.TxTokenBuilder).NewTokenTransfer(from, to, amount, input)
	require.Nil(tx)
	require.EqualError(err, "asset is not of type token")

	// invalid from, to
	builder, _ = NewTxBuilder(xc.AssetConfig{
		Type:     xc.AssetTypeToken,
		Contract: "4zMMC9srt5Ri5X14GAgXhaHii3GnPAEERYPJgZJDncDU",
		Decimals: 6,
	})
	from = xc.Address("from")
	to = xc.Address("to")
	amount = xc.AmountBlockchain{}
	input = TxInput{}
	tx, err = builder.(xc.TxTokenBuilder).NewTokenTransfer(from, to, amount, input)
	require.Nil(tx)
	require.EqualError(err, "invalid length, expected 32, got 3")

	from = xc.Address("Hzn3n914JaSpnxo5mBbmuCDmGL6mxWN9Ac2HzEXFSGtb")
	tx, err = builder.(xc.TxTokenBuilder).NewTokenTransfer(from, to, amount, input)
	require.Nil(tx)
	require.EqualError(err, "invalid length, expected 32, got 2")

	// invalid asset config
	builder, _ = NewTxBuilder(xc.AssetConfig{
		Type:     xc.AssetTypeToken,
		Contract: "contract",
		Decimals: 6,
	})
	tx, err = builder.(xc.TxTokenBuilder).NewTokenTransfer(from, to, amount, input)
	require.Nil(tx)
	require.EqualError(err, "invalid length, expected 32, got 6")
}

func (s *CrosschainTestSuite) TestNewTransfer() {
	require := s.Require()
	builder, _ := NewTxBuilder(xc.AssetConfig{})
	from := xc.Address("Hzn3n914JaSpnxo5mBbmuCDmGL6mxWN9Ac2HzEXFSGtb")
	to := xc.Address("BWbmXj5ckAaWCAtzMZ97qnJhBAKegoXtgNrv9BUpAB11")
	amount := xc.NewAmountBlockchainFromUint64(1200000) // 1.2 SOL
	input := TxInput{}
	tx, err := builder.NewTransfer(from, to, amount, input)
	require.Nil(err)
	require.NotNil(tx)
	solTx := tx.(*Tx).SolTx
	require.Equal(0, len(solTx.Signatures))
	require.Equal(1, len(solTx.Message.Instructions))
	require.Equal(uint16(0x2), solTx.Message.Instructions[0].ProgramIDIndex) // system tx
}

func (s *CrosschainTestSuite) TestNewTransferAsToken() {
	require := s.Require()
	builder, _ := NewTxBuilder(xc.AssetConfig{
		Type:     xc.AssetTypeToken,
		Contract: "4zMMC9srt5Ri5X14GAgXhaHii3GnPAEERYPJgZJDncDU",
		Decimals: 6,
	})
	from := xc.Address("Hzn3n914JaSpnxo5mBbmuCDmGL6mxWN9Ac2HzEXFSGtb")
	to := xc.Address("BWbmXj5ckAaWCAtzMZ97qnJhBAKegoXtgNrv9BUpAB11")
	amount := xc.NewAmountBlockchainFromUint64(1200000) // 1.2 SOL
	input := TxInput{}
	tx, err := builder.NewTransfer(from, to, amount, input)
	require.Nil(err)
	require.NotNil(tx)
	solTx := tx.(*Tx).SolTx
	require.Equal(0, len(solTx.Signatures))
	require.Equal(1, len(solTx.Message.Instructions))
	require.Equal(uint16(0x4), solTx.Message.Instructions[0].ProgramIDIndex) // token tx
}
