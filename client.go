package crosschain

import "context"

// Client is a client that can fetch data and submit tx to a public blockchain
type Client interface {
	FetchTxInput(ctx context.Context, from Address, to Address) (TxInput, error)
	FetchTxInfo(ctx context.Context, txHash TxHash) (TxInfo, error)
	SubmitTx(ctx context.Context, tx Tx) error
}

// GasEstimator is a specific Client that can estimate gas - not implemented yet
type GasEstimator interface {
}

// ClientBalance is a specific Client that can fetch balances - not implemented yet
type ClientBalance interface {
	// Fetch the balance of the asset that this client is configured for
	FetchBalance(ctx context.Context, address Address) (AmountBlockchain, error)
	FetchNativeBalance(ctx context.Context, address Address) (AmountBlockchain, error)
}

type FullClient interface {
	Client
	GasEstimator
	ClientBalance
}
