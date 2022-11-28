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
	asset, err := xc.GetAssetConfig("", "APTOS")
	if err != nil {
		panic("unsupported asset")
	}
	// address := xc.MustAddress(asset, "Hzn3n914JaSpnxo5mBbmuCDmGL6mxWN9Ac2HzEXFSGtb")
	// address := xc.MustAddress(asset, "xpla1hdvf6vv5amc7wp84js0ls27apekwxpr0ge96kg")
	// address := xc.MustAddress(asset, "terra1hw58t56mzszlnnkjak83ul8ff437ylrz57xj4v")
	address := xc.MustAddress(asset, "0xa589a80d61ec380c24a5fdda109c3848c082584e6cb725e5ab19b18354b2ab85")

	// fetch tx info
	client, _ := xc.NewClient(asset)
	balance, err := client.(crosschain.ClientBalance).FetchBalance(ctx, address)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", balance)
	humanBalance, _ := xc.ConvertAmountToHuman(asset, balance)
	fmt.Printf("%s\n", humanBalance)
}
