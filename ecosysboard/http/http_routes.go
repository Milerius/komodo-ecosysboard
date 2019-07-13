/******************************************************************************
 * Copyright © 2013-2019 The Komodo Platform Developers.                      *
 *                                                                            *
 * See the AUTHORS, DEVELOPER-AGREEMENT and LICENSE files at                  *
 * the top-level directory of this distribution for the individual copyright  *
 * holder information and the developer policies on copyright and licensing.  *
 *                                                                            *
 * Unless otherwise agreed in a custom licensing agreement, no part of the    *
 * Komodo Platform software, including this file may be copied, modified,     *
 * propagated or distributed except according to the terms contained in the   *
 * LICENSE file                                                               *
 *                                                                            *
 * Removal or modification of this copyright notice is prohibited.            *
 *                                                                            *
 ******************************************************************************/

package http

import (
	"github.com/fasthttp/router"
)

const (
	CoingGeckoEndpoint       = "https://api.coingecko.com/api/v3/"
	CoinpaprikaEndpoint      = "https://api.coinpaprika.com/v1/"
	DexStatsExplorerEndpoint = ".explorer.dexstats.info/insight-api-komodo/"
)

func InitRooter() *router.Router {
	r := router.New()

	//! Coingecko
	r.GET("/api/v1/coingecko/ping", PingCoingecko)

	//! Dexstats
	r.GET("/api/v1/dexstats/:coin/addr/:addrstr", AddressDetailsDexstats)
	r.GET("/api/v1/dexstats/:coin/tx/:txid", GetTransactionDetailsDexstats)
	r.GET("/api/v1/dexstats/:coin/addrs/:address/utxo", UTXODetailsDexstats)
	r.GET("/api/v1/dexstats/:coin/txsblock/:hash", TransactionByBlockDexstats)
	r.GET("/api/v1/dexstats/:coin/txsaddress/:address", TransactionByAddressDexstats)
	r.GET("/api/v1/dexstats/:coin/block/:hash", BlockDetailsDexstats)
	r.GET("/api/v1/dexstats/:coin/block-index/:blockheight", BlockHashFromHeightDexstats)
	r.GET("/api/v1/dexstats/:coin/status/:query", DiagnosticInfoFromNodeDexstats)
	r.GET("/api/v1/dexstats/:coin/sync", NodeSyncStatusDexstats)
	r.GET("/api/v1/dexstats/:coin/peer", NodePeerStatusDexstats)
	r.POST("/api/v1/dexstats/:coin/search", SearchOnDexstats)

	//! Coinpaprika
	r.GET("/api/v1/coinpaprika/tickers", TickersCoinpaprika)

	r.GET("/api/v1/tickers", TickersKomodoEcosystem)
	return r
}
