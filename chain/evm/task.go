package evm

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	xc "github.com/jumpcrypto/crosschain"
	"github.com/shopspring/decimal"
	"golang.org/x/crypto/sha3"
)

func (txBuilder TxBuilder) NewTask(from xc.Address, to xc.Address, amount xc.AmountBlockchain, input xc.TxInput) (xc.Tx, error) {
	txInput := input.(*TxInput)
	task := txBuilder.Asset.(*xc.TaskConfig)

	switch task.Code {
	case "ProxyTransferTx":
		return txBuilder.BuildProxyTransferTx(from, to, amount, txInput)
	case "WormholeTransferTx":
		return txBuilder.BuildWormholeTransferTx(from, to, amount, txInput)
	}
	return txBuilder.BuildTaskTx(from, to, amount, txInput)
}

func (txBuilder TxBuilder) BuildTaskPayload(taskFrom xc.Address, taskTo xc.Address, taskAmount xc.AmountBlockchain, input *TxInput) (string, xc.AmountBlockchain, []byte, error) {
	asset := txBuilder.Asset.GetAssetConfig()
	task := txBuilder.Asset.GetTask()

	// value, either tx value (for payable functions) or 0
	valueZero := xc.NewAmountBlockchainFromUint64(0)
	valueTx := taskAmount
	value := valueTx
	valueConsumed := false

	// tx.to, typically contract address
	to := asset.Contract

	// data
	var data []byte

	// on EVM we expect only 1 operation
	if len(task.Operations) != 1 {
		return to, value, data, fmt.Errorf("expected 1 operation, got %d", len(task.Operations))
	}

	op := task.Operations[0]

	// override to
	switch contract := op.Contract.(type) {
	case nil:
		// pass
	case string:
		if contract == "dst_asset" {
			to = task.DstAsset.GetAssetConfig().Contract
		} else if contract != "" {
			to = contract
		}
	case map[interface{}]interface{}:
		nativeAsset := string(asset.NativeAsset)
		for k, v := range contract {
			// map keys are lowercase
			if strings.EqualFold(k.(string), nativeAsset) {
				to = v.(string)
			}
		}
	default:
		return to, value, data, fmt.Errorf("invalid config for task=%s contract type=%T", task.ID(), contract)
	}

	// methodID == function signature
	methodID, err := hex.DecodeString(op.Signature)
	if err != nil || len(methodID) != 4 {
		return to, value, data, fmt.Errorf("invalid task signature: %s", op.Signature)
	}
	data = append(data, methodID...)

	userPassedParamIndex := 0
	// iterate over operation params, matching them up to user-passed params
	for _, p := range op.Params {
		if p.Bind != "" {
			// binds
			switch p.Bind {
			case "amount":
				// amount is encoded as uint256
				paddedValue := common.LeftPadBytes(valueTx.Int().Bytes(), 32)
				data = append(data, paddedValue...)
				valueConsumed = true
			case "from":
				addr := common.HexToAddress(string(taskFrom))
				paddedAddr := common.LeftPadBytes(addr.Bytes(), 32)
				data = append(data, paddedAddr...)
			case "to":
				addr := common.HexToAddress(string(taskTo))
				paddedAddr := common.LeftPadBytes(addr.Bytes(), 32)
				data = append(data, paddedAddr...)
			case "contract":
				addr := common.HexToAddress(asset.Contract)
				paddedAddr := common.LeftPadBytes(addr.Bytes(), 32)
				data = append(data, paddedAddr...)
			}
		} else {
			var valStr string

			// get the param -- it's either user-passed or a default
			if p.Value != nil {
				switch valType := p.Value.(type) {
				case string:
					valStr = valType
				case map[interface{}]interface{}:
					nativeAsset := string(asset.NativeAsset)
					if p.Match == "dst_asset" {
						nativeAsset = string(task.DstAsset.GetNativeAsset().NativeAsset)
					}
					for k, v := range valType {
						// map keys are lowercase
						if strings.EqualFold(k.(string), nativeAsset) {
							valStr = fmt.Sprintf("%v", v)
						}
					}
				default:
					return to, value, data, fmt.Errorf("invalid config for task=%s value type=%T", task.ID(), valType)
				}
			} else {
				// no default param, first check that the user passed in the param
				if userPassedParamIndex >= len(input.Params) {
					return to, value, data, fmt.Errorf("not enough params passed in for this task")
				}
				valStr = input.Params[userPassedParamIndex]
				userPassedParamIndex++
			}

			// now we have the param in valStr -- we need to properly encode it
			switch p.Type {
			case "uint256":
				valBig := new(big.Int)
				if strings.HasPrefix(valStr, "0x") {
					// number in hex format
					_, ok := valBig.SetString(valStr, 0)
					if !ok {
						return to, value, data, fmt.Errorf("invalid task param, expected hex uint256: %s", valStr)
					}
				} else {
					// number in decimal format
					valDecimal, err := decimal.NewFromString(valStr)
					if err != nil {
						return to, value, data, fmt.Errorf("invalid task param, expected decimal: %s", valStr)
					}
					valBig = valDecimal.BigInt()
				}
				paddedValue := common.LeftPadBytes(valBig.Bytes(), 32)
				data = append(data, paddedValue...)
			case "address":
				addr := common.HexToAddress(valStr)
				paddedAddr := common.LeftPadBytes(addr.Bytes(), 32)
				data = append(data, paddedAddr...)
			}
		}
	}

	if valueConsumed {
		value = valueZero
	}

	return to, value, data, nil
}

