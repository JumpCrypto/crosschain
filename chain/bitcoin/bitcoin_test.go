package bitcoin

import (
	"context"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/btcsuite/btcd/wire"
	xc "github.com/jumpcrypto/crosschain"
	"github.com/jumpcrypto/crosschain/test"
	"github.com/stretchr/testify/suite"
)

var UXTO_ASSETS []xc.NativeAsset = []xc.NativeAsset{
	xc.BTC,
	xc.BCH,
	xc.DOGE,
	xc.LTC,
}

type CrosschainTestSuite struct {
	suite.Suite
	Ctx context.Context
}

func (s *CrosschainTestSuite) SetupTest() {
	s.Ctx = context.Background()
}

func TestBitcoinTestSuite(t *testing.T) {
	suite.Run(t, new(CrosschainTestSuite))
}

// Address

func (s *CrosschainTestSuite) TestNewAddressBuilder() {
	require := s.Require()
	for _, nativeAsset := range UXTO_ASSETS {
		builder, err := NewAddressBuilder(&xc.AssetConfig{NativeAsset: nativeAsset})
		require.NotNil(builder)
		require.NoError(err)
	}
}

func (s *CrosschainTestSuite) TestGetAddressFromPublicKey() {
	require := s.Require()
	for _, nativeAsset := range UXTO_ASSETS {
		builder, err := NewAddressBuilder(&xc.AssetConfig{
			Net:         "testnet",
			NativeAsset: nativeAsset,
		})
		require.NoError(err)
		pubkey, err := base64.RawStdEncoding.DecodeString("AptrsfXbXbvnsWxobWNFoUXHLO5nmgrQb3PDmGGu1CSS")
		require.NoError(err)
		fmt.Println("checking address for ", nativeAsset)
		switch nativeAsset {
		case xc.BTC, xc.BCH:
			address, err := builder.GetAddressFromPublicKey(pubkey)
			require.NoError(err)
			require.Equal(xc.Address("mhYWE7RrYCgbq4RJDaqZp8fvzVmYnPVnFD"), address)
		case xc.DOGE:
			address, err := builder.GetAddressFromPublicKey(pubkey)
			require.NoError(err)
			require.Equal(xc.Address("nWDiCL2RxZcMTvhUGRWCnPDWFWHSCfkhoz"), address)
		case xc.LTC:
			address, err := builder.GetAddressFromPublicKey(pubkey)
			require.NoError(err)
			require.Equal(xc.Address("mhYWE7RrYCgbq4RJDaqZp8fvzVmYnPVnFD"), address)
		default:
			panic("need to add address test case for " + nativeAsset)
		}
	}
}

func (s *CrosschainTestSuite) TestGetAllPossibleAddressesFromPublicKey() {
	require := s.Require()
	builder, err := NewAddressBuilder(&xc.AssetConfig{
		Net:         "testnet",
		NativeAsset: "BTC",
	})
	require.NoError(err)
	pubkey, err := base64.RawStdEncoding.DecodeString("AptrsfXbXbvnsWxobWNFoUXHLO5nmgrQb3PDmGGu1CSS")
	require.NoError(err)
	addresses, err := builder.GetAllPossibleAddressesFromPublicKey(pubkey)
	require.NoError(err)

	validated_p2pkh := false
	validated_p2wkh := false

	fmt.Println(addresses)
	for _, addr := range addresses {
		if addr.Address == "mhYWE7RrYCgbq4RJDaqZp8fvzVmYnPVnFD" {
			require.Equal(xc.AddressTypeP2PKH, addr.Type)
			validated_p2pkh = true
		} else if addr.Address == "tb1qzca49vcyxkt989qcmhjfp7wyze7n9pq50k2cfd" {
			require.Equal(xc.AddressTypeP2WPKH, addr.Type)
			validated_p2wkh = true
		} else {
			panic("unexpected address generated: " + addr.Address)
		}
	}
	require.True(validated_p2pkh)
	require.True(validated_p2wkh)
}

// TxBuilder

