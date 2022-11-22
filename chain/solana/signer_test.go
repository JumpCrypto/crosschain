package solana

import (
	"encoding/hex"

	xc "github.com/jumpcrypto/crosschain"
)

func (s *CrosschainTestSuite) TestNewSigner() {
	require := s.Require()
	signer, err := NewSigner(xc.AssetConfig{})
	require.Nil(err)
	require.NotNil(signer)
}

func (s *CrosschainTestSuite) TestImportPrivateKey() {
	require := s.Require()
	signer, _ := NewSigner(xc.AssetConfig{})
	key, err := signer.ImportPrivateKey("key")
	require.Equal(key, xc.PrivateKey{0x2, 0x3d, 0xa6})
	require.Nil(err)
}

func (s *CrosschainTestSuite) TestSign() {
	require := s.Require()

	// https://ed25519.cr.yp.to/python/sign.input
	vectors := []struct {
		pri string
		pub string
		msg string
		sig string
	}{
		{
			"9d61b19deffd5a60ba844af492ec2cc44449c5697b326919703bac031cae7f60d75a980182b10ab7d54bfed3c964073a0ee172f3daa62325af021a68f707511a",
			"d75a980182b10ab7d54bfed3c964073a0ee172f3daa62325af021a68f707511a",
			"",
			"e5564300c360ac729086e2cc806e828a84877f1eb8e5d974d873e065224901555fb8821590a33bacc61e39701cf9b46bd25bf5f0595bbe24655141438e7a100b",
		},
		{
			"940c89fe40a81dafbdb2416d14ae469119869744410c3303bfaa0241dac57800a2eb8c0501e30bae0cf842d2bde8dec7386f6b7fc3981b8c57c9792bb94cf2dd",
			"a2eb8c0501e30bae0cf842d2bde8dec7386f6b7fc3981b8c57c9792bb94cf2dd",
			"b87d3813e03f58cf19fd0b6395",
			"d8bb64aad8c9955a115a793addd24f7f2b077648714f49c4694ec995b330d09d640df310f447fd7b6cb5c14f9fe9f490bcf8cfadbfd2169c8ac20d3b8af49a0c",
		},
	}

	for _, v := range vectors {
		signer, _ := NewSigner(xc.AssetConfig{})
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
	signer, _ := NewSigner(xc.AssetConfig{})
	require.Panics(func() {
		signer.Sign(xc.PrivateKey{}, xc.TxDataToSign{})
	})
}
