package aptos

import (
	"context"
	"errors"

	"github.com/coming-chat/go-aptos/aptosclient"
	xc "github.com/jumpcrypto/crosschain"
)

// TxInput for Aptos
type TxInput struct {
	xc.TxInputEnvelope
}

// NewTxInput returns a new Aptos TxInput
func NewTxInput() *TxInput {
	return &TxInput{
		TxInputEnvelope: *xc.NewTxInputEnvelope(xc.DriverAptos),
	}
}

// Client for Aptos
type Client struct {
	Asset       *xc.AssetConfig
	AptosClient *aptosclient.RestClient
}

var _ xc.FullClient = &Client{}

// NewClient returns a new Aptos Client
func NewClient(cfgI xc.ITask) (*Client, error) {
	asset := cfgI.GetAssetConfig()
	cfg := cfgI.GetNativeAsset()
	client, err := aptosclient.Dial(context.Background(), cfg.URL)
	return &Client{
		Asset:       asset,
		AptosClient: client,
	}, err
}

// FetchTxInput returns tx input for a Aptos tx
func (client *Client) FetchTxInput(ctx context.Context, from xc.Address, _ xc.Address) (xc.TxInput, error) {
	return &TxInput{}, errors.New("not implemented")
}

// SubmitTx submits a Aptos tx
func (client *Client) SubmitTx(ctx context.Context, txInput xc.Tx) error {
	return errors.New("not implemented")
}

// FetchTxInfo returns tx info for a Aptos tx
func (client *Client) FetchTxInfo(ctx context.Context, txHash xc.TxHash) (xc.TxInfo, error) {
	return xc.TxInfo{}, errors.New("not implemented")
}

// FetchBalance fetches balance for an Aptos address
func (client *Client) FetchBalance(ctx context.Context, address xc.Address) (xc.AmountBlockchain, error) {
	if client.Asset.Type == xc.AssetTypeNative {
		return client.FetchNativeBalance(ctx, address)
	}
	balance, err := client.AptosClient.BalanceOf(string(address), client.Asset.Contract)
	if err != nil {
		return xc.NewAmountBlockchainFromUint64(0), err
	}
	return xc.AmountBlockchain(*balance), nil
}

// FetchNativeBalance fetches the native asset balance for an Aptos address
func (client *Client) FetchNativeBalance(ctx context.Context, address xc.Address) (xc.AmountBlockchain, error) {
	balance, err := client.AptosClient.AptosBalanceOf(string(address))
	if err != nil {
		return xc.NewAmountBlockchainFromUint64(0), err
	}
	return xc.AmountBlockchain(*balance), nil
}
