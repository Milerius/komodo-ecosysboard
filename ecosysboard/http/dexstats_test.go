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
	"testing"
	"time"

	"github.com/milerius/komodo-ecosysboard/ecosysboard/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/valyala/fasthttp"
)

type HTTPDexstatsTestSuite struct {
	suite.Suite
	strPort string
}

func (suite *HTTPDexstatsTestSuite) SetupTest() {
	port := GetFirstOpenPort()
	cfg := &config.Config{HTTPPort: port}
	suite.strPort = fmt.Sprintf("%d", port)
	go LaunchServer(cfg)
	time.Sleep(3 * time.Second)
}

func (suite *HTTPDexstatsTestSuite) TestAddressDetailsDexstats() {
	statusCode, body, err := fasthttp.Get(nil, "http://127.0.0.1:"+suite.strPort+"/api/v1/dexstats/kmd/addr/RSp8vhyL6hN3yqn5V1qje62pBgBE9fv3Eh")
	suite.finalizeTest(err, statusCode, body)
}

func (suite *HTTPDexstatsTestSuite) TestGetTransactionDetailsDexstats() {
	statusCode, body, err := fasthttp.Get(nil, "http://127.0.0.1:"+suite.strPort+"/api/v1/dexstats/kmd/tx/11ef4a504b4b5573bf9311c9f84e263f5535ec8a671e79d746769bda4b83fcb1")
	suite.finalizeTest(err, statusCode, body)
}

func (suite *HTTPDexstatsTestSuite) TestUTXODetailsDexstats() {
	statusCode, body, err := fasthttp.Get(nil, "http://127.0.0.1:"+suite.strPort+"/api/v1/dexstats/kmd/addrs/RSXGTHQSqwcMw1vowKfEE7sQ8fAmv1tmso/utxo")
	suite.finalizeTest(err, statusCode, body)
}

func (suite *HTTPDexstatsTestSuite) TestTransactionByBlockDexstats() {
	statusCode, body, err := fasthttp.Get(nil, "http://127.0.0.1:"+suite.strPort+"/api/v1/dexstats/kmd/txsblock/0b8c2a49d20fa270324c494e9019b11cfdc81eb90611072c61d1934bcda4de41")
	suite.finalizeTest(err, statusCode, body)
}

func (suite *HTTPDexstatsTestSuite) TestBlockDetailsDexstats() {
	statusCode, body, err := fasthttp.Get(nil, "http://127.0.0.1:"+suite.strPort+"/api/v1/dexstats/kmd/block/0b8c2a49d20fa270324c494e9019b11cfdc81eb90611072c61d1934bcda4de41")
	suite.finalizeTest(err, statusCode, body)
}

func (suite *HTTPDexstatsTestSuite) TestBlockHashFromHeightDexstats() {
	statusCode, body, err := fasthttp.Get(nil, "http://127.0.0.1:"+suite.strPort+"/api/v1/dexstats/kmd/block-index/1436987")
	suite.finalizeTest(err, statusCode, body)
}

func (suite *HTTPDexstatsTestSuite) TestNodeSyncStatusDexstats() {
	statusCode, body, err := fasthttp.Get(nil, "http://127.0.0.1:"+suite.strPort+"/api/v1/dexstats/kmd/sync")
	suite.finalizeTest(err, statusCode, body)
}

func (suite *HTTPDexstatsTestSuite) TestNodePeerStatusDexstats() {
	statusCode, body, err := fasthttp.Get(nil, "http://127.0.0.1:"+suite.strPort+"/api/v1/dexstats/kmd/peer")
	suite.finalizeTest(err, statusCode, body)
}

func (suite *HTTPDexstatsTestSuite) TestDiagnosticInfoFromNodeDexstats() {
	statusCode, body, err := fasthttp.Get(nil, "http://127.0.0.1:"+suite.strPort+"/api/v1/dexstats/kmd/status/getInfo")
	suite.finalizeTest(err, statusCode, body)
	statusCode, body, err = fasthttp.Get(nil, "http://127.0.0.1:"+suite.strPort+"/api/v1/dexstats/kmd/status/getDifficulty")
	suite.finalizeTest(err, statusCode, body)
	statusCode, body, err = fasthttp.Get(nil, "http://127.0.0.1:"+suite.strPort+"/api/v1/dexstats/kmd/status/getBestBlockHash")
	suite.finalizeTest(err, statusCode, body)
	statusCode, body, err = fasthttp.Get(nil, "http://127.0.0.1:"+suite.strPort+"/api/v1/dexstats/kmd/status/getLastBlockHash")
	suite.finalizeTest(err, statusCode, body)
}

func (suite *HTTPDexstatsTestSuite) TestTransactionByAddressDexstats() {
	statusCode, body, err := fasthttp.Get(nil, "http://127.0.0.1:"+suite.strPort+"/api/v1/dexstats/kmd/txsaddress/RMbNsa4Nf3BAd16BQaAAmfzAgnuorUDrCr")
	suite.finalizeTest(err, statusCode, body)
}

func (suite *HTTPDexstatsTestSuite) finalizeTest(err error, statusCode int, body []byte) {
	if err != nil {
		suite.T().Logf("err: %v", err)
	}
	assert.EqualValuesf(suite.T(), 200, statusCode, "status code should be 200")
	assert.Nilf(suite.T(), err, "err should be nil")
	assert.NotEmptyf(suite.T(), body, "body should not be empty")
}

func TestHTTPDexstatsTestSuite(t *testing.T) {
	suite.Run(t, new(HTTPDexstatsTestSuite))
}
