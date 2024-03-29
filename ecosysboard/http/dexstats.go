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
	"github.com/KomodoPlatform/komodo-ecosysboard/ecosysboard/config"
	"github.com/KomodoPlatform/komodo-ecosysboard/ecosysboard/utils"
	"github.com/kpango/glg"
	"github.com/valyala/fasthttp"
	"net/http"
	"strconv"
)

type BlockDetails struct {
	Hash              string   `json:"hash"`
	Size              int      `json:"size"`
	Height            int      `json:"height"`
	Version           int      `json:"version"`
	Merkleroot        string   `json:"merkleroot"`
	Tx                []string `json:"tx"`
	Time              int      `json:"time"`
	Nonce             string   `json:"nonce"`
	Solution          string   `json:"solution"`
	Bits              string   `json:"bits"`
	Difficulty        float64  `json:"difficulty"`
	Chainwork         string   `json:"chainwork"`
	Confirmations     int      `json:"confirmations"`
	Previousblockhash string   `json:"previousblockhash"`
	Nextblockhash     string   `json:"nextblockhash"`
	Reward            float64  `json:"reward"`
	IsMainChain       bool     `json:"isMainChain"`
	PoolInfo          struct {
		PoolName string `json:"poolName"`
		URL      string `json:"url"`
	} `json:"poolInfo"`
}

type StatusLastBlockHash struct {
	SyncTipHash   string `json:"syncTipHash"`
	Lastblockhash string `json:"lastblockhash"`
}

type StatusBestBlockHash struct {
	Bestblockhash string `json:"bestblockhash"`
}

type StatusDifficulty struct {
	Difficulty float64 `json:"difficulty"`
}

type StatusInfo struct {
	Info struct {
		Version         int     `json:"version"`
		Protocolversion int     `json:"protocolversion"`
		Blocks          int     `json:"blocks"`
		Timeoffset      int     `json:"timeoffset"`
		Connections     int     `json:"connections"`
		Proxy           string  `json:"proxy"`
		Difficulty      float64 `json:"difficulty"`
		Testnet         bool    `json:"testnet"`
		Relayfee        float64 `json:"relayfee"`
		Errors          string  `json:"errors"`
		Notarized       int     `json:"notarized"`
		Network         string  `json:"network"`
	} `json:"info"`
}

type NodeSync struct {
	Status           string      `json:"status"`
	BlockChainHeight int         `json:"blockChainHeight"`
	SyncPercentage   int         `json:"syncPercentage"`
	Height           int         `json:"height"`
	Error            interface{} `json:"error"`
	Type             string      `json:"type"`
}

type StatusGlobal struct {
	BestBlockHash StatusBestBlockHash
	Difficulty    StatusDifficulty
	LastBlockHash StatusLastBlockHash
	Infos         StatusInfo
}

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
	InternalExecGet(fullEndpoint, ctx, true)
}

func UTXODetailsDexstats(ctx *fasthttp.RequestCtx) {
	coinName := ctx.UserValue("coin")
	address := ctx.UserValue("address")
	fullEndpoint := "http://" + coinName.(string) + DexStatsExplorerEndpoint + "/addrs/" + address.(string) + "/utxo"
	InternalExecGet(fullEndpoint, ctx, true)
}

func AddressDetailsDexstats(ctx *fasthttp.RequestCtx) {
	coinName := ctx.UserValue("coin")
	addrValue := ctx.UserValue("addrstr")
	fullEndpoint := "http://" + coinName.(string) + DexStatsExplorerEndpoint + "/addr/" + addrValue.(string)
	InternalExecGet(fullEndpoint, ctx, true)
}

func TransactionByBlockDexstats(ctx *fasthttp.RequestCtx) {
	coinName := ctx.UserValue("coin")
	hashName := ctx.UserValue("hash")
	fullEndpoint := "http://" + coinName.(string) + DexStatsExplorerEndpoint + "/txs?block=" + hashName.(string)
	InternalExecGet(fullEndpoint, ctx, true)
}

func TransactionByAddressDexstats(ctx *fasthttp.RequestCtx) {
	coinName := ctx.UserValue("coin")
	addressName := ctx.UserValue("address")
	fullEndpoint := "http://" + coinName.(string) + DexStatsExplorerEndpoint + "/txs?address=" + addressName.(string)
	InternalExecGet(fullEndpoint, ctx, true)
}

