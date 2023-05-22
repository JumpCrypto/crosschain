package factory

import (
	"encoding/hex"
	"strings"

	xc "github.com/jumpcrypto/crosschain"
	"github.com/jumpcrypto/crosschain/chain/evm"
)

func (s *CrosschainTestSuite) TestEthWrap() {
	require := s.Require()
	asset, err := s.Factory.GetAssetConfig("WETH", "ETH")
	require.Nil(err)
	require.NotNil(asset)
	task, err := s.Factory.GetTaskConfig("eth-wrap", "ETH")
	require.Nil(err)
	require.NotNil(task)

	txBuilder, err := s.Factory.NewTxBuilder(task)
	require.Nil(err)
	require.NotNil(txBuilder)

	txInput := evm.TxInput{}
	tx, err := txBuilder.NewTransfer("from", "to", xc.NewAmountBlockchainFromUint64(123), &txInput)
	require.Nil(err)
	evmTx := tx.(*evm.Tx).EthTx
	require.Equal(uint8(0x2), evmTx.Type())
	require.Equal(uint64(800_000), evmTx.Gas())

	require.Equal(asset.GetAssetConfig().Contract, evmTx.To().String())
	require.Equal("123", evmTx.Value().String())
	expectedData := "d0e30db0"
	require.Equal(expectedData, hex.EncodeToString(evmTx.Data()))
}

func (s *CrosschainTestSuite) TestEthWrapPassingWETH() {
	require := s.Require()
	asset, err := s.Factory.GetAssetConfig("WETH", "ETH")
	require.Nil(err)
	require.NotNil(asset)
	task, err := s.Factory.GetTaskConfig("eth-wrap", "WETH.ETH")
	require.Nil(err)
	require.NotNil(task)

	txBuilder, err := s.Factory.NewTxBuilder(task)
	require.Nil(err)
	require.NotNil(txBuilder)

	txInput := evm.TxInput{}
	tx, err := txBuilder.NewTransfer("from", "to", xc.NewAmountBlockchainFromUint64(123), &txInput)
	require.Nil(err)
	evmTx := tx.(*evm.Tx).EthTx
	require.Equal(uint8(0x2), evmTx.Type())
	require.Equal(uint64(800_000), evmTx.Gas())

	require.Equal(asset.GetAssetConfig().Contract, evmTx.To().String())
	require.Equal("123", evmTx.Value().String())
	expectedData := "d0e30db0"
	require.Equal(expectedData, hex.EncodeToString(evmTx.Data()))
}

func (s *CrosschainTestSuite) TestEthUnwrap() {
	require := s.Require()
	asset, err := s.Factory.GetAssetConfig("WETH", "ETH")
	require.Nil(err)
	require.NotNil(asset)
	task, err := s.Factory.GetTaskConfig("eth-unwrap", "WETH.ETH")
	require.Nil(err)
	require.NotNil(task)

	txBuilder, err := s.Factory.NewTxBuilder(task)
	require.Nil(err)
	require.NotNil(txBuilder)

	txInput := evm.TxInput{}
	tx, err := txBuilder.NewTransfer("from", "to", xc.NewAmountBlockchainFromUint64(0x123), &txInput)
	require.Nil(err)
	evmTx := tx.(*evm.Tx).EthTx
	require.Equal(uint8(0x2), evmTx.Type())
	require.Equal(uint64(800_000), evmTx.Gas())

	require.Equal(asset.GetAssetConfig().Contract, evmTx.To().String())
	require.Equal("0", evmTx.Value().String())
	expectedData := "2e1a7d4d" +
		"0000000000000000000000000000000000000000000000000000000000000123"
	require.Equal(expectedData, hex.EncodeToString(evmTx.Data()))
}

