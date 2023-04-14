package evm

import xc "github.com/jumpcrypto/crosschain"

func (s *CrosschainTestSuite) TestTxHashEmpty() {
	require := s.Require()
	tx := Tx{}
	require.Equal(xc.TxHash(""), tx.Hash())
}

func (s *CrosschainTestSuite) TestTxSighashesEmpty() {
	require := s.Require()
	tx := Tx{}
	sighashes, err := tx.Sighashes()
	require.NotNil(sighashes)
	require.EqualError(err, "transaction not initialized")
}

func (s *CrosschainTestSuite) TestTxAddSignatureEmpty() {
	require := s.Require()
	tx := Tx{}
	err := tx.AddSignatures([]xc.TxSignature{}...)
	require.EqualError(err, "transaction not initialized")
}
