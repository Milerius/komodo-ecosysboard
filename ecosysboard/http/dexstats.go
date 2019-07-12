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
	"encoding/json"
	"fmt"
	"github.com/kpango/glg"
	"github.com/milerius/komodo-ecosysboard/ecosysboard/config"
	"github.com/milerius/komodo-ecosysboard/ecosysboard/utils"
	"github.com/valyala/fasthttp"
	"net/http"
)

type SearchRequestDexstatsJson struct {
	Input string `json:"input"`
}

type BlockHashFromBlockHeightJson struct {
	BlockHash string `json:"blockHash"`
}

type SearchAnswerDexstatsJson struct {
	URLToRedirect string `json:"url_to_redirect"`
}

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

func searchAsABlockHeight(endpoint string, searchInput SearchRequestDexstatsJson, ctx *fasthttp.RequestCtx, answerSearch *SearchAnswerDexstatsJson, coinName interface{}) bool {
	endpoint = endpoint + "/block-index/" + searchInput.Input
	statusCode, body, _ := fasthttp.Get(nil, endpoint)
	if statusCode != 200 {
		ctx.SetStatusCode(statusCode)
		return false
	}
	answer := BlockHashFromBlockHeightJson{}
	_ = json.Unmarshal(body, &answer)
	answerSearch.URLToRedirect = "http://" + coinName.(string) + ".explorer.dexstats.info/block/" + answer.BlockHash
	return true
}

func searchAsABlockOrAsATransaction(endpoint string, searchInput SearchRequestDexstatsJson, answerSearch *SearchAnswerDexstatsJson, coinName interface{}, ctx *fasthttp.RequestCtx) bool {
	finalEndpoint := endpoint + "/tx/" + searchInput.Input
	statusCode, _, _ := fasthttp.Get(nil, finalEndpoint)
	if statusCode == 200 {
		answerSearch.URLToRedirect = "http://" + coinName.(string) + ".explorer.dexstats.info/tx/" + searchInput.Input
	} else {
		//! It's may be a blockhash
		finalEndpoint = endpoint + "/block/" + searchInput.Input
		statusCode, _, _ = fasthttp.Get(nil, finalEndpoint)
		if statusCode == 200 {
			answerSearch.URLToRedirect = "http://" + coinName.(string) + ".explorer.dexstats.info/block/" + searchInput.Input
		} else {
			ctx.SetStatusCode(statusCode)
			return false
		}
	}
	return true
}

func SearchOnDexstats(ctx *fasthttp.RequestCtx) {
	coinName := ctx.UserValue("coin")
	searchInput := SearchRequestDexstatsJson{}
	err := json.Unmarshal(ctx.PostBody(), &searchInput)
	if err != nil {
		_ = glg.Errorf("cannot unmarshal the post body: %v", err)
		ctx.SetStatusCode(http.StatusBadRequest)
		return
	}
	_ = glg.Debugf("receive: %v", searchInput)
	answerSearch := SearchAnswerDexstatsJson{}
	endpoint := "http://localhost:" + fmt.Sprintf("%d", config.GConfig.HTTPPort) + "/api/v1/dexstats/" + coinName.(string)
	if utils.IsLookLikeABlock(searchInput.Input) {
		if !searchAsABlockHeight(endpoint, searchInput, ctx, &answerSearch, coinName) {
			return
		}
	} else if utils.IsLookLikeAKomodoAddress(searchInput.Input) {
		//! We can try before to get address details information let me know in the review
		answerSearch.URLToRedirect = "http://" + coinName.(string) + ".explorer.dexstats.info/address/" + searchInput.Input
	} else if utils.IsLookLikeABlockHashOrTransactionId(searchInput.Input) {
		if !searchAsABlockOrAsATransaction(endpoint, searchInput, &answerSearch, coinName, ctx) {
			return
		}
	} else {
		ctx.SetStatusCode(http.StatusBadRequest)
		return
	}
	_ = glg.Debugf("answerURLInformation: %v", answerSearch)
	bodyContents, _ := json.Marshal(answerSearch)
	ctx.SetContentType("application/json")
	ctx.SetStatusCode(200)
	ctx.SetBodyString(string(bodyContents))
}
