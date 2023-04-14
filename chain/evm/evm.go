package evm

import "strings"

func TrimPrefixes(addressOrTxHash string) string {
	str := strings.TrimPrefix(addressOrTxHash, "0x")
	str = strings.TrimPrefix(str, "xdc")
	return str
}
