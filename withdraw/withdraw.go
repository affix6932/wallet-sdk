package withdraw

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
	"github.com/shopspring/decimal"

	sdk "github.com/affix6932/wallet-sdk"
)

type (
	Withdraw struct {
		w   *sdk.WalletClient
		url string
	}
	DoWithdrawReq struct {
		RequestId string          `json:"requestId"`
		Amount    decimal.Decimal `json:"amount"`
		Coin      string          `json:"coin"`
		Network   string          `json:"network"`
		To        string          `json:"to"`
		Tag       string          `json:"tag"`
	}
	DoWithdrawResp struct{}
	QueryReq       struct {
		RequestId string `json:"requestId"`
	}
	QueryResp struct {
		RequestId string          `json:"requestId"`
		Amount    decimal.Decimal `json:"amount"`
		Coin      string          `json:"coin"`
		Network   string          `json:"network"`
		To        string          `json:"to"`
		From      string          `json:"from"`
		Tag       string          `json:"tag"`
		Hash      string          `json:"hash"`
		State     int             `json:"state"`
		Gas       decimal.Decimal `json:"gas"`
		Broker    string          `json:"broker"`
	}

	Resp[T QueryResp | DoWithdrawResp] struct {
		Data T      `json:"data"`
		Code int    `json:"code"`
		Msg  string `json:"msg"`
	}
)

const (
	routeQueryWithdraw   = "/v1/api/withdraw/query_detail"
	routeWithdraw        = "/v1/api/withdraw"
	routeFinanceWithdraw = "/v1/api/financial_withdraw"
)

func NewWithdraw(w *sdk.WalletClient, url string) *Withdraw {
	return &Withdraw{w: w, url: url}
}

func buildReq[T *DoWithdrawReq | *QueryReq](ctx context.Context, req T, baseUrl, router string) (*http.Request, error) {
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

func getErr[T QueryResp | DoWithdrawResp](resp *Resp[T]) error {
	if resp.Code == 0 {
		return nil
	}
	return errors.New(resp.Msg)
}

func (d *Withdraw) QueryDetail(ctx context.Context, req *QueryReq) (*QueryResp, error) {
	r, err := buildReq(ctx, req, d.url, routeQueryWithdraw)
	if err != nil {
		return nil, err
	}
	body, err := d.w.Post(ctx, r)
	if err != nil {
		return nil, err
	}
	var tmp string
	if err := json.Unmarshal(body, &tmp); err != nil {
		return nil, err
	}
	var resp = &Resp[QueryResp]{}
	if err := json.Unmarshal([]byte(tmp), &resp); err != nil {
		return nil, err
	}
	if err = getErr(resp); err != nil {
		return nil, err
	}
	return &resp.Data, nil
}

func (d *Withdraw) DoWithdraw(ctx context.Context, req *DoWithdrawReq) (*DoWithdrawResp, error) {
	r, err := buildReq(ctx, req, d.url, routeWithdraw)
	if err != nil {
		return nil, err
	}
	body, err := d.w.Post(ctx, r)
	if err != nil {
		return nil, err
	}

	var tmp string
	if err := json.Unmarshal(body, &tmp); err != nil {
		return nil, err
	}
	var resp = &Resp[DoWithdrawResp]{}
	if err := json.Unmarshal([]byte(tmp), resp); err != nil {
		return nil, err
	}
	if err = getErr(resp); err != nil {
		return nil, err
	}
	return &resp.Data, nil
}

func (d *Withdraw) DoFinanceWithdraw(ctx context.Context, req *DoWithdrawReq) (*DoWithdrawResp, error) {
	r, err := buildReq(ctx, req, d.url, routeFinanceWithdraw)
	if err != nil {
		return nil, err
	}
	body, err := d.w.Post(ctx, r)
	if err != nil {
		return nil, err
	}

	var tmp string
	if err := json.Unmarshal(body, &tmp); err != nil {
		return nil, err
	}
	var resp = &Resp[DoWithdrawResp]{}
	if err := json.Unmarshal([]byte(tmp), resp); err != nil {
		return nil, err
	}
	if err = getErr(resp); err != nil {
		return nil, err
	}
	return &resp.Data, nil
}
