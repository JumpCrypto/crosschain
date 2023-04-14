# ERC-20

Tutorial: https://goethereumbook.org/en/smart-contract-read-erc20

Notes:
- Installed `solcjs` but couldn't really compile the example -- downloaded the abi directly
- Important: build `abigen` with the same version of `go-ethereum` as multichain is using
- Run `abigen --abi=erc20_sol_ERC20.abi --pkg=erc20 --out=erc20.go`
