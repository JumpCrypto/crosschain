package sui

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/btcsuite/btcutil/base58"
	"github.com/coming-chat/go-sui/types"
	"github.com/jumpcrypto/crosschain/chain/sui/generated/bcs"
)

type ObjectRef struct {
	Field0 bcs.ObjectID
	Field1 bcs.SequenceNumber
	Field2 bcs.ObjectDigest
}

func decodeHex(str string) ([]byte, error) {
	if strings.HasPrefix(str, "0x") {
		return hex.DecodeString(str[2:])
	}
	return hex.DecodeString(str)
}

func toObjectID(slice []byte) (bcs.ObjectID, error) {
	var array [32]byte
	if n := copy(array[:], slice); n != 32 {
		return bcs.ObjectID{}, fmt.Errorf("ObjectID must have 32 bytes: %v", slice)
	}
	return bcs.ObjectID{
		Value: array,
	}, nil
}

func hexToObjectID(str string) (bcs.ObjectID, error) {
	bytes, err := decodeHex(str)
	if err != nil {
		return bcs.ObjectID{}, err
	}

	return toObjectID(bytes)
}

func u64ToPure(x uint64) *bcs.CallArg__Pure {
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, x)
	pure := bcs.CallArg__Pure(bytes)
	return &pure
}

func coinToObject(coin *types.Coin) (*bcs.ObjectArg__ImmOrOwnedObject, error) {
	id, err := hexToObjectID(coin.CoinObjectId.String())
	if err != nil {
		return &bcs.ObjectArg__ImmOrOwnedObject{}, fmt.Errorf("could not decode coin id: %v", err)
	}
	digest, err := base58ToObjectDigest(coin.Digest)
	if err != nil {
		return &bcs.ObjectArg__ImmOrOwnedObject{}, fmt.Errorf("could not decode coin digest: %v", err)
	}
	seq := coin.Version.BigInt().Uint64()

	return &bcs.ObjectArg__ImmOrOwnedObject{
		Field0: id,
		Field1: bcs.SequenceNumber(seq),
		Field2: digest,
	}, nil
}
func mustCoinToObject(coin *types.Coin) *bcs.ObjectArg__ImmOrOwnedObject {
	obj, err := coinToObject(coin)
	if err != nil {
		panic(err)
	}
	return obj
}
func hexToPure(str string) (*bcs.CallArg__Pure, error) {
	bytes, err := decodeHex(str)
	if err != nil {
		return &bcs.CallArg__Pure{}, err
	}
	pure := bcs.CallArg__Pure(bytes)
	return &pure, nil
}

func mustHexToPure(str string) *bcs.CallArg__Pure {
	data, err := hexToPure(str)
	if err != nil {
		panic(err)
	}
	return data
}

func toAddress(slice []byte) bcs.SuiAddress {
	var array [32]byte
	for i := 0; i < len(slice); i += 1 {
		array[i] = slice[i]
	}
	return array
}

func hexToAddress(str string) (bcs.SuiAddress, error) {
	bytes, err := decodeHex(str)
	if err != nil {
		return bcs.SuiAddress{}, err
	}
	return toAddress(bytes), nil
}

func base58ToBytes(str string) ([]byte, error) {
	bytes := base58.Decode(str)
	if len(bytes) == 0 {
		return bytes, fmt.Errorf("failed to decode base58 string %q", str)
	}
	return bytes, nil
}

func base58ToObjectDigest(str string) (bcs.ObjectDigest, error) {
	bytes, err := base58ToBytes(str)
	if err != nil {
		return bcs.ObjectDigest{}, err
	}
	return bcs.ObjectDigest{
		Value: bytes,
	}, nil
}

func ArgumentInput(index uint16) *bcs.Argument__Input {
	x := bcs.Argument__Input(index)
	return &x
}

func ArgumentResult(index uint16) *bcs.Argument__Result {
	x := bcs.Argument__Result(index)
	return &x
}

// Strip the coin::Coin<_> wrapper if present
func NormalizeCoinContract(contract string) string {
	if strings.HasPrefix(contract, "coin::Coin<") {
		contract = strings.Replace(contract, "coin::Coin<", "", 1)
		contract = strings.Replace(contract, ">", "", 1)
	}
	return contract
}
