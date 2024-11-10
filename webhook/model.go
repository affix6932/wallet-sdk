package webhook

import (
	"github.com/shopspring/decimal"
)

// DepositCallbackMsg deposit callback
type DepositCallbackMsg struct {
	Chain     string          `json:"chain" `
	Hash      string          `json:"hash" `
	Address   string          `json:"address" `
	Coin      string          `json:"coin" `
	Amount    decimal.Decimal `json:"amount" `
	Tag       string          `json:"tag" `
	RequestId string          `json:"requestId" `
}

// WithdrawCallbackMsg callback
type WithdrawCallbackMsg struct {
	Chain     string          `json:"chain"`
	Coin      string          `json:"coin"`
	From      string          `json:"from"`
	To        string          `json:"to"`
	Amount    decimal.Decimal `json:"amount"`
	Tag       string          `json:"tag" `
	RequestID string          `json:"requestId"`
	Hash      string          `json:"hash"`
	GasFee    decimal.Decimal `json:"gas"`
}