func (s *CrosschainTestSuite) TestNewTxBuilder() {
	require := s.Require()
	for _, nativeAsset := range UXTO_ASSETS {
		builder, err := NewTxBuilder(&xc.AssetConfig{NativeAsset: nativeAsset})
		require.NotNil(builder)
		require.NoError(err)
	}
}

func (s *CrosschainTestSuite) TestFetchTxInput() {
	require := s.Require()

	server, close := test.MockHTTP(&s.Suite, []string{
		// fetch UnspentOutputs
		`{"data":{"mpjwFvP88ZwAt3wEHY6irKkGhxcsv22BP6":{"address":{"type":"pubkeyhash","script_hex":"76a914652dac91ff1b130616cb11ce33b0ac2f1b4df89188ac","balance":2392235,"balance_usd":0,"received":35323650,"received_usd":0,"spent":32931415,"spent_usd":0,"output_count":12,"unspent_output_count":1,"first_seen_receiving":"2023-04-12 16:16:31","last_seen_receiving":"2023-04-13 15:15:24","first_seen_spending":"2023-04-12 22:28:01","last_seen_spending":"2023-04-13 15:15:24","scripthash_type":null,"transaction_count":12},"transactions":["c4979460bb03a1877bbf23571c83edbd02cb4da20049916fa6c5fbf77470e027","bb84ca051f523835f6b74a09264869f48c585a9e664c450de99905a59f8f410d","81ad62df63a86d2aa4cf524fdff4a6d85f9e51437137c3294482d29499e04e1a","c5175760193f00c33291669e0f3e2628fc1c1aaa083e29ae7f3ed23e2da4cf56","46be2eb86cbc249f0e2c43430fff35cc626a814b014ba6809bb7c03124662efa","3d380e087e07ab392de5e7653ccea054c3394c95f9378463522c8a094a21b584","cc0746e4dfb5e5da26f27810d29666be44b825bacfbec297454ea3e0903a2440","e7d3bf1722af2fcb0ec27b03ca32ec6079d45640e02a7fe3a43947d20a84285e","8b503826b34e7c44e5b0fda2ab378cbbef3992765322e60cc7b27ad777f05202","49d5fd5d9c6909b7a4e0af010a0693244cd0067c7d7cc16ec5948fc779638310","3013bb3657c545c881cac232ee6341e57656669e464103d8af3c1ddf859bde06","b054cd53be7fc8cb75d33372c0b8867f17a4cd49c367d413d0147719fd14c5f6"],"utxo":[{"block_id":-1,"transaction_hash":"c4979460bb03a1877bbf23571c83edbd02cb4da20049916fa6c5fbf77470e027","index":1,"value":2392235}]}},"context":{"code":200,"source":"D","limit":"100,100","offset":"0,0","results":1,"state":2428749,"market_price_usd":30494,"cache":{"live":true,"duration":20,"since":"2023-04-13 15:16:43","until":"2023-04-13 15:17:03","time":null},"api":{"version":"2.0.95-ie","last_major_update":"2022-11-07 02:00:00","next_major_update":null,"documentation":"https:\/\/blockchair.com\/api\/docs","notice":"Please note that on November 7th, 2022 public support for the following blockchains was dropped: EOS, Bitcoin SV"},"servers":"API4,TBTC0","time":0.8323581218719482,"render_time":0.00943303108215332,"full_time":0.8417911529541016,"request_cost":1}}`,
		// fetch blockinfo (for estimate gas)
		`{"data":{"blocks":2428756,"transactions":65332308,"outputs":173266703,"circulation":2099211173546005,"blocks_24h":125,"transactions_24h":9376,"difficulty":104649090.3851,"volume_24h":9305170450343,"mempool_transactions":91,"mempool_size":25369,"mempool_tps":0.21666666666666667,"mempool_total_fee_usd":0,"best_block_height":2428755,"best_block_hash":"00000000000000171993c83855edbbdb4b596a80d7979b9906199a152c02e602","best_block_time":"2023-04-13 15:55:31","blockchain_size":28727851118,"average_transaction_fee_24h":5919,"inflation_24h":305175750,"median_transaction_fee_24h":208,"cdd_24h":10812.844745459975,"mempool_outputs":308,"largest_transaction_24h":{"hash":"bb7fb631e27a18b8802ead03f3ee14b69ae71edb845697f98e0f072b845b0be4","value_usd":0},"hashrate_24h":"650455020414125","inflation_usd_24h":0,"average_transaction_fee_usd_24h":0,"median_transaction_fee_usd_24h":0,"market_price_usd":0,"market_price_btc":0,"market_price_usd_change_24h_percentage":0,"market_cap_usd":0,"market_dominance_percentage":0,"next_retarget_time_estimate":"2023-04-17 02:55:52","next_difficulty_estimate":107945581,"suggested_transaction_fee_per_byte_sat":1,"hodling_addresses":10010301},"context":{"code":200,"source":"A","state":2428755,"market_price_usd":30467,"cache":{"live":false,"duration":"Ignore","since":"2023-04-13 16:03:49","until":"2023-04-13 16:05:00","time":2.86102294921875e-6},"api":{"version":"2.0.95-ie","last_major_update":"2022-11-07 02:00:00","next_major_update":null,"documentation":"https:\/\/blockchair.com\/api\/docs","notice":"Please note that on November 7th, 2022 public support for the following blockchains was dropped: EOS, Bitcoin SV"},"servers":"API4,TBTC0","time":1.8575801849365234,"render_time":0.007244110107421875,"full_time":0.007246971130371094,"request_cost":1}}`,
	})
	defer close()
	asset := &xc.AssetConfig{NativeAsset: xc.BTC, URL: server.URL, Net: "testnet"}
	client, _ := NewClient(asset)

	from := xc.Address("mpjwFvP88ZwAt3wEHY6irKkGhxcsv22BP6")
	to := xc.Address("tb1qtpqqpgadjr2q3f4wrgd6ndclqtfg7cz5evtvs0")
	input, err := client.FetchTxInput(s.Ctx, from, to)
	require.NotNil(input)
	fmt.Println(input)
	fmt.Println(err)
	require.NoError(err)
	btcInput := input.(*TxInput)
	fmt.Println(btcInput)
	require.Len(btcInput.UnspentOutputs, 1)
	require.EqualValues(2392235, btcInput.UnspentOutputs[0].Value.Uint64())
	require.EqualValues(1, btcInput.UnspentOutputs[0].Index)
	// string should be reversed
	require.EqualValues("27e07074f7fbc5a66f914900a24dcb02bded831c5723bf7b87a103bb609497c4", hex.EncodeToString(btcInput.UnspentOutputs[0].Hash))
	require.EqualValues(1, btcInput.GasPricePerByte.Uint64())
}

