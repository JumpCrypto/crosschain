package evm

import (
	xc "github.com/jumpcrypto/crosschain"
)

func (s *CrosschainTestSuite) TestNewLegacyClient() {
	require := s.Require()
	client, err := NewLegacyClient(xc.AssetConfig{})
	require.NotNil(client)
	require.EqualError(err, "not implemented")
}
