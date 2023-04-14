package crosschain

import "context"

// Client is a client that can fetch data and submit tx to a public blockchain
type Client interface {
	FetchTxInput(ctx context.Context, from Address, to Address) (TxInput, error)
	FetchTxInfo(ctx context.Context, txHash TxHash) (TxInfo, error)
	SubmitTx(ctx context.Context, tx Tx) error
}

type EstimateGasFunc func(native NativeAsset) (AmountBlockchain, error)

// GasEstimator is a specific Client that can estimate gas - not implemented yet
type GasEstimator interface {
	EstimateGas(ctx context.Context) (AmountBlockchain, error)
	RegisterEstimateGasCallback(fn EstimateGasFunc)
}

// ClientBalance is a specific Client that can fetch balances
type ClientBalance interface {
	// Fetch the balance of the asset that this client is configured for
	FetchBalance(ctx context.Context, address Address) (AmountBlockchain, error)
	FetchNativeBalance(ctx context.Context, address Address) (AmountBlockchain, error)
}

type FullClient interface {
	Client
	ClientBalance
}

type FullClientWithGas interface {
	Client
	ClientBalance
	GasEstimator
}

type ClientError string

// A transaction terminally failed due to no balance
const NoBalance ClientError = "NoBalance"

// A transaction terminally failed due to no balance after accounting for gas cost
const NoBalanceForGas ClientError = "NoBalanceForGas"

// A transaction terminally failed due to another reason
const TransactionFailure ClientError = "TransactionFailure"

// A transaction failed to submit because it already exists
const TransactionExists ClientError = "TransactionExists"

// A network error occured -- there may be nothing wrong with the transaction
const NetworkError ClientError = "NetworkError"

// No outcome for this error known
const UnknownError ClientError = "UnknownError"
