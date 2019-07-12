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

// #cgo CFLAGS: -O2 -Wall
// #include "magic_port.h"
import "C"

import (
	"github.com/kpango/glg"
	"github.com/valyala/fasthttp"
)

func GetFirstOpenPort() int {
	port := C.get_first_open_port()
	return int(port)
}

func InternalExecGet(finalEndpoint string, ctx *fasthttp.RequestCtx) {
	client := fasthttp.Client{}
	req := fasthttp.AcquireRequest()
	req.Header.SetMethod("GET")
	req.URI().Update(finalEndpoint)
	res := fasthttp.AcquireResponse()
	_ = client.Do(req, res)
	ctx.SetStatusCode(res.StatusCode())
	ctx.SetBodyString(string(res.Body()))
	_ = glg.Debugf("http response: %s", string(res.Body()))
	fasthttp.ReleaseRequest(req)
	fasthttp.ReleaseResponse(res)
}
