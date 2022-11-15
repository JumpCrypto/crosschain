package evm

import (
	"context"
	"errors"

	xc "github.com/jumpcrypto/crosschain"
)

// Client for EVM
type Client struct {
}

// TxInput for EVM
type TxInput struct {
	xc.TxInput
}

// NewClient returns a new EVM Client
func NewClient(cfg xc.AssetConfig) (*Client, error) {
	return &Client{}, errors.New("not implemented")
}

// NewLegacyClient returns a new EVM Client for legacy tx
func NewLegacyClient(cfg xc.AssetConfig) (*Client, error) {
	return &Client{}, errors.New("not implemented")
}

// FetchTxInput returns tx input for a EVM tx
func (client *Client) FetchTxInput(ctx context.Context, from xc.Address, _ xc.Address) (xc.TxInput, error) {
	return TxInput{}, errors.New("not implemented")
}

// SubmitTx submits a EVM tx
func (client *Client) SubmitTx(ctx context.Context, txInput xc.Tx) error {
	return errors.New("not implemented")
}

// FetchTxInfo returns tx info for a EVM tx
func (client *Client) FetchTxInfo(ctx context.Context, txHash xc.TxHash) (xc.TxInfo, error) {
	return xc.TxInfo{}, errors.New("not implemented")
}
