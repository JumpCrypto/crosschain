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
	CosmosTx        types.Tx
	ParsedTransfers []types.Msg
	// aux fields
	CosmosTxBuilder client.TxBuilder
	CosmosTxEncoder types.TxEncoder
	SigsV2          []signingtypes.SignatureV2
	TxDataToSign    []byte
}

var _ xc.Tx = Tx{}

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

func reserializeSig(signature []byte) []byte {
	if len(signature) <= 64 {
		return signature
	}
	// this, e.g., drops the recovery bit that's typical in bitcoin/evm signatures
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
			tx.ParsedTransfers = append(tx.ParsedTransfers, msg)
		}
	}
}

// From returns the from address of a Tx
func (tx Tx) From() xc.Address {
	for _, parsedTransfer := range tx.ParsedTransfers {
		switch tf := parsedTransfer.(type) {
		case *banktypes.MsgSend:
			from := tf.FromAddress
			return xc.Address(from)
		}
	}
	return xc.Address("")
}

// To returns the to address of a Tx
func (tx Tx) To() xc.Address {
	for _, parsedTransfer := range tx.ParsedTransfers {
		switch tf := parsedTransfer.(type) {
		case *banktypes.MsgSend:
			to := tf.ToAddress
			return xc.Address(to)
		}
	}
	return xc.Address("")
}

// ContractAddress returns the contract address of a Tx, if any
func (tx Tx) ContractAddress() xc.ContractAddress {
	for _, parsedTransfer := range tx.ParsedTransfers {
		switch tf := parsedTransfer.(type) {
		case *banktypes.MsgSend:
			denom := tf.Amount[0].Denom
			// remove native assets to be coherent with other chains
			if len(denom) < LEN_NATIVE_ASSET {
				denom = ""
			}
			return xc.ContractAddress(denom)
		}
	}
	return xc.ContractAddress("")
}

// Amount returns the amount of a Tx
func (tx Tx) Amount() xc.AmountBlockchain {
	for _, parsedTransfer := range tx.ParsedTransfers {
		switch tf := parsedTransfer.(type) {
		case *banktypes.MsgSend:
			amount := tf.Amount[0].Amount.BigInt()
			return xc.AmountBlockchain(*amount)
		}
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

// Sources returns the sources of a Tx
func (tx Tx) Sources() []*xc.TxInfoEndpoint {
	sources := []*xc.TxInfoEndpoint{}
	for _, parsedTransfer := range tx.ParsedTransfers {
		switch tf := parsedTransfer.(type) {
		case *banktypes.MsgSend:
			from := tf.FromAddress
			sources = append(sources, &xc.TxInfoEndpoint{
				Address: xc.Address(from),
			})
			// currently assume/support single-source transfers
			return sources
		}
	}
	return sources
}

// Destinations returns the destinations of a Tx
func (tx Tx) Destinations() []*xc.TxInfoEndpoint {
	destinations := []*xc.TxInfoEndpoint{}
	for _, parsedTransfer := range tx.ParsedTransfers {
		switch tf := parsedTransfer.(type) {
		case *banktypes.MsgSend:
			to := tf.ToAddress
			denom := tf.Amount[0].Denom
			amount := tf.Amount[0].Amount.BigInt()
			destinations = append(destinations, &xc.TxInfoEndpoint{
				Address:         xc.Address(to),
				ContractAddress: xc.ContractAddress(denom),
				Amount:          xc.AmountBlockchain(*amount),
			})
		}
	}
	return destinations
}
