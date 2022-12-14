package solana

import (
	"errors"

	xc "github.com/jumpcrypto/crosschain"

	"github.com/gagliardetto/solana-go"
	ata "github.com/gagliardetto/solana-go/programs/associated-token-account"
	"github.com/gagliardetto/solana-go/programs/system"
	"github.com/gagliardetto/solana-go/programs/token"
)

// TxBuilder for Solana
type TxBuilder struct {
	Asset xc.AssetConfig
}

// NewTxBuilder creates a new Solana TxBuilder
func NewTxBuilder(asset xc.AssetConfig) (xc.TxBuilder, error) {
	return TxBuilder{
		Asset: asset,
	}, nil
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
	accountFrom, err := solana.PublicKeyFromBase58(string(from))
	if err != nil {
		return nil, err
	}
	accountTo, err := solana.PublicKeyFromBase58(string(to))
	if err != nil {
		return nil, err
	}

	// txLog := map[string]string{
	// 	"type":      "system.Transfer",
	// 	"lamports":  amount.String(),
	// 	"funding":   accountFrom.String(),
	// 	"recipient": accountTo.String(),
	// }
	// log.Print(txLog)

	tx, err := solana.NewTransaction(
		[]solana.Instruction{
			system.NewTransferInstruction(
				amount.Uint64(),
				accountFrom,
				accountTo,
			).Build(),
		},
		input.(*TxInput).RecentBlockHash,
		solana.TransactionPayer(accountFrom),
	)
	if err != nil {
		return nil, err
	}
	return &Tx{
		SolTx: tx,
	}, nil
}

// NewTokenTransfer creates a new transfer for a token asset
func (txBuilder TxBuilder) NewTokenTransfer(from xc.Address, to xc.Address, amount xc.AmountBlockchain, input xc.TxInput) (xc.Tx, error) {
	if txBuilder.Asset.Type != xc.AssetTypeToken {
		return nil, errors.New("asset is not of type token")
	}
	txInput := input.(*TxInput)

	contract := txBuilder.Asset.Contract
	decimals := uint8(txBuilder.Asset.Decimals)

	accountFrom, err := solana.PublicKeyFromBase58(string(from))
	if err != nil {
		return nil, err
	}

	accountContract, err := solana.PublicKeyFromBase58(string(contract))
	if err != nil {
		return nil, err
	}

	accountTo, err := solana.PublicKeyFromBase58(string(to))
	if err != nil {
		return nil, err
	}

	ataFromStr, err := FindAssociatedTokenAddress(string(from), string(contract))
	if err != nil {
		return nil, err
	}
	ataFrom := solana.MustPublicKeyFromBase58(ataFromStr)

	ataTo := accountTo
	if !txInput.ToIsATA {
		ataToStr, err := FindAssociatedTokenAddress(string(to), string(contract))
		if err != nil {
			return nil, err
		}
		ataTo = solana.MustPublicKeyFromBase58(ataToStr)
	}

	// txLog := map[string]string{
	// 	"type":     "token.TransferChecked",
	// 	"amount":   amount.String(),
	// 	"decimals": strconv.Itoa(int(decimals)),
	// 	"source":   ataFrom.String(),
	// 	"mint":     accountContract.String(),
	// 	"dest":     ataTo.String(),
	// 	"owner":    accountFrom.String(),
	// }
	// log.Print(txLog)

	instructions := []solana.Instruction{}
	if txInput.ShouldCreateATA {
		instructions = append(instructions,
			ata.NewCreateInstruction(
				accountFrom,
				accountTo,
				accountContract,
			).Build(),
		)
	}
	instructions = append(instructions,
		token.NewTransferCheckedInstruction(
			amount.Uint64(),
			decimals,
			ataFrom,
			accountContract,
			ataTo,
			accountFrom,
			[]solana.PublicKey{
				accountFrom,
			},
		).Build(),
	)
	tx, err := solana.NewTransaction(
		instructions,
		txInput.RecentBlockHash,
		solana.TransactionPayer(accountFrom),
	)
	if err != nil {
		return nil, err
	}
	return &Tx{
		SolTx: tx,
	}, nil
}
