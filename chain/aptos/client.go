package aptos

import (
	"context"
	"errors"
	"fmt"

	"github.com/coming-chat/go-aptos/aptosclient"
	"github.com/coming-chat/go-aptos/aptostypes"
	xc "github.com/jumpcrypto/crosschain"
)

// Client for Aptos
type Client struct {
	Asset           xc.ITask
	AptosClient     *aptosclient.RestClient
	EstimateGasFunc xc.EstimateGasFunc
}

var _ xc.FullClientWithGas = &Client{}

// NewClient returns a new Aptos Client
func NewClient(cfgI xc.ITask) (*Client, error) {
	cfg := cfgI.GetNativeAsset()
	client, err := aptosclient.Dial(context.Background(), cfg.URL)
	return &Client{
		Asset:       cfgI,
		AptosClient: client,
	}, err
}

// FetchTxInput returns tx input for a Aptos tx
func (client *Client) FetchTxInput(ctx context.Context, from xc.Address, _ xc.Address) (xc.TxInput, error) {
	ledger, err := client.AptosClient.LedgerInfo()
	if err != nil {
		return &TxInput{}, err
	}
	acc, err := client.AptosClient.GetAccount(string(from))
	if err != nil {
		return &TxInput{}, err
	}
	gas_price, err := client.EstimateGas(ctx)
	if err != nil {
		return &TxInput{}, err
	}

	return &TxInput{
		TxInputEnvelope: xc.TxInputEnvelope{
			Type: xc.DriverAptos,
		},
		SequenceNumber: acc.SequenceNumber,
		ChainId:        ledger.ChainId,
		GasLimit:       2000,
		Timestamp:      ledger.LedgerTimestamp,
		GasPrice:       gas_price.Uint64(),
	}, nil
}

// SubmitTx submits a Aptos tx
func (client *Client) SubmitTx(ctx context.Context, tx xc.Tx) error {
	tx_bz, err := tx.Serialize()
	if err != nil {
		return err
	}
	newTxn, err := client.AptosClient.SubmitSignedBCSTransaction(tx_bz)
	_ = newTxn
	return err
}

// FetchTxInfo returns tx info for a Aptos tx
func (client *Client) FetchTxInfo(ctx context.Context, txHash xc.TxHash) (xc.TxInfo, error) {

	tx, err := client.AptosClient.GetTransactionByHash(string(txHash))
	if err != nil {
		return xc.TxInfo{}, err
	}
	block, err := client.AptosClient.GetBlockByVersion(fmt.Sprintf("%d", tx.Version), false)
	if err != nil {
		return xc.TxInfo{}, err
	}
	ledger, err := client.AptosClient.LedgerInfo()
	if err != nil {
		return xc.TxInfo{}, err
	}

	tx_height := block.BlockHeight
	now_height := ledger.BlockHeight
	confirmations := now_height - tx_height

	unit_price := tx.GasUnitPrice
	gas_used := tx.GasUsed
	feeu256 := xc.NewAmountBlockchainFromUint64(gas_used * unit_price)

	return xc.TxInfo{
		To:            toFromTxPayload(tx.Payload),
		From:          xc.Address(tx.Sender),
		Amount:        valueFromTxPayload(tx.Payload),
		Fee:           feeu256,
		Confirmations: int64(confirmations),
		BlockHash:     fmt.Sprintf("%d", tx.Version),
		// convert usec to sec
		BlockTime:   int64((tx.Timestamp / 1000) / 1000),
		TxID:        tx.Hash,
		BlockIndex:  int64(tx.Version),
		ExplorerURL: fmt.Sprintf("/txn/%d?network=%s", tx.Version, client.Asset.GetNativeAsset().Net),
	}, nil
}

// FetchBalance fetches balance for an Aptos address
func (client *Client) FetchBalance(ctx context.Context, address xc.Address) (xc.AmountBlockchain, error) {
	if token, ok := client.Asset.(*xc.TokenAssetConfig); ok {
		balance, err := client.AptosClient.BalanceOf(string(address), token.Contract)
		if err != nil {
			return xc.NewAmountBlockchainFromUint64(0), err
		}
		return xc.AmountBlockchain(*balance), err
	}
	return client.FetchNativeBalance(ctx, address)
}

// FetchNativeBalance fetches the native asset balance for an Aptos address
func (client *Client) FetchNativeBalance(ctx context.Context, address xc.Address) (xc.AmountBlockchain, error) {
	balance, err := client.AptosClient.AptosBalanceOf(string(address))
	if err != nil {
		return xc.NewAmountBlockchainFromUint64(0), err
	}
	return xc.AmountBlockchain(*balance), nil
}

func (client *Client) RegisterEstimateGasCallback(estimateGas xc.EstimateGasFunc) {
	client.EstimateGasFunc = estimateGas
}

func (client *Client) EstimateGas(ctx context.Context) (xc.AmountBlockchain, error) {
	// invoke EstimateGasFunc callback, if registered
	if client.EstimateGasFunc != nil {
		nativeAsset := client.Asset.GetNativeAsset().NativeAsset
		res, err := client.EstimateGasFunc(nativeAsset)
		if err != nil {
			// continue with default implementation as fallback
		} else {
			return res, err
		}
	}
	// estimate using last 1 blocks
	zero := xc.NewAmountBlockchainFromUint64(0)
	res, err := client.AptosClient.LedgerInfo()
	if err != nil {
		return zero, err
	}
	height := res.BlockHeight
	if height < 500 {
		return zero, errors.New("the chain is too young")
	}
	attempts := 10

	// let's download the last 50 transactions
	transactions := []aptostypes.Transaction{}
	for len(transactions) < 50 && height > 0 {
		block, err := client.AptosClient.GetBlockByHeight(fmt.Sprintf("%d", height), true)
		height = height - 1
		if err != nil {
			// Sometimes a block doesn't exist..
			// so we'll tolerate up to 10 times of this in a row.
			attempts = attempts - 1
			if attempts <= 0 {
				return zero, err
			}
			continue
		}
		l1 := len(transactions)
		for _, tx := range block.Transactions {
			if tx.GasUnitPrice != 0 {
				transactions = append(transactions, tx)
			}
		}
		l2 := len(transactions)
		if l1 == l2 {
			// if the block was empty, count as a failed attempt so we will terminate
			attempts = attempts - 1
			if attempts <= 0 {
				break
			}
			continue
		}
		attempts = 10
	}
	totalUnitPrice := uint64(0)
	for _, tx := range transactions {
		totalUnitPrice += tx.GasUnitPrice
	}

	// use default of 0.000001 fee per gas
	if totalUnitPrice == 0 {
		return xc.NewAmountBlockchainFromUint64(100), nil
	}

	averUnitPrice := float32(totalUnitPrice) / float32(len(transactions))
	// pay 25% premium
	averUnitPrice = averUnitPrice * 1.25
	// truncate
	unit_price := xc.NewAmountBlockchainFromUint64(uint64(averUnitPrice))

	return unit_price, nil
}
