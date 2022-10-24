# Crosschain

<!-- # [![Go Reference](https://pkg.go.dev/badge/github.com/jumpcrypto/crosschain.svg)](https://pkg.go.dev/github.com/jumpcrypto/crosschain) -->

**Disclaimer. This alpha software has been open sourced. All software and code are provided “as is,” without any warranty of any kind, and should be used at your own risk.**

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

- [ ] Bitcoin
- [ ] Bitcoin derived: Bitcoin Cash, Dogecoin
- [ ] Ethereum
- [ ] EVMs: Polygon, Binance Smart Chain, ...
- [x] Solana
- [ ] Cosmos
- [ ] Cosmos derived: Terra

### Assets

- [x] Native assets
- [x] Tokens
- [ ] NFTs
- [ ] Liquidity pools

### Operations

- [x] Transfers (native transfers, token transfers)
- [ ] Smart contract interactions

## Contribute

We welcome contribution, whether in form of bug fixed, documentation, new chains, new functionality.

Just open an issue to discuss what you'd like to contribute and then submit a PR.
