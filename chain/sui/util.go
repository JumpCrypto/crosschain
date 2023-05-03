package sui

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"

	"github.com/btcsuite/btcutil/base58"
	"github.com/jumpcrypto/crosschain/chain/sui/generated/bcs"
)

type ObjectRef struct {
	Field0 bcs.ObjectID
	Field1 bcs.SequenceNumber
	Field2 bcs.ObjectDigest
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
	bytes, err := hex.DecodeString(str)
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

func hexToPure(str string) (*bcs.CallArg__Pure, error) {
	bytes, err := hex.DecodeString(str)
	if err != nil {
		return &bcs.CallArg__Pure{}, err
	}
	pure := bcs.CallArg__Pure(bytes)
	return &pure, nil
}

func toAddress(slice []byte) bcs.SuiAddress {
	var array [32]byte
	for i := 0; i < len(slice); i += 1 {
		array[i] = slice[i]
	}
	return array
}

func hexToAddress(str string) (bcs.SuiAddress, error) {
	bytes, err := hex.DecodeString(str)
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
