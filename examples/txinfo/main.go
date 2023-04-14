package main

import (
	"context"
	"fmt"

	"github.com/jumpcrypto/crosschain"
	"github.com/jumpcrypto/crosschain/factory"
)

func TxInfo(ctx context.Context, xc *factory.Factory, nativeAsset string, txHash string) {
	// get asset model, including config data
	// asset is used to create client, builder, signer, etc.
	asset, err := xc.GetAssetConfig("", nativeAsset)
	if err != nil {
		panic("unsupported asset")
	}

	// fetch tx info
	client, _ := xc.NewClient(asset)
	info, err := client.FetchTxInfo(ctx, crosschain.TxHash(txHash))
	if err != nil {
		panic(err)
	}
	info, _ = xc.EnrichDestinations(asset, info)
	fmt.Printf("%+v\n", info)
}

func main() {
	// initialize crosschain
	xc := factory.NewDefaultFactory()
	ctx := context.Background()
	TxInfo(ctx, xc, "INJ", "a8522e7c84d462011265cb6038b5a52f3028d4acff14d9565b3cbec0f4a2f800")
	TxInfo(ctx, xc, "BTC", "40a83018604b67cfc681b3ad7a8e3a9985f03060ba00df7e6a09af1edc93510a")
	TxInfo(ctx, xc, "ETH", "0x40ecf9af59bbf9ed1f4ea75610af87f801618ffa8902cd210477076a56b36f61")
}
