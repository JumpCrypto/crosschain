package cosmos

import (
	"context"
	"encoding/hex"
	"testing"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/evmos/ethermint/crypto/ethsecp256k1"
	"github.com/stretchr/testify/suite"

	xc "github.com/jumpcrypto/crosschain"
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

func (s *CrosschainTestSuite) TestIsNativeAsset() {
	require := s.Require()

	is := isNativeAsset(xc.AssetConfig{Type: xc.AssetTypeNative})
	require.True(is)

	is = isNativeAsset(xc.AssetConfig{Type: xc.AssetTypeToken, Contract: "uluna"})
	require.True(is)

	is = isNativeAsset(xc.AssetConfig{Type: xc.AssetTypeToken, Contract: "a-valid-long-contract"})
	require.False(is)

	// edge cases
	is = isNativeAsset(xc.AssetConfig{})
	require.True(is)

	is = isNativeAsset(xc.AssetConfig{Type: xc.AssetTypeToken, Contract: "uluna"})
	require.True(is)
}

func (s *CrosschainTestSuite) TestIsEVMOS() {
	require := s.Require()
	is := isEVMOS(xc.ETH)
	require.False(is)

	is = isEVMOS(xc.ATOM)
	require.False(is)

	is = isEVMOS(xc.LUNA)
	require.False(is)

	is = isEVMOS(xc.XPLA)
	require.True(is)
}

func (s *CrosschainTestSuite) TestGetPublicKey() {
	require := s.Require()

	pubKey := getPublicKey(xc.LUNA, []byte{})
	require.Exactly(&secp256k1.PubKey{Key: []byte{}}, pubKey)

	pubKey = getPublicKey(xc.XPLA, []byte{})
	require.Exactly(&ethsecp256k1.PubKey{Key: []byte{}}, pubKey)
}

func (s *CrosschainTestSuite) TestGetSighash() {
	require := s.Require()

	sighash := getSighash(xc.LUNA, []byte{})
	// echo -n '' | openssl dgst -sha256
	require.Exactly("e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855", hex.EncodeToString(sighash))

	sighash = getSighash(xc.XPLA, []byte{})
	require.Exactly("c5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470", hex.EncodeToString(sighash))
}
