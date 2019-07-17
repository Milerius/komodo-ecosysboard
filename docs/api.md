# Rest Api Doc

## search dexstats information

Retrieves an url from an explorer from a block, addresses or transaction

**URL** : `/api/v1/dexstats/:coin/search`

**Method** : `POST`

**Auth required** : No

**Permissions required** : None

**Data constraints**

```json
{
    "input": "[valid input (transaction, block hash, block height, or address]"
}
```

**Data example**

```json
{
    "input": "06d6747a49097830574cf8d33e399d8a8679e457493cd17390a80d0f916287bc"
}
```

## Success Response

**Code** : `200 OK`

**Content example**

```json
{
    "url_to_redirect": "http://kmd.explorer.dexstats.info/block/06d6747a49097830574cf8d33e399d8a8679e457493cd17390a80d0f916287bc"
}
```

Curl command: `curl -X POST http://127.0.0.1:8080/api/v1/dexstats/kmd/search -H 'Content-Type: application/json'  -d '{"input": "06d6747a49097830574cf8d33e399d8a8679e457493cd17390a80d0f916287bc"}'`

## get information about all coins

Get all the information about all the coins from the **komodo ecosystem**: ticker, last block, status, sync, notarized hash

**URL** : `/api/v1/tickers/`

**Method** : `GET`

**Auth required** : No

**Permissions required** : None

### Success Response

**Code** : `200 OK`

**Content examples**

