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
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
	"testing"
)

func TestPingCoingecko(t *testing.T) {
	port := GetFirstOpenPort()
	cfg := &config.Config{HTTPPort: port}
	strPort := fmt.Sprintf("%d", port)
	go LaunchServer(cfg)
	statusCode, body, err := fasthttp.Get(nil, "http://127.0.0.1:"+strPort+"/api/v1/coingecko/ping")
	assert.EqualValuesf(t, 200, statusCode, "status code should be 200")
	assert.Nilf(t, err, "err should be nil")
	assert.NotEmptyf(t, body, "body should not be empty")
}
