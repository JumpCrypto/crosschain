package crosschain

// TxBuilder is a Builder that can transfer assets
type TxBuilder interface {
	NewTransfer(from Address, to Address, amount AmountBlockchain, input TxInput) (Tx, error)
}

// TxTokenBuilder is a Builder that can transfer token assets, in addition to native assets
type TxTokenBuilder interface {
	TxBuilder
	NewNativeTransfer(from Address, to Address, amount AmountBlockchain, input TxInput) (Tx, error)
	NewTokenTransfer(from Address, to Address, amount AmountBlockchain, input TxInput) (Tx, error)
}
