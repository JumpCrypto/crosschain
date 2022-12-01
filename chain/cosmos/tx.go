package cosmos

import (
	"encoding/hex"

	xc "github.com/jumpcrypto/crosschain"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/types"
	signingtypes "github.com/cosmos/cosmos-sdk/types/tx/signing"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/tendermint/tendermint/crypto/tmhash"
)

// Tx for Cosmos
type Tx struct {
	CosmosTx       types.Tx
	ParsedTransfer types.Msg
	// aux fields
	CosmosTxBuilder client.TxBuilder
	CosmosTxEncoder types.TxEncoder
	SigsV2          []signingtypes.SignatureV2
	TxDataToSign    []byte
}

// Hash returns the tx hash or id
func (tx Tx) Hash() xc.TxHash {
	txID := tmhash.Sum(tx.Serialize())
	return xc.TxHash(hex.EncodeToString(txID))
}

// Sighash returns the tx payload to sign, aka sighash
func (tx Tx) Sighash() (xc.TxDataToSign, error) {
	return xc.TxDataToSign(tx.TxDataToSign), nil
}

// AddSignature adds a signature to Tx
func (tx Tx) AddSignature(signature xc.TxSignature) error {
	data := tx.SigsV2[0].Data
	signMode := data.(*signingtypes.SingleSignatureData).SignMode
	tx.SigsV2[0].Data = &signingtypes.SingleSignatureData{
		SignMode:  signMode,
		Signature: signature,
	}
	tx.CosmosTxBuilder.SetSignatures(tx.SigsV2...)
	return nil
}

// Serialize serializes a Tx
func (tx Tx) Serialize() []byte {
	if tx.CosmosTxEncoder == nil || tx.CosmosTxBuilder == nil {
		return []byte{}
	}
	serialized, _ := tx.CosmosTxEncoder(tx.CosmosTxBuilder.GetTx())
	return serialized
}

// ParseTransfer parses a Tx as a transfer
// Currently only banktypes.MsgSend is implemented, i.e. only native tokens
func (tx *Tx) ParseTransfer() {
	for _, msg := range tx.CosmosTx.GetMsgs() {
		switch msg := msg.(type) {
		case *banktypes.MsgSend:
			tx.ParsedTransfer = msg
		}
	}
}

// From returns the from address of a Tx
func (tx Tx) From() xc.Address {
	switch tf := tx.ParsedTransfer.(type) {
	case *banktypes.MsgSend:
		from := tf.FromAddress
		return xc.Address(from)
	}
	return xc.Address("")
}

// To returns the to address of a Tx
func (tx Tx) To() xc.Address {
	switch tf := tx.ParsedTransfer.(type) {
	case *banktypes.MsgSend:
		to := tf.ToAddress
		return xc.Address(to)
	}
	return xc.Address("")
}

// ContractAddress returns the contract address of a Tx, if any
func (tx Tx) ContractAddress() xc.ContractAddress {
	// not implemented
	return xc.ContractAddress("")
}

// Amount returns the amount of a Tx
func (tx Tx) Amount() xc.AmountBlockchain {
	switch tf := tx.ParsedTransfer.(type) {
	case *banktypes.MsgSend:
		amount := tf.Amount[0].Amount.BigInt()
		return xc.AmountBlockchain(*amount)
	}
	return xc.NewAmountBlockchainFromUint64(0)
}

// Fee returns the fee of a Tx
func (tx Tx) Fee() xc.AmountBlockchain {
	switch tf := tx.CosmosTx.(type) {
	case types.FeeTx:
		fee := tf.GetFee()[0].Amount.BigInt()
		return xc.AmountBlockchain(*fee)
	}
	return xc.NewAmountBlockchainFromUint64(0)
}