func (s *CrosschainTestSuite) TestNewNativeTransfer() {
	require := s.Require()
	asset := &xc.AssetConfig{NativeAsset: xc.BTC, Net: "testnet"}
	builder, _ := NewTxBuilder(asset)
	from := xc.Address("mpjwFvP88ZwAt3wEHY6irKkGhxcsv22BP6")
	to := xc.Address("tb1qtpqqpgadjr2q3f4wrgd6ndclqtfg7cz5evtvs0")
	amount := xc.NewAmountBlockchainFromUint64(1)
	input := &TxInput{
		UnspentOutputs: []Output{{
			Value: xc.NewAmountBlockchainFromUint64(1000),
		}},
		GasPricePerByte: xc.NewAmountBlockchainFromUint64(1),
	}
	tf, err := builder.(xc.TxTokenBuilder).NewNativeTransfer(from, to, amount, input)
	require.NoError(err)
	require.NotNil(tf)
	hash := tf.Hash()
	require.Len(hash, 64)

	// Having not enough balance for fees will be an error
	input_small := &TxInput{
		UnspentOutputs: []Output{{
			Value: xc.NewAmountBlockchainFromUint64(5),
		}},
		GasPricePerByte: xc.NewAmountBlockchainFromUint64(1),
	}
	_, err = builder.(xc.TxTokenBuilder).NewNativeTransfer(from, to, amount, input_small)
	require.Error(err)

	// add signature
	sig := []byte{}
	for i := 0; i < 65; i++ {
		sig = append(sig, byte(i))
	}
	err = tf.AddSignatures(xc.TxSignature(sig))
	require.NoError(err)

	ser, err := tf.Serialize()
	require.NoError(err)
	require.True(len(ser) > 64)
}

