package sdk

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"time"

	"go.opentelemetry.io/otel/trace"

	"github.com/pkg/errors"
)

const (
	Wsign   = "w-sign"
	Wbroker = "w-broker"
	Wts     = "w-ts"
	Wnonce  = "w-nonce"
	Wsecret = "w-secret"

	wtest = "w-test"
)

type GWResp struct {
	Code int             `json:"code"`
	Msg  string          `json:"msg"`
	Data json.RawMessage `json:"data"`
}

func (w *WalletClient) postWithEncrypt(ctx context.Context, req *http.Request) ([]byte, error) {
	cli := w.client

	// set trace
	tr := cli.provider.Tracer("w-sdk")
	var span trace.Span
	if req.Context() != nil {
		ctx = req.Context()
	}
	ctx, span = tr.Start(ctx, "postWithEncrypt")
	defer span.End()
	req = req.WithContext(ctx)
	traceID := span.SpanContext().TraceID().String()

	encrypt := w.encrypt
	defer req.Body.Close()
	body, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, errors.Wrap(err, "read request body, traceID. traceID: "+traceID)
	}

	secret := generateRandomString(letters, 16)
	cipher, err := encrypt.AESEncryptECB([]byte(secret), body)
	if err != nil {
		return nil, errors.Wrap(err, "encrypt body err. traceID: "+traceID)
	}

	cipherS, err := w.encrypt.Encrypt([]byte(secret))
	if err != nil {
		return nil, errors.Wrap(err, "encrypt secret err. traceID: "+traceID)
	}

	req.Header.Set(Wsecret, cipherS)
	req.Header.Set(Wnonce, generateRandomString(digits, 6))
	req.Header.Set(Wts, strconv.Itoa(int(time.Now().UnixMilli())))
	req.Header.Set(Wbroker, w.customer)
	sign(req, body)

	req.Body = io.NopCloser(bytes.NewReader(cipher))
	req.ContentLength = int64(len(cipher))
	req.Header.Set("Content-Length", strconv.Itoa(len(cipher)))
	req.Header.Set("Content-Type", "application/json")

	resp, err := cli.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "http post request err. traceID: "+traceID)
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "read response body. traceID: "+traceID)
	}
	gwResp := &GWResp{}
	err = json.Unmarshal(b, gwResp)
	if err != nil {
		return nil, errors.Wrap(err, "json unmarshal err. traceID: "+traceID)
	}
	if gwResp.Code != 0 {
		return nil, errors.WithStack(errors.New(gwResp.Msg))
	}
	return gwResp.Data, nil
}

func (w *WalletClient) postWithoutEncrypt(ctx context.Context, req *http.Request) ([]byte, error) {
	cli := w.client
	req.Header.Set(Wbroker, w.customer)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set(wtest, "1")

	resp, err := cli.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "http post request err")
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "read response body")
	}
	gwResp := &GWResp{}
	err = json.Unmarshal(b, gwResp)
	if err != nil {
		return nil, errors.Wrap(err, "json unmarshal err")
	}
	if gwResp.Code != 0 {
		return nil, errors.WithStack(errors.New(gwResp.Msg))
	}
	return gwResp.Data, nil
}

func (w *WalletClient) Post(ctx context.Context, req *http.Request) ([]byte, error) {
	if w.isTest {
		return w.postWithoutEncrypt(ctx, req)
	}
	return w.postWithEncrypt(ctx, req)
}
