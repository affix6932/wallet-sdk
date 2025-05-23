package sdk

import (
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"net/http"
	"os"

	"github.com/pkg/errors"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"

	"github.com/affix6932/wallet-sdk/encrypt"
)

func init() {
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{}, propagation.Baggage{}))
}

type WalletClient struct {
	client httpClient

	encrypt encrypt.Encrypt

	isTest bool

	customer string
}

type httpClient struct {
	*http.Client
	provider *sdktrace.TracerProvider
}

type Option func(cfg *config)

func WithCustomer(c string) Option {
	return func(cfg *config) {
		cfg.customer = c
	}
}

func WithOtelProvider(op *sdktrace.TracerProvider) Option {
	return func(cfg *config) {
		cfg.provider = op
	}
}

func WithCertPath(caPath, certPath, keyPath string) Option {
	return func(cfg *config) {
		cfg.caCertPath = caPath
		cfg.certPath = certPath
		cfg.keyPath = keyPath
	}
}

func WithCertBytes(caBytes, certBytes, keyBytes []byte) Option {
	return func(cfg *config) {
		cfg.caCert = caBytes
		cfg.cert = certBytes
		cfg.key = keyBytes
	}
}

func WithTest(b bool) Option {
	return func(cfg *config) {
		cfg.isTest = b
	}
}

func WithSecretPath(secretPath string) Option {
	return func(cfg *config) {
		cfg.secretPath = secretPath
	}
}
func WithSecretBytes(secretBytes []byte) Option {
	return func(cfg *config) {
		cfg.secretBytes = secretBytes
	}
}

func Init(ops ...Option) (*WalletClient, error) {
	cfg := &config{}
	for _, op := range ops {
		op(cfg)
	}

	if cfg.provider == nil {
		cfg.provider = sdktrace.NewTracerProvider(sdktrace.WithSampler(sdktrace.AlwaysSample()))
	}

	if len(cfg.customer) == 0 {
		return nil, ConfigMissCustomerErr
	}

	// not test mode must have ca/cert/key
	if !cfg.isTest && !(isByte(cfg) || isFilePath(cfg)) {
		return nil, ConfigMissCertErr
	}

	var (
		secret *rsa.PublicKey
		err    error
	)
	if !cfg.isTest {
		secret, err = encrypt.ToRsaPriKey(cfg.secretPath, cfg.secretBytes)
		if err != nil {
			return nil, errors.WithMessage(err, "encrypt secret failed")
		}
	}
	var client *http.Client

	// test mode
	if cfg.isTest {
		client = &http.Client{
			Transport: otelhttp.NewTransport(http.DefaultTransport),
		}
		return &WalletClient{client: httpClient{client, cfg.provider}, isTest: cfg.isTest, customer: cfg.customer}, nil
	}
	clientCert, caPool, err := initCert(cfg)
	if err != nil {
		return nil, err
	}

	tlsCfg := &tls.Config{
		Certificates: []tls.Certificate{*clientCert},
		RootCAs:      caPool,
	}

	tr := &http.Transport{TLSClientConfig: tlsCfg, Proxy: http.ProxyFromEnvironment}
	client = &http.Client{Transport: otelhttp.NewTransport(tr)}
	return &WalletClient{client: httpClient{client, cfg.provider}, encrypt: encrypt.Init(secret), isTest: cfg.isTest, customer: cfg.customer}, nil
}

// use path
func isFilePath(cfg *config) bool {
	return !(len(cfg.caCertPath) == 0 || len(cfg.certPath) == 0 || len(cfg.keyPath) == 0)
}

// use byte
func isByte(cfg *config) bool {
	return !(len(cfg.key) == 0 || len(cfg.caCert) == 0 || len(cfg.cert) == 0)
}

func initCert(cfg *config) (*tls.Certificate, *x509.CertPool, error) {
	var (
		clientCert tls.Certificate
		err        error
	)
	if isByte(cfg) {
		clientCert, err = tls.X509KeyPair(cfg.cert, cfg.key)
	} else {
		clientCert, err = tls.LoadX509KeyPair(cfg.certPath, cfg.keyPath)
	}

	if err != nil {
		return nil, nil, errors.WithMessage(err, "init client cert failed")
	}

	caCert := cfg.caCert
	if isFilePath(cfg) {
		caCert, err = os.ReadFile(cfg.caCertPath)
		if err != nil {
			return nil, nil, errors.WithMessage(err, "read caCert failed")
		}
	}

	caCertPool := x509.NewCertPool()
	if ok := caCertPool.AppendCertsFromPEM(caCert); !ok {
		return nil, nil, errors.WithMessage(err, "parse caCert failed")
	}

	return &clientCert, caCertPool, nil
}
