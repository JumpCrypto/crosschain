package sui

import (
	"errors"
	"fmt"

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
	if len(local_input.Pubkey) == 0 {
		return &Tx{}, errors.New("must set public key on TxInput for SUI")
	}

	// from = xc.Address(strings.Replace(string(from), "0x", "", 1))
	// to = xc.Address(strings.Replace(string(to), "0x", "", 1))

	fromData, err := hexToAddress(string(from))
	if err != nil {
		return &Tx{}, fmt.Errorf("could not decode from address: %v", err)
	}
	toPure, err := hexToPure(string(to))
	if err != nil {
		return &Tx{}, fmt.Errorf("could not decode to address: %v", err)
	}

	gasObjectId, err := hexToObjectID(local_input.GasCoin.CoinObjectId.String())
	// gasObjectId, err := hexToObjectID("0x26b1fb28f2b0d543b17ffd0034c6e446cf998dc05ffe376ee8b9e00f4934a21c")
	if err != nil {
		return &Tx{}, fmt.Errorf("could not decode gas coin object id: %v", err)
	}
	gasDigest, err := base58ToObjectDigest(local_input.GasCoin.Digest)
	// gasDigest, err := base58ToObjectDigest("GtxXCYfW7v3Z7HLQdCNQ6ZMVuK6Ek8VDwr1AVE423CTd")
	if err != nil {
		return &Tx{}, fmt.Errorf("could not decode gas coin digest: %v", err)
	}
	gasVersion := local_input.GasCoin.Version.BigInt().Uint64()

	local_input.ExcludeGasCoin()
	// Our gas budget should be the minimum of:
	//  - normal budget (2sui)
	//  - balance of the gas coin
	//  - total sui balance remainder after transfering `amount`.

	normal_budget := local_input.GasBudget
	gas_coin_balance := local_input.GasCoin.Balance.Uint64()
	if local_input.TotalBalance().Uint64() < amount.Uint64() {
		return &Tx{}, fmt.Errorf("not enough funds to send after paying for sui gas: budget=%d tf=%d", local_input.GasBudget, amount.Uint64())
	}
	total_remainder := local_input.TotalBalance().Uint64() - amount.Uint64()

	budget := normal_budget
	if gas_coin_balance < budget {
		budget = gas_coin_balance
	}
	if total_remainder < budget {
		budget = total_remainder
	}

	// lower the gas budget to whatever balance is on the gas coin.  no need to split it.
	local_input.GasBudget = budget

	// Our universal transaction goes like this:
	// I. We start with the gas coin and we split it.
	//  	a. The primary of the split is used for paying gas.  It can't be using for anything else
	//	 	   or Sui we have multiple mutating errors.  It should should have enough to cover our total
	//		   gas budget.
	//	    b. The result of the split gets used in the result of the transaction IFF it's the same currency/type.
	// II. We merge the rest of our coins together into one coin (up to say, 50).
	// III. We split this result merged coin into another coin that is the amount we wish to transfer.
	// IV. We send this newly minted coin.
	// So there should always be 4 tx total.

	cmd_inputs := []bcs.CallArg{}
	commands := []bcs.Command{}
	var gasRemainderResult bcs.Argument
	// I. Split the gas coin if necessary
	// Check to see if we can afford the gas budget.
	if len(local_input.Coins) > 0 && local_input.Coins[0].CoinType == local_input.GasCoin.CoinType {
		// Split off the remainder from gas budget
		remainder := local_input.GasCoin.Balance.Uint64() - local_input.GasBudget
		cmd_inputs = append(cmd_inputs, u64ToPure(remainder))

		commands = append(commands, &bcs.Command__SplitCoins{
			Field0: &bcs.Argument__GasCoin{},
			Field1: []bcs.Argument{
				ArgumentInput(uint16(0)),
			},
		})
		gasRemainderResult = ArgumentResult(uint16(len(commands) - 1))

	}

	primaryCoinInput := gasRemainderResult

	// II. merge together rest of coins if needed
	if len(local_input.Coins) > 0 {
		// The first coin becomes our "primary coin"
		primaryCoinInput = ArgumentInput(uint16(len(cmd_inputs)))

		obj, err := coinToObject(local_input.Coins[0])
		if err != nil {
			return &Tx{}, err
		}
		cmd_inputs = append(cmd_inputs, &bcs.CallArg__Object{
			Value: obj,
		})

		merge_inputs := []bcs.Argument{}

		if gasRemainderResult != nil {
			merge_inputs = append(merge_inputs, ArgumentResult(uint16(len(commands)-1)))
		}

		for i, coin := range local_input.Coins[1:] {
			if i > MaxCoinObjects {
				break
			}
			obj, err := coinToObject(coin)
			if err != nil {
				return &Tx{}, err
			}
			merge_inputs = append(merge_inputs, ArgumentInput(uint16(len(cmd_inputs))))

			cmd_inputs = append(cmd_inputs, &bcs.CallArg__Object{
				Value: obj,
			})
		}
		commands = append(commands, &bcs.Command__MergeCoins{
			Field0: primaryCoinInput,
			Field1: merge_inputs,
		})
	}

	if primaryCoinInput == nil {
		// if we only have one coin.. than the primary coin should be the gas coin
		primaryCoinInput = &bcs.Argument__GasCoin{}
	}

	// now let's spend the primary coin by splitting `amt` from it
	commands = append(commands, &bcs.Command__SplitCoins{
		Field0: primaryCoinInput,
		Field1: []bcs.Argument{
			// the last input has the amount
			ArgumentInput(uint16(len(cmd_inputs))),
		},
	})
	cmd_inputs = append(cmd_inputs, u64ToPure(amount.Uint64()))

	// send the new split object
	commands = append(commands, &bcs.Command__TransferObjects{
		Field0: []bcs.Argument{
			// last cmd result has the coin to send
			ArgumentResult(uint16(len(commands) - 1)),
		},
		Field1: ArgumentInput(uint16(len(cmd_inputs))),
	})
	cmd_inputs = append(cmd_inputs, toPure)

	// expires after current epoch
	expiration := bcs.TransactionExpiration__Epoch(local_input.CurrentEpoch)

	gasCoin := ObjectRef{
		Field0: gasObjectId,
		Field1: bcs.SequenceNumber(gasVersion),
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
			Expiration: &expiration,
			Kind: &bcs.TransactionKind__ProgrammableTransaction{
				Value: bcs.ProgrammableTransaction{
					Inputs:   cmd_inputs,
					Commands: commands,
				},
			},
		},
	}

	return &Tx{
		Tx:         tx,
		public_key: local_input.Pubkey,
	}, nil
}

func (txBuilder TxBuilder) NewNativeTransfer(from xc.Address, to xc.Address, amount xc.AmountBlockchain, input xc.TxInput) (xc.Tx, error) {
	return txBuilder.NewTransfer(from, to, amount, input)
}
func (txBuilder TxBuilder) NewTokenTransfer(from xc.Address, to xc.Address, amount xc.AmountBlockchain, input xc.TxInput) (xc.Tx, error) {
	// The token is already in the coins in the tx_input so txbuilding is the exact same.
	return txBuilder.NewTransfer(from, to, amount, input)
}
