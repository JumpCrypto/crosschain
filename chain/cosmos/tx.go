package cosmos

import (
	"encoding/hex"
	"errors"
	"math/big"

	"github.com/btcsuite/btcd/btcec"
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
	serialized, err := tx.Serialize()
	if err != nil || serialized == nil || len(serialized) == 0 {
		return ""
	}
	txID := tmhash.Sum(serialized)
	return xc.TxHash(hex.EncodeToString(txID))
}

// Sighashes returns the tx payload to sign, aka sighash
func (tx Tx) Sighashes() ([]xc.TxDataToSign, error) {
	if tx.TxDataToSign == nil {
		return nil, errors.New("transaction not initialized")
	}
	return []xc.TxDataToSign{tx.TxDataToSign}, nil
}

func signatureFromBytes(sigStr []byte) *btcec.Signature {
	return &btcec.Signature{
		R: new(big.Int).SetBytes(sigStr[:32]),
		S: new(big.Int).SetBytes(sigStr[32:64]),
	}
}

// Serialize signature to R || S.
// R, S are padded to 32 bytes respectively.
func serializeSig(sig *btcec.Signature) []byte {
	rBytes := sig.R.Bytes()
	sBytes := sig.S.Bytes()
	sigBytes := make([]byte, 64)
	// 0 pad the byte arrays from the left if they aren't big enough.
	copy(sigBytes[32-len(rBytes):32], rBytes)
	copy(sigBytes[64-len(sBytes):64], sBytes)
	return sigBytes
}

func reserializeSig(signature []byte) []byte {
	return serializeSig(signatureFromBytes(signature))
}

// AddSignatures adds a signature to Tx
func (tx Tx) AddSignatures(signatures ...xc.TxSignature) error {
	if tx.SigsV2 == nil || len(tx.SigsV2) < 1 || tx.CosmosTxBuilder == nil {
		return errors.New("transaction not initialized")
	}
	if len(signatures) != len(tx.SigsV2) {
		return errors.New("invalid signatures size")
	}
	for i, signature := range signatures {
		data := tx.SigsV2[i].Data
		signMode := data.(*signingtypes.SingleSignatureData).SignMode
		tx.SigsV2[i].Data = &signingtypes.SingleSignatureData{
			SignMode:  signMode,
			Signature: reserializeSig(signature),
		}
	}
	return tx.CosmosTxBuilder.SetSignatures(tx.SigsV2...)
}

// Serialize serializes a Tx
func (tx Tx) Serialize() ([]byte, error) {
	if tx.CosmosTxEncoder == nil {
		return []byte{}, errors.New("transaction not initialized")
	}

	// if CosmosTxBuilder is set, prioritize GetTx()
	txToEncode := tx.CosmosTx
	if tx.CosmosTxBuilder != nil {
		txToEncode = tx.CosmosTxBuilder.GetTx()
	}

	if txToEncode == nil {
		return []byte{}, errors.New("transaction not initialized")
	}
	serialized, err := tx.CosmosTxEncoder(txToEncode)
	return serialized, err
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
