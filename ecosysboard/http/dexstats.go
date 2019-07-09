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
	"github.com/valyala/fasthttp"
)

func GetTransactionDetailsDexstats(ctx *fasthttp.RequestCtx) {
	coinName := ctx.UserValue("coin")
	txId := ctx.UserValue("txid")
	fullEndpoint := "http://" + coinName.(string) + DexStatsExplorerEndpoint + "/tx/" + txId.(string)
	InternalExecGet(fullEndpoint, ctx)
}

func UTXODetailsDexstats(ctx *fasthttp.RequestCtx) {
	coinName := ctx.UserValue("coin")
	address := ctx.UserValue("address")
	fullEndpoint := "http://" + coinName.(string) + DexStatsExplorerEndpoint + "/addrs/" + address.(string) + "/utxo"
	InternalExecGet(fullEndpoint, ctx)
}

func AddressDetailsDexstats(ctx *fasthttp.RequestCtx) {
	coinName := ctx.UserValue("coin")
	addrValue := ctx.UserValue("addrstr")
	fullEndpoint := "http://" + coinName.(string) + DexStatsExplorerEndpoint + "/addr/" + addrValue.(string)
	InternalExecGet(fullEndpoint, ctx)
}
