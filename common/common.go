package common

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
	"github.com/shopspring/decimal"

	sdk "github.com/affix6932/wallet-sdk"
)

type common struct {
	w   *sdk.WalletClient
	url string
}

type Common = *common

func NewCommon(w *sdk.WalletClient, url string) Common {
	return &common{w, url}
}

const (
	queryExchangeRate = "/v1/api/exchange_rate"
	queryCurrentBlock = "/v1/api/current_block"
	queryGasFee       = "/v1/api/gasfee"
)

type QueryExchangeRateRequest struct {
	Symbol string `json:"symbol"`
}
type QueryExchangeRateResponse struct {
	Price     decimal.Decimal `json:"price"`
	UpdatedAt int64           `json:"updated_at"`
}

type QueryCurrentBlockRequest struct {
	Network string `json:"network"`
}
type QueryCurrentBlockResponse struct {
	Block int `json:"block"`
}

type QueryGasFeeRequest struct {
	Network string `json:"network"`
}

type QueryGasFeeResponse struct {
	Token    string          `json:"token"`
	Network  string          `json:"network"`
	Fee      decimal.Decimal `json:"fee"`
	GasToken string          `json:"gas_token"`
}

type Resp[T QueryExchangeRateResponse | QueryCurrentBlockResponse | []*QueryGasFeeResponse] struct {
	Data T      `json:"data"`
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (c *common) QueryExchangeRate(ctx context.Context, symbol string) (*QueryExchangeRateResponse, error) {
	req := &QueryExchangeRateRequest{
		Symbol: symbol,
	}
	r, err := buildReq(ctx, req, c.url, queryExchangeRate)
	if err != nil {
		return nil, err
	}
	body, err := c.w.Post(ctx, r)
	if err != nil {
		return nil, err
	}

	var tmp string
	if err := json.Unmarshal(body, &tmp); err != nil {
		return nil, err
	}
	var resp = &Resp[QueryExchangeRateResponse]{}
	if err := json.Unmarshal([]byte(tmp), resp); err != nil {
		return nil, err
	}
	if err = getErr(resp); err != nil {
		return nil, err
	}
	return &resp.Data, nil
}

func (c *common) QueryCurrentBlock(ctx context.Context, network string) (*QueryCurrentBlockResponse, error) {
	req := &QueryCurrentBlockRequest{
		Network: network,
	}
	r, err := buildReq(ctx, req, c.url, queryCurrentBlock)
	if err != nil {
		return nil, err
	}
	body, err := c.w.Post(ctx, r)
	if err != nil {
		return nil, err
	}

	var tmp string
	if err := json.Unmarshal(body, &tmp); err != nil {
		return nil, err
	}
	var resp = &Resp[QueryCurrentBlockResponse]{}
	if err := json.Unmarshal([]byte(tmp), resp); err != nil {
		return nil, err
	}
	if err = getErr(resp); err != nil {
		return nil, err
	}
	return &resp.Data, nil
}

func (c *common) QueryGas(ctx context.Context, network string) ([]*QueryGasFeeResponse, error) {
	req := &QueryGasFeeRequest{
		Network: network,
	}
	r, err := buildReq(ctx, req, c.url, queryGasFee)
	if err != nil {
		return nil, err
	}
	body, err := c.w.Post(ctx, r)
	if err != nil {
		return nil, err
	}

	var tmp string
	if err := json.Unmarshal(body, &tmp); err != nil {
		return nil, err
	}

	var resp = &Resp[[]*QueryGasFeeResponse]{}
	if err := json.Unmarshal([]byte(tmp), resp); err != nil {
		return nil, err
	}
	if err = getErr(resp); err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func buildReq[T *QueryExchangeRateRequest | *QueryCurrentBlockRequest | *QueryGasFeeRequest](ctx context.Context, req T, baseUrl, router string) (*http.Request, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	r, err := http.NewRequest(http.MethodPost, baseUrl+router, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	r = r.WithContext(ctx)
	return r, nil
}

func getErr[T QueryExchangeRateResponse | QueryCurrentBlockResponse | []*QueryGasFeeResponse](resp *Resp[T]) error {
	if resp.Code == 0 {
		return nil
	}
	return errors.New(resp.Msg)
}
