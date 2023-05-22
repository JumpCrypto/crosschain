package evm

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	xc "github.com/jumpcrypto/crosschain"
	"github.com/jumpcrypto/crosschain/chain/evm/erc20"
)

const DEFAULT_GAS_PRICE = 20_000_000_000
const DEFAULT_GAS_TIP = 3_000_000_000

var ERC20 abi.ABI

func init() {
	var err error
	ERC20, err = abi.JSON(strings.NewReader(erc20.Erc20ABI))
	if err != nil {
		panic(err)
	}
}

// Client for EVM
type Client struct {
	Asset           xc.ITask
	EthClient       *ethclient.Client
	RpcClient       *rpc.Client
	ChainId         *big.Int
	Interceptor     *HttpInterceptor
	EstimateGasFunc xc.EstimateGasFunc
	Legacy          bool
}

var _ xc.FullClientWithGas = &Client{}

// TxInput for EVM
type TxInput struct {
	xc.TxInputEnvelope
	Nonce    uint64
	GasLimit uint64
	// DynamicFeeTx
	GasTipCap xc.AmountBlockchain // maxPriorityFeePerGas
	GasFeeCap xc.AmountBlockchain // maxFeePerGas
	// LegacyTx
	GasPrice xc.AmountBlockchain // wei per gas
	// Task params
	Params []string
}

func NewTxInput() *TxInput {
	return &TxInput{
		TxInputEnvelope: xc.TxInputEnvelope{
			Type: xc.DriverEVM,
		},
	}
}

// Interceptor
type HttpInterceptor struct {
	core    http.RoundTripper
	enabled bool
}

func (i *HttpInterceptor) Enable() {
	i.enabled = true
}
func (i *HttpInterceptor) Disable() {
	i.enabled = false
}

func (i HttpInterceptor) RoundTrip(req *http.Request) (*http.Response, error) {
	defer func() {
		_ = req.Body.Close()
	}()

	// log.Printf("[EVM] req: %+v", req)
	res, err := i.core.RoundTrip(req)
	if err != nil {
		return nil, err
	}
	if i.enabled {
		defer func() {
			_ = res.Body.Close()
		}()
		body, _ := ioutil.ReadAll(res.Body)
		bodyStr := string(body)

		newStr := ""
		// KLAY issue
		if strings.Contains(bodyStr, "type\":\"TxTypeLegacyTransaction") {
			log.Print("Replacing KLAY TxTypeLegacyTransaction")
			newStr = strings.Replace(bodyStr, "TxTypeLegacyTransaction", "0x0", 1)
			newStr = strings.Replace(newStr, "\"V\"", "\"v\"", 1)
			newStr = strings.Replace(newStr, "\"R\"", "\"r\"", 1)
			newStr = strings.Replace(newStr, "\"S\"", "\"s\"", 1)
			newStr = strings.Replace(newStr, "\"signatures\":[{", "", 1)
			newStr = strings.Replace(newStr, "}]", ",\"cumulativeGasUsed\":\"0x0\"", 1)
		}
		if strings.Contains(bodyStr, "parentHash") {
			log.Print("Adding KLAY/CELO sha3Uncles")
			newStr = strings.Replace(bodyStr, "parentHash", "gasLimit\":\"0x0\",\"difficulty\":\"0x0\",\"miner\":\"0x0000000000000000000000000000000000000000\",\"sha3Uncles\":\"0x0000000000000000000000000000000000000000000000000000000000000000\",\"parentHash", 1)
		}
		if newStr != "" {
			newBody := []byte(newStr)
			res.Body = io.NopCloser(bytes.NewReader(newBody))
			res.ContentLength = int64(len(newBody))
			res.Header.Set("Content-Length", fmt.Sprintf("%d", res.ContentLength))
		}
	}
	newRes := res

	return newRes, nil
}

func configToEVMClientURL(cfgI xc.ITask) string {
	cfg := cfgI.GetNativeAsset()
	if cfg.Provider == "infura" {
		return cfg.URL + "/" + cfg.AuthSecret
	}
	return cfg.URL
}

// NewClient returns a new EVM Client
func NewClient(asset xc.ITask) (*Client, error) {
	nativeAsset := asset.GetNativeAsset()
	url := configToEVMClientURL(asset)

	// c, err := rpc.DialContext(context.Background(), url)
	interceptor := &HttpInterceptor{http.DefaultTransport, false}
	httpClient := &http.Client{
		Transport: interceptor,
	}
	c, err := rpc.DialHTTPWithClient(url, httpClient)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("dialing url: %v", nativeAsset.URL))
	}

	client := ethclient.NewClient(c)
	return &Client{
		Asset:           asset,
		EthClient:       client,
		RpcClient:       c,
		ChainId:         nil,
		Interceptor:     interceptor,
		EstimateGasFunc: nil,
		Legacy:          false,
	}, nil
}

