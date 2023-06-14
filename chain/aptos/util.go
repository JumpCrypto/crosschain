package aptos

import (
	"encoding/hex"
	"strings"

	"github.com/coming-chat/go-aptos/aptostypes"
	transactionbuilder "github.com/coming-chat/go-aptos/transaction_builder"
	"github.com/coming-chat/lcs"
	xc "github.com/jumpcrypto/crosschain"
	"github.com/sirupsen/logrus"
)

func mustDecodeHex(h string) []byte {
	h = strings.Replace(h, "0x", "", 1)
	bz, err := hex.DecodeString(h)
	if err != nil {
		panic(err)
	}
	return bz
}

func valueFromTxPayload(payload transactionbuilder.TransactionPayload) xc.AmountBlockchain {
	zero := xc.NewAmountBlockchainFromUint64(0)
	switch payload := payload.(type) {
	case *transactionbuilder.TransactionPayloadEntryFunction:
		if len(payload.Args) > 1 {
			amount := uint64(0)
			err := lcs.Unmarshal(payload.Args[1], &amount)
			if err != nil {
				return zero
			}
			return xc.NewAmountBlockchainFromUint64(amount)
		}
	case *aptostypes.Payload:
		if payload.Function == "0x1::aptos_account::batch_transfer_coins" {
			// TODO - handle the case. Just skip for now
			logrus.Error("0x1::aptos_account::batch_transfer_coins not yet supported")
		} else if len(payload.Arguments) > 1 {
			return xc.NewAmountBlockchainFromStr(payload.Arguments[1].(string))
		}
	default:
		logrus.Errorf("unrecognized payload type: %T\n", payload)
	}
	return zero
}

func toFromTxPayload(payload transactionbuilder.TransactionPayload) xc.Address {
	switch payload := payload.(type) {
	case *transactionbuilder.TransactionPayloadEntryFunction:
		if len(payload.Args) > 0 {
			to_addr := payload.Args[0]
			return xc.Address(hex.EncodeToString(to_addr[:]))
		}
	case *aptostypes.Payload:
		if payload.Function == "0x1::aptos_account::batch_transfer_coins" {
			// TODO - handle the case. Just skip for now
			logrus.Error("0x1::aptos_account::batch_transfer_coins not yet supported")
		} else if len(payload.Arguments) > 0 {
			to_addr := payload.Arguments[0].(string)
			return xc.Address(to_addr)
		}
	default:
		logrus.Errorf("unrecognized payload type: %T\n", payload)
	}
	return ""
}
