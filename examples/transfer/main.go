package main

import (
	"context"
	"fmt"
	"time"

	"github.com/jumpcrypto/crosschain/factory"
)

func main() {
	// initialize crosschain
	xc := factory.NewDefaultFactory()
	ctx := context.Background()

	// get asset model, including config data
	// asset is used to create client, builder, signer, etc.
	asset, err := xc.GetAssetConfig("", "SOL")
	if err != nil {
		panic("unsupported asset")
	}

	panic("Please edit examples/transfer/main.go to set your testnet address and key")
	// set your own private key and address
	// you can get them, for example, from your Phantom wallet
	fromPrivateKey := xc.MustPrivateKey(asset, "...")
	from := xc.MustAddress(asset, "...")
	to := xc.MustAddress(asset, "Hzn3n914JaSpnxo5mBbmuCDmGL6mxWN9Ac2HzEXFSGtb")
	amount := xc.MustAmountBlockchain(asset, "0.005")

	// to create a tx, we typically need some input from the blockchain
	// e.g., nonce for Ethereum, recent block for Solana, gas data, ...
	// (network needed)
	client, _ := xc.NewClient(asset)

	input, err := client.FetchTxInput(ctx, from, to)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", input)

	// create tx
	// (no network, no private key needed)
	builder, _ := xc.NewTxBuilder(asset)
	tx, err := builder.NewTransfer(from, to, amount, input)
	if err != nil {
		panic(err)
	}
	sighash, err := tx.Sighash()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", tx)
	fmt.Printf("signing: %x\n", sighash)

	// sign the tx sighash
	// (private key needed)
	// for Solana, this is equivalent to:
	// fromPrivateKey := base58.Decode("...") // key exported from Phantom
	// signature := ed25519.Sign(ed25519.PrivateKey(fromPrivateKey), []byte(sighash))
	signer, _ := xc.NewSigner(asset)
	signature, err := signer.Sign(fromPrivateKey, sighash)
	if err != nil {
		panic(err)
	}
	fmt.Printf("signature: %x\n", signature)

	// complete the tx by adding its signature
	// (no network, no private key needed)
	err = tx.AddSignature(signature)
	if err != nil {
		panic(err)
	}

	// submit the tx, wait a bit, fetch the tx info
	// (network needed)
	fmt.Printf("tx id: %s\n", tx.Hash())
	err = client.SubmitTx(ctx, tx)
	if err != nil {
		panic(err)
	}
	fmt.Println("Zzz...")
	time.Sleep(20 * time.Second)
	info, err := client.FetchTxInfo(ctx, tx.Hash())
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", info)
}
