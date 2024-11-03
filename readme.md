## support
### Deposit.GetNewAddress
### Deposit.QueryDetail

### Withdraw.QueryDetail
### Withdraw.DoWithdraw

## example
```go
    // prod
	ops := []sdk.Option{
		sdk.WithSecretPath("../test/public_key.pem"),
		sdk.WithCertPath(("../test/server/ca.crt"), ("../test/server/test_client.crt"), ("../test/server/test_client.key")),
		sdk.WithCustomer("a"),
	}
	w, err := sdk.Init(ops...)
	if err != nil {
		t.Fatal(err)
	}

	d := NewDeposit(w, URL)
	resp, err := d.GetNewAddress(context.Background(), &GetNewAddrReq{
		Network:   "TON",
		RequestId: "12345",
	})

```

```go
    // test
	ops := []sdk.Option{
		sdk.WithCustomer("a"),
		sdk.WithTest(true),
	}
	w, err := sdk.Init(ops...)
	if err != nil {
		t.Fatal(err)
	}

	d := NewDeposit(URL)
	resp, err := d.GetNewAddress(context.Background(), &GetNewAddrReq{
		Network:   "TON",
		RequestId: "12345",
	})
	if err != nil {
		t.Fatal(err)
	}
```

