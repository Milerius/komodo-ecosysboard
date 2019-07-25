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

package config

import (
	"encoding/json"
	"errors"
	"github.com/google/go-cmp/cmp"
	"io/ioutil"
	"os"
	"sort"
)

type Config struct {
	HTTPPort       int    `json:"http_port"`
	GitBotClientId string `json:"git_bot_client_id"`
	GitBotSecretId string `json:"git_bot_secret_id"`
	Coins          []struct {
		Coin          string   `json:"coin"`
		CoinPaprikaID string   `json:"coin_paprika_id"`
		GitReposList  []string `json:"git"`
	} `json:"coins"`
	GitReposDirectory string `json:"git_repos_location"`
}

var GConfig *Config

func LoadConfig(ConfigPath string) (*Config, error) {
	if _, err := os.Stat(ConfigPath); os.IsNotExist(err) {
		return nil, errors.New("configuration in the specified path doesn't exist: ")
	}
	file, err := ioutil.ReadFile(ConfigPath)
	if err != nil {
		return nil, err
	}

	cfg := &Config{}
	err = json.Unmarshal([]byte(file), cfg)

	if err != nil {
		return nil, err
	}

	if cmp.Equal(Config{}, *cfg) {
		return nil, errors.New("configuration seem's to be wrong, or not well-formed")
	}
	sort.SliceStable(cfg.Coins, func(i, j int) bool {
		return cfg.Coins[i].Coin < cfg.Coins[j].Coin
	})
	GConfig = cfg
	return cfg, nil
}
