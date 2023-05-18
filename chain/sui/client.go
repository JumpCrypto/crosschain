package sui

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/coming-chat/go-sui/client"
	"github.com/coming-chat/go-sui/types"
	xc "github.com/jumpcrypto/crosschain"
)

// Client for Sui
type Client struct {
	Asset           xc.ITask
	SuiClient       *client.Client
	EstimateGasFunc xc.EstimateGasFunc
}

// NewClient returns a new Aptos Client
func NewClient(cfgI xc.ITask) (*Client, error) {
	cfg := cfgI.GetNativeAsset()
	client, err := client.Dial(cfg.URL)
	return &Client{
		Asset:     cfgI,
		SuiClient: client,
	}, err
}

var _ xc.FullClientWithGas = &Client{}

type SuiMethod string

var (
	// getTransactionBlock SuiMethod = "sui_getTransactionBlock"
	getCheckpoint  SuiMethod = "sui_getCheckpoint"
	getCheckpoints SuiMethod = "sui_getCheckpoints"
	MaxCoinObjects int       = 50
)

func (m SuiMethod) String() string {
	return string(m)
}

type Checkpoint struct {
	Epoch                    string `json:"epoch"`
	SequenceNumber           string `json:"sequenceNumber"`
	Digest                   string `json:"digest"`
	NetworkTotalTransactions string `json:"networkTotalTransactions"`
	PreviousDigest           string `json:"PreviousDigest"`
	TimestampMs              string `json:"timestampMs"`
}

func (ch *Checkpoint) GetEpoch() uint64 {
	return xc.NewAmountBlockchainFromStr(ch.Epoch).Uint64()
}
func (ch *Checkpoint) GetSequenceNumber() uint64 {
	return xc.NewAmountBlockchainFromStr(ch.SequenceNumber).Uint64()
}

type Checkpoints struct {
	Data []*Checkpoint `json:"data"`
}

func (c *Client) FetchLatestCheckpoint(ctx context.Context) (*Checkpoint, error) {
	resp := &Checkpoints{}
	// get last 1 checkpoint, descending order
	err := c.SuiClient.CallContext(ctx, resp, getCheckpoints, nil, 1, true)
	if len(resp.Data) == 0 {
		return &Checkpoint{}, errors.New("no checkpoints yet")
	}
	return resp.Data[0], err
}

func (c *Client) FetchCheckpoint(ctx context.Context, checkpoint uint64) (*Checkpoint, error) {
	resp := &Checkpoint{}
	// get last 1 checkpoint, descending order
	err := c.SuiClient.CallContext(ctx, resp, getCheckpoint, fmt.Sprintf("%d", checkpoint))
	return resp, err
}

