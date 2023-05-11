package evm

import (
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	xc "github.com/jumpcrypto/crosschain"
	"github.com/jumpcrypto/crosschain/chain/evm/erc20"
)

// Tx for EVM
type Tx struct {
	EthTx  *types.Transaction
	Signer types.Signer
	// parsed info
}

var _ xc.Tx = &Tx{}

type parsedTxInfo struct {
	Sources      []*xc.TxInfoEndpoint
	Destinations []*xc.TxInfoEndpoint
}

// Hash returns the tx hash or id
func (tx Tx) Hash() xc.TxHash {
	if tx.EthTx != nil {
		return xc.TxHash(tx.EthTx.Hash().Hex())
	}
	return xc.TxHash("")
}

// Sighashes returns the tx payload to sign, aka sighash
func (tx Tx) Sighashes() ([]xc.TxDataToSign, error) {
	if tx.EthTx == nil {
		return []xc.TxDataToSign{}, errors.New("transaction not initialized")
	}
	sighash := tx.Signer.Hash(tx.EthTx).Bytes()
	return []xc.TxDataToSign{sighash}, nil
}

// AddSignatures adds a signature to Tx
func (tx *Tx) AddSignatures(signatures ...xc.TxSignature) error {
	if tx.EthTx == nil {
		return errors.New("transaction not initialized")
	}

	signedTx, err := tx.EthTx.WithSignature(tx.Signer, signatures[0])
	if err != nil {
		return err
	}
	tx.EthTx = signedTx
	return nil
}

// Serialize returns the serialized tx
func (tx Tx) Serialize() ([]byte, error) {
	if tx.EthTx == nil {
		return []byte{}, errors.New("transaction not initialized")
	}
	return tx.EthTx.MarshalBinary()
}

// ParseTransfer parses a tx and extracts higher-level transfer information
func (tx *Tx) ParseTransfer(receipt *types.Receipt, nativeAsset xc.NativeAsset) parsedTxInfo {
	// 1. first try parsing as an abi we natively support.
	if tx.IsContract() {
		info, err := tx.ParseMultisendTransferTx()
		if err != nil {
			// ignore
		} else {
			return info
		}

		info, err = tx.ParseERC20TransferTx()
		if err != nil {
			// ignore
		} else {
			return info
		}
	}

	// 2. try parsing using the logs
	infoLogs := tx.parseReceipt(receipt, nativeAsset)
	if len(infoLogs.Destinations) > 0 {
		return infoLogs
	}

	// 3. use to/from/amount from the tf
	return parsedTxInfo{
		Sources: []*xc.TxInfoEndpoint{{
			Address: tx.From(),
		}},
		Destinations: []*xc.TxInfoEndpoint{{
			Address: tx.To(),
		}},
	}
}

func (tx *Tx) parseReceipt(receipt *types.Receipt, nativeAsset xc.NativeAsset) parsedTxInfo {
	loggedSources := []*xc.TxInfoEndpoint{}
	loggedDestinations := []*xc.TxInfoEndpoint{}
	for _, log := range receipt.Logs {
		event, _ := ERC20.EventByID(log.Topics[0])
		if event != nil && event.RawName == "Transfer" {
			erc20, _ := erc20.NewErc20(receipt.ContractAddress, nil)
			tf, err := erc20.ParseTransfer(*log)
			if err != nil {
				fmt.Println("could not parse log: ", log.Index)
				continue
			}
			loggedDestinations = append(loggedDestinations, &xc.TxInfoEndpoint{
				Address:         xc.Address(tf.To.String()),
				ContractAddress: xc.ContractAddress(log.Address.String()),
				Amount:          xc.AmountBlockchain(*tf.Tokens),
				NativeAsset:     nativeAsset,
			})
			loggedSources = append(loggedSources, &xc.TxInfoEndpoint{
				Address:         xc.Address(tf.From.String()),
				ContractAddress: xc.ContractAddress(log.Address.String()),
				Amount:          xc.AmountBlockchain(*tf.Tokens),
				NativeAsset:     nativeAsset,
			})
		}
	}
	return parsedTxInfo{
		Sources:      loggedSources,
		Destinations: loggedDestinations,
	}
}

// IsContract returns whether a tx is a contract or native transfer
func (tx Tx) IsContract() bool {
	if tx.EthTx == nil {
		return false
	}
	payload := tx.EthTx.Data()
	return len(payload) > 0
}

// From is the sender of a transfer
func (tx Tx) From() xc.Address {
	if tx.EthTx == nil || tx.Signer == nil {
		return xc.Address("")
	}

	from, err := types.Sender(tx.Signer, tx.EthTx)
	if err != nil {
		return xc.Address("")
	}
	return xc.Address(from.String())
}

// To is the account receiving a transfer
func (tx Tx) To() xc.Address {
	if tx.EthTx == nil {
		return xc.Address("")
	}
	if tx.IsContract() {
		info, err := tx.ParseERC20TransferTx()
		if err != nil {
			// ignore
		} else {
			// single token transfers have a single destination
			// we will opt to use instead.
			return info.Destinations[0].Address
		}
	}
	return xc.Address(tx.EthTx.To().String())
}

