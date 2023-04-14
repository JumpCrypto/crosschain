package evm

import (
	"encoding/hex"

	xc "github.com/jumpcrypto/crosschain"
)

func (s *CrosschainTestSuite) TestNewSigner() {
	require := s.Require()
	signer, err := NewSigner(&xc.AssetConfig{})
	require.Nil(err)
	require.NotNil(signer)
}

func (s *CrosschainTestSuite) TestImportPrivateKey() {
	require := s.Require()
	signer, _ := NewSigner(&xc.AssetConfig{})
	key, err := signer.ImportPrivateKey("289c28")
	require.Equal(key, xc.PrivateKey{0x28, 0x9c, 0x28})
	require.Nil(err)
}

func (s *CrosschainTestSuite) TestSign() {
	require := s.Require()

	// https://github.com/ethereum/go-ethereum/blob/v1.11.5/crypto/crypto_test.go
	vectors := []struct {
		pri string
		pub string
		msg string
		sig string
	}{
		{
			"289c2857d4598e37fb9647507e47a309d6133539bf21a8b9cb6df88fd5232032",
			"",
			"c5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470", // keccak256("")
			"b415397b439cc1eaab587a70717499b56b6cbe63037c241b2eaca2e833a6da097002b11c9611964e97212c82eab9613531f40e065d4d32e32ef31d68fedd977501",
		},
		{
			"289c2857d4598e37fb9647507e47a309d6133539bf21a8b9cb6df88fd5232032",
			"",
			"41b1a0649752af1b28b3dc29a1556eee781e4a4c3a1f7f53f90fa834de098c4d", // keccak256("foo")
			"d155e94305af7e07dd8c32873e5c03cb95c9e05960ef85be9c07f671da58c73718c19adc397a211aa9e87e519e2038c5a3b658618db335f74f800b8e0cfeef4401",
		},
	}

	for _, v := range vectors {
		signer, _ := NewSigner(&xc.AssetConfig{})
		bytesPri, _ := hex.DecodeString(v.pri)
		bytesMsg, _ := hex.DecodeString(v.msg)
		sig, err := signer.Sign(xc.PrivateKey(bytesPri), xc.TxDataToSign(bytesMsg))
		require.Nil(err)
		require.NotNil(sig)
		require.Equal(v.sig, hex.EncodeToString(sig))
	}
}

func (s *CrosschainTestSuite) TestSignErr() {
	require := s.Require()
	signer, _ := NewSigner(&xc.AssetConfig{})

	sig, err := signer.Sign(xc.PrivateKey{}, xc.TxDataToSign{})
	require.NotNil(err)
	require.Equal(sig, xc.TxSignature([]byte{}))
}
