package evm

import (
	"errors"

	xc "github.com/jumpcrypto/crosschain"
	"github.com/jumpcrypto/crosschain/test"
)

func (s *CrosschainTestSuite) TestNewClient() {
	require := s.Require()
	client, err := NewClient(&xc.NativeAssetConfig{})
	require.NotNil(client)
	require.False(client.Legacy)
	require.Nil(err)
}

func (s *CrosschainTestSuite) TestNewLegacyClient() {
	require := s.Require()
	client, err := NewLegacyClient(&xc.NativeAssetConfig{})
	require.NotNil(client)
	require.True(client.Legacy)
	require.Nil(err)
}

func (s *CrosschainTestSuite) TestAccountBalance() {
	require := s.Require()

	vectors := []struct {
		resp interface{}
		val  string
		err  string
	}{
		{
			`"0x123"`,
			"291",
			"",
		},
		{
			`null`,
			"0",
			"cannot unmarshal non-string into Go value of type",
		},
		{
			`{}`,
			"0",
			"cannot unmarshal non-string into Go value of type",
		},
		{
			errors.New(`{"message": "custom RPC error", "code": 123}`),
			"",
			"custom RPC error",
		},
	}

	for _, v := range vectors {
		server, close := test.MockJSONRPC(&s.Suite, v.resp)
		defer close()

		client, _ := NewClient(&xc.NativeAssetConfig{URL: server.URL, Type: xc.AssetTypeNative})
		from := xc.Address("0x0eC9f48533bb2A03F53F341EF5cc1B057892B10B")
		balance, err := client.FetchBalance(s.Ctx, from)

		if v.err != "" {
			require.Equal("0", balance.String())
			require.ErrorContains(err, v.err)
		} else {
			require.Nil(err)
			require.NotNil(balance)
			require.Equal(v.val, balance.String())
		}
	}
}

// func (s *CrosschainTestSuite) TestFetchTxInput() {
// 	require := s.Require()
// 	client, _ := NewClient(xc.AssetConfig{})
// 	from := xc.Address("from")
// 	input, err := client.FetchTxInput(s.Ctx, from, "")
// 	require.NotNil(input)
// 	require.EqualError(err, "not implemented")
// }

func (s *CrosschainTestSuite) TestSubmitTx() {
	require := s.Require()
	server, close := test.MockJSONRPC(&s.Suite, `{}`)
	defer close()
	client, _ := NewClient(&xc.NativeAssetConfig{NativeAsset: xc.ETH, URL: server.URL})
	err := client.SubmitTx(s.Ctx, &test.MockXcTx{
		SerializedSignedTx: []byte{1, 2, 3, 4},
	})
	require.NoError(err)
}

// func (s *CrosschainTestSuite) TestFetchTxInfo() {
// 	require := s.Require()
// 	client, _ := NewClient(xc.AssetConfig{})
// 	info, err := client.FetchTxInfo(s.Ctx, xc.TxHash("hash"))
// 	require.NotNil(info)
// 	require.EqualError(err, "not implemented")
// }