```json
[
  {
    "ticker": {
      "id": "kmd-komodo",
      "name": "Komodo",
      "symbol": "KMD",
      "rank": 45,
      "circulating_supply": 115014195,
      "total_supply": 115014195,
      "max_supply": 0,
      "beta_value": 1.01215,
      "last_updated": "2019-07-16T20:28:03Z",
      "quotes": {
        "USD": {
          "price": 1.23522226,
          "volume_24h": 5612652.6659111,
          "volume_24h_change_24h": -41.25,
          "market_cap": 142068093,
          "market_cap_change_24h": -6.1,
          "percent_change_1h": 2,
          "percent_change_12h": -9.67,
          "percent_change_24h": -6.11,
          "percent_change_7d": -19.58,
          "percent_change_30d": -22.2,
          "percent_change_1y": -23.01,
          "ath_price": 15.4149,
          "ath_date": "2017-12-21T08:04:00Z",
          "percent_from_price_ath": -91.99
        }
      }
    },
    "block_last_hash": "0057f55901bd484b99351d0f78b6c46c51991ffa5670d51429db0d0bb92eb0ac",
    "status": {
      "info": {
        "version": 2001526,
        "protocolversion": 170007,
        "blocks": 1445762,
        "timeoffset": 0,
        "connections": 107,
        "proxy": "",
        "difficulty": 227217655.8847262,
        "testnet": false,
        "relayfee": 1e-06,
        "errors": "",
        "notarized": 1445740,
        "network": "livenet"
      }
    },
    "node_is_online": true,
    "node_is_synced": true,
    "notarizedhash": "00000000d85406d8618a80d54d440467ad3e14478d07c872fbd7f2149788befe",
    "notarizedtxid": [
      "6d05b20f94275b651755fcb8bec325792dd704997b3fa225937aa046780e5258",
      "e81cab15a3c960ffefb0f0c732d25b26f69a8acd262927fd69b1c683f76b17eb",
      "385793e322c946f8bb73b7c3718c2a53e0da5995d919b8d2eb67166e7b8a856e",
      "55a5d72b1f38d41aef052e58ac732095768adabb4e0817f699a684905e8d676f",
      "7e562873ae2b076a54400d317383864bd989c4ba838d0cb341c8c98d461d60a6",
      "e5f186a5d52adbcf851773df6297545fd6f44ac03b62adac95ade3d2f0683c90",
      "48305b8936080a56b93b7c9aa20b6b46659c5933b92348c56c2af0de5930f8c9",
      "7ae1780f32916a792f1256aa1c8602f47acbdaff6e6cfb64eeccb8f60265b5e1",
      "7a3fe3b3531528af135a51f92e03e990ae08606debac6f33874bc1560c2c7ced",
      "f61bdcd7f1cfdd2f57627f364a018e8a43e7e74bbdf503e064b7e49956802566",
      "80af84ceea4f1e4ee8617bd2624f75c72bab8eec9dfd490492d4dae2e43ff38b",
      "10eb0caf861d1773604952ebdfaeb22b85dd56bb092ccdc789a45ad789e0e1cc",
      "46939bd658fae12df343f5f4a022416a113ea531ce5e2f72ed51fb0e2abd2d55",
      "e5faf709c0aa9bd1a6760e23c530334949e4ea792b014e7c6c537c33b5aec888",
      "8d15a92ed8948bcfb8cd217178a1e7b52c121197827ad2c7e8f3afb0c5afcc02",
      "38625cdf56df69014762e21de2380a28c161790aded7a305a0ef92226bad1e90",
      "9488579a817487769c4d5475e3e186b385d434f5658e4ec1abc27584fb39ccbb",
      "3b67600ea0079080328870248ea2844942b23191e0afd6898029843e25b915e7",
      "e5b12d5e6d5601e3f69854fbf1f78b13872792595766c82532bb69d04bf82009",
      "0b911dafea7bcb889155f97eb49d6f93a4381a8ff1959dfaf06785f1fa11f45f",
      "e74809d30fd7dc70f871b4b8d067320d7b2303b17e400a33ead4f92df979b31e",
      "6aaf8db0530ef9ed72c7f04b520cf38e7c69085abfb473a6dd7dd114aa12b867",
      "17322ebaa664f2cb8ca84cbb81a28dab42ab685a126e0de17b331a44f408b208",
      "93909611030f07bb10c327b0a79bb9f43c3d1ba8a116fccdb2e881d7c62e5b6e",
      "30ba2d2154edb74dd08db72b036fbaf43c2d6be5d5c3f09852663eae7a67ae7f",
      "6a212b8ee7f397f68c9dfac6d3df1b07700329cefcb41f74f5b8f02b0146a23a",
      "d885b24835a653d4e16643a8d8326f3b7dd1e84383939058360a9e3d71b26488"
    ]
  },
  {
    "ticker": {
      "id": "vrsc-verus-coin",
      "name": "Verus Coin",
      "symbol": "VRSC",
      "rank": 391,
      "circulating_supply": 42766571,
      "total_supply": 0,
      "max_supply": 83540184,
      "beta_value": 0.670258,
      "last_updated": "2019-07-16T20:27:44Z",
      "quotes": {
        "USD": {
          "price": 0.15010592,
          "volume_24h": 2275.52840557,
          "volume_24h_change_24h": 5841.72,
          "market_cap": 6419515,
          "market_cap_change_24h": -1.87,
          "percent_change_1h": 0.74,
          "percent_change_12h": 6.8,
          "percent_change_24h": -1.94,
          "percent_change_7d": -31.88,
          "percent_change_30d": -27.27,
          "percent_change_1y": 0,
          "ath_price": 0.3854075,
          "ath_date": "2019-06-30T19:42:24Z",
          "percent_from_price_ath": -61.05
        }
      }
    },
    "block_last_hash": "000000000000fbde6f4b7b16b6e2993dd278ab4c0deb8dd2cfd0dce984a179f7",
    "status": {
      "info": {
        "version": 60011,
        "protocolversion": 170008,
        "blocks": 588382,
        "timeoffset": 0,
        "connections": 21,
        "proxy": "",
        "difficulty": 1706297021564.92,
        "testnet": false,
        "relayfee": 1e-06,
        "errors": "",
        "notarized": 0,
        "network": "livenet"
      }
    },
    "node_is_online": true,
    "node_is_synced": true,
    "notarizedhash": "027e3758c3a65b12aa1046462b486d0a63bfa1beae327897f56c5cfb7daaae71",
    "notarizedtxid": [
      "4a5e1e4baab89f3a32518a88c31bc87f618f76673e2cc77ab2127b7afdeda33b"
    ]
  }
]
```

Curl command: `curl http://127.0.0.1:8080/api/v1/tickers`

## get information about one specific coin

Get all the information about a coin from the **komodo ecosystem**: ticker, last block, status, sync, notarized hash

**URL** : `/api/v1/tickers/:coin`

**Method** : `GET`

**Auth required** : No

**Permissions required** : None

### Success Response

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

### Error Response

For a coin with ID `nonexistent` that exists in the **komodo** ecosystem.

**Code** : `404 Not found`

**Content examples**

```json
{
 "error": "This coin does not seem to be part of the komodo ecosystem"
}
```

Curl command: `curl http://127.0.0.1:8080/api/v1/tickers/nonexistent`
