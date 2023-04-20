package bitcoin

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"

	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
	xc "github.com/jumpcrypto/crosschain"
	log "github.com/sirupsen/logrus"
)

// A specific output from a transaction
type Outpoint struct {
	Hash  []byte `json:"hash"`
	Index uint32 `json:"index"`
}

type Output struct {
	Outpoint     `json:"outpoint"`
	Value        xc.AmountBlockchain `json:"value"`
	PubKeyScript []byte              `json:"pubKeyScript"`
}

type Input struct {
	Output    `json:"output"`
	SigScript []byte     `json:"sigScript,omitempty"`
	Address   xc.Address `json:"address,omitempty"`
}

type Recipient struct {
	To    xc.Address          `json:"to"`
	Value xc.AmountBlockchain `json:"value"`
}

// Tx for Bitcoin
type Tx struct {
	msgTx      *wire.MsgTx
	signed     bool
	recipients []Recipient

	amount xc.AmountBlockchain
	input  TxInput
	from   xc.Address
	to     xc.Address
	isBch  bool
}

var _ xc.Tx = &Tx{}

// Hash returns the tx hash or id
func (tx *Tx) Hash() xc.TxHash {
	return tx.txHashReversed()
}

func (tx *Tx) txHashReversed() xc.TxHash {
	txHash := tx.txHashNormalBytes()

	size := len(txHash)
	txHashReversed := make([]byte, size)
	copy(txHashReversed[:], txHash[:])
	for i := 0; i < size/2; i++ {
		txHashReversed[i], txHashReversed[size-1-i] = txHashReversed[size-1-i], txHashReversed[i]
	}
	return xc.TxHash(hex.EncodeToString(txHashReversed))
}
func (tx *Tx) txHashNormal() xc.TxHash {
	txhash := tx.txHashNormalBytes()
	return xc.TxHash(hex.EncodeToString(txhash[:]))
}
func (tx *Tx) txHashNormalBytes() []byte {
	txhash := tx.msgTx.TxHash()
	return txhash[:]
}

func bzToString(bz []byte) string {
	return base64.RawStdEncoding.EncodeToString(bz)
}

// Sighashes returns the tx payload to sign, aka sighash
func (tx *Tx) Sighashes() ([]xc.TxDataToSign, error) {
	sighashes := make([]xc.TxDataToSign, len(tx.input.Inputs))

	for i, txin := range tx.input.Inputs {
		pubKeyScript := txin.PubKeyScript
		sigScript := txin.SigScript
		value := txin.Value.Uint64()

		var hash []byte
		var err error
		log.Debugf("Sighashes params: sigScript=%s IsPayToWitnessPubKeyHash(pubKeyScript)=%t", bzToString(sigScript), txscript.IsPayToWitnessPubKeyHash(pubKeyScript))
		if tx.isBch {
			if sigScript == nil {
				hash = CalculateBchBip143Sighash(pubKeyScript, txscript.NewTxSigHashes(tx.msgTx), txscript.SigHashAll, tx.msgTx, i, int64(value))
			} else {
				hash = CalculateBchBip143Sighash(sigScript, txscript.NewTxSigHashes(tx.msgTx), txscript.SigHashAll, tx.msgTx, i, int64(value))
			}
		} else {
			if sigScript == nil {
				if txscript.IsPayToWitnessPubKeyHash(pubKeyScript) {
					log.Debugf("CalcWitnessSigHash with pubKeyScript: %s", base64.RawURLEncoding.EncodeToString(pubKeyScript))
					hash, err = txscript.CalcWitnessSigHash(pubKeyScript, txscript.NewTxSigHashes(tx.msgTx), txscript.SigHashAll, tx.msgTx, i, int64(value))
				} else {
					log.Debugf("CalcSignatureHash with pubKeyScript: %s", base64.RawURLEncoding.EncodeToString(pubKeyScript))
					hash, err = txscript.CalcSignatureHash(pubKeyScript, txscript.SigHashAll, tx.msgTx, i)
				}
			} else {
				if txscript.IsPayToWitnessScriptHash(pubKeyScript) {
					log.Debugf("CalcWitnessSigHash with sigScript: %s", base64.RawURLEncoding.EncodeToString(sigScript))
					hash, err = txscript.CalcWitnessSigHash(sigScript, txscript.NewTxSigHashes(tx.msgTx), txscript.SigHashAll, tx.msgTx, i, int64(value))
				} else {
					log.Debugf("CalcSignatureHash with sigScript: %s", base64.RawURLEncoding.EncodeToString(sigScript))
					hash, err = txscript.CalcSignatureHash(sigScript, txscript.SigHashAll, tx.msgTx, i)
				}
			}
		}
		if err != nil {
			return []xc.TxDataToSign{}, err
		}

		sighashes[i] = hash
	}

	return sighashes, nil
}

