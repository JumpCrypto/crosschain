package bitcoin

import (
	"errors"
	"fmt"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
	"github.com/btcsuite/btcutil"
	xc "github.com/jumpcrypto/crosschain"
)

const TxVersion int32 = 2

// TxBuilder for Bitcoin
type TxBuilder struct {
	Asset  *xc.AssetConfig
	Params *chaincfg.Params
	isBch  bool
}

// NewTxBuilder creates a new Bitcoin TxBuilder
func NewTxBuilder(cfgI xc.ITask) (xc.TxBuilder, error) {
	asset := cfgI.GetNativeAsset()
	params, err := GetParams(asset)
	if err != nil {
		return TxBuilder{}, err
	}
	return TxBuilder{
		Asset:  asset,
		Params: params,
		isBch:  asset.NativeAsset == xc.BCH,
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

	var local_input TxInput
	var ok bool
	// Either ptr or full type is okay.
	if local_input, ok = input.(TxInput); !ok {
		var ptr *TxInput
		if ptr, ok = (input.(*TxInput)); !ok {
			return &Tx{}, errors.New("xc.TxInput is not from a bitcoin chain")
		}
		local_input = *ptr
	}
	// Only need to save min utxo for the transfer.
	totalSpend := local_input.allocateMinUtxoSet(amount, 10)
	local_input.UnspentOutputs = []Output{}

	gasPrice := local_input.GasPricePerByte
	// 255 for bitcoin, 300 for bch
	estimatedTxBytesLength := xc.NewAmountBlockchainFromUint64(uint64(255 * len(local_input.Inputs)))
	if txBuilder.Asset.NativeAsset == xc.BCH {
		estimatedTxBytesLength = xc.NewAmountBlockchainFromUint64(uint64(300 * len(local_input.Inputs)))
	}
	fee := gasPrice.Mul(&estimatedTxBytesLength)

	transferAmountAndFee := amount.Add(&fee)
	unspentAmountMinusTransferAndFee := totalSpend.Sub(&transferAmountAndFee)
	recipients := []Recipient{
		{
			To:    to,
			Value: amount,
		},
		{
			To:    from,
			Value: unspentAmountMinusTransferAndFee,
		},
	}

	msgTx := wire.NewMsgTx(TxVersion)

	for _, input := range local_input.Inputs {
		hash := chainhash.Hash{}
		copy(hash[:], input.Hash)
		msgTx.AddTxIn(wire.NewTxIn(wire.NewOutPoint(&hash, input.Index), nil, nil))
	}

	// Outputs
	for _, recipient := range recipients {
		addr, err := btcutil.DecodeAddress(string(recipient.To), txBuilder.Params)
		fmt.Println("trying to decode ", recipient.To)
		if err != nil {
			// try to decode as BCH
			bchaddr, err2 := DecodeBchAddress(string(recipient.To), txBuilder.Params)
			if err2 != nil {
				return nil, err
			}
			addr, err = BchAddressFromBytes(bchaddr, txBuilder.Params)
			if err != nil {
				return nil, err
			}
		}
		script, err := txscript.PayToAddrScript(addr)
		if err != nil {
			fmt.Println(err)
			fmt.Println("trying to paytoaddr", recipient.To)
			return nil, err
		}
		value := recipient.Value.Int().Int64()
		if value < 0 {
			return nil, fmt.Errorf("expected value >= 0, got value %v", value)
		}
		msgTx.AddTxOut(wire.NewTxOut(value, script))
	}

	tx := Tx{
		msgTx: msgTx,

		from:   from,
		to:     to,
		amount: amount,
		input:  local_input,

		recipients: recipients,
		isBch:      txBuilder.isBch,
	}
	return &tx, nil
}

// NewTokenTransfer creates a new transfer for a token asset
func (txBuilder TxBuilder) NewTokenTransfer(from xc.Address, to xc.Address, amount xc.AmountBlockchain, input xc.TxInput) (xc.Tx, error) {
	return nil, errors.New("not implemented")
}
