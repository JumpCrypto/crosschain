package newchain

import (
	"errors"

	xc "github.com/jumpcrypto/crosschain"
)

// TxBuilder for Template
type TxBuilder struct {
	Asset xc.AssetConfig
}

// NewTxBuilder creates a new Template TxBuilder
func NewTxBuilder(asset xc.AssetConfig) (xc.TxBuilder, error) {
	return TxBuilder{
		Asset: asset,
	}, errors.New("not implemented")
}

// NewTransfer creates a new transfer for an Asset, either native or token
func (txBuilder TxBuilder) NewTransfer(from xc.Address, to xc.Address, amount xc.AmountBlockchain, input xc.TxInput) (xc.Tx, error) {
	if txBuilder.Asset.Type == xc.AssetTypeToken {
		return txBuilder.NewTokenTransfer(from, to, amount, input)
	}
	return txBuilder.NewNativeTransfer(from, to, amount, input)
}

// NewNativeTransfer creates a new transfer for a native asset
func (txBuilder TxBuilder) NewNativeTransfer(from xc.Address, to xc.Address, amount xc.AmountBlockchain, input xc.TxInput) (xc.Tx, error) {
	return nil, errors.New("not implemented")
}

// NewTokenTransfer creates a new transfer for a token asset
func (txBuilder TxBuilder) NewTokenTransfer(from xc.Address, to xc.Address, amount xc.AmountBlockchain, input xc.TxInput) (xc.Tx, error) {
	return nil, errors.New("not implemented")
}
