package evm

import (
	"encoding/hex"

	"github.com/ethereum/go-ethereum/common"
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
	bytes, _ := hex.DecodeString("04760c4460e5336ac9bbd87952a3c7ec4363fc0a97bd31c86430806e287b437fd1b01abc6e1db640cf3106b520344af1d58b00b57823db3e1407cbc433e1b6d04d")
	address, err := builder.GetAddressFromPublicKey(bytes)
	require.NoError(err)
	require.Equal(xc.Address("0x5891906fEf64A5ae924C7Fc5ed48c0F64a55fCe1"), address)

	bytes_compressed, _ := hex.DecodeString("0229f11138ff637ecef0d1878fb5aff4175e96c0758f2f32c004c8e9791e8750ab")
	address, err = builder.GetAddressFromPublicKey(bytes_compressed)
	require.NoError(err)
	require.Equal(xc.Address("0xCc10cd3f77d370F7893E94e4eEb48Fb9553B7a5B"), address)
}

func (s *CrosschainTestSuite) TestGetAddressFromPublicKeyErr() {
	require := s.Require()
	builder, _ := NewAddressBuilder(&xc.AssetConfig{})

	address, err := builder.GetAddressFromPublicKey([]byte{})
	require.Equal(xc.Address(""), address)
	require.EqualError(err, "invalid secp256k1 public key")

	address, err = builder.GetAddressFromPublicKey([]byte{1, 2, 3})
	require.Equal(xc.Address(""), address)
	require.EqualError(err, "invalid secp256k1 public key")
}

func (s *CrosschainTestSuite) TestGetAllPossibleAddressesFromPublicKey() {
	require := s.Require()
	builder, _ := NewAddressBuilder(&xc.AssetConfig{})
	bytes, _ := hex.DecodeString("04760c4460e5336ac9bbd87952a3c7ec4363fc0a97bd31c86430806e287b437fd1b01abc6e1db640cf3106b520344af1d58b00b57823db3e1407cbc433e1b6d04d")
	addresses, err := builder.GetAllPossibleAddressesFromPublicKey(bytes)
	require.Nil(err)
	require.Equal(1, len(addresses))
	require.Equal(xc.Address("0x5891906fEf64A5ae924C7Fc5ed48c0F64a55fCe1"), addresses[0].Address)
	require.Equal(xc.AddressTypeDefault, addresses[0].Type)
}

func (s *CrosschainTestSuite) TestHexToAddress() {
	require := s.Require()
	address, err := HexToAddress("0x5891906fEf64A5ae924C7Fc5ed48c0F64a55fCe1")
	require.Nil(err)
	require.Equal(common.Address{0x58, 0x91, 0x90, 0x6f, 0xEf, 0x64, 0xA5, 0xae, 0x92, 0x4C, 0x7F, 0xc5, 0xed, 0x48, 0xc0, 0xF6, 0x4a, 0x55, 0xfC, 0xe1}, address)

	// common.HexToAddress adds a 0 if the size is not even
	address, err = HexToAddress("0x891906fEf64A5ae924C7Fc5ed48c0F64a55fCe1")
	require.Nil(err)
	require.Equal(common.Address{0x8, 0x91, 0x90, 0x6f, 0xEf, 0x64, 0xA5, 0xae, 0x92, 0x4C, 0x7F, 0xc5, 0xed, 0x48, 0xc0, 0xF6, 0x4a, 0x55, 0xfC, 0xe1}, address)

	// xdc instead of 0x
	address, err = HexToAddress("xdc5891906fEf64A5ae924C7Fc5ed48c0F64a55fCe1")
	require.Nil(err)
	require.Equal(common.Address{0x58, 0x91, 0x90, 0x6f, 0xEf, 0x64, 0xA5, 0xae, 0x92, 0x4C, 0x7F, 0xc5, 0xed, 0x48, 0xc0, 0xF6, 0x4a, 0x55, 0xfC, 0xe1}, address)

	// this should probably never happen in practise, but just to test the implementation
	address, err = HexToAddress("0xxdc5891906fEf64A5ae924C7Fc5ed48c0F64a55fCe1")
	require.Nil(err)
	require.Equal(common.Address{0x58, 0x91, 0x90, 0x6f, 0xEf, 0x64, 0xA5, 0xae, 0x92, 0x4C, 0x7F, 0xc5, 0xed, 0x48, 0xc0, 0xF6, 0x4a, 0x55, 0xfC, 0xe1}, address)
}
