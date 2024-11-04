## Common

W-Broker must in header.

Current support:

| chain | coin |
|-------|------|
| TON   | TON  |
| TON   | USDT |
| TRC20 | USDT |

### Deposit.GetNewAddress

path: `/v1/api/deposit/get_new_address`

req:

| name      | type   | comment | require |
|-----------|--------|---------|---------|
| requestId | string | uniq id | y       |

resp:

| name    | type   | comment |  |
|---------|--------|---------|--|
| address | string | address |
| tag     | string | tag     |

### Deposit.QueryDetail

path: 	`/v1/api/deposit/query_detail`

req:

| name    | type   | comment      | require |
|---------|--------|--------------|---------|
| txId    | string | uniq id      | y       |
| address | string | user address | y       |
| tag     | string | user tag     | n       |

resp:

| name    | type    | comment                  | |
|---------|---------|--------------------------|-|
| chain   | string  | chainName                | |
| hash    | string  | hash from chain explorer | |
| address | string  | address                  | |
| tag     | string  | tag                      | |
| coin    | string  | coin, USDT/TON/...       | |
| amount  | decimal | 1.23456789               | |
| blockNo | string  | confirmCnt               | |
| txId    | string  | uniqueID                 | |

### Withdraw.QueryDetail

path: 	`/v1/api/withdraw/query_detail`

req:

| name      | type   | comment | require |
|-----------|--------|---------|---------|
| requestId | string | uniq id | y       |

resp:

| name      | type    | comment                                | |
|-----------|---------|----------------------------------------|-|
| requestId | string  | uniq id                                |
| amount    | decimal |                                        |
| coin      | string  |                                        |
| network   | string  |                                        |
| to        | string  |                                        |
| tag       | string  |                                        |
| hash      | string  |                                        |
| state     | int     | 0: pending, 1: process, 2: succ 3:fail |
| gas       | decimal |                                        |

### Withdraw.DoWithdraw
path:     `/v1/api/withdraw`

req:

| name      | type   | comment | require |
|-----------|--------|---------|---------|
| requestId | string | uniq id | y       |

resp:

| name      | type    | comment | |
|-----------|---------|---------|-|
| requestId | string  | uniq id |
| amount    | decimal |         |
| coin      | string  |         |
| network   | string  |         |
| to        | string  |         |
| tag       | string  |         |

## sdk example

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

