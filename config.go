package sdk

import sdktrace "go.opentelemetry.io/otel/sdk/trace"

type config struct {
	caCert []byte
	cert   []byte
	key    []byte

	caCertPath string
	certPath   string
	keyPath    string

	isTest bool

	customer string

	secretPath  string
	secretBytes []byte

	provider *sdktrace.TracerProvider
}
