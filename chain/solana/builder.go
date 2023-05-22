package solana

import (
	"errors"
	"fmt"

	"github.com/gagliardetto/solana-go"
	ata "github.com/gagliardetto/solana-go/programs/associated-token-account"
	"github.com/gagliardetto/solana-go/programs/system"
	"github.com/gagliardetto/solana-go/programs/token"
	xc "github.com/jumpcrypto/crosschain"
)

// TxBuilder for Solana
type TxBuilder struct {
	Asset xc.ITask
}

// NewTxBuilder creates a new Solana TxBuilder
func NewTxBuilder(asset xc.ITask) (xc.TxBuilder, error) {
	return TxBuilder{
		Asset: asset,
	}, nil
}

// NewTransfer creates a new transfer for an Asset, either native or token
func (txBuilder TxBuilder) NewTransfer(from xc.Address, to xc.Address, amount xc.AmountBlockchain, input xc.TxInput) (xc.Tx, error) {
	if _, ok := txBuilder.Asset.(*xc.TaskConfig); ok {
		return txBuilder.NewTask(from, to, amount, input)
	}

	if asset, ok := txBuilder.Asset.(*xc.TokenAssetConfig); ok {
		if asset.Type == xc.AssetTypeNative {
			return txBuilder.NewNativeTransfer(from, to, amount, input)
		} else {
			return txBuilder.NewTokenTransfer(from, to, amount, input)
		}
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
	asset := txBuilder.Asset.GetAssetConfig()
	if asset.Type != xc.AssetTypeToken {
		if _, ok := txBuilder.Asset.(*xc.TokenAssetConfig); !ok {
			return nil, errors.New("asset is not of type token")
		}
	}
	txInput := input.(*TxInput)

	contract := asset.Contract
	if token, ok := txBuilder.Asset.(*xc.TokenAssetConfig); ok && contract == "" {
		contract = token.Contract
	}
	decimals := uint8(asset.Decimals)

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
	return txBuilder.buildSolanaTx(instructions, accountFrom, txInput)
}

func (txBuilder TxBuilder) buildSolanaTx(instructions []solana.Instruction, accountFrom solana.PublicKey, txInput *TxInput) (xc.Tx, error) {
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

func (txBuilder TxBuilder) NewTask(from xc.Address, to xc.Address, amount xc.AmountBlockchain, input xc.TxInput) (xc.Tx, error) {
	txInput := input.(*TxInput)
	task := txBuilder.Asset.(*xc.TaskConfig)
	switch task.Code {
	case "WrapTx":
		return txBuilder.BuildWrapTx(from, to, amount, txInput)
	case "UnwrapEverythingTx":
		return txBuilder.BuildUnwrapEverythingTx(from, to, amount, txInput)
	}
	return &Tx{}, fmt.Errorf("not implemented task: '%s'", txBuilder.Asset.ID())
}

func (txBuilder TxBuilder) BuildWrapTx(from xc.Address, to xc.Address, amount xc.AmountBlockchain, txInput *TxInput) (xc.Tx, error) {
	// use the dst asset
	task := txBuilder.Asset.(*xc.TaskConfig)
	asset := task.DstAsset.GetAssetConfig()

	accountFrom, err := solana.PublicKeyFromBase58(string(from))
	if err != nil {
		return nil, err
	}

	contract := asset.Contract
	accountContract, err := solana.PublicKeyFromBase58(string(contract))
	if err != nil {
		return nil, err
	}

	ataFromStr, err := FindAssociatedTokenAddress(string(from), string(contract))
	if err != nil {
		return nil, err
	}
	ataFrom := solana.MustPublicKeyFromBase58(ataFromStr)

	// instructions to:
	// - transfer to the ATA (system.NewTransferInstruction())
	// - create the ATA (associatedtokenaccount.NewCreateInstruction())
	instructions := []solana.Instruction{
		ata.NewCreateInstruction(
			accountFrom,
			accountFrom,
			accountContract,
		).Build(),
		system.NewTransferInstruction(
			amount.Uint64(),
			accountFrom,
			ataFrom,
		).Build(),
	}

	return txBuilder.buildSolanaTx(instructions, accountFrom, txInput)
}

func (txBuilder TxBuilder) BuildUnwrapEverythingTx(from xc.Address, to xc.Address, amount xc.AmountBlockchain, txInput *TxInput) (xc.Tx, error) {
	asset := txBuilder.Asset.GetAssetConfig()
	accountFrom, err := solana.PublicKeyFromBase58(string(from))
	if err != nil {
		return nil, err
	}

	contract := asset.Contract
	ataFromStr, err := FindAssociatedTokenAddress(string(from), string(contract))
	if err != nil {
		return nil, err
	}
	ataFrom := solana.MustPublicKeyFromBase58(ataFromStr)

	// instructions to:
	// - close the ATA (token.NewCloseAccountInstruction()) -- unwraps everything into from account
	instructions := []solana.Instruction{
		token.NewCloseAccountInstruction(ataFrom, accountFrom, accountFrom, nil).Build(),
	}

	return txBuilder.buildSolanaTx(instructions, accountFrom, txInput)
}
