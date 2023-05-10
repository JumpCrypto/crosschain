package factory

import (
	"fmt"
	"testing"

	xc "github.com/jumpcrypto/crosschain"
	"github.com/jumpcrypto/crosschain/chain/aptos"
	"github.com/jumpcrypto/crosschain/chain/bitcoin"
	"github.com/jumpcrypto/crosschain/chain/cosmos"
	"github.com/jumpcrypto/crosschain/chain/evm"
	"github.com/jumpcrypto/crosschain/chain/solana"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/suite"
)

type CrosschainTestSuite struct {
	suite.Suite
	Factory          *Factory
	TestNativeAssets []xc.NativeAsset
	TestAssetConfigs []xc.ITask
}

func (s *CrosschainTestSuite) SetupTest() {
	s.Factory = NewDefaultFactory()
	s.TestNativeAssets = []xc.NativeAsset{
		xc.ETH,
		xc.MATIC,
		xc.BNB,
		xc.SOL,
		// xc.ATOM,
	}
	for _, native := range s.TestNativeAssets {
		assetConfig, _ := s.Factory.GetAssetConfig("", string(native))
		s.TestAssetConfigs = append(s.TestAssetConfigs, assetConfig)
	}

}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(CrosschainTestSuite))
}

// NewObject functions

func (s *CrosschainTestSuite) TestNewDefaultFactory() {
	require := s.Require()
	require.NotNil(s.Factory)
}

func (s *CrosschainTestSuite) TestNewClient() {
	require := s.Require()
	for _, asset := range s.TestAssetConfigs {
		client, _ := s.Factory.NewClient(asset)
		require.NotNil(client)
	}

	asset, _ := s.Factory.PutAssetConfig(&xc.AssetConfig{Asset: "TEST"})
	_, err := s.Factory.NewClient(asset)
	require.EqualError(err, "unsupported asset")
}

func (s *CrosschainTestSuite) TestNewTxBuilder() {
	require := s.Require()
	for _, asset := range s.TestAssetConfigs {
		builder, _ := s.Factory.NewTxBuilder(asset)
		require.NotNil(builder)
	}

	asset, _ := s.Factory.PutAssetConfig(&xc.AssetConfig{Asset: "TEST"})
	_, err := s.Factory.NewTxBuilder(asset)
	require.EqualError(err, "unsupported asset")
}

func (s *CrosschainTestSuite) TestNewSigner() {
	require := s.Require()
	for _, asset := range s.TestAssetConfigs {
		signer, _ := s.Factory.NewSigner(asset)
		require.NotNil(signer)
	}

	asset, _ := s.Factory.PutAssetConfig(&xc.AssetConfig{Asset: "TEST"})
	_, err := s.Factory.NewSigner(asset)
	require.EqualError(err, "unsupported asset")
}

func (s *CrosschainTestSuite) TestNewAddressBuilder() {
	require := s.Require()
	for _, asset := range s.TestAssetConfigs {
		builder, _ := s.Factory.NewAddressBuilder(asset)
		require.NotNil(builder)
	}

	asset, _ := s.Factory.PutAssetConfig(&xc.AssetConfig{Asset: "TEST"})
	_, err := s.Factory.NewAddressBuilder(asset)
	require.EqualError(err, "unsupported asset")
}

// GetObject functions (excluding config)

func (s *CrosschainTestSuite) TestGetAddressFromPublicKey() {
	require := s.Require()
	for _, asset := range s.TestAssetConfigs {
		address, _ := s.Factory.GetAddressFromPublicKey(asset, []byte{})
		require.NotNil(address)
	}
}

func (s *CrosschainTestSuite) TestGetAllPossibleAddressesFromPublicKey() {
	require := s.Require()
	for _, asset := range s.TestAssetConfigs {
		addresses, _ := s.Factory.GetAllPossibleAddressesFromPublicKey(asset, []byte{})
		require.NotNil(addresses)
	}
}

// MustObject functions

