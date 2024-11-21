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
)

type QueryExchangeRateRequest struct {
	Symbol string `json:"symbol"`
}
type QueryExchangeRateResponse struct {
	Price     decimal.Decimal `json:"price"`
	UpdatedAt int64           `json:"updated_at"`
}
type Resp[T QueryExchangeRateResponse] struct {
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

func buildReq[T *QueryExchangeRateRequest](ctx context.Context, req T, baseUrl, router string) (*http.Request, error) {
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

func getErr[T QueryExchangeRateResponse](resp *Resp[T]) error {
	if resp.Code == 0 {
		return nil
	}
	return errors.New(resp.Msg)
}
