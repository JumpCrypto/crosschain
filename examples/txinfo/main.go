package main

import (
	"context"
	"fmt"

	"github.com/jumpcrypto/crosschain"
	"github.com/jumpcrypto/crosschain/factory"
)

func main() {
	// initialize crosschain
	xc := factory.NewDefaultFactory()
	ctx := context.Background()

	// get asset model, including config data
	// asset is used to create client, builder, signer, etc.
	// txHash := "22nfxos4FYb8xzQs4vkg1EDnqPAfSUt1UH2H3NvUXbZLpUuLHFXLKGUGTT4jWu3PRx6WT9u2hB4gmArQ18AAmfq5"
	// asset, err := xc.GetAssetConfig("", "SOL")
	txHash := "C302C19E895A72F602801C1BBFDAA13EBD7E319A4EB7C6485E260183F8099EE8"
	asset, err := xc.GetAssetConfig("", "LUNA")
	// txHash := "b10cf7cc68ba761307d1b0a07fdb1671e6917ac3d2c2dc9e7ed74ad9f506aa6e"
	// asset, err := xc.GetAssetConfig("", "XPLA")
	if err != nil {
		panic("unsupported asset")
	}

	// fetch tx info
	client, _ := xc.NewClient(asset)
	info, err := client.FetchTxInfo(ctx, crosschain.TxHash(txHash))
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", info)
}