func (s *CrosschainTestSuite) TestMustAmountBlockchain() {
	require := s.Require()
	for _, asset := range s.TestAssetConfigs {
		asset := asset.GetAssetConfig()
		amount := s.Factory.MustAmountBlockchain(asset, "10.3")

		var expected xc.AmountBlockchain
		if asset.Decimals == 6 {
			expected = xc.NewAmountBlockchainFromUint64(10300000)
		}
		if asset.Decimals == 9 {
			expected = xc.NewAmountBlockchainFromUint64(10300000000)
		}
		if asset.Decimals == 12 {
			expected = xc.NewAmountBlockchainFromUint64(10300000000 * 1000)
		}
		if asset.Decimals == 18 {
			expected = xc.NewAmountBlockchainFromUint64(10300000000 * 1000000000)
		}

		require.Equal(expected, amount, "Error on: "+asset.NativeAsset)
	}
}

func (s *CrosschainTestSuite) TestMustAddress() {
	require := s.Require()
	for _, asset := range s.TestAssetConfigs {
		asset := asset.GetAssetConfig()
		address := s.Factory.MustAddress(asset, "myaddress") // trivial impl
		require.Equal(xc.Address("myaddress"), address, "Error on: "+asset.NativeAsset)
	}
}

func (s *CrosschainTestSuite) TestMustPrivateKey() {
	require := s.Require()
	for _, asset := range s.TestAssetConfigs {
		asset := asset.GetAssetConfig()
		if asset.NativeAsset != xc.SOL {
			continue
		}
		privateKey := s.Factory.MustPrivateKey(asset, "myprivatekey")
		require.NotNil(privateKey, "Error on: "+asset.NativeAsset)
	}
}

// Convert functions

func (s *CrosschainTestSuite) TestConvertAmountToHuman() {
	require := s.Require()
	var amountBlockchain xc.AmountBlockchain
	for _, asset := range s.TestAssetConfigs {
		asset := asset.GetAssetConfig()
		if asset.Decimals == 6 {
			amountBlockchain = xc.NewAmountBlockchainFromUint64(10300000)
		}
		if asset.Decimals == 9 {
			amountBlockchain = xc.NewAmountBlockchainFromUint64(10300000000)
		}
		if asset.Decimals == 12 {
			amountBlockchain = xc.NewAmountBlockchainFromUint64(10300000000 * 1000)
		}
		if asset.Decimals == 18 {
			amountBlockchain = xc.NewAmountBlockchainFromUint64(10300000000 * 1000000000)
		}

		amount, err := s.Factory.ConvertAmountToHuman(asset, amountBlockchain)
		require.Nil(err)
		require.Equal("10.3", amount.String(), "Error on: "+asset.NativeAsset)
	}
	asset, _ := s.Factory.PutAssetConfig(&xc.AssetConfig{Asset: "TEST", Decimals: 0})
	amountBlockchain = xc.NewAmountBlockchainFromUint64(103)
	amount, err := s.Factory.ConvertAmountToHuman(asset, amountBlockchain)
	require.NoError(err)
	require.EqualValues("103", amount.String())
}

func (s *CrosschainTestSuite) TestConvertAmountToBlockchain() {
	require := s.Require()
	amountDecimal, _ := decimal.NewFromString("10.3")
	amountHuman := xc.AmountHumanReadable(amountDecimal)

	var expected xc.AmountBlockchain
	for _, asset := range s.TestAssetConfigs {
		asset := asset.GetAssetConfig()
		amount, err := s.Factory.ConvertAmountToBlockchain(asset, amountHuman)

		if asset.Decimals == 6 {
			expected = xc.NewAmountBlockchainFromUint64(10300000)
		}
		if asset.Decimals == 9 {
			expected = xc.NewAmountBlockchainFromUint64(10300000000)
		}
		if asset.Decimals == 12 {
			expected = xc.NewAmountBlockchainFromUint64(10300000000 * 1000)
		}
		if asset.Decimals == 18 {
			expected = xc.NewAmountBlockchainFromUint64(10300000000 * 1000000000)
		}

		require.Nil(err)
		require.Equal(expected, amount, "Error on: "+asset.NativeAsset)
	}
}