// AddSignatures adds a signature to Tx
func (tx *Tx) AddSignatures(signatures ...xc.TxSignature) error {
	if tx.signed {
		return fmt.Errorf("already signed")
	}
	if len(signatures) != len(tx.msgTx.TxIn) {
		return fmt.Errorf("expected %v signatures, got %v signatures", len(tx.msgTx.TxIn), len(signatures))
	}

	for i, rsvBytes := range signatures {
		var err error
		rsv := [65]byte{}
		if len(rsvBytes) != 65 && len(rsvBytes) != 64 {
			return errors.New("signature must be 64 or 65 length serialized bytestring of r,s, and recovery byte")
		}
		copy(rsv[:], rsvBytes)

		// Decode the signature and the pubkey script.
		r := new(big.Int).SetBytes(rsv[:32])
		s := new(big.Int).SetBytes(rsv[32:64])
		signature := btcec.Signature{
			R: r,
			S: s,
		}
		pubKeyScript := tx.input.Inputs[i].Output.PubKeyScript
		sigScript := tx.input.Inputs[i].SigScript

		// Support segwit.
		if sigScript == nil {
			if txscript.IsPayToWitnessPubKeyHash(pubKeyScript) || txscript.IsPayToWitnessScriptHash(pubKeyScript) {
				log.Debug("append signature (segwit)")
				tx.msgTx.TxIn[i].Witness = wire.TxWitness([][]byte{append(signature.Serialize(), byte(txscript.SigHashAll)), tx.input.FromPublicKey})
				continue
			}
		} else {
			if txscript.IsPayToWitnessScriptHash(sigScript) {
				log.Debug("append signature + sigscript (segwit)")
				tx.msgTx.TxIn[i].Witness = wire.TxWitness([][]byte{append(signature.Serialize(), byte(txscript.SigHashAll)), tx.input.FromPublicKey, sigScript})
				continue
			}
		}

		// Support non-segwit
		builder := txscript.NewScriptBuilder()
		sigHashByte := txscript.SigHashAll
		if tx.isBch {
			sigHashByte = sigHashByte | SighashForkID
		}
		builder.AddData(append(signature.Serialize(), byte(sigHashByte)))
		builder.AddData(tx.input.FromPublicKey)
		log.Debug("append signature (non-segwit)")
		if sigScript != nil {
			log.Debug("append sigScript (non-segwit)")
			builder.AddData(sigScript)
		}
		tx.msgTx.TxIn[i].SignatureScript, err = builder.Script()
		if err != nil {
			return err
		}
	}

	tx.signed = true
	return nil
}

func (tx *Tx) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := tx.msgTx.Serialize(buf); err != nil {
		return []byte{}, err
	}
	return buf.Bytes(), nil
}

// Outputs returns the UTXO outputs in the underlying transaction.
func (tx *Tx) Outputs() ([]Output, error) {
	hash := tx.txHashNormal()
	outputs := make([]Output, len(tx.msgTx.TxOut))
	for i := range outputs {
		outputs[i].Outpoint = Outpoint{
			Hash:  []byte(hash),
			Index: uint32(i),
		}
		outputs[i].PubKeyScript = tx.msgTx.TxOut[i].PkScript
		if tx.msgTx.TxOut[i].Value < 0 {
			return nil, fmt.Errorf("bad output %v: value is less than zero", i)
		}
		outputs[i].Value = xc.NewAmountBlockchainFromUint64(uint64(tx.msgTx.TxOut[i].Value))
	}
	return outputs, nil
}

// Heuristic to determine the sender of a transaction by
// using the largest utxo input and taking it's spender.
func (tx *Tx) DetectFrom() (string, xc.AmountBlockchain) {
	from := ""
	max := xc.NewAmountBlockchainFromUint64(0)
	totalIn := xc.NewAmountBlockchainFromUint64(0)
	for _, input := range tx.input.Inputs {
		value := input.Output.Value
		if value.Cmp(&max) > 0 {
			max = value
			from = string(input.Address)
		}
		fmt.Println("inputfrom: ", input.Address)
		totalIn = totalIn.Add(&value)
	}
	return from, totalIn
}

func (tx *Tx) DetectToAndAmount(from string, expectedTo string) (string, xc.AmountBlockchain, xc.AmountBlockchain) {
	to := expectedTo
	amount := xc.NewAmountBlockchainFromUint64(0)
	totalOut := xc.NewAmountBlockchainFromUint64(0)

	for _, recipient := range tx.recipients {
		addr := string(recipient.To)
		value := recipient.Value

		// if we know "to", we add the value(s)
		if expectedTo != "" && addr == expectedTo {
			amount = amount.Add(&value)
		}

		// if we don't know "to", we set "to" as anything different than "from"
		if expectedTo == "" && addr != from {
			amount = value
			to = addr
		}
		fmt.Println("recipient to: ", recipient.To)

		totalOut = totalOut.Add(&value)
	}
	return to, amount, totalOut
}
func (tx *Tx) IsBch() bool {
	return tx.isBch
}
