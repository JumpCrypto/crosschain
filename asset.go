package crosschain

import (
	"fmt"
	"strings"
)

// Asset is an asset on a blockchain. It can be a token or native asset.
type Asset string

// NativeAsset is an asset on a blockchain used to pay gas fees.
// In Crosschain, for simplicity, a NativeAsset represents a chain.
type NativeAsset Asset

// List of supported NativeAsset
const (
	// UTXO
	BCH  = NativeAsset("BCH")  // Bitcoin Cash
	BTC  = NativeAsset("BTC")  // Bitcoin
	DOGE = NativeAsset("DOGE") // Dogecoin
	LTC  = NativeAsset("LTC")  // Litecoin

	// Account-based
	ACA    = NativeAsset("ACA")    // Acala
	APTOS  = NativeAsset("APTOS")  // APTOS
	ArbETH = NativeAsset("ArbETH") // Arbitrum Ether
	AurETH = NativeAsset("AurETH") // Aurora
	AVAX   = NativeAsset("AVAX")   // Avalanche
	BNB    = NativeAsset("BNB")    // Binance Coin
	CELO   = NativeAsset("CELO")   // Celo
	CHZ    = NativeAsset("CHZ")    // Chiliz
	ETC    = NativeAsset("ETC")    // Ethereum Classic
	ETH    = NativeAsset("ETH")    // Ether
	FTM    = NativeAsset("FTM")    // Fantom
	KAR    = NativeAsset("KAR")    // Karura
	KLAY   = NativeAsset("KLAY")   // Klaytn
	MATIC  = NativeAsset("MATIC")  // Matic PoS (Polygon)
	OptETH = NativeAsset("OptETH") // Optimism
	ROSE   = NativeAsset("ROSE")   // Rose (Oasis)
	SOL    = NativeAsset("SOL")    // Solana
)

// AssetType is the type of an asset, either native or token
type AssetType string

// List of supported AssetType
const (
	AssetTypeNative = AssetType("native")
	AssetTypeToken  = AssetType("token")
)

// AssetType returns the type of an Asset
func (asset Asset) AssetType() AssetType {
	switch native := NativeAsset(asset); native {
	case BCH, BTC, DOGE, LTC:
		return AssetTypeNative
	case ACA,
		APTOS,
		ArbETH,
		AurETH,
		AVAX,
		BNB,
		CELO,
		CHZ,
		ETC,
		ETH,
		FTM,
		KAR,
		KLAY,
		MATIC,
		OptETH,
		ROSE,
		SOL:
		return AssetTypeNative
	default:
		return AssetTypeToken
	}
}

// ChainType is the type of a chain
type ChainType string

// List of supported ChainType
const (
	ChainTypeUnknown = ChainType("unknown")
	ChainTypeUTXO    = ChainType("utxo")
	ChainTypeAccount = ChainType("account")
)

// ChainType returns the type of a chain, represented as its NativeAsset
func (native NativeAsset) ChainType() ChainType {
	switch native {
	case BCH, BTC, DOGE, LTC:
		return ChainTypeUTXO
	case ACA,
		APTOS,
		ArbETH,
		AurETH,
		AVAX,
		BNB,
		CELO,
		CHZ,
		ETC,
		ETH,
		FTM,
		KAR,
		KLAY,
		MATIC,
		OptETH,
		ROSE,
		SOL:
		return ChainTypeAccount
	default:
		return ChainTypeUnknown
	}
}

// Driver is the type of a chain
type Driver string

// List of supported Driver
const (
	DriverAptos       = Driver("aptos")
	DriverBitcoin     = Driver("bitcoin")
	DriverCosmos      = Driver("cosmos")
	DriverCosmosEvmos = Driver("evmos")
	DriverEVM         = Driver("evm")
	DriverEVMLegacy   = Driver("evm-legacy")
	DriverSolana      = Driver("solana")
)

// AssetID is an internal identifier for each asset
// Examples: ETH, USDC, USDC.SOL - see tests for details
type AssetID string