func (s *CrosschainTestSuite) TestNewTokenTransfer() {
	require := s.Require()
	asset := &xc.AssetConfig{NativeAsset: xc.BTC, Net: "testnet"}
	builder, _ := NewTxBuilder(asset)
	from := xc.Address("mpjwFvP88ZwAt3wEHY6irKkGhxcsv22BP6")
	to := xc.Address("tb1qtpqqpgadjr2q3f4wrgd6ndclqtfg7cz5evtvs0")
	amount := xc.NewAmountBlockchainFromUint64(1)
	input := &TxInput{
		UnspentOutputs: []Output{{
			Value: xc.NewAmountBlockchainFromUint64(1000),
		}},
		GasPricePerByte: xc.NewAmountBlockchainFromUint64(1),
	}
	tf, err := builder.(xc.TxTokenBuilder).NewTokenTransfer(from, to, amount, input)
	require.Nil(tf)
	require.EqualError(err, "not implemented")
}

func (s *CrosschainTestSuite) TestNewTransfer() {
	require := s.Require()
	asset := &xc.AssetConfig{NativeAsset: xc.BTC, Net: "testnet"}
	builder, _ := NewTxBuilder(asset)
	from := xc.Address("mpjwFvP88ZwAt3wEHY6irKkGhxcsv22BP6")
	to := xc.Address("tb1qtpqqpgadjr2q3f4wrgd6ndclqtfg7cz5evtvs0")
	amount := xc.NewAmountBlockchainFromUint64(1)
	input := &TxInput{
		UnspentOutputs: []Output{{
			Value: xc.NewAmountBlockchainFromUint64(1000),
		}},
		GasPricePerByte: xc.NewAmountBlockchainFromUint64(1),
	}
	tf, err := builder.NewTransfer(from, to, amount, input)
	require.NotNil(tf)
	require.NoError(err)
}

func (s *CrosschainTestSuite) TestNewTransfer_Token() {
	require := s.Require()
	asset := &xc.AssetConfig{NativeAsset: xc.BTC, Net: "testnet", Type: xc.AssetTypeToken}
	builder, _ := NewTxBuilder(asset)
	from := xc.Address("mpjwFvP88ZwAt3wEHY6irKkGhxcsv22BP6")
	to := xc.Address("tb1qtpqqpgadjr2q3f4wrgd6ndclqtfg7cz5evtvs0")
	amount := xc.AmountBlockchain{}
	input := TxInput{}
	tf, err := builder.NewTransfer(from, to, amount, input)
	require.Nil(tf)
	require.EqualError(err, "not implemented")
}

// Client

func (s *CrosschainTestSuite) TestNewClient() {
	require := s.Require()
	asset := &xc.AssetConfig{NativeAsset: xc.BTC, Net: "testnet"}
	client, err := NewClient(asset)
	require.NotNil(client)
	require.NoError(err)
}

func (s *CrosschainTestSuite) TestSubmitTx() {
	require := s.Require()
	server, close := test.MockHTTP(&s.Suite, []string{
		// transaction submitted
		`{"data":{"transaction_hash":"227178d784150211e8ea5a586ee75bc97655e61f02bc8c07557e475cfecea3cd"},"context":{"code":200,"source":"R","state":2428749,"market_price_usd":30494,"cache":{"live":true,"duration":20,"since":"2023-04-13 15:16:43","until":"2023-04-13 15:17:03","time":null},"api":{"version":"2.0.95-ie","last_major_update":"2022-11-07 02:00:00","next_major_update":null,"documentation":"https:\/\/blockchair.com\/api\/docs","notice":"Please note that on November 7th, 2022 public support for the following blockchains was dropped: EOS, Bitcoin SV"},"servers":"API4,TBTC0","time":0.0487971305847168,"render_time":0.007144927978515625,"full_time":0.05594205856323242,"request_cost":1}}`,
	})
	defer close()
	asset := &xc.AssetConfig{NativeAsset: xc.BTC, URL: server.URL, Net: "testnet"}
	client, err := NewClient(asset)
	require.NoError(err)
	err = client.SubmitTx(s.Ctx, &Tx{
		msgTx: wire.NewMsgTx(2),
	})
	require.NoError(err)
}

