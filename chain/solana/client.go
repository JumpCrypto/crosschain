package solana

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	xc "github.com/jumpcrypto/crosschain"

	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/token"
	"github.com/gagliardetto/solana-go/rpc"
)

// Client for Solana
type Client struct {
	SolClient *rpc.Client
}

// TxInput for Solana
type TxInput struct {
	xc.TxInput
	RecentBlockHash solana.Hash
}

// NewClient returns a new JSON-RPC Client to the Solana node
func NewClient(cfg xc.AssetConfig) (*Client, error) {
	solClient := rpc.New(cfg.URL)
	return &Client{
		SolClient: solClient,
	}, nil
}

// FetchTxInput returns tx input for a Solana tx, namely a RecentBlockHash
func (client *Client) FetchTxInput(ctx context.Context, from xc.Address) (xc.TxInput, error) {
	// GetRecentBlockhash will be deprecated - GetLatestBlockhash already tested, just switch
	// recent, err := client.SolClient.GetLatestBlockhash(context.Background(), rpc.CommitmentFinalized)
	recent, err := client.SolClient.GetRecentBlockhash(context.Background(), rpc.CommitmentFinalized)
	if err != nil {
		return nil, err
	}
	if recent == nil || recent.Value == nil {
		return nil, errors.New("error fetching blockhash")
	}
	return TxInput{
		RecentBlockHash: recent.Value.Blockhash,
	}, nil
}

// SubmitTx submits a Solana tx
func (client *Client) SubmitTx(ctx context.Context, txInput xc.Tx) error {
	tx := txInput.(*Tx)
	_, err := client.SolClient.SendTransactionWithOpts(
		context.Background(),
		tx.SolTx,
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
		context.Background(),
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
		// recent, err := client.SolClient.GetLatestBlockhash(context.Background(), rpc.CommitmentFinalized)
		recent, err := client.SolClient.GetRecentBlockhash(context.Background(), rpc.CommitmentFinalized)
		if err != nil {
			// ignore
		} else {
			result.Confirmations = int64(recent.Context.Slot) - result.BlockIndex
		}
	}
	result.Fee = xc.NewAmountBlockchainFromUint64(meta.Fee)

	result.TxID = string(txHash)
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
	result.Amount = tx.Value()

	return result, nil
}

// AccountBalance returns the SOL balance for an account
func (client *Client) AccountBalance(ctx context.Context, addr xc.Address) (xc.AmountBlockchain, error) {
	out, err := client.SolClient.GetBalance(
		context.Background(),
		solana.MustPublicKeyFromBase58(string(addr)),
		rpc.CommitmentFinalized,
	)
	if err != nil {
		return xc.NewAmountBlockchainFromUint64(0), fmt.Errorf("failed to get balance for '%v': %v", addr, err)
	}
	if out == nil {
		return xc.NewAmountBlockchainFromUint64(0), nil
	}

	return xc.NewAmountBlockchainFromUint64(out.Value), nil
}

// TokenBalance returns the token balance for an account and token
func (client *Client) TokenBalance(ctx context.Context, addr xc.Address, contract string) (xc.AmountBlockchain, error) {
	address := solana.MustPublicKeyFromBase58(string(addr))
	mint := solana.MustPublicKeyFromBase58(contract)
	associatedAddr, _, err := solana.FindAssociatedTokenAddress(address, mint)
	if err != nil {
		return xc.NewAmountBlockchainFromUint64(0), fmt.Errorf("failed to get ATT for '%v': %v", addr, err)
	}
	out, err := client.SolClient.GetTokenAccountBalance(
		context.Background(),
		associatedAddr,
		rpc.CommitmentFinalized,
	)
	if err != nil {
		if strings.Contains(err.Error(), "could not find account") {
			// account not found => balance is 0
			return xc.NewAmountBlockchainFromUint64(0), nil
		}
		return xc.NewAmountBlockchainFromUint64(0), fmt.Errorf("failed to get balance for '%v': %v", addr, err)
	}
	if out == nil || out.Value == nil {
		return xc.NewAmountBlockchainFromUint64(0), nil
	}

	balance, err := strconv.ParseUint(out.Value.Amount, 10, 64)
	return xc.NewAmountBlockchainFromUint64(balance), err
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
