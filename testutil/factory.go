package testutil

import (
	xc "github.com/jumpcrypto/crosschain"
	"github.com/jumpcrypto/crosschain/factory"
)

// TestFactory for unit tests
type TestFactory struct {
	DefaultFactory factory.FactoryContext

	NewClientFunc               func(asset xc.AssetConfig) (xc.Client, error)
	NewTxBuilderFunc            func(asset xc.AssetConfig) (xc.TxBuilder, error)
	NewSignerFunc               func(asset xc.AssetConfig) (xc.Signer, error)
	NewAddressBuilderFunc       func(asset xc.AssetConfig) (xc.AddressBuilder, error)
	GetAddressFromPublicKeyFunc func(asset xc.AssetConfig, publicKey []byte) (xc.Address, error)
}

var _ factory.FactoryContext = &TestFactory{}

// NewClient creates a new Client
func (f *TestFactory) NewClient(asset xc.AssetConfig) (xc.Client, error) {
	if f.NewClientFunc != nil {
		return f.NewClientFunc(asset)
	}
	return f.DefaultFactory.NewClient(asset)
}

// NewTxBuilder creates a new TxBuilder
func (f *TestFactory) NewTxBuilder(asset xc.AssetConfig) (xc.TxBuilder, error) {
	if f.NewTxBuilderFunc != nil {
		return f.NewTxBuilderFunc(asset)
	}
	return f.DefaultFactory.NewTxBuilder(asset)
}

// NewSigner creates a new Signer
func (f *TestFactory) NewSigner(asset xc.AssetConfig) (xc.Signer, error) {
	if f.NewSignerFunc != nil {
		return f.NewSignerFunc(asset)
	}
	return f.DefaultFactory.NewSigner(asset)
}

// NewAddressBuilder creates a new AddressBuilder
func (f *TestFactory) NewAddressBuilder(asset xc.AssetConfig) (xc.AddressBuilder, error) {
	if f.NewAddressBuilderFunc != nil {
		return f.NewAddressBuilderFunc(asset)
	}
	return f.DefaultFactory.NewAddressBuilder(asset)

}

// GetAddressFromPublicKey returns an Address given a public key
func (f *TestFactory) GetAddressFromPublicKey(asset xc.AssetConfig, publicKey []byte) (xc.Address, error) {
	if f.GetAddressFromPublicKeyFunc != nil {
		return f.GetAddressFromPublicKeyFunc(asset, publicKey)
	}
	return f.DefaultFactory.GetAddressFromPublicKey(asset, publicKey)
}

// GetAllPossibleAddressesFromPublicKey returns all PossibleAddress(es) given a public key
func (f *TestFactory) GetAllPossibleAddressesFromPublicKey(asset xc.AssetConfig, publicKey []byte) ([]xc.PossibleAddress, error) {
	if f.GetAddressFromPublicKeyFunc != nil {
		return f.GetAllPossibleAddressesFromPublicKey(asset, publicKey)
	}
	return f.DefaultFactory.GetAllPossibleAddressesFromPublicKey(asset, publicKey)
}

// ConvertAmountToHuman converts an AmountBlockchain into AmountHumanReadable, dividing by the appropriate number of decimals
func (f *TestFactory) ConvertAmountToHuman(asset xc.AssetConfig, blockchainAmount xc.AmountBlockchain) (xc.AmountHumanReadable, error) {
	return f.DefaultFactory.ConvertAmountToHuman(asset, blockchainAmount)
}

// ConvertAmountToBlockchain converts an AmountHumanReadable into AmountBlockchain, multiplying by the appropriate number of decimals
func (f *TestFactory) ConvertAmountToBlockchain(asset xc.AssetConfig, humanAmount xc.AmountHumanReadable) (xc.AmountBlockchain, error) {
	return f.DefaultFactory.ConvertAmountToBlockchain(asset, humanAmount)
}

// ConvertAmountStrToBlockchain converts a string representing an AmountHumanReadable into AmountBlockchain, multiplying by the appropriate number of decimals
func (f *TestFactory) ConvertAmountStrToBlockchain(asset xc.AssetConfig, humanAmountStr string) (xc.AmountBlockchain, error) {
	return f.DefaultFactory.ConvertAmountStrToBlockchain(asset, humanAmountStr)
}

// GetAssetID returns a canonical AssetID
func (f *TestFactory) GetAssetID(asset string, nativeAsset string) xc.AssetID {
	return f.DefaultFactory.GetAssetID(asset, nativeAsset)
}

// GetAssetConfig returns an AssetConfig by asset and native asset (chain)
func (f *TestFactory) GetAssetConfig(asset string, nativeAsset string) (xc.AssetConfig, error) {
	return f.DefaultFactory.GetAssetConfig(asset, nativeAsset)
}

// GetAssetConfigByContract returns an AssetConfig by contract and native asset (chain)
func (f *TestFactory) GetAssetConfigByContract(contract string, nativeAsset string) (xc.AssetConfig, error) {
	return f.DefaultFactory.GetAssetConfigByContract(contract, nativeAsset)
}

// EnrichAssetConfig augments a partial AssetConfig, for example if some info is stored in a db and other in a config file
func (f *TestFactory) EnrichAssetConfig(partialCfg xc.AssetConfig) (xc.AssetConfig, error) {
	return f.DefaultFactory.EnrichAssetConfig(partialCfg)
}

// PutAssetConfig adds an AssetConfig to the current Config cache
func (f *TestFactory) PutAssetConfig(config xc.AssetConfig) (xc.AssetConfig, error) {
	return f.DefaultFactory.PutAssetConfig(config)
}

// Config returns the Config
func (f *TestFactory) Config() interface{} {
	return f.DefaultFactory.Config()
}

// MustAddress coverts a string to Address, panic if error
func (f *TestFactory) MustAddress(asset xc.AssetConfig, addressStr string) xc.Address {
	return f.DefaultFactory.MustAddress(asset, addressStr)
}

// MustAmountBlockchain coverts a string into AmountBlockchain, panic if error
func (f *TestFactory) MustAmountBlockchain(asset xc.AssetConfig, humanAmountStr string) xc.AmountBlockchain {
	return f.DefaultFactory.MustAmountBlockchain(asset, humanAmountStr)

}

// MustPrivateKey coverts a string into PrivateKey, panic if error
func (f *TestFactory) MustPrivateKey(asset xc.AssetConfig, privateKeyStr string) xc.PrivateKey {
	return f.DefaultFactory.MustPrivateKey(asset, privateKeyStr)
}

// NewDefaultFactory creates a new Factory
func NewDefaultFactory() TestFactory {
	f := factory.NewDefaultFactory()
	return TestFactory{
		DefaultFactory: f,
	}
}