func (s *CrosschainTestSuite) TestFetchTxInfo() {
	require := s.Require()
	server, close := test.MockHTTP(&s.Suite, []string{
		// tx info
		`{"data":{"227178d784150211e8ea5a586ee75bc97655e61f02bc8c07557e475cfecea3cd":{"transaction":{"block_id":2428751,"id":65331999,"hash":"227178d784150211e8ea5a586ee75bc97655e61f02bc8c07557e475cfecea3cd","date":"2023-04-13","time":"2023-04-13 15:29:58","size":255,"weight":1020,"version":2,"lock_time":0,"is_coinbase":false,"has_witness":false,"input_count":1,"output_count":2,"input_total":2392235,"input_total_usd":0,"output_total":2391980,"output_total_usd":0,"fee":255,"fee_usd":0,"fee_per_kb":1000,"fee_per_kb_usd":0,"fee_per_kwu":250,"fee_per_kwu_usd":0,"cdd_total":0,"is_rbf":false},"inputs":[{"block_id":2428751,"transaction_id":65331998,"index":1,"transaction_hash":"c4979460bb03a1877bbf23571c83edbd02cb4da20049916fa6c5fbf77470e027","date":"2023-04-13","time":"2023-04-13 15:29:58","value":2392235,"value_usd":0,"recipient":"mpjwFvP88ZwAt3wEHY6irKkGhxcsv22BP6","type":"pubkeyhash","script_hex":"76a914652dac91ff1b130616cb11ce33b0ac2f1b4df89188ac","is_from_coinbase":false,"is_spendable":null,"is_spent":true,"spending_block_id":2428751,"spending_transaction_id":65331999,"spending_index":0,"spending_transaction_hash":"227178d784150211e8ea5a586ee75bc97655e61f02bc8c07557e475cfecea3cd","spending_date":"2023-04-13","spending_time":"2023-04-13 15:29:58","spending_value_usd":0,"spending_sequence":4294967295,"spending_signature_hex":"483045022100ad6a8b65d8eeeecf729d1ff6a8af95f3e9049944d3ca9ef5a9f7410dab494fc80220350555ff6b2911abbeda5190bbe3220d57e703feaf4edf393edc3fa2ec390bdf014104e6d880f2d81328599fd482d6e1e3a4ff5698dabccd1969d88a4134c113e17e3df7497d8133d5b3146f79b841e4e3c9e8d07c8a61cf423399e597352da50510e2","spending_witness":"","lifespan":0,"cdd":0}],"outputs":[{"block_id":2428751,"transaction_id":65331999,"index":0,"transaction_hash":"227178d784150211e8ea5a586ee75bc97655e61f02bc8c07557e475cfecea3cd","date":"2023-04-13","time":"2023-04-13 15:29:58","value":100000,"value_usd":0,"recipient":"tb1qtpqqpgadjr2q3f4wrgd6ndclqtfg7cz5evtvs0","type":"witness_v0_keyhash","script_hex":"0014584000a3ad90d408a6ae1a1ba9b71f02d28f6054","is_from_coinbase":false,"is_spendable":null,"is_spent":true,"spending_block_id":2428757,"spending_transaction_id":65332310,"spending_index":2,"spending_transaction_hash":"5e87a2a3d459c438cde63e536f40124f2acaf8d0158931144698da58b9476a0b","spending_date":"2023-04-13","spending_time":"2023-04-13 16:15:41","spending_value_usd":0,"spending_sequence":4294967294,"spending_signature_hex":"","spending_witness":"30440220044a4568409ced4aa381d8bdf3d3af23f063aec0a44feb009a7f345fc2ca25d10220485d78aafc21debedbe60f5196d8f67bbfc397815025c15de16bfbfa9b6fede201,0294fb77024c22c688ca1f548b4878d5a0cc24cb57aec750c87b19c1b780baa7a8","lifespan":2743,"cdd":0},{"block_id":2428751,"transaction_id":65331999,"index":1,"transaction_hash":"227178d784150211e8ea5a586ee75bc97655e61f02bc8c07557e475cfecea3cd","date":"2023-04-13","time":"2023-04-13 15:29:58","value":2291980,"value_usd":0,"recipient":"mpjwFvP88ZwAt3wEHY6irKkGhxcsv22BP6","type":"pubkeyhash","script_hex":"76a914652dac91ff1b130616cb11ce33b0ac2f1b4df89188ac","is_from_coinbase":false,"is_spendable":null,"is_spent":true,"spending_block_id":2428751,"spending_transaction_id":65332000,"spending_index":0,"spending_transaction_hash":"129b73fa6de24c8ebbbca2b4d8f4702ddbf1e02cb6d22cc5cf743d2a92b87880","spending_date":"2023-04-13","spending_time":"2023-04-13 15:29:58","spending_value_usd":0,"spending_sequence":4294967295,"spending_signature_hex":"473044022059a5def9aa5436c923e12f56ad48a75dcaf8e7667191e4ec34e9213482769fe9022063f8f52c7d21031b0d0db04ffb32277447835a87913435f60f439f63133627a4014104e6d880f2d81328599fd482d6e1e3a4ff5698dabccd1969d88a4134c113e17e3df7497d8133d5b3146f79b841e4e3c9e8d07c8a61cf423399e597352da50510e2","spending_witness":"","lifespan":0,"cdd":0}]}},"context":{"code":200,"source":"D","results":1,"state":2428762,"market_price_usd":30389,"cache":{"live":true,"duration":20,"since":"2023-04-13 17:27:13","until":"2023-04-13 17:27:33","time":null},"api":{"version":"2.0.95-ie","last_major_update":"2022-11-07 02:00:00","next_major_update":null,"documentation":"https:\/\/blockchair.com\/api\/docs","notice":"Please note that on November 7th, 2022 public support for the following blockchains was dropped: EOS, Bitcoin SV"},"servers":"API4,TBTC0","time":1.1531751155853271,"render_time":0.049282073974609375,"full_time":1.2024571895599365,"request_cost":1}}`,
	})
	defer close()
	asset := &xc.AssetConfig{NativeAsset: xc.BTC, URL: server.URL, Net: "testnet"}
	client, err := NewClient(asset)
	require.NoError(err)
	info, err := client.FetchTxInfo(s.Ctx, xc.TxHash("227178d784150211e8ea5a586ee75bc97655e61f02bc8c07557e475cfecea3cd"))
	require.NotNil(info)
	require.NoError(err)
	require.EqualValues(100000, info.Amount.Uint64())
	require.EqualValues("227178d784150211e8ea5a586ee75bc97655e61f02bc8c07557e475cfecea3cd", info.TxID)
	require.Len(info.Sources, 1)
	// destination should not include the change
	require.Len(info.Destinations, 1)
	require.EqualValues("tb1qtpqqpgadjr2q3f4wrgd6ndclqtfg7cz5evtvs0", info.Destinations[0].Address)
	require.EqualValues("mpjwFvP88ZwAt3wEHY6irKkGhxcsv22BP6", info.Sources[0].Address)
	require.EqualValues(100000, info.Destinations[0].Amount.Uint64())
	require.EqualValues(xc.TxStatusSuccess, info.Status)
	require.EqualValues(12, info.Confirmations)
	require.EqualValues(255, info.Fee.Uint64())
}