func (s *CrosschainTestSuite) TestConvertAmountStrToBlockchain() {
	require := s.Require()
	var expected xc.AmountBlockchain
	for _, asset := range s.TestAssetConfigs {
		asset := asset.GetAssetConfig()
		amount, err := s.Factory.ConvertAmountStrToBlockchain(asset, "10.3")

		if asset.Decimals == 6 {
			expected = xc.NewAmountBlockchainFromUint64(10_300_000)
		}
		if asset.Decimals == 9 {
			expected = xc.NewAmountBlockchainFromUint64(10_300_000_000)
		}
		if asset.Decimals == 12 {
			expected = xc.NewAmountBlockchainFromUint64(10_300_000_000_000)
		}
		if asset.Decimals == 18 {
			expected = xc.NewAmountBlockchainFromUint64(10_300_000_000_000_000_000)
		}

		require.Nil(err)
		require.Equal(expected, amount, "Error on: "+asset.NativeAsset)
	}

	asset, _ := s.Factory.PutAssetConfig(&xc.AssetConfig{Asset: "TEST", Decimals: 0})
	amount, err := s.Factory.ConvertAmountStrToBlockchain(asset, "103")
	require.NoError(err)
	require.EqualValues(103, amount.Uint64())
}

func (s *CrosschainTestSuite) TestConvertAmountStrToBlockchainErr() {
	require := s.Require()
	for _, asset := range s.TestAssetConfigs {
		amount, err := s.Factory.ConvertAmountStrToBlockchain(asset, "")
		require.EqualError(err, "can't convert  to decimal")
		require.Equal(xc.NewAmountBlockchainFromUint64(0), amount)

		_, err = s.Factory.ConvertAmountStrToBlockchain(asset, "err")
		require.EqualError(err, "can't convert err to decimal: exponent is not numeric")
		require.Equal(xc.NewAmountBlockchainFromUint64(0), amount)
	}
}

// Config functions

func (s *CrosschainTestSuite) TestEnrichAssetConfig() {
	require := s.Require()

	assetCfgI, _ := s.Factory.GetAssetConfig("USDC", "SOL")
	assetCfg := assetCfgI.(*xc.TokenAssetConfig)
	assetCfg.URL = ""
	assetCfgEnriched, err := s.Factory.EnrichAssetConfig(assetCfg)
	require.Nil(err)
	require.NotNil(assetCfgEnriched)
	require.Equal("USDC", assetCfgEnriched.Asset)
	require.NotEqual("", assetCfgEnriched.URL)
	require.NotEqual("", assetCfgEnriched.Driver)

	assetCfg.URL = ""
	assetCfg.Chain = "TEST"
	assetCfgEnriched, err = s.Factory.EnrichAssetConfig(assetCfg)
	require.EqualError(err, "unsupported native asset: TEST")
	require.NotNil(assetCfgEnriched)
	require.Equal("USDC", assetCfgEnriched.Asset)
	require.Equal("", assetCfgEnriched.URL)
	require.Equal("TEST", assetCfgEnriched.Chain)
	require.NotEqual("", assetCfgEnriched.Driver)

	assetCfg.URL = ""
	assetCfg.Chain = ""
	assetCfgEnriched, err = s.Factory.EnrichAssetConfig(assetCfg)
	require.EqualError(err, "unsupported native asset: (empty)")
	require.NotNil(assetCfgEnriched)
	require.Equal("USDC", assetCfgEnriched.Asset)
	require.Equal("", assetCfgEnriched.URL)
	require.Equal("", assetCfgEnriched.Chain)
	require.NotEqual("", assetCfgEnriched.Driver)
}

func (s *CrosschainTestSuite) TestGetAssetID() {
	require := s.Require()
	assetID := xc.GetAssetIDFromAsset("USDC", "SOL")
	require.Equal(xc.AssetID("USDC.SOL"), assetID)
}

func (s *CrosschainTestSuite) TestGetAssetConfig() {
	require := s.Require()
	task, err := s.Factory.GetAssetConfig("USDC", "SOL")
	asset := task.GetAssetConfig()
	require.Nil(err)
	require.NotNil(asset)
	require.Equal("USDC", asset.Asset)
	require.Equal(xc.SOL, asset.NativeAsset)
}

