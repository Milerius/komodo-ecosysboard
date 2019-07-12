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
	"net/http"
	"strconv"
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
	time.Sleep(10 * time.Millisecond)
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

func (suite *HTTPDexstatsTestSuite) TestSearchOnDexstats() {
	//fasthttp.Pos
	config.GConfig = new(config.Config)
	value, _ := strconv.Atoi(suite.strPort)
	config.GConfig.HTTPPort = value
	client := fasthttp.Client{}
	type args struct {
		Input string
	}
	tests := []struct {
		name            string
		args            args
		expected_status int
		body_contains   string
	}{
		{"Fake input", args{`bad input`}, http.StatusBadRequest, ""},
		{"Another fake input", args{`{"input": "fake input"}`}, http.StatusBadRequest, ""},
		{"Search block by valid block height", args{`{"input": "1439517"}`}, http.StatusOK, "block"},
		{"Search block by invalid block height", args{`{"input": "2439517"}`}, http.StatusBadRequest, ""},
		{"Search valid block by block hash", args{`{"input": "06d6747a49097830574cf8d33e399d8a8679e457493cd17390a80d0f916287bc"}`}, http.StatusOK, "block"},
		{"Search non existent block by block hash", args{`{"input": "0c64df81b9a66ca7882fd2b7968b562aa5394789aa0ab8444fb159862514b336"}`}, http.StatusNotFound, ""},
		{"Search valid address", args{`{"input": "RKdXYhrQxB3LtwGpysGenKFHFTqSi5g7EF"}`}, http.StatusOK, "address"},
		{"Search valid transaction", args{`{"input": "8b9478ddd6c3cae81fce0db9bb25fadece76e403eb5470c0515be99139b52042"}`}, http.StatusOK, "tx"},
		{"Search non existent transaction", args{`{"input": "8b9478ddd6c3cae81fce0db9bb25fadece76e403eb5470c0515be99139b52043"}`}, http.StatusNotFound, ""},
	}
	for _, tt := range tests {
		req := fasthttp.AcquireRequest()
		req.URI().Update("http://127.0.0.1:" + suite.strPort + "/api/v1/dexstats/kmd/search")
		req.Header.Add("Content-Type", "application/json")
		req.Header.SetMethod("POST")
		req.SetBodyString(tt.args.Input)
		res := fasthttp.AcquireResponse()
		_ = client.Do(req, res)
		suite.T().Logf("resp: %s", string(res.Body()))
		assert.EqualValues(suite.T(), tt.expected_status, res.StatusCode())
		if len(tt.body_contains) > 0 {
			assert.Containsf(suite.T(), string(res.Body()), tt.body_contains, "should contains: %s", tt.body_contains)
		}
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(res)
	}
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
