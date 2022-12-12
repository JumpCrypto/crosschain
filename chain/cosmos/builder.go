package cosmos

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/types"
	signingtypes "github.com/cosmos/cosmos-sdk/types/tx/signing"
	"github.com/cosmos/cosmos-sdk/x/auth/signing"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	ethermintCodec "github.com/evmos/ethermint/encoding/codec"
	xc "github.com/jumpcrypto/crosschain"
	terraApp "github.com/terra-money/core/app"
	wasmtypes "github.com/terra-money/core/x/wasm/types"
)

// TxBuilder for Cosmos
type TxBuilder struct {
	xc.TxBuilder
	Asset           xc.AssetConfig
	CosmosTxConfig  client.TxConfig
	CosmosTxBuilder client.TxBuilder
}

// NewTxBuilder creates a new Cosmos TxBuilder
func NewTxBuilder(asset xc.AssetConfig) (xc.TxBuilder, error) {
	cosmosCfg := terraApp.MakeEncodingConfig()
	ethermintCodec.RegisterInterfaces(cosmosCfg.InterfaceRegistry)

	return TxBuilder{
		Asset:           asset,
		CosmosTxConfig:  cosmosCfg.TxConfig,
		CosmosTxBuilder: cosmosCfg.TxConfig.NewTxBuilder(),
	}, nil
}

// NewTransfer creates a new transfer for an Asset, either native or token
func (txBuilder TxBuilder) NewTransfer(from xc.Address, to xc.Address, amount xc.AmountBlockchain, input xc.TxInput) (xc.Tx, error) {
	if isNativeAsset(txBuilder.Asset) {
		return txBuilder.NewNativeTransfer(from, to, amount, input)
	}
	return txBuilder.NewTokenTransfer(from, to, amount, input)
}

// NewNativeTransfer creates a new transfer for a native asset
func (txBuilder TxBuilder) NewNativeTransfer(from xc.Address, to xc.Address, amount xc.AmountBlockchain, input xc.TxInput) (xc.Tx, error) {
	txInput := input.(*TxInput)
	asset := txBuilder.Asset
	amountInt := big.Int(amount)

	if txInput.GasLimit == 0 {
		txInput.GasLimit = 110_000
	}

	msgSend := &banktypes.MsgSend{
		FromAddress: string(from),
		ToAddress:   string(to),
		Amount: types.Coins{
			{
				Denom:  asset.ChainCoin,
				Amount: types.NewIntFromBigInt(&amountInt),
			},
		},
	}

	return txBuilder.createTxWithMsg(from, to, amount, txInput, msgSend)
}

// NewTokenTransfer creates a new transfer for a token asset
func (txBuilder TxBuilder) NewTokenTransfer(from xc.Address, to xc.Address, amount xc.AmountBlockchain, input xc.TxInput) (xc.Tx, error) {
	txInput := input.(*TxInput)
	asset := txBuilder.Asset

	// Terra Classic: most tokens are actually native tokens
	// in crosschain we can treat them interchangeably as native or non-native assets
	// if contract isn't a valid address, they're native tokens
	if isNativeAsset(asset) {
		return txBuilder.NewNativeTransfer(from, to, amount, input)
	}

	if txInput.GasLimit == 0 {
		txInput.GasLimit = 900_000
	}

	contractTransferMsg := fmt.Sprintf(`{"transfer": {"amount": "%s", "recipient": "%s"}}`, amount.String(), to)
	msgSend := &wasmtypes.MsgExecuteContract{
		Sender:     string(from),
		Contract:   asset.Contract,
		ExecuteMsg: json.RawMessage(contractTransferMsg),
	}

	return txBuilder.createTxWithMsg(from, to, amount, txInput, msgSend)
}

func accAddressFromBech32WithPrefix(address string, prefix string) ([]byte, error) {
	if len(strings.TrimSpace(address)) == 0 {
		return nil, errors.New("empty address string is not allowed")
	}

	addressBytes, err := types.GetFromBech32(address, prefix)
	if err != nil {
		return nil, err
	}

	err = types.VerifyAddressFormat(addressBytes)
	if err != nil {
		return nil, err
	}

	return addressBytes, nil
}

// createTxWithMsg creates a new Tx given Cosmos Msg
func (txBuilder TxBuilder) createTxWithMsg(from xc.Address, to xc.Address, amount xc.AmountBlockchain, input *TxInput, msg types.Msg) (xc.Tx, error) {
	asset := txBuilder.Asset
	cosmosTxConfig := txBuilder.CosmosTxConfig
	cosmosBuilder := txBuilder.CosmosTxBuilder

	err := cosmosBuilder.SetMsgs(msg)
	if err != nil {
		return nil, err
	}

	_, err = accAddressFromBech32WithPrefix(string(from), asset.ChainPrefix)
	if err != nil {
		return nil, err
	}

	_, err = accAddressFromBech32WithPrefix(string(to), asset.ChainPrefix)
	if err != nil {
		return nil, err
	}

	cosmosBuilder.SetMemo(input.Memo)
	cosmosBuilder.SetGasLimit(input.GasLimit)
	cosmosBuilder.SetFeeAmount(types.Coins{
		{
			Denom:  asset.ChainCoin,
			Amount: types.NewIntFromUint64(uint64(input.GasPrice * float64(input.GasLimit))),
		},
	})

	sigMode := signingtypes.SignMode_SIGN_MODE_DIRECT
	sigsV2 := []signingtypes.SignatureV2{
		{
			PubKey: input.FromPublicKey,
			Data: &signingtypes.SingleSignatureData{
				SignMode:  sigMode,
				Signature: nil,
			},
			Sequence: input.Sequence,
		},
	}
	err = cosmosBuilder.SetSignatures(sigsV2...)
	if err != nil {
		return nil, err
	}

	signerData := signing.SignerData{
		AccountNumber: input.AccountNumber,
		ChainID:       asset.ChainIDStr,
		Sequence:      input.Sequence,
	}
	sighashData, err := cosmosTxConfig.SignModeHandler().GetSignBytes(sigMode, signerData, cosmosBuilder.GetTx())
	if err != nil {
		return nil, err
	}

	return &Tx{
		CosmosTx:        cosmosBuilder.GetTx(),
		ParsedTransfer:  msg,
		CosmosTxBuilder: cosmosBuilder,
		CosmosTxEncoder: cosmosTxConfig.TxEncoder(),
		SigsV2:          sigsV2,
		TxDataToSign:    getSighash(asset.NativeAsset, sighashData),
	}, nil
}
