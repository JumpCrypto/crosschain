package crosschain

import "context"

// Client is a client that can fetch data and submit tx to a public blockchain
type Client interface {
	// UpdateAsset configures the client to use a different asset on the same chain
	UpdateAsset(assetCfg AssetConfig) error
	FetchTxInput(ctx context.Context, from Address, to Address) (TxInput, error)
	FetchTxInfo(ctx context.Context, txHash TxHash) (TxInfo, error)
	SubmitTx(ctx context.Context, tx Tx) error
}

// GasEstimator is a specific Client that can estimate gas - not implemented yet
type GasEstimator interface {
}

// ClientBalance is a specific Client that can fetch balances - not implemented yet
type ClientBalance interface {
	FetchBalance(ctx context.Context, address Address) (AmountBlockchain, error)

	// Fetch Balance for a specific asset on the same chain
	FetchBalanceForAsset(ctx context.Context, address Address, assetCfg AssetConfig) (AmountBlockchain, error)

	FetchNativeBalance(ctx context.Context, address Address) (AmountBlockchain, error)
}
