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
	"github.com/valyala/fasthttp"
	"net/http"
	"strings"
	"sync"
)

type CoinInfos struct {
	Ticker        CoinpaprikaTickerData `json:"ticker"`
	BlockLastHash string                `json:"block_last_hash"`
	BlockHeight   int                   `json:"block_height"`
}

var komodoCoinsToCoinpaprikaRegistry = map[string]string{
	"kmd":      "kmd-komodo",
	"vrsc":     "vrsc-verus-coin",
	"k64":      "k64-komodore64",
	"rick":     "test coin",
	"revs":     "revs",
	"supernet": "unity-supernet",
}

func TickersKomodoEcosystem(ctx *fasthttp.RequestCtx) {
	tickers := make([]CoinpaprikaTickerData, 0, len(komodoCoinsToCoinpaprikaRegistry))
	mutex := sync.RWMutex{}
	var wg sync.WaitGroup
	wg.Add(len(komodoCoinsToCoinpaprikaRegistry))
	for key, value := range komodoCoinsToCoinpaprikaRegistry {
		go func(key string, value string) {
			defer wg.Done()
			res := CTickerCoinpaprika(value)
			if value == "test coin" || res.Symbol == "" {
				res.Symbol = strings.ToUpper(key)
			}
			mutex.Lock()
			tickers = append(tickers, *res)
			mutex.Unlock()
		}(key, value)
	}
	wg.Wait()
	if len(tickers) == 0 {
		ctx.SetStatusCode(http.StatusBadRequest)
		return
	}
	_ = glg.Debug("tickers komodo: %v", tickers)
	ctx.SetStatusCode(200)
	jsonTicker, _ := json.Marshal(tickers)
	ctx.SetBodyString(string(jsonTicker))
}
