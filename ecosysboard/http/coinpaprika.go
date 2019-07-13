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
	"time"
)

type CoinpaprikaTickerData struct {
	ID                string    `json:"id"`
	Name              string    `json:"name"`
	Symbol            string    `json:"symbol"`
	Rank              int       `json:"rank"`
	CirculatingSupply int       `json:"circulating_supply"`
	TotalSupply       int       `json:"total_supply"`
	MaxSupply         int       `json:"max_supply"`
	BetaValue         float64   `json:"beta_value"`
	LastUpdated       time.Time `json:"last_updated"`
	Quotes            struct {
		USD struct {
			Price               float64   `json:"price"`
			Volume24H           float64   `json:"volume_24h"`
			Volume24HChange24H  float64   `json:"volume_24h_change_24h"`
			MarketCap           int64     `json:"market_cap"`
			MarketCapChange24H  float64   `json:"market_cap_change_24h"`
			PercentChange1H     float64   `json:"percent_change_1h"`
			PercentChange12H    float64   `json:"percent_change_12h"`
			PercentChange24H    float64   `json:"percent_change_24h"`
			PercentChange7D     float64   `json:"percent_change_7d"`
			PercentChange30D    float64   `json:"percent_change_30d"`
			PercentChange1Y     float64   `json:"percent_change_1y"`
			AthPrice            float64   `json:"ath_price"`
			AthDate             time.Time `json:"ath_date"`
			PercentFromPriceAth float64   `json:"percent_from_price_ath"`
		} `json:"USD"`
	} `json:"quotes"`
}

func CTickersCoinpaprika() []CoinpaprikaTickerData {
	var tickers []CoinpaprikaTickerData
	finalEndpoint := CoinpaprikaEndpoint + "tickers"
	req, res := InternalExecGet(finalEndpoint, nil, false)
	if res.StatusCode() == 200 {
		_ = json.Unmarshal(res.Body(), &tickers)
		_ = glg.Debugf("tickers: %v", tickers)
	}

	ReleaseInternalExecGet(req, res)
	return tickers
}

func CTickerCoinpaprika(coinsId string) *CoinpaprikaTickerData {
	ticker := &CoinpaprikaTickerData{}
	finalEndpoint := CoinpaprikaEndpoint + "tickers/" + coinsId
	req, res := InternalExecGet(finalEndpoint, nil, false)
	if res.StatusCode() == 200 {
		_ = json.Unmarshal(res.Body(), &ticker)
		_ = glg.Debugf("ticker: %v", *ticker)
	}
	ReleaseInternalExecGet(req, res)
	return ticker
}

func TickersCoinpaprika(ctx *fasthttp.RequestCtx) {
	finalEndpoint := CoinpaprikaEndpoint + "tickers"
	InternalExecGet(finalEndpoint, ctx, true)
}