func CBlockHashFromHeightDexstats(coinName string, blockHeight string) BlockHashFromBlockHeightJson {
	blockHash := BlockHashFromBlockHeightJson{}
	fullEndpoint := "http://" + coinName + DexStatsExplorerEndpoint + "/block-index/" + blockHeight
	req, res := InternalExecGet(fullEndpoint, nil, false)
	_ = json.Unmarshal(res.Body(), &blockHash)
	ReleaseInternalExecGet(req, res)
	return blockHash
}

func BlockHashFromHeightDexstats(ctx *fasthttp.RequestCtx) {
	coinName := ctx.UserValue("coin")
	blockHeight := ctx.UserValue("blockheight")
	fullEndpoint := "http://" + coinName.(string) + DexStatsExplorerEndpoint + "/block-index/" + blockHeight.(string)
	InternalExecGet(fullEndpoint, ctx, true)
}

func DiagnosticInfoFromNodeDexstats(ctx *fasthttp.RequestCtx) {
	coinName := ctx.UserValue("coin")
	query := ctx.UserValue("query")
	fullEndpoint := "http://" + coinName.(string) + DexStatsExplorerEndpoint + "status?q=" + query.(string)
	InternalExecGet(fullEndpoint, ctx, true)
}

func CDiagnosticInfoFromNodeDexstats(statusType string, coinName string) StatusGlobal {
	status := StatusGlobal{}
	fullEndpoint := "http://" + coinName + DexStatsExplorerEndpoint + "status?q=" + statusType
	req, res := InternalExecGet(fullEndpoint, nil, false)
	switch statusType {
	case "getInfo":
		statusInfo := StatusInfo{}
		_ = json.Unmarshal(res.Body(), &statusInfo)
		status.Infos = statusInfo
		break
	case "getLastBlockHash":
		statusLastBlockHash := StatusLastBlockHash{}
		_ = json.Unmarshal(res.Body(), &statusLastBlockHash)
		status.LastBlockHash = statusLastBlockHash
		break
	case "getBestBlockHash":
		statusBestBlockHash := StatusBestBlockHash{}
		_ = json.Unmarshal(res.Body(), &statusBestBlockHash)
		status.BestBlockHash = statusBestBlockHash
		break
	default:
	}
	ReleaseInternalExecGet(req, res)
	return status
}

func CNodeSyncStatusDexstats(coinName string) NodeSync {
	nodeRes := NodeSync{}
	fullEndpoint := "http://" + coinName + DexStatsExplorerEndpoint + "/sync"
	req, res := InternalExecGet(fullEndpoint, nil, false)
	_ = json.Unmarshal(res.Body(), &nodeRes)
	ReleaseInternalExecGet(req, res)
	return nodeRes
}

func NodeSyncStatusDexstats(ctx *fasthttp.RequestCtx) {
	coinName := ctx.UserValue("coin")
	fullEndpoint := "http://" + coinName.(string) + DexStatsExplorerEndpoint + "/sync"
	InternalExecGet(fullEndpoint, ctx, true)
}

func NodePeerStatusDexstats(ctx *fasthttp.RequestCtx) {
	coinName := ctx.UserValue("coin")
	fullEndpoint := "http://" + coinName.(string) + DexStatsExplorerEndpoint + "/peer"
	InternalExecGet(fullEndpoint, ctx, true)
}

func CBlockDetailsDexstats(coinName string, blockHash string) BlockDetails {
	blockDetails := BlockDetails{}
	fullEndpoint := "http://" + coinName + DexStatsExplorerEndpoint + "/block/" + blockHash
	req, res := InternalExecGet(fullEndpoint, nil, false)
	_ = json.Unmarshal(res.Body(), &blockDetails)
	ReleaseInternalExecGet(req, res)
	return blockDetails
}

func CGetSupplyDexstats(coinName string) (float64, error, int) {
	fullEndpoint := "http://" + coinName + DexStatsExplorerWithoutInsightEndpoint + "/supply"
	req, res := InternalExecGet(fullEndpoint, nil, false)
	res.StatusCode()
	supply, err := strconv.ParseFloat(string(res.Body()), 64)
	if err != nil {
		return 0, err, res.StatusCode()
	}
	ReleaseInternalExecGet(req, res)
	return supply, nil, http.StatusOK
}

func BlockDetailsDexstats(ctx *fasthttp.RequestCtx) {
	coinName := ctx.UserValue("coin")
	hashName := ctx.UserValue("hash")
	fullEndpoint := "http://" + coinName.(string) + DexStatsExplorerEndpoint + "/block/" + hashName.(string)
	InternalExecGet(fullEndpoint, ctx, true)
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
