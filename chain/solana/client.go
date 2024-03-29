package solana

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	xc "github.com/jumpcrypto/crosschain"

	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/token"
	"github.com/gagliardetto/solana-go/rpc"
)

// TxInput for Solana
type TxInput struct {
	xc.TxInputEnvelope
	RecentBlockHash solana.Hash
	ToIsATA         bool
	ShouldCreateATA bool
}

// NewTxInput returns a new Solana TxInput
func NewTxInput() *TxInput {
	return &TxInput{
		TxInputEnvelope: *xc.NewTxInputEnvelope(xc.DriverSolana),
	}
}

// Client for Solana
type Client struct {
	SolClient *rpc.Client
	Asset     xc.ITask
}

var _ xc.Client = &Client{}

// NewClient returns a new JSON-RPC Client to the Solana node
func NewClient(cfgI xc.ITask) (*Client, error) {
	cfg := cfgI.GetNativeAsset()
	solClient := rpc.New(cfg.URL)
	return &Client{
		SolClient: solClient,
		Asset:     cfgI,
	}, nil
}

// FetchTxInput returns tx input for a Solana tx, namely a RecentBlockHash
func (client *Client) FetchTxInput(ctx context.Context, from xc.Address, to xc.Address) (xc.TxInput, error) {
	txInput := NewTxInput()
	asset := client.Asset

	// get recent block hash (i.e. nonce)
	// GetRecentBlockhash will be deprecated - GetLatestBlockhash already tested, just switch
	// recent, err := client.SolClient.GetLatestBlockhash(ctx, rpc.CommitmentFinalized)
	recent, err := client.SolClient.GetRecentBlockhash(ctx, rpc.CommitmentFinalized)
	if err != nil {
		return nil, err
	}
	if recent == nil || recent.Value == nil {
		return nil, errors.New("error fetching blockhash")
	}
	txInput.RecentBlockHash = recent.Value.Blockhash
	contract := ""
	if token, ok := asset.(*xc.TokenAssetConfig); ok {
		contract = token.Contract
	} else {
		// TODO remove native asset
		if asset.GetAssetConfig().Contract != "" {
			contract = asset.GetAssetConfig().Contract
		} else {
			// native transfer
			return txInput, nil
		}
	}

	// get account info - check if to is an owner or ata
	accountTo, err := solana.PublicKeyFromBase58(string(to))
	if err != nil {
		return nil, err
	}
	res, err := client.SolClient.GetAccountInfo(ctx, accountTo)
	if err != nil {
		txInput.ToIsATA = true
	} else {
		ownerAddr := res.Value.Owner.String()
		// log.Println("owner", ownerAddr)
		sysAddr := "11111111111111111111111111111111"
		if ownerAddr != sysAddr {
			// The field "to" is not an owner address, therefore is a (possibly custom) ATA
			txInput.ToIsATA = true
		}
	}

	// for tokens, get ata account info
	ataTo := accountTo
	if !txInput.ToIsATA {
		ataToStr, err := FindAssociatedTokenAddress(string(to), contract)
		if err != nil {
			return nil, err
		}
		ataTo = solana.MustPublicKeyFromBase58(ataToStr)
	}
	_, err = client.SolClient.GetAccountInfo(ctx, ataTo)
	if err != nil {
		// if the ATA doesn't exist yet, we will create when sending tokens
		txInput.ShouldCreateATA = true
	}

	return txInput, nil
}

func (client *Client) SubmitTx(ctx context.Context, txInput xc.Tx) error {
	txData, err := txInput.Serialize()
	if err != nil {
		return fmt.Errorf("send transaction: encode transaction: %w", err)
	}

	_, err = client.SolClient.SendEncodedTransactionWithOpts(
		ctx,
		base64.StdEncoding.EncodeToString(txData),
		rpc.TransactionOpts{
			SkipPreflight:       false,
			PreflightCommitment: rpc.CommitmentFinalized,
		},
	)
	return err
}

func (client *Client) getParsedAccountInfo(ctx context.Context, ataPubKey string) (*token.Account, error) {
	ata := solana.MustPublicKeyFromBase58(ataPubKey)
	res, err := client.SolClient.GetAccountInfo(ctx, ata)
	if err != nil {
		return nil, err
	}
	account := res.Value
	data := account.Data.GetBinary()
	var tokenAcct token.Account
	if err := bin.NewBinDecoder(data).Decode(&tokenAcct); err != nil {
		return nil, err
	}
	return &tokenAcct, nil
}

