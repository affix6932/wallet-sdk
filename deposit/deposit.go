package deposit

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"

	sdk "wallet-sdk"
)

type (
	QueryDetailReq struct {
	}

	QueryDetailResp struct{}
	GetNewAddrReq   struct {
		Network   string `json:"network"`
		RequestId string `json:"requestId"`
	}
	GetNewAddrResp struct {
		Address string `json:"address"`
		Tag     string `json:"tag"`
	}

	Resp[T QueryDetailResp | GetNewAddrResp] struct {
		Data T      `json:"data"`
		Code int    `json:"code"`
		Msg  string `json:"msg"`
	}
)

const (
	queryDetailRoute = "/v1/api/deposit/query_detail"
	getNewAddrRoute  = "/v1/api/deposit/get_new_address"
)

func NewDeposit(w *sdk.WalletClient, url string) *Deposit {
	return &Deposit{w: w, url: url}
}

type Deposit struct {
	w   *sdk.WalletClient
	url string
}

func (d *Deposit) QueryDetail(ctx context.Context, req *QueryDetailReq) (*QueryDetailResp, error) {
	r, err := buildReq(ctx, req, d.url, queryDetailRoute)
	if err != nil {
		return nil, err
	}
	body, err := d.w.Post(ctx, r)
	if err != nil {
		return nil, err
	}
	var resp = &Resp[QueryDetailResp]{}
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, err
	}
	if err = getErr(resp); err != nil {
		return nil, err
	}
	return &resp.Data, nil
}

func (d *Deposit) GetNewAddress(ctx context.Context, req *GetNewAddrReq) (*GetNewAddrResp, error) {
	r, err := buildReq(ctx, req, d.url, getNewAddrRoute)
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
	var resp = &Resp[GetNewAddrResp]{}
	if err := json.Unmarshal([]byte(tmp), resp); err != nil {
		return nil, err
	}
	if err = getErr(resp); err != nil {
		return nil, err
	}
	return &resp.Data, nil
}

func buildReq[T *QueryDetailReq | *GetNewAddrReq](ctx context.Context, req T, baseUrl, router string) (*http.Request, error) {
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

func getErr[T QueryDetailResp | GetNewAddrResp](resp *Resp[T]) error {
	if resp.Code == 0 {
		return nil
	}
	return errors.New(resp.Msg)
}
