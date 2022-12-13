package cosmos

import (
	"encoding/hex"

	xc "github.com/jumpcrypto/crosschain"
)

func (s *CrosschainTestSuite) TestNewAddressBuilder() {
	require := s.Require()
	builder, err := NewAddressBuilder(xc.AssetConfig{})
	require.Nil(err)
	require.NotNil(builder)
}

func (s *CrosschainTestSuite) TestGetAddressFromPublicKey() {
	require := s.Require()
	builder, _ := NewAddressBuilder(xc.AssetConfig{NativeAsset: "LUNA", ChainPrefix: "terra"})
	bytes, _ := hex.DecodeString("02FCF724C97DFFAC2021EFA1818C2FEF3BCBB753CA22913A8DB5E79EC4A3DEE0D1")
	address, err := builder.GetAddressFromPublicKey(bytes)
	require.Nil(err)
	require.Equal(xc.Address("terra1dp3q305hgttt8n34rt8rg9xpanc42z4ye7upfg"), address)
}

func (s *CrosschainTestSuite) TestGetAddressFromPublicKeyEvmos() {
	require := s.Require()
	builder, _ := NewAddressBuilder(xc.AssetConfig{NativeAsset: "XPLA", ChainPrefix: "xpla", Driver: string(xc.DriverCosmosEvmos)})
	bytes, _ := hex.DecodeString("02E8445082A72F29B75CA48748A914DF60622A609CACFCE8ED0E35804560741D29")
	address, err := builder.GetAddressFromPublicKey(bytes)
	require.Nil(err)
	require.Equal(xc.Address("xpla1r56x9533ntqtlsd99cth48fhyjf82gfstgvk9m"), address)
}

func (s *CrosschainTestSuite) TestGetAddressFromPublicKeyErr() {
	require := s.Require()
	builder, _ := NewAddressBuilder(xc.AssetConfig{})

	require.Panics(func() {
		// cosmos-sdk panics with "length of pubkey is incorrect"
		_, _ = builder.GetAddressFromPublicKey([]byte{})
	})

	require.Panics(func() {
		// cosmos-sdk panics with "length of pubkey is incorrect"
		_, _ = builder.GetAddressFromPublicKey([]byte{1, 2, 3})
	})

	// AssetConfig.ChainPrefix is needed to bech32ify
	pubKeyBytes, _ := hex.DecodeString("02E8445082A72F29B75CA48748A914DF60622A609CACFCE8ED0E35804560741D29")
	address, err := builder.GetAddressFromPublicKey(pubKeyBytes)
	require.Equal(xc.Address(""), address)
	require.EqualError(err, "prefix cannot be empty")

	// cosmos-sdk doesn't check if pubkey is on the curve
	builder, _ = NewAddressBuilder(xc.AssetConfig{NativeAsset: "LUNA", ChainPrefix: "terra"})
	bytes, _ := hex.DecodeString("001122334455667788990011223344556677889900112233445566778899001122")
	address, err = builder.GetAddressFromPublicKey(bytes)
	require.Nil(err)
	require.Equal(xc.Address("terra1hw58t56mzszlnnkjak83ul8ff437ylrz57xj4v"), address)

	// ethermint doesn't check if pubkey is on the curve,
	// but it attempts to decompress the point to generate the address
	// therefore indirectly it catches the error
	builder, _ = NewAddressBuilder(xc.AssetConfig{NativeAsset: "XPLA", ChainPrefix: "xpla", Driver: string(xc.DriverCosmosEvmos)})
	bytes, _ = hex.DecodeString("001122334455667788990011223344556677889900112233445566778899001122")
	address, err = builder.GetAddressFromPublicKey(bytes)
	require.ErrorContains(err, "addresses cannot be empty")
	require.Equal(xc.Address(""), address)
}

func (s *CrosschainTestSuite) TestGetAllPossibleAddressesFromPublicKey() {
	require := s.Require()
	builder, _ := NewAddressBuilder(xc.AssetConfig{NativeAsset: "LUNA", ChainPrefix: "terra"})
	bytes, _ := hex.DecodeString("02E8445082A72F29B75CA48748A914DF60622A609CACFCE8ED0E35804560741D29")
	addresses, err := builder.GetAllPossibleAddressesFromPublicKey(bytes)
	require.Nil(err)
	require.Equal(1, len(addresses))
	require.Equal(xc.Address("terra1mzqd0kynsjzsnf3d37m5uvs53kkxssf0aasf27"), addresses[0].Address)
	require.Equal(xc.AddressTypeDefault, addresses[0].Type)
}