func (s *CrosschainTestSuite) TestGetAssetConfigEdgeCases() {
	require := s.Require()
	task, err := s.Factory.GetAssetConfig("", "")
	asset := task.GetAssetConfig()
	require.NotNil(err)
	require.NotNil(asset)
	require.Equal("", asset.Asset)
	require.Equal(xc.NativeAsset(""), asset.NativeAsset)
}

func (s *CrosschainTestSuite) TestGetTaskConfig() {
	require := s.Require()
	asset, err := s.Factory.GetTaskConfig("sol-wrap", "SOL")
	require.Nil(err)
	require.NotNil(asset)
}

func (s *CrosschainTestSuite) TestGetTaskConfigEdgeCases() {
	require := s.Require()
	singleAsset, _ := s.Factory.GetAssetConfig("USDC", "SOL")
	asset, err := s.Factory.GetTaskConfig("", "USDC.SOL")
	require.Nil(err)
	require.NotNil(singleAsset)
	require.NotNil(asset)
	require.Equal(singleAsset, asset)
}

func (s *CrosschainTestSuite) TestGetMultiAssetConfig() {
	require := s.Require()
	asset, err := s.Factory.GetMultiAssetConfig("SOL", "WSOL.SOL")
	require.Nil(err)
	require.NotNil(asset)
}

func (s *CrosschainTestSuite) TestGetMultiAssetConfigEdgeCases() {
	require := s.Require()
	singleAsset, _ := s.Factory.GetAssetConfig("USDC", "SOL")
	tasks, err := s.Factory.GetMultiAssetConfig("USDC.SOL", "")
	require.Nil(err)
	require.NotNil(singleAsset)
	require.NotNil(tasks)
	require.NotNil(tasks[0])
	require.Equal(singleAsset, tasks[0])
}

func (s *CrosschainTestSuite) TestGetAssetConfigByContract() {
	require := s.Require()

	assetI, err := s.Factory.GetAssetConfigByContract("0xB4FBF271143F4FBf7B91A5ded31805e42b2208d6", "ETH")
	asset := assetI.GetAssetConfig()
	require.Nil(err)
	require.NotNil(asset)
	require.Equal("WETH", asset.Asset)

	assetI, err = s.Factory.GetAssetConfigByContract("4zMMC9srt5Ri5X14GAgXhaHii3GnPAEERYPJgZJDncDU", "SOL")
	asset = assetI.GetAssetConfig()
	require.Nil(err)
	require.NotNil(asset)
	require.Equal("USDC", asset.Asset)

	assetI, err = s.Factory.GetAssetConfigByContract("0x123456", "ETH")
	asset = assetI.GetAssetConfig()
	require.EqualError(err, "invalid contract: '0x123456'")
	require.NotNil(asset)
	require.Equal("", asset.Asset)
}

func (s *CrosschainTestSuite) TestNormalizeAddressString() {
	require := s.Require()
	address := ""

	address = NormalizeAddressString("myaddress", "BTC")
	require.Equal("myaddress", address) // no normalization

	address = NormalizeAddressString("bitcoincash:myaddress", "BCH")
	require.Equal("myaddress", address)

	address = NormalizeAddressString("0x0ECE", "")
	require.Equal("0x0ece", address) // lowercase

	address = NormalizeAddressString("0x0ECE", "ETH")
	require.Equal("0x0ece", address)
}

func (s *CrosschainTestSuite) TestPutAssetConfig() {
	require := s.Require()
	assetName := "TEST"

	assetI, err := s.Factory.GetAssetConfig(assetName, "")
	require.EqualError(err, "invalid asset: 'TEST.ETH'")
	require.NotNil(assetI)

	assetI, err = s.Factory.PutAssetConfig(&xc.TokenAssetConfig{Asset: assetName, Chain: "ETH"})
	fmt.Println(assetI)
	asset := assetI.GetAssetConfig()
	require.Nil(err)
	require.Equal(assetName, asset.Asset)

	assetI, err = s.Factory.GetAssetConfig("TEST", "")
	asset = assetI.GetAssetConfig()
	require.Nil(err)
	require.Equal(assetName, asset.Asset)
}

