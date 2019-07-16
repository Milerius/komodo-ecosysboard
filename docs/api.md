# get information about one specific coin

Get all the information about a coin from the **komodo ecosystem**: ticker, last block, status, sync, notarized hash

**URL** : `/api/v1/tickers/:coin`

**Method** : `GET`

**Auth required** : No

**Permissions required** : None

## Success Response

**Code** : `200 OK`

**Content examples**

For a coin with ID `kmd` that exists in the **komodo** ecosystem.

```json
{
  "ticker": {
    "id": "kmd-komodo",
    "name": "Komodo",
    "symbol": "KMD",
    "rank": 45,
    "circulating_supply": 115003971,
    "total_supply": 115003971,
    "max_supply": 0,
    "beta_value": 1.01008,
    "last_updated": "2019-07-16T16:14:30Z",
    "quotes": {
      "USD": {
        "price": 1.30777359,
        "volume_24h": 5906789.3910286,
        "volume_24h_change_24h": -45.03,
        "market_cap": 150399156,
        "market_cap_change_24h": -1.91,
        "percent_change_1h": -1.55,
        "percent_change_12h": 3.46,
        "percent_change_24h": -1.91,
        "percent_change_7d": -15.64,
        "percent_change_30d": -15.4,
        "percent_change_1y": -18.05,
        "ath_price": 15.4149,
        "ath_date": "2017-12-21T08:04:00Z",
        "percent_from_price_ath": -91.52
      }
    }
  },
  "block_last_hash": "05eb998b46a6e67891e799f1007574ba144c62fbb4d8d358c4354a14c6aa0914",
  "status": {
    "info": {
      "version": 2001526,
      "protocolversion": 170007,
      "blocks": 1445516,
      "timeoffset": 0,
      "connections": 95,
      "proxy": "",
      "difficulty": 272890374.5314302,
      "testnet": false,
      "relayfee": 1e-06,
      "errors": "",
      "notarized": 1445500,
      "network": "livenet"
    }
  },
  "node_is_online": true,
  "node_is_synced": true,
  "notarizedhash": "0e87172ff6332df2ace74767eaa5cce1c247e76b22b014237a39991a51ad46fd",
  "notarizedtxid": [
    "d2b636b04981cb9af1df3aa13bc25336a24d5c45123ed22e33077c2c8d11f5c2",
    "d6d5adf13399a4f3f976c5e5a47e6a376463b4f27cbf44cfc0d33e98f31bc069",
    "6fcfe91bcd151ba4f1c741924da018a0ed95cf364ef93881dca1f80a5ce49cb8",
    "1e33b7ca09b24fabf9e9f151a86b37e16f27b4e406b2b1570da8cbae51686f85",
    "8a13c2ce50be7926b28ad91945c7ceefc8c703f49f0ab6999bda2e7df00d6112",
    "86cccd60100ecb6ac904aa837c68906ac6fb2e97db95b1ca28f8981377284bb2",
    "bbc6cde82664185d9fb9294c6314014a77d44891c516198f6487d7754729e67b"
  ]
}
```

Curl command: `curl http://127.0.0.1:8080/api/v1/tickers/kmd`

## Error Response

For a coin with ID `nonexistent` that exists in the **komodo** ecosystem.

**Code** : `404 Not found`

**Content examples**

```json
{
 "error": "This coin does not seem to be part of the komodo ecosystem"
}
```

Curl command: `curl http://127.0.0.1:8080/api/v1/tickers/nonexistent`
