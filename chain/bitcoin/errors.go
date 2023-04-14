package bitcoin

import (
	"strings"

	xc "github.com/jumpcrypto/crosschain"
)

func CheckError(err error) xc.ClientError {
	msg := strings.ToLower(err.Error())
	if strings.Contains(msg, "txn-mempool-conflict") ||
		strings.Contains(msg, "bad-txns-inputs-missingorspent") {
		return xc.TransactionFailure
	}
	if strings.Contains(msg, "response body closed") ||
		strings.Contains(msg, "not found") ||
		strings.Contains(msg, "eof") {
		return xc.NetworkError
	}
	if strings.Contains(msg, "transaction already in block chain") ||
		strings.Contains(msg, "already known") {
		return xc.TransactionExists
	}
	return xc.UnknownError
}
