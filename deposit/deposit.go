package deposit

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
	QueryDetailReq struct {
		RequestId string `json:"requestId"`
		Address   string `json:"address"`
		Tag       string `json:"tag"`
	}

	QueryDetailResp struct {
		Chain        string          `json:"chain"`
		Hash         string          `json:"hash"`
		Address      string          `json:"address"`
		From         string          `json:"from"`
		Tag          string          `json:"tag"`
		Coin         string          `json:"coin"`
		Amount       decimal.Decimal `json:"amount"`
		Confirm      string          `json:"blockNo"`
		RequestID    string          `json:"requestId"`
		FiatAmount   decimal.Decimal `json:"fiatAmount"`
		Symbol       string          `json:"symbol"`
		ExchangeRate decimal.Decimal `json:"exchangeRate"`
		Broker       string          `json:"broker"`
	}
	GetNewAddrReq struct {
		Network   string `json:"network"`
		RequestId string `json:"requestId"`
	}
	GetNewAddrResp struct {
		Address string `json:"address"`
		Tag     string `json:"tag"`
	}

	QueryDetailReqByTxID struct {
		TxId string `json:"txId"`
	}

	MinLimitReq struct {
		Coin   string          `json:"coin"`
		Chain  string          `json:"chain"`
		Amount decimal.Decimal `json:"amount,omitempty"`
	}

	MinLimitResp struct {
		CurrentAmount decimal.Decimal `json:"currentAmount"`
		DefaultAmount decimal.Decimal `json:"defaultAmount"`
	}

	Resp[T QueryDetailResp | GetNewAddrResp | MinLimitResp | json.RawMessage] struct {
		Data T      `json:"data"`
		Code int    `json:"code"`
		Msg  string `json:"msg"`
	}
)

const (
	queryDetailRoute       = "/v1/api/deposit/query_detail"
	getNewAddrRoute        = "/v1/api/deposit/get_new_address"
	queryDetailByTxIDRoute = "/v1/api/deposit/query_detail_by_txid"
	getDepositMinLimit     = "/v1/api/deposit/get_deposit_min_limit"
	setDepositMinLimit     = "/v1/api/deposit/set_deposit_min_limit"
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
	var tmp string
	if err := json.Unmarshal(body, &tmp); err != nil {
		return nil, err
	}
	var resp = &Resp[QueryDetailResp]{}
	if err := json.Unmarshal([]byte(tmp), &resp); err != nil {
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

func (d *Deposit) QueryDetailByTxID(ctx context.Context, req *QueryDetailReqByTxID) (*QueryDetailResp, error) {
	r, err := buildReq(ctx, req, d.url, queryDetailByTxIDRoute)
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
	var resp = &Resp[QueryDetailResp]{}
	if err := json.Unmarshal([]byte(tmp), resp); err != nil {
		return nil, err
	}
	if err = getErr(resp); err != nil {
		return nil, err
	}
	return &resp.Data, nil
}

func (d *Deposit) GetDepositMinLimit(ctx context.Context, req *MinLimitReq) (*MinLimitResp, error) {
	r, err := buildReq(ctx, req, d.url, getDepositMinLimit)
	if err != nil {
		return nil, err
	}
	body, err := d.w.Post(ctx, r)
	if err != nil {
		return nil, err
	}
	var resp = &Resp[MinLimitResp]{}

	return getResp(body, resp)
}

func (d *Deposit) SetDepositMinLimit(ctx context.Context, req *MinLimitReq) error {
	r, err := buildReq(ctx, req, d.url, setDepositMinLimit)
	if err != nil {
		return err
	}
	body, err := d.w.Post(ctx, r)
	if err != nil {
		return err
	}

	var resp = &Resp[json.RawMessage]{}
	_, err = getResp(body, resp)
	return err
}

func buildReq[T *QueryDetailReq | *GetNewAddrReq | *QueryDetailReqByTxID | *MinLimitReq](ctx context.Context, req T, baseUrl, router string) (*http.Request, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	r, err := http.NewRequest(http.MethodPost, baseUrl+router, bytes.NewBuffer(body))
	if err != nil {
		return nil, errors.WithStack(err)
	}
	r = r.WithContext(ctx)
	return r, nil
}

func getErr[T QueryDetailResp | GetNewAddrResp | MinLimitResp | json.RawMessage](resp *Resp[T]) error {
	if resp.Code == 0 {
		return nil
	}
	return errors.New(resp.Msg)
}

func getResp[T QueryDetailResp | GetNewAddrResp | MinLimitResp | json.RawMessage](body []byte, resp *Resp[T]) (*T, error) {
	var tmp string
	if err := json.Unmarshal(body, &tmp); err != nil {
		return nil, errors.Wrap(err, string(body))
	}
	if err := json.Unmarshal([]byte(tmp), resp); err != nil {
		return nil, err
	}
	if err := getErr(resp); err != nil {
		return nil, err
	}
	return &resp.Data, nil
}
