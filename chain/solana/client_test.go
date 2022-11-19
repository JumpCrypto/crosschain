package solana

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"

	"net/http/httptest"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/token"
	"github.com/gagliardetto/solana-go/rpc"
	xc "github.com/jumpcrypto/crosschain"
)

func wrapRPCResult(res string) string {
	return `{"jsonrpc":"2.0","result":` + res + `,"id":0}`
}

func wrapRPCError(err string) string {
	return `{"jsonrpc":"2.0","error":` + err + `,"id":0}`
}

type MockJSONRPCServer struct {
	*httptest.Server
	body    []byte
	counter int
}

func MockJSONRPC(s *CrosschainTestSuite, response interface{}) (mock *MockJSONRPCServer, close func()) {
	require := s.Require()
	mock = &MockJSONRPCServer{
		Server: httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			curResponse := response
			if a, ok := response.([]string); ok {
				curResponse = a[mock.counter]
			}
			mock.counter++

			// error => send RPC error
			if e, ok := curResponse.(error); ok {
				rw.WriteHeader(400)
				rw.Write([]byte(wrapRPCError(e.Error())))
				return
			}

			// string => convert to JSON
			if s, ok := curResponse.(string); ok {
				curResponse = json.RawMessage(wrapRPCResult(s))
			}

			var err error
			mock.body, err = ioutil.ReadAll(req.Body)
			log.Println("rpc>>", string(mock.body))
			require.NoError(err)

			// JSON input, or serializable into JSON
			var responseBody []byte
			if v, ok := curResponse.(json.RawMessage); ok {
				responseBody = v
			} else {
				responseBody, err = json.Marshal(curResponse)
				require.NoError(err)
			}
			log.Println("<<rpc", string(responseBody))
			rw.Write(responseBody)
		})),
	}
	return mock, func() { mock.Close() }
}

func (s *CrosschainTestSuite) TestNewClient() {
	require := s.Require()
	client, err := NewClient(xc.AssetConfig{})
	require.NotNil(client)
	require.Nil(err)
}

func (s *CrosschainTestSuite) TestFindAssociatedTokenAddress() {
	require := s.Require()

	ata, err := FindAssociatedTokenAddress("Hzn3n914JaSpnxo5mBbmuCDmGL6mxWN9Ac2HzEXFSGtb", "4zMMC9srt5Ri5X14GAgXhaHii3GnPAEERYPJgZJDncDU")
	require.Nil(err)
	require.Equal("DvSgNMRxVSMBpLp4hZeBrmQo8ZRFne72actTZ3PYE3AA", ata)

	ata, err = FindAssociatedTokenAddress("", "4zMMC9srt5Ri5X14GAgXhaHii3GnPAEERYPJgZJDncDU")
	require.ErrorContains(err, "zero length string")
	require.Equal("", ata)

	ata, err = FindAssociatedTokenAddress("Hzn3n914JaSpnxo5mBbmuCDmGL6mxWN9Ac2HzEXFSGtb", "xxx")
	require.ErrorContains(err, "invalid length")
	require.Equal("", ata)
}