// ChainID returns the ChainID
func (client *Client) ChainID() (*big.Int, error) {
	var err error
	if client.ChainId == nil {
		client.ChainId, err = client.EthClient.ChainID(context.Background())
	}
	return client.ChainId, err
}

// NewLegacyClient returns a new EVM Client for legacy tx
func NewLegacyClient(cfg xc.ITask) (*Client, error) {
	client, err := NewClient(cfg)
	if err != nil {
		return nil, err
	}
	client.Legacy = true
	return client, nil
}

// FetchTxInput returns tx input for a EVM tx
func (client *Client) FetchTxInput(ctx context.Context, from xc.Address, _ xc.Address) (xc.TxInput, error) {
	nativeAsset := client.Asset.GetNativeAsset()
	task := client.Asset.GetTask()

	zero := xc.NewAmountBlockchainFromUint64(0)
	result := NewTxInput()
	result.GasTipCap = xc.NewAmountBlockchainFromUint64(DEFAULT_GAS_TIP)
	result.GasFeeCap = zero
	result.GasPrice = zero

	// Nonce
	var targetAddr common.Address
	var err error
	if task != nil && task.Signer != "" {
		targetAddr, err = HexToAddress(xc.Address(task.Signer))
	} else {
		targetAddr, err = HexToAddress(from)
	}
	if err != nil {
		return zero, fmt.Errorf("bad to address '%v': %v", from, err)
	}
	nonce, err := client.EthClient.NonceAt(ctx, targetAddr, nil)
	if err != nil {
		return result, err
	}
	result.Nonce = nonce

	// Gas
	if !nativeAsset.NoGasFees {
		gas, err := client.EstimateGas(ctx)
		if err != nil {
			// pass, return err later
		}
		result.GasPrice = gas  // legacy
		result.GasFeeCap = gas // new
	} else {
		result.GasTipCap = zero
	}

	return result, err
}

// SubmitTx submits a EVM tx
func (client *Client) SubmitTx(ctx context.Context, tx xc.Tx) error {
	switch tx := tx.(type) {
	case *Tx:
		err := client.EthClient.SendTransaction(ctx, tx.EthTx)
		if err != nil {
			return fmt.Errorf(fmt.Sprintf("sending transaction '%v': %v", tx.Hash(), err))
		}
		return nil
	default:
		bz, err := tx.Serialize()
		if err != nil {
			return err
		}
		return client.RpcClient.CallContext(ctx, nil, "eth_sendRawTransaction", hexutil.Encode(bz))
	}
}

// FetchTxInfo returns tx info for a EVM tx
func (client *Client) FetchTxInfo(ctx context.Context, txHashStr xc.TxHash) (xc.TxInfo, error) {
	nativeAsset := client.Asset.GetNativeAsset()
	txHashHex := TrimPrefixes(string(txHashStr))
	txHash := common.HexToHash(txHashHex)

	result := xc.TxInfo{
		TxID:        txHashHex,
		ExplorerURL: nativeAsset.ExplorerURL + "/tx/0x" + txHashHex,
	}

	tx, pending, err := client.EthClient.TransactionByHash(ctx, txHash)
	if err != nil {
		// TODO retry only for KLAY
		client.Interceptor.Enable()
		tx, pending, err = client.EthClient.TransactionByHash(ctx, txHash)
		client.Interceptor.Disable()
		if err != nil {
			return result, fmt.Errorf(fmt.Sprintf("fetching tx by hash '%s': %v", txHashStr, err))
		}
	}

	chainID := new(big.Int).SetInt64(nativeAsset.ChainID)
	// chainID, err := client.EthClient.ChainID(ctx)
	// if err != nil {
	// 	return result, fmt.Errorf("fetching chain ID: %v", err)
	// }

	// If the transaction is still pending, return an empty txInfo.
	if pending {
		return result, nil
	}

	receipt, err := client.EthClient.TransactionReceipt(ctx, txHash)
	if err != nil {
		// TODO retry only for KLAY
		client.Interceptor.Enable()
		receipt, err = client.EthClient.TransactionReceipt(ctx, txHash)
		client.Interceptor.Disable()
		if err != nil {
			return result, fmt.Errorf("fetching receipt for tx %v : %v", txHashStr, err)
		}
	}

	// if no receipt, tx has 0 confirmations
	if receipt == nil {
		return result, nil
	}

	// reverted tx
	result.BlockIndex = receipt.BlockNumber.Int64()
	result.BlockHash = receipt.BlockHash.Hex()
	gasUsed := receipt.GasUsed
	if receipt.Status == 0 {
		result.Status = xc.TxStatusFailure
	}

	// tx confirmed
	currentHeader, err := client.EthClient.HeaderByNumber(ctx, receipt.BlockNumber)
	if err != nil {
		client.Interceptor.Enable()
		currentHeader, err = client.EthClient.HeaderByNumber(ctx, receipt.BlockNumber)
		client.Interceptor.Disable()
		if err != nil {
			return result, fmt.Errorf("fetching current header: (%T) %v", err, err)
		}
	}
	result.BlockTime = int64(currentHeader.Time)
	var baseFee uint64
	if currentHeader.BaseFee != nil {
		baseFee = currentHeader.BaseFee.Uint64()
	}

	latestHeader, err := client.EthClient.HeaderByNumber(ctx, nil)
	if err != nil {
		client.Interceptor.Enable()
		latestHeader, err = client.EthClient.HeaderByNumber(ctx, nil)
		client.Interceptor.Disable()
		if err != nil {
			return result, fmt.Errorf("fetching latest header: %v", err)
		}
	}
	result.Confirmations = latestHeader.Number.Int64() - receipt.BlockNumber.Int64()

	// // tx confirmed
	confirmedTx := Tx{
		EthTx:  tx,
		Signer: types.LatestSignerForChainID(chainID),
	}

	info := confirmedTx.ParseTransfer(receipt, nativeAsset.NativeAsset)

	result.From = confirmedTx.From()
	result.To = confirmedTx.To()
	result.ContractAddress = confirmedTx.ContractAddress()
	result.Amount = confirmedTx.Amount()
	result.Fee = confirmedTx.Fee(baseFee, gasUsed)
	result.Sources = info.Sources
	result.Destinations = info.Destinations

	return result, nil
}

