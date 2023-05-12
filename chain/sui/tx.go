package sui

import (
	"encoding/base64"
	"encoding/hex"
	"sort"

	"github.com/btcsuite/btcutil/base58"
	"github.com/coming-chat/go-sui/types"
	xc "github.com/jumpcrypto/crosschain"
	"github.com/jumpcrypto/crosschain/chain/sui/generated/bcs"
	"golang.org/x/crypto/blake2b"
)

// Tx for Template
type TxInput struct {
	xc.TxInputEnvelope
	GasBudget uint64
	GasPrice  uint64
	Pubkey    []byte
	// Native Sui object that we can use to pay gas with
	GasCoin types.Coin
	// All objects (native or token)
	Coins []*types.Coin
	// current epoch
	CurrentEpoch uint64
}

var _ xc.TxInputWithPublicKey = &TxInput{}

func (input *TxInput) SetPublicKey(pubkey xc.PublicKey) error {
	input.Pubkey = pubkey
	return nil
}

func (input *TxInput) SetPublicKeyFromStr(pubkeyStr string) error {
	var err error
	var pubkey []byte
	if len(pubkeyStr) == 64 || len(pubkeyStr) == 128 {
		pubkey, err = hex.DecodeString(pubkeyStr)
		if err != nil {
			return err
		}
	} else {
		pubkey, err = base64.StdEncoding.DecodeString(pubkeyStr)
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

func (input *TxInput) ExcludeGasCoin() {
	for i, coin := range input.Coins {
		if coin.CoinObjectId.String() == input.GasCoin.CoinObjectId.String() {
			// drop it
			input.Coins = append(input.Coins[:i], input.Coins[i+1:]...)
			break
		}
	}
}

func (input *TxInput) TotalBalance() xc.AmountBlockchain {
	amount := xc.NewAmountBlockchainFromUint64(0)
	coinType := ""
	for _, coin := range input.Coins {
		coinType = coin.CoinType
		coinBal := xc.NewAmountBlockchainFromUint64(coin.Balance.Uint64())
		amount = amount.Add(&coinBal)
	}
	// add gas coin if it's same type
	if coinType == "" || coinType == input.GasCoin.CoinType {
		coinBal := xc.NewAmountBlockchainFromUint64(input.GasCoin.Balance.Uint64())
		amount = amount.Add(&coinBal)
	}
	return amount
}

// Sort coins in place from highest to lowest
func (input *TxInput) SortCoins() {
	SortCoins(input.Coins)
}

func (input *TxInput) IsNativeTransfer() bool {
	if len(input.Coins) > 0 && input.Coins[0].CoinType != input.GasCoin.CoinType {
		return false
	}
	return true
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
	// Input      TxInput
	signatures [][]byte
	public_key []byte
	Tx         bcs.TransactionData__V1
}

var _ xc.Tx = &Tx{}
var _ SignatureGetter = &Tx{}

// Hash returns the tx hash or id
func (tx Tx) Hash() xc.TxHash {
	typeTag := "TransactionData::"
	bz, err := tx.Serialize()
	if err != nil {
		panic(err)
	}
	tohash := append([]byte(typeTag), bz...)
	hash := blake2b.Sum256(tohash)
	hash_b58 := base58.Encode(hash[:])
	return xc.TxHash(hash_b58)
}
func (tx Tx) Sighashes() ([]xc.TxDataToSign, error) {
	bytes, err := tx.Serialize()
	// 0 = transaction data, 0 = V0 intent version, 0 = sui
	// https://github.com/MystenLabs/sui/blob/a78b9e3f8a212924848f540da5a2587526525853/sdk/typescript/src/utils/intent.ts#L26
	intent := []byte{0, 0, 0}
	msg := append(intent, bytes...)
	hash := blake2b.Sum256(msg)

	if err != nil {
		return []xc.TxDataToSign{}, err
	}
	return []xc.TxDataToSign{hash[:]}, nil
}
func (tx *Tx) AddSignatures(signatures ...xc.TxSignature) error {
	for _, sig := range signatures {
		// sui expects signature to be {0, signature, public_key}
		sui_sig := []byte{0}
		sui_sig = append(sui_sig, sig...)
		sui_sig = append(sui_sig, tx.public_key...)
		tx.signatures = append(tx.signatures, sui_sig)
	}
	return nil
}
func (tx Tx) Serialize() ([]byte, error) {
	bytes, err := tx.Tx.BcsSerialize()
	if err != nil {
		return bytes, err
	}
	return bytes, nil
}
func (tx Tx) GetSignatures() [][]byte {
	return tx.signatures
}
