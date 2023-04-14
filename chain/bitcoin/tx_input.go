package bitcoin

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"sort"

	xc "github.com/jumpcrypto/crosschain"
	log "github.com/sirupsen/logrus"
)

// TxInput for Bitcoin
type TxInput struct {
	unspentOutputs  []Output
	inputs          []Input
	FromPublicKey   []byte
	gasPricePerByte xc.AmountBlockchain
}

var _ xc.TxInputWithPublicKey = &TxInput{}

func (txInput *TxInput) SetPublicKey(publicKeyBytes xc.PublicKey) error {
	txInput.FromPublicKey = publicKeyBytes
	return nil
}

func (txInput *TxInput) SetPublicKeyFromStr(publicKeyStr string) error {
	publicKeyBytes, err := base64.StdEncoding.DecodeString(publicKeyStr)
	if err != nil {
		return fmt.Errorf("invalid public key %v: %v", publicKeyStr, err)
	}
	err = txInput.SetPublicKey(publicKeyBytes)

	return err
}

// 1. sort unspentOutputs from lowest to highest
// 2. grab the minimum amount of UTXO needed to satify amount
// 3. tack on the smallest utxo's until `minUtxo` is reached.
// This ensures a small number of UTXO are used for transaction while also consolidating some
// smaller utxo into the transaction.
// Returns the total balance of the min utxo set.  txInput.inputs are updated to the new set.
func (txInput *TxInput) allocateMinUtxoSet(targetAmount xc.AmountBlockchain, minUtxo int) *xc.AmountBlockchain {
	balance := xc.NewAmountBlockchainFromUint64(0)

	// 1. sort from lowest to higher
	if len(txInput.unspentOutputs) > 1 {
		sort.Slice(txInput.unspentOutputs, func(i, j int) bool {
			return txInput.unspentOutputs[i].Value.Cmp(&txInput.unspentOutputs[j].Value) <= 0
		})
	}

	inputs := []Input{}
	lenUTXOIndex := len(txInput.unspentOutputs) - 1
	for balance.Cmp(&targetAmount) < 0 && lenUTXOIndex >= 0 {
		o := txInput.unspentOutputs[lenUTXOIndex]
		log.Infof("unspent output h2l: %s (%s)", hex.EncodeToString(o.PubKeyScript), o.Value.String())
		balance = balance.Add(&o.Value)

		inputs = append(inputs, Input{
			Output: o,
		})
		lenUTXOIndex--
	}

	// add the smallest utxo until we reach `minUtxo` inputs
	// lenUTXOIndex wasn't used, so i can grow up to lenUTXOIndex (included)
	i := 0
	for len(inputs) < minUtxo && i < lenUTXOIndex {
		o := txInput.unspentOutputs[i]
		log.Infof("unspent output l2h: %s (%s)", hex.EncodeToString(o.PubKeyScript), o.Value.String())
		balance = balance.Add(&o.Value)
		inputs = append(inputs, Input{
			Output: o,
		})
		i++
	}
	txInput.inputs = inputs
	return &balance
}