func (txBuilder TxBuilder) BuildTaskTx(from xc.Address, to xc.Address, amount xc.AmountBlockchain, input xc.TxInput) (xc.Tx, error) {
	txInput := input.(*TxInput)
	asset := txBuilder.Asset.GetAssetConfig()

	txInput.GasLimit = 800_000
	if asset.NativeAsset == xc.KLAY {
		txInput.GasLimit = 2_000_000
	}
	if asset.NativeAsset == xc.ArbETH {
		txInput.GasLimit = 20_000_000
	}

	contract, value, payload, err := txBuilder.BuildTaskPayload(from, to, amount, txInput)
	if err != nil {
		return nil, err
	}

	return txBuilder.buildEvmTxWithPayload(xc.Address(contract), value, payload, txInput)
}

func (txBuilder TxBuilder) BuildProxyPayload(contract xc.ContractAddress, to xc.Address, amount xc.AmountBlockchain) ([]byte, error) {
	transferFnSignature := []byte("sendETH(uint256,address)")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]

	toAddress, err := HexToAddress(to)
	if err != nil {
		return nil, err
	}
	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)

	paddedAmount := common.LeftPadBytes(amount.Int().Bytes(), 32)

	var paddedContractAddress []byte
	if contract != "" {
		// log.Printf("sending token=%s", contract)
		transferFnSignature := []byte("sendTokens(address,uint256,address)")
		hash := sha3.NewLegacyKeccak256()
		hash.Write(transferFnSignature)
		methodID = hash.Sum(nil)[:4]

		contractAddress, err := HexToAddress(xc.Address(contract))
		if err != nil {
			return nil, err
		}
		paddedContractAddress = common.LeftPadBytes(contractAddress.Bytes(), 32)
	}
	// log.Print("Proxy methodID: ", hexutil.Encode(methodID)) // 0xa9059cbb for ERC20 transfer

	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedContractAddress...)
	data = append(data, paddedAmount...)
	data = append(data, paddedAddress...)

	return data, nil
}

func (txBuilder TxBuilder) BuildProxyTransferTx(from xc.Address, to xc.Address, amount xc.AmountBlockchain, input xc.TxInput) (xc.Tx, error) {
	txInput := input.(*TxInput)
	txInput.GasLimit = 400_000
	asset := txBuilder.Asset.GetAssetConfig()

	zero := xc.NewAmountBlockchainFromUint64(0)
	payload, err := txBuilder.BuildProxyPayload(xc.ContractAddress(asset.Contract), to, amount)
	if err != nil {
		return nil, err
	}

	// This is a special kind of smart contract that holds funds, but delegates signatures to an external signer
	// The Transfer.from is the smart contract, i.e. the Ethereum tx dst address
	// The Ethereum tx src address is the signer, specified in the Task definition (used, e.g. by the client to fetch the nonce, etc.)
	// The Transfer.to and/or Transfer.Asset.Contract (= Task.SrcAsset.Contract) are serialized in the payload.
	return txBuilder.buildEvmTxWithPayload(from, zero, payload, txInput)
}

func (txBuilder TxBuilder) BuildWormholePayload(taskFrom xc.Address, taskTo xc.Address, taskAmount xc.AmountBlockchain, txInput *TxInput) (string, xc.AmountBlockchain, []byte, error) {
	task := txBuilder.Asset.GetTask()

	contract, value, payload, err := txBuilder.BuildTaskPayload(taskFrom, taskTo, taskAmount, txInput)
	if err != nil {
		return contract, value, payload, err
	}

	// compute arbiterFee
	dstAsset := task.DstAsset.(*xc.TokenAssetConfig)
	priceUSD := dstAsset.Metadata.PriceUSD
	if priceUSD.String() == "0" {
		return contract, value, payload, fmt.Errorf("token price for %s is required to calculate arbiter fee", dstAsset.ID())
	}
	defaultArbiterFeeUsdStr, ok := task.DefaultParams["arbiter_fee_usd"]
	if !ok {
		return contract, value, payload, fmt.Errorf("invalid config: wormhole-transfer requires default_params.arbiter_fee_usd")
	}
	defaultArbiterFeeUsd := xc.NewAmountHumanReadableFromStr(fmt.Sprintf("%v", defaultArbiterFeeUsdStr))
	numTokens := defaultArbiterFeeUsd.Div(priceUSD)

	// - name: arbiterFee
	//   type: uint256
	arbiterFee := numTokens.ToBlockchain(dstAsset.Decimals)
	paddedValue := common.LeftPadBytes(arbiterFee.Int().Bytes(), 32)
	payload = append(payload, paddedValue...)

	// - name: nonce
	//   type: uint32
	nonceBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(nonceBytes, uint32(txInput.Nonce))
	paddedValue = common.LeftPadBytes(nonceBytes, 32)
	payload = append(payload, paddedValue...)

	return contract, value, payload, nil
}

func (txBuilder TxBuilder) BuildWormholeTransferTx(from xc.Address, to xc.Address, amount xc.AmountBlockchain, input xc.TxInput) (xc.Tx, error) {
	txInput := input.(*TxInput)
	asset := txBuilder.Asset.GetAssetConfig()

	txInput.GasLimit = 800_000
	if asset.NativeAsset == xc.KLAY {
		txInput.GasLimit = 2_000_000
	}

	contract, value, payload, err := txBuilder.BuildWormholePayload(from, to, amount, txInput)
	if err != nil {
		return nil, err
	}

	return txBuilder.buildEvmTxWithPayload(xc.Address(contract), value, payload, txInput)
}