/*
curl https://api.devnet.solana.com -X POST -H "Content-Type: application/json" -d '

	{
	  "jsonrpc": "2.0",
	  "id": 1,
	  "method": "getAccountInfo",
	  "params": [
	    "Hzn3n914JaSpnxo5mBbmuCDmGL6mxWN9Ac2HzEXFSGtb", {"encoding": "base64"}
	  ]
	}

'
*/
func (s *CrosschainTestSuite) TestFetchTxInput() {
	require := s.Require()

	vectors := []struct {
		asset           xc.AssetConfig
		resp            interface{}
		blockHash       string
		toIsATA         bool
		shouldCreateATA bool
		err             string
	}{
		{
			xc.AssetConfig{},
			[]string{
				// valid owner account
				`{"context":{"apiVersion":"1.13.3","slot":175635504},"value":{"data":["","base64"],"executable":false,"lamports":1860881440,"owner":"11111111111111111111111111111111","rentEpoch":371}}`,
				// valid blockhash
				`{"context":{"slot":83986105},"value":{"blockhash":"DvLEyV2GHk86K5GojpqnRsvhfMF5kdZomKMnhVpvHyqK","feeCalculator":{"lamportsPerSignature":5000}}}`,
			},
			"DvLEyV2GHk86K5GojpqnRsvhfMF5kdZomKMnhVpvHyqK",
			false,
			false,
			"",
		},
		{
			xc.AssetConfig{},
			[]string{
				// empty owner account
				`{"context":{"apiVersion":"1.13.3","slot":175636079},"value":null}`,
				// valid blockhash
				`{"context":{"slot":83986105},"value":{"blockhash":"DvLEyV2GHk86K5GojpqnRsvhfMF5kdZomKMnhVpvHyqK","feeCalculator":{"lamportsPerSignature":5000}}}`,
			},
			"DvLEyV2GHk86K5GojpqnRsvhfMF5kdZomKMnhVpvHyqK",
			false,
			false,
			"",
		},
		{
			xc.AssetConfig{},
			[]string{
				// valid ATA
				`{"context":{"apiVersion":"1.13.3","slot":175635873},"value":{"data":["O0Qss5EhV/E6kz0BNCgtAytf/s0Botvxt3kGCN8ALqctdvBIL2OuEzV5LBqS2x3308rEBwESq+xcukVQUYDkgpg6AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA","base64"],"executable":false,"lamports":2039280,"owner":"TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA","rentEpoch":0}}`,
				// valid blockhash
				`{"context":{"slot":83986105},"value":{"blockhash":"DvLEyV2GHk86K5GojpqnRsvhfMF5kdZomKMnhVpvHyqK","feeCalculator":{"lamportsPerSignature":5000}}}`,
			},
			"DvLEyV2GHk86K5GojpqnRsvhfMF5kdZomKMnhVpvHyqK",
			true,
			false,
			"",
		},
		{
			xc.AssetConfig{Type: xc.AssetTypeToken, Contract: "4zMMC9srt5Ri5X14GAgXhaHii3GnPAEERYPJgZJDncDU"},
			[]string{
				// valid owner account
				`{"context":{"apiVersion":"1.13.3","slot":175635504},"value":{"data":["","base64"],"executable":false,"lamports":1860881440,"owner":"11111111111111111111111111111111","rentEpoch":371}}`,
				// valid ATA
				`{"context":{"apiVersion":"1.13.3","slot":175635873},"value":{"data":["O0Qss5EhV/E6kz0BNCgtAytf/s0Botvxt3kGCN8ALqctdvBIL2OuEzV5LBqS2x3308rEBwESq+xcukVQUYDkgpg6AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA","base64"],"executable":false,"lamports":2039280,"owner":"TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA","rentEpoch":0}}`,
				// valid blockhash
				`{"context":{"slot":83986105},"value":{"blockhash":"DvLEyV2GHk86K5GojpqnRsvhfMF5kdZomKMnhVpvHyqK","feeCalculator":{"lamportsPerSignature":5000}}}`,
			},
			"DvLEyV2GHk86K5GojpqnRsvhfMF5kdZomKMnhVpvHyqK",
			false,
			false,
			"",
		},
		{
			xc.AssetConfig{Type: xc.AssetTypeToken, Contract: "4zMMC9srt5Ri5X14GAgXhaHii3GnPAEERYPJgZJDncDU"},
			[]string{
				// valid owner account
				`{"context":{"apiVersion":"1.13.3","slot":175635504},"value":{"data":["","base64"],"executable":false,"lamports":1860881440,"owner":"11111111111111111111111111111111","rentEpoch":371}}`,
				// empty ATA
				`{"context":{"apiVersion":"1.13.3","slot":175636079},"value":null}`,
				// valid blockhash
				`{"context":{"slot":83986105},"value":{"blockhash":"DvLEyV2GHk86K5GojpqnRsvhfMF5kdZomKMnhVpvHyqK","feeCalculator":{"lamportsPerSignature":5000}}}`,
			},
			"DvLEyV2GHk86K5GojpqnRsvhfMF5kdZomKMnhVpvHyqK",
			false,
			true,
			"",
		},
		{
			xc.AssetConfig{},
			[]string{
				// empty owner account
				`{"context":{"apiVersion":"1.13.3","slot":175636079},"value":null}`,
				// invalid blockhash
				`{"context":{"slot":83986105},"value":{"blockhash":"error","feeCalculator":{"lamportsPerSignature":5000}}}`,
			},
			"",
			false,
			false,
			"rpc.GetRecentBlockhashResult",
		},
		{
			xc.AssetConfig{},
			`null`,
			"",
			false,
			false,
			"error fetching blockhash",
		},
		{
			xc.AssetConfig{},
			`{}`,
			"",
			false,
			false,
			"error fetching blockhash",
		},
		{
			xc.AssetConfig{},
			errors.New(`{"message": "custom RPC error", "code": 123}`),
			"",
			false,
			false,
			"custom RPC error",
		},
	}

	for _, v := range vectors {
		server, close := MockJSONRPC(s, v.resp)
		defer close()

		v.asset.URL = server.URL
		client, _ := NewClient(v.asset)
		from := xc.Address("from")
		to := xc.Address("Hzn3n914JaSpnxo5mBbmuCDmGL6mxWN9Ac2HzEXFSGtb")
		input, err := client.FetchTxInput(s.Ctx, from, to)

		if v.err != "" {
			require.Nil(input)
			require.ErrorContains(err, v.err)
		} else {
			require.Nil(err)
			require.NotNil(input)
			require.Equal(v.toIsATA, input.(*TxInput).ToIsATA, "ToIsATA")
			require.Equal(v.shouldCreateATA, input.(*TxInput).ShouldCreateATA, "ShouldCreateATA")
			require.Equal(v.blockHash, input.(*TxInput).RecentBlockHash.String())
		}
	}
}