// AssetConfig is the model used to represent an asset read from config file or db
type AssetConfig struct {
	// 	[[silochain.beta.chains]]
	//     asset = "eth"
	//     net = "mainnet"
	//     url = "http://7.125.36.22:8089"
	//
	//   [[silochain.beta.chains]]
	//     asset = "usdc"
	//     chain = "eth"
	//     net = "mainnet"
	//     contract = "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"
	//     decimals = 6
	Asset                string  `yaml:"asset"`
	Driver               string  `yaml:"driver"`
	Net                  string  `yaml:"net"`
	URL                  string  `yaml:"url"`
	FcdURL               string  `yaml:"fcd_url"`
	Auth                 string  `yaml:"auth"`
	Provider             string  `yaml:"provider"`
	ChainID              int64   `yaml:"chain_id"`
	ChainIDStr           string  `yaml:"chain_id_str"`
	ChainName            string  `yaml:"chain_name"`
	ChainPrefix          string  `yaml:"chain_prefix"`
	ChainCoin            string  `yaml:"chain_coin"`
	ChainCoinHDPath      uint32  `yaml:"chain_coin_hd_path"`
	ChainGasPriceDefault float64 `yaml:"chain_gas_price_default"`
	ChainGasMultiplier   float64 `yaml:"chain_gas_multiplier"`
	ExplorerURL          string  `yaml:"explorer_url"`

	// Tokens
	Chain    string `yaml:"chain"`
	Contract string `yaml:"contract"`
	Decimals int32  `yaml:"decimals"`
	Name     string `yaml:"name"`

	// Not used for serde
	ID          AssetID     `yaml:"-"`
	AuthSecret  string      `yaml:"-"`
	NativeAsset NativeAsset `yaml:"-"`
	Type        AssetType   `yaml:"-"`
}

// Config is the full config containing all Assets
type Config struct {
	AllAssets []AssetConfig `yaml:"chains"`
}

func (c AssetConfig) String() string {
	// do NOT print AuthSecret
	return fmt.Sprintf("net: %s, url: %s, auth: %s, provider: %s", c.Net, c.URL, c.Auth, c.Provider)
}

func parseAssetAndNativeAsset(asset string, nativeAsset string) (string, string) {
	if asset == "" && nativeAsset == "" {
		return "", ""
	}
	if asset == "" && nativeAsset != "" {
		asset = nativeAsset
	}

	assetSplit := strings.Split(asset, ".")
	if len(assetSplit) == 2 && Asset(assetSplit[1]).AssetType() == AssetTypeNative {
		asset = assetSplit[0]
		if nativeAsset == "" {
			nativeAsset = assetSplit[1]
		}
	}
	validNative := Asset(asset).AssetType() == AssetTypeNative

	if nativeAsset == "" {
		if validNative {
			nativeAsset = asset
		} else {
			nativeAsset = "ETH"
		}
	}
	nativeAsset = strings.ToUpper(nativeAsset)

	return asset, nativeAsset
}

// GetAssetIDFromAsset return the canonical AssetID given two input strings asset, nativeAsset.
// Input can come from user input.
// Examples:
// - GetAssetIDFromAsset("USDC", "") -> "USDC"
// - GetAssetIDFromAsset("USDC", "ETH") -> "USDC"
// - GetAssetIDFromAsset("USDC", "SOL") -> "USDC.SOL"
// - GetAssetIDFromAsset("USDC.SOL", "") -> "USDC.SOL"
// See tests for more examples.
func GetAssetIDFromAsset(asset string, nativeAsset string) AssetID {
	// id is SYMBOL for ERC20 and SYMBOL.CHAIN for others
	// e.g. BTC, ETH, USDC, SOL, USDC.SOL
	asset, nativeAsset = parseAssetAndNativeAsset(asset, nativeAsset)
	asset = strings.ToUpper(asset)
	validNative := Asset(asset).AssetType() == AssetTypeNative

	// native asset, e.g. BTC, ETH, SOL
	if asset == nativeAsset {
		return AssetID(asset)
	}
	if nativeAsset == "ETH" && !validNative {
		return AssetID(asset)
	}
	// token, e.g. USDC, USDC.SOL
	return AssetID(asset + "." + nativeAsset)
}