// FetchTxInfo returns tx info for a Solana tx
func (client *Client) FetchTxInfo(ctx context.Context, txHash xc.TxHash) (xc.TxInfo, error) {
	result := xc.TxInfo{}

	txSig, err := solana.SignatureFromBase58(string(txHash))
	if err != nil {
		return result, err
	}

	res, err := client.SolClient.GetTransaction(
		ctx,
		txSig,
		&rpc.GetTransactionOpts{
			Encoding:   solana.EncodingBase64,
			Commitment: rpc.CommitmentFinalized,
		},
	)
	if err != nil {
		return result, err
	}
	if res == nil || res.Transaction == nil {
		return result, errors.New("invalid transaction in response")
	}

	solTx, err := solana.TransactionFromDecoder(bin.NewBinDecoder(res.Transaction.GetBinary()))
	if err != nil {
		return result, err
	}
	// log.Print(solTx)
	tx := &Tx{
		SolTx: solTx,
	}
	meta := res.Meta

	if res.Slot > 0 {
		result.BlockIndex = int64(res.Slot)
		if res.BlockTime != nil {
			result.BlockTime = int64(*res.BlockTime)
		}

		// GetRecentBlockhash will be deprecated - GetLatestBlockhash already tested, just switch
		// recent, err := client.SolClient.GetLatestBlockhash(ctx, rpc.CommitmentFinalized)
		recent, err := client.SolClient.GetRecentBlockhash(ctx, rpc.CommitmentFinalized)
		if err != nil {
			// ignore
		} else {
			result.Confirmations = int64(recent.Context.Slot) - result.BlockIndex
		}
	}
	result.Fee = xc.NewAmountBlockchainFromUint64(meta.Fee)

	result.TxID = string(txHash)
	result.ExplorerURL = client.Asset.GetNativeAsset().ExplorerURL + "/tx/" + result.TxID + "?cluster=" + client.Asset.GetNativeAsset().Net
	tx.ParseTransfer()

	// first, check associated token account
	// if detected, call client.GetAccountInfo() to retrieve ATA data
	result.ToAlt = tx.ToAlt()
	if result.ToAlt != "" {
		// ignore err, ata == nil is ok
		ata, _ := client.getParsedAccountInfo(ctx, string(result.ToAlt))
		tx.SetAssociatedTokenAccount(ata)
	}

	// parse tx info - this should happen after ATA is set
	// (in most cases it works also in case or error)
	result.From = tx.From()
	result.To = tx.To()
	result.ContractAddress = tx.ContractAddress()
	result.Amount = tx.Amount()

	return result, nil
}

// FindAssociatedTokenAddress returns the associated token account (ATA) for a given account and token
func FindAssociatedTokenAddress(addr string, contract string) (string, error) {
	address, err := solana.PublicKeyFromBase58(addr)
	if err != nil {
		return "", err
	}
	mint, err := solana.PublicKeyFromBase58(contract)
	if err != nil {
		return "", err
	}
	associatedAddr, _, err := solana.FindAssociatedTokenAddress(address, mint)
	if err != nil {
		return "", err
	}
	return associatedAddr.String(), nil
}

// FetchNativeBalance fetches account balance for a Solana address
func (client *Client) FetchNativeBalance(ctx context.Context, address xc.Address) (xc.AmountBlockchain, error) {
	zero := xc.NewAmountBlockchainFromUint64(0)
	out, err := client.SolClient.GetBalance(
		ctx,
		solana.MustPublicKeyFromBase58(string(address)),
		rpc.CommitmentFinalized,
	)
	if err != nil {
		return zero, fmt.Errorf("failed to get balance for '%v': %v", address, err)
	}
	if out == nil {
		return xc.NewAmountBlockchainFromUint64(0), nil
	}

	return xc.NewAmountBlockchainFromUint64(out.Value), nil
}

// FetchBalance fetches token balance for a Solana address
func (client *Client) FetchBalance(ctx context.Context, address xc.Address) (xc.AmountBlockchain, error) {
	return client.FetchBalanceForAsset(ctx, address, client.Asset)
}

// FetchBalanceForAsset fetches a specific token balance which may not be the asset configured for the client
func (client *Client) FetchBalanceForAsset(ctx context.Context, address xc.Address, assetCfg xc.ITask) (xc.AmountBlockchain, error) {
	if token, ok := assetCfg.(*xc.TokenAssetConfig); ok {
		return client.fetchContractBalance(ctx, address, token.Contract)
	}
	// TODO remove assetconfig
	if assetCfg.GetAssetConfig().Contract != "" {
		return client.fetchContractBalance(ctx, address, assetCfg.GetAssetConfig().Contract)
	}
	return client.FetchNativeBalance(ctx, address)
}

// fetchContractBalance fetches a specific token balance for a Solana address
func (client *Client) fetchContractBalance(ctx context.Context, address xc.Address, contract string) (xc.AmountBlockchain, error) {
	zero := xc.NewAmountBlockchainFromUint64(0)
	ataStr, err := FindAssociatedTokenAddress(string(address), contract)
	if err != nil {
		return zero, err
	}
	ata := solana.MustPublicKeyFromBase58(ataStr)

	out, err := client.SolClient.GetTokenAccountBalance(
		ctx,
		ata,
		rpc.CommitmentFinalized,
	)
	if err != nil {
		if strings.Contains(err.Error(), "could not find account") {
			// account not found => balance is 0
			return zero, nil
		}
		return zero, fmt.Errorf("failed to get balance for '%v': %v", address, err)
	}
	if out == nil || out.Value == nil {
		return xc.NewAmountBlockchainFromUint64(0), nil
	}

	balance := xc.NewAmountBlockchainFromStr(out.Value.Amount)
	return balance, nil
}