func (s *CrosschainTestSuite) TestProxyTransfer() {
	require := s.Require()
	task, err := s.Factory.GetTaskConfig("proxy-transfer-eth", "ETH")
	require.Nil(err)
	require.NotNil(task)

	txBuilder, err := s.Factory.NewTxBuilder(task)
	require.Nil(err)
	require.NotNil(txBuilder)

	txInput := evm.TxInput{}
	from := "0x0eC9f48533bb2A03F53F341EF5cc1B057892B10B"
	to := "a0a5C02F0371cCc142ad5AD170C291c86c3E6379"
	tx, err := txBuilder.NewTransfer(xc.Address(from), xc.Address(to), xc.NewAmountBlockchainFromUint64(0x123), &txInput)
	require.Nil(err)
	evmTx := tx.(*evm.Tx).EthTx
	require.Equal(uint8(0x2), evmTx.Type())
	require.Equal(uint64(400_000), evmTx.Gas())

	require.Equal(from, evmTx.To().String())
	require.Equal("0", evmTx.Value().String())
	expectedData := "c664c714" +
		"0000000000000000000000000000000000000000000000000000000000000123" +
		"000000000000000000000000" + strings.ToLower(to)
	require.Equal(expectedData, hex.EncodeToString(evmTx.Data()))
}

func (s *CrosschainTestSuite) TestProxyTransferToken() {
	require := s.Require()
	asset, err := s.Factory.GetAssetConfig("USDC", "ETH")
	require.Nil(err)
	require.NotNil(asset)
	task, err := s.Factory.GetTaskConfig("proxy-transfer-erc20", "USDC.ETH")
	require.Nil(err)
	require.NotNil(task)

	txBuilder, err := s.Factory.NewTxBuilder(task)
	require.Nil(err)
	require.NotNil(txBuilder)

	txInput := evm.TxInput{}
	from := "0x0eC9f48533bb2A03F53F341EF5cc1B057892B10B"
	to := "a0a5C02F0371cCc142ad5AD170C291c86c3E6379"
	tx, err := txBuilder.NewTransfer(xc.Address(from), xc.Address(to), xc.NewAmountBlockchainFromUint64(0x123), &txInput)
	require.Nil(err)
	evmTx := tx.(*evm.Tx).EthTx
	require.Equal(uint8(0x2), evmTx.Type())
	require.Equal(uint64(400_000), evmTx.Gas())

	require.Equal(from, evmTx.To().String())
	require.Equal("0", evmTx.Value().String())
	expectedData := "4217e287" +
		"000000000000000000000000" + strings.TrimPrefix(asset.GetAssetConfig().Contract, "0x") +
		"0000000000000000000000000000000000000000000000000000000000000123" +
		"000000000000000000000000" + strings.ToLower(to)
	require.Equal(expectedData, hex.EncodeToString(evmTx.Data()))
}

func (s *CrosschainTestSuite) TestERC20Transfer() {
	require := s.Require()
	asset, err := s.Factory.GetAssetConfig("USDC", "ETH")
	require.Nil(err)
	require.NotNil(asset)
	task, err := s.Factory.GetTaskConfig("erc20-transfer", "USDC.ETH")
	require.Nil(err)
	require.NotNil(task)

	txBuilder, err := s.Factory.NewTxBuilder(task)
	require.Nil(err)
	require.NotNil(txBuilder)

	txInput := evm.TxInput{}
	from := "0x0eC9f48533bb2A03F53F341EF5cc1B057892B10B"
	to := "a0a5C02F0371cCc142ad5AD170C291c86c3E6379"
	tx, err := txBuilder.NewTransfer(xc.Address(from), xc.Address(to), xc.NewAmountBlockchainFromUint64(0x123), &txInput)
	require.Nil(err)
	evmTx := tx.(*evm.Tx).EthTx
	require.Equal(uint8(0x2), evmTx.Type())
	require.Equal(uint64(800_000), evmTx.Gas())

	require.Equal(asset.GetAssetConfig().Contract, strings.ToLower(evmTx.To().String()))
	require.Equal("0", evmTx.Value().String())
	expectedData := "a9059cbb" +
		"000000000000000000000000" + strings.ToLower(to) +
		"0000000000000000000000000000000000000000000000000000000000000123"
	require.Equal(expectedData, hex.EncodeToString(evmTx.Data()))

	// test that a token transfer produces the same result (except for gas limit)
	txBuilder, err = s.Factory.NewTxBuilder(asset)
	require.Nil(err)
	require.NotNil(txBuilder)

	tx2, err := txBuilder.NewTransfer(xc.Address(from), xc.Address(to), xc.NewAmountBlockchainFromUint64(0x123), &txInput)
	require.Nil(err)
	evmTx2 := tx2.(*evm.Tx).EthTx

	require.Equal(evmTx.To().String(), evmTx2.To().String())
	require.Equal(evmTx.Value(), evmTx2.Value())
	require.Equal(hex.EncodeToString(evmTx.Data()), hex.EncodeToString(evmTx2.Data()))
}

