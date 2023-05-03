package sui

import (
	"encoding/base64"
	"encoding/hex"
	"sort"

	"github.com/coming-chat/go-sui/types"
	xc "github.com/jumpcrypto/crosschain"
	"github.com/jumpcrypto/crosschain/chain/sui/generated/bcs"
)

// Tx for Template
type TxInput struct {
	xc.TxInputEnvelope
	GasBudget uint64
	GasPrice  uint64
	ChainId   int
	Pubkey    []byte
	// Native Sui object that we can use to pay gas with
	GasCoin types.Coin
	// All objects (native or token)
	Coins []*types.Coin
}

var _ xc.TxInputWithPublicKey = &TxInput{}

func (input *TxInput) SetPublicKey(pubkey xc.PublicKey) error {
	input.Pubkey = pubkey
	return nil
}

func (input *TxInput) SetPublicKeyFromStr(pubkeyStr string) error {
	var err error
	var pubkey []byte
	if len(pubkeyStr) == 128 || len(pubkeyStr) == 130 {
		pubkey, err = hex.DecodeString(pubkeyStr)
		if err != nil {
			return err
		}
	} else {
		pubkey, err = base64.RawStdEncoding.DecodeString(pubkeyStr)
		if err != nil {
			return err
		}
	}
	input.Pubkey = pubkey
	return nil
}

func SortCoins(coins []*types.Coin) {
	sort.Slice(coins, func(i, j int) bool {
		return coins[i].Balance.Decimal().Cmp(coins[j].Balance.Decimal()) > 0
	})
}

// Sort coins in place from highest to lowest
func (input *TxInput) SortCoins() {
	SortCoins(input.Coins)
}

func NewTxInput() *TxInput {
	return &TxInput{
		TxInputEnvelope: xc.TxInputEnvelope{
			Type: xc.DriverSui,
		},
	}
}

type SignatureGetter interface {
	GetSignatures() [][]byte
}
type Tx struct {
	Input      TxInput
	signatures [][]byte
	tx         bcs.TransactionData__V1
}

var _ xc.Tx = &Tx{}
var _ SignatureGetter = &Tx{}

// Hash returns the tx hash or id
func (tx Tx) Hash() xc.TxHash {
	panic("not implemented")
}
func (tx Tx) Sighashes() ([]xc.TxDataToSign, error) {
	panic("not implemented")
}
func (tx *Tx) AddSignatures(signatures ...xc.TxSignature) error {
	panic("not implemented")
}
func (tx Tx) Serialize() ([]byte, error) {
	panic("not implemented")
}
func (tx Tx) GetSignatures() [][]byte {
	return tx.signatures
}
