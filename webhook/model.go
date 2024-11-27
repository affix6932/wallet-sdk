package webhook

import (
	"github.com/shopspring/decimal"
)

// DepositCallbackMsg deposit callback
type DepositCallbackMsg struct {
	Chain        string          `json:"chain"`
	Hash         string          `json:"hash"`
	Address      string          `json:"address"`
	Coin         string          `json:"coin"`
	Amount       decimal.Decimal `json:"amount"`
	Tag          string          `json:"tag"`
	RequestId    string          `json:"requestId"`
	From         string          `json:"from"`
	FiatAmount   decimal.Decimal `json:"fiatAmount"`
	Symbol       string          `json:"symbol"`
	ExchangeRate decimal.Decimal `json:"exchangeRate"`
	Broker       string          `json:"broker"`
}

// WithdrawCallbackMsg callback
type WithdrawCallbackMsg struct {
	Network   string          `json:"network"`
	Coin      string          `json:"coin"`
	From      string          `json:"from"`
	To        string          `json:"to"`
	Amount    decimal.Decimal `json:"amount"`
	Tag       string          `json:"tag" `
	RequestID string          `json:"requestId"`
	Hash      string          `json:"hash"`
	GasFee    decimal.Decimal `json:"gas"`
	State     int             `json:"state"`
	Broker    string          `json:"broker"`
}

type CollectCallbackMsg struct {
	RequestId string          `json:"requestId"`
	Amount    decimal.Decimal `json:"amount"`
	Coin      string          `json:"coin"`
	Chain     string          `json:"chain"`
	To        string          `json:"to"`
	Tag       string          `json:"tag"`
	Hash      string          `json:"hash"`
	From      string          `json:"from"`
	Gas       decimal.Decimal `json:"gas"`
}
