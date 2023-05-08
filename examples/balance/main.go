package main

import (
	"context"
	"fmt"

	"github.com/jumpcrypto/crosschain"
	"github.com/jumpcrypto/crosschain/factory"
)

func getBalanceSui(ctx context.Context, xc *factory.Factory) {
	asset, err := xc.GetAssetConfig("", "SUI")
	if err != nil {
		panic("unsupported asset")
	}
	address := xc.MustAddress(asset, "0x7d20dcdb2bca4f508ea9613994683eb4e76e9c4ed371169677c1be02aaf0b58e")
	client, _ := xc.NewClient(asset)
	balance, err := client.(crosschain.ClientBalance).FetchBalance(ctx, address)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Sui machine: %s\n", balance)
	humanBalance, _ := xc.ConvertAmountToHuman(asset, balance)
	fmt.Printf("Sui: %s\n", humanBalance)
}

func getBalanceAptos(ctx context.Context, xc *factory.Factory) {
	asset, err := xc.GetAssetConfig("", "APTOS")
	if err != nil {
		panic("unsupported asset")
	}
	address := xc.MustAddress(asset, "0xa589a80d61ec380c24a5fdda109c3848c082584e6cb725e5ab19b18354b2ab85")
	client, _ := xc.NewClient(asset)
	balance, err := client.(crosschain.ClientBalance).FetchBalance(ctx, address)
	if err != nil {
		panic(err)
	}
	fmt.Printf("APTOS octas: %s\n", balance)
	humanBalance, _ := xc.ConvertAmountToHuman(asset, balance)
	fmt.Printf("APTOS: %s\n", humanBalance)
}

func getBalanceInjective(ctx context.Context, xc *factory.Factory) {
	asset, err := xc.GetAssetConfig("", "INJ")
	if err != nil {
		panic("unsupported asset")
	}
	address := xc.MustAddress(asset, "inj162x3ax7z6ksquhshlqh6d498kr60qdx7wqf9we")
	client, _ := xc.NewClient(asset)
	balance, err := client.(crosschain.ClientBalance).FetchBalance(ctx, address)
	if err != nil {
		panic(err)
	}
	fmt.Printf("uinj: %s\n", balance)
	humanBalance, _ := xc.ConvertAmountToHuman(asset, balance)
	fmt.Printf("inj: %s\n", humanBalance)

	token, err := xc.GetAssetConfig("USDT", "INJ")
	if err != nil {
		panic(err)
	}
	client, _ = xc.NewClient(token)
	balance, err = client.(crosschain.ClientBalance).FetchBalance(ctx, address)
	if err != nil {
		panic(err)
	}
	humanBalance, _ = xc.ConvertAmountToHuman(token, balance)
	fmt.Printf("USDT.INJ: %s\n", humanBalance)
}

func getBalanceEthereum(ctx context.Context, xc *factory.Factory) {
	asset, err := xc.GetAssetConfig("", "ETH")
	if err != nil {
		panic("unsupported asset")
	}
	address := xc.MustAddress(asset, "0x0eC9f48533bb2A03F53F341EF5cc1B057892B10B")
	client, _ := xc.NewClient(asset)
	balance, err := client.(crosschain.ClientBalance).FetchBalance(ctx, address)
	if err != nil {
		panic(err)
	}
	fmt.Printf("gwei: %s\n", balance)
	humanBalance, _ := xc.ConvertAmountToHuman(asset, balance)
	fmt.Printf("ETH: %s\n", humanBalance)

	token, err := xc.GetAssetConfig("USDC", "ETH")
	if err != nil {
		panic(err)
	}
	client, _ = xc.NewClient(token)
	balance, err = client.(crosschain.ClientBalance).FetchBalance(ctx, address)
	if err != nil {
		panic(err)
	}
	humanBalance, _ = xc.ConvertAmountToHuman(token, balance)
	fmt.Printf("USDC.ETH: %s\n", humanBalance)
}

func getBalanceBitcoin(ctx context.Context, xc *factory.Factory) {
	asset, err := xc.GetAssetConfig("", "BTC")
	if err != nil {
		panic("unsupported asset")
	}
	address := xc.MustAddress(asset, "tb1q3ayl6ml4054lr2snhm5qx8mv6epulnqhv6k2kk")
	client, _ := xc.NewClient(asset)
	balance, err := client.(crosschain.ClientBalance).FetchBalance(ctx, address)
	if err != nil {
		panic(err)
	}
	fmt.Printf("satoshi: %s\n", balance)
	humanBalance, _ := xc.ConvertAmountToHuman(asset, balance)
	fmt.Printf("BTC: %s\n", humanBalance)
}

func main() {
	// initialize crosschain
	xc := factory.NewDefaultFactory()
	ctx := context.Background()
	getBalanceSui(ctx, xc)
	// getBalanceAptos(ctx, xc)
	// getBalanceInjective(ctx, xc)
	// getBalanceEthereum(ctx, xc)
	// getBalanceBitcoin(ctx, xc)
}
