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
	"fmt"
	"github.com/kpango/glg"
	"github.com/valyala/fasthttp"
)

func GetFirstOpenPort() int {
	port := C.get_first_open_port()
	return int(port)
}

func InternalExecGet(finalEndpoint string, ctx *fasthttp.RequestCtx) {
	status, body, err := fasthttp.Get(nil, finalEndpoint)
	if err != nil {
		_ = glg.Error(err)
		ctx.SetStatusCode(status)
		return
	}
	if status != 200 {
		_ = glg.Error("status code is not 200")
	}
	ctx.SetStatusCode(status)
	_, _ = ctx.Write(body)
	_ = glg.Info("http response: ", string(body))
	fmt.Println(len(body))
}