// EstimateGas estimates gas price for a Cosmos chain
func (client *Client) EstimateGas(ctx context.Context) (xc.AmountBlockchain, error) {
	asset := client.Asset.GetNativeAsset()

	// invoke EstimateGasFunc callback, if registered
	if client.EstimateGasFunc != nil {
		nativeAsset := asset.NativeAsset
		res, err := client.EstimateGasFunc(nativeAsset)
		if err != nil {
			// continue with default implementation as fallback
		} else {
			return res, err
		}
	}

	// KLAY has fixed gas price of 250 ston
	if asset.NativeAsset == xc.KLAY {
		return xc.NewAmountBlockchainFromUint64(250_000_000_000), nil
	}

	// legacy gas estimation via SuggestGasPrice
	var baseFee uint64
	if client.Legacy {
		baseFeeInt, err := client.EthClient.SuggestGasPrice(ctx)
		if err != nil {
			// pass
		} else {
			baseFee = baseFeeInt.Uint64()
		}
	} else {
		latest, err := client.EthClient.HeaderByNumber(ctx, nil)
		if err != nil {
			// pass
		} else {
			baseFee = latest.BaseFee.Uint64()
		}
	}

	if baseFee < DEFAULT_GAS_PRICE {
		baseFee = DEFAULT_GAS_PRICE
	}
	multiplier := 2.0
	if asset.ChainGasMultiplier > 0.0 {
		multiplier = asset.ChainGasMultiplier
	}

	gasPrice := (uint64)((float64)(baseFee+DEFAULT_GAS_TIP) * multiplier)
	return xc.NewAmountBlockchainFromUint64(gasPrice), nil
}

// RegisterEstimateGasCallback registers a callback to get gas price
func (client *Client) RegisterEstimateGasCallback(fn xc.EstimateGasFunc) {
	client.EstimateGasFunc = fn
}

// Fetch the balance of the native asset that this client is configured for
func (client *Client) FetchNativeBalance(ctx context.Context, address xc.Address) (xc.AmountBlockchain, error) {
	zero := xc.NewAmountBlockchainFromUint64(0)
	targetAddr, err := HexToAddress(address)
	if err != nil {
		return zero, fmt.Errorf("bad to address '%v': %v", address, err)
	}
	balance, err := client.EthClient.BalanceAt(ctx, targetAddr, nil)
	if err != nil {
		return zero, fmt.Errorf("failed to get balance for '%v': %v", address, err)
	}

	return xc.AmountBlockchain(*balance), nil
}

// Fetch the balance of the asset that this client is configured for
func (client *Client) FetchBalance(ctx context.Context, address xc.Address) (xc.AmountBlockchain, error) {
	// native
	if _, ok := client.Asset.(*xc.NativeAssetConfig); ok {
		return client.FetchNativeBalance(ctx, address)
	}

	// token
	asset := client.Asset.GetAssetConfig()
	zero := xc.NewAmountBlockchainFromUint64(0)
	tokenAddress, _ := HexToAddress(xc.Address(asset.Contract))
	instance, err := erc20.NewErc20(tokenAddress, client.EthClient)
	if err != nil {
		return zero, err
	}

	dstAddress, _ := HexToAddress(address)
	balance, err := instance.BalanceOf(&bind.CallOpts{}, dstAddress)
	if err != nil {
		return zero, err
	}
	return xc.AmountBlockchain(*balance), nil
}
