package sui

import (
	"encoding/hex"

	xc "github.com/jumpcrypto/crosschain"
)

func (s *CrosschainTestSuite) TestNewAddressBuilder() {
	require := s.Require()
	builder, err := NewAddressBuilder(&xc.AssetConfig{})
	require.Nil(err)
	require.NotNil(builder)
}

func (s *CrosschainTestSuite) TestGetAddressFromPublicKey() {
	require := s.Require()
	builder, _ := NewAddressBuilder(&xc.AssetConfig{})
	bytes, _ := hex.DecodeString("E0651D94176024B0C137C23A782D50291C04C8B5BCEDD4A7CD066BF4C0D21B22")
	address, err := builder.GetAddressFromPublicKey(bytes)
	require.Nil(err)
	require.Equal(xc.Address("0xa589a80d61ec380c24a5fdda109c3848c082584e6cb725e5ab19b18354b2ab85"), address)

	bytes, _ = hex.DecodeString("E0651D94176024B0C137C23A782D50291C04C8B5BCEDD4A7CD066BF4C0D21B2200")
	address, err = builder.GetAddressFromPublicKey(bytes)
	require.Nil(err)
	require.Equal(xc.Address("0xa589a80d61ec380c24a5fdda109c3848c082584e6cb725e5ab19b18354b2ab85"), address)
}

func (s *CrosschainTestSuite) TestGetAddressFromPublicKeyErr() {
	require := s.Require()
	builder, _ := NewAddressBuilder(&xc.AssetConfig{})

	address, err := builder.GetAddressFromPublicKey([]byte{})
	require.Equal(xc.Address(""), address)
	require.EqualError(err, "invalid length for ed25519 public key")

	address, err = builder.GetAddressFromPublicKey([]byte{1, 2, 3})
	require.Equal(xc.Address(""), address)
	require.EqualError(err, "invalid length for ed25519 public key")

	bytes, _ := hex.DecodeString("E0651D94176024B0C137C23A782D50291C04C8B5BCEDD4A7CD066BF4C0D21B2201")
	address, err = builder.GetAddressFromPublicKey(bytes)
	require.Equal(xc.Address(""), address)
	require.EqualError(err, "invalid format for ed25519 public key")
}

func (s *CrosschainTestSuite) TestGetAllPossibleAddressesFromPublicKey() {
	require := s.Require()
	builder, _ := NewAddressBuilder(&xc.AssetConfig{NativeAsset: "LUNA", ChainPrefix: "terra"})
	bytes, _ := hex.DecodeString("E0651D94176024B0C137C23A782D50291C04C8B5BCEDD4A7CD066BF4C0D21B22")
	addresses, err := builder.GetAllPossibleAddressesFromPublicKey(bytes)
	require.Nil(err)
	require.Equal(1, len(addresses))
	require.Equal(xc.Address("0xa589a80d61ec380c24a5fdda109c3848c082584e6cb725e5ab19b18354b2ab85"), addresses[0].Address)
	require.Equal(xc.AddressTypeDefault, addresses[0].Type)
}
