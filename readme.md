[golang sdk example](#sdk-example)

[document](#common)

[java sdk](https://github.com/affix6932/wallet-sdk-java)

```
                   ,.-^^-._          ┌─┐                                                                                                  
                  |-.____.-|         ║"│                                                                                                  
                  |        |         └┬┘                                                                                                  
                  |        |         ┌┼┐                                                                                                  
                  |        |          │             ┌────────┐                     ┌──────┐           ┌────┐           ┌──────┐           
                  '-.____.-'         ┌┴┐            │Platform│                     │Wallet│           │Addr│           │Chains│           
                     DB             User            └────┬───┘                     └───┬──┘           └──┬─┘           └───┬──┘           
                      │               │                  │                             │                 │                 │              
          ╔═══════════╪══════════╤════╪══════════════════╪═════════════════════════════╪═════════════════╪═════╗           │              
          ║ GET&DISPLAY ADDRESS  │    │                  │                             │                 │     ║           │              
          ╟──────────────────────┘    │  𝟏 do deposit    │                             │                 │     ║           │              
          ║           │               │─────────────────>│                             │                 │     ║           │              
          ║           │               │                  │                             │                 │     ║           │              
          ║           │               │                  │                             │                 │     ║           │              
          ║           │               │   ╔═════════════╤╪═════════════════════════════╪═════════════╗   │     ║           │              
          ║           │               │   ║ FIRST TIME  ││                             │             ║   │     ║           │              
          ║           │               │   ╟─────────────┘│   𝟐 get deposit address     │             ║   │     ║           │              
          ║           │               │   ║              │────────────────────────────>│             ║   │     ║           │              
          ║           │               │   ║              │                             │             ║   │     ║           │              
          ║           │               │   ║              │         𝟑 address           │             ║   │     ║           │              
          ║           │               │   ║              │<─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ │             ║   │     ║           │              
          ║           │               │   ╚══════════════╪═════════════════════════════╪═════════════╝   │     ║           │              
          ║           │               │                  │                             │                 │     ║           │              
          ║           │           𝟒 save||get            │                             │                 │     ║           │              
          ║           │<─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─│                             │                 │     ║           │              
          ║           │               │                  │                             │                 │     ║           │              
          ║           │               │𝟓 display address │                             │                 │     ║           │              
          ║           │               │<─ ─ ─ ─ ─ ─ ─ ─ ─│                             │                 │     ║           │              
          ╚═══════════╪═══════════════╪══════════════════╪═════════════════════════════╪═════════════════╪═════╝           │              
                      │               │                  │                             │                 │                 │              
                      │               │                  │                             │                 │                 │              
                      │  ╔══════════╤═╪══════════════════╪═════════════════════════════╪═════════════════╪═════════════════╪═════════════╗
                      │  ║ DEPOSIT  │ │                  │                             │                 │                 │             ║
                      │  ╟──────────┘ │  𝟔 get address   │                             │                 │                 │             ║
                      │  ║            │─────────────────>│                             │                 │                 │             ║
                      │  ║            │                  │                             │                 │                 │             ║
                      │  ║            │    𝟕 address     │                             │                 │                 │             ║
                      │  ║            │<─ ─ ─ ─ ─ ─ ─ ─ ─│                             │                 │                 │             ║
                      │  ║            │                  │                             │                 │                 │             ║
                      │  ║            │                  │        𝟖 transfer           │                 │                 │             ║
                      │  ║            │─────────────────────────────────────────────────────────────────>│                 │             ║
                      │  ║            │                  │                             │                 │                 │             ║
                      │  ║            │                  │                             │           𝟗 scan addr             │             ║
                      │  ║            │                  │                             │──────────────────────────────────>│             ║
                      │  ║            │                  │                             │                 │                 │             ║
                      │  ║            │                  │                             │       𝟏𝟎 get deposit info         │             ║
                      │  ║            │                  │                             │<─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ │             ║
                      │  ║            │                  │                             │                 │                 │             ║
                      │  ║            │                  │    𝟏𝟏 webhook callback      │                 │                 │             ║
                      │  ║            │                  │<─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ │                 │                 │             ║
                      │  ║            │                  │                             │                 │                 │             ║
                      │  ║            │                  │𝟏𝟐 check  transaction exists │                 │                 │             ║
                      │  ║            │                  │────────────────────────────>│                 │                 │             ║
                      │  ║            │                  │                             │                 │                 │             ║
                      │  ║            │                  │                             │                 │                 │             ║
                      │  ║            │   ╔══════╤═══════╪═════════════════════════════╪═════════════╗   │                 │             ║
                      │  ║            │   ║ ALT  │  transaction exists                 │             ║   │                 │             ║
                      │  ║            │   ╟──────┘       │                             │             ║   │                 │             ║
                      │  ║            │   ║              │    𝟏𝟑 transaction info      │             ║   │                 │             ║
                      │  ║            │   ║              │<─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ │             ║   │                 │             ║
                      │  ║            │   ╠══════════════╪═════════════════════════════╪═════════════╣   │                 │             ║
                      │  ║            │   ║ [not exists] │                             │             ║   │                 │             ║
                      │  ║            │   ║              │          𝟏𝟒 null            │             ║   │                 │             ║
                      │  ║            │   ║              │<─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ │             ║   │                 │             ║
                      │  ║            │   ╚══════════════╪═════════════════════════════╪═════════════╝   │                 │             ║
                      │  ╚════════════╪══════════════════╪═════════════════════════════╪═════════════════╪═════════════════╪═════════════╝
                     DB             User            ┌────┴───┐                     ┌───┴──┐           ┌──┴─┐           ┌───┴──┐           
                   ,.-^^-._          ┌─┐            │Platform│                     │Wallet│           │Addr│           │Chains│           
                  |-.____.-|         ║"│            └────────┘                     └──────┘           └────┘           └──────┘           
                  |        |         └┬┘                                                                                                  
                  |        |         ┌┼┐                                                                                                  
                  |        |          │                                                                                                   
                  '-.____.-'         ┌┴┐
```

withdraw

```
                      ┌─┐                                                       ,.-^^-._                
                      ║"│                                                      |-.____.-|               
                      └┬┘                                                      |        |               
                      ┌┼┐                                                      |        |               
                       │             ┌────────┐                                |        |       ┌──────┐
                      ┌┴┐            │Platform│                                '-.____.-'       │Wallet│
                     User            └────┬───┘                                   DB            └───┬──┘
                       │   𝟏 withdraw     │                                        │                │   
                       │─────────────────>│                                        │                │   
                       │                  │                                        │                │   
                       │                  │𝟐 check balance && freeze/deduct amount │                │   
                       │                  │───────────────────────────────────────>│                │   
                       │                  │                                        │                │   
                       │                  │                                        │                │   
          ╔════════════╪╤═════════════════╪════════════════════════════════════════╪═══════════╗    │   
          ║ NOT ENOUGH  │                 │                                        │           ║    │   
          ╟─────────────┘                 │         𝟑 not enough, cancel           │           ║    │   
          ║            │                  │<─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─│           ║    │   
          ║            │                  │                                        │           ║    │   
          ║            │𝟒 withdraw failed │                                        │           ║    │   
          ║            │<─ ─ ─ ─ ─ ─ ─ ─ ─│                                        │           ║    │   
          ╚════════════╪══════════════════╪════════════════════════════════════════╪═══════════╝    │   
                       │                  │                                        │                │   
                       │                  │                                        │                │   
                       │                  │                                        │                │   
                       │                  │                                        │                │   
                       │                  │                       𝟓 withdraw       │                │   
                       │                  │────────────────────────────────────────────────────────>│   
                       │                  │                                        │                │   
                       │                  │                          𝟔 ok          │                │   
                       │                  │<─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ │   
                       │                  │                                        │                │   
                       │                  │               𝟕 callback withdraw detail                │   
                       │                  │<─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ │   
                       │                  │                                        │                │   
                       │                  │                                        │                │   
                       │   ╔══════════════╤════════════════════════════════════════╪═══════════╗    │   
                       │   ║ DETAIL SUCC  │                                        │           ║    │   
                       │   ╟──────────────┘            𝟖 deduct amount             │           ║    │   
                       │   ║              │───────────────────────────────────────>│           ║    │   
                       │   ╠══════════════╪════════════════════════════════════════╪═══════════╣    │   
                       │   ║ [failed]     │                                        │           ║    │   
                       │   ║              │         𝟗 unfeeze/add amount           │           ║    │   
                       │   ║              │───────────────────────────────────────>│           ║    │   
                       │   ╚══════════════╪════════════════════════════════════════╪═══════════╝    │   
                     User            ┌────┴───┐                                   DB            ┌───┴──┐
                      ┌─┐            │Platform│                                 ,.-^^-._        │Wallet│
                      ║"│            └────────┘                                |-.____.-|       └──────┘
                      └┬┘                                                      |        |               
                      ┌┼┐                                                      |        |               
                       │                                                       |        |               
                      ┌┴┐                                                      '-.____.-'               
```

## Common

W-Broker must in header.

Current support:

| chain | coin |
|-------|------|
| TON   | TON  |
| TON   | USDT |
| TRON  | USDT |

resp always like:

```go
// resp from gateway: {
"code": 0, // from gateway. 0: success, others: failed,
"data": "{\"msg\":\"\",\"code\":0,\"data\":{\"chain\":\"TON\",\"hash\":\"og6hrZpiFxjsKsfRAr+CIQd\",\"address\":\"EQAkhkg79yqAqbG67Ch5m8j\",\"tag\":\"2\",\"coin\":\"USDT\",\"amount\":\"0.222222\",\"blockNo\":\"\",\"txId\":\"da2c87300d993cc924b9085af5f5183a\"}}"
}

api resp body in data:
{
"msg": "",
"code": 0,
"data": {
"chain": "TON",
"hash": "og6hrZpiFxjsKsfRAr+CIQdZp",
"address": "EQAkhkg79yqAqbG67Ch5m8j",
"tag": "2",
"coin": "USDT",
"amount": "0.222222",
"blockNo": "",
"txId": "da2c87300d993cc924b9085af5f5183a"
}
}

```

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

path:    `/v1/api/deposit/query_detail`

req:

| name      | type   | comment                                          | require |
|-----------|--------|--------------------------------------------------|---------|
| requestId | string | uniq id. bussinessId, not transactionId on chain | y       |
| address   | string | user address                                     | y       |
| tag       | string | user tag                                         | n       |

resp:

| name         | type           | comment                  |   |
|--------------|----------------|--------------------------|---|
| chain        | string         | chainName                |   |
| hash         | string         | hash from chain explorer |   |
| address      | string         | address                  |   |
| from         | string         | from                     |   |
| tag          | string         | tag                      |   |
| coin         | string         | coin, USDT/TON/...       |   |
| amount       | decimal        | 1.23456789               |   |
| blockNo      | string         | block info               |   |
| requestId    | string         | uniqueID                 |   |
| exchangeRate | decimal(40,18) |                          | - |
| fiatAmount   | decimal(40,18) |                          | - |
| symbol       | string         |                          | - |

### Deposit.QueryDetailByTxId

path:    `/v1/api/deposit/query_detail_by_txid`

req:

| name | type   | comment             | require |
|------|--------|---------------------|---------|
| txId | string | txid(hash) on chain | y       |

resp:

| name         | type           | comment                  |   |
|--------------|----------------|--------------------------|---|
| chain        | string         | chainName                |   |
| hash         | string         | hash from chain explorer |   |
| address      | string         | address                  |   |
| from         | string         | from                     |   |
| tag          | string         | tag                      |   |
| coin         | string         | coin, USDT/TON/...       |   |
| amount       | decimal        | 1.23456789               |   |
| blockNo      | string         | block info               |   |
| requestId    | string         | uniqueID                 |   |
| exchangeRate | decimal(40,18) |                          | - |
| fiatAmount   | decimal(40,18) |                          | - |
| symbol       | string         |                          | - |

### Deposit Callback Struct

| name         | type           | comment                  |   |
|--------------|----------------|--------------------------|---|
| chain        | string         | chainName                |   |
| hash         | string         | hash from chain explorer |   |
| address      | string         | address                  |   |
| from         | string         | from                     |   |
| tag          | string         | tag                      |   |
| coin         | string         | coin, USDT/TON/...       |   |
| amount       | decimal(40,18) | 0.123456                 |   |
| exchangeRate | decimal(40,18) |                          | - |
| fiatAmount   | decimal(40,18) |                          | - |
| symbol       | string         |                          | - |
| requestId    | string         | uniq ID                  | - |

### Withdraw.QueryDetail

path:    `/v1/api/withdraw/query_detail`

req:

| name      | type   | comment | require |
|-----------|--------|---------|---------|
| requestId | string | uniq id | y       |

resp:

| name         | type           | comment                                |   |
|--------------|----------------|----------------------------------------|---|
| requestId    | string         | uniq id                                |
| amount       | decimal        |                                        |
| coin         | string         |                                        |
| network      | string         |                                        |
| to           | string         |                                        |
| tag          | string         |                                        |
| hash         | string         |                                        |
| state        | int            | 0: pending, 1: process, 2: succ 3:fail |
| gas          | decimal        |                                        |
| exchangeRate | decimal(40,18) |                                        | - |
| fiatAmount   | decimal(40,18) |                                        | - |
| symbol       | string         |                                        | - |

### Withdraw.DoWithdraw

path:     `/v1/api/withdraw`

req:

| name      | type    | comment | require |
|-----------|---------|---------|---------|
| requestId | string  | uniq id | y       |
| amount    | decimal |         | y       |
| coin      | string  |         | y       |
| network   | string  |         | y       |
| to        | string  |         | y       |
| tag       | string  |         | y       |

resp:

| name | type | comment | |
|------|------|---------|-|
| code | int  |         |
| msg  | int  |         |

### Withdraw.DoWithdrawSymbol

path:     `/v1/api/withdraw_symbol`

req:

| name       | type    | comment           | require |
|------------|---------|-------------------|---------|
| requestId  | string  | uniq id           | y       |
| fiatAmount | decimal |                   | y       |
| symbol     | string  | USDT/EUR,USDT/BRL | y       |
| network    | string  |                   | y       |
| to         | string  |                   | y       |
| tag        | string  |                   | y       |

resp:

| name | type | comment | |
|------|------|---------|-|
| code | int  |         |
| msg  | int  |         |

### Withdraw Callback Struct

| name         | type           | comment       |   |
|--------------|----------------|---------------|---|
| network      | string         |               | - |
| coin         | string         |               | - |
| from         | string         |               | - |
| to           | string         |               | - |
| amount       | decimal(40,18) |               | - |
| requestId    | string         | uniq id       | - |
| hash         | string         | hash on chain | - |
| gas          | decimal(40,18) |               | - |
| tag          | string         |               | - |
| exchangeRate | decimal(40,18) |               | - |
| fiatAmount   | decimal(40,18) |               | - |
| symbol       | string         |               | - |

### Exchange Rate

path: `/v1/api/exchange_rate`

req:

| name   | type   | comment                        | require |
|--------|--------|--------------------------------|---------|
| symbol | string | USDT/USD,USDT/EUR,USDT/BRL ... | y       |

resp:

| name      | type            | comment                   |            |
|-----------|-----------------|---------------------------|------------|
| price     | decimal(40, 18) | USDT to USD/EUR/BRL price | 0.9999     |
| updatedAt | int             | timestamp(10)             | 1731470665 |

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

## java api

WIP
