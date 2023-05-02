# Generated Go bindings for Sui Rust types

* The Rust codebase for Sui has a bunch of important types (i.e. Rust enums or
Rust structs). The most important one of them is
[`TransactionData`](https://github.com/MystenLabs/sui/blob/f4d0fb2870d22046859776570e73aa477e7081ce/crates/sui-types/src/messages.rs#L984),
which is used in Sui transactions. Specifically, the [`sui_executeTransactionBlock`](https://docs.sui.io/sui-jsonrpc#sui_executeTransactionBlock) endpoint in a full node expects a `tx_bytes` parameter; this parameter is a BCS-encoded instance of the `TransactionData` struct, in base64.

In other words, if you want to execute a Sui transaction (from Rust), you need to:
1) Create an instance of `TransactionData` in memory;
2) Serialize that `TransactionData` into a vector of bytes using BCS;
3) Encode that vector of bytes into base64;
4) Pass that base64 string to the `sui_executeTransactionBlock` endpoint.

This directory contains Go types (i.e. structs and interfaces) that mirror their corresponding Rust types, and can be used to generate valid BCS data from Go. This is achieved by using https://github.com/zefchain/serde-reflection, like this:


1) The Sui team exposes https://github.com/MystenLabs/sui/blob/main/crates/sui-core/tests/staged/sui.yaml, a YAML dump of all their Rust structs, generated with `serde-reflection`. We copy that file here as `sui_types.yaml`.
2) We then use `serde-generate` (also part of https://github.com/zefchain/serde-reflection) to transform that YAML dump into auto-generated Go code that we can use to build and serialize Sui Rust types.

For posterity, these are the commands I ran to generate all files in this directory:

```bash
cp ~/code/github.com/MystenLabs/sui/crates/sui-core/tests/staged/sui.yaml sui_types.yaml

serdegen --language go --target-source-dir ./generated --module-name bcs --with-runtimes bcs -- sui_types.yaml

gofmt -w generated/bcs/lib.go
```