func (c *Client) FetchTxInfo(ctx context.Context, txHash xc.TxHash) (xc.TxInfo, error) {
	opts := types.SuiTransactionBlockResponseOptions{
		ShowInput:          true,
		ShowEffects:        true,
		ShowObjectChanges:  true,
		ShowBalanceChanges: true,
		// do we need events?
		ShowEvents: true,
	}
	resp, err := c.SuiClient.GetTransactionBlock(ctx, string(txHash), opts)
	if err != nil {
		return xc.TxInfo{}, err
	}

	// get latest checkpoint so we can compute our confirmations
	latestCheckpoint, err := c.FetchLatestCheckpoint(ctx)
	if err != nil {
		return xc.TxInfo{}, err
	}
	if resp.Checkpoint == nil {
		return xc.TxInfo{}, errors.New("sui endpoint failed to provide checkpoint")
	}
	txCheckpoint, err := c.FetchCheckpoint(ctx, resp.Checkpoint.Uint64())
	if err != nil {
		return xc.TxInfo{}, err
	}
	// latestCheckpoint.Epoch
	sources := []*xc.TxInfoEndpoint{}
	destinations := []*xc.TxInfoEndpoint{}

	from := ""
	to := ""
	contract := ""
	destinationAmount := xc.NewAmountBlockchainFromUint64(0)
	totalSent := xc.NewAmountBlockchainFromUint64(0)
	totalReceived := xc.NewAmountBlockchainFromUint64(0)

	for _, bal := range resp.BalanceChanges {
		amt := xc.NewAmountBlockchainFromStr(bal.Amount)
		asset := ""
		contract = bal.CoinType
		if strings.HasSuffix(bal.CoinType, "sui::SUI") && (strings.HasPrefix(bal.CoinType, "0x0000000000000000000000000000000000000000000000000000000000") || len(contract) < 16) {
			contract = ""
			asset = "SUI"
		}
		if amt.Sign() < 0 {
			from = bal.Owner.AddressOwner.String()
			abs := amt.Abs()
			totalSent = totalSent.Add(&abs)
			sources = append(sources, &xc.TxInfoEndpoint{
				Asset:           xc.Asset(asset),
				ContractAddress: xc.ContractAddress(contract),
				Amount:          abs,
				Address:         xc.Address(from),
				NativeAsset:     c.Asset.GetNativeAsset().NativeAsset,
			})
		} else {
			to = bal.Owner.AddressOwner.String()
			destinationAmount = amt
			totalReceived = totalReceived.Add(&amt)
			destinations = append(destinations, &xc.TxInfoEndpoint{
				Asset:           xc.Asset(asset),
				ContractAddress: xc.ContractAddress(contract),
				Amount:          amt,
				Address:         xc.Address(to),
				NativeAsset:     c.Asset.GetNativeAsset().NativeAsset,
			})
		}
	}

	// fee is difference between total sent and recieved in balance changes
	fee := totalSent.Sub(&totalReceived)

	status := xc.TxStatusSuccess
	if resp.Effects.Data.V1.Status.Error != "" {
		status = xc.TxStatusFailure
	}

	return xc.TxInfo{
		BlockHash:       txCheckpoint.Digest,
		TxID:            resp.Digest,
		From:            xc.Address(from),
		To:              xc.Address(to),
		ContractAddress: xc.ContractAddress(contract),
		Amount:          destinationAmount,
		Fee:             fee,
		BlockTime:       resp.TimestampMs.Int64(),
		BlockIndex:      resp.Checkpoint.Int64(),
		Confirmations:   int64(latestCheckpoint.GetSequenceNumber()) - int64(txCheckpoint.GetSequenceNumber()),

		ExplorerURL:  fmt.Sprintf("https://explorer.sui.io/txblock/%s?network=%s", resp.Digest, c.Asset.GetAssetConfig().Net),
		Sources:      sources,
		Destinations: destinations,
		Error:        resp.Effects.Data.V1.Status.Error,
		Status:       status,
	}, nil
}

func (c *Client) RegisterEstimateGasCallback(estimateGas xc.EstimateGasFunc) {
	c.EstimateGasFunc = estimateGas
}

func (c *Client) EstimateGas(ctx context.Context) (xc.AmountBlockchain, error) {
	if c.EstimateGasFunc != nil {
		nativeAsset := c.Asset.GetNativeAsset().NativeAsset
		res, err := c.EstimateGasFunc(nativeAsset)
		if err != nil {
			// continue with default implementation as fallback
		} else {
			return res, err
		}
	}

	ref, err := c.SuiClient.GetReferenceGasPrice(ctx)
	if err != nil {
		return xc.NewAmountBlockchainFromUint64(0), err
	}
	return xc.NewAmountBlockchainFromUint64(ref.Uint64()), nil
}

func (c *Client) GetAllCoinsFor(ctx context.Context, address xc.Address, contract string) ([]*types.Coin, error) {
	fromData, err := types.NewHexData(string(address))
	if err != nil {
		return []*types.Coin{}, err
	}

	all_coins := []*types.Coin{}

	var next *types.HexData
	for {
		coins, err := c.SuiClient.GetCoins(ctx, *fromData, &contract, next, 250)
		if err != nil {
			return []*types.Coin{}, err
		}
		for _, coin := range coins.Data {
			c := coin
			all_coins = append(all_coins, &c)
		}
		next = coins.NextCursor
		if next == nil || !coins.HasNextPage {
			break
		}
	}
	return all_coins, nil

}

