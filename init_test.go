package sdk

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
)

func TestInit(t *testing.T) {
	fn := func(path string) []byte {
		b, err := os.ReadFile(path)
		if err != nil {
			t.Fatalf("fail: %n", err)
		}
		return b
	}

	tests := []struct {
		name    string
		ops     []Option
		wantErr bool
	}{
		{
			name: "byFilePath",
			ops: []Option{
				WithCertPath("test/test_ca.crt", "test/test_client.crt", "test/test_client.key"),
				WithCustomer("a"),
				WithSecretPath("test/public_key.pem"),
			},
			wantErr: false,
		},
		{
			name: "byByte",
			ops: []Option{
				WithCertBytes(fn("test/test_ca.crt"), fn("test/test_client.crt"), fn("test/test_client.key")),
				WithSecretBytes(fn("test/public_key.pem")),
				WithCustomer("a"),
			},
			wantErr: false,
		},
		{
			name: "testMode",
			ops: []Option{
				WithCustomer("a"),
				WithSecretPath("test/public_key.pem"),
				WithTest(true),
			},
			wantErr: false,
		},
		{
			name:    "nil cfg",
			ops:     nil,
			wantErr: true,
		},
		{
			name: "err cfg",
			ops: []Option{
				WithCertBytes(fn("test/test_ca.crt"), nil, nil),

				WithCustomer("a"),
			},
			wantErr: true,
		},
		{
			name: "ca path err",
			ops: []Option{
				WithCertPath(("test_ca.crt"), ("test_client.crt"), ("test_client.key")),
				WithSecretPath("test/public_key.pem"),

				WithCustomer("a"),
			},
			wantErr: true,
		},
		{
			name: "ca err",
			ops: []Option{
				WithCertBytes(fn("test/test_ca.crt")[:10], nil, nil),
				WithSecretPath("test/public_key.pem"),

				WithCustomer("a"),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Init(tt.ops...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Init() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestClientAndServer(t *testing.T) {
	cfg := &config{
		caCertPath: "test/test_ca.crt",
		certPath:   "test/test_client.crt",
		keyPath:    "test/test_client.key",
		customer:   "a",
	}
	go func() {
		caCert, err := ioutil.ReadFile(cfg.caCertPath)
		if err != nil {
			t.Fatal(err)
		}

		// 创建 CA 证书池
		caCertPool := x509.NewCertPool()
		caCertPool.AppendCertsFromPEM(caCert)

		serverCert, err := tls.LoadX509KeyPair("test/test_server.crt", "test/test_server.key")
		if err != nil {
			t.Fatal(err)
		}

		config := &tls.Config{
			Certificates: []tls.Certificate{serverCert},
			ClientAuth:   tls.RequireAndVerifyClientCert,
			ClientCAs:    caCertPool,
		}

		fn := func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "ok")
		}
		// 创建 HTTPS 服务器
		server := &http.Server{
			Addr:      ":8443",
			TLSConfig: config,
			Handler:   http.HandlerFunc(fn),
		}
		server.ListenAndServeTLS("", "")
	}()

	ops := []Option{WithCertPath(("test/test_ca.crt"), ("test/test_client.crt"), ("test/test_client.key")),
		WithCustomer("a"),
		WithSecretPath("test/public_key.pem"),
	}
	cli, err := Init(ops...)
	if err != nil {
		t.Fatal(err)
	}
	resp, err := cli.client.Get("https://localhost:8443/")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	if string(body) != "ok" {
		t.Fatal(string(body))
	}
	t.Log(string(body))
}
