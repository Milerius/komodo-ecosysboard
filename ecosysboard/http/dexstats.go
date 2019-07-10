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

func TransactionByBlockDexstats(ctx *fasthttp.RequestCtx) {
	coinName := ctx.UserValue("coin")
	hashName := ctx.UserValue("hash")
	fullEndpoint := "http://" + coinName.(string) + DexStatsExplorerEndpoint + "/txs?block=" + hashName.(string)
	InternalExecGet(fullEndpoint, ctx)
}

func TransactionByAddressDexstats(ctx *fasthttp.RequestCtx) {
	coinName := ctx.UserValue("coin")
	addressName := ctx.UserValue("address")
	fullEndpoint := "http://" + coinName.(string) + DexStatsExplorerEndpoint + "/txs?address=" + addressName.(string)
	InternalExecGet(fullEndpoint, ctx)
}

func BlockHashFromHeightDexstats(ctx *fasthttp.RequestCtx) {
	coinName := ctx.UserValue("coin")
	blockHeight := ctx.UserValue("blockheight")
	fullEndpoint := "http://" + coinName.(string) + DexStatsExplorerEndpoint + "/block-index/" + blockHeight.(string)
	InternalExecGet(fullEndpoint, ctx)
}

func DiagnosticInfoFromNodeDexstats(ctx *fasthttp.RequestCtx) {
	coinName := ctx.UserValue("coin")
	query := ctx.UserValue("query")
	fullEndpoint := "http://" + coinName.(string) + DexStatsExplorerEndpoint + "/status?:q=" + query.(string)
	InternalExecGet(fullEndpoint, ctx)
}

func NodeSyncStatusDexstats(ctx *fasthttp.RequestCtx) {
	coinName := ctx.UserValue("coin")
	fullEndpoint := "http://" + coinName.(string) + DexStatsExplorerEndpoint + "/sync"
	InternalExecGet(fullEndpoint, ctx)
}

func NodePeerStatusDexstats(ctx *fasthttp.RequestCtx) {
	coinName := ctx.UserValue("coin")
	fullEndpoint := "http://" + coinName.(string) + DexStatsExplorerEndpoint + "/peer"
	InternalExecGet(fullEndpoint, ctx)
}

func BlockDetailsDexstats(ctx *fasthttp.RequestCtx) {
	coinName := ctx.UserValue("coin")
	hashName := ctx.UserValue("hash")
	fullEndpoint := "http://" + coinName.(string) + DexStatsExplorerEndpoint + "/block/" + hashName.(string)
	InternalExecGet(fullEndpoint, ctx)
}
