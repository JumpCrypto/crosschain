package factory

import (
	"testing"

	xc "github.com/jumpcrypto/crosschain"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/suite"
)

type CrosschainTestSuite struct {
	suite.Suite
	Factory          Factory
	TestNativeAssets []xc.NativeAsset
	TestAssetConfigs []xc.AssetConfig
}

func (s *CrosschainTestSuite) SetupTest() {
	s.Factory = NewDefaultFactory()
	s.TestNativeAssets = []xc.NativeAsset{
		xc.ETH,
		xc.MATIC,
		xc.BNB,
		xc.SOL,
		xc.ATOM,
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

	asset, _ := s.Factory.PutAssetConfig(xc.AssetConfig{Asset: "TEST"})
	_, err := s.Factory.NewClient(asset)
	require.EqualError(err, "unsupported asset")
}

func (s *CrosschainTestSuite) TestNewTxBuilder() {
	require := s.Require()
	for _, asset := range s.TestAssetConfigs {
		builder, _ := s.Factory.NewTxBuilder(asset)
		require.NotNil(builder)
	}

	asset, _ := s.Factory.PutAssetConfig(xc.AssetConfig{Asset: "TEST"})
	_, err := s.Factory.NewTxBuilder(asset)
	require.EqualError(err, "unsupported asset")
}

func (s *CrosschainTestSuite) TestNewSigner() {
	require := s.Require()
	for _, asset := range s.TestAssetConfigs {
		signer, _ := s.Factory.NewSigner(asset)
		require.NotNil(signer)
	}

	asset, _ := s.Factory.PutAssetConfig(xc.AssetConfig{Asset: "TEST"})
	_, err := s.Factory.NewSigner(asset)
	require.EqualError(err, "unsupported asset")
}

func (s *CrosschainTestSuite) TestNewAddressBuilder() {
	require := s.Require()
	for _, asset := range s.TestAssetConfigs {
		builder, _ := s.Factory.NewAddressBuilder(asset)
		require.NotNil(builder)
	}

	asset, _ := s.Factory.PutAssetConfig(xc.AssetConfig{Asset: "TEST"})
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
		address := s.Factory.MustAddress(asset, "myaddress") // trivial impl
		require.Equal(xc.Address("myaddress"), address, "Error on: "+asset.NativeAsset)
	}
}

func (s *CrosschainTestSuite) TestMustPrivateKey() {
	require := s.Require()
	for _, asset := range s.TestAssetConfigs {
		privateKey := s.Factory.MustPrivateKey(asset, "myprivatekey")
		require.NotNil(privateKey, "Error on: "+asset.NativeAsset)
	}
}

// Convert functions

func (s *CrosschainTestSuite) TestConvertAmountToHuman() {
	require := s.Require()
	var amountBlockchain xc.AmountBlockchain
	for _, asset := range s.TestAssetConfigs {
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
}

func (s *CrosschainTestSuite) TestConvertAmountToBlockchain() {
	require := s.Require()
	amountDecimal, _ := decimal.NewFromString("10.3")
	amountHuman := xc.AmountHumanReadable(amountDecimal)

	var expected xc.AmountBlockchain
	for _, asset := range s.TestAssetConfigs {
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
		amount, err := s.Factory.ConvertAmountStrToBlockchain(asset, "10.3")

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

	asset, _ := s.Factory.PutAssetConfig(xc.AssetConfig{Asset: "TEST"})
	_, err := s.Factory.ConvertAmountStrToBlockchain(asset, "10.3")
	require.EqualError(err, "unsupported asset")
}

func (s *CrosschainTestSuite) TestConvertAmountStrToBlockchainErr() {
	require := s.Require()
	for _, asset := range s.TestAssetConfigs {
		_, err := s.Factory.ConvertAmountStrToBlockchain(asset, "err")
		require.EqualError(err, "can't convert err to decimal: exponent is not numeric")
	}
}

// Config functions

func (s *CrosschainTestSuite) TestEnrichAssetConfig() {
	require := s.Require()

	assetCfg, _ := s.Factory.GetAssetConfig("USDC", "SOL")
	assetCfg.URL = ""
	assetCfgEnriched, err := s.Factory.EnrichAssetConfig(assetCfg)
	require.Nil(err)
	require.NotNil(assetCfgEnriched)
	require.Equal("USDC", assetCfgEnriched.Asset)
	require.NotEqual("", assetCfgEnriched.URL)

	assetCfg.URL = ""
	assetCfg.Chain = "TEST"
	assetCfgEnriched, err = s.Factory.EnrichAssetConfig(assetCfg)
	require.EqualError(err, "unsupported native asset: TEST")
	require.NotNil(assetCfgEnriched)
	require.Equal("USDC", assetCfgEnriched.Asset)
	require.Equal("", assetCfgEnriched.URL)
	require.Equal("TEST", assetCfgEnriched.Chain)

	assetCfg.URL = ""
	assetCfg.Chain = ""
	assetCfgEnriched, err = s.Factory.EnrichAssetConfig(assetCfg)
	require.EqualError(err, "unsupported native asset: (empty)")
	require.NotNil(assetCfgEnriched)
	require.Equal("USDC", assetCfgEnriched.Asset)
	require.Equal("", assetCfgEnriched.URL)
	require.Equal("", assetCfgEnriched.Chain)
}

func (s *CrosschainTestSuite) TestGetAssetID() {
	require := s.Require()
	assetID := s.Factory.GetAssetID("USDC", "SOL")
	require.Equal(xc.AssetID("USDC.SOL"), assetID)
}

func (s *CrosschainTestSuite) TestGetAssetConfig() {
	require := s.Require()
	asset, err := s.Factory.GetAssetConfig("USDC", "SOL")
	require.Nil(err)
	require.NotNil(asset)
	require.Equal("USDC", asset.Asset)
	require.Equal(xc.SOL, asset.NativeAsset)
}

func (s *CrosschainTestSuite) TestGetAssetConfigByContract() {
	require := s.Require()

	asset, err := s.Factory.GetAssetConfigByContract("0xc778417E063141139Fce010982780140Aa0cD5Ab", "ETH")
	require.Nil(err)
	require.NotNil(asset)
	require.Equal("WETH", asset.Asset)

	asset, err = s.Factory.GetAssetConfigByContract("4zMMC9srt5Ri5X14GAgXhaHii3GnPAEERYPJgZJDncDU", "SOL")
	require.Nil(err)
	require.NotNil(asset)
	require.Equal("USDC", asset.Asset)

	asset, err = s.Factory.GetAssetConfigByContract("0x123456", "ETH")
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

	asset, err := s.Factory.GetAssetConfig(assetName, "")
	require.EqualError(err, "invalid asset: 'TEST'")
	require.NotNil(asset)

	asset, err = s.Factory.PutAssetConfig(xc.AssetConfig{Asset: assetName})
	require.Nil(err)
	require.Equal(assetName, asset.Asset)

	asset, err = s.Factory.GetAssetConfig("TEST", "")
	require.Nil(err)
	require.Equal(assetName, asset.Asset)
}

func (s *CrosschainTestSuite) TestConfig() {
	require := s.Require()
	cfg := s.Factory.Config()
	require.NotNil(cfg)
}