// Signer

func (s *CrosschainTestSuite) TestNewSigner() {
	require := s.Require()
	signer, err := NewSigner(&xc.AssetConfig{})
	require.NotNil(signer)
	require.NoError(err)
}

func (s *CrosschainTestSuite) TestImportPrivateKeyAndSign() {
	require := s.Require()
	signer, _ := NewSigner(&xc.AssetConfig{})
	keyBz := []byte{}
	for i := 0; i < 32; i++ {
		keyBz = append(keyBz, byte(i))
	}
	key, err := signer.ImportPrivateKey(hex.EncodeToString(keyBz))
	require.NotNil(key)
	require.NoError(err)

	sig, err := signer.Sign(key, keyBz)
	require.NotNil(sig)
	require.NoError(err)
}

// Tx

func (s *CrosschainTestSuite) TestTxHash() {
	require := s.Require()

	asset := &xc.AssetConfig{NativeAsset: xc.BTC, Net: "testnet"}
	builder, _ := NewTxBuilder(asset)
	from := xc.Address("mpjwFvP88ZwAt3wEHY6irKkGhxcsv22BP6")
	to := xc.Address("tb1qtpqqpgadjr2q3f4wrgd6ndclqtfg7cz5evtvs0")
	amount := xc.NewAmountBlockchainFromUint64(1)
	input := &TxInput{
		UnspentOutputs: []Output{{
			Value: xc.NewAmountBlockchainFromUint64(1000),
		}},
		GasPricePerByte: xc.NewAmountBlockchainFromUint64(1),
	}
	tf, err := builder.(xc.TxTokenBuilder).NewNativeTransfer(from, to, amount, input)
	require.NoError(err)

	tx := tf.(*Tx)
	require.Equal(xc.TxHash("0ebdd0e519cf4bf67ac4d924c07e3312483b09844c9f16f46c04f5fe1500c788"), tx.Hash())
}

