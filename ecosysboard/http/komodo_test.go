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
	"github.com/stretchr/testify/suite"
	"github.com/valyala/fasthttp"
	"testing"
	"time"
)

type HTTPKomodoTestSuite struct {
	suite.Suite
	strPort string
}

func (suite *HTTPKomodoTestSuite) finalizeTests(url string) {
	client := fasthttp.Client{}
	req := fasthttp.AcquireRequest()
	req.Header.SetMethod("GET")
	req.URI().Update(url)
	res := fasthttp.AcquireResponse()
	_ = client.Do(req, res)
	if len(string(res.Body())) < 500 {
		suite.T().Logf("http response: %s", string(res.Body()))
	}
	assert.EqualValuesf(suite.T(), 200, res.StatusCode(), "status code should be 200")
	assert.NotEmptyf(suite.T(), res.Body(), "body should not be empty")
	fasthttp.ReleaseRequest(req)
	fasthttp.ReleaseResponse(res)
}

func (suite *HTTPKomodoTestSuite) SetupTest() {
	port := GetFirstOpenPort()
	cfg := &config.Config{HTTPPort: port}
	suite.strPort = fmt.Sprintf("%d", port)
	go LaunchServer(cfg)
	time.Sleep(10 * time.Millisecond)
}

func TestHTTPKomodoTestSuite(t *testing.T) {
	suite.Run(t, new(HTTPKomodoTestSuite))
}

func (suite *HTTPCoinpaprikaTestSuite) TestTickersKomodoEcosystem() {
	suite.finalizeTests("http://127.0.0.1:" + suite.strPort + "/api/v1/tickers")
}
