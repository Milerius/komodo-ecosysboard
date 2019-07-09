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
	"github.com/milerius/komodo-ecosysboard/ecosysboard/config"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
	"testing"
	"time"
)

var port int
var strPort string

func SetupTest() {
	port = GetFirstOpenPort()
	strPort = fmt.Sprintf("%d", port)
}

func AddressDetailsDexstatsTest(t *testing.T) {
	statusCode, body, err := fasthttp.Get(nil, "http://127.0.0.1:"+strPort+"/api/v1/dexstats/addr/kmd/RSp8vhyL6hN3yqn5V1qje62pBgBE9fv3Eh")
	if err != nil {
		t.Logf("err: %v", err)
	}
	assert.EqualValuesf(t, 200, statusCode, "status code should be 200")
	assert.Nilf(t, err, "err should be nil")
	assert.NotEmptyf(t, body, "body should not be empty")
}
func GetTransactionDetailsDexstatsTest(t *testing.T) {
	statusCode, body, err := fasthttp.Get(nil, "http://127.0.0.1:"+strPort+"/api/v1/dexstats/tx/kmd/11ef4a504b4b5573bf9311c9f84e263f5535ec8a671e79d746769bda4b83fcb1")
	if err != nil {
		t.Logf("err: %v", err)
	}
	assert.EqualValuesf(t, 200, statusCode, "status code should be 200")
	assert.Nilf(t, err, "err should be nil")
	assert.NotEmptyf(t, body, "body should not be empty")
}

func UTXODetailsDexstatsTest(t *testing.T) {
	statusCode, body, err := fasthttp.Get(nil, "http://127.0.0.1:"+strPort+"/api/v1/dexstats/addrs/kmd/RSXGTHQSqwcMw1vowKfEE7sQ8fAmv1tmso/utxo")
	if err != nil {
		t.Logf("err: %v", err)
	}
	assert.EqualValuesf(t, 200, statusCode, "status code should be 200")
	assert.Nilf(t, err, "err should be nil")
	assert.NotEmptyf(t, body, "body should not be empty")
}

func TestHTTPDexstatsTestSuite(t *testing.T) {
	SetupTest()
	cfg := &config.Config{HTTPPort: port}
	go LaunchServer(cfg)
	UTXODetailsDexstatsTest(t)
	time.Sleep(1 * time.Second)
	GetTransactionDetailsDexstatsTest(t)
	time.Sleep(1 * time.Second)
	AddressDetailsDexstatsTest(t)
}
