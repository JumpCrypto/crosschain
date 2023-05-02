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
	// TxInfo(ctx, xc, "INJ", "a8522e7c84d462011265cb6038b5a52f3028d4acff14d9565b3cbec0f4a2f800")
	// TxInfo(ctx, xc, "SOL", "3mDyJibiCCXEgfyYcW21Cu9o89qZsgpr2J3n3fpF2EzoT9psDmfyFq3Lv5MxbvvrjkLVnk2KC1TAe7vSTjFcyGHV")
	TxInfo(ctx, xc, "APTOS", "0x15940935f6317d7a42085855aa8167106aff03aeff5528bed51da015940d3221")
}
