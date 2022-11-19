package cosmos

import (
	"context"
	"errors"
	"net/http"
	"net/url"

	xc "github.com/jumpcrypto/crosschain"

	cosmClient "github.com/cosmos/cosmos-sdk/client"
	rpchttp "github.com/tendermint/tendermint/rpc/client/http"
	terraApp "github.com/terra-money/core/app"
)

// Client for Cosmos
type Client struct {
	cfg    xc.AssetConfig
	ctx    cosmClient.Context
	prefix string // e.g., terra
	coin   string // e.g., uluna, uusd
}

// TxInput for Cosmos
type TxInput struct {
	xc.TxInput
}

// NewClient returns a new Client
func NewClient(cfg xc.AssetConfig) (*Client, error) {
	host := cfg.URL
	httpClient, err := rpchttp.NewWithClient(
		host,
		"websocket",
		&http.Client{
			// Timeout: opts.Timeout,

			// We override the transport layer with a custom implementation as
			// there is an issue with the Cosmos SDK that causes it to
			// incorrectly parse URLs.
			Transport: newTransport(host, &http.Transport{}),
		})
	if err != nil {
		panic(err)
	}

	cosmosCfg := terraApp.MakeEncodingConfig()
	cliCtx := cosmClient.Context{}.
		WithClient(httpClient).
		WithCodec(cosmosCfg.Marshaler).
		WithTxConfig(cosmosCfg.TxConfig).
		WithLegacyAmino(cosmosCfg.Amino).
		WithInterfaceRegistry(cosmosCfg.InterfaceRegistry).
		WithChainID(string(cfg.ChainIDStr))

	// cliCtx.Client
	// app := terraApp.NewTerraApp()
	// app.E

	return &Client{
		cfg:    cfg,
		ctx:    cliCtx,
		prefix: cfg.ChainPrefix,
	}, nil
}

type transport struct {
	remote string
	proxy  http.RoundTripper
}

func newTransport(remote string, proxy http.RoundTripper) *transport {
	return &transport{
		remote: remote,
		proxy:  proxy,
	}
}

func (t *transport) RoundTrip(req *http.Request) (*http.Response, error) {
	u, err := url.Parse(t.remote)
	if err != nil {
		return nil, err
	}
	req.URL = u
	req.Host = u.Host

	// Proxy request.
	return t.proxy.RoundTrip(req)
}

// FetchTxInput returns tx input for a Cosmos tx
func (client *Client) FetchTxInput(ctx context.Context, from xc.Address, _ xc.Address) (xc.TxInput, error) {
	return TxInput{}, errors.New("not implemented")
}

// SubmitTx submits a Cosmos tx
func (client *Client) SubmitTx(ctx context.Context, txInput xc.Tx) error {
	return errors.New("not implemented")
}

// FetchTxInfo returns tx info for a Cosmos tx
func (client *Client) FetchTxInfo(ctx context.Context, txHash xc.TxHash) (xc.TxInfo, error) {
	return xc.TxInfo{}, errors.New("not implemented")
}
