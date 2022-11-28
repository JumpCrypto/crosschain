package aptos

import (
	xc "github.com/jumpcrypto/crosschain"
	"github.com/jumpcrypto/crosschain/test"
)

func (s *CrosschainTestSuite) TestNewClient() {
	require := s.Require()
	resp := `{"chain_id":38,"epoch":"133","ledger_version":"13087045","oldest_ledger_version":"0","ledger_timestamp":"1669676013555573","node_role":"full_node","oldest_block_height":"0","block_height":"5435983","git_hash":"2c74a456298fcd520241a562119b6fe30abdaae2"}`
	server, close := test.MockHTTP(&s.Suite, resp)
	defer close()

	client, err := NewClient(xc.AssetConfig{URL: server.URL})
	require.NotNil(client)
	require.Nil(err)
}

func (s *CrosschainTestSuite) TestFetchTxInput() {
	// require := s.Require()

	// vectors := []struct {
	// 	asset           xc.AssetConfig
	// 	resp            interface{}
	// 	blockHash       string
	// 	toIsATA         bool
	// 	shouldCreateATA bool
	// 	err             string
	// }{
	// 	{
	// 		xc.AssetConfig{},
	// 		// valid blockhash
	// 		`{"context":{"slot":83986105},"value":{"blockhash":"DvLEyV2GHk86K5GojpqnRsvhfMF5kdZomKMnhVpvHyqK","feeCalculator":{"lamportsPerSignature":5000}}}`,
	// 		"DvLEyV2GHk86K5GojpqnRsvhfMF5kdZomKMnhVpvHyqK",
	// 		false,
	// 		false,
	// 		"",
	// 	},
	// 	{
	// 		xc.AssetConfig{Type: xc.AssetTypeToken, Contract: "4zMMC9srt5Ri5X14GAgXhaHii3GnPAEERYPJgZJDncDU"},
	// 		[]string{
	// 			// valid blockhash
	// 			`{"context":{"slot":83986105},"value":{"blockhash":"DvLEyV2GHk86K5GojpqnRsvhfMF5kdZomKMnhVpvHyqK","feeCalculator":{"lamportsPerSignature":5000}}}`,
	// 			// valid owner account
	// 			`{"context":{"apiVersion":"1.13.3","slot":175635504},"value":{"data":["","base64"],"executable":false,"lamports":1860881440,"owner":"11111111111111111111111111111111","rentEpoch":371}}`,
	// 			// valid ATA
	// 			`{"context":{"apiVersion":"1.13.3","slot":175635873},"value":{"data":["O0Qss5EhV/E6kz0BNCgtAytf/s0Botvxt3kGCN8ALqctdvBIL2OuEzV5LBqS2x3308rEBwESq+xcukVQUYDkgpg6AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA","base64"],"executable":false,"lamports":2039280,"owner":"TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA","rentEpoch":0}}`,
	// 		},
	// 		"DvLEyV2GHk86K5GojpqnRsvhfMF5kdZomKMnhVpvHyqK",
	// 		false,
	// 		false,
	// 		"",
	// 	},
	// 	{
	// 		xc.AssetConfig{Type: xc.AssetTypeToken, Contract: "4zMMC9srt5Ri5X14GAgXhaHii3GnPAEERYPJgZJDncDU"},
	// 		[]string{
	// 			// valid blockhash
	// 			`{"context":{"slot":83986105},"value":{"blockhash":"DvLEyV2GHk86K5GojpqnRsvhfMF5kdZomKMnhVpvHyqK","feeCalculator":{"lamportsPerSignature":5000}}}`,
	// 			// valid owner account
	// 			`{"context":{"apiVersion":"1.13.3","slot":175635504},"value":{"data":["","base64"],"executable":false,"lamports":1860881440,"owner":"11111111111111111111111111111111","rentEpoch":371}}`,
	// 			// empty ATA
	// 			`{"context":{"apiVersion":"1.13.3","slot":175636079},"value":null}`,
	// 		},
	// 		"DvLEyV2GHk86K5GojpqnRsvhfMF5kdZomKMnhVpvHyqK",
	// 		false,
	// 		true,
	// 		"",
	// 	},
	// 	{
	// 		xc.AssetConfig{Type: xc.AssetTypeToken, Contract: "4zMMC9srt5Ri5X14GAgXhaHii3GnPAEERYPJgZJDncDU"},
	// 		[]string{
	// 			// valid blockhash
	// 			`{"context":{"slot":83986105},"value":{"blockhash":"DvLEyV2GHk86K5GojpqnRsvhfMF5kdZomKMnhVpvHyqK","feeCalculator":{"lamportsPerSignature":5000}}}`,
	// 			// valid ATA
	// 			`{"context":{"apiVersion":"1.13.3","slot":175635873},"value":{"data":["O0Qss5EhV/E6kz0BNCgtAytf/s0Botvxt3kGCN8ALqctdvBIL2OuEzV5LBqS2x3308rEBwESq+xcukVQUYDkgpg6AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA","base64"],"executable":false,"lamports":2039280,"owner":"TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA","rentEpoch":0}}`,
	// 			// valid ATA
	// 			`{"context":{"apiVersion":"1.13.3","slot":175635873},"value":{"data":["O0Qss5EhV/E6kz0BNCgtAytf/s0Botvxt3kGCN8ALqctdvBIL2OuEzV5LBqS2x3308rEBwESq+xcukVQUYDkgpg6AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA","base64"],"executable":false,"lamports":2039280,"owner":"TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA","rentEpoch":0}}`,
	// 		},
	// 		"DvLEyV2GHk86K5GojpqnRsvhfMF5kdZomKMnhVpvHyqK",
	// 		true,
	// 		false,
	// 		"",
	// 	},
	// 	{
	// 		xc.AssetConfig{Type: xc.AssetTypeToken, Contract: "4zMMC9srt5Ri5X14GAgXhaHii3GnPAEERYPJgZJDncDU"},
	// 		[]string{
	// 			// valid blockhash
	// 			`{"context":{"slot":83986105},"value":{"blockhash":"DvLEyV2GHk86K5GojpqnRsvhfMF5kdZomKMnhVpvHyqK","feeCalculator":{"lamportsPerSignature":5000}}}`,
	// 			// empty ATA
	// 			`{"context":{"apiVersion":"1.13.3","slot":175636079},"value":null}`,
	// 			// empty ATA
	// 			`{"context":{"apiVersion":"1.13.3","slot":175636079},"value":null}`,
	// 		},
	// 		"DvLEyV2GHk86K5GojpqnRsvhfMF5kdZomKMnhVpvHyqK",
	// 		true,
	// 		true,
	// 		"",
	// 	},
	// 	{
	// 		xc.AssetConfig{},
	// 		[]string{
	// 			// invalid blockhash
	// 			`{"context":{"slot":83986105},"value":{"blockhash":"error","feeCalculator":{"lamportsPerSignature":5000}}}`,
	// 		},
	// 		"",
	// 		false,
	// 		false,
	// 		"rpc.GetRecentBlockhashResult",
	// 	},
	// 	{
	// 		xc.AssetConfig{Type: xc.AssetTypeToken, Contract: "invalid-contract"},
	// 		[]string{
	// 			// valid blockhash
	// 			`{"context":{"slot":83986105},"value":{"blockhash":"DvLEyV2GHk86K5GojpqnRsvhfMF5kdZomKMnhVpvHyqK","feeCalculator":{"lamportsPerSignature":5000}}}`,
	// 			// valid owner account
	// 			`{"context":{"apiVersion":"1.13.3","slot":175635504},"value":{"data":["","base64"],"executable":false,"lamports":1860881440,"owner":"11111111111111111111111111111111","rentEpoch":371}}`,
	// 		},
	// 		"DvLEyV2GHk86K5GojpqnRsvhfMF5kdZomKMnhVpvHyqK",
	// 		true,
	// 		true,
	// 		"decode: invalid base58 digit",
	// 	},
	// 	{
	// 		xc.AssetConfig{},
	// 		`null`,
	// 		"",
	// 		false,
	// 		false,
	// 		"error fetching blockhash",
	// 	},
	// 	{
	// 		xc.AssetConfig{},
	// 		`{}`,
	// 		"",
	// 		false,
	// 		false,
	// 		"error fetching blockhash",
	// 	},
	// 	{
	// 		xc.AssetConfig{},
	// 		errors.New(`{"message": "custom RPC error", "code": 123}`),
	// 		"",
	// 		false,
	// 		false,
	// 		"custom RPC error",
	// 	},
	// }

	// for _, v := range vectors {
	// 	server, close := test.MockJSONRPC(&s.Suite, v.resp)
	// 	defer close()

	// 	v.asset.URL = server.URL
	// 	client, _ := NewClient(v.asset)
	// 	from := xc.Address("from")
	// 	to := xc.Address("Hzn3n914JaSpnxo5mBbmuCDmGL6mxWN9Ac2HzEXFSGtb")
	// 	input, err := client.FetchTxInput(s.Ctx, from, to)

	// 	if v.err != "" {
	// 		require.Nil(input)
	// 		require.ErrorContains(err, v.err)
	// 	} else {
	// 		require.Nil(err)
	// 		require.NotNil(input)
	// 		require.Equal(v.toIsATA, input.(*TxInput).ToIsATA, "ToIsATA")
	// 		require.Equal(v.shouldCreateATA, input.(*TxInput).ShouldCreateATA, "ShouldCreateATA")
	// 		require.Equal(v.blockHash, input.(*TxInput).RecentBlockHash.String())
	// 	}
	// }
}