func (s *CrosschainTestSuite) TestConfig() {
	require := s.Require()
	cfg := s.Factory.Config()
	require.NotNil(cfg)
}

func (s *CrosschainTestSuite) TestTxInputSerDeser() {
	require := s.Require()

	// Solana
	inputSolana := solana.NewTxInput()
	inputSolana.RecentBlockHash = [32]byte{1, 2, 3}
	inputSolana.ToIsATA = true
	inputSolana.ShouldCreateATA = true
	ser, err := s.Factory.MarshalTxInput(inputSolana)
	require.NoError(err)

	deser, err := s.Factory.UnmarshalTxInput(ser)
	require.NoError(err)
	typedSolana := deser.(*solana.TxInput)
	require.NotNil(typedSolana)
	require.Equal(inputSolana, typedSolana)

	// Cosmos
	inputCosmos := cosmos.NewTxInput()
	inputCosmos.FromPublicKey = []byte{1, 2, 3}
	inputCosmos.AccountNumber = 1
	inputCosmos.Sequence = 2
	inputCosmos.GasLimit = 3
	inputCosmos.GasPrice = 4.5
	inputCosmos.Memo = "memo"
	ser, err = s.Factory.MarshalTxInput(inputCosmos)
	require.NoError(err)

	deser, err = s.Factory.UnmarshalTxInput(ser)
	require.NoError(err)
	typedCosmos := deser.(*cosmos.TxInput)
	require.NotNil(typedCosmos)
	expected := inputCosmos
	require.Equal(expected, typedCosmos)

	inputBtc := bitcoin.NewTxInput()
	inputBtc.UnspentOutputs = []bitcoin.Output{
		{
			Outpoint: bitcoin.Outpoint{
				Index: 1,
			},
			Value: xc.NewAmountBlockchainFromUint64(100),
		},
		{
			Outpoint: bitcoin.Outpoint{
				Index: 2,
			},
			Value: xc.NewAmountBlockchainFromUint64(200),
		},
	}
	btcBz, err := MarshalTxInput(inputBtc)
	require.NoError(err)
	inputBtc2, err := UnmarshalTxInput(btcBz)
	require.NoError(err)

	require.Equal(inputBtc.UnspentOutputs[0].Value.String(), "100")
	require.Equal(inputBtc.UnspentOutputs[1].Value.String(), "200")
	require.EqualValues(inputBtc2.(*bitcoin.TxInput).UnspentOutputs[0].Outpoint.Index, 1)
	require.EqualValues(inputBtc2.(*bitcoin.TxInput).UnspentOutputs[1].Outpoint.Index, 2)
	fmt.Println("SERDE:\n", string(btcBz))
	require.Equal(inputBtc2.(*bitcoin.TxInput).UnspentOutputs[0].Value.String(), "100")
	require.Equal(inputBtc2.(*bitcoin.TxInput).UnspentOutputs[1].Value.String(), "200")

}

func (s *CrosschainTestSuite) TestAllTxInputSerDeser() {
	require := s.Require()
	for _, driver := range xc.SupportedDrivers {
		var input xc.TxInput
		switch driver {
		case xc.DriverEVM, xc.DriverEVMLegacy:
			input = evm.NewTxInput()
		case xc.DriverCosmos, xc.DriverCosmosEvmos:
			input = cosmos.NewTxInput()
		case xc.DriverSolana:
			input = solana.NewTxInput()
		case xc.DriverAptos:
			input = aptos.NewTxInput()
		case xc.DriverBitcoin:
			input = bitcoin.NewTxInput()
		case xc.DriverSui:
			input = bitcoin.NewTxInput()
		default:
			require.Fail("must add driver to test: " + string(driver))
		}
		bz, err := MarshalTxInput(input)
		require.NoError(err)
		_, err = UnmarshalTxInput(bz)
		require.NoError(err)
	}
}

func (s *CrosschainTestSuite) TestSigAlg() {
	require := s.Require()
	for _, driver := range xc.SupportedDrivers {
		require.NotEmpty(driver.SignatureAlgorithm())
	}
}
