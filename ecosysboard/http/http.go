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
	"fmt"
	"github.com/KomodoPlatform/komodo-ecosysboard/ecosysboard/config"
	"github.com/kpango/glg"
	"github.com/valyala/fasthttp"
)

func LaunchServer(cfg *config.Config) {
	router := InitRooter()
	glg.Fatal(fasthttp.ListenAndServe(":"+fmt.Sprintf("%d", cfg.HTTPPort), router.Handler))
}
