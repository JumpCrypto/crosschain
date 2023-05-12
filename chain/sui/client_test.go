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
		name string
		resp interface{}
		val  xc.TxInfo
		err  string
	}{
		{
			"sui_deposit",
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
						Address:     "0x7d20dcdb2bca4f508ea9613994683eb4e76e9c4ed371169677c1be02aaf0b58e",
						Amount:      xc.NewAmountBlockchainFromStr("10001997880"),
						Asset:       "SUI",
						NativeAsset: "SUI",
					},
				},
				Destinations: []*xc.TxInfoEndpoint{
					{
						Address:     "0xbb8a8269cf96ba2ec27dc9becd79836394dbe7946c7ac211928be4a0b1de66b9",
						Amount:      xc.NewAmountBlockchainFromStr("10000000000"),
						Asset:       "SUI",
						NativeAsset: "SUI",
					},
				},
				Fee:    xc.NewAmountBlockchainFromStr("1997880"),
				Amount: xc.NewAmountBlockchainFromStr("10000000000"),
			},
			"",
		},
		// Test that a rejected transaction is reported as failed.
		{
			"sui_failed_tx",
			[]string{
				// grab the tx
				`{"digest":"9aFfSrP7jvvteSS4q8L8RMC71NbfBeK1FK8aWcw8c8py","transaction":{"data":{"messageVersion":"v1","transaction":{"kind":"ProgrammableTransaction","inputs":[{"type":"pure","valueType":"u64","value":"31835657560"},{"type":"object","objectType":"immOrOwnedObject","objectId":"0x3150377d1db0395abfd3b19cfeca94eaf5987a12b95a0aab431195e77399f092","version":"22585","digest":"4Q9Bp5Auw4VLSCUPMrZDdik59ka8y84C4kNjtR4Lbu1M"},{"type":"object","objectType":"immOrOwnedObject","objectId":"0x5a23ee6e22faa7017b11ad24e7c8ced1d33465cfd06656bc028eb21c6f4cad97","version":"23425","digest":"8ijx3ir4ANizqUpHSx5HW1vgiZFX1Jdr652hWBA4biyT"},{"type":"object","objectType":"immOrOwnedObject","objectId":"0x7d1775d7f791554b25933fde2b91d578ddc2874d1f402b55a7b8f5fb270b845d","version":"23506","digest":"AHMmVxis2no8PKF6WR6eu1AmmRFpNWrJFwJJT4XMkEuY"},{"type":"object","objectType":"immOrOwnedObject","objectId":"0x92e60b8b39e5f3ecd57f6ed98de382549f50ab64ddfe8643b8c9b4b12a77cee1","version":"22513","digest":"BUahUimbYrX8w7zhZZRCbgkLvnJfLnqwAnS3cTNEruaT"},{"type":"object","objectType":"immOrOwnedObject","objectId":"0xd65a04388a0b9378e87d0195e98bd0f9f7b460aac22ebd89fb3ba19e1759f414","version":"23506","digest":"6JRrZPgQZ8EHaf2qbwhzWmDopZeGhcUraqy8aXJXVa5M"},{"type":"object","objectType":"immOrOwnedObject","objectId":"0xe33119108d864f4169d7ed7fa963f51aaed7ef7107d8785cca237916e5079d7c","version":"23425","digest":"7tdHk5JuiUgvBQ1sHPrzah77ffEPrA4DfnndLJ19xE5r"},{"type":"object","objectType":"immOrOwnedObject","objectId":"0x8192d5c2b5722c60866761927d5a0737cd55d0c2b1150eabf818253795b38998","version":"1852497","digest":"FL9GS4b72Ay3Lwc55Q9A9DMLDTXUQ5ancKnTfL6WD8JL"},{"type":"pure","valueType":"u64","value":"100000000000"},{"type":"pure","valueType":"address","value":"0x7d20dcdb2bca4f508ea9613994683eb4e76e9c4ed371169677c1be02aaf0b58e"}],"transactions":[{"SplitCoins":["GasCoin",[{"Input":0}]]},{"MergeCoins":[{"Input":1},[{"Result":0},{"Input":2},{"Input":3},{"Input":4},{"Input":5},{"Input":6},{"Input":7}]]},{"SplitCoins":[{"Input":1},[{"Input":8}]]},{"TransferObjects":[[{"Result":2}],{"Input":9}]}]},"sender":"0xbb8a8269cf96ba2ec27dc9becd79836394dbe7946c7ac211928be4a0b1de66b9","gasData":{"payment":[{"objectId":"0xc587db1fbe680b769c1a562a09f2c871a087bafa542c7cb73db6064e2b791bdf","version":1852497,"digest":"7epS94m8djHYKu4V29DSTfkm6mJ6TvZAzwY6ntA65v9A"}],"owner":"0xbb8a8269cf96ba2ec27dc9becd79836394dbe7946c7ac211928be4a0b1de66b9","price":"1000","budget":"2000000000"}},"txSignatures":["ADHNGErTgWHJCs8hm3k9ZyTI8AEUoHrSuyC6VQF7voUzohX6Xfe20Nk/thH/ADNkARWbKYq0kJ0VKaRNpL1e7AhqA6rdJ6N1PDry1nZZFSjz2CCfM3uVBhY0ebxeYfZ+vQ=="]},"effects":{"messageVersion":"v1","status":{"status":"failure","error":"InsufficientCoinBalance in command 2"},"executedEpoch":"25","gasUsed":{"computationCost":"1000000","storageCost":"7904000","storageRebate":"7824960","nonRefundableStorageFee":"79040"},"modifiedAtVersions":[{"objectId":"0x3150377d1db0395abfd3b19cfeca94eaf5987a12b95a0aab431195e77399f092","sequenceNumber":"22585"},{"objectId":"0x5a23ee6e22faa7017b11ad24e7c8ced1d33465cfd06656bc028eb21c6f4cad97","sequenceNumber":"23425"},{"objectId":"0x7d1775d7f791554b25933fde2b91d578ddc2874d1f402b55a7b8f5fb270b845d","sequenceNumber":"23506"},{"objectId":"0x8192d5c2b5722c60866761927d5a0737cd55d0c2b1150eabf818253795b38998","sequenceNumber":"1852497"},{"objectId":"0x92e60b8b39e5f3ecd57f6ed98de382549f50ab64ddfe8643b8c9b4b12a77cee1","sequenceNumber":"22513"},{"objectId":"0xc587db1fbe680b769c1a562a09f2c871a087bafa542c7cb73db6064e2b791bdf","sequenceNumber":"1852497"},{"objectId":"0xd65a04388a0b9378e87d0195e98bd0f9f7b460aac22ebd89fb3ba19e1759f414","sequenceNumber":"23506"},{"objectId":"0xe33119108d864f4169d7ed7fa963f51aaed7ef7107d8785cca237916e5079d7c","sequenceNumber":"23425"}],"transactionDigest":"9aFfSrP7jvvteSS4q8L8RMC71NbfBeK1FK8aWcw8c8py","mutated":[{"owner":{"AddressOwner":"0xbb8a8269cf96ba2ec27dc9becd79836394dbe7946c7ac211928be4a0b1de66b9"},"reference":{"objectId":"0x3150377d1db0395abfd3b19cfeca94eaf5987a12b95a0aab431195e77399f092","version":1852498,"digest":"De3ysFkPDxrzVMW46Jzike9hJ8xgAsy4ZCcctcLuDe9A"}},{"owner":{"AddressOwner":"0xbb8a8269cf96ba2ec27dc9becd79836394dbe7946c7ac211928be4a0b1de66b9"},"reference":{"objectId":"0x5a23ee6e22faa7017b11ad24e7c8ced1d33465cfd06656bc028eb21c6f4cad97","version":1852498,"digest":"537fcd4aKnd4cEDzV9fwvihH91x1F7BxHYNHLDZhhvyJ"}},{"owner":{"AddressOwner":"0xbb8a8269cf96ba2ec27dc9becd79836394dbe7946c7ac211928be4a0b1de66b9"},"reference":{"objectId":"0x7d1775d7f791554b25933fde2b91d578ddc2874d1f402b55a7b8f5fb270b845d","version":1852498,"digest":"6kcb5Pr9bwaaaDWBCEqe6eAfGEhVBHoyJkaaxesHL9J7"}},{"owner":{"AddressOwner":"0xbb8a8269cf96ba2ec27dc9becd79836394dbe7946c7ac211928be4a0b1de66b9"},"reference":{"objectId":"0x8192d5c2b5722c60866761927d5a0737cd55d0c2b1150eabf818253795b38998","version":1852498,"digest":"7HANAnn32L4c6aRKnurxtpncPdRjEuPA59dEJih7APkE"}},{"owner":{"AddressOwner":"0xbb8a8269cf96ba2ec27dc9becd79836394dbe7946c7ac211928be4a0b1de66b9"},"reference":{"objectId":"0x92e60b8b39e5f3ecd57f6ed98de382549f50ab64ddfe8643b8c9b4b12a77cee1","version":1852498,"digest":"85wiAzekdiq6kzc4L8WqHgzYUToEQevYNQaVLR92Jcax"}},{"owner":{"AddressOwner":"0xbb8a8269cf96ba2ec27dc9becd79836394dbe7946c7ac211928be4a0b1de66b9"},"reference":{"objectId":"0xc587db1fbe680b769c1a562a09f2c871a087bafa542c7cb73db6064e2b791bdf","version":1852498,"digest":"29FvSiZFhZ7z7HKEMFAza1H38m8hZogBCMD9TEmAtq3a"}},{"owner":{"AddressOwner":"0xbb8a8269cf96ba2ec27dc9becd79836394dbe7946c7ac211928be4a0b1de66b9"},"reference":{"objectId":"0xd65a04388a0b9378e87d0195e98bd0f9f7b460aac22ebd89fb3ba19e1759f414","version":1852498,"digest":"4s2P3qRZxQGpZhXz4EG5t8VDcrmzN3yaRqXF3AwQCnMs"}},{"owner":{"AddressOwner":"0xbb8a8269cf96ba2ec27dc9becd79836394dbe7946c7ac211928be4a0b1de66b9"},"reference":{"objectId":"0xe33119108d864f4169d7ed7fa963f51aaed7ef7107d8785cca237916e5079d7c","version":1852498,"digest":"7sTJRtJypM1pTGUwJxDcqqvoTzDMWKwnzrLmUoJwWPmy"}}],"gasObject":{"owner":{"AddressOwner":"0xbb8a8269cf96ba2ec27dc9becd79836394dbe7946c7ac211928be4a0b1de66b9"},"reference":{"objectId":"0xc587db1fbe680b769c1a562a09f2c871a087bafa542c7cb73db6064e2b791bdf","version":1852498,"digest":"29FvSiZFhZ7z7HKEMFAza1H38m8hZogBCMD9TEmAtq3a"}},"dependencies":["3DB59fmaogiJYpv1sk4qqDUs9GBr3su1TXvz9CyxsGeL","3ZQj4vfygTFYJ37Vxhd4CiPHSx8Zs43LT289bnCP1asz","52b5qwqzoUFnooftzhXgN7okpxhzipYtU5NEbHTvWM5Y","7g2RPre2F7WJxYBG5urbZvjKev8YpfMxRT8GL8HCshv5","AR6416r8xmdCMoXxDNvUyRwn9jcfRNU6ZABnKDe7g8pQ","BvcAM1TZATGL4FfXJjDWSA1Y6QcRRNCriLSm4SPddU4G","CeU2wgD7P1Zn6WJdtrkytsS2aqH3mLqkwTiGQgtpjeAN"]},"events":[],"objectChanges":[{"type":"mutated","sender":"0xbb8a8269cf96ba2ec27dc9becd79836394dbe7946c7ac211928be4a0b1de66b9","owner":{"AddressOwner":"0xbb8a8269cf96ba2ec27dc9becd79836394dbe7946c7ac211928be4a0b1de66b9"},"objectType":"0x0000000000000000000000000000000000000000000000000000000000000002::coin::Coin<0x0000000000000000000000000000000000000000000000000000000000000002::sui::SUI>","objectId":"0x3150377d1db0395abfd3b19cfeca94eaf5987a12b95a0aab431195e77399f092","version":"1852498","previousVersion":"22585","digest":"De3ysFkPDxrzVMW46Jzike9hJ8xgAsy4ZCcctcLuDe9A"},{"type":"mutated","sender":"0xbb8a8269cf96ba2ec27dc9becd79836394dbe7946c7ac211928be4a0b1de66b9","owner":{"AddressOwner":"0xbb8a8269cf96ba2ec27dc9becd79836394dbe7946c7ac211928be4a0b1de66b9"},"objectType":"0x0000000000000000000000000000000000000000000000000000000000000002::coin::Coin<0x0000000000000000000000000000000000000000000000000000000000000002::sui::SUI>","objectId":"0x5a23ee6e22faa7017b11ad24e7c8ced1d33465cfd06656bc028eb21c6f4cad97","version":"1852498","previousVersion":"23425","digest":"537fcd4aKnd4cEDzV9fwvihH91x1F7BxHYNHLDZhhvyJ"},{"type":"mutated","sender":"0xbb8a8269cf96ba2ec27dc9becd79836394dbe7946c7ac211928be4a0b1de66b9","owner":{"AddressOwner":"0xbb8a8269cf96ba2ec27dc9becd79836394dbe7946c7ac211928be4a0b1de66b9"},"objectType":"0x0000000000000000000000000000000000000000000000000000000000000002::coin::Coin<0x0000000000000000000000000000000000000000000000000000000000000002::sui::SUI>","objectId":"0x7d1775d7f791554b25933fde2b91d578ddc2874d1f402b55a7b8f5fb270b845d","version":"1852498","previousVersion":"23506","digest":"6kcb5Pr9bwaaaDWBCEqe6eAfGEhVBHoyJkaaxesHL9J7"},{"type":"mutated","sender":"0xbb8a8269cf96ba2ec27dc9becd79836394dbe7946c7ac211928be4a0b1de66b9","owner":{"AddressOwner":"0xbb8a8269cf96ba2ec27dc9becd79836394dbe7946c7ac211928be4a0b1de66b9"},"objectType":"0x0000000000000000000000000000000000000000000000000000000000000002::coin::Coin<0x0000000000000000000000000000000000000000000000000000000000000002::sui::SUI>","objectId":"0x8192d5c2b5722c60866761927d5a0737cd55d0c2b1150eabf818253795b38998","version":"1852498","previousVersion":"1852497","digest":"7HANAnn32L4c6aRKnurxtpncPdRjEuPA59dEJih7APkE"},{"type":"mutated","sender":"0xbb8a8269cf96ba2ec27dc9becd79836394dbe7946c7ac211928be4a0b1de66b9","owner":{"AddressOwner":"0xbb8a8269cf96ba2ec27dc9becd79836394dbe7946c7ac211928be4a0b1de66b9"},"objectType":"0x0000000000000000000000000000000000000000000000000000000000000002::coin::Coin<0x0000000000000000000000000000000000000000000000000000000000000002::sui::SUI>","objectId":"0x92e60b8b39e5f3ecd57f6ed98de382549f50ab64ddfe8643b8c9b4b12a77cee1","version":"1852498","previousVersion":"22513","digest":"85wiAzekdiq6kzc4L8WqHgzYUToEQevYNQaVLR92Jcax"},{"type":"mutated","sender":"0xbb8a8269cf96ba2ec27dc9becd79836394dbe7946c7ac211928be4a0b1de66b9","owner":{"AddressOwner":"0xbb8a8269cf96ba2ec27dc9becd79836394dbe7946c7ac211928be4a0b1de66b9"},"objectType":"0x0000000000000000000000000000000000000000000000000000000000000002::coin::Coin<0x0000000000000000000000000000000000000000000000000000000000000002::sui::SUI>","objectId":"0xc587db1fbe680b769c1a562a09f2c871a087bafa542c7cb73db6064e2b791bdf","version":"1852498","previousVersion":"1852497","digest":"29FvSiZFhZ7z7HKEMFAza1H38m8hZogBCMD9TEmAtq3a"},{"type":"mutated","sender":"0xbb8a8269cf96ba2ec27dc9becd79836394dbe7946c7ac211928be4a0b1de66b9","owner":{"AddressOwner":"0xbb8a8269cf96ba2ec27dc9becd79836394dbe7946c7ac211928be4a0b1de66b9"},"objectType":"0x0000000000000000000000000000000000000000000000000000000000000002::coin::Coin<0x0000000000000000000000000000000000000000000000000000000000000002::sui::SUI>","objectId":"0xd65a04388a0b9378e87d0195e98bd0f9f7b460aac22ebd89fb3ba19e1759f414","version":"1852498","previousVersion":"23506","digest":"4s2P3qRZxQGpZhXz4EG5t8VDcrmzN3yaRqXF3AwQCnMs"},{"type":"mutated","sender":"0xbb8a8269cf96ba2ec27dc9becd79836394dbe7946c7ac211928be4a0b1de66b9","owner":{"AddressOwner":"0xbb8a8269cf96ba2ec27dc9becd79836394dbe7946c7ac211928be4a0b1de66b9"},"objectType":"0x0000000000000000000000000000000000000000000000000000000000000002::coin::Coin<0x0000000000000000000000000000000000000000000000000000000000000002::sui::SUI>","objectId":"0xe33119108d864f4169d7ed7fa963f51aaed7ef7107d8785cca237916e5079d7c","version":"1852498","previousVersion":"23425","digest":"7sTJRtJypM1pTGUwJxDcqqvoTzDMWKwnzrLmUoJwWPmy"}],"balanceChanges":[{"owner":{"AddressOwner":"0xbb8a8269cf96ba2ec27dc9becd79836394dbe7946c7ac211928be4a0b1de66b9"},"coinType":"0x0000000000000000000000000000000000000000000000000000000000000002::sui::SUI","amount":"-1079040"}],"timestampMs":"1683667808812","checkpoint":"2652859"}`,
				// grab the latest checkpoint
				`{"data":[{"epoch":"18","sequenceNumber":"1969067","digest":"Cji8yjGsk9Yg5sxUB4iWCQSiEjWHvVTzKMqBChvzNHjV","networkTotalTransactions":"4644633","previousDigest":"wG1Garuf7TkmtjdER4GFkiK9Q2id7ddnk4nAEXu7aDQ","epochRollingGasCostSummary":{"computationCost":"256732544695150","storageCost":"915446949600","storageRebate":"873308640996","nonRefundableStorageFee":"8821299404"},"timestampMs":"1683136992429","transactions":["22W8Rbz4NnYp2Jj7bSt15XrLBMH8Q6cW9HhB8WhAnUs4","3nfvfoJNytiWGDrrfJevXPfGjr1wsaWMB2yBKXDoqsAe","747otLtjqoiDuSKWSuvjCZWYGhNLrnV4fgGkXofNCHf9","A9kji7Sm4Sb4KWK1UudQ9SBymrtYWhCuccPfqpyRhaB6","G6ryjgaWHiRfSq7ULy9LVbp6P2JdphuBfuFRphvCNJHx"],"checkpointCommitments":[],"validatorSignature":"kaUtWOGEuc9DPgOueMBX5NaGDIjbCrXVVolb4jQCX6B02TCXFTAs9wtpRoOitmNv"}],"nextCursor":"1969067","hasNextPage":true}`,
				// grab the checkpoint for the tx
				`{"epoch":"18","sequenceNumber":"1953362","digest":"BHeEq9rUuc2kdh1k7vk4oN22oK7TtpTpVuaM32fs6UrB","networkTotalTransactions":"4594815","previousDigest":"QnwDCY5dJgpY9rRGbYseG8pQgdacHGtXgjC9TPVyKh8","epochRollingGasCostSummary":{"computationCost":"256697889685794","storageCost":"781589532800","storageRebate":"744825062520","nonRefundableStorageFee":"7523485480"},"timestampMs":"1683124849673","transactions":["UMeT2asZr2hsqT3n3vHmfE4UaNMjkVpoQvUpkRpfddd","3V6WU3ofLCd4pfujYnsEY95y4DxQhW8JM6m2Hy8XTbqg","7NHKPiWjTPM84ZBdi9othG3HCw2LUbxfFe183SWdTwhh","9i1bZ9mtMjMsjzvLRwP2YF9DUvLztiY1QbQVWfgGNrRJ","A5w6t3SkpHTUmCA14L7q9bnVQXsKWv5XRrKTP7wM7xCd","FohJy8o9qf3Qwag6wdWA6XCJeJEvjUmMzWgDhVPxMUnU","GAgWaGyeUoRSKGJQS3Mbs3PYeqfFLHtHpyXmHXGAiMAe","Hbo9qbYRj5XncBtpyzEro2aXYgWiThFEJTj8XmR5ArBg","J2Vkui75vgoLvCmNiREVKwpeTVPCq5EQ71i2ETahP6R9"],"checkpointCommitments":[],"validatorSignature":"k7orjPUoopsGMd6eR3JDie70DwppJ3t/F2BVMNDN06vX1FMDgscBuXf970TXoD9z"}`,
			},
			xc.TxInfo{
				BlockHash:     "BHeEq9rUuc2kdh1k7vk4oN22oK7TtpTpVuaM32fs6UrB",
				TxID:          "9aFfSrP7jvvteSS4q8L8RMC71NbfBeK1FK8aWcw8c8py",
				ExplorerURL:   "https://explorer.sui.io/txblock/9aFfSrP7jvvteSS4q8L8RMC71NbfBeK1FK8aWcw8c8py?network=devnet",
				From:          "0xbb8a8269cf96ba2ec27dc9becd79836394dbe7946c7ac211928be4a0b1de66b9",
				To:            "",
				BlockIndex:    2652859,
				BlockTime:     1683667808812,
				Confirmations: 15705,
				Sources: []*xc.TxInfoEndpoint{
					{
						Address:     "0xbb8a8269cf96ba2ec27dc9becd79836394dbe7946c7ac211928be4a0b1de66b9",
						Amount:      xc.NewAmountBlockchainFromStr("1079040"),
						Asset:       "SUI",
						NativeAsset: "SUI",
					},
				},
				Destinations: []*xc.TxInfoEndpoint{},
				Fee:          xc.NewAmountBlockchainFromStr("1079040"),
				Error:        "InsufficientCoinBalance in command 2",
				Status:       xc.TxStatusFailure,
			},
			"",
		},
		{
			"sui_token_usdc_deposit",
			[]string{
				// sui_getTransactionBlock
				`{"digest":"7uepPpd7LLqittQmViGyobWrTYv5RDZCeyh6Ja8ZJCWP","transaction":{"data":{"messageVersion":"v1","transaction":{"kind":"ProgrammableTransaction","inputs":[{"type":"object","objectType":"immOrOwnedObject","objectId":"0x5b72d2e0bb0a6a45421b3f474bb97aa3b63a1ce2a14991e68a1d96be4d2f19b5","version":"207749","digest":"EQZA7D3mRhLUNtsQpfe4QetzainUAncyP7Piwpx5FtEv"},{"type":"object","objectType":"immOrOwnedObject","objectId":"0xe5c09ef8b4ccc651dba40286eb212c75c9b196c680ef6417ef8fbe9b527ef67e","version":"207747","digest":"J83CQYgkKVqNk3MTdEKc3XimtuGxKuQGVYSVeQqXft7i"},{"type":"object","objectType":"immOrOwnedObject","objectId":"0x282fbf63a36defad50d93da60a63e1a70900c8bd9011403d6d5b5303e652dc62","version":"207746","digest":"7c1zLpCL9fbYB2oh7xhTLBm68zvMLca1tjaFUvekqAm2"},{"type":"object","objectType":"immOrOwnedObject","objectId":"0x928dbb9c08f8ac17ac362ec14fdf5354f46f8550648d3b4858d2febc16ad6c9e","version":"207748","digest":"EJMHGdHAnzJGuotYFxbciXBi9hjseBS5pq2UcDK3PU8a"},{"type":"pure","valueType":"u64","value":"1500000000000"},{"type":"pure","valueType":"address","value":"0xfe33ab3ab64a92088402fc22d850f04f0770d899695104447ffd93d7b83cfeb8"}],"transactions":[{"MergeCoins":[{"Input":0},[{"Input":1},{"Input":2},{"Input":3}]]},{"SplitCoins":[{"Input":0},[{"Input":4}]]},{"TransferObjects":[[{"Result":1}],{"Input":5}]}]},"sender":"0xbb8a8269cf96ba2ec27dc9becd79836394dbe7946c7ac211928be4a0b1de66b9","gasData":{"payment":[{"objectId":"0x34f60fd2a191693f538c75a224b66afb3e7f1ccaea898aff2bc442bed59ec162","version":207746,"digest":"xZ3CBnj1N7VrfZMDs6151tUnbWrqjEkDKAGJTGULZQH"}],"owner":"0xbb8a8269cf96ba2ec27dc9becd79836394dbe7946c7ac211928be4a0b1de66b9","price":"1000","budget":"2000000000"}},"txSignatures":["AF85s1fZeT7kUoElI75iyOuh+o16N6H2TJcMiQ6BTt2ftkf6/2bj9WmpdFbDmiSUiXDoST9BJaXCkzN/hDt09QZqA6rdJ6N1PDry1nZZFSjz2CCfM3uVBhY0ebxeYfZ+vQ=="]},"effects":{"messageVersion":"v1","status":{"status":"success"},"executedEpoch":"1","gasUsed":{"computationCost":"1000000","storageCost":"3663200","storageRebate":"6275016","nonRefundableStorageFee":"63384"},"modifiedAtVersions":[{"objectId":"0x34f60fd2a191693f538c75a224b66afb3e7f1ccaea898aff2bc442bed59ec162","sequenceNumber":"207746"},{"objectId":"0x5b72d2e0bb0a6a45421b3f474bb97aa3b63a1ce2a14991e68a1d96be4d2f19b5","sequenceNumber":"207749"},{"objectId":"0x282fbf63a36defad50d93da60a63e1a70900c8bd9011403d6d5b5303e652dc62","sequenceNumber":"207746"},{"objectId":"0x928dbb9c08f8ac17ac362ec14fdf5354f46f8550648d3b4858d2febc16ad6c9e","sequenceNumber":"207748"},{"objectId":"0xe5c09ef8b4ccc651dba40286eb212c75c9b196c680ef6417ef8fbe9b527ef67e","sequenceNumber":"207747"}],"transactionDigest":"7uepPpd7LLqittQmViGyobWrTYv5RDZCeyh6Ja8ZJCWP","created":[{"owner":{"AddressOwner":"0xfe33ab3ab64a92088402fc22d850f04f0770d899695104447ffd93d7b83cfeb8"},"reference":{"objectId":"0xc2d4733ab534984c5aede3466208b19739cf469182b64ff4f27d70d70cb19eae","version":207750,"digest":"6VEwDv3JRH7yFUxg9LNaB8SdndJDw2aC64gjKme5pHSn"}}],"mutated":[{"owner":{"AddressOwner":"0xbb8a8269cf96ba2ec27dc9becd79836394dbe7946c7ac211928be4a0b1de66b9"},"reference":{"objectId":"0x34f60fd2a191693f538c75a224b66afb3e7f1ccaea898aff2bc442bed59ec162","version":207750,"digest":"9GXKkHGN2iF8gp5knA5FpWTVbdxJpHbAYBquNJ6QWXMW"}},{"owner":{"AddressOwner":"0xbb8a8269cf96ba2ec27dc9becd79836394dbe7946c7ac211928be4a0b1de66b9"},"reference":{"objectId":"0x5b72d2e0bb0a6a45421b3f474bb97aa3b63a1ce2a14991e68a1d96be4d2f19b5","version":207750,"digest":"Garu9gSWgCKNfn2GEYXBqyd4o6ddDTNuXP6hZ9VuV7aN"}}],"deleted":[{"objectId":"0x282fbf63a36defad50d93da60a63e1a70900c8bd9011403d6d5b5303e652dc62","version":207750,"digest":"7gyGAp71YXQRoxmFBaHxofQXAipvgHyBKPyxmdSJxyvz"},{"objectId":"0x928dbb9c08f8ac17ac362ec14fdf5354f46f8550648d3b4858d2febc16ad6c9e","version":207750,"digest":"7gyGAp71YXQRoxmFBaHxofQXAipvgHyBKPyxmdSJxyvz"},{"objectId":"0xe5c09ef8b4ccc651dba40286eb212c75c9b196c680ef6417ef8fbe9b527ef67e","version":207750,"digest":"7gyGAp71YXQRoxmFBaHxofQXAipvgHyBKPyxmdSJxyvz"}],"gasObject":{"owner":{"AddressOwner":"0xbb8a8269cf96ba2ec27dc9becd79836394dbe7946c7ac211928be4a0b1de66b9"},"reference":{"objectId":"0x34f60fd2a191693f538c75a224b66afb3e7f1ccaea898aff2bc442bed59ec162","version":207750,"digest":"9GXKkHGN2iF8gp5knA5FpWTVbdxJpHbAYBquNJ6QWXMW"}},"dependencies":["648EuGHEB2dmLsVgCU6NS7HEts67A6UCu2kGGFHverpB","7aoUNiaP3WNsm1vb8URx9AEKS2bcqzh9wA1jhJb3oa5K","9AC68uvDUXu4s2nJd8QpTCEsbNgnM1K63Cbtn7MaA5Zr","HQ6jnXza2rHGyqcBXTfG4NyNGhyNZksNiWhKXfGnJafY"]},"events":[],"objectChanges":[{"type":"mutated","sender":"0xbb8a8269cf96ba2ec27dc9becd79836394dbe7946c7ac211928be4a0b1de66b9","owner":{"AddressOwner":"0xbb8a8269cf96ba2ec27dc9becd79836394dbe7946c7ac211928be4a0b1de66b9"},"objectType":"0x0000000000000000000000000000000000000000000000000000000000000002::coin::Coin<0x0000000000000000000000000000000000000000000000000000000000000002::sui::SUI>","objectId":"0x34f60fd2a191693f538c75a224b66afb3e7f1ccaea898aff2bc442bed59ec162","version":"207750","previousVersion":"207746","digest":"9GXKkHGN2iF8gp5knA5FpWTVbdxJpHbAYBquNJ6QWXMW"},{"type":"mutated","sender":"0xbb8a8269cf96ba2ec27dc9becd79836394dbe7946c7ac211928be4a0b1de66b9","owner":{"AddressOwner":"0xbb8a8269cf96ba2ec27dc9becd79836394dbe7946c7ac211928be4a0b1de66b9"},"objectType":"0x0000000000000000000000000000000000000000000000000000000000000002::coin::Coin<0x3821e4ae13d37a1c55a03a86eab613450c1302e6b4df461e1c79bdf8381dde47::iusdc::IUSDC>","objectId":"0x5b72d2e0bb0a6a45421b3f474bb97aa3b63a1ce2a14991e68a1d96be4d2f19b5","version":"207750","previousVersion":"207749","digest":"Garu9gSWgCKNfn2GEYXBqyd4o6ddDTNuXP6hZ9VuV7aN"},{"type":"created","sender":"0xbb8a8269cf96ba2ec27dc9becd79836394dbe7946c7ac211928be4a0b1de66b9","owner":{"AddressOwner":"0xfe33ab3ab64a92088402fc22d850f04f0770d899695104447ffd93d7b83cfeb8"},"objectType":"0x0000000000000000000000000000000000000000000000000000000000000002::coin::Coin<0x3821e4ae13d37a1c55a03a86eab613450c1302e6b4df461e1c79bdf8381dde47::iusdc::IUSDC>","objectId":"0xc2d4733ab534984c5aede3466208b19739cf469182b64ff4f27d70d70cb19eae","version":"207750","digest":"6VEwDv3JRH7yFUxg9LNaB8SdndJDw2aC64gjKme5pHSn"}],"balanceChanges":[{"owner":{"AddressOwner":"0xbb8a8269cf96ba2ec27dc9becd79836394dbe7946c7ac211928be4a0b1de66b9"},"coinType":"0x0000000000000000000000000000000000000000000000000000000000000002::sui::SUI","amount":"1611816"},{"owner":{"AddressOwner":"0xbb8a8269cf96ba2ec27dc9becd79836394dbe7946c7ac211928be4a0b1de66b9"},"coinType":"0x3821e4ae13d37a1c55a03a86eab613450c1302e6b4df461e1c79bdf8381dde47::iusdc::IUSDC","amount":"-1500000000000"},{"owner":{"AddressOwner":"0xfe33ab3ab64a92088402fc22d850f04f0770d899695104447ffd93d7b83cfeb8"},"coinType":"0x3821e4ae13d37a1c55a03a86eab613450c1302e6b4df461e1c79bdf8381dde47::iusdc::IUSDC","amount":"1500000000000"}],"timestampMs":"1683903105617","checkpoint":"213114"}`,
				// sui_getCheckpoints
				`{"data":[{"epoch":"1","sequenceNumber":"221952","digest":"8KTk1Bhd55Shki8dHNeirwcx3wt7R4nseZLsCKibaBT9","networkTotalTransactions":"1747735","previousDigest":"Foah2YvwaXrgFWid41NLdVZrUHgvYE7LWKrgGtBUDHSm","epochRollingGasCostSummary":{"computationCost":"851240304460","storageCost":"1968707883600","storageRebate":"1895302770372","nonRefundableStorageFee":"19144472428"},"timestampMs":"1683909941286","transactions":["WZgXY77rYkat7HRRVHF1TZ3YSSX9qgJ34pDdcQ5FBFP","qsP5dbGkk94ZhLJLLjBZSW6T4Df2SGWjXhyREQWXbCg","3iuvvr2NFh1b7CeGu23QYnuAb4927QhH53Jz12sWihjF","EaHUFdPMRDLDG57CiL1zqQbGNmk8gWuXYT2hFDTXAZw8"],"checkpointCommitments":[],"validatorSignature":"gLdMOYtmxpSc2W6Q0QfjnLA3y9BFQCsYZRI5o2UlAwTFLhbwhY7a5UbKmO3BmCpw"}],"nextCursor":"221952","hasNextPage":true}`,
				// sui_getCheckpoint
				`{"epoch":"1","sequenceNumber":"213114","digest":"8S4qJohEmkdebSt2t9nHmtWquK39KZVJK5bKrARYNBYP","networkTotalTransactions":"1673546","previousDigest":"93hv9hGLRfRZXt78mV6fytQYWXM9grW81H72eyqJYDco","epochRollingGasCostSummary":{"computationCost":"785450001460","storageCost":"1810842150400","storageRebate":"1744678497672","nonRefundableStorageFee":"17623015128"},"timestampMs":"1683903105617","transactions":["26jr9a6RR8G9TKpXySwnArwMTtqGSoN5BPddp9tsuKLo","2fy33DVLoKw1Qfe4PBEpyYZyTDSsy4UTVLZvUsEASEZt","3SV6qmb4wXKYWTnivjHQL8xH41MRq7Vj7hsZsspWfERf","4d3MhqoRk2gESQUoVNyNDWQxGVLoTWbFxhdhjtsNcEPb","7uepPpd7LLqittQmViGyobWrTYv5RDZCeyh6Ja8ZJCWP","8xUKJfoJKyC8GbMPnNM4iWjEER5CEsru6XQ4adax3HBN","DrJegZcM4kWEMyEcJp18UiymaqFnWZBpTUhPGSKW2gir","EKrnLwFDrtYVcmPC8pFwZAJQQC4e1LhorPbHa3EG6E2H","G9tReX5eo1REF1Vxrn4yxGd4WM7PsVZ1oMRQBnNUnYWh"],"checkpointCommitments":[],"validatorSignature":"iEOUB8jhgUky5jAE45hhAuQ0o4lUjphOTc79jtTjBo/aajRlO2UMIuzXUXAgBwyq"}`,
			},
			xc.TxInfo{
				BlockHash:       "8S4qJohEmkdebSt2t9nHmtWquK39KZVJK5bKrARYNBYP",
				TxID:            "7uepPpd7LLqittQmViGyobWrTYv5RDZCeyh6Ja8ZJCWP",
				ExplorerURL:     "https://explorer.sui.io/txblock/7uepPpd7LLqittQmViGyobWrTYv5RDZCeyh6Ja8ZJCWP?network=devnet",
				From:            "0xbb8a8269cf96ba2ec27dc9becd79836394dbe7946c7ac211928be4a0b1de66b9",
				To:              "0xfe33ab3ab64a92088402fc22d850f04f0770d899695104447ffd93d7b83cfeb8",
				BlockIndex:      213114,
				BlockTime:       1683903105617,
				Confirmations:   8838,
				ContractAddress: "0x3821e4ae13d37a1c55a03a86eab613450c1302e6b4df461e1c79bdf8381dde47::iusdc::IUSDC",
				Sources: []*xc.TxInfoEndpoint{
					{
						Address:         "0xbb8a8269cf96ba2ec27dc9becd79836394dbe7946c7ac211928be4a0b1de66b9",
						Amount:          xc.NewAmountBlockchainFromStr("1500000000000"),
						NativeAsset:     "SUI",
						ContractAddress: "0x3821e4ae13d37a1c55a03a86eab613450c1302e6b4df461e1c79bdf8381dde47::iusdc::IUSDC",
					},
				},
				Destinations: []*xc.TxInfoEndpoint{
					{
						Address:     "0xbb8a8269cf96ba2ec27dc9becd79836394dbe7946c7ac211928be4a0b1de66b9",
						Amount:      xc.NewAmountBlockchainFromStr("1611816"),
						Asset:       "SUI",
						NativeAsset: "SUI",
					},
					{
						Address:         "0xfe33ab3ab64a92088402fc22d850f04f0770d899695104447ffd93d7b83cfeb8",
						Amount:          xc.NewAmountBlockchainFromStr("1500000000000"),
						NativeAsset:     "SUI",
						ContractAddress: "0x3821e4ae13d37a1c55a03a86eab613450c1302e6b4df461e1c79bdf8381dde47::iusdc::IUSDC",
					},
				},
				// believe it or not SUI will rebate (pay you) fee if you merge enough coins
				Fee:    xc.NewAmountBlockchainFromStr("-1611816"),
				Amount: xc.NewAmountBlockchainFromStr("1500000000000"),
			},
			"",
		},
	}

	for _, v := range vectors {
		fmt.Println("testing ", v.name)
		server, close := test.MockJSONRPC(&s.Suite, v.resp)
		defer close()
		asset := &xc.AssetConfig{NativeAsset: xc.SUI, Net: "devnet", URL: server.URL}

		asset.URL = server.URL
		client, _ := NewClient(asset)
		txInfo, err := client.FetchTxInfo(s.Ctx, xc.TxHash(""))

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

func coinObject(contract string, id string, digest string, amount uint64, version int) *types.Coin {
	coinId, err := decodeHex(id)
	if err != nil {
		panic(err)
	}
	var bal types.SafeSuiBigInt[uint64] = types.NewSafeSuiBigInt(amount)
	return &types.Coin{
		CoinType:     contract,
		CoinObjectId: coinId,
		Digest:       digest,
		Balance:      bal,
		// fixed for test
		Version: decimal.NewFromInt(int64(version)),
	}
}
func suiCoin(id string, digest string, amount uint64, version int) *types.Coin {
	return coinObject("0x2::sui::SUI", id, digest, amount, version)
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
		token     *xc.TokenAssetConfig
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
			nil,
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
				&bcs.CallArg__Object{Value: mustCoinToObject(suiCoin("0x1cdc19f7751451412d090632bb1ca2c845a9c8f6cd8798d99d304571cfea1ca6", "u6uSbWNMxkRkCqkjSTbsMeWMYB2VK7pbAo6vFoaMzSo", 2001904720, 1852477))},
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
			nil,
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
				&bcs.CallArg__Object{Value: mustCoinToObject(suiCoin("0xc587db1fbe680b769c1a562a09f2c871a087bafa542c7cb73db6064e2b791bdf", "7Y2zjQxn2wj5jhrvS5NBKCFJDzWHZ4UMG7XJNNioNgTS", 1897841920, 1852477))},
				&bcs.CallArg__Object{Value: mustCoinToObject(suiCoin("0x87bae5d7376e857106f7908eab6f7106ea3f7c2a1b3349f99925bb12631b1ff0", "9GeMg1yw4J9ck62XR3KHXi72kfVVeuqfAcK5rL3hRdVK", 1500000000, 1852477))},
				&bcs.CallArg__Object{Value: mustCoinToObject(suiCoin("0x5c22194a002befba3d34d26036d4c440f86099ac4cb9b8aaeca22fb379229237", "3t7sWDqfyKvbGxtnS1GwMV2kgdasLgmDTRJ7MHhoyCz3", 1300000000, 1852477))},
				&bcs.CallArg__Object{Value: mustCoinToObject(suiCoin("0xa47d2121ef5ca77d83723d72a6b70c3bce15a2f438294f2d0fbcb530ab5d0b27", "EYMFpVaEcfdv8kv1hxZz8y884z2fhQJt8d3G1zKBYf6m", 1200000000, 1852477))},
				&bcs.CallArg__Object{Value: mustCoinToObject(suiCoin("0x3a705927f067d86aedf19bbe84bb15cedbf613c3ac5c88b8508d8ce3f9bbbf7c", "GreYy8apDQHR7zwsZLHZ6bfQAiQ12xE4TNcm4vznpNUM", 1000000000, 1852477))},
				&bcs.CallArg__Object{Value: mustCoinToObject(suiCoin("0xeb8b3e2e9c446f25f29fee61c43583b9d7fbfed600a83219ca99d1ee681ac958", "DkV4WuN3ZPLHfF87otc23aVzwfbJQWyP171YhEyJQG5Q", 1000000000, 1852477))},
				&bcs.CallArg__Object{Value: mustCoinToObject(suiCoin("0xfb889571ed135b9bd1c1fd7d00d69694305bb74113efa087c18d6444528da091", "8XS5rHkHwoYtSq2Fg49NbZsxRT6tM5n31wx3bKMjpRDy", 900000000, 1852477))},
				&bcs.CallArg__Object{Value: mustCoinToObject(suiCoin("0x1cdc19f7751451412d090632bb1ca2c845a9c8f6cd8798d99d304571cfea1ca6", "SNiJ8aV9rerhbVTwZikSAWVgJPhx9jxaPXdGcfeYut9", 45035120, 1852477))},
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
			nil,
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
			nil,
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

		// Test with sending almost all of the balance and expect that the gas_budget is reduced to the remainder accordingly.
		{
			"Test_gas_budget_remainder",
			"95.0",
			nil,
			[]string{
				// get coins
				` {"data":[
					{"coinType":"0x2::sui::SUI","coinObjectId":"0x3150377d1db0395abfd3b19cfeca94eaf5987a12b95a0aab431195e77399f092","version":"1852505","digest":"7xvFhTk5r3RCLQPcybUeTuwAUKAy8ozXN5EbKsnvp9Qb","balance":"10000000000","previousTransaction":"APAAcvLGmcXFTjMwv7iAJ2hwETQyQFDkfVzs4tEyE43F"},
					{"coinType":"0x2::sui::SUI","coinObjectId":"0x5a23ee6e22faa7017b11ad24e7c8ced1d33465cfd06656bc028eb21c6f4cad97","version":"1852505","digest":"67APB8hNkhBWmARr49RRXwQGgCC3A8VMxcLKbftUiYQF","balance":"10000000000","previousTransaction":"APAAcvLGmcXFTjMwv7iAJ2hwETQyQFDkfVzs4tEyE43F"},
					{"coinType":"0x2::sui::SUI","coinObjectId":"0x7d1775d7f791554b25933fde2b91d578ddc2874d1f402b55a7b8f5fb270b845d","version":"1852505","digest":"GP2Uc7u6uCa8QnxQz2kkiPnz3hBaS3z17vcuVyLYGDPh","balance":"10000000000","previousTransaction":"APAAcvLGmcXFTjMwv7iAJ2hwETQyQFDkfVzs4tEyE43F"},
					{"coinType":"0x2::sui::SUI","coinObjectId":"0x8192d5c2b5722c60866761927d5a0737cd55d0c2b1150eabf818253795b38998","version":"1852505","digest":"5MYcnxjPkzxG3bwPaLkDKG9snzeZVLFwQ25pePL1vDH7","balance":"1997992240","previousTransaction":"APAAcvLGmcXFTjMwv7iAJ2hwETQyQFDkfVzs4tEyE43F"},
					{"coinType":"0x2::sui::SUI","coinObjectId":"0x92e60b8b39e5f3ecd57f6ed98de382549f50ab64ddfe8643b8c9b4b12a77cee1","version":"1852505","digest":"HveRBbXj1nLo6rQioHAKV9TtmqnQvvuuKfG5BMDbM3TX","balance":"10000000000","previousTransaction":"APAAcvLGmcXFTjMwv7iAJ2hwETQyQFDkfVzs4tEyE43F"},
					{"coinType":"0x2::sui::SUI","coinObjectId":"0xc587db1fbe680b769c1a562a09f2c871a087bafa542c7cb73db6064e2b791bdf","version":"1852505","digest":"ENUoMU2gFeLZEPxxQxdMwrNGtvJLFry2HvgXVCnCB1k9","balance":"33827025240","previousTransaction":"APAAcvLGmcXFTjMwv7iAJ2hwETQyQFDkfVzs4tEyE43F"},
					{"coinType":"0x2::sui::SUI","coinObjectId":"0xd65a04388a0b9378e87d0195e98bd0f9f7b460aac22ebd89fb3ba19e1759f414","version":"1852505","digest":"DSoevdHtFsNV8M1rCeFMW4GU8sEczqtJYJnm921Vv7gF","balance":"10000000000","previousTransaction":"APAAcvLGmcXFTjMwv7iAJ2hwETQyQFDkfVzs4tEyE43F"},
					{"coinType":"0x2::sui::SUI","coinObjectId":"0xe33119108d864f4169d7ed7fa963f51aaed7ef7107d8785cca237916e5079d7c","version":"1852505","digest":"FUzqDFoq73G1WecQacFv2WmX83ezk16DGZkkzLTwCbvJ","balance":"10000000000","previousTransaction":"APAAcvLGmcXFTjMwv7iAJ2hwETQyQFDkfVzs4tEyE43F"}],"nextCursor":"0xe33119108d864f4169d7ed7fa963f51aaed7ef7107d8785cca237916e5079d7c","hasNextPage":false}`,
				// get checkpoint
				`{"data":[{"epoch":"21","sequenceNumber":"2206686","digest":"HtsAAgd1ajMR8qMocnNF6XbAtiBHrxdauGhWtXqKouF3","networkTotalTransactions":"5164703","previousDigest":"H8oYvb73KoG7TWXpw4JPy2qZk7ddvHY3rYQ8kHcNmcua","epochRollingGasCostSummary":{"computationCost":"130960164300","storageCost":"499151462400","storageRebate":"422717709348","nonRefundableStorageFee":"4269875852"},"timestampMs":"1683320609521","transactions":["3yVjcHqKwLN8K8TrZZZMpMUp4VSGg4LRp4uuzvvzzrFD","Cv2NH6zJiRJMtPMzxzZABgDpBfNmb9eniWW9t5v2kPtz","GJaDtfzHap6V8ARdQTstkJm7PiWsEXWkUapXHA2nbmbD"],"checkpointCommitments":[],"validatorSignature":"i3aT5RVtIOvX0pEc/HU+xFTHbw2zV5SdT7q5n6GfS+e85CtkC8qqseeK2Hx9Nhia"}],"nextCursor":"2206686","hasNextPage":true}`,
				//reference gas
				"1000",
				// submit tx
				`{"digest":"5NVoZeHas2s7go1wiSMXtM2g1KitDwu2eksvEt1jRcWj","confirmedLocalExecution":true}`,
			},
			// split, merge, split, transfer
			[]bcs.CallArg{
				// remainder split
				u64ToPure(33002007760),
				// merged coins after gas coin (sorted by value)
				&bcs.CallArg__Object{Value: mustCoinToObject(suiCoin("0x3150377d1db0395abfd3b19cfeca94eaf5987a12b95a0aab431195e77399f092", "7xvFhTk5r3RCLQPcybUeTuwAUKAy8ozXN5EbKsnvp9Qb", 10000000000, 1852505))},
				&bcs.CallArg__Object{Value: mustCoinToObject(suiCoin("0x5a23ee6e22faa7017b11ad24e7c8ced1d33465cfd06656bc028eb21c6f4cad97", "67APB8hNkhBWmARr49RRXwQGgCC3A8VMxcLKbftUiYQF", 10000000000, 1852505))},
				&bcs.CallArg__Object{Value: mustCoinToObject(suiCoin("0x7d1775d7f791554b25933fde2b91d578ddc2874d1f402b55a7b8f5fb270b845d", "GP2Uc7u6uCa8QnxQz2kkiPnz3hBaS3z17vcuVyLYGDPh", 10000000000, 1852505))},
				&bcs.CallArg__Object{Value: mustCoinToObject(suiCoin("0x92e60b8b39e5f3ecd57f6ed98de382549f50ab64ddfe8643b8c9b4b12a77cee1", "HveRBbXj1nLo6rQioHAKV9TtmqnQvvuuKfG5BMDbM3TX", 10000000000, 1852505))},
				&bcs.CallArg__Object{Value: mustCoinToObject(suiCoin("0xd65a04388a0b9378e87d0195e98bd0f9f7b460aac22ebd89fb3ba19e1759f414", "DSoevdHtFsNV8M1rCeFMW4GU8sEczqtJYJnm921Vv7gF", 10000000000, 1852505))},
				&bcs.CallArg__Object{Value: mustCoinToObject(suiCoin("0xe33119108d864f4169d7ed7fa963f51aaed7ef7107d8785cca237916e5079d7c", "FUzqDFoq73G1WecQacFv2WmX83ezk16DGZkkzLTwCbvJ", 10000000000, 1852505))},
				&bcs.CallArg__Object{Value: mustCoinToObject(suiCoin("0x8192d5c2b5722c60866761927d5a0737cd55d0c2b1150eabf818253795b38998", "5MYcnxjPkzxG3bwPaLkDKG9snzeZVLFwQ25pePL1vDH7", 1997992240, 1852505))},
				// split amt (transfer amount)
				u64ToPure(95_000_000_000),
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
					},
				},
				&bcs.Command__SplitCoins{
					Field0: ArgumentInput(1),
					Field1: []bcs.Argument{
						ArgumentInput(8),
					},
				},
				&bcs.Command__TransferObjects{
					Field0: []bcs.Argument{
						ArgumentResult(2),
					},
					Field1: ArgumentInput(9),
				},
			},
			825017480,
			nil,
		},

		// Test sending a usdc token when there's 1 usdc coin
		{
			"Test_1_usdc_coin",
			"3.0",
			&xc.TokenAssetConfig{
				Asset:    "USDC",
				Decimals: 9,
				Contract: "0x3821e4ae13d37a1c55a03a86eab613450c1302e6b4df461e1c79bdf8381dde47::iusdc::IUSDC",
				Type:     xc.AssetTypeToken,
			},
			[]string{
				// suix_getCoins (token)
				`{"data":[
					{"coinType":"0x3821e4ae13d37a1c55a03a86eab613450c1302e6b4df461e1c79bdf8381dde47::iusdc::IUSDC","coinObjectId":"0x70136e3f64bea493b5c73e5e2fa03beb36678fc2f6df471c17e7ae8cd34baac5","version":"207737","digest":"D7N9aUM4Cn6ukSFnbDUu2VACa7Q6FLzyj9znSL3W4YHt","balance":"2516489093104","previousTransaction":"HhCq8usSNk4DTmaKCR17AZC4dbwEw1rSFFuFTchm61iF"}
				],"nextCursor":"0x70136e3f64bea493b5c73e5e2fa03beb36678fc2f6df471c17e7ae8cd34baac5","hasNextPage":false}`,
				// sui_getCheckpoints
				`{"data":[{"epoch":"21","sequenceNumber":"210947","digest":"8S1Rd5KsqGH2JEoTnJHroz5RDD5oPxAxerXKBomJ4h8y","networkTotalTransactions":"1654947","previousDigest":"FDpz4PTTBu32XnHjo5JQ4gQfpUrUPivcWQXWvP8b1caH","epochRollingGasCostSummary":{"computationCost":"767194001460","storageCost":"1767919045200","storageRebate":"1704412479924","nonRefundableStorageFee":"17216287676"},"timestampMs":"1683901433754","transactions":["2XoiF5ueWohVK1sYKgVJZN1GjB4fXNrbhJZqDt2CPAWq","2c5CvBw57EX2kLXAqzqA1JVpPSb9RcZZ3wW6FEvaCRRs","4YsGHEsMFeMUNWALBuECduwFCAkb2VFjx993fetvLLsS","4d8NLyNCk2XQiiA4T7t6ah9NC9eWnHhXhycmSEgS9ivG","BHdwEXSpi1HfyAJH5HuducZ43iRuKgEm8Ukm4j3srQDt","Bhm8dMMtHxJqaKEPogBQKiAKYTsfTzTjURXwvqa3mGbz","DqFPyPoKGv7TwTaThrCUqQPCUuykUfzRGsE1QdqKjWbx","EnPgKYQpmbzzyke742XEH2ncitBnJ5d968yMXRdEiCNw"],"checkpointCommitments":[],"validatorSignature":"jgxvZ1V0HlIWovRMHNYppUsNyq6EKgt7KwZRDJiOspO3ScIs9Z53kcJ8HJUiFk49"}],"nextCursor":"210947","hasNextPage":true}`,
				// suix_getCoins (gas)
				`{"data":[{"coinType":"0x2::sui::SUI","coinObjectId":"0x34f60fd2a191693f538c75a224b66afb3e7f1ccaea898aff2bc442bed59ec162","version":"207737","digest":"ACRr1x7hC7CRfPu9a7gecLkStHdu6eGNK81SqvCGVJM1","balance":"6892967516","previousTransaction":"HhCq8usSNk4DTmaKCR17AZC4dbwEw1rSFFuFTchm61iF"}],"nextCursor":"0x34f60fd2a191693f538c75a224b66afb3e7f1ccaea898aff2bc442bed59ec162","hasNextPage":false}`,
				// suix_getReferenceGasPrice
				`"1000"`,
				// sui_executeTransactionBlock
				`{"digest":"HAKa4YPcFYT4M1LkYvUE6u8nLjgJX6cwmyR4LHGNXqYe","confirmedLocalExecution":true}`,
			},
			// split, transfer
			[]bcs.CallArg{
				&bcs.CallArg__Object{Value: mustCoinToObject(coinObject(
					"0x3821e4ae13d37a1c55a03a86eab613450c1302e6b4df461e1c79bdf8381dde47::iusdc::IUSDC",
					"0x70136e3f64bea493b5c73e5e2fa03beb36678fc2f6df471c17e7ae8cd34baac5",
					"D7N9aUM4Cn6ukSFnbDUu2VACa7Q6FLzyj9znSL3W4YHt",
					10000000000,
					207737,
				))},
				// split amt (transfer amount)
				u64ToPure(3_000_000_000),
				// destination address
				mustHexToPure(to),
			},
			[]bcs.Command{
				&bcs.Command__SplitCoins{
					Field0: ArgumentInput(0),
					Field1: []bcs.Argument{ArgumentInput(1)},
				},
				&bcs.Command__TransferObjects{
					Field0: []bcs.Argument{
						ArgumentResult(0),
					},
					Field1: ArgumentInput(2),
				},
			},
			2_000_000_000,
			nil,
		},

		// Test sending a usdc token when there's multiple usdc coin
		{
			"Test_many_usdc_coin",
			"1500.0",
			&xc.TokenAssetConfig{
				Asset:    "USDC",
				Decimals: 9,
				Contract: "0x3821e4ae13d37a1c55a03a86eab613450c1302e6b4df461e1c79bdf8381dde47::iusdc::IUSDC",
				Type:     xc.AssetTypeToken,
			},
			[]string{
				// suix_getCoins
				`{"data":[
					{"coinType":"0x3821e4ae13d37a1c55a03a86eab613450c1302e6b4df461e1c79bdf8381dde47::iusdc::IUSDC","coinObjectId":"0x282fbf63a36defad50d93da60a63e1a70900c8bd9011403d6d5b5303e652dc62","version":"207746","digest":"7c1zLpCL9fbYB2oh7xhTLBm68zvMLca1tjaFUvekqAm2","balance":"18489093104","previousTransaction":"648EuGHEB2dmLsVgCU6NS7HEts67A6UCu2kGGFHverpB"},
					{"coinType":"0x3821e4ae13d37a1c55a03a86eab613450c1302e6b4df461e1c79bdf8381dde47::iusdc::IUSDC","coinObjectId":"0x5b72d2e0bb0a6a45421b3f474bb97aa3b63a1ce2a14991e68a1d96be4d2f19b5","version":"207749","digest":"EQZA7D3mRhLUNtsQpfe4QetzainUAncyP7Piwpx5FtEv","balance":"1000000000000","previousTransaction":"7aoUNiaP3WNsm1vb8URx9AEKS2bcqzh9wA1jhJb3oa5K"},
					{"coinType":"0x3821e4ae13d37a1c55a03a86eab613450c1302e6b4df461e1c79bdf8381dde47::iusdc::IUSDC","coinObjectId":"0x928dbb9c08f8ac17ac362ec14fdf5354f46f8550648d3b4858d2febc16ad6c9e","version":"207748","digest":"EJMHGdHAnzJGuotYFxbciXBi9hjseBS5pq2UcDK3PU8a","balance":"1000000000","previousTransaction":"9AC68uvDUXu4s2nJd8QpTCEsbNgnM1K63Cbtn7MaA5Zr"},
					{"coinType":"0x3821e4ae13d37a1c55a03a86eab613450c1302e6b4df461e1c79bdf8381dde47::iusdc::IUSDC","coinObjectId":"0xe5c09ef8b4ccc651dba40286eb212c75c9b196c680ef6417ef8fbe9b527ef67e","version":"207747","digest":"J83CQYgkKVqNk3MTdEKc3XimtuGxKuQGVYSVeQqXft7i","balance":"500000000000","previousTransaction":"HQ6jnXza2rHGyqcBXTfG4NyNGhyNZksNiWhKXfGnJafY"}
					],"nextCursor":"0xe5c09ef8b4ccc651dba40286eb212c75c9b196c680ef6417ef8fbe9b527ef67e","hasNextPage":false}`,
				// sui_getCheckpoints
				`{"data":[{"epoch":"21","sequenceNumber":"213110","digest":"7s32SP7EX4G6kwz6oF92j6aTH536tQctu8QZtaWnMjLh","networkTotalTransactions":"1673509","previousDigest":"FijCctQxLyzJRerPBPKoHzpSgk6ZXcVJ4jj5oGmWPgo2","epochRollingGasCostSummary":{"computationCost":"785413001460","storageCost":"1810760678400","storageRebate":"1744598088684","nonRefundableStorageFee":"17622202916"},"timestampMs":"1683903102105","transactions":["HiTCaHRPiQQ2JzPd74u8ubFA9qTqdzFBQNe4Kk8mjkY","3hSj8MncixFy4wNpFRGsPuGGPNgLNfHvws4sZ4iqqcQA","5Cz9paQe5tUSanZtJ22w6kyhSeovdndLLgsmpTK1PShe","5cDUSNmx1Wz1pcW33J2hHqtptFdQrAV6JeHJjvRhwBh3","7AZJHhWcfyB8ddgHEg4JCMXwfcKkeEZGDmW4tNkMtpa6","E66oHc4eS3yXVMAHwdUooVHnDRqg9vjKgWjhnUBR6GZp","FFPoHULMzkFr2SafckJ5yxcRL4UX493BJNLpjRjRG9E8","G5cXK2kYiyf8GjmTuGvEi4m3r991Go2kQz5jPGiJQfbx","HEsfY9HEC5rAJkq5ofFMRrpW9w9dF82Jv29azn1FqjaX","Hq1odW3nkVH6FeZ8Baw9NRLq43TnQP6V2ojVhmuDa21p"],"checkpointCommitments":[],"validatorSignature":"o026PhUt43ZhM9EAvHHGYyn/xAdsCWlFp0Q6KqcRiA8Md1f9U+vVRGk8JUro6P1m"}],"nextCursor":"213110","hasNextPage":true}`,
				// suix_getCoins
				`{"data":[{"coinType":"0x2::sui::SUI","coinObjectId":"0x34f60fd2a191693f538c75a224b66afb3e7f1ccaea898aff2bc442bed59ec162","version":"207746","digest":"xZ3CBnj1N7VrfZMDs6151tUnbWrqjEkDKAGJTGULZQH","balance":"3890508188","previousTransaction":"648EuGHEB2dmLsVgCU6NS7HEts67A6UCu2kGGFHverpB"}],"nextCursor":"0x34f60fd2a191693f538c75a224b66afb3e7f1ccaea898aff2bc442bed59ec162","hasNextPage":false}`,
				// suix_getReferenceGasPrice
				`"1000"`,
				// sui_executeTransactionBlock
				`{"digest":"7uepPpd7LLqittQmViGyobWrTYv5RDZCeyh6Ja8ZJCWP","confirmedLocalExecution":true}`,
			},
			// merge, split, transfer
			[]bcs.CallArg{
				&bcs.CallArg__Object{Value: mustCoinToObject(coinObject(
					"0x3821e4ae13d37a1c55a03a86eab613450c1302e6b4df461e1c79bdf8381dde47::iusdc::IUSDC",
					"0x5b72d2e0bb0a6a45421b3f474bb97aa3b63a1ce2a14991e68a1d96be4d2f19b5",
					"EQZA7D3mRhLUNtsQpfe4QetzainUAncyP7Piwpx5FtEv",
					1000000000000,
					207749,
				))},
				&bcs.CallArg__Object{Value: mustCoinToObject(coinObject(
					"0x3821e4ae13d37a1c55a03a86eab613450c1302e6b4df461e1c79bdf8381dde47::iusdc::IUSDC",
					"0xe5c09ef8b4ccc651dba40286eb212c75c9b196c680ef6417ef8fbe9b527ef67e",
					"J83CQYgkKVqNk3MTdEKc3XimtuGxKuQGVYSVeQqXft7i",
					500000000000,
					207747,
				))},

				&bcs.CallArg__Object{Value: mustCoinToObject(coinObject(
					"0x3821e4ae13d37a1c55a03a86eab613450c1302e6b4df461e1c79bdf8381dde47::iusdc::IUSDC",
					"0x282fbf63a36defad50d93da60a63e1a70900c8bd9011403d6d5b5303e652dc62",
					"7c1zLpCL9fbYB2oh7xhTLBm68zvMLca1tjaFUvekqAm2",
					18489093104,
					207746,
				))},
				&bcs.CallArg__Object{Value: mustCoinToObject(coinObject(
					"0x3821e4ae13d37a1c55a03a86eab613450c1302e6b4df461e1c79bdf8381dde47::iusdc::IUSDC",
					"0x928dbb9c08f8ac17ac362ec14fdf5354f46f8550648d3b4858d2febc16ad6c9e",
					"EJMHGdHAnzJGuotYFxbciXBi9hjseBS5pq2UcDK3PU8a",
					1000000000,
					207748,
				))},

				// split amt (transfer amount)
				u64ToPure(1_500_000_000_000),
				// destination address
				mustHexToPure(to),
			},
			[]bcs.Command{
				&bcs.Command__MergeCoins{
					Field0: ArgumentInput(0),
					Field1: []bcs.Argument{
						ArgumentInput(1),
						ArgumentInput(2),
						ArgumentInput(3),
					},
				},
				&bcs.Command__SplitCoins{
					Field0: ArgumentInput(0),
					Field1: []bcs.Argument{ArgumentInput(4)},
				},
				&bcs.Command__TransferObjects{
					Field0: []bcs.Argument{
						ArgumentResult(1),
					},
					Field1: ArgumentInput(5),
				},
			},
			2_000_000_000,
			nil,
		},
	}

	for _, v := range vectors {
		fmt.Println("Running ", v.name)
		server, close := test.MockJSONRPC(&s.Suite, v.resp)
		defer close()
		nativeAsset := &xc.NativeAssetConfig{NativeAsset: xc.SUI, Net: "devnet", URL: server.URL}
		nativeAsset.URL = server.URL
		var asset xc.ITask = nativeAsset
		if v.token != nil {
			v.token.NativeAssetConfig = nativeAsset
			asset = v.token
			// asset =
		}
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
			fmt.Println("checking command", i)
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
