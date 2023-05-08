package sui

import (
	"context"
	"errors"
	"fmt"

	"github.com/coming-chat/go-sui/types"
	xc "github.com/jumpcrypto/crosschain"
	"github.com/jumpcrypto/crosschain/chain/sui/generated/bcs"
	"github.com/jumpcrypto/crosschain/test"
	"github.com/shopspring/decimal"
)

func (s *CrosschainTestSuite) TestFetchTxInfo() {
	require := s.Require()

	vectors := []struct {
		tx   string
		resp interface{}
		val  xc.TxInfo
		err  string
	}{
		{
			"",
			[]string{
				// grab the tx
				`{"digest":"J2Vkui75vgoLvCmNiREVKwpeTVPCq5EQ71i2ETahP6R9","transaction":{"data":{"messageVersion":"v1","transaction":{"kind":"ProgrammableTransaction","inputs":[{"type":"pure","valueType":"u64","value":"10000000000"},{"type":"pure","valueType":"address","value":"0xbb8a8269cf96ba2ec27dc9becd79836394dbe7946c7ac211928be4a0b1de66b9"}],"transactions":[{"SplitCoins":["GasCoin",[{"Input":0}]]},{"TransferObjects":[[{"NestedResult":[0,0]}],{"Input":1}]}]},"sender":"0x7d20dcdb2bca4f508ea9613994683eb4e76e9c4ed371169677c1be02aaf0b58e","gasData":{"payment":[{"objectId":"0xbddb28b55556649dd58e27b39ea80c57295b869a770cb0c04e8ab30cb3a358d8","version":22995,"digest":"FfBTsF3cCgYrN7GeZDMkuLzuyjrZTwdEsdPEokLy4cdQ"}],"owner":"0x7d20dcdb2bca4f508ea9613994683eb4e76e9c4ed371169677c1be02aaf0b58e","price":"1000","budget":"10000000000"}},"txSignatures":["AJE+bRLErdYJJLUKAheAbt+rAIFAM/JRaNPnDafZky4hnjjvmsyWVRymbxqmuaLagV6nQgP7e/bmhUFIUTed9gISBz31NeFTlDTV+RW4oXQeICAR/h+E3u6xe3MRyQsRGw=="]},"effects":{"messageVersion":"v1","status":{"status":"success"},"executedEpoch":"18","gasUsed":{"computationCost":"1000000","storageCost":"1976000","storageRebate":"978120","nonRefundableStorageFee":"9880"},"modifiedAtVersions":[{"objectId":"0xbddb28b55556649dd58e27b39ea80c57295b869a770cb0c04e8ab30cb3a358d8","sequenceNumber":"22995"}],"transactionDigest":"J2Vkui75vgoLvCmNiREVKwpeTVPCq5EQ71i2ETahP6R9","created":[{"owner":{"AddressOwner":"0xbb8a8269cf96ba2ec27dc9becd79836394dbe7946c7ac211928be4a0b1de66b9"},"reference":{"objectId":"0xe6dd381983b77040780e98f9b0a9b12ed3dc8223d1f1dda607120fd007d3ce6b","version":22996,"digest":"ALf7a4D7bpJCvhL4pW2dtk5ZbyjphQcz49nTY9ZP4tCG"}}],"mutated":[{"owner":{"AddressOwner":"0x7d20dcdb2bca4f508ea9613994683eb4e76e9c4ed371169677c1be02aaf0b58e"},"reference":{"objectId":"0xbddb28b55556649dd58e27b39ea80c57295b869a770cb0c04e8ab30cb3a358d8","version":22996,"digest":"9ATwa4EctZHbK2RSEqsrsM6pohyCBR62DDwoFuDUUhVU"}}],"gasObject":{"owner":{"AddressOwner":"0x7d20dcdb2bca4f508ea9613994683eb4e76e9c4ed371169677c1be02aaf0b58e"},"reference":{"objectId":"0xbddb28b55556649dd58e27b39ea80c57295b869a770cb0c04e8ab30cb3a358d8","version":22996,"digest":"9ATwa4EctZHbK2RSEqsrsM6pohyCBR62DDwoFuDUUhVU"}},"dependencies":["CA12cnDvch6aj9WxqThhnbBZ9uVsKG2fFWvc7tfHKQ2n"]},"events":[],"objectChanges":[{"type":"mutated","sender":"0x7d20dcdb2bca4f508ea9613994683eb4e76e9c4ed371169677c1be02aaf0b58e","owner":{"AddressOwner":"0x7d20dcdb2bca4f508ea9613994683eb4e76e9c4ed371169677c1be02aaf0b58e"},"objectType":"0x2::coin::Coin<0x2::sui::SUI>","objectId":"0xbddb28b55556649dd58e27b39ea80c57295b869a770cb0c04e8ab30cb3a358d8","version":"22996","previousVersion":"22995","digest":"9ATwa4EctZHbK2RSEqsrsM6pohyCBR62DDwoFuDUUhVU"},{"type":"created","sender":"0x7d20dcdb2bca4f508ea9613994683eb4e76e9c4ed371169677c1be02aaf0b58e","owner":{"AddressOwner":"0xbb8a8269cf96ba2ec27dc9becd79836394dbe7946c7ac211928be4a0b1de66b9"},"objectType":"0x2::coin::Coin<0x2::sui::SUI>","objectId":"0xe6dd381983b77040780e98f9b0a9b12ed3dc8223d1f1dda607120fd007d3ce6b","version":"22996","digest":"ALf7a4D7bpJCvhL4pW2dtk5ZbyjphQcz49nTY9ZP4tCG"}],"balanceChanges":[{"owner":{"AddressOwner":"0x7d20dcdb2bca4f508ea9613994683eb4e76e9c4ed371169677c1be02aaf0b58e"},"coinType":"0x2::sui::SUI","amount":"-10001997880"},{"owner":{"AddressOwner":"0xbb8a8269cf96ba2ec27dc9becd79836394dbe7946c7ac211928be4a0b1de66b9"},"coinType":"0x2::sui::SUI","amount":"10000000000"}],"timestampMs":"1683124849673","checkpoint":"1953362"}`,
				// grab the latest checkpoint
				`{"data":[{"epoch":"18","sequenceNumber":"1969067","digest":"Cji8yjGsk9Yg5sxUB4iWCQSiEjWHvVTzKMqBChvzNHjV","networkTotalTransactions":"4644633","previousDigest":"wG1Garuf7TkmtjdER4GFkiK9Q2id7ddnk4nAEXu7aDQ","epochRollingGasCostSummary":{"computationCost":"256732544695150","storageCost":"915446949600","storageRebate":"873308640996","nonRefundableStorageFee":"8821299404"},"timestampMs":"1683136992429","transactions":["22W8Rbz4NnYp2Jj7bSt15XrLBMH8Q6cW9HhB8WhAnUs4","3nfvfoJNytiWGDrrfJevXPfGjr1wsaWMB2yBKXDoqsAe","747otLtjqoiDuSKWSuvjCZWYGhNLrnV4fgGkXofNCHf9","A9kji7Sm4Sb4KWK1UudQ9SBymrtYWhCuccPfqpyRhaB6","G6ryjgaWHiRfSq7ULy9LVbp6P2JdphuBfuFRphvCNJHx"],"checkpointCommitments":[],"validatorSignature":"kaUtWOGEuc9DPgOueMBX5NaGDIjbCrXVVolb4jQCX6B02TCXFTAs9wtpRoOitmNv"}],"nextCursor":"1969067","hasNextPage":true}`,
				// grab the checkpoint for the tx
				`{"epoch":"18","sequenceNumber":"1953362","digest":"BHeEq9rUuc2kdh1k7vk4oN22oK7TtpTpVuaM32fs6UrB","networkTotalTransactions":"4594815","previousDigest":"QnwDCY5dJgpY9rRGbYseG8pQgdacHGtXgjC9TPVyKh8","epochRollingGasCostSummary":{"computationCost":"256697889685794","storageCost":"781589532800","storageRebate":"744825062520","nonRefundableStorageFee":"7523485480"},"timestampMs":"1683124849673","transactions":["UMeT2asZr2hsqT3n3vHmfE4UaNMjkVpoQvUpkRpfddd","3V6WU3ofLCd4pfujYnsEY95y4DxQhW8JM6m2Hy8XTbqg","7NHKPiWjTPM84ZBdi9othG3HCw2LUbxfFe183SWdTwhh","9i1bZ9mtMjMsjzvLRwP2YF9DUvLztiY1QbQVWfgGNrRJ","A5w6t3SkpHTUmCA14L7q9bnVQXsKWv5XRrKTP7wM7xCd","FohJy8o9qf3Qwag6wdWA6XCJeJEvjUmMzWgDhVPxMUnU","GAgWaGyeUoRSKGJQS3Mbs3PYeqfFLHtHpyXmHXGAiMAe","Hbo9qbYRj5XncBtpyzEro2aXYgWiThFEJTj8XmR5ArBg","J2Vkui75vgoLvCmNiREVKwpeTVPCq5EQ71i2ETahP6R9"],"checkpointCommitments":[],"validatorSignature":"k7orjPUoopsGMd6eR3JDie70DwppJ3t/F2BVMNDN06vX1FMDgscBuXf970TXoD9z"}`,
			},
			xc.TxInfo{
				BlockHash:     "BHeEq9rUuc2kdh1k7vk4oN22oK7TtpTpVuaM32fs6UrB",
				TxID:          "J2Vkui75vgoLvCmNiREVKwpeTVPCq5EQ71i2ETahP6R9",
				ExplorerURL:   "https://explorer.sui.io/txblock/J2Vkui75vgoLvCmNiREVKwpeTVPCq5EQ71i2ETahP6R9?network=devnet",
				From:          "0x7d20dcdb2bca4f508ea9613994683eb4e76e9c4ed371169677c1be02aaf0b58e",
				To:            "0xbb8a8269cf96ba2ec27dc9becd79836394dbe7946c7ac211928be4a0b1de66b9",
				BlockIndex:    1953362,
				BlockTime:     1683124849673,
				Confirmations: 15705,
				Sources: []*xc.TxInfoEndpoint{
					{
						Address: "0x7d20dcdb2bca4f508ea9613994683eb4e76e9c4ed371169677c1be02aaf0b58e",
						Amount:  xc.NewAmountBlockchainFromStr("10001997880"),
						Asset:   "SUI",
					},
				},
				Destinations: []*xc.TxInfoEndpoint{
					{
						Address: "0xbb8a8269cf96ba2ec27dc9becd79836394dbe7946c7ac211928be4a0b1de66b9",
						Amount:  xc.NewAmountBlockchainFromStr("10000000000"),
						Asset:   "SUI",
					},
				},
				Fee:    xc.NewAmountBlockchainFromStr("1997880"),
				Amount: xc.NewAmountBlockchainFromStr("10000000000"),
			},
			"",
		},
	}

	for _, v := range vectors {
		server, close := test.MockJSONRPC(&s.Suite, v.resp)
		defer close()
		asset := &xc.AssetConfig{NativeAsset: xc.SUI, Net: "devnet", URL: server.URL}

		asset.URL = server.URL
		client, _ := NewClient(asset)
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

func suiCoin(id string, digest string, amount uint64, version_maybe ...int) *types.Coin {
	coinId, err := decodeHex(id)
	if err != nil {
		panic(err)
	}
	version := 1852477
	if len(version_maybe) > 0 {
		version = version_maybe[0]
	}
	var bal types.SafeSuiBigInt[uint64] = types.NewSafeSuiBigInt(amount)
	return &types.Coin{
		CoinType:     "0x2::sui::SUI",
		CoinObjectId: coinId,
		Digest:       digest,
		Balance:      bal,
		// fixed for test
		Version: decimal.NewFromInt(int64(version)),
	}
}

func (s *CrosschainTestSuite) TestTransfer() {
	require := s.Require()

	from := "0xbb8a8269cf96ba2ec27dc9becd79836394dbe7946c7ac211928be4a0b1de66b9"
	from_pk := "6a03aadd27a3753c3af2d676591528f3d8209f337b9506163479bc5e61f67ebd"
	to := "0xaa8a8269cf96ba2ec27dc9becd79836394dbe7946c7ac211928be4a0b1de6600"

	vectors := []struct {
		// in SUI, not mist.
		name      string
		amount    string
		resp      interface{}
		inputs    []bcs.CallArg
		commands  []bcs.Command
		gasBudget uint64
		err       error
	}{
		// Test with 2 sui coins
		{
			"Test_with_2_sui_coin",
			"1.5",
			[]string{
				// get coins
				`{"data":[
					{"coinType":"0x2::sui::SUI","coinObjectId":"0x1cdc19f7751451412d090632bb1ca2c845a9c8f6cd8798d99d304571cfea1ca6","version":"1852477","digest":"u6uSbWNMxkRkCqkjSTbsMeWMYB2VK7pbAo6vFoaMzSo","balance":"2001904720","previousTransaction":"AtPwJTvPfAd47yjBmJCGCJEB7E2XmoJ6aB23XX1o6c4M"},
					{"coinType":"0x2::sui::SUI","coinObjectId":"0x418ca9b7e3bf4bd3ecdb2d45daae92b2428a3488670e28a620ee7ee870f46b2d","version":"1852477","digest":"SwXnkbcrycgr6unAXdcJQ5jfo9dMNMkztMWc3ZxNjL3","balance":"28969157920","previousTransaction":"AtPwJTvPfAd47yjBmJCGCJEB7E2XmoJ6aB23XX1o6c4M"}
				],"nextCursor":"0x418ca9b7e3bf4bd3ecdb2d45daae92b2428a3488670e28a620ee7ee870f46b2d","hasNextPage":false}`,
				// get checkpoint
				`{"data":[{"epoch":"21","sequenceNumber":"2206686","digest":"HtsAAgd1ajMR8qMocnNF6XbAtiBHrxdauGhWtXqKouF3","networkTotalTransactions":"5164703","previousDigest":"H8oYvb73KoG7TWXpw4JPy2qZk7ddvHY3rYQ8kHcNmcua","epochRollingGasCostSummary":{"computationCost":"130960164300","storageCost":"499151462400","storageRebate":"422717709348","nonRefundableStorageFee":"4269875852"},"timestampMs":"1683320609521","transactions":["3yVjcHqKwLN8K8TrZZZMpMUp4VSGg4LRp4uuzvvzzrFD","Cv2NH6zJiRJMtPMzxzZABgDpBfNmb9eniWW9t5v2kPtz","GJaDtfzHap6V8ARdQTstkJm7PiWsEXWkUapXHA2nbmbD"],"checkpointCommitments":[],"validatorSignature":"i3aT5RVtIOvX0pEc/HU+xFTHbw2zV5SdT7q5n6GfS+e85CtkC8qqseeK2Hx9Nhia"}],"nextCursor":"2206686","hasNextPage":true}`,
				//reference gas
				"1000",
				// submit tx
				`{"digest":"5bKyJZUyqHV4aDwQSR9hsiBJXpfTycDoP2NG59bL6p1E","confirmedLocalExecution":true}`,
			},
			// split, merge, split, transfer
			[]bcs.CallArg{
				// remainder split (gas coin balance - gas budget)
				u64ToPure(28969157920 - 2_000_000_000),
				// merged coins after gas coin
				&bcs.CallArg__Object{Value: mustCoinToObject(suiCoin("0x1cdc19f7751451412d090632bb1ca2c845a9c8f6cd8798d99d304571cfea1ca6", "u6uSbWNMxkRkCqkjSTbsMeWMYB2VK7pbAo6vFoaMzSo", 2001904720))},
				// split amt (transfer amount)
				u64ToPure(1_500_000_000),
				// destination address
				mustHexToPure(to),
			},
			[]bcs.Command{
				&bcs.Command__SplitCoins{
					Field0: &bcs.Argument__GasCoin{},
					Field1: []bcs.Argument{ArgumentInput(0)},
				},
				&bcs.Command__MergeCoins{
					Field0: ArgumentInput(1),
					Field1: []bcs.Argument{
						ArgumentResult(0),
					},
				},
				&bcs.Command__SplitCoins{
					Field0: ArgumentInput(1),
					Field1: []bcs.Argument{
						ArgumentInput(2),
					},
				},
				&bcs.Command__TransferObjects{
					Field0: []bcs.Argument{
						ArgumentResult(2),
					},
					Field1: ArgumentInput(3),
				},
			},
			2_000_000_000,
			nil,
		},

		// Test with >>2 sui coins
		{
			"Test_with_many_sui_coin",
			"3",
			[]string{
				// get coins
				`{"data":[
					{"coinType":"0x2::sui::SUI","coinObjectId":"0x1cdc19f7751451412d090632bb1ca2c845a9c8f6cd8798d99d304571cfea1ca6","version":"1852477","digest":"SNiJ8aV9rerhbVTwZikSAWVgJPhx9jxaPXdGcfeYut9","balance": "45035120" ,"previousTransaction":"D85viPWoceLm1siButSgA9Z7fyfRR7GcnZvti1HgmXp8"},
					{"coinType":"0x2::sui::SUI","coinObjectId":"0x3a705927f067d86aedf19bbe84bb15cedbf613c3ac5c88b8508d8ce3f9bbbf7c","version":"1852477","digest":"GreYy8apDQHR7zwsZLHZ6bfQAiQ12xE4TNcm4vznpNUM","balance":"1000000000","previousTransaction":"D85viPWoceLm1siButSgA9Z7fyfRR7GcnZvti1HgmXp8"},
					{"coinType":"0x2::sui::SUI","coinObjectId":"0x5c22194a002befba3d34d26036d4c440f86099ac4cb9b8aaeca22fb379229237","version":"1852477","digest":"3t7sWDqfyKvbGxtnS1GwMV2kgdasLgmDTRJ7MHhoyCz3","balance":"1300000000","previousTransaction":"D85viPWoceLm1siButSgA9Z7fyfRR7GcnZvti1HgmXp8"},
					{"coinType":"0x2::sui::SUI","coinObjectId":"0x87bae5d7376e857106f7908eab6f7106ea3f7c2a1b3349f99925bb12631b1ff0","version":"1852477","digest":"9GeMg1yw4J9ck62XR3KHXi72kfVVeuqfAcK5rL3hRdVK","balance":"1500000000","previousTransaction":"D85viPWoceLm1siButSgA9Z7fyfRR7GcnZvti1HgmXp8"},
					{"coinType":"0x2::sui::SUI","coinObjectId":"0xa47d2121ef5ca77d83723d72a6b70c3bce15a2f438294f2d0fbcb530ab5d0b27","version":"1852477","digest":"EYMFpVaEcfdv8kv1hxZz8y884z2fhQJt8d3G1zKBYf6m","balance":"1200000000","previousTransaction":"D85viPWoceLm1siButSgA9Z7fyfRR7GcnZvti1HgmXp8"},
					{"coinType":"0x2::sui::SUI","coinObjectId":"0xc587db1fbe680b769c1a562a09f2c871a087bafa542c7cb73db6064e2b791bdf","version":"1852477","digest":"7Y2zjQxn2wj5jhrvS5NBKCFJDzWHZ4UMG7XJNNioNgTS","balance":"1897841920","previousTransaction":"D85viPWoceLm1siButSgA9Z7fyfRR7GcnZvti1HgmXp8"},
					{"coinType":"0x2::sui::SUI","coinObjectId":"0xeb8b3e2e9c446f25f29fee61c43583b9d7fbfed600a83219ca99d1ee681ac958","version":"1852477","digest":"DkV4WuN3ZPLHfF87otc23aVzwfbJQWyP171YhEyJQG5Q","balance":"1000000000","previousTransaction":"D85viPWoceLm1siButSgA9Z7fyfRR7GcnZvti1HgmXp8"},
					{"coinType":"0x2::sui::SUI","coinObjectId":"0xefa9b328eb6b955be0a23033c405a4281b0ebca01844e2fc963b8a7408062323","version":"1852477","digest":"8KXUUETW9jss1Z7Sj7wSECJmDbxNqj97mP9u4y5wsbMz","balance":"5000000000","previousTransaction":"HcfDapsNNdsUHYh1pau7Tio3Gvi7GqhphzBVjvYeGU16"},
					{"coinType":"0x2::sui::SUI","coinObjectId":"0xfb889571ed135b9bd1c1fd7d00d69694305bb74113efa087c18d6444528da091","version":"1852477","digest":"8XS5rHkHwoYtSq2Fg49NbZsxRT6tM5n31wx3bKMjpRDy","balance":"900000000","previousTransaction":"D85viPWoceLm1siButSgA9Z7fyfRR7GcnZvti1HgmXp8"}],
					"nextCursor":"0xfb889571ed135b9bd1c1fd7d00d69694305bb74113efa087c18d6444528da091","hasNextPage":false}`,
				// get checkpoint
				`{"data":[{"epoch":"21","sequenceNumber":"2212095","digest":"3Gfav3tbk6vgpLp456y6MynYTAzoar6wodfNq5Ahw1M9","networkTotalTransactions":"5173764","previousDigest":"6iVUW6jom9Z8jUdDbseW6q2ESiaVcVGddtKwZX9i7yNt","epochRollingGasCostSummary":{"computationCost":"3402552000","storageCost":"19937429600","storageRebate":"18713346696","nonRefundableStorageFee":"189023704"},"timestampMs":"1683324797581","transactions":["9o478feJK21ao7Z9GUq51WC6jFWmNjj8MKzDcjzAkVvk"],"checkpointCommitments":[],"validatorSignature":"q5LhcNeDArLCUy2kTP8jpwj/vexiFaKnXR/v1UtzVujHzna6SIgrGWsENOebd3+z"}],"nextCursor":"2212095","hasNextPage":true}`,
				//reference gas
				"1000",
				// submit tx
				`{"digest":"BzJbzapMeyC1QrdC5Q7H4okbxyZQJ9MntWaTBHesi3cW","confirmedLocalExecution":true}`,
			},
			// split, merge, split, transfer
			[]bcs.CallArg{
				// remainder split
				u64ToPure(3000000000),
				// merged coins after gas coin (sorted by value)
				&bcs.CallArg__Object{Value: mustCoinToObject(suiCoin("0xc587db1fbe680b769c1a562a09f2c871a087bafa542c7cb73db6064e2b791bdf", "7Y2zjQxn2wj5jhrvS5NBKCFJDzWHZ4UMG7XJNNioNgTS", 1897841920))},
				&bcs.CallArg__Object{Value: mustCoinToObject(suiCoin("0x87bae5d7376e857106f7908eab6f7106ea3f7c2a1b3349f99925bb12631b1ff0", "9GeMg1yw4J9ck62XR3KHXi72kfVVeuqfAcK5rL3hRdVK", 1500000000))},
				&bcs.CallArg__Object{Value: mustCoinToObject(suiCoin("0x5c22194a002befba3d34d26036d4c440f86099ac4cb9b8aaeca22fb379229237", "3t7sWDqfyKvbGxtnS1GwMV2kgdasLgmDTRJ7MHhoyCz3", 1300000000))},
				&bcs.CallArg__Object{Value: mustCoinToObject(suiCoin("0xa47d2121ef5ca77d83723d72a6b70c3bce15a2f438294f2d0fbcb530ab5d0b27", "EYMFpVaEcfdv8kv1hxZz8y884z2fhQJt8d3G1zKBYf6m", 1200000000))},
				&bcs.CallArg__Object{Value: mustCoinToObject(suiCoin("0x3a705927f067d86aedf19bbe84bb15cedbf613c3ac5c88b8508d8ce3f9bbbf7c", "GreYy8apDQHR7zwsZLHZ6bfQAiQ12xE4TNcm4vznpNUM", 1000000000))},
				&bcs.CallArg__Object{Value: mustCoinToObject(suiCoin("0xeb8b3e2e9c446f25f29fee61c43583b9d7fbfed600a83219ca99d1ee681ac958", "DkV4WuN3ZPLHfF87otc23aVzwfbJQWyP171YhEyJQG5Q", 1000000000))},
				&bcs.CallArg__Object{Value: mustCoinToObject(suiCoin("0xfb889571ed135b9bd1c1fd7d00d69694305bb74113efa087c18d6444528da091", "8XS5rHkHwoYtSq2Fg49NbZsxRT6tM5n31wx3bKMjpRDy", 900000000))},
				&bcs.CallArg__Object{Value: mustCoinToObject(suiCoin("0x1cdc19f7751451412d090632bb1ca2c845a9c8f6cd8798d99d304571cfea1ca6", "SNiJ8aV9rerhbVTwZikSAWVgJPhx9jxaPXdGcfeYut9", 45035120))},
				// split amt (transfer amount)
				u64ToPure(3_000_000_000),
				// destination address
				mustHexToPure(to),
			},
			[]bcs.Command{
				&bcs.Command__SplitCoins{
					Field0: &bcs.Argument__GasCoin{},
					Field1: []bcs.Argument{ArgumentInput(0)},
				},
				&bcs.Command__MergeCoins{
					Field0: ArgumentInput(1),
					Field1: []bcs.Argument{
						ArgumentResult(0),
						ArgumentInput(2),
						ArgumentInput(3),
						ArgumentInput(4),
						ArgumentInput(5),
						ArgumentInput(6),
						ArgumentInput(7),
						ArgumentInput(8),
					},
				},
				&bcs.Command__SplitCoins{
					Field0: ArgumentInput(1),
					Field1: []bcs.Argument{
						ArgumentInput(9),
					},
				},
				&bcs.Command__TransferObjects{
					Field0: []bcs.Argument{
						ArgumentResult(2),
					},
					Field1: ArgumentInput(10),
				},
			},
			2_000_000_000,
			nil,
		},
		// Test with 1 sui coin
		{
			"Test_with_1_sui_coin",
			"1.0",
			[]string{
				// get coins
				`{"data":[
					{"coinType":"0x2::sui::SUI","coinObjectId":"0xc587db1fbe680b769c1a562a09f2c871a087bafa542c7cb73db6064e2b791bdf","version":"1852491","digest":"GBm2HRW1WvNRrGX5iM3syjbD1PeaWQs69s42wJEam7HY","balance":"1845686480","previousTransaction":"4qkLLVGsxNwvvpJMwSbCh4jFmC9J8Cb1x1zhNaC7k5cK"}
					],"nextCursor":"0xc587db1fbe680b769c1a562a09f2c871a087bafa542c7cb73db6064e2b791bdf","hasNextPage":false}`,
				// get checkpoint
				`{"data":[{"epoch":"21","sequenceNumber":"2206686","digest":"HtsAAgd1ajMR8qMocnNF6XbAtiBHrxdauGhWtXqKouF3","networkTotalTransactions":"5164703","previousDigest":"H8oYvb73KoG7TWXpw4JPy2qZk7ddvHY3rYQ8kHcNmcua","epochRollingGasCostSummary":{"computationCost":"130960164300","storageCost":"499151462400","storageRebate":"422717709348","nonRefundableStorageFee":"4269875852"},"timestampMs":"1683320609521","transactions":["3yVjcHqKwLN8K8TrZZZMpMUp4VSGg4LRp4uuzvvzzrFD","Cv2NH6zJiRJMtPMzxzZABgDpBfNmb9eniWW9t5v2kPtz","GJaDtfzHap6V8ARdQTstkJm7PiWsEXWkUapXHA2nbmbD"],"checkpointCommitments":[],"validatorSignature":"i3aT5RVtIOvX0pEc/HU+xFTHbw2zV5SdT7q5n6GfS+e85CtkC8qqseeK2Hx9Nhia"}],"nextCursor":"2206686","hasNextPage":true}`,
				//reference gas
				"1000",
				// submit tx
				`{"digest":"5bKyJZUyqHV4aDwQSR9hsiBJXpfTycDoP2NG59bL6p1E","confirmedLocalExecution":true}`,
			},
			// split, merge, split, transfer
			[]bcs.CallArg{
				// no split of gas coin
				// no merge coin this time

				// split amt (transfer amount)
				u64ToPure(1_000_000_000),
				// destination address
				mustHexToPure(to),
			},
			[]bcs.Command{
				&bcs.Command__SplitCoins{
					Field0: &bcs.Argument__GasCoin{},
					Field1: []bcs.Argument{
						ArgumentInput(0),
					},
				},
				&bcs.Command__TransferObjects{
					Field0: []bcs.Argument{
						ArgumentResult(0),
					},
					Field1: ArgumentInput(1),
				},
			},
			// this gas budget is lower because we don't have 2 sui.
			// so it should be the leftover.
			1845686480 - 1_000_000_000,
			nil,
		},
		// Test with 1 sui coin, no balance
		{
			"Test_with_no_balance",
			"10.0",
			[]string{
				// get coins
				`{"data":[
					{"coinType":"0x2::sui::SUI","coinObjectId":"0xc587db1fbe680b769c1a562a09f2c871a087bafa542c7cb73db6064e2b791bdf","version":"1852491","digest":"GBm2HRW1WvNRrGX5iM3syjbD1PeaWQs69s42wJEam7HY","balance":"1845686480","previousTransaction":"4qkLLVGsxNwvvpJMwSbCh4jFmC9J8Cb1x1zhNaC7k5cK"}
					],"nextCursor":"0xc587db1fbe680b769c1a562a09f2c871a087bafa542c7cb73db6064e2b791bdf","hasNextPage":false}`,
				// get checkpoint
				`{"data":[{"epoch":"21","sequenceNumber":"2206686","digest":"HtsAAgd1ajMR8qMocnNF6XbAtiBHrxdauGhWtXqKouF3","networkTotalTransactions":"5164703","previousDigest":"H8oYvb73KoG7TWXpw4JPy2qZk7ddvHY3rYQ8kHcNmcua","epochRollingGasCostSummary":{"computationCost":"130960164300","storageCost":"499151462400","storageRebate":"422717709348","nonRefundableStorageFee":"4269875852"},"timestampMs":"1683320609521","transactions":["3yVjcHqKwLN8K8TrZZZMpMUp4VSGg4LRp4uuzvvzzrFD","Cv2NH6zJiRJMtPMzxzZABgDpBfNmb9eniWW9t5v2kPtz","GJaDtfzHap6V8ARdQTstkJm7PiWsEXWkUapXHA2nbmbD"],"checkpointCommitments":[],"validatorSignature":"i3aT5RVtIOvX0pEc/HU+xFTHbw2zV5SdT7q5n6GfS+e85CtkC8qqseeK2Hx9Nhia"}],"nextCursor":"2206686","hasNextPage":true}`,
				//reference gas
				"1000",
				// submit tx
				`{"digest":"5bKyJZUyqHV4aDwQSR9hsiBJXpfTycDoP2NG59bL6p1E","confirmedLocalExecution":true}`,
			},
			// split, merge, split, transfer
			[]bcs.CallArg{},
			[]bcs.Command{},
			0,
			errors.New("not enough funds"),
		},
	}

	for _, v := range vectors {
		fmt.Println("Running ", v.name)
		server, close := test.MockJSONRPC(&s.Suite, v.resp)
		defer close()
		asset := &xc.AssetConfig{NativeAsset: xc.SUI, Net: "devnet", URL: server.URL}

		asset.URL = server.URL
		client, err := NewClient(asset)
		require.NoError(err)
		input, err := client.FetchTxInput(context.Background(), xc.Address(from), xc.Address(to))
		require.NoError(err)
		local_input := input.(*TxInput)
		local_input.SetPublicKeyFromStr(from_pk)

		// check that the gas coin was not also included in
		// the list of coins to spend.
		for _, coin := range local_input.Coins {
			require.NotEqualValues(coin.CoinObjectId, local_input.GasCoin.CoinObjectId)
		}

		builder, err := NewTxBuilder(asset)
		require.NoError(err)
		amount_human_dec, err := decimal.NewFromString(v.amount)
		require.NoError(err)
		amount_machine := xc.AmountHumanReadable(amount_human_dec).ToBlockchain(9)

		tx, err := builder.NewTransfer(xc.Address(from), xc.Address(to), amount_machine, input)
		if v.err == nil {
			require.NoError(err)
		} else {
			require.ErrorContains(err, v.err.Error())
			continue
		}
		suiTx := tx.(*Tx).Tx
		// check various properties of the sui tx
		fromData, _ := hexToAddress(string(from))
		expiration := bcs.TransactionExpiration__Epoch(21)
		gasObjectId, _ := hexToObjectID(local_input.GasCoin.CoinObjectId.String())
		gasDigest, _ := base58ToObjectDigest(local_input.GasCoin.Digest)
		gasVersion := local_input.GasCoin.Version.BigInt().Uint64()

		gasCoin := ObjectRef{
			Field0: gasObjectId,
			Field1: bcs.SequenceNumber(gasVersion),
			Field2: gasDigest,
		}

		require.EqualValues(suiTx.Value.Expiration, &expiration)

		require.EqualValues(suiTx.Value.GasData.Budget, v.gasBudget)
		require.EqualValues(suiTx.Value.GasData.Price, 1_000)
		require.EqualValues(suiTx.Value.GasData.Owner, fromData)
		require.EqualValues(suiTx.Value.GasData.Payment, []struct {
			Field0 bcs.ObjectID
			Field1 bcs.SequenceNumber
			Field2 bcs.ObjectDigest
		}{
			gasCoin,
		})

		commands := suiTx.Value.Kind.(*bcs.TransactionKind__ProgrammableTransaction).Value.Commands
		inputs := suiTx.Value.Kind.(*bcs.TransactionKind__ProgrammableTransaction).Value.Inputs

		require.Len(commands, len(v.commands))
		require.Len(inputs, len(v.inputs))
		for i, cmd := range v.commands {
			require.Equal(cmd, commands[i])
		}

		for i, inp := range v.inputs {
			fmt.Println("checking input ", i)
			require.Equal(inp, inputs[i])
		}

		err = client.SubmitTx(context.Background(), tx)
		require.NoError(err)
	}
}

func (s *CrosschainTestSuite) TestFetchBalance() {
	require := s.Require()

	from := xc.Address("0xbb8a8269cf96ba2ec27dc9becd79836394dbe7946c7ac211928be4a0b1de66b9")

	vectors := []struct {
		// in SUI, not mist.
		name   string
		amount string
		resp   interface{}
		err    error
	}{
		// Test with 2 sui coins
		{
			"fetch_balance_2_coins",
			"35.8336498",
			[]string{
				// get coins
				`{"data":[{"coinType":"0x2::sui::SUI","coinObjectId":"0x8192d5c2b5722c60866761927d5a0737cd55d0c2b1150eabf818253795b38998","version":"1852497","digest":"FL9GS4b72Ay3Lwc55Q9A9DMLDTXUQ5ancKnTfL6WD8JL","balance":"1997992240","previousTransaction":"7g2RPre2F7WJxYBG5urbZvjKev8YpfMxRT8GL8HCshv5"},{"coinType":"0x2::sui::SUI","coinObjectId":"0xc587db1fbe680b769c1a562a09f2c871a087bafa542c7cb73db6064e2b791bdf","version":"1852497","digest":"7epS94m8djHYKu4V29DSTfkm6mJ6TvZAzwY6ntA65v9A","balance":"33835657560","previousTransaction":"7g2RPre2F7WJxYBG5urbZvjKev8YpfMxRT8GL8HCshv5"}],"nextCursor":"0xc587db1fbe680b769c1a562a09f2c871a087bafa542c7cb73db6064e2b791bdf","hasNextPage":false}`,
			},
			nil,
		},
		{
			"fetch_balance_4_coins",
			"71.6672996",
			[]string{
				// get coins
				`{"data":[{"coinType":"0x2::sui::SUI","coinObjectId":"0x8192d5c2b5722c60866761927d5a0737cd55d0c2b1150eabf818253795b38998","version":"1852497","digest":"FL9GS4b72Ay3Lwc55Q9A9DMLDTXUQ5ancKnTfL6WD8JL","balance":"1997992240","previousTransaction":"7g2RPre2F7WJxYBG5urbZvjKev8YpfMxRT8GL8HCshv5"},{"coinType":"0x2::sui::SUI","coinObjectId":"0xc587db1fbe680b769c1a562a09f2c871a087bafa542c7cb73db6064e2b791bdf","version":"1852497","digest":"7epS94m8djHYKu4V29DSTfkm6mJ6TvZAzwY6ntA65v9A","balance":"33835657560","previousTransaction":"7g2RPre2F7WJxYBG5urbZvjKev8YpfMxRT8GL8HCshv5"}],"nextCursor":"0xc587db1fbe680b769c1a562a09f2c871a087bafa542c7cb73db6064e2b791bdf","hasNextPage":true}`,
				`{"data":[{"coinType":"0x2::sui::SUI","coinObjectId":"0x8192d5c2b5722c60866761927d5a0737cd55d0c2b1150eabf818253795b38990","version":"1852497","digest":"FL9GS4b72Ay3Lwc55Q9A9DMLDTXUQ5ancKnTfL6WD8JL","balance":"1997992240","previousTransaction":"7g2RPre2F7WJxYBG5urbZvjKev8YpfMxRT8GL8HCshv5"},{"coinType":"0x2::sui::SUI","coinObjectId":"0xc587db1fbe680b769c1a562a09f2c871a087bafa542c7cb73db6064e2b791bd0","version":"1852497","digest":"7epS94m8djHYKu4V29DSTfkm6mJ6TvZAzwY6ntA65v9A","balance":"33835657560","previousTransaction":"7g2RPre2F7WJxYBG5urbZvjKev8YpfMxRT8GL8HCshv5"}],"nextCursor":"0xc587db1fbe680b769c1a562a09f2c871a087bafa542c7cb73db6064e2b791bdf","hasNextPage":false}`,
			},
			nil,
		},
	}

	for _, v := range vectors {
		fmt.Println("Running ", v.name)
		ctx := context.Background()
		server, close := test.MockJSONRPC(&s.Suite, v.resp)
		defer close()
		asset := &xc.AssetConfig{NativeAsset: xc.SUI, Net: "devnet", URL: server.URL}

		asset.URL = server.URL
		client, err := NewClient(asset)
		require.NoError(err)

		bal, err := client.FetchBalance(ctx, xc.Address(from))

		if v.err != nil {
			require.ErrorContains(err, v.err.Error())
			continue
		}
		require.NoError(err)
		require.EqualValues(v.amount, bal.ToHuman(9).String())
	}
}