func (s *CrosschainTestSuite) TestSubmitTxErr() {
	// require := s.Require()

	// client, _ := NewClient(xc.AssetConfig{})
	// tx := &Tx{
	// 	SolTx:                  &solana.Transaction{},
	// 	ParsedSolTx:            &rpc.ParsedTransaction{},
	// 	associatedTokenAccount: &token.Account{},
	// 	parsedTransfer:         nil,
	// }
	// err := client.SubmitTx(s.Ctx, tx)
	// require.ErrorContains(err, "unsupported protocol scheme")
}

func (s *CrosschainTestSuite) TestFetchTxInfo() {
	// require := s.Require()

	// vectors := []struct {
	// 	tx   string
	// 	resp interface{}
	// 	val  xc.TxInfo
	// 	err  string
	// }{
	// 	{
	// 		// 1 SOL
	// 		"5U2YvvKUS6NUrDAJnABHjx2szwLCVmg8LCRK9BDbZwVAbf2q5j8D9Sc9kUoqanoqpn6ZpDguY3rip9W7N7vwCjSw",
	// 		[]string{
	// 			`{"blockTime":1650017168,"meta":{"err":null,"fee":5000,"innerInstructions":[],"loadedAddresses":{"readonly":[],"writable":[]},"logMessages":["Program 11111111111111111111111111111111 invoke [1]","Program 11111111111111111111111111111111 success"],"postBalances":[19921026477997237,1869985000,1],"postTokenBalances":[],"preBalances":[19921027478002237,869985000,1],"preTokenBalances":[],"rewards":[],"status":{"Ok":null}},"slot":128184605,"transaction":["Ad9f9FfCzdIyQqsm7dCzCNeEmfKMbUPhhRScrNuIs12xcfF3nkjOIiTMgLm5zkbdgHWDGQaLCOrjSxTcLNBwqwABAAEDeXJtpS2Z1gsH6tc7L28L9gg8yFx3qU401pHXj4vK/sn8iAhjIZAIQGI1+kyPuyqG09p7Z2Lqw5MjsqHYxASkFAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAkyu+8VadWPShFvQQKPdmQ5srpSxowzCLu+orIeRxb2cBAgIAAQwCAAAAAMqaOwAAAAA=","base64"]}`,
	// 			`{"context":{"slot":128184606},"value":{"blockhash":"DvLEyV2GHk86K5GojpqnRsvhfMF5kdZomKMnhVpvHyqK","feeCalculator":{"lamportsPerSignature":5000}}}`,
	// 		},
	// 		xc.TxInfo{
	// 			TxID:            "5U2YvvKUS6NUrDAJnABHjx2szwLCVmg8LCRK9BDbZwVAbf2q5j8D9Sc9kUoqanoqpn6ZpDguY3rip9W7N7vwCjSw",
	// 			ExplorerURL:     "/tx/5U2YvvKUS6NUrDAJnABHjx2szwLCVmg8LCRK9BDbZwVAbf2q5j8D9Sc9kUoqanoqpn6ZpDguY3rip9W7N7vwCjSw?cluster=",
	// 			From:            "9B5XszUGdMaxCZ7uSQhPzdks5ZQSmWxrmzCSvtJ6Ns6g",
	// 			To:              "Hzn3n914JaSpnxo5mBbmuCDmGL6mxWN9Ac2HzEXFSGtb",
	// 			ToAlt:           "",
	// 			ContractAddress: "",
	// 			Amount:          xc.NewAmountBlockchainFromUint64(1000000000),
	// 			Fee:             xc.NewAmountBlockchainFromUint64(5000),
	// 			BlockIndex:      128184605,
	// 			BlockTime:       1650017168,
	// 			Confirmations:   1,
	// 		},
	// 		"",
	// 	},
	// 	{
	// 		// 0.12 SOL
	// 		"3XRGeupw3XacNQ4op3TQdWJsX3VvSnzQdjBvQDjGHaTCZs1eJzbuVn67RThFXEBSDBvoCXT5eX7rU1frQLni5AKb",
	// 		[]string{
	// 			`{"blockTime":1645123751,"meta":{"err":null,"fee":5000,"innerInstructions":[],"loadedAddresses":{"readonly":[],"writable":[]},"logMessages":["Program 11111111111111111111111111111111 invoke [1]","Program 11111111111111111111111111111111 success"],"postBalances":[879990000,1420000000,1],"postTokenBalances":[],"preBalances":[999995000,1300000000,1],"preTokenBalances":[],"rewards":[],"status":{"Ok":null}},"slot":115310825,"transaction":["AX5EBZa5UnMbHNgzEDz8dn1mcrTjLwLsLC3Ph3tMgQshAb2hEkbkkUQleXVJqmcTYmxnnw3jIXOjfR3lGvw8pQoBAAED/IgIYyGQCEBiNfpMj7sqhtPae2di6sOTI7Kh2MQEpBR3FzzGpO7sbgIIhX1XFeQKpFBxBTrVYewdaBjV/jf96AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAZ3rYIt4WDe4pwTzQI6YOAbSxt/Orf5UkTzqKqXN1KMoBAgIAAQwCAAAAAA4nBwAAAAA=","base64"]}`,
	// 			`{"context":{"slot":115310827},"value":{"blockhash":"DvLEyV2GHk86K5GojpqnRsvhfMF5kdZomKMnhVpvHyqK","feeCalculator":{"lamportsPerSignature":5000}}}`,
	// 		},
	// 		xc.TxInfo{
	// 			TxID:            "3XRGeupw3XacNQ4op3TQdWJsX3VvSnzQdjBvQDjGHaTCZs1eJzbuVn67RThFXEBSDBvoCXT5eX7rU1frQLni5AKb",
	// 			ExplorerURL:     "/tx/3XRGeupw3XacNQ4op3TQdWJsX3VvSnzQdjBvQDjGHaTCZs1eJzbuVn67RThFXEBSDBvoCXT5eX7rU1frQLni5AKb?cluster=",
	// 			From:            "Hzn3n914JaSpnxo5mBbmuCDmGL6mxWN9Ac2HzEXFSGtb",
	// 			To:              "91t4uSdtBiftqsB24W2fRXFCXjUyc6xY3WMGFedAaTHh",
	// 			ToAlt:           "",
	// 			ContractAddress: "",
	// 			Amount:          xc.NewAmountBlockchainFromUint64(120000000),
	// 			Fee:             xc.NewAmountBlockchainFromUint64(5000),
	// 			BlockIndex:      115310825,
	// 			BlockTime:       1645123751,
	// 			Confirmations:   2,
	// 			Status:          0,
	// 		},
	// 		"",
	// 	},
	// 	{
	// 		// 0.001 USDC
	// 		"ZJaJTB5oLfPrzEsFE2cEa94KdNb6SGvqMgaLdtqoYFnaqo4zAncVPjkpDqPbVPv85S68zNcaTyYobDcPJuRfhrX",
	// 		[]string{
	// 			`{"blockTime":1645120351,"meta":{"err":null,"fee":5000,"innerInstructions":[{"index":0,"instructions":[{"accounts":[0,1],"data":"3Bxs4h24hBtQy9rw","programIdIndex":5},{"accounts":[1],"data":"9krTDU2LzCSUJuVZ","programIdIndex":5},{"accounts":[1],"data":"SYXsBSQy3GeifSEQSGvTbrPNposbSAiSoh1YA85wcvGKSnYg","programIdIndex":5},{"accounts":[1,4,3,7],"data":"2","programIdIndex":6}]}],"loadedAddresses":{"readonly":[],"writable":[]},"logMessages":["Program ATokenGPvbdGVxr1b2hvZbsiqW5xWH25efTNsLJA8knL invoke [1]","Program log: Transfer 2039280 lamports to the associated token account","Program 11111111111111111111111111111111 invoke [2]","Program 11111111111111111111111111111111 success","Program log: Allocate space for the associated token account","Program 11111111111111111111111111111111 invoke [2]","Program 11111111111111111111111111111111 success","Program log: Assign the associated token account to the SPL Token program","Program 11111111111111111111111111111111 invoke [2]","Program 11111111111111111111111111111111 success","Program log: Initialize the associated token account","Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA invoke [2]","Program log: Instruction: InitializeAccount","Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA consumed 3297 of 169352 compute units","Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA success","Program ATokenGPvbdGVxr1b2hvZbsiqW5xWH25efTNsLJA8knL consumed 34626 of 200000 compute units","Program ATokenGPvbdGVxr1b2hvZbsiqW5xWH25efTNsLJA8knL success","Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA invoke [1]","Program log: Instruction: TransferChecked","Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA consumed 3414 of 200000 compute units","Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA success"],"postBalances":[4995907578480,2039280,2002039280,0,1461600,1,953185920,1009200,898174080],"postTokenBalances":[{"accountIndex":1,"mint":"4zMMC9srt5Ri5X14GAgXhaHii3GnPAEERYPJgZJDncDU","owner":"Hzn3n914JaSpnxo5mBbmuCDmGL6mxWN9Ac2HzEXFSGtb","uiTokenAmount":{"amount":"1000000","decimals":6,"uiAmount":1.0,"uiAmountString":"1"}},{"accountIndex":2,"mint":"4zMMC9srt5Ri5X14GAgXhaHii3GnPAEERYPJgZJDncDU","owner":"HzcTrHjkEhjFTHEsC6Dsv8DXCh21WgujD4s5M15Sm94g","uiTokenAmount":{"amount":"9437986064320000","decimals":6,"uiAmount":9437986064.32,"uiAmountString":"9437986064.32"}}],"preBalances":[4995909622760,0,2002039280,0,1461600,1,953185920,1009200,898174080],"preTokenBalances":[{"accountIndex":2,"mint":"4zMMC9srt5Ri5X14GAgXhaHii3GnPAEERYPJgZJDncDU","owner":"HzcTrHjkEhjFTHEsC6Dsv8DXCh21WgujD4s5M15Sm94g","uiTokenAmount":{"amount":"9437986065320000","decimals":6,"uiAmount":9437986065.32,"uiAmountString":"9437986065.32"}}],"rewards":[],"status":{"Ok":null}},"slot":115302132,"transaction":["AWHQR1wAwmzbcOtoAVvDi4tkCQD/n6Pnks/038BYBd5o3oh+QZDKHR0Onl9j+AFGp5wziV4cS96gDbzm4RYebwsBAAYJ/H0x63TSPiBNuaFOZG+ZK2YJNAoqn9Gpp2i1PA5g++W//Q+h80WCx4fF9949OkDj+1D/UUOKp6XPufludArJk2as/hRuCozLqVIK567QNivd/o4dGEy7JdaJah9qp8KU/IgIYyGQCEBiNfpMj7sqhtPae2di6sOTI7Kh2MQEpBQ7RCyzkSFX8TqTPQE0KC0DK1/+zQGi2/G3eQYI3wAupwAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABt324ddloZPZy+FGzut5rBy0he1fWzeROoz1hX7/AKkGp9UXGSxcUSGMyUw9SvF/WNruCJuh/UTj29mKAAAAAIyXJY9OJInxuz0QKRSODYMLWhOZ2v8QhASOe9jb6fhZNT+YJRqOG5qSK+OHTJJvIkfUNScIwvHTlo4R2xjWx7cCCAcAAQMEBQYHAAYFAgQBAAAKDEBCDwAAAAAABg==","base64"]}`,
	// 			`{"context":{"slot":115302135},"value":{"blockhash":"DvLEyV2GHk86K5GojpqnRsvhfMF5kdZomKMnhVpvHyqK","feeCalculator":{"lamportsPerSignature":5000}}}`,
	// 			`{"context":{"apiVersion":"1.13.2","slot":169710435},"value":{"data":["O0Qss5EhV/E6kz0BNCgtAytf/s0Botvxt3kGCN8ALqf8iAhjIZAIQGI1+kyPuyqG09p7Z2Lqw5MjsqHYxASkFAA1DAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA","base64"],"executable":false,"lamports":2039280,"owner":"TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA","rentEpoch":371}}`,
	// 		},
	// 		xc.TxInfo{
	// 			TxID:            "ZJaJTB5oLfPrzEsFE2cEa94KdNb6SGvqMgaLdtqoYFnaqo4zAncVPjkpDqPbVPv85S68zNcaTyYobDcPJuRfhrX",
	// 			ExplorerURL:     "/tx/ZJaJTB5oLfPrzEsFE2cEa94KdNb6SGvqMgaLdtqoYFnaqo4zAncVPjkpDqPbVPv85S68zNcaTyYobDcPJuRfhrX?cluster=",
	// 			From:            "HzcTrHjkEhjFTHEsC6Dsv8DXCh21WgujD4s5M15Sm94g",
	// 			To:              "Hzn3n914JaSpnxo5mBbmuCDmGL6mxWN9Ac2HzEXFSGtb",
	// 			ToAlt:           "DvSgNMRxVSMBpLp4hZeBrmQo8ZRFne72actTZ3PYE3AA",
	// 			ContractAddress: "4zMMC9srt5Ri5X14GAgXhaHii3GnPAEERYPJgZJDncDU",
	// 			Amount:          xc.NewAmountBlockchainFromUint64(1000000),
	// 			Fee:             xc.NewAmountBlockchainFromUint64(5000),
	// 			BlockIndex:      115302132,
	// 			BlockTime:       1645120351,
	// 			Confirmations:   3,
	// 		},
	// 		"",
	// 	},
	// 	{
	// 		// 0.0002 USDC
	// 		"5ZrG8iS4RxLXDRQEWkAoddWHzkS1fA1m6ppxaAekgGzskhcFqjkw1ZaFCsLorbhY5V4YUUkjE3SLY2JNLyVanxrM",
	// 		[]string{
	// 			`{"blockTime":1645121566,"meta":{"err":null,"fee":5000,"innerInstructions":[],"loadedAddresses":{"readonly":[],"writable":[]},"logMessages":["Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA invoke [1]","Program log: Instruction: TransferChecked","Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA consumed 3285 of 200000 compute units","Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA success"],"postBalances":[999995000,2039280,2039280,1461600,953185920],"postTokenBalances":[{"accountIndex":1,"mint":"4zMMC9srt5Ri5X14GAgXhaHii3GnPAEERYPJgZJDncDU","owner":"91t4uSdtBiftqsB24W2fRXFCXjUyc6xY3WMGFedAaTHh","uiTokenAmount":{"amount":"1200000","decimals":6,"uiAmount":1.2,"uiAmountString":"1.2"}},{"accountIndex":2,"mint":"4zMMC9srt5Ri5X14GAgXhaHii3GnPAEERYPJgZJDncDU","owner":"Hzn3n914JaSpnxo5mBbmuCDmGL6mxWN9Ac2HzEXFSGtb","uiTokenAmount":{"amount":"800000","decimals":6,"uiAmount":0.8,"uiAmountString":"0.8"}}],"preBalances":[1000000000,2039280,2039280,1461600,953185920],"preTokenBalances":[{"accountIndex":1,"mint":"4zMMC9srt5Ri5X14GAgXhaHii3GnPAEERYPJgZJDncDU","owner":"91t4uSdtBiftqsB24W2fRXFCXjUyc6xY3WMGFedAaTHh","uiTokenAmount":{"amount":"1000000","decimals":6,"uiAmount":1.0,"uiAmountString":"1"}},{"accountIndex":2,"mint":"4zMMC9srt5Ri5X14GAgXhaHii3GnPAEERYPJgZJDncDU","owner":"Hzn3n914JaSpnxo5mBbmuCDmGL6mxWN9Ac2HzEXFSGtb","uiTokenAmount":{"amount":"1000000","decimals":6,"uiAmount":1.0,"uiAmountString":"1"}}],"rewards":[],"status":{"Ok":null}},"slot":115305244,"transaction":["AeRlYJzP6QA39jotdCCkfUXkwSfdikREPsky500UJkzv9WjwqWJRE1AVBY7gaDVCoHUXzBdRP2HqHa+yk6AOpAIBAAIF/IgIYyGQCEBiNfpMj7sqhtPae2di6sOTI7Kh2MQEpBRSZ7LN9mT5H4I0p8JRr5XuoV6s6jPU4jO/AY/AlB3Scr/9D6HzRYLHh8X33j06QOP7UP9RQ4qnpc+5+W50CsmTO0Qss5EhV/E6kz0BNCgtAytf/s0Botvxt3kGCN8ALqcG3fbh12Whk9nL4UbO63msHLSF7V9bN5E6jPWFfv8AqSRIX/y7kCeqnCwBTRUU6oea3zqDrhoNbdKX6z3IVPmkAQQEAgMBAAoMQA0DAAAAAAAG","base64"]}`,
	// 			`{"context":{"slot":115305248},"value":{"blockhash":"DvLEyV2GHk86K5GojpqnRsvhfMF5kdZomKMnhVpvHyqK","feeCalculator":{"lamportsPerSignature":5000}}}`,
	// 			`{"context":{"apiVersion":"1.13.2","slot":169710192},"value":{"data":["O0Qss5EhV/E6kz0BNCgtAytf/s0Botvxt3kGCN8ALqd3FzzGpO7sbgIIhX1XFeQKpFBxBTrVYewdaBjV/jf96IBPEgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA","base64"],"executable":false,"lamports":2039280,"owner":"TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA","rentEpoch":371}}`,
	// 		},
	// 		xc.TxInfo{
	// 			TxID:            "5ZrG8iS4RxLXDRQEWkAoddWHzkS1fA1m6ppxaAekgGzskhcFqjkw1ZaFCsLorbhY5V4YUUkjE3SLY2JNLyVanxrM",
	// 			ExplorerURL:     "/tx/5ZrG8iS4RxLXDRQEWkAoddWHzkS1fA1m6ppxaAekgGzskhcFqjkw1ZaFCsLorbhY5V4YUUkjE3SLY2JNLyVanxrM?cluster=",
	// 			From:            "Hzn3n914JaSpnxo5mBbmuCDmGL6mxWN9Ac2HzEXFSGtb",
	// 			To:              "91t4uSdtBiftqsB24W2fRXFCXjUyc6xY3WMGFedAaTHh",
	// 			ToAlt:           "6Yg9GttAiHjbHMoiomBuGBDULP7HxQyez45dEiR9CJqw",
	// 			ContractAddress: "4zMMC9srt5Ri5X14GAgXhaHii3GnPAEERYPJgZJDncDU",
	// 			Amount:          xc.NewAmountBlockchainFromUint64(200000),
	// 			Fee:             xc.NewAmountBlockchainFromUint64(5000),
	// 			BlockIndex:      115305244,
	// 			BlockTime:       1645121566,
	// 			Confirmations:   4,
	// 		},
	// 		"",
	// 	},
	// 	{
	// 		"5U2YvvKUS6NUrDAJnABHjx2szwLCVmg8LCRK9BDbZwVAbf2q5j8D9Sc9kUoqanoqpn6ZpDguY3rip9W7N7vwCjSw",
	// 		`{}`,
	// 		xc.TxInfo{},
	// 		"invalid transaction in response",
	// 	},
	// 	{
	// 		"5U2YvvKUS6NUrDAJnABHjx2szwLCVmg8LCRK9BDbZwVAbf2q5j8D9Sc9kUoqanoqpn6ZpDguY3rip9W7N7vwCjSw",
	// 		`null`,
	// 		xc.TxInfo{},
	// 		"not found",
	// 	},
	// 	{
	// 		"5U2YvvKUS6NUrDAJnABHjx2szwLCVmg8LCRK9BDbZwVAbf2q5j8D9Sc9kUoqanoqpn6ZpDguY3rip9W7N7vwCjSw",
	// 		errors.New(`{"message": "custom RPC error", "code": 123}`),
	// 		xc.TxInfo{},
	// 		"custom RPC error",
	// 	},
	// 	{
	// 		"",
	// 		"",
	// 		xc.TxInfo{},
	// 		"zero length string",
	// 	},
	// 	{
	// 		"invalid-sig",asset.go
	// 		"",
	// 		xc.TxInfo{},
	// 		"invalid base58 digit",
	// 	},
	// 	{
	// 		// 1 SOL
	// 		"5U2YvvKUS6NUrDAJnABHjx2szwLCVmg8LCRK9BDbZwVAbf2q5j8D9Sc9kUoqanoqpn6ZpDguY3rip9W7N7vwCjSw",
	// 		[]string{
	// 			`{"blockTime":1650017168,"meta":{"err":null,"fee":5000,"innerInstructions":[],"loadedAddresses":{"readonly":[],"writable":[]},"logMessages":["Program 11111111111111111111111111111111 invoke [1]","Program 11111111111111111111111111111111 success"],"postBalances":[19921026477997237,1869985000,1],"postTokenBalances":[],"preBalances":[19921027478002237,869985000,1],"preTokenBalances":[],"rewards":[],"status":{"Ok":null}},"slot":128184605,"transaction":["invalid-binary","base64"]}`,
	// 			`{"context":{"slot":128184606},"value":{"blockhash":"DvLEyV2GHk86K5GojpqnRsvhfMF5kdZomKMnhVpvHyqK","feeCalculator":{"lamportsPerSignature":5000}}}`,
	// 		},
	// 		xc.TxInfo{},
	// 		"illegal base64 data",
	// 	},
	// }

	// for _, v := range vectors {
	// 	server, close := test.MockJSONRPC(&s.Suite, v.resp)
	// 	defer close()

	// 	client, _ := NewClient(xc.AssetConfig{URL: server.URL})
	// 	txInfo, err := client.FetchTxInfo(s.Ctx, xc.TxHash(v.tx))

	// 	if v.err != "" {
	// 		require.Equal(xc.TxInfo{}, txInfo)
	// 		require.ErrorContains(err, v.err)
	// 	} else {
	// 		require.Nil(err)
	// 		require.NotNil(txInfo)
	// 		require.Equal(v.val, txInfo)
	// 	}
	// }
}

