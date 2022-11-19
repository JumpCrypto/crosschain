package solana

import (
	"encoding/hex"

	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
)

func (s *CrosschainTestSuite) TestTx() {
	require := s.Require()
	vectors := []struct {
		hash            string
		bin             string
		from            string
		to              string
		amount          string
		contract        string
		ata             string
		recentBlockhash string
		sighash         string
	}{
		{
			// received 1 SOL from faucet
			"5U2YvvKUS6NUrDAJnABHjx2szwLCVmg8LCRK9BDbZwVAbf2q5j8D9Sc9kUoqanoqpn6ZpDguY3rip9W7N7vwCjSw",
			"01df5ff457c2cdd23242ab26edd0b308d78499f28c6d43e185149cacdb88b35db171f1779e48ce2224cc80b9b9ce46dd80758319068b08eae34b14dc2cd070ab000100010379726da52d99d60b07ead73b2f6f0bf6083cc85c77a94e34d691d78f8bcafec9fc880863219008406235fa4c8fbb2a86d3da7b6762eac39323b2a1d8c404a4140000000000000000000000000000000000000000000000000000000000000000932bbef1569d58f4a116f41028f766439b2ba52c68c3308bbbea2b21e4716f6701020200010c0200000000ca9a3b00000000",
			"9B5XszUGdMaxCZ7uSQhPzdks5ZQSmWxrmzCSvtJ6Ns6g",
			"Hzn3n914JaSpnxo5mBbmuCDmGL6mxWN9Ac2HzEXFSGtb",
			"1000000000",
			"",
			"",
			"AuVczqLUL8EtFX64vBKbEp2nC1XWB5G589bChZ1zXjre",
			"0100010379726da52d99d60b07ead73b2f6f0bf6083cc85c77a94e34d691d78f8bcafec9fc880863219008406235fa4c8fbb2a86d3da7b6762eac39323b2a1d8c404a4140000000000000000000000000000000000000000000000000000000000000000932bbef1569d58f4a116f41028f766439b2ba52c68c3308bbbea2b21e4716f6701020200010c0200000000ca9a3b00000000",
		},
		{
			// sent 0.01 SOL
			"2xya3FpeCzrNNhBRUxXkNYoN8uKSrTfYHGExGWdq7qv4KK5QZqeJ7vRD41wkE4bVnA4xFkqVWrdPVCSB5ES6zsgL",
			"016249babbf7957c48553193f732ecaa99911e000f19a22fad842ade5b673dcac4bb21f6a2d571e6157e0fbf43e4737045bae63a03a758071fbc9c455bccc5c60901000103fc880863219008406235fa4c8fbb2a86d3da7b6762eac39323b2a1d8c404a4149c2a065622d0167ca0225ee7ce4cfb435c826245a4a1eaea52e038e365b1ba200000000000000000000000000000000000000000000000000000000000000000dc41aee9064607cdb2190b3db02177f694b3ae7729ae8f326becdc5fb9e0749001020200010c020000008096980000000000",
			"Hzn3n914JaSpnxo5mBbmuCDmGL6mxWN9Ac2HzEXFSGtb",
			"BWbmXj5ckAaWCAtzMZ97qnJhBAKegoXtgNrv9BUpAB11",
			"10000000",
			"",
			"",
			"FpnnUiTYBKD5VJKd2ZaPAimqnpf1gMLgeqM4CDGNfhyZ",
			"01000103fc880863219008406235fa4c8fbb2a86d3da7b6762eac39323b2a1d8c404a4149c2a065622d0167ca0225ee7ce4cfb435c826245a4a1eaea52e038e365b1ba200000000000000000000000000000000000000000000000000000000000000000dc41aee9064607cdb2190b3db02177f694b3ae7729ae8f326becdc5fb9e0749001020200010c020000008096980000000000",
		},
		{
			// received 1 USDC from faucet (TokenTransferChecked)
			"2xRfP3gythPLBNUFtfNQ48nD9QBJjvh6AEz6xPgnrGkpQy1RBEvKveMUjix7t3rdPCpemdcS4DaKMLT5t27f92yQ",
			"0161d0475c00c26cdb70eb68015bc38b8b640900ff9fa3e792cff4dfc05805de68de887e4190ca1d1d0e9e5f63f80146a79c33895e1c4bdea00dbce6e1161e6f0b01000609fc7d31eb74d23e204db9a14e646f992b6609340a2a9fd1a9a768b53c0e60fbe5bffd0fa1f34582c787c5f7de3d3a40e3fb50ff51438aa7a5cfb9f96e740ac99366acfe146e0a8ccba9520ae7aed0362bddfe8e1d184cbb25d6896a1f6aa7c294fc880863219008406235fa4c8fbb2a86d3da7b6762eac39323b2a1d8c404a4143b442cb3912157f13a933d0134282d032b5ffecd01a2dbf1b7790608df002ea7000000000000000000000000000000000000000000000000000000000000000006ddf6e1d765a193d9cbe146ceeb79ac1cb485ed5f5b37913a8cf5857eff00a906a7d517192c5c51218cc94c3d4af17f58daee089ba1fd44e3dbd98a000000008c97258f4e2489f1bb3d1029148e0d830b5a1399daff1084048e7bd8dbe9f859353f98251a8e1b9a922be3874c926f2247d4352708c2f1d3968e11db18d6c7b70208070001030405060700060502040100000a0c40420f000000000006",
			"HzcTrHjkEhjFTHEsC6Dsv8DXCh21WgujD4s5M15Sm94g",
			"",
			"1000000",
			"4zMMC9srt5Ri5X14GAgXhaHii3GnPAEERYPJgZJDncDU",
			"DvSgNMRxVSMBpLp4hZeBrmQo8ZRFne72actTZ3PYE3AA",
			"4arrNA8SxQd8Y9kr2ZjiQuZJUJPnMo87PBWwrBhQBMyx",
			"01000609fc7d31eb74d23e204db9a14e646f992b6609340a2a9fd1a9a768b53c0e60fbe5bffd0fa1f34582c787c5f7de3d3a40e3fb50ff51438aa7a5cfb9f96e740ac99366acfe146e0a8ccba9520ae7aed0362bddfe8e1d184cbb25d6896a1f6aa7c294fc880863219008406235fa4c8fbb2a86d3da7b6762eac39323b2a1d8c404a4143b442cb3912157f13a933d0134282d032b5ffecd01a2dbf1b7790608df002ea7000000000000000000000000000000000000000000000000000000000000000006ddf6e1d765a193d9cbe146ceeb79ac1cb485ed5f5b37913a8cf5857eff00a906a7d517192c5c51218cc94c3d4af17f58daee089ba1fd44e3dbd98a000000008c97258f4e2489f1bb3d1029148e0d830b5a1399daff1084048e7bd8dbe9f859353f98251a8e1b9a922be3874c926f2247d4352708c2f1d3968e11db18d6c7b70208070001030405060700060502040100000a0c40420f000000000006",
		},
		{
			// sent 0.2 USDC (TokenTransferChecked)
			"5ZrG8iS4RxLXDRQEWkAoddWHzkS1fA1m6ppxaAekgGzskhcFqjkw1ZaFCsLorbhY5V4YUUkjE3SLY2JNLyVanxrM",
			"01e465609ccfe90037f63a2d7420a47d45e4c127dd8a44443ec932e74d14264ceff568f0a96251135015058ee0683542a07517cc17513f61ea1dafb293a00ea40201000205fc880863219008406235fa4c8fbb2a86d3da7b6762eac39323b2a1d8c404a4145267b2cdf664f91f8234a7c251af95eea15eacea33d4e233bf018fc0941dd272bffd0fa1f34582c787c5f7de3d3a40e3fb50ff51438aa7a5cfb9f96e740ac9933b442cb3912157f13a933d0134282d032b5ffecd01a2dbf1b7790608df002ea706ddf6e1d765a193d9cbe146ceeb79ac1cb485ed5f5b37913a8cf5857eff00a924485ffcbb9027aa9c2c014d1514ea879adf3a83ae1a0d6dd297eb3dc854f9a4010404020301000a0c400d03000000000006",
			"Hzn3n914JaSpnxo5mBbmuCDmGL6mxWN9Ac2HzEXFSGtb",
			"",
			"200000",
			"4zMMC9srt5Ri5X14GAgXhaHii3GnPAEERYPJgZJDncDU",
			"6Yg9GttAiHjbHMoiomBuGBDULP7HxQyez45dEiR9CJqw",
			"3SdgnLoSP9nzDdNyiccukAd3bmQXj8fX7ArMadM9ge4j",
			"01000205fc880863219008406235fa4c8fbb2a86d3da7b6762eac39323b2a1d8c404a4145267b2cdf664f91f8234a7c251af95eea15eacea33d4e233bf018fc0941dd272bffd0fa1f34582c787c5f7de3d3a40e3fb50ff51438aa7a5cfb9f96e740ac9933b442cb3912157f13a933d0134282d032b5ffecd01a2dbf1b7790608df002ea706ddf6e1d765a193d9cbe146ceeb79ac1cb485ed5f5b37913a8cf5857eff00a924485ffcbb9027aa9c2c014d1514ea879adf3a83ae1a0d6dd297eb3dc854f9a4010404020301000a0c400d03000000000006",
		},
		{
			// received 0.2 USDC (TokenTransfer)
			"5imoPkDcXUnRx1H5g4R7kudAJzTaVqK6nSTn6G95yj11jzSVXPGZnCKA4fsKxzqKgkoou3VKqw6q7uH7fuYzfbsL",
			"01ec17478d45c674e1a1fa4cf9a3d4b745fb7652951ef7b9ca2d2b3a4de0e112aa248fbc20c4646652e9f261849ca45a43e3d72dfdd444275b0083ea290aac4d0f0100030652986010573739df4b58ba50e39cf3f335b89cc7d1cb1d32b5de04efa068c939686d2a4ce8a5af61fceb001cab6d18c60a1f06fb51dd49cb6e7a90add1b4e073a2cf14f40c19ea7389494738ca2ce35caf2439a07273452ff89029f655e3321431cadce2aa36ec04604683eaa6f1362c329e1191042842a04e09b82b759fb32449b9ea6ac2a16091c35eff193fcc65fa6bac99d17a0c83f6b5341d8f8a650bef06ddf6e1d765a193d9cbe146ceeb79ac1cb485ed5f5b37913a8cf5857eff00a9c6e03adae0f5ee96a2345b4fb7ca0efa8f96d29e523d5e06df4034aa704f8422020301042000000000000000000000000000000000000000000000000000000000000000000503010200090300fbd4315d000000",
			"6ZRCB7AAqGre6c72PRz3MHLC73VMYvJ8bi9KHf1HFpNk",
			"",
			"400268000000",
			"",
			"BxYC6Su89aLQa8pZTo1EgcDziMBviZVNLH92r3MeFUQK",
			"EPL2z7YjrZwJk47BBm7kS3uoHPx6Dbsj1WYvPKGcd3d3",
			"0100030652986010573739df4b58ba50e39cf3f335b89cc7d1cb1d32b5de04efa068c939686d2a4ce8a5af61fceb001cab6d18c60a1f06fb51dd49cb6e7a90add1b4e073a2cf14f40c19ea7389494738ca2ce35caf2439a07273452ff89029f655e3321431cadce2aa36ec04604683eaa6f1362c329e1191042842a04e09b82b759fb32449b9ea6ac2a16091c35eff193fcc65fa6bac99d17a0c83f6b5341d8f8a650bef06ddf6e1d765a193d9cbe146ceeb79ac1cb485ed5f5b37913a8cf5857eff00a9c6e03adae0f5ee96a2345b4fb7ca0efa8f96d29e523d5e06df4034aa704f8422020301042000000000000000000000000000000000000000000000000000000000000000000503010200090300fbd4315d000000",
		},
		{
			"",
			"",
			"",
			"",
			"0",
			"",
			"",
			"",
			"",
		},
	}

	for _, v := range vectors {
		bytes, _ := hex.DecodeString(v.bin)
		solTx, _ := solana.TransactionFromDecoder(bin.NewBinDecoder(bytes))
		tx := &Tx{
			SolTx: solTx,
		}
		tx.ParseTransfer()

		// basic info
		require.Equal(v.hash, string(tx.Hash()))
		require.Equal(v.from, string(tx.From()))
		require.Equal(v.to, string(tx.To()))
		require.Equal(v.amount, tx.Amount().String())
		require.Equal(v.contract, string(tx.ContractAddress()))
		require.Equal(v.ata, string(tx.ToAlt()))
		require.Equal(v.recentBlockhash, string(tx.RecentBlockhash()))

		// sighash
		sighash, err := tx.Sighash()
		if v.hash == "" {
			require.EqualError(err, "transaction not initialized")
			require.Nil(tx.SolTx)
		} else {
			require.Nil(err)
			require.Equal(v.sighash, hex.EncodeToString(sighash))

			// remove signature
			require.Equal(1, len(tx.SolTx.Signatures))
			sig := tx.SolTx.Signatures[0][:]
			tx.SolTx.Signatures = tx.SolTx.Signatures[:0]
			require.Equal(0, len(tx.SolTx.Signatures))
			require.Equal("", string(tx.Hash()))

			// // readd signature
			err = tx.AddSignature(sig)
			require.Nil(err)
			require.Equal(v.hash, string(tx.Hash()))
		}
	}
}

func (s *CrosschainTestSuite) TestTxHashErr() {
	require := s.Require()

	tx := Tx{}
	hash := tx.Hash()
	require.Equal("", string(hash))
}

func (s *CrosschainTestSuite) TestTxSighashErr() {
	require := s.Require()

	tx := Tx{}
	sighash, err := tx.Sighash()
	require.EqualError(err, "transaction not initialized")
	require.Nil(sighash)
}

func (s *CrosschainTestSuite) TestTxAddSignatureErr() {
	require := s.Require()

	tx := Tx{}
	err := tx.AddSignature([]byte{})
	require.EqualError(err, "invalid signature (0): ")

	err = tx.AddSignature([]byte{1, 2, 3})
	require.EqualError(err, "invalid signature (3): 010203")

	bytes := make([]byte, 64)
	err = tx.AddSignature(bytes)
	require.EqualError(err, "transaction not initialized")
}