func (s *CrosschainTestSuite) TestWormholeApprove() {
	// get wormhole pipeline config (approve > transfer)
	require := s.Require()
	srcAsset, err := s.Factory.GetAssetConfig("WETH", "ETH")
	require.Nil(err)
	require.NotNil(srcAsset)
	dstAsset, err := s.Factory.GetAssetConfig("WETH", "MATIC")
	require.Nil(err)
	require.NotNil(dstAsset)

	tasks, err := s.Factory.GetMultiAssetConfig("WETH.ETH", "WETH.MATIC")
	require.Nil(err)
	require.Equal(2, len(tasks))

	// test wormhole approve only
	task := tasks[0]
	txBuilder, err := s.Factory.NewTxBuilder(task)
	require.Nil(err)
	require.NotNil(txBuilder)

	txInput := evm.TxInput{}
	tx, err := txBuilder.NewTransfer("from", "0x0eC9f48533bb2A03F53F341EF5cc1B057892B10B", xc.NewAmountBlockchainFromUint64(0x123), &txInput)
	require.Nil(err)
	evmTx := tx.(*evm.Tx).EthTx
	require.Equal(uint8(0x2), evmTx.Type())
	require.Equal(uint64(800_000), evmTx.Gas())

	require.Equal(srcAsset.GetAssetConfig().Contract, evmTx.To().String())
	require.Equal("0", evmTx.Value().String())
	expectedData := "095ea7b3" +
		"0000000000000000000000003ee18b2214aff97000d974cf647e7c347e8fa585" +
		"0000000000000000000000000000000000000000000000000000000000000123"
	require.Equal(expectedData, hex.EncodeToString(evmTx.Data()))
}

func (s *CrosschainTestSuite) TestWormholeTransfer() {
	// get wormhole pipeline config (approve > transfer)
	require := s.Require()
	srcAsset, err := s.Factory.GetAssetConfig("WETH", "ETH")
	require.Nil(err)
	require.NotNil(srcAsset)
	dstAsset, err := s.Factory.GetAssetConfig("WETH", "MATIC")
	require.Nil(err)
	require.NotNil(dstAsset)

	tasks, err := s.Factory.GetMultiAssetConfig("WETH.ETH", "WETH.MATIC")
	require.Nil(err)
	require.Equal(2, len(tasks))

	// test wormhole transfer only
	task := tasks[1]
	txBuilder, err := s.Factory.NewTxBuilder(task)
	require.Nil(err)
	require.NotNil(txBuilder)

	txInput := evm.TxInput{Nonce: 0x234}
	to := "0x0eC9f48533bb2A03F53F341EF5cc1B057892B10B"
	_, err = txBuilder.NewTransfer("from", xc.Address(to), xc.NewAmountBlockchainFromUint64(0x123), &txInput)
	require.Error(err, "token price for WETH.MATIC is required to calculate arbiter fee")

	task.GetTask().DstAsset.(*xc.TokenAssetConfig).PriceUSD = xc.NewAmountHumanReadableFromStr("2.5")
	tx, err := txBuilder.NewTransfer("from", xc.Address(to), xc.NewAmountBlockchainFromUint64(0x123), &txInput)
	require.Nil(err)
	evmTx := tx.(*evm.Tx).EthTx
	require.Equal(uint8(0x2), evmTx.Type())
	require.Equal(uint64(800_000), evmTx.Gas())

	require.Equal("0x3ee18B2214AFF97000D974cf647E7C347E8fa585", evmTx.To().String())
	require.Equal("0", evmTx.Value().String())
	expectedData := "0f5287b0" +
		"000000000000000000000000" + strings.TrimPrefix(srcAsset.GetAssetConfig().Contract, "0x") + // token
		"0000000000000000000000000000000000000000000000000000000000000123" + // amount
		"0000000000000000000000000000000000000000000000000000000000000005" + // chain
		"000000000000000000000000" + strings.TrimPrefix(to, "0x") + // recipient
		"0000000000000000000000000000000000000000000000001BC16D674EC80000" + // arbiterFee
		"0000000000000000000000000000000000000000000000000000000000000234" // nonce
	require.Equal(strings.ToLower(expectedData), hex.EncodeToString(evmTx.Data()))
}