// Amount returns the tx amount
func (tx Tx) Amount() xc.AmountBlockchain {
	if tx.EthTx == nil {
		return xc.NewAmountBlockchainFromUint64(0)
	}
	info, err := tx.ParseERC20TransferTx()
	if err != nil {
		// ignore
	} else {
		// if this is a erc20 transfer, we use it's amount
		return info.Destinations[0].Amount
	}
	return xc.AmountBlockchain(*tx.EthTx.Value())
}

// ContractAddress returns the contract address for a token transfer
func (tx Tx) ContractAddress() xc.ContractAddress {
	if tx.IsContract() {
		return xc.ContractAddress(tx.EthTx.To().String())
	}
	return xc.ContractAddress("")
}

// Fee returns the fee associated to the tx
func (tx Tx) Fee(baseFeeUint uint64, gasUsedUint uint64) xc.AmountBlockchain {
	// from Etherscan: (BaseFee + MaxPriority)*GasUsed
	maxPriority := xc.AmountBlockchain(*tx.EthTx.GasTipCap())
	gasUsed := xc.NewAmountBlockchainFromUint64(gasUsedUint)
	baseFee := xc.NewAmountBlockchainFromUint64(baseFeeUint)
	baseFeeAndPriority := baseFee.Add(&maxPriority)
	fee1 := gasUsed.Mul(&baseFeeAndPriority)

	// old gas price * gas used
	gasPrice := xc.AmountBlockchain(*tx.EthTx.GasPrice())
	fee2 := gasPrice.Mul(&gasUsed)

	if fee1.Cmp(&fee2) < 0 {
		return fee1
	}
	return fee2
}

// ParseERC20TransferTx parses the tx payload as ERC20 transfer
func (tx Tx) ParseERC20TransferTx() (parsedTxInfo, error) {
	payload := tx.EthTx.Data()
	if len(payload) != 4+32*2 || hex.EncodeToString(payload[:4]) != "a9059cbb" {
		return parsedTxInfo{}, errors.New("payload is not ERC20.transfer(address,uint256)")
	}

	var buf1 [20]byte
	copy(buf1[:], payload[4+12:4+32])
	to := xc.Address(common.Address(buf1).String())
	if !strings.HasPrefix(string(to), "0x") {
		to = "0x" + to
	}

	var buf2 [32]byte
	copy(buf2[:], payload[4+32:4+2*32])
	amount := new(big.Int).SetBytes(buf2[:])

	return parsedTxInfo{
		// the from should be the tx sender
		Sources: []*xc.TxInfoEndpoint{{
			Address: tx.From(),
			Amount:  xc.AmountBlockchain(*amount),
		}},
		// destination
		Destinations: []*xc.TxInfoEndpoint{{
			Address:         to,
			ContractAddress: tx.ContractAddress(),
			Amount:          xc.AmountBlockchain(*amount),
		}},
	}, nil
}

// ParseMultisendTransferTx parses the tx payload as multi-send transfer
func (tx Tx) ParseMultisendTransferTx() (parsedTxInfo, error) {
	res := parsedTxInfo{}
	payload := tx.EthTx.Data()

	abiSigETH := "1a1da075"
	abiSigERC20 := "ca350aa6"
	if len(payload) < 4 {
		return res, errors.New("evm payload is miscellaneous")
	}
	abiSig := hex.EncodeToString(payload[:4])
	if abiSig != abiSigETH && abiSig != abiSigERC20 {
		// log.Printf("invalid abi=%s", abiSig)
		return res, errors.New("payload is not multisend (1a1da075 or ca350aa6)")
	}

	offset := 4 + 32*2 // ignore first 2 params

	// read array len
	var buf20 [20]byte
	var buf [32]byte
	copy(buf[:], payload[offset:offset+32])
	offset += 32

	arrayLen := int(new(big.Int).SetBytes(buf[:]).Int64())
	numParams := 2
	if abiSig == abiSigERC20 {
		numParams = 3
	}
	if len(payload) != offset+numParams*32*arrayLen {
		// log.Printf("invalid payload len=%d / %d", len(buf), offset+numParams*32*arrayLen)
		return res, errors.New("payload is not multisend (len)")
	}

	for i := 0; i < arrayLen; i++ {
		res.Destinations = append(res.Destinations, &xc.TxInfoEndpoint{})
		if abiSig == abiSigETH {
			res.Destinations[i].Asset = "ETH"
			// to
			copy(buf20[:], payload[offset+12:offset+32])
			res.Destinations[i].Address = xc.Address("0x" + common.Address(buf20).String())
			offset += 32
			// amount
			copy(buf[:], payload[offset:offset+32])
			res.Destinations[i].Amount = xc.AmountBlockchain(*new(big.Int).SetBytes(buf[:]))
			offset += 32
		} else if abiSig == abiSigERC20 {
			// asset
			copy(buf20[:], payload[offset+12:offset+32])
			res.Destinations[i].ContractAddress = xc.ContractAddress("0x" + common.Address(buf20).String())
			offset += 32
			// to
			copy(buf20[:], payload[offset+12:offset+32])
			res.Destinations[i].Address = xc.Address("0x" + common.Address(buf20).String())
			offset += 32
			// amount
			copy(buf[:], payload[offset:offset+32])
			res.Destinations[i].Amount = xc.AmountBlockchain(*new(big.Int).SetBytes(buf[:]))
			offset += 32
		}
	}

	// a single source, the contract
	res.Sources = append(res.Sources, &xc.TxInfoEndpoint{
		Address: xc.Address(tx.ContractAddress()),
	})

	return res, nil
}
