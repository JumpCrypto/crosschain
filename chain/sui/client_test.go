package sui

import (
	xc "github.com/jumpcrypto/crosschain"
	"github.com/jumpcrypto/crosschain/test"
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
			xc.TxInfo{},
			"",
		},
	}

	for _, v := range vectors {
		server, close := test.MockHTTP(&s.Suite, v.resp)
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
