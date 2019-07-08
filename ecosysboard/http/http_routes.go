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
	"github.com/valyala/fasthttp"
)

const (
	CoingGeckoEndpoint = "https://api.coingecko.com/api/v3/"
)

func Index(ctx *fasthttp.RequestCtx) {
	_, _ = ctx.WriteString("Welcome!")
	ctx.SetStatusCode(200)
}

func InitRooter() *router.Router {
	r := router.New()
	r.GET("/", Index)
	r.GET("/api/v1/coingecko/ping", PingCoingecko)
	return r
}