func (s *CrosschainTestSuite) TestTxSighashes() {
	require := s.Require()
	tx := Tx{}
	sighashes, err := tx.Sighashes()
	require.NotNil(sighashes)
	require.NoError(err)
}

func (s *CrosschainTestSuite) TestTxAddSignature() {
	require := s.Require()
	asset := &xc.AssetConfig{NativeAsset: xc.BTC, Net: "testnet"}
	builder, _ := NewTxBuilder(asset)
	from := xc.Address("mpjwFvP88ZwAt3wEHY6irKkGhxcsv22BP6")
	to := xc.Address("tb1qtpqqpgadjr2q3f4wrgd6ndclqtfg7cz5evtvs0")
	amount := xc.NewAmountBlockchainFromUint64(10)
	input := TxInput{
		UnspentOutputs: []Output{{
			Value: xc.NewAmountBlockchainFromUint64(1000),
		}},
	}
	err := input.SetPublicKey(xc.PublicKey{})
	require.NoError(err)
	tf, err := builder.(xc.TxTokenBuilder).NewNativeTransfer(from, to, amount, input)
	require.NoError(err)

	tx := tf.(*Tx)
	err = tx.AddSignatures([]xc.TxSignature{
		[]byte{1, 2, 3, 4},
	}...)
	require.ErrorContains(err, "signature must be 64 or 65 length")
	sig := []byte{}
	for i := 0; i < 65; i++ {
		sig = append(sig, byte(i))
	}
	err = tx.AddSignatures([]xc.TxSignature{
		sig,
	}...)
	require.NoError(err)

	// can't sign multiple times in a row
	err = tx.AddSignatures([]xc.TxSignature{
		sig,
	}...)
	require.ErrorContains(err, "already signed")

	// must have a signature for each input needed
	tf, _ = builder.(xc.TxTokenBuilder).NewNativeTransfer(from, to, amount, input)
	err = tf.(*Tx).AddSignatures([]xc.TxSignature{
		sig, sig,
	}...)
	require.ErrorContains(err, "expected 1 signatures, got 2 signatures")

	// 2 inputs = 2 sigs
	amount = xc.NewAmountBlockchainFromUint64(15000)
	input = TxInput{
		UnspentOutputs: []Output{{
			Value: xc.NewAmountBlockchainFromUint64(10000),
		},
			{
				Value: xc.NewAmountBlockchainFromUint64(10000),
			},
		},
	}
	tf, _ = builder.(xc.TxTokenBuilder).NewNativeTransfer(from, to, amount, input)
	require.Len(tf.(*Tx).input.Inputs, 2)
	err = tf.(*Tx).AddSignatures([]xc.TxSignature{
		sig, sig,
	}...)
	require.NoError(err)
}
