package crosschain

func (s *CrosschainTestSuite) TestTypesAssetVsNativeAsset() {
	require := s.Require()
	require.Equal(NativeAsset("SOL"), SOL)
	require.NotEqual(Asset("SOL"), SOL)
}

func (s *CrosschainTestSuite) TestAssetType() {
	require := s.Require()
	require.Equal(AssetTypeNative, Asset(BTC).AssetType())
	require.Equal(AssetTypeNative, Asset(ETH).AssetType())
	require.Equal(AssetTypeNative, Asset("ETH").AssetType())
	require.Equal(AssetTypeToken, Asset("WETH").AssetType())
	require.Equal(AssetTypeToken, Asset("unknown").AssetType())
}

func (s *CrosschainTestSuite) TestChainType() {
	require := s.Require()
	require.Equal(ChainTypeUTXO, BTC.ChainType())
	require.Equal(ChainTypeAccount, ETH.ChainType())
	require.Equal(ChainTypeAccount, NativeAsset("ETH").ChainType())
	require.Equal(ChainTypeUnknown, NativeAsset("unknown").ChainType())
}

func (s *CrosschainTestSuite) TestAssetConfig() {
	require := s.Require()
	assetConfig := AssetConfig{
		Asset:       "myasset",
		NativeAsset: "native",
		Net:         "mynet",
		URL:         "myurl",
		Auth:        "myauth",
		Provider:    "myprovider",
		AuthSecret:  "SECRET",
	}
	require.Equal("NativeAssetConfig(id=myasset asset=myasset chainId=0 driver= type= chainCoin= prefix= net=mynet url=myurl auth=myauth provider=myprovider native_asset=native)", assetConfig.String())
	require.NotContains(assetConfig.String(), "SECRET")
}

func (s *CrosschainTestSuite) TestParseAssetAndNativeAsset() {
	require := s.Require()
	var asset, native string

	asset, native = parseAssetAndNativeAsset("", "SOL")
	require.Equal("SOL", asset)
	require.Equal("SOL", native)

	asset, native = parseAssetAndNativeAsset("", "ETH")
	require.Equal("ETH", asset)
	require.Equal("ETH", native)

	asset, native = parseAssetAndNativeAsset("ETH", "ETH")
	require.Equal("ETH", asset)
	require.Equal("ETH", native)

	asset, native = parseAssetAndNativeAsset("USDC", "SOL")
	require.Equal("USDC", asset)
	require.Equal("SOL", native)

	asset, native = parseAssetAndNativeAsset("USDC", "ETH")
	require.Equal("USDC", asset)
	require.Equal("ETH", native)

	asset, native = parseAssetAndNativeAsset("USDC", "")
	require.Equal("USDC", asset)
	require.Equal("ETH", native)

	asset, native = parseAssetAndNativeAsset("USDC.SOL", "")
	require.Equal("USDC", asset)
	require.Equal("SOL", native)

	// WETH
	asset, native = parseAssetAndNativeAsset("WETH", "")
	require.Equal("WETH", asset)
	require.Equal("ETH", native)

	asset, native = parseAssetAndNativeAsset("WETH.ETH", "")
	require.Equal("WETH", asset)
	require.Equal("ETH", native)

	asset, native = parseAssetAndNativeAsset("WETH", "ETH")
	require.Equal("WETH", asset)
	require.Equal("ETH", native)

	asset, native = parseAssetAndNativeAsset("WETH", "MATIC")
	require.Equal("WETH", asset)
	require.Equal("MATIC", native)

	asset, native = parseAssetAndNativeAsset("WETH.MATIC", "")
	require.Equal("WETH", asset)
	require.Equal("MATIC", native)

	asset, native = parseAssetAndNativeAsset("WETH.SOL", "")
	require.Equal("WETH", asset)
	require.Equal("SOL", native)

	asset, native = parseAssetAndNativeAsset("WETH", "SOL")
	require.Equal("WETH", asset)
	require.Equal("SOL", native)
}

func (s *CrosschainTestSuite) TestParseAssetAndNativeAssetEdgeCases() {
	require := s.Require()
	var asset, native string

	asset, native = parseAssetAndNativeAsset("", "")
	require.Equal("", asset)
	require.Equal("", native)

	asset, native = parseAssetAndNativeAsset("", "test")
	require.Equal("test", asset)
	require.Equal("test", native)

	asset, native = parseAssetAndNativeAsset("USDC.sol", "") // invalid
	require.Equal("USDC.sol", asset)
	require.Equal("ETH", native)

	asset, native = parseAssetAndNativeAsset("USDC.WETH", "") // invalid
	require.Equal("USDC.WETH", asset)
	require.Equal("ETH", native)

	asset, native = parseAssetAndNativeAsset("USDC.ETH.SOL", "") // invalid
	require.Equal("USDC.ETH.SOL", asset)
	require.Equal("ETH", native)
}

func (s *CrosschainTestSuite) TestGetAssetIDFromAsset() {
	require := s.Require()

	require.Equal(AssetID(""), GetAssetIDFromAsset("", ""))

	require.Equal(AssetID("SOL"), GetAssetIDFromAsset("", "SOL"))
	require.Equal(AssetID("SOL"), GetAssetIDFromAsset("SOL", ""))
	require.Equal(AssetID("SOL"), GetAssetIDFromAsset("SOL", "SOL"))
	require.Equal(AssetID("SOL"), GetAssetIDFromAsset("SOL.SOL", ""))

	require.Equal(AssetID("ETH"), GetAssetIDFromAsset("", "ETH"))
	require.Equal(AssetID("ETH"), GetAssetIDFromAsset("ETH", ""))
	require.Equal(AssetID("ETH"), GetAssetIDFromAsset("ETH", "ETH"))
	require.Equal(AssetID("ETH"), GetAssetIDFromAsset("ETH.ETH", ""))

	require.Equal(AssetID("USDC.ETH"), GetAssetIDFromAsset("USDC", ""))
	require.Equal(AssetID("USDC.ETH"), GetAssetIDFromAsset("USDC", "ETH"))
	require.Equal(AssetID("USDC.ETH"), GetAssetIDFromAsset("USDC.ETH", ""))
	require.Equal(AssetID("USDC.SOL"), GetAssetIDFromAsset("USDC", "SOL"))
	require.Equal(AssetID("USDC.SOL"), GetAssetIDFromAsset("USDC.SOL", ""))

	require.Equal(AssetID("ArbETH"), GetAssetIDFromAsset("", "ArbETH"))
	require.Equal(AssetID("WETH.ArbETH"), GetAssetIDFromAsset("WETH", "ArbETH"))
	require.Equal(AssetID("WETH.ArbETH"), GetAssetIDFromAsset("WETH.ArbETH", ""))

	require.Equal(AssetID("INJ"), GetAssetIDFromAsset("INJ", "INJ"))
	require.Equal(AssetID("INJ.ETH"), GetAssetIDFromAsset("INJ", "ETH"))
	require.Equal(AssetID("INJ.SOL"), GetAssetIDFromAsset("INJ", "SOL"))

	require.Equal(AssetID("TEST.ETH"), GetAssetIDFromAsset("TEST", ""))
}
