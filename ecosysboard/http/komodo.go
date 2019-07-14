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
	"github.com/kpango/glg"
	"github.com/milerius/komodo-ecosysboard/ecosysboard/config"
	"github.com/valyala/fasthttp"
	"net/http"
	"strings"
	"sync"
)

type CoinInfos struct {
	Ticker        CoinpaprikaTickerData `json:"ticker"`
	BlockLastHash string                `json:"block_last_hash"`
	BlockInfo     StatusInfo            `json:"status"`
}

func AllInformationsKomodoEcosystem(ctx *fasthttp.RequestCtx) {
	coinInfos := make([]CoinInfos, 0, len(config.GConfig.Coins))
	mutex := sync.RWMutex{}
	var wg sync.WaitGroup
	wg.Add(len(config.GConfig.Coins))
	for _, value := range config.GConfig.Coins {
		go func(key string, value string) {
			currentCoin := CoinInfos{}
			defer wg.Done()

			//! Ticker
			res := CTickerCoinpaprika(value)
			if value == "test coin" || res.Symbol == "" {
				res.Symbol = strings.ToUpper(key)
			}

			//! Last block hash
			currentCoin.BlockLastHash = CDiagnosticInfoFromNodeDexstats("getLastBlockHash", key).LastBlockHash.Lastblockhash
			currentCoin.BlockInfo = CDiagnosticInfoFromNodeDexstats("getInfo", key).Infos
			currentCoin.Ticker = *res
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
