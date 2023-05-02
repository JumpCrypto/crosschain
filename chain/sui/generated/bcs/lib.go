package bcs


import (
	"fmt"
	"github.com/novifinancial/serde-reflection/serde-generate/runtime/golang/serde"
	"github.com/novifinancial/serde-reflection/serde-generate/runtime/golang/bcs"
)


type AbortLocation interface {
	isAbortLocation()
	Serialize(serializer serde.Serializer) error
	BcsSerialize() ([]byte, error)
}

func DeserializeAbortLocation(deserializer serde.Deserializer) (AbortLocation, error) {
	index, err := deserializer.DeserializeVariantIndex()
	if err != nil { return nil, err }

	switch index {
	case 0:
		if val, err := load_AbortLocation__Module(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 1:
		if val, err := load_AbortLocation__Script(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	default:
		return nil, fmt.Errorf("Unknown variant index for AbortLocation: %d", index)
	}
}

func BcsDeserializeAbortLocation(input []byte) (AbortLocation, error) {
	if input == nil {
		var obj AbortLocation
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input);
	obj, err := DeserializeAbortLocation(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type AbortLocation__Module struct {
	Value ModuleId
}

func (*AbortLocation__Module) isAbortLocation() {}

func (obj *AbortLocation__Module) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(0)
	if err := obj.Value.Serialize(serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *AbortLocation__Module) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_AbortLocation__Module(deserializer serde.Deserializer) (AbortLocation__Module, error) {
	var obj AbortLocation__Module
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := DeserializeModuleId(deserializer); err == nil { obj.Value = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type AbortLocation__Script struct {
}

func (*AbortLocation__Script) isAbortLocation() {}

func (obj *AbortLocation__Script) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(1)
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *AbortLocation__Script) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_AbortLocation__Script(deserializer serde.Deserializer) (AbortLocation__Script, error) {
	var obj AbortLocation__Script
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type AccountAddress [32]uint8

func (obj *AccountAddress) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	if err := serialize_array32_u8_array((([32]uint8)(*obj)), serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *AccountAddress) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func DeserializeAccountAddress(deserializer serde.Deserializer) (AccountAddress, error) {
	var obj [32]uint8
	if err := deserializer.IncreaseContainerDepth(); err != nil { return (AccountAddress)(obj), err }
	if val, err := deserialize_array32_u8_array(deserializer); err == nil { obj = val } else { return ((AccountAddress)(obj)), err }
	deserializer.DecreaseContainerDepth()
	return (AccountAddress)(obj), nil
}

func BcsDeserializeAccountAddress(input []byte) (AccountAddress, error) {
	if input == nil {
		var obj AccountAddress
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input);
	obj, err := DeserializeAccountAddress(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type Argument interface {
	isArgument()
	Serialize(serializer serde.Serializer) error
	BcsSerialize() ([]byte, error)
}

func DeserializeArgument(deserializer serde.Deserializer) (Argument, error) {
	index, err := deserializer.DeserializeVariantIndex()
	if err != nil { return nil, err }

	switch index {
	case 0:
		if val, err := load_Argument__GasCoin(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 1:
		if val, err := load_Argument__Input(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 2:
		if val, err := load_Argument__Result(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 3:
		if val, err := load_Argument__NestedResult(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	default:
		return nil, fmt.Errorf("Unknown variant index for Argument: %d", index)
	}
}

func BcsDeserializeArgument(input []byte) (Argument, error) {
	if input == nil {
		var obj Argument
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input);
	obj, err := DeserializeArgument(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type Argument__GasCoin struct {
}

func (*Argument__GasCoin) isArgument() {}

func (obj *Argument__GasCoin) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(0)
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *Argument__GasCoin) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_Argument__GasCoin(deserializer serde.Deserializer) (Argument__GasCoin, error) {
	var obj Argument__GasCoin
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type Argument__Input uint16

func (*Argument__Input) isArgument() {}

func (obj *Argument__Input) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(1)
	if err := serializer.SerializeU16(((uint16)(*obj))); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *Argument__Input) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_Argument__Input(deserializer serde.Deserializer) (Argument__Input, error) {
	var obj uint16
	if err := deserializer.IncreaseContainerDepth(); err != nil { return (Argument__Input)(obj), err }
	if val, err := deserializer.DeserializeU16(); err == nil { obj = val } else { return ((Argument__Input)(obj)), err }
	deserializer.DecreaseContainerDepth()
	return (Argument__Input)(obj), nil
}

type Argument__Result uint16

func (*Argument__Result) isArgument() {}

func (obj *Argument__Result) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(2)
	if err := serializer.SerializeU16(((uint16)(*obj))); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *Argument__Result) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_Argument__Result(deserializer serde.Deserializer) (Argument__Result, error) {
	var obj uint16
	if err := deserializer.IncreaseContainerDepth(); err != nil { return (Argument__Result)(obj), err }
	if val, err := deserializer.DeserializeU16(); err == nil { obj = val } else { return ((Argument__Result)(obj)), err }
	deserializer.DecreaseContainerDepth()
	return (Argument__Result)(obj), nil
}

type Argument__NestedResult struct {
	Field0 uint16
	Field1 uint16
}

func (*Argument__NestedResult) isArgument() {}

func (obj *Argument__NestedResult) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(3)
	if err := serializer.SerializeU16(obj.Field0); err != nil { return err }
	if err := serializer.SerializeU16(obj.Field1); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *Argument__NestedResult) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_Argument__NestedResult(deserializer serde.Deserializer) (Argument__NestedResult, error) {
	var obj Argument__NestedResult
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := deserializer.DeserializeU16(); err == nil { obj.Field0 = val } else { return obj, err }
	if val, err := deserializer.DeserializeU16(); err == nil { obj.Field1 = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type AuthorityPublicKeyBytes []byte

func (obj *AuthorityPublicKeyBytes) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	if err := serializer.SerializeBytes((([]byte)(*obj))); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *AuthorityPublicKeyBytes) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func DeserializeAuthorityPublicKeyBytes(deserializer serde.Deserializer) (AuthorityPublicKeyBytes, error) {
	var obj []byte
	if err := deserializer.IncreaseContainerDepth(); err != nil { return (AuthorityPublicKeyBytes)(obj), err }
	if val, err := deserializer.DeserializeBytes(); err == nil { obj = val } else { return ((AuthorityPublicKeyBytes)(obj)), err }
	deserializer.DecreaseContainerDepth()
	return (AuthorityPublicKeyBytes)(obj), nil
}

func BcsDeserializeAuthorityPublicKeyBytes(input []byte) (AuthorityPublicKeyBytes, error) {
	if input == nil {
		var obj AuthorityPublicKeyBytes
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input);
	obj, err := DeserializeAuthorityPublicKeyBytes(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type CallArg interface {
	isCallArg()
	Serialize(serializer serde.Serializer) error
	BcsSerialize() ([]byte, error)
}

func DeserializeCallArg(deserializer serde.Deserializer) (CallArg, error) {
	index, err := deserializer.DeserializeVariantIndex()
	if err != nil { return nil, err }

	switch index {
	case 0:
		if val, err := load_CallArg__Pure(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 1:
		if val, err := load_CallArg__Object(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	default:
		return nil, fmt.Errorf("Unknown variant index for CallArg: %d", index)
	}
}

func BcsDeserializeCallArg(input []byte) (CallArg, error) {
	if input == nil {
		var obj CallArg
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input);
	obj, err := DeserializeCallArg(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type CallArg__Pure []uint8

func (*CallArg__Pure) isCallArg() {}

func (obj *CallArg__Pure) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(0)
	if err := serialize_vector_u8((([]uint8)(*obj)), serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *CallArg__Pure) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_CallArg__Pure(deserializer serde.Deserializer) (CallArg__Pure, error) {
	var obj []uint8
	if err := deserializer.IncreaseContainerDepth(); err != nil { return (CallArg__Pure)(obj), err }
	if val, err := deserialize_vector_u8(deserializer); err == nil { obj = val } else { return ((CallArg__Pure)(obj)), err }
	deserializer.DecreaseContainerDepth()
	return (CallArg__Pure)(obj), nil
}

type CallArg__Object struct {
	Value ObjectArg
}

func (*CallArg__Object) isCallArg() {}

func (obj *CallArg__Object) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(1)
	if err := obj.Value.Serialize(serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *CallArg__Object) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_CallArg__Object(deserializer serde.Deserializer) (CallArg__Object, error) {
	var obj CallArg__Object
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := DeserializeObjectArg(deserializer); err == nil { obj.Value = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type ChangeEpoch struct {
	Epoch uint64
	ProtocolVersion ProtocolVersion
	StorageCharge uint64
	ComputationCharge uint64
	StorageRebate uint64
	NonRefundableStorageFee uint64
	EpochStartTimestampMs uint64
	SystemPackages []struct {Field0 SequenceNumber; Field1 [][]uint8; Field2 []ObjectID}
}

func (obj *ChangeEpoch) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	if err := serializer.SerializeU64(obj.Epoch); err != nil { return err }
	if err := obj.ProtocolVersion.Serialize(serializer); err != nil { return err }
	if err := serializer.SerializeU64(obj.StorageCharge); err != nil { return err }
	if err := serializer.SerializeU64(obj.ComputationCharge); err != nil { return err }
	if err := serializer.SerializeU64(obj.StorageRebate); err != nil { return err }
	if err := serializer.SerializeU64(obj.NonRefundableStorageFee); err != nil { return err }
	if err := serializer.SerializeU64(obj.EpochStartTimestampMs); err != nil { return err }
	if err := serialize_vector_tuple3_SequenceNumber_vector_vector_u8_vector_ObjectID(obj.SystemPackages, serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *ChangeEpoch) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func DeserializeChangeEpoch(deserializer serde.Deserializer) (ChangeEpoch, error) {
	var obj ChangeEpoch
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := deserializer.DeserializeU64(); err == nil { obj.Epoch = val } else { return obj, err }
	if val, err := DeserializeProtocolVersion(deserializer); err == nil { obj.ProtocolVersion = val } else { return obj, err }
	if val, err := deserializer.DeserializeU64(); err == nil { obj.StorageCharge = val } else { return obj, err }
	if val, err := deserializer.DeserializeU64(); err == nil { obj.ComputationCharge = val } else { return obj, err }
	if val, err := deserializer.DeserializeU64(); err == nil { obj.StorageRebate = val } else { return obj, err }
	if val, err := deserializer.DeserializeU64(); err == nil { obj.NonRefundableStorageFee = val } else { return obj, err }
	if val, err := deserializer.DeserializeU64(); err == nil { obj.EpochStartTimestampMs = val } else { return obj, err }
	if val, err := deserialize_vector_tuple3_SequenceNumber_vector_vector_u8_vector_ObjectID(deserializer); err == nil { obj.SystemPackages = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

func BcsDeserializeChangeEpoch(input []byte) (ChangeEpoch, error) {
	if input == nil {
		var obj ChangeEpoch
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input);
	obj, err := DeserializeChangeEpoch(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type CheckpointCommitment interface {
	isCheckpointCommitment()
	Serialize(serializer serde.Serializer) error
	BcsSerialize() ([]byte, error)
}

func DeserializeCheckpointCommitment(deserializer serde.Deserializer) (CheckpointCommitment, error) {
	index, err := deserializer.DeserializeVariantIndex()
	if err != nil { return nil, err }

	switch index {
	case 0:
		if val, err := load_CheckpointCommitment__EcmhLiveObjectSetDigest(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	default:
		return nil, fmt.Errorf("Unknown variant index for CheckpointCommitment: %d", index)
	}
}

func BcsDeserializeCheckpointCommitment(input []byte) (CheckpointCommitment, error) {
	if input == nil {
		var obj CheckpointCommitment
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input);
	obj, err := DeserializeCheckpointCommitment(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type CheckpointCommitment__EcmhLiveObjectSetDigest struct {
	Value ECMHLiveObjectSetDigest
}

func (*CheckpointCommitment__EcmhLiveObjectSetDigest) isCheckpointCommitment() {}

func (obj *CheckpointCommitment__EcmhLiveObjectSetDigest) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(0)
	if err := obj.Value.Serialize(serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *CheckpointCommitment__EcmhLiveObjectSetDigest) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_CheckpointCommitment__EcmhLiveObjectSetDigest(deserializer serde.Deserializer) (CheckpointCommitment__EcmhLiveObjectSetDigest, error) {
	var obj CheckpointCommitment__EcmhLiveObjectSetDigest
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := DeserializeECMHLiveObjectSetDigest(deserializer); err == nil { obj.Value = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type CheckpointContents interface {
	isCheckpointContents()
	Serialize(serializer serde.Serializer) error
	BcsSerialize() ([]byte, error)
}

func DeserializeCheckpointContents(deserializer serde.Deserializer) (CheckpointContents, error) {
	index, err := deserializer.DeserializeVariantIndex()
	if err != nil { return nil, err }

	switch index {
	case 0:
		if val, err := load_CheckpointContents__V1(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	default:
		return nil, fmt.Errorf("Unknown variant index for CheckpointContents: %d", index)
	}
}

func BcsDeserializeCheckpointContents(input []byte) (CheckpointContents, error) {
	if input == nil {
		var obj CheckpointContents
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input);
	obj, err := DeserializeCheckpointContents(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type CheckpointContents__V1 struct {
	Value CheckpointContentsV1
}

func (*CheckpointContents__V1) isCheckpointContents() {}

func (obj *CheckpointContents__V1) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(0)
	if err := obj.Value.Serialize(serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *CheckpointContents__V1) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_CheckpointContents__V1(deserializer serde.Deserializer) (CheckpointContents__V1, error) {
	var obj CheckpointContents__V1
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := DeserializeCheckpointContentsV1(deserializer); err == nil { obj.Value = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type CheckpointContentsDigest struct {
	Value Digest
}

func (obj *CheckpointContentsDigest) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	if err := obj.Value.Serialize(serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *CheckpointContentsDigest) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func DeserializeCheckpointContentsDigest(deserializer serde.Deserializer) (CheckpointContentsDigest, error) {
	var obj CheckpointContentsDigest
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := DeserializeDigest(deserializer); err == nil { obj.Value = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

func BcsDeserializeCheckpointContentsDigest(input []byte) (CheckpointContentsDigest, error) {
	if input == nil {
		var obj CheckpointContentsDigest
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input);
	obj, err := DeserializeCheckpointContentsDigest(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type CheckpointContentsV1 struct {
	Transactions []ExecutionDigests
	UserSignatures [][]GenericSignature
}

func (obj *CheckpointContentsV1) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	if err := serialize_vector_ExecutionDigests(obj.Transactions, serializer); err != nil { return err }
	if err := serialize_vector_vector_GenericSignature(obj.UserSignatures, serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *CheckpointContentsV1) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func DeserializeCheckpointContentsV1(deserializer serde.Deserializer) (CheckpointContentsV1, error) {
	var obj CheckpointContentsV1
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := deserialize_vector_ExecutionDigests(deserializer); err == nil { obj.Transactions = val } else { return obj, err }
	if val, err := deserialize_vector_vector_GenericSignature(deserializer); err == nil { obj.UserSignatures = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

func BcsDeserializeCheckpointContentsV1(input []byte) (CheckpointContentsV1, error) {
	if input == nil {
		var obj CheckpointContentsV1
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input);
	obj, err := DeserializeCheckpointContentsV1(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type CheckpointDigest struct {
	Value Digest
}

func (obj *CheckpointDigest) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	if err := obj.Value.Serialize(serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *CheckpointDigest) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func DeserializeCheckpointDigest(deserializer serde.Deserializer) (CheckpointDigest, error) {
	var obj CheckpointDigest
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := DeserializeDigest(deserializer); err == nil { obj.Value = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

func BcsDeserializeCheckpointDigest(input []byte) (CheckpointDigest, error) {
	if input == nil {
		var obj CheckpointDigest
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input);
	obj, err := DeserializeCheckpointDigest(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type CheckpointSummary struct {
	Epoch uint64
	SequenceNumber uint64
	NetworkTotalTransactions uint64
	ContentDigest CheckpointContentsDigest
	PreviousDigest *CheckpointDigest
	EpochRollingGasCostSummary GasCostSummary
	TimestampMs uint64
	CheckpointCommitments []CheckpointCommitment
	EndOfEpochData *EndOfEpochData
	VersionSpecificData []uint8
}

func (obj *CheckpointSummary) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	if err := serializer.SerializeU64(obj.Epoch); err != nil { return err }
	if err := serializer.SerializeU64(obj.SequenceNumber); err != nil { return err }
	if err := serializer.SerializeU64(obj.NetworkTotalTransactions); err != nil { return err }
	if err := obj.ContentDigest.Serialize(serializer); err != nil { return err }
	if err := serialize_option_CheckpointDigest(obj.PreviousDigest, serializer); err != nil { return err }
	if err := obj.EpochRollingGasCostSummary.Serialize(serializer); err != nil { return err }
	if err := serializer.SerializeU64(obj.TimestampMs); err != nil { return err }
	if err := serialize_vector_CheckpointCommitment(obj.CheckpointCommitments, serializer); err != nil { return err }
	if err := serialize_option_EndOfEpochData(obj.EndOfEpochData, serializer); err != nil { return err }
	if err := serialize_vector_u8(obj.VersionSpecificData, serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *CheckpointSummary) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func DeserializeCheckpointSummary(deserializer serde.Deserializer) (CheckpointSummary, error) {
	var obj CheckpointSummary
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := deserializer.DeserializeU64(); err == nil { obj.Epoch = val } else { return obj, err }
	if val, err := deserializer.DeserializeU64(); err == nil { obj.SequenceNumber = val } else { return obj, err }
	if val, err := deserializer.DeserializeU64(); err == nil { obj.NetworkTotalTransactions = val } else { return obj, err }
	if val, err := DeserializeCheckpointContentsDigest(deserializer); err == nil { obj.ContentDigest = val } else { return obj, err }
	if val, err := deserialize_option_CheckpointDigest(deserializer); err == nil { obj.PreviousDigest = val } else { return obj, err }
	if val, err := DeserializeGasCostSummary(deserializer); err == nil { obj.EpochRollingGasCostSummary = val } else { return obj, err }
	if val, err := deserializer.DeserializeU64(); err == nil { obj.TimestampMs = val } else { return obj, err }
	if val, err := deserialize_vector_CheckpointCommitment(deserializer); err == nil { obj.CheckpointCommitments = val } else { return obj, err }
	if val, err := deserialize_option_EndOfEpochData(deserializer); err == nil { obj.EndOfEpochData = val } else { return obj, err }
	if val, err := deserialize_vector_u8(deserializer); err == nil { obj.VersionSpecificData = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

func BcsDeserializeCheckpointSummary(input []byte) (CheckpointSummary, error) {
	if input == nil {
		var obj CheckpointSummary
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input);
	obj, err := DeserializeCheckpointSummary(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type Command interface {
	isCommand()
	Serialize(serializer serde.Serializer) error
	BcsSerialize() ([]byte, error)
}

func DeserializeCommand(deserializer serde.Deserializer) (Command, error) {
	index, err := deserializer.DeserializeVariantIndex()
	if err != nil { return nil, err }

	switch index {
	case 0:
		if val, err := load_Command__MoveCall(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 1:
		if val, err := load_Command__TransferObjects(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 2:
		if val, err := load_Command__SplitCoins(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 3:
		if val, err := load_Command__MergeCoins(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 4:
		if val, err := load_Command__Publish(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 5:
		if val, err := load_Command__MakeMoveVec(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 6:
		if val, err := load_Command__Upgrade(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	default:
		return nil, fmt.Errorf("Unknown variant index for Command: %d", index)
	}
}

func BcsDeserializeCommand(input []byte) (Command, error) {
	if input == nil {
		var obj Command
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input);
	obj, err := DeserializeCommand(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type Command__MoveCall struct {
	Value ProgrammableMoveCall
}

func (*Command__MoveCall) isCommand() {}

func (obj *Command__MoveCall) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(0)
	if err := obj.Value.Serialize(serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *Command__MoveCall) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_Command__MoveCall(deserializer serde.Deserializer) (Command__MoveCall, error) {
	var obj Command__MoveCall
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := DeserializeProgrammableMoveCall(deserializer); err == nil { obj.Value = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type Command__TransferObjects struct {
	Field0 []Argument
	Field1 Argument
}

func (*Command__TransferObjects) isCommand() {}

func (obj *Command__TransferObjects) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(1)
	if err := serialize_vector_Argument(obj.Field0, serializer); err != nil { return err }
	if err := obj.Field1.Serialize(serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *Command__TransferObjects) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_Command__TransferObjects(deserializer serde.Deserializer) (Command__TransferObjects, error) {
	var obj Command__TransferObjects
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := deserialize_vector_Argument(deserializer); err == nil { obj.Field0 = val } else { return obj, err }
	if val, err := DeserializeArgument(deserializer); err == nil { obj.Field1 = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type Command__SplitCoins struct {
	Field0 Argument
	Field1 []Argument
}

func (*Command__SplitCoins) isCommand() {}

func (obj *Command__SplitCoins) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(2)
	if err := obj.Field0.Serialize(serializer); err != nil { return err }
	if err := serialize_vector_Argument(obj.Field1, serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *Command__SplitCoins) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_Command__SplitCoins(deserializer serde.Deserializer) (Command__SplitCoins, error) {
	var obj Command__SplitCoins
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := DeserializeArgument(deserializer); err == nil { obj.Field0 = val } else { return obj, err }
	if val, err := deserialize_vector_Argument(deserializer); err == nil { obj.Field1 = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type Command__MergeCoins struct {
	Field0 Argument
	Field1 []Argument
}

func (*Command__MergeCoins) isCommand() {}

func (obj *Command__MergeCoins) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(3)
	if err := obj.Field0.Serialize(serializer); err != nil { return err }
	if err := serialize_vector_Argument(obj.Field1, serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *Command__MergeCoins) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_Command__MergeCoins(deserializer serde.Deserializer) (Command__MergeCoins, error) {
	var obj Command__MergeCoins
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := DeserializeArgument(deserializer); err == nil { obj.Field0 = val } else { return obj, err }
	if val, err := deserialize_vector_Argument(deserializer); err == nil { obj.Field1 = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type Command__Publish struct {
	Field0 [][]uint8
	Field1 []ObjectID
}

func (*Command__Publish) isCommand() {}

func (obj *Command__Publish) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(4)
	if err := serialize_vector_vector_u8(obj.Field0, serializer); err != nil { return err }
	if err := serialize_vector_ObjectID(obj.Field1, serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *Command__Publish) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_Command__Publish(deserializer serde.Deserializer) (Command__Publish, error) {
	var obj Command__Publish
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := deserialize_vector_vector_u8(deserializer); err == nil { obj.Field0 = val } else { return obj, err }
	if val, err := deserialize_vector_ObjectID(deserializer); err == nil { obj.Field1 = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type Command__MakeMoveVec struct {
	Field0 *TypeTag
	Field1 []Argument
}

func (*Command__MakeMoveVec) isCommand() {}

func (obj *Command__MakeMoveVec) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(5)
	if err := serialize_option_TypeTag(obj.Field0, serializer); err != nil { return err }
	if err := serialize_vector_Argument(obj.Field1, serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *Command__MakeMoveVec) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_Command__MakeMoveVec(deserializer serde.Deserializer) (Command__MakeMoveVec, error) {
	var obj Command__MakeMoveVec
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := deserialize_option_TypeTag(deserializer); err == nil { obj.Field0 = val } else { return obj, err }
	if val, err := deserialize_vector_Argument(deserializer); err == nil { obj.Field1 = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type Command__Upgrade struct {
	Field0 [][]uint8
	Field1 []ObjectID
	Field2 ObjectID
	Field3 Argument
}

func (*Command__Upgrade) isCommand() {}

func (obj *Command__Upgrade) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(6)
	if err := serialize_vector_vector_u8(obj.Field0, serializer); err != nil { return err }
	if err := serialize_vector_ObjectID(obj.Field1, serializer); err != nil { return err }
	if err := obj.Field2.Serialize(serializer); err != nil { return err }
	if err := obj.Field3.Serialize(serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *Command__Upgrade) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_Command__Upgrade(deserializer serde.Deserializer) (Command__Upgrade, error) {
	var obj Command__Upgrade
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := deserialize_vector_vector_u8(deserializer); err == nil { obj.Field0 = val } else { return obj, err }
	if val, err := deserialize_vector_ObjectID(deserializer); err == nil { obj.Field1 = val } else { return obj, err }
	if val, err := DeserializeObjectID(deserializer); err == nil { obj.Field2 = val } else { return obj, err }
	if val, err := DeserializeArgument(deserializer); err == nil { obj.Field3 = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type CommandArgumentError interface {
	isCommandArgumentError()
	Serialize(serializer serde.Serializer) error
	BcsSerialize() ([]byte, error)
}

func DeserializeCommandArgumentError(deserializer serde.Deserializer) (CommandArgumentError, error) {
	index, err := deserializer.DeserializeVariantIndex()
	if err != nil { return nil, err }

	switch index {
	case 0:
		if val, err := load_CommandArgumentError__TypeMismatch(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 1:
		if val, err := load_CommandArgumentError__InvalidBcsBytes(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 2:
		if val, err := load_CommandArgumentError__InvalidUsageOfPureArg(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 3:
		if val, err := load_CommandArgumentError__InvalidArgumentToPrivateEntryFunction(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 4:
		if val, err := load_CommandArgumentError__IndexOutOfBounds(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 5:
		if val, err := load_CommandArgumentError__SecondaryIndexOutOfBounds(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 6:
		if val, err := load_CommandArgumentError__InvalidResultArity(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 7:
		if val, err := load_CommandArgumentError__InvalidGasCoinUsage(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 8:
		if val, err := load_CommandArgumentError__InvalidValueUsage(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 9:
		if val, err := load_CommandArgumentError__InvalidObjectByValue(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 10:
		if val, err := load_CommandArgumentError__InvalidObjectByMutRef(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	default:
		return nil, fmt.Errorf("Unknown variant index for CommandArgumentError: %d", index)
	}
}

func BcsDeserializeCommandArgumentError(input []byte) (CommandArgumentError, error) {
	if input == nil {
		var obj CommandArgumentError
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input);
	obj, err := DeserializeCommandArgumentError(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type CommandArgumentError__TypeMismatch struct {
}

func (*CommandArgumentError__TypeMismatch) isCommandArgumentError() {}

func (obj *CommandArgumentError__TypeMismatch) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(0)
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *CommandArgumentError__TypeMismatch) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_CommandArgumentError__TypeMismatch(deserializer serde.Deserializer) (CommandArgumentError__TypeMismatch, error) {
	var obj CommandArgumentError__TypeMismatch
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type CommandArgumentError__InvalidBcsBytes struct {
}

func (*CommandArgumentError__InvalidBcsBytes) isCommandArgumentError() {}

func (obj *CommandArgumentError__InvalidBcsBytes) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(1)
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *CommandArgumentError__InvalidBcsBytes) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_CommandArgumentError__InvalidBcsBytes(deserializer serde.Deserializer) (CommandArgumentError__InvalidBcsBytes, error) {
	var obj CommandArgumentError__InvalidBcsBytes
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type CommandArgumentError__InvalidUsageOfPureArg struct {
}

func (*CommandArgumentError__InvalidUsageOfPureArg) isCommandArgumentError() {}

func (obj *CommandArgumentError__InvalidUsageOfPureArg) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(2)
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *CommandArgumentError__InvalidUsageOfPureArg) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_CommandArgumentError__InvalidUsageOfPureArg(deserializer serde.Deserializer) (CommandArgumentError__InvalidUsageOfPureArg, error) {
	var obj CommandArgumentError__InvalidUsageOfPureArg
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type CommandArgumentError__InvalidArgumentToPrivateEntryFunction struct {
}

func (*CommandArgumentError__InvalidArgumentToPrivateEntryFunction) isCommandArgumentError() {}

func (obj *CommandArgumentError__InvalidArgumentToPrivateEntryFunction) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(3)
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *CommandArgumentError__InvalidArgumentToPrivateEntryFunction) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_CommandArgumentError__InvalidArgumentToPrivateEntryFunction(deserializer serde.Deserializer) (CommandArgumentError__InvalidArgumentToPrivateEntryFunction, error) {
	var obj CommandArgumentError__InvalidArgumentToPrivateEntryFunction
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type CommandArgumentError__IndexOutOfBounds struct {
	Idx uint16
}

func (*CommandArgumentError__IndexOutOfBounds) isCommandArgumentError() {}

func (obj *CommandArgumentError__IndexOutOfBounds) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(4)
	if err := serializer.SerializeU16(obj.Idx); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *CommandArgumentError__IndexOutOfBounds) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_CommandArgumentError__IndexOutOfBounds(deserializer serde.Deserializer) (CommandArgumentError__IndexOutOfBounds, error) {
	var obj CommandArgumentError__IndexOutOfBounds
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := deserializer.DeserializeU16(); err == nil { obj.Idx = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type CommandArgumentError__SecondaryIndexOutOfBounds struct {
	ResultIdx uint16
	SecondaryIdx uint16
}

func (*CommandArgumentError__SecondaryIndexOutOfBounds) isCommandArgumentError() {}

func (obj *CommandArgumentError__SecondaryIndexOutOfBounds) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(5)
	if err := serializer.SerializeU16(obj.ResultIdx); err != nil { return err }
	if err := serializer.SerializeU16(obj.SecondaryIdx); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *CommandArgumentError__SecondaryIndexOutOfBounds) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_CommandArgumentError__SecondaryIndexOutOfBounds(deserializer serde.Deserializer) (CommandArgumentError__SecondaryIndexOutOfBounds, error) {
	var obj CommandArgumentError__SecondaryIndexOutOfBounds
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := deserializer.DeserializeU16(); err == nil { obj.ResultIdx = val } else { return obj, err }
	if val, err := deserializer.DeserializeU16(); err == nil { obj.SecondaryIdx = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type CommandArgumentError__InvalidResultArity struct {
	ResultIdx uint16
}

func (*CommandArgumentError__InvalidResultArity) isCommandArgumentError() {}

func (obj *CommandArgumentError__InvalidResultArity) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(6)
	if err := serializer.SerializeU16(obj.ResultIdx); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *CommandArgumentError__InvalidResultArity) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_CommandArgumentError__InvalidResultArity(deserializer serde.Deserializer) (CommandArgumentError__InvalidResultArity, error) {
	var obj CommandArgumentError__InvalidResultArity
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := deserializer.DeserializeU16(); err == nil { obj.ResultIdx = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type CommandArgumentError__InvalidGasCoinUsage struct {
}

func (*CommandArgumentError__InvalidGasCoinUsage) isCommandArgumentError() {}

func (obj *CommandArgumentError__InvalidGasCoinUsage) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(7)
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *CommandArgumentError__InvalidGasCoinUsage) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_CommandArgumentError__InvalidGasCoinUsage(deserializer serde.Deserializer) (CommandArgumentError__InvalidGasCoinUsage, error) {
	var obj CommandArgumentError__InvalidGasCoinUsage
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type CommandArgumentError__InvalidValueUsage struct {
}

func (*CommandArgumentError__InvalidValueUsage) isCommandArgumentError() {}

func (obj *CommandArgumentError__InvalidValueUsage) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(8)
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *CommandArgumentError__InvalidValueUsage) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_CommandArgumentError__InvalidValueUsage(deserializer serde.Deserializer) (CommandArgumentError__InvalidValueUsage, error) {
	var obj CommandArgumentError__InvalidValueUsage
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type CommandArgumentError__InvalidObjectByValue struct {
}

func (*CommandArgumentError__InvalidObjectByValue) isCommandArgumentError() {}

func (obj *CommandArgumentError__InvalidObjectByValue) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(9)
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *CommandArgumentError__InvalidObjectByValue) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_CommandArgumentError__InvalidObjectByValue(deserializer serde.Deserializer) (CommandArgumentError__InvalidObjectByValue, error) {
	var obj CommandArgumentError__InvalidObjectByValue
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type CommandArgumentError__InvalidObjectByMutRef struct {
}

func (*CommandArgumentError__InvalidObjectByMutRef) isCommandArgumentError() {}

func (obj *CommandArgumentError__InvalidObjectByMutRef) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(10)
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *CommandArgumentError__InvalidObjectByMutRef) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_CommandArgumentError__InvalidObjectByMutRef(deserializer serde.Deserializer) (CommandArgumentError__InvalidObjectByMutRef, error) {
	var obj CommandArgumentError__InvalidObjectByMutRef
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type CompressedSignature interface {
	isCompressedSignature()
	Serialize(serializer serde.Serializer) error
	BcsSerialize() ([]byte, error)
}

func DeserializeCompressedSignature(deserializer serde.Deserializer) (CompressedSignature, error) {
	index, err := deserializer.DeserializeVariantIndex()
	if err != nil { return nil, err }

	switch index {
	case 0:
		if val, err := load_CompressedSignature__Ed25519(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 1:
		if val, err := load_CompressedSignature__Secp256k1(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	default:
		return nil, fmt.Errorf("Unknown variant index for CompressedSignature: %d", index)
	}
}

func BcsDeserializeCompressedSignature(input []byte) (CompressedSignature, error) {
	if input == nil {
		var obj CompressedSignature
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input);
	obj, err := DeserializeCompressedSignature(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type CompressedSignature__Ed25519 [64]uint8

func (*CompressedSignature__Ed25519) isCompressedSignature() {}

func (obj *CompressedSignature__Ed25519) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(0)
	if err := serialize_array64_u8_array((([64]uint8)(*obj)), serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *CompressedSignature__Ed25519) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_CompressedSignature__Ed25519(deserializer serde.Deserializer) (CompressedSignature__Ed25519, error) {
	var obj [64]uint8
	if err := deserializer.IncreaseContainerDepth(); err != nil { return (CompressedSignature__Ed25519)(obj), err }
	if val, err := deserialize_array64_u8_array(deserializer); err == nil { obj = val } else { return ((CompressedSignature__Ed25519)(obj)), err }
	deserializer.DecreaseContainerDepth()
	return (CompressedSignature__Ed25519)(obj), nil
}

type CompressedSignature__Secp256k1 [64]uint8

func (*CompressedSignature__Secp256k1) isCompressedSignature() {}

func (obj *CompressedSignature__Secp256k1) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(1)
	if err := serialize_array64_u8_array((([64]uint8)(*obj)), serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *CompressedSignature__Secp256k1) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_CompressedSignature__Secp256k1(deserializer serde.Deserializer) (CompressedSignature__Secp256k1, error) {
	var obj [64]uint8
	if err := deserializer.IncreaseContainerDepth(); err != nil { return (CompressedSignature__Secp256k1)(obj), err }
	if val, err := deserialize_array64_u8_array(deserializer); err == nil { obj = val } else { return ((CompressedSignature__Secp256k1)(obj)), err }
	deserializer.DecreaseContainerDepth()
	return (CompressedSignature__Secp256k1)(obj), nil
}

type ConsensusCommitPrologue struct {
	Epoch uint64
	Round uint64
	CommitTimestampMs uint64
}

func (obj *ConsensusCommitPrologue) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	if err := serializer.SerializeU64(obj.Epoch); err != nil { return err }
	if err := serializer.SerializeU64(obj.Round); err != nil { return err }
	if err := serializer.SerializeU64(obj.CommitTimestampMs); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *ConsensusCommitPrologue) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func DeserializeConsensusCommitPrologue(deserializer serde.Deserializer) (ConsensusCommitPrologue, error) {
	var obj ConsensusCommitPrologue
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := deserializer.DeserializeU64(); err == nil { obj.Epoch = val } else { return obj, err }
	if val, err := deserializer.DeserializeU64(); err == nil { obj.Round = val } else { return obj, err }
	if val, err := deserializer.DeserializeU64(); err == nil { obj.CommitTimestampMs = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

func BcsDeserializeConsensusCommitPrologue(input []byte) (ConsensusCommitPrologue, error) {
	if input == nil {
		var obj ConsensusCommitPrologue
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input);
	obj, err := DeserializeConsensusCommitPrologue(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type Data interface {
	isData()
	Serialize(serializer serde.Serializer) error
	BcsSerialize() ([]byte, error)
}

func DeserializeData(deserializer serde.Deserializer) (Data, error) {
	index, err := deserializer.DeserializeVariantIndex()
	if err != nil { return nil, err }

	switch index {
	case 0:
		if val, err := load_Data__Move(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 1:
		if val, err := load_Data__Package(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	default:
		return nil, fmt.Errorf("Unknown variant index for Data: %d", index)
	}
}

func BcsDeserializeData(input []byte) (Data, error) {
	if input == nil {
		var obj Data
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input);
	obj, err := DeserializeData(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type Data__Move struct {
	Value MoveObject
}

func (*Data__Move) isData() {}

func (obj *Data__Move) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(0)
	if err := obj.Value.Serialize(serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *Data__Move) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_Data__Move(deserializer serde.Deserializer) (Data__Move, error) {
	var obj Data__Move
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := DeserializeMoveObject(deserializer); err == nil { obj.Value = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type Data__Package struct {
	Value MovePackage
}

func (*Data__Package) isData() {}

func (obj *Data__Package) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(1)
	if err := obj.Value.Serialize(serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *Data__Package) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_Data__Package(deserializer serde.Deserializer) (Data__Package, error) {
	var obj Data__Package
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := DeserializeMovePackage(deserializer); err == nil { obj.Value = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type DeleteKind interface {
	isDeleteKind()
	Serialize(serializer serde.Serializer) error
	BcsSerialize() ([]byte, error)
}

func DeserializeDeleteKind(deserializer serde.Deserializer) (DeleteKind, error) {
	index, err := deserializer.DeserializeVariantIndex()
	if err != nil { return nil, err }

	switch index {
	case 0:
		if val, err := load_DeleteKind__Normal(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 1:
		if val, err := load_DeleteKind__UnwrapThenDelete(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 2:
		if val, err := load_DeleteKind__Wrap(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	default:
		return nil, fmt.Errorf("Unknown variant index for DeleteKind: %d", index)
	}
}

func BcsDeserializeDeleteKind(input []byte) (DeleteKind, error) {
	if input == nil {
		var obj DeleteKind
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input);
	obj, err := DeserializeDeleteKind(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type DeleteKind__Normal struct {
}

func (*DeleteKind__Normal) isDeleteKind() {}

func (obj *DeleteKind__Normal) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(0)
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *DeleteKind__Normal) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_DeleteKind__Normal(deserializer serde.Deserializer) (DeleteKind__Normal, error) {
	var obj DeleteKind__Normal
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type DeleteKind__UnwrapThenDelete struct {
}

func (*DeleteKind__UnwrapThenDelete) isDeleteKind() {}

func (obj *DeleteKind__UnwrapThenDelete) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(1)
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *DeleteKind__UnwrapThenDelete) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_DeleteKind__UnwrapThenDelete(deserializer serde.Deserializer) (DeleteKind__UnwrapThenDelete, error) {
	var obj DeleteKind__UnwrapThenDelete
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type DeleteKind__Wrap struct {
}

func (*DeleteKind__Wrap) isDeleteKind() {}

func (obj *DeleteKind__Wrap) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(2)
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *DeleteKind__Wrap) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_DeleteKind__Wrap(deserializer serde.Deserializer) (DeleteKind__Wrap, error) {
	var obj DeleteKind__Wrap
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type Digest []byte

func (obj *Digest) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	if err := serializer.SerializeBytes((([]byte)(*obj))); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *Digest) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func DeserializeDigest(deserializer serde.Deserializer) (Digest, error) {
	var obj []byte
	if err := deserializer.IncreaseContainerDepth(); err != nil { return (Digest)(obj), err }
	if val, err := deserializer.DeserializeBytes(); err == nil { obj = val } else { return ((Digest)(obj)), err }
	deserializer.DecreaseContainerDepth()
	return (Digest)(obj), nil
}

func BcsDeserializeDigest(input []byte) (Digest, error) {
	if input == nil {
		var obj Digest
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input);
	obj, err := DeserializeDigest(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type ECMHLiveObjectSetDigest struct {
	Digest Digest
}

func (obj *ECMHLiveObjectSetDigest) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	if err := obj.Digest.Serialize(serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *ECMHLiveObjectSetDigest) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func DeserializeECMHLiveObjectSetDigest(deserializer serde.Deserializer) (ECMHLiveObjectSetDigest, error) {
	var obj ECMHLiveObjectSetDigest
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := DeserializeDigest(deserializer); err == nil { obj.Digest = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

func BcsDeserializeECMHLiveObjectSetDigest(input []byte) (ECMHLiveObjectSetDigest, error) {
	if input == nil {
		var obj ECMHLiveObjectSetDigest
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input);
	obj, err := DeserializeECMHLiveObjectSetDigest(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type EmptySignInfo struct {
}

func (obj *EmptySignInfo) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *EmptySignInfo) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func DeserializeEmptySignInfo(deserializer serde.Deserializer) (EmptySignInfo, error) {
	var obj EmptySignInfo
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

func BcsDeserializeEmptySignInfo(input []byte) (EmptySignInfo, error) {
	if input == nil {
		var obj EmptySignInfo
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input);
	obj, err := DeserializeEmptySignInfo(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type EndOfEpochData struct {
	NextEpochCommittee []struct {Field0 AuthorityPublicKeyBytes; Field1 uint64}
	NextEpochProtocolVersion ProtocolVersion
	EpochCommitments []CheckpointCommitment
}

func (obj *EndOfEpochData) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	if err := serialize_vector_tuple2_AuthorityPublicKeyBytes_u64(obj.NextEpochCommittee, serializer); err != nil { return err }
	if err := obj.NextEpochProtocolVersion.Serialize(serializer); err != nil { return err }
	if err := serialize_vector_CheckpointCommitment(obj.EpochCommitments, serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *EndOfEpochData) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func DeserializeEndOfEpochData(deserializer serde.Deserializer) (EndOfEpochData, error) {
	var obj EndOfEpochData
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := deserialize_vector_tuple2_AuthorityPublicKeyBytes_u64(deserializer); err == nil { obj.NextEpochCommittee = val } else { return obj, err }
	if val, err := DeserializeProtocolVersion(deserializer); err == nil { obj.NextEpochProtocolVersion = val } else { return obj, err }
	if val, err := deserialize_vector_CheckpointCommitment(deserializer); err == nil { obj.EpochCommitments = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

func BcsDeserializeEndOfEpochData(input []byte) (EndOfEpochData, error) {
	if input == nil {
		var obj EndOfEpochData
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input);
	obj, err := DeserializeEndOfEpochData(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type Envelope struct {
	Data SenderSignedData
	AuthSignature EmptySignInfo
}

func (obj *Envelope) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	if err := obj.Data.Serialize(serializer); err != nil { return err }
	if err := obj.AuthSignature.Serialize(serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *Envelope) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func DeserializeEnvelope(deserializer serde.Deserializer) (Envelope, error) {
	var obj Envelope
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := DeserializeSenderSignedData(deserializer); err == nil { obj.Data = val } else { return obj, err }
	if val, err := DeserializeEmptySignInfo(deserializer); err == nil { obj.AuthSignature = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

func BcsDeserializeEnvelope(input []byte) (Envelope, error) {
	if input == nil {
		var obj Envelope
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input);
	obj, err := DeserializeEnvelope(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type ExecutionData struct {
	Transaction Envelope
	Effects TransactionEffects
}

func (obj *ExecutionData) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	if err := obj.Transaction.Serialize(serializer); err != nil { return err }
	if err := obj.Effects.Serialize(serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *ExecutionData) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func DeserializeExecutionData(deserializer serde.Deserializer) (ExecutionData, error) {
	var obj ExecutionData
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := DeserializeEnvelope(deserializer); err == nil { obj.Transaction = val } else { return obj, err }
	if val, err := DeserializeTransactionEffects(deserializer); err == nil { obj.Effects = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

func BcsDeserializeExecutionData(input []byte) (ExecutionData, error) {
	if input == nil {
		var obj ExecutionData
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input);
	obj, err := DeserializeExecutionData(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type ExecutionDigests struct {
	Transaction TransactionDigest
	Effects TransactionEffectsDigest
}

func (obj *ExecutionDigests) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	if err := obj.Transaction.Serialize(serializer); err != nil { return err }
	if err := obj.Effects.Serialize(serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *ExecutionDigests) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func DeserializeExecutionDigests(deserializer serde.Deserializer) (ExecutionDigests, error) {
	var obj ExecutionDigests
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := DeserializeTransactionDigest(deserializer); err == nil { obj.Transaction = val } else { return obj, err }
	if val, err := DeserializeTransactionEffectsDigest(deserializer); err == nil { obj.Effects = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

func BcsDeserializeExecutionDigests(input []byte) (ExecutionDigests, error) {
	if input == nil {
		var obj ExecutionDigests
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input);
	obj, err := DeserializeExecutionDigests(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type ExecutionFailureStatus interface {
	isExecutionFailureStatus()
	Serialize(serializer serde.Serializer) error
	BcsSerialize() ([]byte, error)
}

func DeserializeExecutionFailureStatus(deserializer serde.Deserializer) (ExecutionFailureStatus, error) {
	index, err := deserializer.DeserializeVariantIndex()
	if err != nil { return nil, err }

	switch index {
	case 0:
		if val, err := load_ExecutionFailureStatus__InsufficientGas(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 1:
		if val, err := load_ExecutionFailureStatus__InvalidGasObject(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 2:
		if val, err := load_ExecutionFailureStatus__InvariantViolation(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 3:
		if val, err := load_ExecutionFailureStatus__FeatureNotYetSupported(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 4:
		if val, err := load_ExecutionFailureStatus__MoveObjectTooBig(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 5:
		if val, err := load_ExecutionFailureStatus__MovePackageTooBig(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 6:
		if val, err := load_ExecutionFailureStatus__CircularObjectOwnership(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 7:
		if val, err := load_ExecutionFailureStatus__InsufficientCoinBalance(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 8:
		if val, err := load_ExecutionFailureStatus__CoinBalanceOverflow(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 9:
		if val, err := load_ExecutionFailureStatus__PublishErrorNonZeroAddress(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 10:
		if val, err := load_ExecutionFailureStatus__SuiMoveVerificationError(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 11:
		if val, err := load_ExecutionFailureStatus__MovePrimitiveRuntimeError(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 12:
		if val, err := load_ExecutionFailureStatus__MoveAbort(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 13:
		if val, err := load_ExecutionFailureStatus__VmVerificationOrDeserializationError(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 14:
		if val, err := load_ExecutionFailureStatus__VmInvariantViolation(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 15:
		if val, err := load_ExecutionFailureStatus__FunctionNotFound(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 16:
		if val, err := load_ExecutionFailureStatus__ArityMismatch(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 17:
		if val, err := load_ExecutionFailureStatus__TypeArityMismatch(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 18:
		if val, err := load_ExecutionFailureStatus__NonEntryFunctionInvoked(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 19:
		if val, err := load_ExecutionFailureStatus__CommandArgumentError(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 20:
		if val, err := load_ExecutionFailureStatus__TypeArgumentError(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 21:
		if val, err := load_ExecutionFailureStatus__UnusedValueWithoutDrop(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 22:
		if val, err := load_ExecutionFailureStatus__InvalidPublicFunctionReturnType(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 23:
		if val, err := load_ExecutionFailureStatus__InvalidTransferObject(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 24:
		if val, err := load_ExecutionFailureStatus__EffectsTooLarge(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 25:
		if val, err := load_ExecutionFailureStatus__PublishUpgradeMissingDependency(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 26:
		if val, err := load_ExecutionFailureStatus__PublishUpgradeDependencyDowngrade(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 27:
		if val, err := load_ExecutionFailureStatus__PackageUpgradeError(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 28:
		if val, err := load_ExecutionFailureStatus__WrittenObjectsTooLarge(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 29:
		if val, err := load_ExecutionFailureStatus__CertificateDenied(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	default:
		return nil, fmt.Errorf("Unknown variant index for ExecutionFailureStatus: %d", index)
	}
}

func BcsDeserializeExecutionFailureStatus(input []byte) (ExecutionFailureStatus, error) {
	if input == nil {
		var obj ExecutionFailureStatus
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input);
	obj, err := DeserializeExecutionFailureStatus(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type ExecutionFailureStatus__InsufficientGas struct {
}

func (*ExecutionFailureStatus__InsufficientGas) isExecutionFailureStatus() {}

func (obj *ExecutionFailureStatus__InsufficientGas) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(0)
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *ExecutionFailureStatus__InsufficientGas) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_ExecutionFailureStatus__InsufficientGas(deserializer serde.Deserializer) (ExecutionFailureStatus__InsufficientGas, error) {
	var obj ExecutionFailureStatus__InsufficientGas
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type ExecutionFailureStatus__InvalidGasObject struct {
}

func (*ExecutionFailureStatus__InvalidGasObject) isExecutionFailureStatus() {}

func (obj *ExecutionFailureStatus__InvalidGasObject) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(1)
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *ExecutionFailureStatus__InvalidGasObject) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_ExecutionFailureStatus__InvalidGasObject(deserializer serde.Deserializer) (ExecutionFailureStatus__InvalidGasObject, error) {
	var obj ExecutionFailureStatus__InvalidGasObject
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type ExecutionFailureStatus__InvariantViolation struct {
}

func (*ExecutionFailureStatus__InvariantViolation) isExecutionFailureStatus() {}

func (obj *ExecutionFailureStatus__InvariantViolation) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(2)
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *ExecutionFailureStatus__InvariantViolation) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_ExecutionFailureStatus__InvariantViolation(deserializer serde.Deserializer) (ExecutionFailureStatus__InvariantViolation, error) {
	var obj ExecutionFailureStatus__InvariantViolation
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type ExecutionFailureStatus__FeatureNotYetSupported struct {
}

func (*ExecutionFailureStatus__FeatureNotYetSupported) isExecutionFailureStatus() {}

func (obj *ExecutionFailureStatus__FeatureNotYetSupported) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(3)
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *ExecutionFailureStatus__FeatureNotYetSupported) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_ExecutionFailureStatus__FeatureNotYetSupported(deserializer serde.Deserializer) (ExecutionFailureStatus__FeatureNotYetSupported, error) {
	var obj ExecutionFailureStatus__FeatureNotYetSupported
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type ExecutionFailureStatus__MoveObjectTooBig struct {
	ObjectSize uint64
	MaxObjectSize uint64
}

func (*ExecutionFailureStatus__MoveObjectTooBig) isExecutionFailureStatus() {}

func (obj *ExecutionFailureStatus__MoveObjectTooBig) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(4)
	if err := serializer.SerializeU64(obj.ObjectSize); err != nil { return err }
	if err := serializer.SerializeU64(obj.MaxObjectSize); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *ExecutionFailureStatus__MoveObjectTooBig) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_ExecutionFailureStatus__MoveObjectTooBig(deserializer serde.Deserializer) (ExecutionFailureStatus__MoveObjectTooBig, error) {
	var obj ExecutionFailureStatus__MoveObjectTooBig
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := deserializer.DeserializeU64(); err == nil { obj.ObjectSize = val } else { return obj, err }
	if val, err := deserializer.DeserializeU64(); err == nil { obj.MaxObjectSize = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type ExecutionFailureStatus__MovePackageTooBig struct {
	ObjectSize uint64
	MaxObjectSize uint64
}

func (*ExecutionFailureStatus__MovePackageTooBig) isExecutionFailureStatus() {}

func (obj *ExecutionFailureStatus__MovePackageTooBig) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(5)
	if err := serializer.SerializeU64(obj.ObjectSize); err != nil { return err }
	if err := serializer.SerializeU64(obj.MaxObjectSize); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *ExecutionFailureStatus__MovePackageTooBig) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_ExecutionFailureStatus__MovePackageTooBig(deserializer serde.Deserializer) (ExecutionFailureStatus__MovePackageTooBig, error) {
	var obj ExecutionFailureStatus__MovePackageTooBig
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := deserializer.DeserializeU64(); err == nil { obj.ObjectSize = val } else { return obj, err }
	if val, err := deserializer.DeserializeU64(); err == nil { obj.MaxObjectSize = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type ExecutionFailureStatus__CircularObjectOwnership struct {
	Object ObjectID
}

func (*ExecutionFailureStatus__CircularObjectOwnership) isExecutionFailureStatus() {}

func (obj *ExecutionFailureStatus__CircularObjectOwnership) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(6)
	if err := obj.Object.Serialize(serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *ExecutionFailureStatus__CircularObjectOwnership) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_ExecutionFailureStatus__CircularObjectOwnership(deserializer serde.Deserializer) (ExecutionFailureStatus__CircularObjectOwnership, error) {
	var obj ExecutionFailureStatus__CircularObjectOwnership
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := DeserializeObjectID(deserializer); err == nil { obj.Object = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type ExecutionFailureStatus__InsufficientCoinBalance struct {
}

func (*ExecutionFailureStatus__InsufficientCoinBalance) isExecutionFailureStatus() {}

func (obj *ExecutionFailureStatus__InsufficientCoinBalance) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(7)
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *ExecutionFailureStatus__InsufficientCoinBalance) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_ExecutionFailureStatus__InsufficientCoinBalance(deserializer serde.Deserializer) (ExecutionFailureStatus__InsufficientCoinBalance, error) {
	var obj ExecutionFailureStatus__InsufficientCoinBalance
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type ExecutionFailureStatus__CoinBalanceOverflow struct {
}

func (*ExecutionFailureStatus__CoinBalanceOverflow) isExecutionFailureStatus() {}

func (obj *ExecutionFailureStatus__CoinBalanceOverflow) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(8)
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *ExecutionFailureStatus__CoinBalanceOverflow) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_ExecutionFailureStatus__CoinBalanceOverflow(deserializer serde.Deserializer) (ExecutionFailureStatus__CoinBalanceOverflow, error) {
	var obj ExecutionFailureStatus__CoinBalanceOverflow
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type ExecutionFailureStatus__PublishErrorNonZeroAddress struct {
}

func (*ExecutionFailureStatus__PublishErrorNonZeroAddress) isExecutionFailureStatus() {}

func (obj *ExecutionFailureStatus__PublishErrorNonZeroAddress) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(9)
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *ExecutionFailureStatus__PublishErrorNonZeroAddress) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_ExecutionFailureStatus__PublishErrorNonZeroAddress(deserializer serde.Deserializer) (ExecutionFailureStatus__PublishErrorNonZeroAddress, error) {
	var obj ExecutionFailureStatus__PublishErrorNonZeroAddress
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type ExecutionFailureStatus__SuiMoveVerificationError struct {
}

func (*ExecutionFailureStatus__SuiMoveVerificationError) isExecutionFailureStatus() {}

func (obj *ExecutionFailureStatus__SuiMoveVerificationError) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(10)
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *ExecutionFailureStatus__SuiMoveVerificationError) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_ExecutionFailureStatus__SuiMoveVerificationError(deserializer serde.Deserializer) (ExecutionFailureStatus__SuiMoveVerificationError, error) {
	var obj ExecutionFailureStatus__SuiMoveVerificationError
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type ExecutionFailureStatus__MovePrimitiveRuntimeError struct {
	Value MoveLocationOpt
}

func (*ExecutionFailureStatus__MovePrimitiveRuntimeError) isExecutionFailureStatus() {}

func (obj *ExecutionFailureStatus__MovePrimitiveRuntimeError) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(11)
	if err := obj.Value.Serialize(serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *ExecutionFailureStatus__MovePrimitiveRuntimeError) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_ExecutionFailureStatus__MovePrimitiveRuntimeError(deserializer serde.Deserializer) (ExecutionFailureStatus__MovePrimitiveRuntimeError, error) {
	var obj ExecutionFailureStatus__MovePrimitiveRuntimeError
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := DeserializeMoveLocationOpt(deserializer); err == nil { obj.Value = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type ExecutionFailureStatus__MoveAbort struct {
	Field0 MoveLocation
	Field1 uint64
}

func (*ExecutionFailureStatus__MoveAbort) isExecutionFailureStatus() {}

func (obj *ExecutionFailureStatus__MoveAbort) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(12)
	if err := obj.Field0.Serialize(serializer); err != nil { return err }
	if err := serializer.SerializeU64(obj.Field1); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *ExecutionFailureStatus__MoveAbort) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_ExecutionFailureStatus__MoveAbort(deserializer serde.Deserializer) (ExecutionFailureStatus__MoveAbort, error) {
	var obj ExecutionFailureStatus__MoveAbort
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := DeserializeMoveLocation(deserializer); err == nil { obj.Field0 = val } else { return obj, err }
	if val, err := deserializer.DeserializeU64(); err == nil { obj.Field1 = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type ExecutionFailureStatus__VmVerificationOrDeserializationError struct {
}

func (*ExecutionFailureStatus__VmVerificationOrDeserializationError) isExecutionFailureStatus() {}

func (obj *ExecutionFailureStatus__VmVerificationOrDeserializationError) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(13)
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *ExecutionFailureStatus__VmVerificationOrDeserializationError) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_ExecutionFailureStatus__VmVerificationOrDeserializationError(deserializer serde.Deserializer) (ExecutionFailureStatus__VmVerificationOrDeserializationError, error) {
	var obj ExecutionFailureStatus__VmVerificationOrDeserializationError
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type ExecutionFailureStatus__VmInvariantViolation struct {
}

func (*ExecutionFailureStatus__VmInvariantViolation) isExecutionFailureStatus() {}

func (obj *ExecutionFailureStatus__VmInvariantViolation) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(14)
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *ExecutionFailureStatus__VmInvariantViolation) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_ExecutionFailureStatus__VmInvariantViolation(deserializer serde.Deserializer) (ExecutionFailureStatus__VmInvariantViolation, error) {
	var obj ExecutionFailureStatus__VmInvariantViolation
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type ExecutionFailureStatus__FunctionNotFound struct {
}

func (*ExecutionFailureStatus__FunctionNotFound) isExecutionFailureStatus() {}

func (obj *ExecutionFailureStatus__FunctionNotFound) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(15)
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *ExecutionFailureStatus__FunctionNotFound) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_ExecutionFailureStatus__FunctionNotFound(deserializer serde.Deserializer) (ExecutionFailureStatus__FunctionNotFound, error) {
	var obj ExecutionFailureStatus__FunctionNotFound
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type ExecutionFailureStatus__ArityMismatch struct {
}

func (*ExecutionFailureStatus__ArityMismatch) isExecutionFailureStatus() {}

func (obj *ExecutionFailureStatus__ArityMismatch) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(16)
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *ExecutionFailureStatus__ArityMismatch) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_ExecutionFailureStatus__ArityMismatch(deserializer serde.Deserializer) (ExecutionFailureStatus__ArityMismatch, error) {
	var obj ExecutionFailureStatus__ArityMismatch
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type ExecutionFailureStatus__TypeArityMismatch struct {
}

func (*ExecutionFailureStatus__TypeArityMismatch) isExecutionFailureStatus() {}

func (obj *ExecutionFailureStatus__TypeArityMismatch) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(17)
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *ExecutionFailureStatus__TypeArityMismatch) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_ExecutionFailureStatus__TypeArityMismatch(deserializer serde.Deserializer) (ExecutionFailureStatus__TypeArityMismatch, error) {
	var obj ExecutionFailureStatus__TypeArityMismatch
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type ExecutionFailureStatus__NonEntryFunctionInvoked struct {
}

func (*ExecutionFailureStatus__NonEntryFunctionInvoked) isExecutionFailureStatus() {}

func (obj *ExecutionFailureStatus__NonEntryFunctionInvoked) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(18)
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *ExecutionFailureStatus__NonEntryFunctionInvoked) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_ExecutionFailureStatus__NonEntryFunctionInvoked(deserializer serde.Deserializer) (ExecutionFailureStatus__NonEntryFunctionInvoked, error) {
	var obj ExecutionFailureStatus__NonEntryFunctionInvoked
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type ExecutionFailureStatus__CommandArgumentError struct {
	ArgIdx uint16
	Kind CommandArgumentError
}

func (*ExecutionFailureStatus__CommandArgumentError) isExecutionFailureStatus() {}

func (obj *ExecutionFailureStatus__CommandArgumentError) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(19)
	if err := serializer.SerializeU16(obj.ArgIdx); err != nil { return err }
	if err := obj.Kind.Serialize(serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *ExecutionFailureStatus__CommandArgumentError) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_ExecutionFailureStatus__CommandArgumentError(deserializer serde.Deserializer) (ExecutionFailureStatus__CommandArgumentError, error) {
	var obj ExecutionFailureStatus__CommandArgumentError
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := deserializer.DeserializeU16(); err == nil { obj.ArgIdx = val } else { return obj, err }
	if val, err := DeserializeCommandArgumentError(deserializer); err == nil { obj.Kind = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type ExecutionFailureStatus__TypeArgumentError struct {
	ArgumentIdx uint16
	Kind TypeArgumentError
}

func (*ExecutionFailureStatus__TypeArgumentError) isExecutionFailureStatus() {}

func (obj *ExecutionFailureStatus__TypeArgumentError) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(20)
	if err := serializer.SerializeU16(obj.ArgumentIdx); err != nil { return err }
	if err := obj.Kind.Serialize(serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *ExecutionFailureStatus__TypeArgumentError) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_ExecutionFailureStatus__TypeArgumentError(deserializer serde.Deserializer) (ExecutionFailureStatus__TypeArgumentError, error) {
	var obj ExecutionFailureStatus__TypeArgumentError
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := deserializer.DeserializeU16(); err == nil { obj.ArgumentIdx = val } else { return obj, err }
	if val, err := DeserializeTypeArgumentError(deserializer); err == nil { obj.Kind = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type ExecutionFailureStatus__UnusedValueWithoutDrop struct {
	ResultIdx uint16
	SecondaryIdx uint16
}

func (*ExecutionFailureStatus__UnusedValueWithoutDrop) isExecutionFailureStatus() {}

func (obj *ExecutionFailureStatus__UnusedValueWithoutDrop) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(21)
	if err := serializer.SerializeU16(obj.ResultIdx); err != nil { return err }
	if err := serializer.SerializeU16(obj.SecondaryIdx); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *ExecutionFailureStatus__UnusedValueWithoutDrop) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_ExecutionFailureStatus__UnusedValueWithoutDrop(deserializer serde.Deserializer) (ExecutionFailureStatus__UnusedValueWithoutDrop, error) {
	var obj ExecutionFailureStatus__UnusedValueWithoutDrop
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := deserializer.DeserializeU16(); err == nil { obj.ResultIdx = val } else { return obj, err }
	if val, err := deserializer.DeserializeU16(); err == nil { obj.SecondaryIdx = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type ExecutionFailureStatus__InvalidPublicFunctionReturnType struct {
	Idx uint16
}

func (*ExecutionFailureStatus__InvalidPublicFunctionReturnType) isExecutionFailureStatus() {}

func (obj *ExecutionFailureStatus__InvalidPublicFunctionReturnType) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(22)
	if err := serializer.SerializeU16(obj.Idx); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *ExecutionFailureStatus__InvalidPublicFunctionReturnType) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_ExecutionFailureStatus__InvalidPublicFunctionReturnType(deserializer serde.Deserializer) (ExecutionFailureStatus__InvalidPublicFunctionReturnType, error) {
	var obj ExecutionFailureStatus__InvalidPublicFunctionReturnType
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := deserializer.DeserializeU16(); err == nil { obj.Idx = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type ExecutionFailureStatus__InvalidTransferObject struct {
}

func (*ExecutionFailureStatus__InvalidTransferObject) isExecutionFailureStatus() {}

func (obj *ExecutionFailureStatus__InvalidTransferObject) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(23)
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *ExecutionFailureStatus__InvalidTransferObject) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_ExecutionFailureStatus__InvalidTransferObject(deserializer serde.Deserializer) (ExecutionFailureStatus__InvalidTransferObject, error) {
	var obj ExecutionFailureStatus__InvalidTransferObject
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type ExecutionFailureStatus__EffectsTooLarge struct {
	CurrentSize uint64
	MaxSize uint64
}

func (*ExecutionFailureStatus__EffectsTooLarge) isExecutionFailureStatus() {}

func (obj *ExecutionFailureStatus__EffectsTooLarge) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(24)
	if err := serializer.SerializeU64(obj.CurrentSize); err != nil { return err }
	if err := serializer.SerializeU64(obj.MaxSize); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *ExecutionFailureStatus__EffectsTooLarge) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_ExecutionFailureStatus__EffectsTooLarge(deserializer serde.Deserializer) (ExecutionFailureStatus__EffectsTooLarge, error) {
	var obj ExecutionFailureStatus__EffectsTooLarge
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := deserializer.DeserializeU64(); err == nil { obj.CurrentSize = val } else { return obj, err }
	if val, err := deserializer.DeserializeU64(); err == nil { obj.MaxSize = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type ExecutionFailureStatus__PublishUpgradeMissingDependency struct {
}

func (*ExecutionFailureStatus__PublishUpgradeMissingDependency) isExecutionFailureStatus() {}

func (obj *ExecutionFailureStatus__PublishUpgradeMissingDependency) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(25)
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *ExecutionFailureStatus__PublishUpgradeMissingDependency) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_ExecutionFailureStatus__PublishUpgradeMissingDependency(deserializer serde.Deserializer) (ExecutionFailureStatus__PublishUpgradeMissingDependency, error) {
	var obj ExecutionFailureStatus__PublishUpgradeMissingDependency
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type ExecutionFailureStatus__PublishUpgradeDependencyDowngrade struct {
}

func (*ExecutionFailureStatus__PublishUpgradeDependencyDowngrade) isExecutionFailureStatus() {}

func (obj *ExecutionFailureStatus__PublishUpgradeDependencyDowngrade) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(26)
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *ExecutionFailureStatus__PublishUpgradeDependencyDowngrade) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_ExecutionFailureStatus__PublishUpgradeDependencyDowngrade(deserializer serde.Deserializer) (ExecutionFailureStatus__PublishUpgradeDependencyDowngrade, error) {
	var obj ExecutionFailureStatus__PublishUpgradeDependencyDowngrade
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type ExecutionFailureStatus__PackageUpgradeError struct {
	UpgradeError PackageUpgradeError
}

func (*ExecutionFailureStatus__PackageUpgradeError) isExecutionFailureStatus() {}

func (obj *ExecutionFailureStatus__PackageUpgradeError) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(27)
	if err := obj.UpgradeError.Serialize(serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *ExecutionFailureStatus__PackageUpgradeError) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_ExecutionFailureStatus__PackageUpgradeError(deserializer serde.Deserializer) (ExecutionFailureStatus__PackageUpgradeError, error) {
	var obj ExecutionFailureStatus__PackageUpgradeError
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := DeserializePackageUpgradeError(deserializer); err == nil { obj.UpgradeError = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type ExecutionFailureStatus__WrittenObjectsTooLarge struct {
	CurrentSize uint64
	MaxSize uint64
}

func (*ExecutionFailureStatus__WrittenObjectsTooLarge) isExecutionFailureStatus() {}

func (obj *ExecutionFailureStatus__WrittenObjectsTooLarge) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(28)
	if err := serializer.SerializeU64(obj.CurrentSize); err != nil { return err }
	if err := serializer.SerializeU64(obj.MaxSize); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *ExecutionFailureStatus__WrittenObjectsTooLarge) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_ExecutionFailureStatus__WrittenObjectsTooLarge(deserializer serde.Deserializer) (ExecutionFailureStatus__WrittenObjectsTooLarge, error) {
	var obj ExecutionFailureStatus__WrittenObjectsTooLarge
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := deserializer.DeserializeU64(); err == nil { obj.CurrentSize = val } else { return obj, err }
	if val, err := deserializer.DeserializeU64(); err == nil { obj.MaxSize = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type ExecutionFailureStatus__CertificateDenied struct {
}

func (*ExecutionFailureStatus__CertificateDenied) isExecutionFailureStatus() {}

func (obj *ExecutionFailureStatus__CertificateDenied) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(29)
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *ExecutionFailureStatus__CertificateDenied) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_ExecutionFailureStatus__CertificateDenied(deserializer serde.Deserializer) (ExecutionFailureStatus__CertificateDenied, error) {
	var obj ExecutionFailureStatus__CertificateDenied
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type ExecutionStatus interface {
	isExecutionStatus()
	Serialize(serializer serde.Serializer) error
	BcsSerialize() ([]byte, error)
}

func DeserializeExecutionStatus(deserializer serde.Deserializer) (ExecutionStatus, error) {
	index, err := deserializer.DeserializeVariantIndex()
	if err != nil { return nil, err }

	switch index {
	case 0:
		if val, err := load_ExecutionStatus__Success(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 1:
		if val, err := load_ExecutionStatus__Failure(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	default:
		return nil, fmt.Errorf("Unknown variant index for ExecutionStatus: %d", index)
	}
}

func BcsDeserializeExecutionStatus(input []byte) (ExecutionStatus, error) {
	if input == nil {
		var obj ExecutionStatus
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input);
	obj, err := DeserializeExecutionStatus(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type ExecutionStatus__Success struct {
}

func (*ExecutionStatus__Success) isExecutionStatus() {}

func (obj *ExecutionStatus__Success) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(0)
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *ExecutionStatus__Success) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_ExecutionStatus__Success(deserializer serde.Deserializer) (ExecutionStatus__Success, error) {
	var obj ExecutionStatus__Success
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type ExecutionStatus__Failure struct {
	Error ExecutionFailureStatus
	Command *uint64
}

func (*ExecutionStatus__Failure) isExecutionStatus() {}

func (obj *ExecutionStatus__Failure) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(1)
	if err := obj.Error.Serialize(serializer); err != nil { return err }
	if err := serialize_option_u64(obj.Command, serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *ExecutionStatus__Failure) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_ExecutionStatus__Failure(deserializer serde.Deserializer) (ExecutionStatus__Failure, error) {
	var obj ExecutionStatus__Failure
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := DeserializeExecutionFailureStatus(deserializer); err == nil { obj.Error = val } else { return obj, err }
	if val, err := deserialize_option_u64(deserializer); err == nil { obj.Command = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type FullCheckpointContents struct {
	Transactions []ExecutionData
	UserSignatures [][]GenericSignature
}

func (obj *FullCheckpointContents) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	if err := serialize_vector_ExecutionData(obj.Transactions, serializer); err != nil { return err }
	if err := serialize_vector_vector_GenericSignature(obj.UserSignatures, serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *FullCheckpointContents) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func DeserializeFullCheckpointContents(deserializer serde.Deserializer) (FullCheckpointContents, error) {
	var obj FullCheckpointContents
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := deserialize_vector_ExecutionData(deserializer); err == nil { obj.Transactions = val } else { return obj, err }
	if val, err := deserialize_vector_vector_GenericSignature(deserializer); err == nil { obj.UserSignatures = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

func BcsDeserializeFullCheckpointContents(input []byte) (FullCheckpointContents, error) {
	if input == nil {
		var obj FullCheckpointContents
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input);
	obj, err := DeserializeFullCheckpointContents(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type GasCostSummary struct {
	ComputationCost uint64
	StorageCost uint64
	StorageRebate uint64
	NonRefundableStorageFee uint64
}

func (obj *GasCostSummary) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	if err := serializer.SerializeU64(obj.ComputationCost); err != nil { return err }
	if err := serializer.SerializeU64(obj.StorageCost); err != nil { return err }
	if err := serializer.SerializeU64(obj.StorageRebate); err != nil { return err }
	if err := serializer.SerializeU64(obj.NonRefundableStorageFee); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *GasCostSummary) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func DeserializeGasCostSummary(deserializer serde.Deserializer) (GasCostSummary, error) {
	var obj GasCostSummary
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := deserializer.DeserializeU64(); err == nil { obj.ComputationCost = val } else { return obj, err }
	if val, err := deserializer.DeserializeU64(); err == nil { obj.StorageCost = val } else { return obj, err }
	if val, err := deserializer.DeserializeU64(); err == nil { obj.StorageRebate = val } else { return obj, err }
	if val, err := deserializer.DeserializeU64(); err == nil { obj.NonRefundableStorageFee = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

func BcsDeserializeGasCostSummary(input []byte) (GasCostSummary, error) {
	if input == nil {
		var obj GasCostSummary
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input);
	obj, err := DeserializeGasCostSummary(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type GasData struct {
	Payment []struct {Field0 ObjectID; Field1 SequenceNumber; Field2 ObjectDigest}
	Owner SuiAddress
	Price uint64
	Budget uint64
}

func (obj *GasData) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	if err := serialize_vector_tuple3_ObjectID_SequenceNumber_ObjectDigest(obj.Payment, serializer); err != nil { return err }
	if err := obj.Owner.Serialize(serializer); err != nil { return err }
	if err := serializer.SerializeU64(obj.Price); err != nil { return err }
	if err := serializer.SerializeU64(obj.Budget); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *GasData) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func DeserializeGasData(deserializer serde.Deserializer) (GasData, error) {
	var obj GasData
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := deserialize_vector_tuple3_ObjectID_SequenceNumber_ObjectDigest(deserializer); err == nil { obj.Payment = val } else { return obj, err }
	if val, err := DeserializeSuiAddress(deserializer); err == nil { obj.Owner = val } else { return obj, err }
	if val, err := deserializer.DeserializeU64(); err == nil { obj.Price = val } else { return obj, err }
	if val, err := deserializer.DeserializeU64(); err == nil { obj.Budget = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

func BcsDeserializeGasData(input []byte) (GasData, error) {
	if input == nil {
		var obj GasData
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input);
	obj, err := DeserializeGasData(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type GenericSignature []uint8

func (obj *GenericSignature) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	if err := serialize_vector_u8((([]uint8)(*obj)), serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *GenericSignature) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func DeserializeGenericSignature(deserializer serde.Deserializer) (GenericSignature, error) {
	var obj []uint8
	if err := deserializer.IncreaseContainerDepth(); err != nil { return (GenericSignature)(obj), err }
	if val, err := deserialize_vector_u8(deserializer); err == nil { obj = val } else { return ((GenericSignature)(obj)), err }
	deserializer.DecreaseContainerDepth()
	return (GenericSignature)(obj), nil
}

func BcsDeserializeGenericSignature(input []byte) (GenericSignature, error) {
	if input == nil {
		var obj GenericSignature
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input);
	obj, err := DeserializeGenericSignature(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type GenesisObject interface {
	isGenesisObject()
	Serialize(serializer serde.Serializer) error
	BcsSerialize() ([]byte, error)
}

func DeserializeGenesisObject(deserializer serde.Deserializer) (GenesisObject, error) {
	index, err := deserializer.DeserializeVariantIndex()
	if err != nil { return nil, err }

	switch index {
	case 0:
		if val, err := load_GenesisObject__RawObject(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	default:
		return nil, fmt.Errorf("Unknown variant index for GenesisObject: %d", index)
	}
}

func BcsDeserializeGenesisObject(input []byte) (GenesisObject, error) {
	if input == nil {
		var obj GenesisObject
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input);
	obj, err := DeserializeGenesisObject(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type GenesisObject__RawObject struct {
	Data Data
	Owner Owner
}

func (*GenesisObject__RawObject) isGenesisObject() {}

func (obj *GenesisObject__RawObject) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(0)
	if err := obj.Data.Serialize(serializer); err != nil { return err }
	if err := obj.Owner.Serialize(serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *GenesisObject__RawObject) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_GenesisObject__RawObject(deserializer serde.Deserializer) (GenesisObject__RawObject, error) {
	var obj GenesisObject__RawObject
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := DeserializeData(deserializer); err == nil { obj.Data = val } else { return obj, err }
	if val, err := DeserializeOwner(deserializer); err == nil { obj.Owner = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type GenesisTransaction struct {
	Objects []GenesisObject
}

func (obj *GenesisTransaction) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	if err := serialize_vector_GenesisObject(obj.Objects, serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *GenesisTransaction) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func DeserializeGenesisTransaction(deserializer serde.Deserializer) (GenesisTransaction, error) {
	var obj GenesisTransaction
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := deserialize_vector_GenesisObject(deserializer); err == nil { obj.Objects = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

func BcsDeserializeGenesisTransaction(input []byte) (GenesisTransaction, error) {
	if input == nil {
		var obj GenesisTransaction
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input);
	obj, err := DeserializeGenesisTransaction(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type Identifier string

func (obj *Identifier) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	if err := serializer.SerializeStr(((string)(*obj))); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *Identifier) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func DeserializeIdentifier(deserializer serde.Deserializer) (Identifier, error) {
	var obj string
	if err := deserializer.IncreaseContainerDepth(); err != nil { return (Identifier)(obj), err }
	if val, err := deserializer.DeserializeStr(); err == nil { obj = val } else { return ((Identifier)(obj)), err }
	deserializer.DecreaseContainerDepth()
	return (Identifier)(obj), nil
}

func BcsDeserializeIdentifier(input []byte) (Identifier, error) {
	if input == nil {
		var obj Identifier
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input);
	obj, err := DeserializeIdentifier(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type Intent struct {
	Scope uint8
	Version uint8
	AppId uint8
}

func (obj *Intent) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	if err := serializer.SerializeU8(obj.Scope); err != nil { return err }
	if err := serializer.SerializeU8(obj.Version); err != nil { return err }
	if err := serializer.SerializeU8(obj.AppId); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *Intent) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func DeserializeIntent(deserializer serde.Deserializer) (Intent, error) {
	var obj Intent
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := deserializer.DeserializeU8(); err == nil { obj.Scope = val } else { return obj, err }
	if val, err := deserializer.DeserializeU8(); err == nil { obj.Version = val } else { return obj, err }
	if val, err := deserializer.DeserializeU8(); err == nil { obj.AppId = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

func BcsDeserializeIntent(input []byte) (Intent, error) {
	if input == nil {
		var obj Intent
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input);
	obj, err := DeserializeIntent(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type IntentMessage struct {
	Intent Intent
	Value TransactionData
}

func (obj *IntentMessage) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	if err := obj.Intent.Serialize(serializer); err != nil { return err }
	if err := obj.Value.Serialize(serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *IntentMessage) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func DeserializeIntentMessage(deserializer serde.Deserializer) (IntentMessage, error) {
	var obj IntentMessage
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := DeserializeIntent(deserializer); err == nil { obj.Intent = val } else { return obj, err }
	if val, err := DeserializeTransactionData(deserializer); err == nil { obj.Value = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

func BcsDeserializeIntentMessage(input []byte) (IntentMessage, error) {
	if input == nil {
		var obj IntentMessage
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input);
	obj, err := DeserializeIntentMessage(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type ModuleId struct {
	Address AccountAddress
	Name Identifier
}

func (obj *ModuleId) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	if err := obj.Address.Serialize(serializer); err != nil { return err }
	if err := obj.Name.Serialize(serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *ModuleId) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func DeserializeModuleId(deserializer serde.Deserializer) (ModuleId, error) {
	var obj ModuleId
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := DeserializeAccountAddress(deserializer); err == nil { obj.Address = val } else { return obj, err }
	if val, err := DeserializeIdentifier(deserializer); err == nil { obj.Name = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

func BcsDeserializeModuleId(input []byte) (ModuleId, error) {
	if input == nil {
		var obj ModuleId
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input);
	obj, err := DeserializeModuleId(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type MoveFieldLayout struct {
	Name Identifier
	Layout MoveTypeLayout
}

func (obj *MoveFieldLayout) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	if err := obj.Name.Serialize(serializer); err != nil { return err }
	if err := obj.Layout.Serialize(serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *MoveFieldLayout) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func DeserializeMoveFieldLayout(deserializer serde.Deserializer) (MoveFieldLayout, error) {
	var obj MoveFieldLayout
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := DeserializeIdentifier(deserializer); err == nil { obj.Name = val } else { return obj, err }
	if val, err := DeserializeMoveTypeLayout(deserializer); err == nil { obj.Layout = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

func BcsDeserializeMoveFieldLayout(input []byte) (MoveFieldLayout, error) {
	if input == nil {
		var obj MoveFieldLayout
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input);
	obj, err := DeserializeMoveFieldLayout(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type MoveLocation struct {
	Module ModuleId
	Function uint16
	Instruction uint16
	FunctionName *string
}

func (obj *MoveLocation) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	if err := obj.Module.Serialize(serializer); err != nil { return err }
	if err := serializer.SerializeU16(obj.Function); err != nil { return err }
	if err := serializer.SerializeU16(obj.Instruction); err != nil { return err }
	if err := serialize_option_str(obj.FunctionName, serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *MoveLocation) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func DeserializeMoveLocation(deserializer serde.Deserializer) (MoveLocation, error) {
	var obj MoveLocation
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := DeserializeModuleId(deserializer); err == nil { obj.Module = val } else { return obj, err }
	if val, err := deserializer.DeserializeU16(); err == nil { obj.Function = val } else { return obj, err }
	if val, err := deserializer.DeserializeU16(); err == nil { obj.Instruction = val } else { return obj, err }
	if val, err := deserialize_option_str(deserializer); err == nil { obj.FunctionName = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

func BcsDeserializeMoveLocation(input []byte) (MoveLocation, error) {
	if input == nil {
		var obj MoveLocation
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input);
	obj, err := DeserializeMoveLocation(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type MoveLocationOpt struct {
	Value *MoveLocation
}

func (obj *MoveLocationOpt) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	if err := serialize_option_MoveLocation(obj.Value, serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *MoveLocationOpt) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func DeserializeMoveLocationOpt(deserializer serde.Deserializer) (MoveLocationOpt, error) {
	var obj MoveLocationOpt
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := deserialize_option_MoveLocation(deserializer); err == nil { obj.Value = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

func BcsDeserializeMoveLocationOpt(input []byte) (MoveLocationOpt, error) {
	if input == nil {
		var obj MoveLocationOpt
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input);
	obj, err := DeserializeMoveLocationOpt(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type MoveObject struct {
	Type MoveObjectType
	HasPublicTransfer bool
	Version SequenceNumber
	Contents []byte
}

func (obj *MoveObject) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	if err := obj.Type.Serialize(serializer); err != nil { return err }
	if err := serializer.SerializeBool(obj.HasPublicTransfer); err != nil { return err }
	if err := obj.Version.Serialize(serializer); err != nil { return err }
	if err := serializer.SerializeBytes(obj.Contents); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *MoveObject) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func DeserializeMoveObject(deserializer serde.Deserializer) (MoveObject, error) {
	var obj MoveObject
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := DeserializeMoveObjectType(deserializer); err == nil { obj.Type = val } else { return obj, err }
	if val, err := deserializer.DeserializeBool(); err == nil { obj.HasPublicTransfer = val } else { return obj, err }
	if val, err := DeserializeSequenceNumber(deserializer); err == nil { obj.Version = val } else { return obj, err }
	if val, err := deserializer.DeserializeBytes(); err == nil { obj.Contents = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

func BcsDeserializeMoveObject(input []byte) (MoveObject, error) {
	if input == nil {
		var obj MoveObject
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input);
	obj, err := DeserializeMoveObject(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type MoveObjectType struct {
	Value MoveObjectType_
}

func (obj *MoveObjectType) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	if err := obj.Value.Serialize(serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *MoveObjectType) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func DeserializeMoveObjectType(deserializer serde.Deserializer) (MoveObjectType, error) {
	var obj MoveObjectType
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := DeserializeMoveObjectType_(deserializer); err == nil { obj.Value = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

func BcsDeserializeMoveObjectType(input []byte) (MoveObjectType, error) {
	if input == nil {
		var obj MoveObjectType
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input);
	obj, err := DeserializeMoveObjectType(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type MoveObjectType_ interface {
	isMoveObjectType_()
	Serialize(serializer serde.Serializer) error
	BcsSerialize() ([]byte, error)
}

func DeserializeMoveObjectType_(deserializer serde.Deserializer) (MoveObjectType_, error) {
	index, err := deserializer.DeserializeVariantIndex()
	if err != nil { return nil, err }

	switch index {
	case 0:
		if val, err := load_MoveObjectType___Other(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 1:
		if val, err := load_MoveObjectType___GasCoin(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 2:
		if val, err := load_MoveObjectType___StakedSui(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 3:
		if val, err := load_MoveObjectType___Coin(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	default:
		return nil, fmt.Errorf("Unknown variant index for MoveObjectType_: %d", index)
	}
}

func BcsDeserializeMoveObjectType_(input []byte) (MoveObjectType_, error) {
	if input == nil {
		var obj MoveObjectType_
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input);
	obj, err := DeserializeMoveObjectType_(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type MoveObjectType___Other struct {
	Value StructTag
}

func (*MoveObjectType___Other) isMoveObjectType_() {}

func (obj *MoveObjectType___Other) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(0)
	if err := obj.Value.Serialize(serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *MoveObjectType___Other) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_MoveObjectType___Other(deserializer serde.Deserializer) (MoveObjectType___Other, error) {
	var obj MoveObjectType___Other
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := DeserializeStructTag(deserializer); err == nil { obj.Value = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type MoveObjectType___GasCoin struct {
}

func (*MoveObjectType___GasCoin) isMoveObjectType_() {}

func (obj *MoveObjectType___GasCoin) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(1)
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *MoveObjectType___GasCoin) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_MoveObjectType___GasCoin(deserializer serde.Deserializer) (MoveObjectType___GasCoin, error) {
	var obj MoveObjectType___GasCoin
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type MoveObjectType___StakedSui struct {
}

func (*MoveObjectType___StakedSui) isMoveObjectType_() {}

func (obj *MoveObjectType___StakedSui) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(2)
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *MoveObjectType___StakedSui) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_MoveObjectType___StakedSui(deserializer serde.Deserializer) (MoveObjectType___StakedSui, error) {
	var obj MoveObjectType___StakedSui
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type MoveObjectType___Coin struct {
	Value TypeTag
}

func (*MoveObjectType___Coin) isMoveObjectType_() {}

func (obj *MoveObjectType___Coin) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(3)
	if err := obj.Value.Serialize(serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *MoveObjectType___Coin) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_MoveObjectType___Coin(deserializer serde.Deserializer) (MoveObjectType___Coin, error) {
	var obj MoveObjectType___Coin
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := DeserializeTypeTag(deserializer); err == nil { obj.Value = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type MovePackage struct {
	Id ObjectID
	Version SequenceNumber
	ModuleMap map[string][]byte
	TypeOriginTable []TypeOrigin
	LinkageTable map[ObjectID]UpgradeInfo
}

func (obj *MovePackage) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	if err := obj.Id.Serialize(serializer); err != nil { return err }
	if err := obj.Version.Serialize(serializer); err != nil { return err }
	if err := serialize_map_str_to_bytes(obj.ModuleMap, serializer); err != nil { return err }
	if err := serialize_vector_TypeOrigin(obj.TypeOriginTable, serializer); err != nil { return err }
	if err := serialize_map_ObjectID_to_UpgradeInfo(obj.LinkageTable, serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *MovePackage) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func DeserializeMovePackage(deserializer serde.Deserializer) (MovePackage, error) {
	var obj MovePackage
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := DeserializeObjectID(deserializer); err == nil { obj.Id = val } else { return obj, err }
	if val, err := DeserializeSequenceNumber(deserializer); err == nil { obj.Version = val } else { return obj, err }
	if val, err := deserialize_map_str_to_bytes(deserializer); err == nil { obj.ModuleMap = val } else { return obj, err }
	if val, err := deserialize_vector_TypeOrigin(deserializer); err == nil { obj.TypeOriginTable = val } else { return obj, err }
	if val, err := deserialize_map_ObjectID_to_UpgradeInfo(deserializer); err == nil { obj.LinkageTable = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

func BcsDeserializeMovePackage(input []byte) (MovePackage, error) {
	if input == nil {
		var obj MovePackage
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input);
	obj, err := DeserializeMovePackage(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type MoveStructLayout interface {
	isMoveStructLayout()
	Serialize(serializer serde.Serializer) error
	BcsSerialize() ([]byte, error)
}

func DeserializeMoveStructLayout(deserializer serde.Deserializer) (MoveStructLayout, error) {
	index, err := deserializer.DeserializeVariantIndex()
	if err != nil { return nil, err }

	switch index {
	case 0:
		if val, err := load_MoveStructLayout__Runtime(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 1:
		if val, err := load_MoveStructLayout__WithFields(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 2:
		if val, err := load_MoveStructLayout__WithTypes(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	default:
		return nil, fmt.Errorf("Unknown variant index for MoveStructLayout: %d", index)
	}
}

func BcsDeserializeMoveStructLayout(input []byte) (MoveStructLayout, error) {
	if input == nil {
		var obj MoveStructLayout
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input);
	obj, err := DeserializeMoveStructLayout(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type MoveStructLayout__Runtime []MoveTypeLayout

func (*MoveStructLayout__Runtime) isMoveStructLayout() {}

func (obj *MoveStructLayout__Runtime) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(0)
	if err := serialize_vector_MoveTypeLayout((([]MoveTypeLayout)(*obj)), serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *MoveStructLayout__Runtime) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_MoveStructLayout__Runtime(deserializer serde.Deserializer) (MoveStructLayout__Runtime, error) {
	var obj []MoveTypeLayout
	if err := deserializer.IncreaseContainerDepth(); err != nil { return (MoveStructLayout__Runtime)(obj), err }
	if val, err := deserialize_vector_MoveTypeLayout(deserializer); err == nil { obj = val } else { return ((MoveStructLayout__Runtime)(obj)), err }
	deserializer.DecreaseContainerDepth()
	return (MoveStructLayout__Runtime)(obj), nil
}

type MoveStructLayout__WithFields []MoveFieldLayout

func (*MoveStructLayout__WithFields) isMoveStructLayout() {}

func (obj *MoveStructLayout__WithFields) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(1)
	if err := serialize_vector_MoveFieldLayout((([]MoveFieldLayout)(*obj)), serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *MoveStructLayout__WithFields) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_MoveStructLayout__WithFields(deserializer serde.Deserializer) (MoveStructLayout__WithFields, error) {
	var obj []MoveFieldLayout
	if err := deserializer.IncreaseContainerDepth(); err != nil { return (MoveStructLayout__WithFields)(obj), err }
	if val, err := deserialize_vector_MoveFieldLayout(deserializer); err == nil { obj = val } else { return ((MoveStructLayout__WithFields)(obj)), err }
	deserializer.DecreaseContainerDepth()
	return (MoveStructLayout__WithFields)(obj), nil
}

type MoveStructLayout__WithTypes struct {
	Type StructTag
	Fields []MoveFieldLayout
}

func (*MoveStructLayout__WithTypes) isMoveStructLayout() {}

func (obj *MoveStructLayout__WithTypes) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(2)
	if err := obj.Type.Serialize(serializer); err != nil { return err }
	if err := serialize_vector_MoveFieldLayout(obj.Fields, serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *MoveStructLayout__WithTypes) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_MoveStructLayout__WithTypes(deserializer serde.Deserializer) (MoveStructLayout__WithTypes, error) {
	var obj MoveStructLayout__WithTypes
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := DeserializeStructTag(deserializer); err == nil { obj.Type = val } else { return obj, err }
	if val, err := deserialize_vector_MoveFieldLayout(deserializer); err == nil { obj.Fields = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type MoveTypeLayout interface {
	isMoveTypeLayout()
	Serialize(serializer serde.Serializer) error
	BcsSerialize() ([]byte, error)
}

func DeserializeMoveTypeLayout(deserializer serde.Deserializer) (MoveTypeLayout, error) {
	index, err := deserializer.DeserializeVariantIndex()
	if err != nil { return nil, err }

	switch index {
	case 0:
		if val, err := load_MoveTypeLayout__Bool(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 1:
		if val, err := load_MoveTypeLayout__U8(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 2:
		if val, err := load_MoveTypeLayout__U64(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 3:
		if val, err := load_MoveTypeLayout__U128(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 4:
		if val, err := load_MoveTypeLayout__Address(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 5:
		if val, err := load_MoveTypeLayout__Vector(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 6:
		if val, err := load_MoveTypeLayout__Struct(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 7:
		if val, err := load_MoveTypeLayout__Signer(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 8:
		if val, err := load_MoveTypeLayout__U16(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 9:
		if val, err := load_MoveTypeLayout__U32(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 10:
		if val, err := load_MoveTypeLayout__U256(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	default:
		return nil, fmt.Errorf("Unknown variant index for MoveTypeLayout: %d", index)
	}
}

func BcsDeserializeMoveTypeLayout(input []byte) (MoveTypeLayout, error) {
	if input == nil {
		var obj MoveTypeLayout
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input);
	obj, err := DeserializeMoveTypeLayout(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type MoveTypeLayout__Bool struct {
}

func (*MoveTypeLayout__Bool) isMoveTypeLayout() {}

func (obj *MoveTypeLayout__Bool) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(0)
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *MoveTypeLayout__Bool) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_MoveTypeLayout__Bool(deserializer serde.Deserializer) (MoveTypeLayout__Bool, error) {
	var obj MoveTypeLayout__Bool
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type MoveTypeLayout__U8 struct {
}

func (*MoveTypeLayout__U8) isMoveTypeLayout() {}

func (obj *MoveTypeLayout__U8) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(1)
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *MoveTypeLayout__U8) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_MoveTypeLayout__U8(deserializer serde.Deserializer) (MoveTypeLayout__U8, error) {
	var obj MoveTypeLayout__U8
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type MoveTypeLayout__U64 struct {
}

func (*MoveTypeLayout__U64) isMoveTypeLayout() {}

func (obj *MoveTypeLayout__U64) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(2)
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *MoveTypeLayout__U64) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_MoveTypeLayout__U64(deserializer serde.Deserializer) (MoveTypeLayout__U64, error) {
	var obj MoveTypeLayout__U64
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type MoveTypeLayout__U128 struct {
}

func (*MoveTypeLayout__U128) isMoveTypeLayout() {}

func (obj *MoveTypeLayout__U128) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(3)
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *MoveTypeLayout__U128) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_MoveTypeLayout__U128(deserializer serde.Deserializer) (MoveTypeLayout__U128, error) {
	var obj MoveTypeLayout__U128
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type MoveTypeLayout__Address struct {
}

func (*MoveTypeLayout__Address) isMoveTypeLayout() {}

func (obj *MoveTypeLayout__Address) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(4)
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *MoveTypeLayout__Address) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_MoveTypeLayout__Address(deserializer serde.Deserializer) (MoveTypeLayout__Address, error) {
	var obj MoveTypeLayout__Address
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type MoveTypeLayout__Vector struct {
	Value MoveTypeLayout
}

func (*MoveTypeLayout__Vector) isMoveTypeLayout() {}

func (obj *MoveTypeLayout__Vector) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(5)
	if err := obj.Value.Serialize(serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *MoveTypeLayout__Vector) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_MoveTypeLayout__Vector(deserializer serde.Deserializer) (MoveTypeLayout__Vector, error) {
	var obj MoveTypeLayout__Vector
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := DeserializeMoveTypeLayout(deserializer); err == nil { obj.Value = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type MoveTypeLayout__Struct struct {
	Value MoveStructLayout
}

func (*MoveTypeLayout__Struct) isMoveTypeLayout() {}

func (obj *MoveTypeLayout__Struct) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(6)
	if err := obj.Value.Serialize(serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *MoveTypeLayout__Struct) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_MoveTypeLayout__Struct(deserializer serde.Deserializer) (MoveTypeLayout__Struct, error) {
	var obj MoveTypeLayout__Struct
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := DeserializeMoveStructLayout(deserializer); err == nil { obj.Value = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type MoveTypeLayout__Signer struct {
}

func (*MoveTypeLayout__Signer) isMoveTypeLayout() {}

func (obj *MoveTypeLayout__Signer) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(7)
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *MoveTypeLayout__Signer) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_MoveTypeLayout__Signer(deserializer serde.Deserializer) (MoveTypeLayout__Signer, error) {
	var obj MoveTypeLayout__Signer
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type MoveTypeLayout__U16 struct {
}

func (*MoveTypeLayout__U16) isMoveTypeLayout() {}

func (obj *MoveTypeLayout__U16) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(8)
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *MoveTypeLayout__U16) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_MoveTypeLayout__U16(deserializer serde.Deserializer) (MoveTypeLayout__U16, error) {
	var obj MoveTypeLayout__U16
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type MoveTypeLayout__U32 struct {
}

func (*MoveTypeLayout__U32) isMoveTypeLayout() {}

func (obj *MoveTypeLayout__U32) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(9)
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *MoveTypeLayout__U32) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_MoveTypeLayout__U32(deserializer serde.Deserializer) (MoveTypeLayout__U32, error) {
	var obj MoveTypeLayout__U32
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type MoveTypeLayout__U256 struct {
}

func (*MoveTypeLayout__U256) isMoveTypeLayout() {}

func (obj *MoveTypeLayout__U256) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(10)
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *MoveTypeLayout__U256) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_MoveTypeLayout__U256(deserializer serde.Deserializer) (MoveTypeLayout__U256, error) {
	var obj MoveTypeLayout__U256
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type MultiSig struct {
	Sigs []CompressedSignature
	Bitmap []byte
	MultisigPk MultiSigPublicKey
}

func (obj *MultiSig) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	if err := serialize_vector_CompressedSignature(obj.Sigs, serializer); err != nil { return err }
	if err := serializer.SerializeBytes(obj.Bitmap); err != nil { return err }
	if err := obj.MultisigPk.Serialize(serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *MultiSig) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func DeserializeMultiSig(deserializer serde.Deserializer) (MultiSig, error) {
	var obj MultiSig
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := deserialize_vector_CompressedSignature(deserializer); err == nil { obj.Sigs = val } else { return obj, err }
	if val, err := deserializer.DeserializeBytes(); err == nil { obj.Bitmap = val } else { return obj, err }
	if val, err := DeserializeMultiSigPublicKey(deserializer); err == nil { obj.MultisigPk = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

func BcsDeserializeMultiSig(input []byte) (MultiSig, error) {
	if input == nil {
		var obj MultiSig
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input);
	obj, err := DeserializeMultiSig(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type MultiSigPublicKey struct {
	PkMap []struct {Field0 string; Field1 uint8}
	Threshold uint16
}

func (obj *MultiSigPublicKey) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	if err := serialize_vector_tuple2_str_u8(obj.PkMap, serializer); err != nil { return err }
	if err := serializer.SerializeU16(obj.Threshold); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *MultiSigPublicKey) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func DeserializeMultiSigPublicKey(deserializer serde.Deserializer) (MultiSigPublicKey, error) {
	var obj MultiSigPublicKey
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := deserialize_vector_tuple2_str_u8(deserializer); err == nil { obj.PkMap = val } else { return obj, err }
	if val, err := deserializer.DeserializeU16(); err == nil { obj.Threshold = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

func BcsDeserializeMultiSigPublicKey(input []byte) (MultiSigPublicKey, error) {
	if input == nil {
		var obj MultiSigPublicKey
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input);
	obj, err := DeserializeMultiSigPublicKey(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type ObjectArg interface {
	isObjectArg()
	Serialize(serializer serde.Serializer) error
	BcsSerialize() ([]byte, error)
}

func DeserializeObjectArg(deserializer serde.Deserializer) (ObjectArg, error) {
	index, err := deserializer.DeserializeVariantIndex()
	if err != nil { return nil, err }

	switch index {
	case 0:
		if val, err := load_ObjectArg__ImmOrOwnedObject(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 1:
		if val, err := load_ObjectArg__SharedObject(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	default:
		return nil, fmt.Errorf("Unknown variant index for ObjectArg: %d", index)
	}
}

func BcsDeserializeObjectArg(input []byte) (ObjectArg, error) {
	if input == nil {
		var obj ObjectArg
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input);
	obj, err := DeserializeObjectArg(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type ObjectArg__ImmOrOwnedObject struct {Field0 ObjectID; Field1 SequenceNumber; Field2 ObjectDigest}

func (*ObjectArg__ImmOrOwnedObject) isObjectArg() {}

func (obj *ObjectArg__ImmOrOwnedObject) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(0)
	if err := serialize_tuple3_ObjectID_SequenceNumber_ObjectDigest(((struct {Field0 ObjectID; Field1 SequenceNumber; Field2 ObjectDigest})(*obj)), serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *ObjectArg__ImmOrOwnedObject) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_ObjectArg__ImmOrOwnedObject(deserializer serde.Deserializer) (ObjectArg__ImmOrOwnedObject, error) {
	var obj struct {Field0 ObjectID; Field1 SequenceNumber; Field2 ObjectDigest}
	if err := deserializer.IncreaseContainerDepth(); err != nil { return (ObjectArg__ImmOrOwnedObject)(obj), err }
	if val, err := deserialize_tuple3_ObjectID_SequenceNumber_ObjectDigest(deserializer); err == nil { obj = val } else { return ((ObjectArg__ImmOrOwnedObject)(obj)), err }
	deserializer.DecreaseContainerDepth()
	return (ObjectArg__ImmOrOwnedObject)(obj), nil
}

type ObjectArg__SharedObject struct {
	Id ObjectID
	InitialSharedVersion SequenceNumber
	Mutable bool
}

func (*ObjectArg__SharedObject) isObjectArg() {}

func (obj *ObjectArg__SharedObject) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(1)
	if err := obj.Id.Serialize(serializer); err != nil { return err }
	if err := obj.InitialSharedVersion.Serialize(serializer); err != nil { return err }
	if err := serializer.SerializeBool(obj.Mutable); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *ObjectArg__SharedObject) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_ObjectArg__SharedObject(deserializer serde.Deserializer) (ObjectArg__SharedObject, error) {
	var obj ObjectArg__SharedObject
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := DeserializeObjectID(deserializer); err == nil { obj.Id = val } else { return obj, err }
	if val, err := DeserializeSequenceNumber(deserializer); err == nil { obj.InitialSharedVersion = val } else { return obj, err }
	if val, err := deserializer.DeserializeBool(); err == nil { obj.Mutable = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type ObjectDigest struct {
	Value Digest
}

func (obj *ObjectDigest) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	if err := obj.Value.Serialize(serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *ObjectDigest) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func DeserializeObjectDigest(deserializer serde.Deserializer) (ObjectDigest, error) {
	var obj ObjectDigest
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := DeserializeDigest(deserializer); err == nil { obj.Value = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

func BcsDeserializeObjectDigest(input []byte) (ObjectDigest, error) {
	if input == nil {
		var obj ObjectDigest
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input);
	obj, err := DeserializeObjectDigest(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type ObjectID struct {
	Value AccountAddress
}

func (obj *ObjectID) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	if err := obj.Value.Serialize(serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *ObjectID) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func DeserializeObjectID(deserializer serde.Deserializer) (ObjectID, error) {
	var obj ObjectID
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := DeserializeAccountAddress(deserializer); err == nil { obj.Value = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

func BcsDeserializeObjectID(input []byte) (ObjectID, error) {
	if input == nil {
		var obj ObjectID
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input);
	obj, err := DeserializeObjectID(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type ObjectInfoRequestKind interface {
	isObjectInfoRequestKind()
	Serialize(serializer serde.Serializer) error
	BcsSerialize() ([]byte, error)
}

func DeserializeObjectInfoRequestKind(deserializer serde.Deserializer) (ObjectInfoRequestKind, error) {
	index, err := deserializer.DeserializeVariantIndex()
	if err != nil { return nil, err }

	switch index {
	case 0:
		if val, err := load_ObjectInfoRequestKind__LatestObjectInfo(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 1:
		if val, err := load_ObjectInfoRequestKind__PastObjectInfoDebug(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	default:
		return nil, fmt.Errorf("Unknown variant index for ObjectInfoRequestKind: %d", index)
	}
}

func BcsDeserializeObjectInfoRequestKind(input []byte) (ObjectInfoRequestKind, error) {
	if input == nil {
		var obj ObjectInfoRequestKind
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input);
	obj, err := DeserializeObjectInfoRequestKind(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type ObjectInfoRequestKind__LatestObjectInfo struct {
}

func (*ObjectInfoRequestKind__LatestObjectInfo) isObjectInfoRequestKind() {}

func (obj *ObjectInfoRequestKind__LatestObjectInfo) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(0)
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *ObjectInfoRequestKind__LatestObjectInfo) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_ObjectInfoRequestKind__LatestObjectInfo(deserializer serde.Deserializer) (ObjectInfoRequestKind__LatestObjectInfo, error) {
	var obj ObjectInfoRequestKind__LatestObjectInfo
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type ObjectInfoRequestKind__PastObjectInfoDebug struct {
	Value SequenceNumber
}

func (*ObjectInfoRequestKind__PastObjectInfoDebug) isObjectInfoRequestKind() {}

func (obj *ObjectInfoRequestKind__PastObjectInfoDebug) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(1)
	if err := obj.Value.Serialize(serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *ObjectInfoRequestKind__PastObjectInfoDebug) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_ObjectInfoRequestKind__PastObjectInfoDebug(deserializer serde.Deserializer) (ObjectInfoRequestKind__PastObjectInfoDebug, error) {
	var obj ObjectInfoRequestKind__PastObjectInfoDebug
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := DeserializeSequenceNumber(deserializer); err == nil { obj.Value = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type Owner interface {
	isOwner()
	Serialize(serializer serde.Serializer) error
	BcsSerialize() ([]byte, error)
}

func DeserializeOwner(deserializer serde.Deserializer) (Owner, error) {
	index, err := deserializer.DeserializeVariantIndex()
	if err != nil { return nil, err }

	switch index {
	case 0:
		if val, err := load_Owner__AddressOwner(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 1:
		if val, err := load_Owner__ObjectOwner(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 2:
		if val, err := load_Owner__Shared(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 3:
		if val, err := load_Owner__Immutable(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	default:
		return nil, fmt.Errorf("Unknown variant index for Owner: %d", index)
	}
}

func BcsDeserializeOwner(input []byte) (Owner, error) {
	if input == nil {
		var obj Owner
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input);
	obj, err := DeserializeOwner(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type Owner__AddressOwner struct {
	Value SuiAddress
}

func (*Owner__AddressOwner) isOwner() {}

func (obj *Owner__AddressOwner) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(0)
	if err := obj.Value.Serialize(serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *Owner__AddressOwner) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_Owner__AddressOwner(deserializer serde.Deserializer) (Owner__AddressOwner, error) {
	var obj Owner__AddressOwner
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := DeserializeSuiAddress(deserializer); err == nil { obj.Value = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type Owner__ObjectOwner struct {
	Value SuiAddress
}

func (*Owner__ObjectOwner) isOwner() {}

func (obj *Owner__ObjectOwner) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(1)
	if err := obj.Value.Serialize(serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *Owner__ObjectOwner) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_Owner__ObjectOwner(deserializer serde.Deserializer) (Owner__ObjectOwner, error) {
	var obj Owner__ObjectOwner
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := DeserializeSuiAddress(deserializer); err == nil { obj.Value = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type Owner__Shared struct {
	InitialSharedVersion SequenceNumber
}

func (*Owner__Shared) isOwner() {}

func (obj *Owner__Shared) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(2)
	if err := obj.InitialSharedVersion.Serialize(serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *Owner__Shared) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_Owner__Shared(deserializer serde.Deserializer) (Owner__Shared, error) {
	var obj Owner__Shared
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := DeserializeSequenceNumber(deserializer); err == nil { obj.InitialSharedVersion = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type Owner__Immutable struct {
}

func (*Owner__Immutable) isOwner() {}

func (obj *Owner__Immutable) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(3)
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *Owner__Immutable) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_Owner__Immutable(deserializer serde.Deserializer) (Owner__Immutable, error) {
	var obj Owner__Immutable
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type PackageUpgradeError interface {
	isPackageUpgradeError()
	Serialize(serializer serde.Serializer) error
	BcsSerialize() ([]byte, error)
}

func DeserializePackageUpgradeError(deserializer serde.Deserializer) (PackageUpgradeError, error) {
	index, err := deserializer.DeserializeVariantIndex()
	if err != nil { return nil, err }

	switch index {
	case 0:
		if val, err := load_PackageUpgradeError__UnableToFetchPackage(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 1:
		if val, err := load_PackageUpgradeError__NotAPackage(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 2:
		if val, err := load_PackageUpgradeError__IncompatibleUpgrade(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 3:
		if val, err := load_PackageUpgradeError__DigestDoesNotMatch(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 4:
		if val, err := load_PackageUpgradeError__UnknownUpgradePolicy(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 5:
		if val, err := load_PackageUpgradeError__PackageIdDoesNotMatch(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	default:
		return nil, fmt.Errorf("Unknown variant index for PackageUpgradeError: %d", index)
	}
}

func BcsDeserializePackageUpgradeError(input []byte) (PackageUpgradeError, error) {
	if input == nil {
		var obj PackageUpgradeError
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input);
	obj, err := DeserializePackageUpgradeError(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type PackageUpgradeError__UnableToFetchPackage struct {
	PackageId ObjectID
}

func (*PackageUpgradeError__UnableToFetchPackage) isPackageUpgradeError() {}

func (obj *PackageUpgradeError__UnableToFetchPackage) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(0)
	if err := obj.PackageId.Serialize(serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *PackageUpgradeError__UnableToFetchPackage) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_PackageUpgradeError__UnableToFetchPackage(deserializer serde.Deserializer) (PackageUpgradeError__UnableToFetchPackage, error) {
	var obj PackageUpgradeError__UnableToFetchPackage
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := DeserializeObjectID(deserializer); err == nil { obj.PackageId = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type PackageUpgradeError__NotAPackage struct {
	ObjectId ObjectID
}

func (*PackageUpgradeError__NotAPackage) isPackageUpgradeError() {}

func (obj *PackageUpgradeError__NotAPackage) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(1)
	if err := obj.ObjectId.Serialize(serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *PackageUpgradeError__NotAPackage) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_PackageUpgradeError__NotAPackage(deserializer serde.Deserializer) (PackageUpgradeError__NotAPackage, error) {
	var obj PackageUpgradeError__NotAPackage
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := DeserializeObjectID(deserializer); err == nil { obj.ObjectId = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type PackageUpgradeError__IncompatibleUpgrade struct {
}

func (*PackageUpgradeError__IncompatibleUpgrade) isPackageUpgradeError() {}

func (obj *PackageUpgradeError__IncompatibleUpgrade) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(2)
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *PackageUpgradeError__IncompatibleUpgrade) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_PackageUpgradeError__IncompatibleUpgrade(deserializer serde.Deserializer) (PackageUpgradeError__IncompatibleUpgrade, error) {
	var obj PackageUpgradeError__IncompatibleUpgrade
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type PackageUpgradeError__DigestDoesNotMatch struct {
	Digest []uint8
}

func (*PackageUpgradeError__DigestDoesNotMatch) isPackageUpgradeError() {}

func (obj *PackageUpgradeError__DigestDoesNotMatch) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(3)
	if err := serialize_vector_u8(obj.Digest, serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *PackageUpgradeError__DigestDoesNotMatch) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_PackageUpgradeError__DigestDoesNotMatch(deserializer serde.Deserializer) (PackageUpgradeError__DigestDoesNotMatch, error) {
	var obj PackageUpgradeError__DigestDoesNotMatch
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := deserialize_vector_u8(deserializer); err == nil { obj.Digest = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type PackageUpgradeError__UnknownUpgradePolicy struct {
	Policy uint8
}

func (*PackageUpgradeError__UnknownUpgradePolicy) isPackageUpgradeError() {}

func (obj *PackageUpgradeError__UnknownUpgradePolicy) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(4)
	if err := serializer.SerializeU8(obj.Policy); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *PackageUpgradeError__UnknownUpgradePolicy) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_PackageUpgradeError__UnknownUpgradePolicy(deserializer serde.Deserializer) (PackageUpgradeError__UnknownUpgradePolicy, error) {
	var obj PackageUpgradeError__UnknownUpgradePolicy
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := deserializer.DeserializeU8(); err == nil { obj.Policy = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type PackageUpgradeError__PackageIdDoesNotMatch struct {
	PackageId ObjectID
	TicketId ObjectID
}

func (*PackageUpgradeError__PackageIdDoesNotMatch) isPackageUpgradeError() {}

func (obj *PackageUpgradeError__PackageIdDoesNotMatch) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(5)
	if err := obj.PackageId.Serialize(serializer); err != nil { return err }
	if err := obj.TicketId.Serialize(serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *PackageUpgradeError__PackageIdDoesNotMatch) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_PackageUpgradeError__PackageIdDoesNotMatch(deserializer serde.Deserializer) (PackageUpgradeError__PackageIdDoesNotMatch, error) {
	var obj PackageUpgradeError__PackageIdDoesNotMatch
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := DeserializeObjectID(deserializer); err == nil { obj.PackageId = val } else { return obj, err }
	if val, err := DeserializeObjectID(deserializer); err == nil { obj.TicketId = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type ProgrammableMoveCall struct {
	Package ObjectID
	Module Identifier
	Function Identifier
	TypeArguments []TypeTag
	Arguments []Argument
}

func (obj *ProgrammableMoveCall) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	if err := obj.Package.Serialize(serializer); err != nil { return err }
	if err := obj.Module.Serialize(serializer); err != nil { return err }
	if err := obj.Function.Serialize(serializer); err != nil { return err }
	if err := serialize_vector_TypeTag(obj.TypeArguments, serializer); err != nil { return err }
	if err := serialize_vector_Argument(obj.Arguments, serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *ProgrammableMoveCall) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func DeserializeProgrammableMoveCall(deserializer serde.Deserializer) (ProgrammableMoveCall, error) {
	var obj ProgrammableMoveCall
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := DeserializeObjectID(deserializer); err == nil { obj.Package = val } else { return obj, err }
	if val, err := DeserializeIdentifier(deserializer); err == nil { obj.Module = val } else { return obj, err }
	if val, err := DeserializeIdentifier(deserializer); err == nil { obj.Function = val } else { return obj, err }
	if val, err := deserialize_vector_TypeTag(deserializer); err == nil { obj.TypeArguments = val } else { return obj, err }
	if val, err := deserialize_vector_Argument(deserializer); err == nil { obj.Arguments = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

func BcsDeserializeProgrammableMoveCall(input []byte) (ProgrammableMoveCall, error) {
	if input == nil {
		var obj ProgrammableMoveCall
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input);
	obj, err := DeserializeProgrammableMoveCall(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type ProgrammableTransaction struct {
	Inputs []CallArg
	Commands []Command
}

func (obj *ProgrammableTransaction) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	if err := serialize_vector_CallArg(obj.Inputs, serializer); err != nil { return err }
	if err := serialize_vector_Command(obj.Commands, serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *ProgrammableTransaction) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func DeserializeProgrammableTransaction(deserializer serde.Deserializer) (ProgrammableTransaction, error) {
	var obj ProgrammableTransaction
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := deserialize_vector_CallArg(deserializer); err == nil { obj.Inputs = val } else { return obj, err }
	if val, err := deserialize_vector_Command(deserializer); err == nil { obj.Commands = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

func BcsDeserializeProgrammableTransaction(input []byte) (ProgrammableTransaction, error) {
	if input == nil {
		var obj ProgrammableTransaction
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input);
	obj, err := DeserializeProgrammableTransaction(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type ProtocolVersion uint64

func (obj *ProtocolVersion) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	if err := serializer.SerializeU64(((uint64)(*obj))); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *ProtocolVersion) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func DeserializeProtocolVersion(deserializer serde.Deserializer) (ProtocolVersion, error) {
	var obj uint64
	if err := deserializer.IncreaseContainerDepth(); err != nil { return (ProtocolVersion)(obj), err }
	if val, err := deserializer.DeserializeU64(); err == nil { obj = val } else { return ((ProtocolVersion)(obj)), err }
	deserializer.DecreaseContainerDepth()
	return (ProtocolVersion)(obj), nil
}

func BcsDeserializeProtocolVersion(input []byte) (ProtocolVersion, error) {
	if input == nil {
		var obj ProtocolVersion
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input);
	obj, err := DeserializeProtocolVersion(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type SenderSignedData []SenderSignedTransaction

func (obj *SenderSignedData) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	if err := serialize_vector_SenderSignedTransaction((([]SenderSignedTransaction)(*obj)), serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *SenderSignedData) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func DeserializeSenderSignedData(deserializer serde.Deserializer) (SenderSignedData, error) {
	var obj []SenderSignedTransaction
	if err := deserializer.IncreaseContainerDepth(); err != nil { return (SenderSignedData)(obj), err }
	if val, err := deserialize_vector_SenderSignedTransaction(deserializer); err == nil { obj = val } else { return ((SenderSignedData)(obj)), err }
	deserializer.DecreaseContainerDepth()
	return (SenderSignedData)(obj), nil
}

func BcsDeserializeSenderSignedData(input []byte) (SenderSignedData, error) {
	if input == nil {
		var obj SenderSignedData
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input);
	obj, err := DeserializeSenderSignedData(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type SenderSignedTransaction struct {
	IntentMessage IntentMessage
	TxSignatures []GenericSignature
}

func (obj *SenderSignedTransaction) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	if err := obj.IntentMessage.Serialize(serializer); err != nil { return err }
	if err := serialize_vector_GenericSignature(obj.TxSignatures, serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *SenderSignedTransaction) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func DeserializeSenderSignedTransaction(deserializer serde.Deserializer) (SenderSignedTransaction, error) {
	var obj SenderSignedTransaction
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := DeserializeIntentMessage(deserializer); err == nil { obj.IntentMessage = val } else { return obj, err }
	if val, err := deserialize_vector_GenericSignature(deserializer); err == nil { obj.TxSignatures = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

func BcsDeserializeSenderSignedTransaction(input []byte) (SenderSignedTransaction, error) {
	if input == nil {
		var obj SenderSignedTransaction
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input);
	obj, err := DeserializeSenderSignedTransaction(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type SequenceNumber uint64

func (obj *SequenceNumber) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	if err := serializer.SerializeU64(((uint64)(*obj))); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *SequenceNumber) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func DeserializeSequenceNumber(deserializer serde.Deserializer) (SequenceNumber, error) {
	var obj uint64
	if err := deserializer.IncreaseContainerDepth(); err != nil { return (SequenceNumber)(obj), err }
	if val, err := deserializer.DeserializeU64(); err == nil { obj = val } else { return ((SequenceNumber)(obj)), err }
	deserializer.DecreaseContainerDepth()
	return (SequenceNumber)(obj), nil
}

func BcsDeserializeSequenceNumber(input []byte) (SequenceNumber, error) {
	if input == nil {
		var obj SequenceNumber
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input);
	obj, err := DeserializeSequenceNumber(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type StructTag struct {
	Address AccountAddress
	Module Identifier
	Name Identifier
	TypeArgs []TypeTag
}

func (obj *StructTag) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	if err := obj.Address.Serialize(serializer); err != nil { return err }
	if err := obj.Module.Serialize(serializer); err != nil { return err }
	if err := obj.Name.Serialize(serializer); err != nil { return err }
	if err := serialize_vector_TypeTag(obj.TypeArgs, serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *StructTag) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func DeserializeStructTag(deserializer serde.Deserializer) (StructTag, error) {
	var obj StructTag
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := DeserializeAccountAddress(deserializer); err == nil { obj.Address = val } else { return obj, err }
	if val, err := DeserializeIdentifier(deserializer); err == nil { obj.Module = val } else { return obj, err }
	if val, err := DeserializeIdentifier(deserializer); err == nil { obj.Name = val } else { return obj, err }
	if val, err := deserialize_vector_TypeTag(deserializer); err == nil { obj.TypeArgs = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

func BcsDeserializeStructTag(input []byte) (StructTag, error) {
	if input == nil {
		var obj StructTag
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input);
	obj, err := DeserializeStructTag(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type SuiAddress [32]uint8

func (obj *SuiAddress) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	if err := serialize_array32_u8_array((([32]uint8)(*obj)), serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *SuiAddress) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func DeserializeSuiAddress(deserializer serde.Deserializer) (SuiAddress, error) {
	var obj [32]uint8
	if err := deserializer.IncreaseContainerDepth(); err != nil { return (SuiAddress)(obj), err }
	if val, err := deserialize_array32_u8_array(deserializer); err == nil { obj = val } else { return ((SuiAddress)(obj)), err }
	deserializer.DecreaseContainerDepth()
	return (SuiAddress)(obj), nil
}

func BcsDeserializeSuiAddress(input []byte) (SuiAddress, error) {
	if input == nil {
		var obj SuiAddress
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input);
	obj, err := DeserializeSuiAddress(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type TransactionData interface {
	isTransactionData()
	Serialize(serializer serde.Serializer) error
	BcsSerialize() ([]byte, error)
}

func DeserializeTransactionData(deserializer serde.Deserializer) (TransactionData, error) {
	index, err := deserializer.DeserializeVariantIndex()
	if err != nil { return nil, err }

	switch index {
	case 0:
		if val, err := load_TransactionData__V1(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	default:
		return nil, fmt.Errorf("Unknown variant index for TransactionData: %d", index)
	}
}

func BcsDeserializeTransactionData(input []byte) (TransactionData, error) {
	if input == nil {
		var obj TransactionData
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input);
	obj, err := DeserializeTransactionData(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type TransactionData__V1 struct {
	Value TransactionDataV1
}

func (*TransactionData__V1) isTransactionData() {}

func (obj *TransactionData__V1) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(0)
	if err := obj.Value.Serialize(serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *TransactionData__V1) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_TransactionData__V1(deserializer serde.Deserializer) (TransactionData__V1, error) {
	var obj TransactionData__V1
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := DeserializeTransactionDataV1(deserializer); err == nil { obj.Value = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type TransactionDataV1 struct {
	Kind TransactionKind
	Sender SuiAddress
	GasData GasData
	Expiration TransactionExpiration
}

func (obj *TransactionDataV1) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	if err := obj.Kind.Serialize(serializer); err != nil { return err }
	if err := obj.Sender.Serialize(serializer); err != nil { return err }
	if err := obj.GasData.Serialize(serializer); err != nil { return err }
	if err := obj.Expiration.Serialize(serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *TransactionDataV1) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func DeserializeTransactionDataV1(deserializer serde.Deserializer) (TransactionDataV1, error) {
	var obj TransactionDataV1
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := DeserializeTransactionKind(deserializer); err == nil { obj.Kind = val } else { return obj, err }
	if val, err := DeserializeSuiAddress(deserializer); err == nil { obj.Sender = val } else { return obj, err }
	if val, err := DeserializeGasData(deserializer); err == nil { obj.GasData = val } else { return obj, err }
	if val, err := DeserializeTransactionExpiration(deserializer); err == nil { obj.Expiration = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

func BcsDeserializeTransactionDataV1(input []byte) (TransactionDataV1, error) {
	if input == nil {
		var obj TransactionDataV1
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input);
	obj, err := DeserializeTransactionDataV1(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type TransactionDigest struct {
	Value Digest
}

func (obj *TransactionDigest) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	if err := obj.Value.Serialize(serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *TransactionDigest) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func DeserializeTransactionDigest(deserializer serde.Deserializer) (TransactionDigest, error) {
	var obj TransactionDigest
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := DeserializeDigest(deserializer); err == nil { obj.Value = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

func BcsDeserializeTransactionDigest(input []byte) (TransactionDigest, error) {
	if input == nil {
		var obj TransactionDigest
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input);
	obj, err := DeserializeTransactionDigest(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type TransactionEffects interface {
	isTransactionEffects()
	Serialize(serializer serde.Serializer) error
	BcsSerialize() ([]byte, error)
}

func DeserializeTransactionEffects(deserializer serde.Deserializer) (TransactionEffects, error) {
	index, err := deserializer.DeserializeVariantIndex()
	if err != nil { return nil, err }

	switch index {
	case 0:
		if val, err := load_TransactionEffects__V1(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	default:
		return nil, fmt.Errorf("Unknown variant index for TransactionEffects: %d", index)
	}
}

func BcsDeserializeTransactionEffects(input []byte) (TransactionEffects, error) {
	if input == nil {
		var obj TransactionEffects
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input);
	obj, err := DeserializeTransactionEffects(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type TransactionEffects__V1 struct {
	Value TransactionEffectsV1
}

func (*TransactionEffects__V1) isTransactionEffects() {}

func (obj *TransactionEffects__V1) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(0)
	if err := obj.Value.Serialize(serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *TransactionEffects__V1) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_TransactionEffects__V1(deserializer serde.Deserializer) (TransactionEffects__V1, error) {
	var obj TransactionEffects__V1
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := DeserializeTransactionEffectsV1(deserializer); err == nil { obj.Value = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type TransactionEffectsDigest struct {
	Value Digest
}

func (obj *TransactionEffectsDigest) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	if err := obj.Value.Serialize(serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *TransactionEffectsDigest) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func DeserializeTransactionEffectsDigest(deserializer serde.Deserializer) (TransactionEffectsDigest, error) {
	var obj TransactionEffectsDigest
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := DeserializeDigest(deserializer); err == nil { obj.Value = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

func BcsDeserializeTransactionEffectsDigest(input []byte) (TransactionEffectsDigest, error) {
	if input == nil {
		var obj TransactionEffectsDigest
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input);
	obj, err := DeserializeTransactionEffectsDigest(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type TransactionEffectsV1 struct {
	Status ExecutionStatus
	ExecutedEpoch uint64
	GasUsed GasCostSummary
	ModifiedAtVersions []struct {Field0 ObjectID; Field1 SequenceNumber}
	SharedObjects []struct {Field0 ObjectID; Field1 SequenceNumber; Field2 ObjectDigest}
	TransactionDigest TransactionDigest
	Created []struct {Field0 struct {Field0 ObjectID; Field1 SequenceNumber; Field2 ObjectDigest}; Field1 Owner}
	Mutated []struct {Field0 struct {Field0 ObjectID; Field1 SequenceNumber; Field2 ObjectDigest}; Field1 Owner}
	Unwrapped []struct {Field0 struct {Field0 ObjectID; Field1 SequenceNumber; Field2 ObjectDigest}; Field1 Owner}
	Deleted []struct {Field0 ObjectID; Field1 SequenceNumber; Field2 ObjectDigest}
	UnwrappedThenDeleted []struct {Field0 ObjectID; Field1 SequenceNumber; Field2 ObjectDigest}
	Wrapped []struct {Field0 ObjectID; Field1 SequenceNumber; Field2 ObjectDigest}
	GasObject struct {Field0 struct {Field0 ObjectID; Field1 SequenceNumber; Field2 ObjectDigest}; Field1 Owner}
	EventsDigest *TransactionEventsDigest
	Dependencies []TransactionDigest
}

func (obj *TransactionEffectsV1) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	if err := obj.Status.Serialize(serializer); err != nil { return err }
	if err := serializer.SerializeU64(obj.ExecutedEpoch); err != nil { return err }
	if err := obj.GasUsed.Serialize(serializer); err != nil { return err }
	if err := serialize_vector_tuple2_ObjectID_SequenceNumber(obj.ModifiedAtVersions, serializer); err != nil { return err }
	if err := serialize_vector_tuple3_ObjectID_SequenceNumber_ObjectDigest(obj.SharedObjects, serializer); err != nil { return err }
	if err := obj.TransactionDigest.Serialize(serializer); err != nil { return err }
	if err := serialize_vector_tuple2_tuple3_ObjectID_SequenceNumber_ObjectDigest_Owner(obj.Created, serializer); err != nil { return err }
	if err := serialize_vector_tuple2_tuple3_ObjectID_SequenceNumber_ObjectDigest_Owner(obj.Mutated, serializer); err != nil { return err }
	if err := serialize_vector_tuple2_tuple3_ObjectID_SequenceNumber_ObjectDigest_Owner(obj.Unwrapped, serializer); err != nil { return err }
	if err := serialize_vector_tuple3_ObjectID_SequenceNumber_ObjectDigest(obj.Deleted, serializer); err != nil { return err }
	if err := serialize_vector_tuple3_ObjectID_SequenceNumber_ObjectDigest(obj.UnwrappedThenDeleted, serializer); err != nil { return err }
	if err := serialize_vector_tuple3_ObjectID_SequenceNumber_ObjectDigest(obj.Wrapped, serializer); err != nil { return err }
	if err := serialize_tuple2_tuple3_ObjectID_SequenceNumber_ObjectDigest_Owner(obj.GasObject, serializer); err != nil { return err }
	if err := serialize_option_TransactionEventsDigest(obj.EventsDigest, serializer); err != nil { return err }
	if err := serialize_vector_TransactionDigest(obj.Dependencies, serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *TransactionEffectsV1) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func DeserializeTransactionEffectsV1(deserializer serde.Deserializer) (TransactionEffectsV1, error) {
	var obj TransactionEffectsV1
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := DeserializeExecutionStatus(deserializer); err == nil { obj.Status = val } else { return obj, err }
	if val, err := deserializer.DeserializeU64(); err == nil { obj.ExecutedEpoch = val } else { return obj, err }
	if val, err := DeserializeGasCostSummary(deserializer); err == nil { obj.GasUsed = val } else { return obj, err }
	if val, err := deserialize_vector_tuple2_ObjectID_SequenceNumber(deserializer); err == nil { obj.ModifiedAtVersions = val } else { return obj, err }
	if val, err := deserialize_vector_tuple3_ObjectID_SequenceNumber_ObjectDigest(deserializer); err == nil { obj.SharedObjects = val } else { return obj, err }
	if val, err := DeserializeTransactionDigest(deserializer); err == nil { obj.TransactionDigest = val } else { return obj, err }
	if val, err := deserialize_vector_tuple2_tuple3_ObjectID_SequenceNumber_ObjectDigest_Owner(deserializer); err == nil { obj.Created = val } else { return obj, err }
	if val, err := deserialize_vector_tuple2_tuple3_ObjectID_SequenceNumber_ObjectDigest_Owner(deserializer); err == nil { obj.Mutated = val } else { return obj, err }
	if val, err := deserialize_vector_tuple2_tuple3_ObjectID_SequenceNumber_ObjectDigest_Owner(deserializer); err == nil { obj.Unwrapped = val } else { return obj, err }
	if val, err := deserialize_vector_tuple3_ObjectID_SequenceNumber_ObjectDigest(deserializer); err == nil { obj.Deleted = val } else { return obj, err }
	if val, err := deserialize_vector_tuple3_ObjectID_SequenceNumber_ObjectDigest(deserializer); err == nil { obj.UnwrappedThenDeleted = val } else { return obj, err }
	if val, err := deserialize_vector_tuple3_ObjectID_SequenceNumber_ObjectDigest(deserializer); err == nil { obj.Wrapped = val } else { return obj, err }
	if val, err := deserialize_tuple2_tuple3_ObjectID_SequenceNumber_ObjectDigest_Owner(deserializer); err == nil { obj.GasObject = val } else { return obj, err }
	if val, err := deserialize_option_TransactionEventsDigest(deserializer); err == nil { obj.EventsDigest = val } else { return obj, err }
	if val, err := deserialize_vector_TransactionDigest(deserializer); err == nil { obj.Dependencies = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

func BcsDeserializeTransactionEffectsV1(input []byte) (TransactionEffectsV1, error) {
	if input == nil {
		var obj TransactionEffectsV1
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input);
	obj, err := DeserializeTransactionEffectsV1(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type TransactionEventsDigest struct {
	Value Digest
}

func (obj *TransactionEventsDigest) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	if err := obj.Value.Serialize(serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *TransactionEventsDigest) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func DeserializeTransactionEventsDigest(deserializer serde.Deserializer) (TransactionEventsDigest, error) {
	var obj TransactionEventsDigest
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := DeserializeDigest(deserializer); err == nil { obj.Value = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

func BcsDeserializeTransactionEventsDigest(input []byte) (TransactionEventsDigest, error) {
	if input == nil {
		var obj TransactionEventsDigest
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input);
	obj, err := DeserializeTransactionEventsDigest(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type TransactionExpiration interface {
	isTransactionExpiration()
	Serialize(serializer serde.Serializer) error
	BcsSerialize() ([]byte, error)
}

func DeserializeTransactionExpiration(deserializer serde.Deserializer) (TransactionExpiration, error) {
	index, err := deserializer.DeserializeVariantIndex()
	if err != nil { return nil, err }

	switch index {
	case 0:
		if val, err := load_TransactionExpiration__None(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 1:
		if val, err := load_TransactionExpiration__Epoch(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	default:
		return nil, fmt.Errorf("Unknown variant index for TransactionExpiration: %d", index)
	}
}

func BcsDeserializeTransactionExpiration(input []byte) (TransactionExpiration, error) {
	if input == nil {
		var obj TransactionExpiration
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input);
	obj, err := DeserializeTransactionExpiration(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type TransactionExpiration__None struct {
}

func (*TransactionExpiration__None) isTransactionExpiration() {}

func (obj *TransactionExpiration__None) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(0)
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *TransactionExpiration__None) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_TransactionExpiration__None(deserializer serde.Deserializer) (TransactionExpiration__None, error) {
	var obj TransactionExpiration__None
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type TransactionExpiration__Epoch uint64

func (*TransactionExpiration__Epoch) isTransactionExpiration() {}

func (obj *TransactionExpiration__Epoch) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(1)
	if err := serializer.SerializeU64(((uint64)(*obj))); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *TransactionExpiration__Epoch) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_TransactionExpiration__Epoch(deserializer serde.Deserializer) (TransactionExpiration__Epoch, error) {
	var obj uint64
	if err := deserializer.IncreaseContainerDepth(); err != nil { return (TransactionExpiration__Epoch)(obj), err }
	if val, err := deserializer.DeserializeU64(); err == nil { obj = val } else { return ((TransactionExpiration__Epoch)(obj)), err }
	deserializer.DecreaseContainerDepth()
	return (TransactionExpiration__Epoch)(obj), nil
}

type TransactionKind interface {
	isTransactionKind()
	Serialize(serializer serde.Serializer) error
	BcsSerialize() ([]byte, error)
}

func DeserializeTransactionKind(deserializer serde.Deserializer) (TransactionKind, error) {
	index, err := deserializer.DeserializeVariantIndex()
	if err != nil { return nil, err }

	switch index {
	case 0:
		if val, err := load_TransactionKind__ProgrammableTransaction(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 1:
		if val, err := load_TransactionKind__ChangeEpoch(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 2:
		if val, err := load_TransactionKind__Genesis(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 3:
		if val, err := load_TransactionKind__ConsensusCommitPrologue(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	default:
		return nil, fmt.Errorf("Unknown variant index for TransactionKind: %d", index)
	}
}

func BcsDeserializeTransactionKind(input []byte) (TransactionKind, error) {
	if input == nil {
		var obj TransactionKind
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input);
	obj, err := DeserializeTransactionKind(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type TransactionKind__ProgrammableTransaction struct {
	Value ProgrammableTransaction
}

func (*TransactionKind__ProgrammableTransaction) isTransactionKind() {}

func (obj *TransactionKind__ProgrammableTransaction) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(0)
	if err := obj.Value.Serialize(serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *TransactionKind__ProgrammableTransaction) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_TransactionKind__ProgrammableTransaction(deserializer serde.Deserializer) (TransactionKind__ProgrammableTransaction, error) {
	var obj TransactionKind__ProgrammableTransaction
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := DeserializeProgrammableTransaction(deserializer); err == nil { obj.Value = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type TransactionKind__ChangeEpoch struct {
	Value ChangeEpoch
}

func (*TransactionKind__ChangeEpoch) isTransactionKind() {}

func (obj *TransactionKind__ChangeEpoch) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(1)
	if err := obj.Value.Serialize(serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *TransactionKind__ChangeEpoch) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_TransactionKind__ChangeEpoch(deserializer serde.Deserializer) (TransactionKind__ChangeEpoch, error) {
	var obj TransactionKind__ChangeEpoch
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := DeserializeChangeEpoch(deserializer); err == nil { obj.Value = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type TransactionKind__Genesis struct {
	Value GenesisTransaction
}

func (*TransactionKind__Genesis) isTransactionKind() {}

func (obj *TransactionKind__Genesis) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(2)
	if err := obj.Value.Serialize(serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *TransactionKind__Genesis) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_TransactionKind__Genesis(deserializer serde.Deserializer) (TransactionKind__Genesis, error) {
	var obj TransactionKind__Genesis
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := DeserializeGenesisTransaction(deserializer); err == nil { obj.Value = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type TransactionKind__ConsensusCommitPrologue struct {
	Value ConsensusCommitPrologue
}

func (*TransactionKind__ConsensusCommitPrologue) isTransactionKind() {}

func (obj *TransactionKind__ConsensusCommitPrologue) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(3)
	if err := obj.Value.Serialize(serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *TransactionKind__ConsensusCommitPrologue) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_TransactionKind__ConsensusCommitPrologue(deserializer serde.Deserializer) (TransactionKind__ConsensusCommitPrologue, error) {
	var obj TransactionKind__ConsensusCommitPrologue
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := DeserializeConsensusCommitPrologue(deserializer); err == nil { obj.Value = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type TypeArgumentError interface {
	isTypeArgumentError()
	Serialize(serializer serde.Serializer) error
	BcsSerialize() ([]byte, error)
}

func DeserializeTypeArgumentError(deserializer serde.Deserializer) (TypeArgumentError, error) {
	index, err := deserializer.DeserializeVariantIndex()
	if err != nil { return nil, err }

	switch index {
	case 0:
		if val, err := load_TypeArgumentError__TypeNotFound(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 1:
		if val, err := load_TypeArgumentError__ConstraintNotSatisfied(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	default:
		return nil, fmt.Errorf("Unknown variant index for TypeArgumentError: %d", index)
	}
}

func BcsDeserializeTypeArgumentError(input []byte) (TypeArgumentError, error) {
	if input == nil {
		var obj TypeArgumentError
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input);
	obj, err := DeserializeTypeArgumentError(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type TypeArgumentError__TypeNotFound struct {
}

func (*TypeArgumentError__TypeNotFound) isTypeArgumentError() {}

func (obj *TypeArgumentError__TypeNotFound) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(0)
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *TypeArgumentError__TypeNotFound) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_TypeArgumentError__TypeNotFound(deserializer serde.Deserializer) (TypeArgumentError__TypeNotFound, error) {
	var obj TypeArgumentError__TypeNotFound
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type TypeArgumentError__ConstraintNotSatisfied struct {
}

func (*TypeArgumentError__ConstraintNotSatisfied) isTypeArgumentError() {}

func (obj *TypeArgumentError__ConstraintNotSatisfied) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(1)
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *TypeArgumentError__ConstraintNotSatisfied) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_TypeArgumentError__ConstraintNotSatisfied(deserializer serde.Deserializer) (TypeArgumentError__ConstraintNotSatisfied, error) {
	var obj TypeArgumentError__ConstraintNotSatisfied
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type TypeOrigin struct {
	ModuleName string
	StructName string
	Package ObjectID
}

func (obj *TypeOrigin) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	if err := serializer.SerializeStr(obj.ModuleName); err != nil { return err }
	if err := serializer.SerializeStr(obj.StructName); err != nil { return err }
	if err := obj.Package.Serialize(serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *TypeOrigin) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func DeserializeTypeOrigin(deserializer serde.Deserializer) (TypeOrigin, error) {
	var obj TypeOrigin
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := deserializer.DeserializeStr(); err == nil { obj.ModuleName = val } else { return obj, err }
	if val, err := deserializer.DeserializeStr(); err == nil { obj.StructName = val } else { return obj, err }
	if val, err := DeserializeObjectID(deserializer); err == nil { obj.Package = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

func BcsDeserializeTypeOrigin(input []byte) (TypeOrigin, error) {
	if input == nil {
		var obj TypeOrigin
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input);
	obj, err := DeserializeTypeOrigin(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type TypeTag interface {
	isTypeTag()
	Serialize(serializer serde.Serializer) error
	BcsSerialize() ([]byte, error)
}

func DeserializeTypeTag(deserializer serde.Deserializer) (TypeTag, error) {
	index, err := deserializer.DeserializeVariantIndex()
	if err != nil { return nil, err }

	switch index {
	case 0:
		if val, err := load_TypeTag__Bool(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 1:
		if val, err := load_TypeTag__U8(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 2:
		if val, err := load_TypeTag__U64(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 3:
		if val, err := load_TypeTag__U128(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 4:
		if val, err := load_TypeTag__Address(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 5:
		if val, err := load_TypeTag__Signer(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 6:
		if val, err := load_TypeTag__Vector(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 7:
		if val, err := load_TypeTag__Struct(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 8:
		if val, err := load_TypeTag__U16(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 9:
		if val, err := load_TypeTag__U32(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 10:
		if val, err := load_TypeTag__U256(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	default:
		return nil, fmt.Errorf("Unknown variant index for TypeTag: %d", index)
	}
}

func BcsDeserializeTypeTag(input []byte) (TypeTag, error) {
	if input == nil {
		var obj TypeTag
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input);
	obj, err := DeserializeTypeTag(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type TypeTag__Bool struct {
}

func (*TypeTag__Bool) isTypeTag() {}

func (obj *TypeTag__Bool) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(0)
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *TypeTag__Bool) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_TypeTag__Bool(deserializer serde.Deserializer) (TypeTag__Bool, error) {
	var obj TypeTag__Bool
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type TypeTag__U8 struct {
}

func (*TypeTag__U8) isTypeTag() {}

func (obj *TypeTag__U8) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(1)
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *TypeTag__U8) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_TypeTag__U8(deserializer serde.Deserializer) (TypeTag__U8, error) {
	var obj TypeTag__U8
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type TypeTag__U64 struct {
}

func (*TypeTag__U64) isTypeTag() {}

func (obj *TypeTag__U64) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(2)
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *TypeTag__U64) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_TypeTag__U64(deserializer serde.Deserializer) (TypeTag__U64, error) {
	var obj TypeTag__U64
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type TypeTag__U128 struct {
}

func (*TypeTag__U128) isTypeTag() {}

func (obj *TypeTag__U128) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(3)
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *TypeTag__U128) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_TypeTag__U128(deserializer serde.Deserializer) (TypeTag__U128, error) {
	var obj TypeTag__U128
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type TypeTag__Address struct {
}

func (*TypeTag__Address) isTypeTag() {}

func (obj *TypeTag__Address) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(4)
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *TypeTag__Address) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_TypeTag__Address(deserializer serde.Deserializer) (TypeTag__Address, error) {
	var obj TypeTag__Address
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type TypeTag__Signer struct {
}

func (*TypeTag__Signer) isTypeTag() {}

func (obj *TypeTag__Signer) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(5)
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *TypeTag__Signer) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_TypeTag__Signer(deserializer serde.Deserializer) (TypeTag__Signer, error) {
	var obj TypeTag__Signer
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type TypeTag__Vector struct {
	Value TypeTag
}

func (*TypeTag__Vector) isTypeTag() {}

func (obj *TypeTag__Vector) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(6)
	if err := obj.Value.Serialize(serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *TypeTag__Vector) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_TypeTag__Vector(deserializer serde.Deserializer) (TypeTag__Vector, error) {
	var obj TypeTag__Vector
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := DeserializeTypeTag(deserializer); err == nil { obj.Value = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type TypeTag__Struct struct {
	Value StructTag
}

func (*TypeTag__Struct) isTypeTag() {}

func (obj *TypeTag__Struct) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(7)
	if err := obj.Value.Serialize(serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *TypeTag__Struct) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_TypeTag__Struct(deserializer serde.Deserializer) (TypeTag__Struct, error) {
	var obj TypeTag__Struct
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := DeserializeStructTag(deserializer); err == nil { obj.Value = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type TypeTag__U16 struct {
}

func (*TypeTag__U16) isTypeTag() {}

func (obj *TypeTag__U16) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(8)
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *TypeTag__U16) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_TypeTag__U16(deserializer serde.Deserializer) (TypeTag__U16, error) {
	var obj TypeTag__U16
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type TypeTag__U32 struct {
}

func (*TypeTag__U32) isTypeTag() {}

func (obj *TypeTag__U32) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(9)
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *TypeTag__U32) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_TypeTag__U32(deserializer serde.Deserializer) (TypeTag__U32, error) {
	var obj TypeTag__U32
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type TypeTag__U256 struct {
}

func (*TypeTag__U256) isTypeTag() {}

func (obj *TypeTag__U256) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(10)
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *TypeTag__U256) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_TypeTag__U256(deserializer serde.Deserializer) (TypeTag__U256, error) {
	var obj TypeTag__U256
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type TypedStoreError interface {
	isTypedStoreError()
	Serialize(serializer serde.Serializer) error
	BcsSerialize() ([]byte, error)
}

func DeserializeTypedStoreError(deserializer serde.Deserializer) (TypedStoreError, error) {
	index, err := deserializer.DeserializeVariantIndex()
	if err != nil { return nil, err }

	switch index {
	case 0:
		if val, err := load_TypedStoreError__RocksDbError(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 1:
		if val, err := load_TypedStoreError__SerializationError(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 2:
		if val, err := load_TypedStoreError__UnregisteredColumn(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 3:
		if val, err := load_TypedStoreError__CrossDbBatch(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 4:
		if val, err := load_TypedStoreError__MetricsReporting(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 5:
		if val, err := load_TypedStoreError__RetryableTransactionError(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	default:
		return nil, fmt.Errorf("Unknown variant index for TypedStoreError: %d", index)
	}
}

func BcsDeserializeTypedStoreError(input []byte) (TypedStoreError, error) {
	if input == nil {
		var obj TypedStoreError
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input);
	obj, err := DeserializeTypedStoreError(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type TypedStoreError__RocksDbError string

func (*TypedStoreError__RocksDbError) isTypedStoreError() {}

func (obj *TypedStoreError__RocksDbError) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(0)
	if err := serializer.SerializeStr(((string)(*obj))); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *TypedStoreError__RocksDbError) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_TypedStoreError__RocksDbError(deserializer serde.Deserializer) (TypedStoreError__RocksDbError, error) {
	var obj string
	if err := deserializer.IncreaseContainerDepth(); err != nil { return (TypedStoreError__RocksDbError)(obj), err }
	if val, err := deserializer.DeserializeStr(); err == nil { obj = val } else { return ((TypedStoreError__RocksDbError)(obj)), err }
	deserializer.DecreaseContainerDepth()
	return (TypedStoreError__RocksDbError)(obj), nil
}

type TypedStoreError__SerializationError string

func (*TypedStoreError__SerializationError) isTypedStoreError() {}

func (obj *TypedStoreError__SerializationError) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(1)
	if err := serializer.SerializeStr(((string)(*obj))); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *TypedStoreError__SerializationError) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_TypedStoreError__SerializationError(deserializer serde.Deserializer) (TypedStoreError__SerializationError, error) {
	var obj string
	if err := deserializer.IncreaseContainerDepth(); err != nil { return (TypedStoreError__SerializationError)(obj), err }
	if val, err := deserializer.DeserializeStr(); err == nil { obj = val } else { return ((TypedStoreError__SerializationError)(obj)), err }
	deserializer.DecreaseContainerDepth()
	return (TypedStoreError__SerializationError)(obj), nil
}

type TypedStoreError__UnregisteredColumn string

func (*TypedStoreError__UnregisteredColumn) isTypedStoreError() {}

func (obj *TypedStoreError__UnregisteredColumn) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(2)
	if err := serializer.SerializeStr(((string)(*obj))); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *TypedStoreError__UnregisteredColumn) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_TypedStoreError__UnregisteredColumn(deserializer serde.Deserializer) (TypedStoreError__UnregisteredColumn, error) {
	var obj string
	if err := deserializer.IncreaseContainerDepth(); err != nil { return (TypedStoreError__UnregisteredColumn)(obj), err }
	if val, err := deserializer.DeserializeStr(); err == nil { obj = val } else { return ((TypedStoreError__UnregisteredColumn)(obj)), err }
	deserializer.DecreaseContainerDepth()
	return (TypedStoreError__UnregisteredColumn)(obj), nil
}

type TypedStoreError__CrossDbBatch struct {
}

func (*TypedStoreError__CrossDbBatch) isTypedStoreError() {}

func (obj *TypedStoreError__CrossDbBatch) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(3)
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *TypedStoreError__CrossDbBatch) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_TypedStoreError__CrossDbBatch(deserializer serde.Deserializer) (TypedStoreError__CrossDbBatch, error) {
	var obj TypedStoreError__CrossDbBatch
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type TypedStoreError__MetricsReporting struct {
}

func (*TypedStoreError__MetricsReporting) isTypedStoreError() {}

func (obj *TypedStoreError__MetricsReporting) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(4)
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *TypedStoreError__MetricsReporting) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_TypedStoreError__MetricsReporting(deserializer serde.Deserializer) (TypedStoreError__MetricsReporting, error) {
	var obj TypedStoreError__MetricsReporting
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type TypedStoreError__RetryableTransactionError struct {
}

func (*TypedStoreError__RetryableTransactionError) isTypedStoreError() {}

func (obj *TypedStoreError__RetryableTransactionError) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(5)
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *TypedStoreError__RetryableTransactionError) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_TypedStoreError__RetryableTransactionError(deserializer serde.Deserializer) (TypedStoreError__RetryableTransactionError, error) {
	var obj TypedStoreError__RetryableTransactionError
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type UpgradeInfo struct {
	UpgradedId ObjectID
	UpgradedVersion SequenceNumber
}

func (obj *UpgradeInfo) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	if err := obj.UpgradedId.Serialize(serializer); err != nil { return err }
	if err := obj.UpgradedVersion.Serialize(serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *UpgradeInfo) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func DeserializeUpgradeInfo(deserializer serde.Deserializer) (UpgradeInfo, error) {
	var obj UpgradeInfo
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := DeserializeObjectID(deserializer); err == nil { obj.UpgradedId = val } else { return obj, err }
	if val, err := DeserializeSequenceNumber(deserializer); err == nil { obj.UpgradedVersion = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

func BcsDeserializeUpgradeInfo(input []byte) (UpgradeInfo, error) {
	if input == nil {
		var obj UpgradeInfo
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input);
	obj, err := DeserializeUpgradeInfo(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}
func serialize_array32_u8_array(value [32]uint8, serializer serde.Serializer) error {
	for _, item := range(value) {
		if err := serializer.SerializeU8(item); err != nil { return err }
	}
	return nil
}

func deserialize_array32_u8_array(deserializer serde.Deserializer) ([32]uint8, error) {
	var obj [32]uint8
	for i := range(obj) {
		if val, err := deserializer.DeserializeU8(); err == nil { obj[i] = val } else { return obj, err }
	}
	return obj, nil
}

func serialize_array64_u8_array(value [64]uint8, serializer serde.Serializer) error {
	for _, item := range(value) {
		if err := serializer.SerializeU8(item); err != nil { return err }
	}
	return nil
}

func deserialize_array64_u8_array(deserializer serde.Deserializer) ([64]uint8, error) {
	var obj [64]uint8
	for i := range(obj) {
		if val, err := deserializer.DeserializeU8(); err == nil { obj[i] = val } else { return obj, err }
	}
	return obj, nil
}

func serialize_map_ObjectID_to_UpgradeInfo(value map[ObjectID]UpgradeInfo, serializer serde.Serializer) error {
	if err := serializer.SerializeLen(uint64(len(value))); err != nil { return err }
	offsets := make([]uint64, len(value))
	count := 0
	for k, v := range(value) {
		offsets[count] = serializer.GetBufferOffset()
		count += 1
		if err := k.Serialize(serializer); err != nil { return err }
		if err := v.Serialize(serializer); err != nil { return err }
	}
	serializer.SortMapEntries(offsets);
	return nil
}

func deserialize_map_ObjectID_to_UpgradeInfo(deserializer serde.Deserializer) (map[ObjectID]UpgradeInfo, error) {
	length, err := deserializer.DeserializeLen()
	if err != nil { return nil, err }
	obj := make(map[ObjectID]UpgradeInfo)
	previous_slice := serde.Slice { 0, 0 }
	for i := 0; i < int(length); i++ {
		var slice serde.Slice
		slice.Start = deserializer.GetBufferOffset()
		var key ObjectID
		if val, err := DeserializeObjectID(deserializer); err == nil { key = val } else { return nil, err }
		slice.End = deserializer.GetBufferOffset()
		if i > 0 {
			err := deserializer.CheckThatKeySlicesAreIncreasing(previous_slice, slice)
			if err != nil { return nil, err }
		}
		previous_slice = slice
		if val, err := DeserializeUpgradeInfo(deserializer); err == nil { obj[key] = val } else { return nil, err }
	}
	return obj, nil
}

func serialize_map_str_to_bytes(value map[string][]byte, serializer serde.Serializer) error {
	if err := serializer.SerializeLen(uint64(len(value))); err != nil { return err }
	offsets := make([]uint64, len(value))
	count := 0
	for k, v := range(value) {
		offsets[count] = serializer.GetBufferOffset()
		count += 1
		if err := serializer.SerializeStr(k); err != nil { return err }
		if err := serializer.SerializeBytes(v); err != nil { return err }
	}
	serializer.SortMapEntries(offsets);
	return nil
}

func deserialize_map_str_to_bytes(deserializer serde.Deserializer) (map[string][]byte, error) {
	length, err := deserializer.DeserializeLen()
	if err != nil { return nil, err }
	obj := make(map[string][]byte)
	previous_slice := serde.Slice { 0, 0 }
	for i := 0; i < int(length); i++ {
		var slice serde.Slice
		slice.Start = deserializer.GetBufferOffset()
		var key string
		if val, err := deserializer.DeserializeStr(); err == nil { key = val } else { return nil, err }
		slice.End = deserializer.GetBufferOffset()
		if i > 0 {
			err := deserializer.CheckThatKeySlicesAreIncreasing(previous_slice, slice)
			if err != nil { return nil, err }
		}
		previous_slice = slice
		if val, err := deserializer.DeserializeBytes(); err == nil { obj[key] = val } else { return nil, err }
	}
	return obj, nil
}

func serialize_option_CheckpointDigest(value *CheckpointDigest, serializer serde.Serializer) error {
	if value != nil {
		if err := serializer.SerializeOptionTag(true); err != nil { return err }
		if err := (*value).Serialize(serializer); err != nil { return err }
	} else {
		if err := serializer.SerializeOptionTag(false); err != nil { return err }
	}
	return nil
}

func deserialize_option_CheckpointDigest(deserializer serde.Deserializer) (*CheckpointDigest, error) {
	tag, err := deserializer.DeserializeOptionTag()
	if err != nil { return nil, err }
	if tag {
		value := new(CheckpointDigest)
		if val, err := DeserializeCheckpointDigest(deserializer); err == nil { *value = val } else { return nil, err }
	        return value, nil
	} else {
		return nil, nil
	}
}

func serialize_option_EndOfEpochData(value *EndOfEpochData, serializer serde.Serializer) error {
	if value != nil {
		if err := serializer.SerializeOptionTag(true); err != nil { return err }
		if err := (*value).Serialize(serializer); err != nil { return err }
	} else {
		if err := serializer.SerializeOptionTag(false); err != nil { return err }
	}
	return nil
}

func deserialize_option_EndOfEpochData(deserializer serde.Deserializer) (*EndOfEpochData, error) {
	tag, err := deserializer.DeserializeOptionTag()
	if err != nil { return nil, err }
	if tag {
		value := new(EndOfEpochData)
		if val, err := DeserializeEndOfEpochData(deserializer); err == nil { *value = val } else { return nil, err }
	        return value, nil
	} else {
		return nil, nil
	}
}

func serialize_option_MoveLocation(value *MoveLocation, serializer serde.Serializer) error {
	if value != nil {
		if err := serializer.SerializeOptionTag(true); err != nil { return err }
		if err := (*value).Serialize(serializer); err != nil { return err }
	} else {
		if err := serializer.SerializeOptionTag(false); err != nil { return err }
	}
	return nil
}

func deserialize_option_MoveLocation(deserializer serde.Deserializer) (*MoveLocation, error) {
	tag, err := deserializer.DeserializeOptionTag()
	if err != nil { return nil, err }
	if tag {
		value := new(MoveLocation)
		if val, err := DeserializeMoveLocation(deserializer); err == nil { *value = val } else { return nil, err }
	        return value, nil
	} else {
		return nil, nil
	}
}

func serialize_option_TransactionEventsDigest(value *TransactionEventsDigest, serializer serde.Serializer) error {
	if value != nil {
		if err := serializer.SerializeOptionTag(true); err != nil { return err }
		if err := (*value).Serialize(serializer); err != nil { return err }
	} else {
		if err := serializer.SerializeOptionTag(false); err != nil { return err }
	}
	return nil
}

func deserialize_option_TransactionEventsDigest(deserializer serde.Deserializer) (*TransactionEventsDigest, error) {
	tag, err := deserializer.DeserializeOptionTag()
	if err != nil { return nil, err }
	if tag {
		value := new(TransactionEventsDigest)
		if val, err := DeserializeTransactionEventsDigest(deserializer); err == nil { *value = val } else { return nil, err }
	        return value, nil
	} else {
		return nil, nil
	}
}

func serialize_option_TypeTag(value *TypeTag, serializer serde.Serializer) error {
	if value != nil {
		if err := serializer.SerializeOptionTag(true); err != nil { return err }
		if err := (*value).Serialize(serializer); err != nil { return err }
	} else {
		if err := serializer.SerializeOptionTag(false); err != nil { return err }
	}
	return nil
}

func deserialize_option_TypeTag(deserializer serde.Deserializer) (*TypeTag, error) {
	tag, err := deserializer.DeserializeOptionTag()
	if err != nil { return nil, err }
	if tag {
		value := new(TypeTag)
		if val, err := DeserializeTypeTag(deserializer); err == nil { *value = val } else { return nil, err }
	        return value, nil
	} else {
		return nil, nil
	}
}

func serialize_option_str(value *string, serializer serde.Serializer) error {
	if value != nil {
		if err := serializer.SerializeOptionTag(true); err != nil { return err }
		if err := serializer.SerializeStr((*value)); err != nil { return err }
	} else {
		if err := serializer.SerializeOptionTag(false); err != nil { return err }
	}
	return nil
}

func deserialize_option_str(deserializer serde.Deserializer) (*string, error) {
	tag, err := deserializer.DeserializeOptionTag()
	if err != nil { return nil, err }
	if tag {
		value := new(string)
		if val, err := deserializer.DeserializeStr(); err == nil { *value = val } else { return nil, err }
	        return value, nil
	} else {
		return nil, nil
	}
}

func serialize_option_u64(value *uint64, serializer serde.Serializer) error {
	if value != nil {
		if err := serializer.SerializeOptionTag(true); err != nil { return err }
		if err := serializer.SerializeU64((*value)); err != nil { return err }
	} else {
		if err := serializer.SerializeOptionTag(false); err != nil { return err }
	}
	return nil
}

func deserialize_option_u64(deserializer serde.Deserializer) (*uint64, error) {
	tag, err := deserializer.DeserializeOptionTag()
	if err != nil { return nil, err }
	if tag {
		value := new(uint64)
		if val, err := deserializer.DeserializeU64(); err == nil { *value = val } else { return nil, err }
	        return value, nil
	} else {
		return nil, nil
	}
}

func serialize_tuple2_AuthorityPublicKeyBytes_u64(value struct {Field0 AuthorityPublicKeyBytes; Field1 uint64}, serializer serde.Serializer) error {
	if err := value.Field0.Serialize(serializer); err != nil { return err }
	if err := serializer.SerializeU64(value.Field1); err != nil { return err }
	return nil
}

func deserialize_tuple2_AuthorityPublicKeyBytes_u64(deserializer serde.Deserializer) (struct {Field0 AuthorityPublicKeyBytes; Field1 uint64}, error) {
	var obj struct {Field0 AuthorityPublicKeyBytes; Field1 uint64}
	if val, err := DeserializeAuthorityPublicKeyBytes(deserializer); err == nil { obj.Field0 = val } else { return obj, err }
	if val, err := deserializer.DeserializeU64(); err == nil { obj.Field1 = val } else { return obj, err }
	return obj, nil
}

func serialize_tuple2_ObjectID_SequenceNumber(value struct {Field0 ObjectID; Field1 SequenceNumber}, serializer serde.Serializer) error {
	if err := value.Field0.Serialize(serializer); err != nil { return err }
	if err := value.Field1.Serialize(serializer); err != nil { return err }
	return nil
}

func deserialize_tuple2_ObjectID_SequenceNumber(deserializer serde.Deserializer) (struct {Field0 ObjectID; Field1 SequenceNumber}, error) {
	var obj struct {Field0 ObjectID; Field1 SequenceNumber}
	if val, err := DeserializeObjectID(deserializer); err == nil { obj.Field0 = val } else { return obj, err }
	if val, err := DeserializeSequenceNumber(deserializer); err == nil { obj.Field1 = val } else { return obj, err }
	return obj, nil
}

func serialize_tuple2_str_u8(value struct {Field0 string; Field1 uint8}, serializer serde.Serializer) error {
	if err := serializer.SerializeStr(value.Field0); err != nil { return err }
	if err := serializer.SerializeU8(value.Field1); err != nil { return err }
	return nil
}

func deserialize_tuple2_str_u8(deserializer serde.Deserializer) (struct {Field0 string; Field1 uint8}, error) {
	var obj struct {Field0 string; Field1 uint8}
	if val, err := deserializer.DeserializeStr(); err == nil { obj.Field0 = val } else { return obj, err }
	if val, err := deserializer.DeserializeU8(); err == nil { obj.Field1 = val } else { return obj, err }
	return obj, nil
}

func serialize_tuple2_tuple3_ObjectID_SequenceNumber_ObjectDigest_Owner(value struct {Field0 struct {Field0 ObjectID; Field1 SequenceNumber; Field2 ObjectDigest}; Field1 Owner}, serializer serde.Serializer) error {
	if err := serialize_tuple3_ObjectID_SequenceNumber_ObjectDigest(value.Field0, serializer); err != nil { return err }
	if err := value.Field1.Serialize(serializer); err != nil { return err }
	return nil
}

func deserialize_tuple2_tuple3_ObjectID_SequenceNumber_ObjectDigest_Owner(deserializer serde.Deserializer) (struct {Field0 struct {Field0 ObjectID; Field1 SequenceNumber; Field2 ObjectDigest}; Field1 Owner}, error) {
	var obj struct {Field0 struct {Field0 ObjectID; Field1 SequenceNumber; Field2 ObjectDigest}; Field1 Owner}
	if val, err := deserialize_tuple3_ObjectID_SequenceNumber_ObjectDigest(deserializer); err == nil { obj.Field0 = val } else { return obj, err }
	if val, err := DeserializeOwner(deserializer); err == nil { obj.Field1 = val } else { return obj, err }
	return obj, nil
}

func serialize_tuple3_ObjectID_SequenceNumber_ObjectDigest(value struct {Field0 ObjectID; Field1 SequenceNumber; Field2 ObjectDigest}, serializer serde.Serializer) error {
	if err := value.Field0.Serialize(serializer); err != nil { return err }
	if err := value.Field1.Serialize(serializer); err != nil { return err }
	if err := value.Field2.Serialize(serializer); err != nil { return err }
	return nil
}

func deserialize_tuple3_ObjectID_SequenceNumber_ObjectDigest(deserializer serde.Deserializer) (struct {Field0 ObjectID; Field1 SequenceNumber; Field2 ObjectDigest}, error) {
	var obj struct {Field0 ObjectID; Field1 SequenceNumber; Field2 ObjectDigest}
	if val, err := DeserializeObjectID(deserializer); err == nil { obj.Field0 = val } else { return obj, err }
	if val, err := DeserializeSequenceNumber(deserializer); err == nil { obj.Field1 = val } else { return obj, err }
	if val, err := DeserializeObjectDigest(deserializer); err == nil { obj.Field2 = val } else { return obj, err }
	return obj, nil
}

func serialize_tuple3_SequenceNumber_vector_vector_u8_vector_ObjectID(value struct {Field0 SequenceNumber; Field1 [][]uint8; Field2 []ObjectID}, serializer serde.Serializer) error {
	if err := value.Field0.Serialize(serializer); err != nil { return err }
	if err := serialize_vector_vector_u8(value.Field1, serializer); err != nil { return err }
	if err := serialize_vector_ObjectID(value.Field2, serializer); err != nil { return err }
	return nil
}

func deserialize_tuple3_SequenceNumber_vector_vector_u8_vector_ObjectID(deserializer serde.Deserializer) (struct {Field0 SequenceNumber; Field1 [][]uint8; Field2 []ObjectID}, error) {
	var obj struct {Field0 SequenceNumber; Field1 [][]uint8; Field2 []ObjectID}
	if val, err := DeserializeSequenceNumber(deserializer); err == nil { obj.Field0 = val } else { return obj, err }
	if val, err := deserialize_vector_vector_u8(deserializer); err == nil { obj.Field1 = val } else { return obj, err }
	if val, err := deserialize_vector_ObjectID(deserializer); err == nil { obj.Field2 = val } else { return obj, err }
	return obj, nil
}

func serialize_vector_Argument(value []Argument, serializer serde.Serializer) error {
	if err := serializer.SerializeLen(uint64(len(value))); err != nil { return err }
	for _, item := range(value) {
		if err := item.Serialize(serializer); err != nil { return err }
	}
	return nil
}

func deserialize_vector_Argument(deserializer serde.Deserializer) ([]Argument, error) {
	length, err := deserializer.DeserializeLen()
	if err != nil { return nil, err }
	obj := make([]Argument, length)
	for i := range(obj) {
		if val, err := DeserializeArgument(deserializer); err == nil { obj[i] = val } else { return nil, err }
	}
	return obj, nil
}

func serialize_vector_CallArg(value []CallArg, serializer serde.Serializer) error {
	if err := serializer.SerializeLen(uint64(len(value))); err != nil { return err }
	for _, item := range(value) {
		if err := item.Serialize(serializer); err != nil { return err }
	}
	return nil
}

func deserialize_vector_CallArg(deserializer serde.Deserializer) ([]CallArg, error) {
	length, err := deserializer.DeserializeLen()
	if err != nil { return nil, err }
	obj := make([]CallArg, length)
	for i := range(obj) {
		if val, err := DeserializeCallArg(deserializer); err == nil { obj[i] = val } else { return nil, err }
	}
	return obj, nil
}

func serialize_vector_CheckpointCommitment(value []CheckpointCommitment, serializer serde.Serializer) error {
	if err := serializer.SerializeLen(uint64(len(value))); err != nil { return err }
	for _, item := range(value) {
		if err := item.Serialize(serializer); err != nil { return err }
	}
	return nil
}

func deserialize_vector_CheckpointCommitment(deserializer serde.Deserializer) ([]CheckpointCommitment, error) {
	length, err := deserializer.DeserializeLen()
	if err != nil { return nil, err }
	obj := make([]CheckpointCommitment, length)
	for i := range(obj) {
		if val, err := DeserializeCheckpointCommitment(deserializer); err == nil { obj[i] = val } else { return nil, err }
	}
	return obj, nil
}

func serialize_vector_Command(value []Command, serializer serde.Serializer) error {
	if err := serializer.SerializeLen(uint64(len(value))); err != nil { return err }
	for _, item := range(value) {
		if err := item.Serialize(serializer); err != nil { return err }
	}
	return nil
}

func deserialize_vector_Command(deserializer serde.Deserializer) ([]Command, error) {
	length, err := deserializer.DeserializeLen()
	if err != nil { return nil, err }
	obj := make([]Command, length)
	for i := range(obj) {
		if val, err := DeserializeCommand(deserializer); err == nil { obj[i] = val } else { return nil, err }
	}
	return obj, nil
}

func serialize_vector_CompressedSignature(value []CompressedSignature, serializer serde.Serializer) error {
	if err := serializer.SerializeLen(uint64(len(value))); err != nil { return err }
	for _, item := range(value) {
		if err := item.Serialize(serializer); err != nil { return err }
	}
	return nil
}

func deserialize_vector_CompressedSignature(deserializer serde.Deserializer) ([]CompressedSignature, error) {
	length, err := deserializer.DeserializeLen()
	if err != nil { return nil, err }
	obj := make([]CompressedSignature, length)
	for i := range(obj) {
		if val, err := DeserializeCompressedSignature(deserializer); err == nil { obj[i] = val } else { return nil, err }
	}
	return obj, nil
}

func serialize_vector_ExecutionData(value []ExecutionData, serializer serde.Serializer) error {
	if err := serializer.SerializeLen(uint64(len(value))); err != nil { return err }
	for _, item := range(value) {
		if err := item.Serialize(serializer); err != nil { return err }
	}
	return nil
}

func deserialize_vector_ExecutionData(deserializer serde.Deserializer) ([]ExecutionData, error) {
	length, err := deserializer.DeserializeLen()
	if err != nil { return nil, err }
	obj := make([]ExecutionData, length)
	for i := range(obj) {
		if val, err := DeserializeExecutionData(deserializer); err == nil { obj[i] = val } else { return nil, err }
	}
	return obj, nil
}

func serialize_vector_ExecutionDigests(value []ExecutionDigests, serializer serde.Serializer) error {
	if err := serializer.SerializeLen(uint64(len(value))); err != nil { return err }
	for _, item := range(value) {
		if err := item.Serialize(serializer); err != nil { return err }
	}
	return nil
}

func deserialize_vector_ExecutionDigests(deserializer serde.Deserializer) ([]ExecutionDigests, error) {
	length, err := deserializer.DeserializeLen()
	if err != nil { return nil, err }
	obj := make([]ExecutionDigests, length)
	for i := range(obj) {
		if val, err := DeserializeExecutionDigests(deserializer); err == nil { obj[i] = val } else { return nil, err }
	}
	return obj, nil
}

func serialize_vector_GenericSignature(value []GenericSignature, serializer serde.Serializer) error {
	if err := serializer.SerializeLen(uint64(len(value))); err != nil { return err }
	for _, item := range(value) {
		if err := item.Serialize(serializer); err != nil { return err }
	}
	return nil
}

func deserialize_vector_GenericSignature(deserializer serde.Deserializer) ([]GenericSignature, error) {
	length, err := deserializer.DeserializeLen()
	if err != nil { return nil, err }
	obj := make([]GenericSignature, length)
	for i := range(obj) {
		if val, err := DeserializeGenericSignature(deserializer); err == nil { obj[i] = val } else { return nil, err }
	}
	return obj, nil
}

func serialize_vector_GenesisObject(value []GenesisObject, serializer serde.Serializer) error {
	if err := serializer.SerializeLen(uint64(len(value))); err != nil { return err }
	for _, item := range(value) {
		if err := item.Serialize(serializer); err != nil { return err }
	}
	return nil
}

func deserialize_vector_GenesisObject(deserializer serde.Deserializer) ([]GenesisObject, error) {
	length, err := deserializer.DeserializeLen()
	if err != nil { return nil, err }
	obj := make([]GenesisObject, length)
	for i := range(obj) {
		if val, err := DeserializeGenesisObject(deserializer); err == nil { obj[i] = val } else { return nil, err }
	}
	return obj, nil
}

func serialize_vector_MoveFieldLayout(value []MoveFieldLayout, serializer serde.Serializer) error {
	if err := serializer.SerializeLen(uint64(len(value))); err != nil { return err }
	for _, item := range(value) {
		if err := item.Serialize(serializer); err != nil { return err }
	}
	return nil
}

func deserialize_vector_MoveFieldLayout(deserializer serde.Deserializer) ([]MoveFieldLayout, error) {
	length, err := deserializer.DeserializeLen()
	if err != nil { return nil, err }
	obj := make([]MoveFieldLayout, length)
	for i := range(obj) {
		if val, err := DeserializeMoveFieldLayout(deserializer); err == nil { obj[i] = val } else { return nil, err }
	}
	return obj, nil
}

func serialize_vector_MoveTypeLayout(value []MoveTypeLayout, serializer serde.Serializer) error {
	if err := serializer.SerializeLen(uint64(len(value))); err != nil { return err }
	for _, item := range(value) {
		if err := item.Serialize(serializer); err != nil { return err }
	}
	return nil
}

func deserialize_vector_MoveTypeLayout(deserializer serde.Deserializer) ([]MoveTypeLayout, error) {
	length, err := deserializer.DeserializeLen()
	if err != nil { return nil, err }
	obj := make([]MoveTypeLayout, length)
	for i := range(obj) {
		if val, err := DeserializeMoveTypeLayout(deserializer); err == nil { obj[i] = val } else { return nil, err }
	}
	return obj, nil
}

func serialize_vector_ObjectID(value []ObjectID, serializer serde.Serializer) error {
	if err := serializer.SerializeLen(uint64(len(value))); err != nil { return err }
	for _, item := range(value) {
		if err := item.Serialize(serializer); err != nil { return err }
	}
	return nil
}

func deserialize_vector_ObjectID(deserializer serde.Deserializer) ([]ObjectID, error) {
	length, err := deserializer.DeserializeLen()
	if err != nil { return nil, err }
	obj := make([]ObjectID, length)
	for i := range(obj) {
		if val, err := DeserializeObjectID(deserializer); err == nil { obj[i] = val } else { return nil, err }
	}
	return obj, nil
}

func serialize_vector_SenderSignedTransaction(value []SenderSignedTransaction, serializer serde.Serializer) error {
	if err := serializer.SerializeLen(uint64(len(value))); err != nil { return err }
	for _, item := range(value) {
		if err := item.Serialize(serializer); err != nil { return err }
	}
	return nil
}

func deserialize_vector_SenderSignedTransaction(deserializer serde.Deserializer) ([]SenderSignedTransaction, error) {
	length, err := deserializer.DeserializeLen()
	if err != nil { return nil, err }
	obj := make([]SenderSignedTransaction, length)
	for i := range(obj) {
		if val, err := DeserializeSenderSignedTransaction(deserializer); err == nil { obj[i] = val } else { return nil, err }
	}
	return obj, nil
}

func serialize_vector_TransactionDigest(value []TransactionDigest, serializer serde.Serializer) error {
	if err := serializer.SerializeLen(uint64(len(value))); err != nil { return err }
	for _, item := range(value) {
		if err := item.Serialize(serializer); err != nil { return err }
	}
	return nil
}

func deserialize_vector_TransactionDigest(deserializer serde.Deserializer) ([]TransactionDigest, error) {
	length, err := deserializer.DeserializeLen()
	if err != nil { return nil, err }
	obj := make([]TransactionDigest, length)
	for i := range(obj) {
		if val, err := DeserializeTransactionDigest(deserializer); err == nil { obj[i] = val } else { return nil, err }
	}
	return obj, nil
}

func serialize_vector_TypeOrigin(value []TypeOrigin, serializer serde.Serializer) error {
	if err := serializer.SerializeLen(uint64(len(value))); err != nil { return err }
	for _, item := range(value) {
		if err := item.Serialize(serializer); err != nil { return err }
	}
	return nil
}

func deserialize_vector_TypeOrigin(deserializer serde.Deserializer) ([]TypeOrigin, error) {
	length, err := deserializer.DeserializeLen()
	if err != nil { return nil, err }
	obj := make([]TypeOrigin, length)
	for i := range(obj) {
		if val, err := DeserializeTypeOrigin(deserializer); err == nil { obj[i] = val } else { return nil, err }
	}
	return obj, nil
}

func serialize_vector_TypeTag(value []TypeTag, serializer serde.Serializer) error {
	if err := serializer.SerializeLen(uint64(len(value))); err != nil { return err }
	for _, item := range(value) {
		if err := item.Serialize(serializer); err != nil { return err }
	}
	return nil
}

func deserialize_vector_TypeTag(deserializer serde.Deserializer) ([]TypeTag, error) {
	length, err := deserializer.DeserializeLen()
	if err != nil { return nil, err }
	obj := make([]TypeTag, length)
	for i := range(obj) {
		if val, err := DeserializeTypeTag(deserializer); err == nil { obj[i] = val } else { return nil, err }
	}
	return obj, nil
}

func serialize_vector_tuple2_AuthorityPublicKeyBytes_u64(value []struct {Field0 AuthorityPublicKeyBytes; Field1 uint64}, serializer serde.Serializer) error {
	if err := serializer.SerializeLen(uint64(len(value))); err != nil { return err }
	for _, item := range(value) {
		if err := serialize_tuple2_AuthorityPublicKeyBytes_u64(item, serializer); err != nil { return err }
	}
	return nil
}

func deserialize_vector_tuple2_AuthorityPublicKeyBytes_u64(deserializer serde.Deserializer) ([]struct {Field0 AuthorityPublicKeyBytes; Field1 uint64}, error) {
	length, err := deserializer.DeserializeLen()
	if err != nil { return nil, err }
	obj := make([]struct {Field0 AuthorityPublicKeyBytes; Field1 uint64}, length)
	for i := range(obj) {
		if val, err := deserialize_tuple2_AuthorityPublicKeyBytes_u64(deserializer); err == nil { obj[i] = val } else { return nil, err }
	}
	return obj, nil
}

func serialize_vector_tuple2_ObjectID_SequenceNumber(value []struct {Field0 ObjectID; Field1 SequenceNumber}, serializer serde.Serializer) error {
	if err := serializer.SerializeLen(uint64(len(value))); err != nil { return err }
	for _, item := range(value) {
		if err := serialize_tuple2_ObjectID_SequenceNumber(item, serializer); err != nil { return err }
	}
	return nil
}

func deserialize_vector_tuple2_ObjectID_SequenceNumber(deserializer serde.Deserializer) ([]struct {Field0 ObjectID; Field1 SequenceNumber}, error) {
	length, err := deserializer.DeserializeLen()
	if err != nil { return nil, err }
	obj := make([]struct {Field0 ObjectID; Field1 SequenceNumber}, length)
	for i := range(obj) {
		if val, err := deserialize_tuple2_ObjectID_SequenceNumber(deserializer); err == nil { obj[i] = val } else { return nil, err }
	}
	return obj, nil
}

func serialize_vector_tuple2_str_u8(value []struct {Field0 string; Field1 uint8}, serializer serde.Serializer) error {
	if err := serializer.SerializeLen(uint64(len(value))); err != nil { return err }
	for _, item := range(value) {
		if err := serialize_tuple2_str_u8(item, serializer); err != nil { return err }
	}
	return nil
}

func deserialize_vector_tuple2_str_u8(deserializer serde.Deserializer) ([]struct {Field0 string; Field1 uint8}, error) {
	length, err := deserializer.DeserializeLen()
	if err != nil { return nil, err }
	obj := make([]struct {Field0 string; Field1 uint8}, length)
	for i := range(obj) {
		if val, err := deserialize_tuple2_str_u8(deserializer); err == nil { obj[i] = val } else { return nil, err }
	}
	return obj, nil
}

func serialize_vector_tuple2_tuple3_ObjectID_SequenceNumber_ObjectDigest_Owner(value []struct {Field0 struct {Field0 ObjectID; Field1 SequenceNumber; Field2 ObjectDigest}; Field1 Owner}, serializer serde.Serializer) error {
	if err := serializer.SerializeLen(uint64(len(value))); err != nil { return err }
	for _, item := range(value) {
		if err := serialize_tuple2_tuple3_ObjectID_SequenceNumber_ObjectDigest_Owner(item, serializer); err != nil { return err }
	}
	return nil
}

func deserialize_vector_tuple2_tuple3_ObjectID_SequenceNumber_ObjectDigest_Owner(deserializer serde.Deserializer) ([]struct {Field0 struct {Field0 ObjectID; Field1 SequenceNumber; Field2 ObjectDigest}; Field1 Owner}, error) {
	length, err := deserializer.DeserializeLen()
	if err != nil { return nil, err }
	obj := make([]struct {Field0 struct {Field0 ObjectID; Field1 SequenceNumber; Field2 ObjectDigest}; Field1 Owner}, length)
	for i := range(obj) {
		if val, err := deserialize_tuple2_tuple3_ObjectID_SequenceNumber_ObjectDigest_Owner(deserializer); err == nil { obj[i] = val } else { return nil, err }
	}
	return obj, nil
}

func serialize_vector_tuple3_ObjectID_SequenceNumber_ObjectDigest(value []struct {Field0 ObjectID; Field1 SequenceNumber; Field2 ObjectDigest}, serializer serde.Serializer) error {
	if err := serializer.SerializeLen(uint64(len(value))); err != nil { return err }
	for _, item := range(value) {
		if err := serialize_tuple3_ObjectID_SequenceNumber_ObjectDigest(item, serializer); err != nil { return err }
	}
	return nil
}

func deserialize_vector_tuple3_ObjectID_SequenceNumber_ObjectDigest(deserializer serde.Deserializer) ([]struct {Field0 ObjectID; Field1 SequenceNumber; Field2 ObjectDigest}, error) {
	length, err := deserializer.DeserializeLen()
	if err != nil { return nil, err }
	obj := make([]struct {Field0 ObjectID; Field1 SequenceNumber; Field2 ObjectDigest}, length)
	for i := range(obj) {
		if val, err := deserialize_tuple3_ObjectID_SequenceNumber_ObjectDigest(deserializer); err == nil { obj[i] = val } else { return nil, err }
	}
	return obj, nil
}

func serialize_vector_tuple3_SequenceNumber_vector_vector_u8_vector_ObjectID(value []struct {Field0 SequenceNumber; Field1 [][]uint8; Field2 []ObjectID}, serializer serde.Serializer) error {
	if err := serializer.SerializeLen(uint64(len(value))); err != nil { return err }
	for _, item := range(value) {
		if err := serialize_tuple3_SequenceNumber_vector_vector_u8_vector_ObjectID(item, serializer); err != nil { return err }
	}
	return nil
}

func deserialize_vector_tuple3_SequenceNumber_vector_vector_u8_vector_ObjectID(deserializer serde.Deserializer) ([]struct {Field0 SequenceNumber; Field1 [][]uint8; Field2 []ObjectID}, error) {
	length, err := deserializer.DeserializeLen()
	if err != nil { return nil, err }
	obj := make([]struct {Field0 SequenceNumber; Field1 [][]uint8; Field2 []ObjectID}, length)
	for i := range(obj) {
		if val, err := deserialize_tuple3_SequenceNumber_vector_vector_u8_vector_ObjectID(deserializer); err == nil { obj[i] = val } else { return nil, err }
	}
	return obj, nil
}

func serialize_vector_u8(value []uint8, serializer serde.Serializer) error {
	if err := serializer.SerializeLen(uint64(len(value))); err != nil { return err }
	for _, item := range(value) {
		if err := serializer.SerializeU8(item); err != nil { return err }
	}
	return nil
}

func deserialize_vector_u8(deserializer serde.Deserializer) ([]uint8, error) {
	length, err := deserializer.DeserializeLen()
	if err != nil { return nil, err }
	obj := make([]uint8, length)
	for i := range(obj) {
		if val, err := deserializer.DeserializeU8(); err == nil { obj[i] = val } else { return nil, err }
	}
	return obj, nil
}

func serialize_vector_vector_GenericSignature(value [][]GenericSignature, serializer serde.Serializer) error {
	if err := serializer.SerializeLen(uint64(len(value))); err != nil { return err }
	for _, item := range(value) {
		if err := serialize_vector_GenericSignature(item, serializer); err != nil { return err }
	}
	return nil
}

func deserialize_vector_vector_GenericSignature(deserializer serde.Deserializer) ([][]GenericSignature, error) {
	length, err := deserializer.DeserializeLen()
	if err != nil { return nil, err }
	obj := make([][]GenericSignature, length)
	for i := range(obj) {
		if val, err := deserialize_vector_GenericSignature(deserializer); err == nil { obj[i] = val } else { return nil, err }
	}
	return obj, nil
}

func serialize_vector_vector_u8(value [][]uint8, serializer serde.Serializer) error {
	if err := serializer.SerializeLen(uint64(len(value))); err != nil { return err }
	for _, item := range(value) {
		if err := serialize_vector_u8(item, serializer); err != nil { return err }
	}
	return nil
}

func deserialize_vector_vector_u8(deserializer serde.Deserializer) ([][]uint8, error) {
	length, err := deserializer.DeserializeLen()
	if err != nil { return nil, err }
	obj := make([][]uint8, length)
	for i := range(obj) {
		if val, err := deserialize_vector_u8(deserializer); err == nil { obj[i] = val } else { return nil, err }
	}
	return obj, nil
}

