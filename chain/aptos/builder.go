package aptos

import (
	"errors"

	transactionbuilder "github.com/coming-chat/go-aptos/transaction_builder"
	xc "github.com/jumpcrypto/crosschain"
)

// TxBuilder for Template
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
	if _, ok := txBuilder.Asset.(*xc.TokenAssetConfig); ok {
		return txBuilder.NewTokenTransfer(from, to, amount, input)
	}
	return txBuilder.NewNativeTransfer(from, to, amount, input)
}

// NewNativeTransfer creates a new transfer for a native asset
func (txBuilder TxBuilder) NewNativeTransfer(from xc.Address, to xc.Address, amount xc.AmountBlockchain, input xc.TxInput) (xc.Tx, error) {
	var local_input TxInput
	var ok bool
	// Either ptr or full type is okay.
	if local_input, ok = input.(TxInput); !ok {
		var ptr *TxInput
		if ptr, ok = (input.(*TxInput)); !ok {
			return &Tx{}, errors.New("xc.TxInput is not from an aptos chain")
		}
		local_input = *ptr
	}
	to_addr := [transactionbuilder.ADDRESS_LENGTH]byte{}
	from_addr := [transactionbuilder.ADDRESS_LENGTH]byte{}
	copy(from_addr[:], mustDecodeHex(string(from)))
	copy(to_addr[:], mustDecodeHex(string(to)))
	toAmountBytes := transactionbuilder.BCSSerializeBasicValue(amount.Int().Uint64())

	chain_id := local_input.ChainId
	moduleName, err := transactionbuilder.NewModuleIdFromString("0x1::aptos_account")
	if err != nil {
		return &Tx{}, err
	}
	payload := transactionbuilder.TransactionPayloadEntryFunction{
		ModuleName:   *moduleName,
		FunctionName: "transfer",
		Args: [][]byte{
			to_addr[:], toAmountBytes,
		},
	}

	return &Tx{
		tx: transactionbuilder.RawTransaction{
			Sender:         from_addr,
			SequenceNumber: local_input.SequenceNumber,
			Payload:        payload,
			MaxGasAmount:   local_input.GasLimit,
			GasUnitPrice:   local_input.GasPrice,
			// ~1 hour expiration
			ExpirationTimestampSecs: local_input.Timestamp + 60*60,
			ChainId:                 uint8(chain_id),
		},
		Input: local_input,
	}, nil
}

// NewTokenTransfer creates a new transfer for a token asset
func (txb *TxBuilder) NewTokenTransfer(from xc.Address, to xc.Address, amount xc.AmountBlockchain, input xc.TxInput) (xc.Tx, error) {
	var local_input TxInput
	var ok bool
	// Either ptr or full type is okay.
	if local_input, ok = input.(TxInput); !ok {
		var ptr *TxInput
		if ptr, ok = (input.(*TxInput)); !ok {
			return &Tx{}, errors.New("xc.TxInput is not from an aptos chain")
		}
		local_input = *ptr
	}
	to_addr := [transactionbuilder.ADDRESS_LENGTH]byte{}
	from_addr := [transactionbuilder.ADDRESS_LENGTH]byte{}
	copy(from_addr[:], mustDecodeHex(string(from)))
	copy(to_addr[:], mustDecodeHex(string(to)))
	toAmountBytes := transactionbuilder.BCSSerializeBasicValue(amount.Int().Uint64())

	contract := ""
	if token, ok := txb.Asset.(*xc.TokenAssetConfig); ok {
		contract = token.Contract
	}
	if contract == "" {
		contract = txb.Asset.GetAssetConfig().Contract
	}

	typeTag, err := transactionbuilder.NewTypeTagStructFromString(contract)
	if err != nil {
		return nil, err
	}

	chain_id := local_input.ChainId
	moduleName, err := transactionbuilder.NewModuleIdFromString("0x1::coin")
	if err != nil {
		return &Tx{}, err
	}
	payload := transactionbuilder.TransactionPayloadEntryFunction{
		ModuleName:   *moduleName,
		FunctionName: "transfer",
		TyArgs:       []transactionbuilder.TypeTag{*typeTag},
		Args: [][]byte{
			to_addr[:], toAmountBytes,
		},
	}
	return &Tx{
		tx: transactionbuilder.RawTransaction{
			Sender:         from_addr,
			SequenceNumber: local_input.SequenceNumber,
			Payload:        payload,
			MaxGasAmount:   local_input.GasLimit,
			GasUnitPrice:   local_input.GasPrice,
			// ~1 hour expiration
			ExpirationTimestampSecs: local_input.Timestamp + 60*60,
			ChainId:                 uint8(chain_id),
		},
		Input: local_input,
	}, nil
}