func (s *CrosschainTestSuite) TestFetchBalance() {
	require := s.Require()

	vectors := []struct {
		asset xc.AssetConfig
		resp  interface{}
		val   string
		err   string
	}{
		{
			xc.AssetConfig{Type: xc.AssetTypeNative},
			`{"type":"0x1::coin::CoinStore<0x1::aptos_coin::AptosCoin>","data":{"coin":{"value":"1000000"},"deposit_events":{"counter":"2","guid":{"id":{"addr":"0xa589a80d61ec380c24a5fdda109c3848c082584e6cb725e5ab19b18354b2ab85","creation_num":"2"}}},"frozen":false,"withdraw_events":{"counter":"0","guid":{"id":{"addr":"0xa589a80d61ec380c24a5fdda109c3848c082584e6cb725e5ab19b18354b2ab85","creation_num":"3"}}}}}`,
			"1000000",
			"",
		},
		{
			xc.AssetConfig{Type: xc.AssetTypeToken},
			`{"type":"0x1::coin::CoinStore<0x1::aptos_coin::AptosCoin>","data":{"coin":{"value":"1000000"},"deposit_events":{"counter":"2","guid":{"id":{"addr":"0xa589a80d61ec380c24a5fdda109c3848c082584e6cb725e5ab19b18354b2ab85","creation_num":"2"}}},"frozen":false,"withdraw_events":{"counter":"0","guid":{"id":{"addr":"0xa589a80d61ec380c24a5fdda109c3848c082584e6cb725e5ab19b18354b2ab85","creation_num":"3"}}}}}`,
			"1000000",
			"",
		},
		{
			xc.AssetConfig{Type: xc.AssetTypeNative},
			`null`,
			"0",
			"",
		},
		// {
		// 	xc.AssetConfig{Type: xc.AssetTypeNative},
		// 	`{}`,
		// 	"0",
		// 	"failed to get account balance",
		// },
		// {
		// 	xc.AssetConfig{Type: xc.AssetTypeNative},
		// 	errors.New(`{"message": "custom RPC error", "code": 123}`),
		// 	"",
		// 	"",
		// },
		// {
		// 	xc.AssetConfig{Type: xc.AssetTypeToken},
		// 	`{"message":"Resource not found by Address(0xa589a80d61ec380c24a5fdda109c3848c082584e6cb725e5ab19b18354b2ab85), Struct tag(0x1::coin::CoinStore) and Ledger version(13106860)","error_code":"resource_not_found","vm_error_code":null}`,
		// 	"0",
		// 	"",
		// },
	}

	for _, v := range vectors {
		resp := `{"chain_id":38,"epoch":"133","ledger_version":"13087045","oldest_ledger_version":"0","ledger_timestamp":"1669676013555573","node_role":"full_node","oldest_block_height":"0","block_height":"5435983","git_hash":"2c74a456298fcd520241a562119b6fe30abdaae2"}`
		server, close := test.MockHTTP(&s.Suite, resp)
		defer close()

		asset := v.asset
		asset.URL = server.URL
		client, _ := NewClient(asset)

		server.Response = v.resp
		from := xc.Address("0xa589a80d61ec380c24a5fdda109c3848c082584e6cb725e5ab19b18354b2ab85")
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
