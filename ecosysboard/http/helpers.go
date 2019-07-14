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

func InternalExecGet(finalEndpoint string, ctx *fasthttp.RequestCtx, shouldRelease bool) (*fasthttp.Request, *fasthttp.Response) {
	_ = glg.Debugf("final endpoint: %s", finalEndpoint)
	status, body, _ := fasthttp.Get(nil, finalEndpoint)
	if ctx != nil {
		ctx.SetStatusCode(status)
		ctx.SetBodyString(string(body))
	}
	_ = glg.Debugf("http response: %s", string(body))
	if !shouldRelease {
		req := fasthttp.AcquireRequest()
		res := fasthttp.AcquireResponse()
		res.SetStatusCode(status)
		res.SetBody(body)
		return req, res
	}
	return nil, nil
}

func ReleaseInternalExecGet(req *fasthttp.Request, res *fasthttp.Response) {
	fasthttp.ReleaseRequest(req)
	fasthttp.ReleaseResponse(res)
}
