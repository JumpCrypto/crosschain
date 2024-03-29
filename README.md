# Crosschain

[![Go Reference](https://pkg.go.dev/badge/github.com/jumpcrypto/crosschain.svg)](https://pkg.go.dev/github.com/jumpcrypto/crosschain)
[![Coverage Status](https://coveralls.io/repos/github/JumpCrypto/crosschain/badge.svg?branch=main)](https://coveralls.io/github/JumpCrypto/crosschain?branch=main)

A Go library to interact with multiple blockchains.

Crosschain main design principle is to isolate network Client, Signer and tx Builder.
This way you can build applications or micro services using just what you need and with the convenience of a unified interface.

## Complete Example

See `examples/transfer/main.go`

Or run:
```
go run ./examples/transfer/main.go
```

<!-- ## [Documentation](https://pkg.go.dev/github.com/jumpcrypto/crosschain) -->

## Features

### Blockchains

- [x] Bitcoin
- [x] Bitcoin derived: Bitcoin Cash, Dogecoin
- [x] Ethereum
- [x] EVMs: Polygon, Binance Smart Chain, ...
- [x] Solana
- [x] Cosmos
- [x] Cosmos derived: Terra, Injective, XPLA, ...
- [ ] Polkadot
- [ ] Aptos
- [ ] Sui

### Assets

- [x] Native assets
- [x] Tokens
- [ ] NFTs
- [ ] Liquidity pools

### Operations

- [x] Balances (native asset, tokens)
- [x] Transfers (native transfers, token transfers)
- [x] Wraps/unwraps: ETH, SOL, ...
- [ ] Swaps
- [x] Crosschain transfers (via bridge): Wormhole
- [x] Tasks (generic smart contract calls, single tx): EVM
- [x] Pipelines (generic smart contract calls, multiple tx): EVM

## Contribute

We welcome contribution, whether in form of bug fixed, documentation, new chains, new functionality.

Just open an issue to discuss what you'd like to contribute and then submit a PR.

**Disclaimer. This alpha software has been open sourced. All software and code are provided “as is,” without any warranty of any kind, and should be used at your own risk.**