func (s *CrosschainTestSuite) TestSubmitTxErr() {
	require := s.Require()

	client, _ := NewClient(xc.AssetConfig{})
	tx := &Tx{
		SolTx:                  &solana.Transaction{},
		ParsedSolTx:            &rpc.ParsedTransaction{},
		associatedTokenAccount: &token.Account{},
		parsedTransfer:         nil,
	}
	err := client.SubmitTx(s.Ctx, tx)
	require.ErrorContains(err, "unsupported protocol scheme")
}

func (s *CrosschainTestSuite) TestAccountBalance() {
	require := s.Require()

	vectors := []struct {
		resp interface{}
		val  string
		err  string
	}{
		{
			`{"value": 123}`,
			"123",
			"",
		},
		{
			`null`,
			"0",
			"",
		},
		{
			`{}`,
			"0",
			"",
		},
		{
			errors.New(`{"message": "custom RPC error", "code": 123}`),
			"",
			"custom RPC error",
		},
	}

	for _, v := range vectors {
		server, close := MockJSONRPC(s, v.resp)
		defer close()

		client, _ := NewClient(xc.AssetConfig{URL: server.URL})
		from := xc.Address("Hzn3n914JaSpnxo5mBbmuCDmGL6mxWN9Ac2HzEXFSGtb")
		balance, err := client.AccountBalance(s.Ctx, from)

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

func (s *CrosschainTestSuite) TestTokenBalance() {
	require := s.Require()

	vectors := []struct {
		resp interface{}
		val  string
		err  string
	}{
		{
			`{"context":{"slot":1114},"value":{"amount":"9864","decimals":2,"uiAmount":98.64,"uiAmountString":"98.64"}}`,
			"9864",
			"",
		},
		{
			`null`,
			"0",
			"",
		},
		{
			`{}`,
			"0",
			"",
		},
		{
			errors.New(`{"message": "custom RPC error", "code": 123}`),
			"",
			"custom RPC error",
		},
	}

	for _, v := range vectors {
		server, close := MockJSONRPC(s, v.resp)
		defer close()

		client, _ := NewClient(xc.AssetConfig{URL: server.URL})
		from := xc.Address("Hzn3n914JaSpnxo5mBbmuCDmGL6mxWN9Ac2HzEXFSGtb")
		balance, err := client.TokenBalance(s.Ctx, from, "4zMMC9srt5Ri5X14GAgXhaHii3GnPAEERYPJgZJDncDU")

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

func (s *CrosschainTestSuite) TestFetchTxInfo() {
	require := s.Require()

	vectors := []struct {
		tx   string
		resp interface{}
		val  xc.TxInfo
		err  string
	}{
		{
			// 1 SOL
			"5U2YvvKUS6NUrDAJnABHjx2szwLCVmg8LCRK9BDbZwVAbf2q5j8D9Sc9kUoqanoqpn6ZpDguY3rip9W7N7vwCjSw",
			[]string{
				`{"blockTime":1650017168,"meta":{"err":null,"fee":5000,"innerInstructions":[],"loadedAddresses":{"readonly":[],"writable":[]},"logMessages":["Program 11111111111111111111111111111111 invoke [1]","Program 11111111111111111111111111111111 success"],"postBalances":[19921026477997237,1869985000,1],"postTokenBalances":[],"preBalances":[19921027478002237,869985000,1],"preTokenBalances":[],"rewards":[],"status":{"Ok":null}},"slot":128184605,"transaction":["Ad9f9FfCzdIyQqsm7dCzCNeEmfKMbUPhhRScrNuIs12xcfF3nkjOIiTMgLm5zkbdgHWDGQaLCOrjSxTcLNBwqwABAAEDeXJtpS2Z1gsH6tc7L28L9gg8yFx3qU401pHXj4vK/sn8iAhjIZAIQGI1+kyPuyqG09p7Z2Lqw5MjsqHYxASkFAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAkyu+8VadWPShFvQQKPdmQ5srpSxowzCLu+orIeRxb2cBAgIAAQwCAAAAAMqaOwAAAAA=","base64"]}`,
				`{"context":{"slot":128184606},"value":{"blockhash":"DvLEyV2GHk86K5GojpqnRsvhfMF5kdZomKMnhVpvHyqK","feeCalculator":{"lamportsPerSignature":5000}}}`,
			},
			xc.TxInfo{
				TxID:            "5U2YvvKUS6NUrDAJnABHjx2szwLCVmg8LCRK9BDbZwVAbf2q5j8D9Sc9kUoqanoqpn6ZpDguY3rip9W7N7vwCjSw",
				ExplorerURL:     "/tx/5U2YvvKUS6NUrDAJnABHjx2szwLCVmg8LCRK9BDbZwVAbf2q5j8D9Sc9kUoqanoqpn6ZpDguY3rip9W7N7vwCjSw?cluster=",
				From:            "9B5XszUGdMaxCZ7uSQhPzdks5ZQSmWxrmzCSvtJ6Ns6g",
				To:              "Hzn3n914JaSpnxo5mBbmuCDmGL6mxWN9Ac2HzEXFSGtb",
				ToAlt:           "",
				ContractAddress: "",
				Amount:          xc.NewAmountBlockchainFromUint64(1000000000),
				Fee:             xc.NewAmountBlockchainFromUint64(5000),
				BlockIndex:      128184605,
				BlockTime:       1650017168,
				Confirmations:   1,
			},
			"",
		},
		{
			// 0.12 SOL
			"3XRGeupw3XacNQ4op3TQdWJsX3VvSnzQdjBvQDjGHaTCZs1eJzbuVn67RThFXEBSDBvoCXT5eX7rU1frQLni5AKb",
			[]string{
				`{"blockTime":1645123751,"meta":{"err":null,"fee":5000,"innerInstructions":[],"loadedAddresses":{"readonly":[],"writable":[]},"logMessages":["Program 11111111111111111111111111111111 invoke [1]","Program 11111111111111111111111111111111 success"],"postBalances":[879990000,1420000000,1],"postTokenBalances":[],"preBalances":[999995000,1300000000,1],"preTokenBalances":[],"rewards":[],"status":{"Ok":null}},"slot":115310825,"transaction":["AX5EBZa5UnMbHNgzEDz8dn1mcrTjLwLsLC3Ph3tMgQshAb2hEkbkkUQleXVJqmcTYmxnnw3jIXOjfR3lGvw8pQoBAAED/IgIYyGQCEBiNfpMj7sqhtPae2di6sOTI7Kh2MQEpBR3FzzGpO7sbgIIhX1XFeQKpFBxBTrVYewdaBjV/jf96AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAZ3rYIt4WDe4pwTzQI6YOAbSxt/Orf5UkTzqKqXN1KMoBAgIAAQwCAAAAAA4nBwAAAAA=","base64"]}`,
				`{"context":{"slot":115310827},"value":{"blockhash":"DvLEyV2GHk86K5GojpqnRsvhfMF5kdZomKMnhVpvHyqK","feeCalculator":{"lamportsPerSignature":5000}}}`,
			},
			xc.TxInfo{
				TxID:            "3XRGeupw3XacNQ4op3TQdWJsX3VvSnzQdjBvQDjGHaTCZs1eJzbuVn67RThFXEBSDBvoCXT5eX7rU1frQLni5AKb",
				ExplorerURL:     "/tx/3XRGeupw3XacNQ4op3TQdWJsX3VvSnzQdjBvQDjGHaTCZs1eJzbuVn67RThFXEBSDBvoCXT5eX7rU1frQLni5AKb?cluster=",
				From:            "Hzn3n914JaSpnxo5mBbmuCDmGL6mxWN9Ac2HzEXFSGtb",
				To:              "91t4uSdtBiftqsB24W2fRXFCXjUyc6xY3WMGFedAaTHh",
				ToAlt:           "",
				ContractAddress: "",
				Amount:          xc.NewAmountBlockchainFromUint64(120000000),
				Fee:             xc.NewAmountBlockchainFromUint64(5000),
				BlockIndex:      115310825,
				BlockTime:       1645123751,
				Confirmations:   2,
				Status:          0,
			},
			"",
		},
		{
			// 0.001 USDC
			"ZJaJTB5oLfPrzEsFE2cEa94KdNb6SGvqMgaLdtqoYFnaqo4zAncVPjkpDqPbVPv85S68zNcaTyYobDcPJuRfhrX",
			[]string{
				`{"blockTime":1645120351,"meta":{"err":null,"fee":5000,"innerInstructions":[{"index":0,"instructions":[{"accounts":[0,1],"data":"3Bxs4h24hBtQy9rw","programIdIndex":5},{"accounts":[1],"data":"9krTDU2LzCSUJuVZ","programIdIndex":5},{"accounts":[1],"data":"SYXsBSQy3GeifSEQSGvTbrPNposbSAiSoh1YA85wcvGKSnYg","programIdIndex":5},{"accounts":[1,4,3,7],"data":"2","programIdIndex":6}]}],"loadedAddresses":{"readonly":[],"writable":[]},"logMessages":["Program ATokenGPvbdGVxr1b2hvZbsiqW5xWH25efTNsLJA8knL invoke [1]","Program log: Transfer 2039280 lamports to the associated token account","Program 11111111111111111111111111111111 invoke [2]","Program 11111111111111111111111111111111 success","Program log: Allocate space for the associated token account","Program 11111111111111111111111111111111 invoke [2]","Program 11111111111111111111111111111111 success","Program log: Assign the associated token account to the SPL Token program","Program 11111111111111111111111111111111 invoke [2]","Program 11111111111111111111111111111111 success","Program log: Initialize the associated token account","Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA invoke [2]","Program log: Instruction: InitializeAccount","Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA consumed 3297 of 169352 compute units","Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA success","Program ATokenGPvbdGVxr1b2hvZbsiqW5xWH25efTNsLJA8knL consumed 34626 of 200000 compute units","Program ATokenGPvbdGVxr1b2hvZbsiqW5xWH25efTNsLJA8knL success","Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA invoke [1]","Program log: Instruction: TransferChecked","Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA consumed 3414 of 200000 compute units","Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA success"],"postBalances":[4995907578480,2039280,2002039280,0,1461600,1,953185920,1009200,898174080],"postTokenBalances":[{"accountIndex":1,"mint":"4zMMC9srt5Ri5X14GAgXhaHii3GnPAEERYPJgZJDncDU","owner":"Hzn3n914JaSpnxo5mBbmuCDmGL6mxWN9Ac2HzEXFSGtb","uiTokenAmount":{"amount":"1000000","decimals":6,"uiAmount":1.0,"uiAmountString":"1"}},{"accountIndex":2,"mint":"4zMMC9srt5Ri5X14GAgXhaHii3GnPAEERYPJgZJDncDU","owner":"HzcTrHjkEhjFTHEsC6Dsv8DXCh21WgujD4s5M15Sm94g","uiTokenAmount":{"amount":"9437986064320000","decimals":6,"uiAmount":9437986064.32,"uiAmountString":"9437986064.32"}}],"preBalances":[4995909622760,0,2002039280,0,1461600,1,953185920,1009200,898174080],"preTokenBalances":[{"accountIndex":2,"mint":"4zMMC9srt5Ri5X14GAgXhaHii3GnPAEERYPJgZJDncDU","owner":"HzcTrHjkEhjFTHEsC6Dsv8DXCh21WgujD4s5M15Sm94g","uiTokenAmount":{"amount":"9437986065320000","decimals":6,"uiAmount":9437986065.32,"uiAmountString":"9437986065.32"}}],"rewards":[],"status":{"Ok":null}},"slot":115302132,"transaction":["AWHQR1wAwmzbcOtoAVvDi4tkCQD/n6Pnks/038BYBd5o3oh+QZDKHR0Onl9j+AFGp5wziV4cS96gDbzm4RYebwsBAAYJ/H0x63TSPiBNuaFOZG+ZK2YJNAoqn9Gpp2i1PA5g++W//Q+h80WCx4fF9949OkDj+1D/UUOKp6XPufludArJk2as/hRuCozLqVIK567QNivd/o4dGEy7JdaJah9qp8KU/IgIYyGQCEBiNfpMj7sqhtPae2di6sOTI7Kh2MQEpBQ7RCyzkSFX8TqTPQE0KC0DK1/+zQGi2/G3eQYI3wAupwAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABt324ddloZPZy+FGzut5rBy0he1fWzeROoz1hX7/AKkGp9UXGSxcUSGMyUw9SvF/WNruCJuh/UTj29mKAAAAAIyXJY9OJInxuz0QKRSODYMLWhOZ2v8QhASOe9jb6fhZNT+YJRqOG5qSK+OHTJJvIkfUNScIwvHTlo4R2xjWx7cCCAcAAQMEBQYHAAYFAgQBAAAKDEBCDwAAAAAABg==","base64"]}`,
				`{"context":{"slot":115302135},"value":{"blockhash":"DvLEyV2GHk86K5GojpqnRsvhfMF5kdZomKMnhVpvHyqK","feeCalculator":{"lamportsPerSignature":5000}}}`,
				`{"context":{"apiVersion":"1.13.2","slot":169710435},"value":{"data":["O0Qss5EhV/E6kz0BNCgtAytf/s0Botvxt3kGCN8ALqf8iAhjIZAIQGI1+kyPuyqG09p7Z2Lqw5MjsqHYxASkFAA1DAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA","base64"],"executable":false,"lamports":2039280,"owner":"TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA","rentEpoch":371}}`,
			},
			xc.TxInfo{
				TxID:            "ZJaJTB5oLfPrzEsFE2cEa94KdNb6SGvqMgaLdtqoYFnaqo4zAncVPjkpDqPbVPv85S68zNcaTyYobDcPJuRfhrX",
				ExplorerURL:     "/tx/ZJaJTB5oLfPrzEsFE2cEa94KdNb6SGvqMgaLdtqoYFnaqo4zAncVPjkpDqPbVPv85S68zNcaTyYobDcPJuRfhrX?cluster=",
				From:            "HzcTrHjkEhjFTHEsC6Dsv8DXCh21WgujD4s5M15Sm94g",
				To:              "Hzn3n914JaSpnxo5mBbmuCDmGL6mxWN9Ac2HzEXFSGtb",
				ToAlt:           "DvSgNMRxVSMBpLp4hZeBrmQo8ZRFne72actTZ3PYE3AA",
				ContractAddress: "4zMMC9srt5Ri5X14GAgXhaHii3GnPAEERYPJgZJDncDU",
				Amount:          xc.NewAmountBlockchainFromUint64(1000000),
				Fee:             xc.NewAmountBlockchainFromUint64(5000),
				BlockIndex:      115302132,
				BlockTime:       1645120351,
				Confirmations:   3,
			},
			"",
		},
		{
			// 0.0002 USDC
			"5ZrG8iS4RxLXDRQEWkAoddWHzkS1fA1m6ppxaAekgGzskhcFqjkw1ZaFCsLorbhY5V4YUUkjE3SLY2JNLyVanxrM",
			[]string{
				`{"blockTime":1645121566,"meta":{"err":null,"fee":5000,"innerInstructions":[],"loadedAddresses":{"readonly":[],"writable":[]},"logMessages":["Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA invoke [1]","Program log: Instruction: TransferChecked","Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA consumed 3285 of 200000 compute units","Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA success"],"postBalances":[999995000,2039280,2039280,1461600,953185920],"postTokenBalances":[{"accountIndex":1,"mint":"4zMMC9srt5Ri5X14GAgXhaHii3GnPAEERYPJgZJDncDU","owner":"91t4uSdtBiftqsB24W2fRXFCXjUyc6xY3WMGFedAaTHh","uiTokenAmount":{"amount":"1200000","decimals":6,"uiAmount":1.2,"uiAmountString":"1.2"}},{"accountIndex":2,"mint":"4zMMC9srt5Ri5X14GAgXhaHii3GnPAEERYPJgZJDncDU","owner":"Hzn3n914JaSpnxo5mBbmuCDmGL6mxWN9Ac2HzEXFSGtb","uiTokenAmount":{"amount":"800000","decimals":6,"uiAmount":0.8,"uiAmountString":"0.8"}}],"preBalances":[1000000000,2039280,2039280,1461600,953185920],"preTokenBalances":[{"accountIndex":1,"mint":"4zMMC9srt5Ri5X14GAgXhaHii3GnPAEERYPJgZJDncDU","owner":"91t4uSdtBiftqsB24W2fRXFCXjUyc6xY3WMGFedAaTHh","uiTokenAmount":{"amount":"1000000","decimals":6,"uiAmount":1.0,"uiAmountString":"1"}},{"accountIndex":2,"mint":"4zMMC9srt5Ri5X14GAgXhaHii3GnPAEERYPJgZJDncDU","owner":"Hzn3n914JaSpnxo5mBbmuCDmGL6mxWN9Ac2HzEXFSGtb","uiTokenAmount":{"amount":"1000000","decimals":6,"uiAmount":1.0,"uiAmountString":"1"}}],"rewards":[],"status":{"Ok":null}},"slot":115305244,"transaction":["AeRlYJzP6QA39jotdCCkfUXkwSfdikREPsky500UJkzv9WjwqWJRE1AVBY7gaDVCoHUXzBdRP2HqHa+yk6AOpAIBAAIF/IgIYyGQCEBiNfpMj7sqhtPae2di6sOTI7Kh2MQEpBRSZ7LN9mT5H4I0p8JRr5XuoV6s6jPU4jO/AY/AlB3Scr/9D6HzRYLHh8X33j06QOP7UP9RQ4qnpc+5+W50CsmTO0Qss5EhV/E6kz0BNCgtAytf/s0Botvxt3kGCN8ALqcG3fbh12Whk9nL4UbO63msHLSF7V9bN5E6jPWFfv8AqSRIX/y7kCeqnCwBTRUU6oea3zqDrhoNbdKX6z3IVPmkAQQEAgMBAAoMQA0DAAAAAAAG","base64"]}`,
				`{"context":{"slot":115305248},"value":{"blockhash":"DvLEyV2GHk86K5GojpqnRsvhfMF5kdZomKMnhVpvHyqK","feeCalculator":{"lamportsPerSignature":5000}}}`,
				`{"context":{"apiVersion":"1.13.2","slot":169710192},"value":{"data":["O0Qss5EhV/E6kz0BNCgtAytf/s0Botvxt3kGCN8ALqd3FzzGpO7sbgIIhX1XFeQKpFBxBTrVYewdaBjV/jf96IBPEgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA","base64"],"executable":false,"lamports":2039280,"owner":"TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA","rentEpoch":371}}`,
			},
			xc.TxInfo{
				TxID:            "5ZrG8iS4RxLXDRQEWkAoddWHzkS1fA1m6ppxaAekgGzskhcFqjkw1ZaFCsLorbhY5V4YUUkjE3SLY2JNLyVanxrM",
				ExplorerURL:     "/tx/5ZrG8iS4RxLXDRQEWkAoddWHzkS1fA1m6ppxaAekgGzskhcFqjkw1ZaFCsLorbhY5V4YUUkjE3SLY2JNLyVanxrM?cluster=",
				From:            "Hzn3n914JaSpnxo5mBbmuCDmGL6mxWN9Ac2HzEXFSGtb",
				To:              "91t4uSdtBiftqsB24W2fRXFCXjUyc6xY3WMGFedAaTHh",
				ToAlt:           "6Yg9GttAiHjbHMoiomBuGBDULP7HxQyez45dEiR9CJqw",
				ContractAddress: "4zMMC9srt5Ri5X14GAgXhaHii3GnPAEERYPJgZJDncDU",
				Amount:          xc.NewAmountBlockchainFromUint64(200000),
				Fee:             xc.NewAmountBlockchainFromUint64(5000),
				BlockIndex:      115305244,
				BlockTime:       1645121566,
				Confirmations:   4,
			},
			"",
		},
		{
			"5U2YvvKUS6NUrDAJnABHjx2szwLCVmg8LCRK9BDbZwVAbf2q5j8D9Sc9kUoqanoqpn6ZpDguY3rip9W7N7vwCjSw",
			`{}`,
			xc.TxInfo{},
			"invalid transaction in response",
		},
		{
			"5U2YvvKUS6NUrDAJnABHjx2szwLCVmg8LCRK9BDbZwVAbf2q5j8D9Sc9kUoqanoqpn6ZpDguY3rip9W7N7vwCjSw",
			`null`,
			xc.TxInfo{},
			"not found",
		},
		{
			"5U2YvvKUS6NUrDAJnABHjx2szwLCVmg8LCRK9BDbZwVAbf2q5j8D9Sc9kUoqanoqpn6ZpDguY3rip9W7N7vwCjSw",
			errors.New(`{"message": "custom RPC error", "code": 123}`),
			xc.TxInfo{},
			"custom RPC error",
		},
	}

	for _, v := range vectors {
		server, close := MockJSONRPC(s, v.resp)
		defer close()

		client, _ := NewClient(xc.AssetConfig{URL: server.URL})
		txInfo, err := client.FetchTxInfo(s.Ctx, xc.TxHash(v.tx))

		if v.err != "" {
			require.Equal(xc.TxInfo{}, txInfo)
			require.ErrorContains(err, v.err)
		} else {
			require.Nil(err)
			require.NotNil(txInfo)
			require.Equal(v.val, txInfo)
		}
	}
}