func (c *Client) FetchTxInput(ctx context.Context, from xc.Address, to xc.Address) (xc.TxInput, error) {
	// native asset SUI
	native := "0x2::sui::SUI"
	contract := native
	if token, ok := c.Asset.(*xc.TokenAssetConfig); ok {
		contract = NormalizeCoinContract(token.Contract)
	}

	all_coins, err := c.GetAllCoinsFor(ctx, from, contract)
	if err != nil {
		return TxInput{}, err
	}

	latestCheckpoint, err := c.FetchLatestCheckpoint(ctx)
	if err != nil {
		return xc.TxInfo{}, err
	}
	epoch := xc.NewAmountBlockchainFromStr(latestCheckpoint.Epoch)

	// store the object id's for the transfer
	input := NewTxInput()
	input.CurrentEpoch = epoch.Uint64()
	input.Coins = all_coins
	input.SortCoins()
	// take max 50 to bound the tx_input size.
	if len(input.Coins) > MaxCoinObjects {
		input.Coins = input.Coins[:MaxCoinObjects]
	}

	if contract == native {
		// gas coin should just be the largest object
		if len(input.Coins) > 0 {
			input.GasCoin = *input.Coins[0]
		}
	} else {
		// we need to fetch our sui objects
		all_sui_coins, err := c.GetAllCoinsFor(ctx, from, native)
		if err != nil {
			return TxInput{}, err
		}
		SortCoins(all_sui_coins)
		if len(all_sui_coins) > 0 {
			input.GasCoin = *all_sui_coins[0]
		}
	}

	gasPrice, err := c.EstimateGas(ctx)
	if err != nil {
		defaultgas := c.Asset.GetNativeAsset().ChainGasPriceDefault
		if defaultgas < 0.1 {
			return input, err
		}
		// use the default
		input.GasPrice = uint64(defaultgas)
	}
	input.GasPrice = gasPrice.Uint64()
	// 2 SUI
	input.GasBudget = 2_000_000_000
	input.ExcludeGasCoin()

	return input, nil
}

// SubmitTx submits a Sui tx
func (c *Client) SubmitTx(ctx context.Context, tx xc.Tx) error {
	tx_bz, err := tx.Serialize()
	if err != nil {
		return err
	}
	var sigs [][]byte
	sigsB64 := []any{}
	if getter, ok := tx.(SignatureGetter); ok {
		sigs = getter.GetSignatures()
	} else {
		return errors.New("cannot get signatures to submit sui transaction, must implement GetSignatures()")
	}
	for _, sig := range sigs {
		sigsB64 = append(sigsB64, types.Base64Data(sig))
	}

	newTxn, err := c.SuiClient.ExecuteTransactionBlock(
		ctx,
		types.Base64Data(tx_bz),
		sigsB64,
		&types.SuiTransactionBlockResponseOptions{},
		types.TxnRequestTypeWaitForLocalExecution,
	)
	_ = newTxn
	return err
}

func (c *Client) FetchBalanceFor(ctx context.Context, address xc.Address, contract string) (xc.AmountBlockchain, error) {
	total := xc.NewAmountBlockchainFromUint64(0)
	contract = NormalizeCoinContract(contract)
	all_coins, err := c.GetAllCoinsFor(ctx, address, contract)
	if err != nil {
		return total, err
	}

	for _, coin := range all_coins {
		amt := xc.NewAmountBlockchainFromUint64(coin.Balance.Uint64())
		total = total.Add(&amt)
	}

	return total, nil
}
func (c *Client) FetchBalance(ctx context.Context, address xc.Address) (xc.AmountBlockchain, error) {
	// native asset SUI
	contract := "0x2::sui::SUI"
	if token, ok := c.Asset.(*xc.TokenAssetConfig); ok {
		contract = token.Contract
	}
	return c.FetchBalanceFor(ctx, address, contract)
}

func (c *Client) FetchNativeBalance(ctx context.Context, address xc.Address) (xc.AmountBlockchain, error) {
	return c.FetchBalanceFor(ctx, address, "0x2::sui::SUI")
}
