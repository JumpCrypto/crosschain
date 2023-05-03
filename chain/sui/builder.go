package sui

import (
	"errors"
	"fmt"
	"strings"

	xc "github.com/jumpcrypto/crosschain"
	"github.com/jumpcrypto/crosschain/chain/sui/generated/bcs"
)

type TxBuilder struct {
	Asset xc.ITask
}

var _ xc.TxTokenBuilder = &TxBuilder{}

// NewTxBuilder creates a new Template TxBuilder
func NewTxBuilder(asset xc.ITask) (xc.TxBuilder, error) {
	return &TxBuilder{
		Asset: asset,
	}, nil
}

// NewTransfer creates a new transfer for an Asset, either native or token
func (txBuilder TxBuilder) NewTransfer(from xc.Address, to xc.Address, amount xc.AmountBlockchain, input xc.TxInput) (xc.Tx, error) {
	var local_input TxInput
	var ok bool
	// Either ptr or full type is okay.
	if local_input, ok = input.(TxInput); !ok {
		var ptr *TxInput
		if ptr, ok = (input.(*TxInput)); !ok {
			return &Tx{}, errors.New("xc.TxInput is not from an sui chain")
		}
		local_input = *ptr
	}
	from = xc.Address(strings.Replace(string(from), "0x", "", 1))
	to = xc.Address(strings.Replace(string(to), "0x", "", 1))

	fromData, err := hexToAddress(string(from))
	if err != nil {
		return &Tx{}, fmt.Errorf("could not decode from address: %v", err)
	}
	toPure, err := hexToPure(string(to))
	if err != nil {
		return &Tx{}, fmt.Errorf("could not decode to address: %v", err)
	}
	gasObjectId, err := hexToObjectID(local_input.GasCoin.CoinObjectId.String())
	if err != nil {
		return &Tx{}, fmt.Errorf("could not decode gas coin object id: %v", err)
	}
	gasDigest, err := base58ToObjectDigest(local_input.GasCoin.CoinObjectId.String())
	if err != nil {
		return &Tx{}, fmt.Errorf("could not decode gas coin digest: %v", err)
	}

	local_input.SortCoins()

	cmd_inputs := []bcs.CallArg{}
	commands := []bcs.Command{
		&bcs.Command__TransferObjects{
			Field0: []bcs.Argument{
				ArgumentResult(0),
			},
			Field1: ArgumentInput(1),
		},
	}
	// first input is our primary object that we're merging into, and that we're sending from.
	first_coin := local_input.Coins[0]
	id, err := hexToPure(first_coin.CoinObjectId.String())
	if err != nil {
		return &Tx{}, fmt.Errorf("could not decode coin id: %v", err)
	}
	cmd_inputs = append(cmd_inputs, id)

	// Let's merge together up to 50 sui inputs into our largest input.
	if len(local_input.Coins) > 1 {
		for i, coin := range local_input.Coins[1:50] {
			id, err := hexToPure(coin.CoinObjectId.String())
			if err != nil {
				return &Tx{}, fmt.Errorf(": %v", err)
			}
			cmd_inputs = append(cmd_inputs, id)
			commands = append(commands, &bcs.Command__MergeCoins{
				Field0: ArgumentInput(0),
				Field1: []bcs.Argument{
					ArgumentInput(uint16(i)),
				},
			})
		}
	}

	// now let's spend the first coin by splitting `amt` from it
	cmd_inputs = append(cmd_inputs, u64ToPure(amount.Uint64()))
	commands = append(commands, &bcs.Command__SplitCoins{
		Field0: ArgumentInput(0),
		Field1: []bcs.Argument{
			// the last input has the amount
			ArgumentInput(uint16(len(cmd_inputs) - 1)),
		},
	})

	// send the new object that is the result of split
	cmd_inputs = append(cmd_inputs, toPure)
	commands = append(commands, &bcs.Command__SplitCoins{
		Field0: ArgumentResult(0),
		Field1: []bcs.Argument{
			// the last input has the destination
			ArgumentInput(uint16(len(cmd_inputs) - 1)),
		},
	})

	gasCoin := ObjectRef{
		Field0: gasObjectId,
		Field1: bcs.SequenceNumber(local_input.GasCoin.Version.BigInt().Uint64()),
		Field2: gasDigest,
	}

	tx := bcs.TransactionData__V1{
		Value: bcs.TransactionDataV1{
			GasData: bcs.GasData{
				Payment: []struct {
					Field0 bcs.ObjectID
					Field1 bcs.SequenceNumber
					Field2 bcs.ObjectDigest
				}{
					gasCoin,
				},
				Owner:  fromData,
				Price:  local_input.GasPrice,
				Budget: local_input.GasBudget,
			},
			Sender:     fromData,
			Expiration: &bcs.TransactionExpiration__None{},
			Kind: &bcs.TransactionKind__ProgrammableTransaction{
				Value: bcs.ProgrammableTransaction{
					Inputs:   cmd_inputs,
					Commands: commands,
				},
			},
		},
	}

	return &Tx{
		Input: local_input,
		tx:    tx,
	}, nil
}

func (txBuilder TxBuilder) NewNativeTransfer(from xc.Address, to xc.Address, amount xc.AmountBlockchain, input xc.TxInput) (xc.Tx, error) {
	return txBuilder.NewTransfer(from, to, amount, input)
}
func (txBuilder TxBuilder) NewTokenTransfer(from xc.Address, to xc.Address, amount xc.AmountBlockchain, input xc.TxInput) (xc.Tx, error) {
	// The token is already in the coins in the tx_input so txbuilding is the exact same.
	return txBuilder.NewTransfer(from, to, amount, input)
}
