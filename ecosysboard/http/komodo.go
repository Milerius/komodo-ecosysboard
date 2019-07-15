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
	"github.com/google/go-cmp/cmp"
	"github.com/kpango/glg"
	"github.com/milerius/komodo-ecosysboard/ecosysboard/config"
	"github.com/valyala/fasthttp"
	"net/http"
	"sort"
	"strings"
	"sync"
)

type CoinInfos struct {
	Ticker                CoinpaprikaTickerData `json:"ticker"`
	BlockLastHash         string                `json:"block_last_hash"`
	BlockInfo             StatusInfo            `json:"status"`
	NodeIsOnline          bool                  `json:"node_is_online"`
	NodeIsSynced          bool                  `json:"node_is_synced"`
	NotarizedHash         string                `json:"notarizedhash"`
	NotarizedTransactions []string              `json:"notarizedtxid"`
}

func getInfoAboutSpecificCoin(key string, value string) CoinInfos {
	currentCoin := CoinInfos{}
	//! Ticker
	res := CTickerCoinpaprika(value)
	if value == "test coin" || res.Symbol == "" {
		res.Symbol = strings.ToUpper(key)
	}
	//! Last block hash
	currentCoin.BlockLastHash = CDiagnosticInfoFromNodeDexstats("getLastBlockHash", key).LastBlockHash.Lastblockhash
	currentCoin.BlockInfo = CDiagnosticInfoFromNodeDexstats("getInfo", key).Infos
	node := CNodeSyncStatusDexstats(key)
	currentCoin.NodeIsSynced = node.Status == "finished" && node.BlockChainHeight == currentCoin.BlockInfo.Info.Blocks
	currentCoin.NodeIsOnline = currentCoin.BlockInfo.Info.Connections > 2
	if currentCoin.NodeIsSynced && currentCoin.NodeIsOnline {
		currentCoin.NotarizedHash = CBlockHashFromHeightDexstats(key, fmt.Sprintf("%d", currentCoin.BlockInfo.Info.Notarized)).BlockHash
		currentCoin.NotarizedTransactions = CBlockDetailsDexstats(key, currentCoin.NotarizedHash).Tx
	}
	currentCoin.Ticker = *res
	return currentCoin
}

func GetInformationForSpecificCoinKomodoEcosystem(ctx *fasthttp.RequestCtx) {
	coinName := ctx.UserValue("coin")
	idx := sort.Search(len(config.GConfig.Coins), func(i int) bool { return config.GConfig.Coins[i].Coin >= coinName.(string) })
	_ = glg.Infof("find needle: %v", config.GConfig.Coins[idx])
	coinInfo := getInfoAboutSpecificCoin(config.GConfig.Coins[idx].Coin, config.GConfig.Coins[idx].CoinPaprikaID)
	if cmp.Equal(CoinInfos{}, coinInfo) {
		ctx.SetStatusCode(http.StatusInternalServerError)
		return
	}
	ctx.SetStatusCode(200)
	coinInfoJson, _ := json.Marshal(coinInfo)
	ctx.SetBodyString(string(coinInfoJson))
}

func AllInformationsKomodoEcosystem(ctx *fasthttp.RequestCtx) {
	coinInfos := make([]CoinInfos, 0, len(config.GConfig.Coins))
	mutex := sync.RWMutex{}
	var wg sync.WaitGroup
	wg.Add(len(config.GConfig.Coins))
	for _, value := range config.GConfig.Coins {
		go func(key string, value string) {
			defer wg.Done()
			currentCoin := getInfoAboutSpecificCoin(key, value)
			mutex.Lock()
			coinInfos = append(coinInfos, currentCoin)
			mutex.Unlock()
		}(value.Coin, value.CoinPaprikaID)
	}
	wg.Wait()
	if len(coinInfos) == 0 {
		ctx.SetStatusCode(http.StatusBadRequest)
		return
	}
	_ = glg.Debug("coinInfos komodo: %v", coinInfos)
	ctx.SetStatusCode(200)
	jsonTicker, _ := json.Marshal(coinInfos)
	ctx.SetBodyString(string(jsonTicker))
}
