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
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"sort"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	_ = os.Chmod("samples/not_good_rights_config.json", 0000)
	type args struct {
		ConfigPath string
	}
	tests := []struct {
		name    string
		args    args
		want    *Config
		wantErr bool
	}{
		{"Wrong path", args{"unexistent/path/"}, nil, true},
		{"Good path", args{"samples/good_config.json"}, &Config{HTTPPort: 8080, Coins: []struct {
			Coin          string `json:"coin"`
			CoinPaprikaID string `json:"coin_paprika_id"`
		}{
			{"kmd", "kmd-komodo"},
			{"k64", "k64-komodore64"},
			{"vrsc", "vrsc-verus-coin"},
			{"rick", "test coin"},
			{"revs", "revs"},
			{"supernet", "unity-supernet"},
		}}, false},
		{"Not enough writes", args{"samples/not_good_rights_config.json"}, nil, true},
		{"Not real json", args{"samples/not_real_json.json"}, nil, true},
		{"Non Complete path", args{"config/config"}, nil, true},
		{"Wrong Json file", args{"samples/wrong_json.json"}, nil, true},
	}
	for _, tt := range tests {
		if tt.want != nil {
			sort.SliceStable(tt.want.Coins, func(i, j int) bool {
				return tt.want.Coins[i].Coin < tt.want.Coins[j].Coin
			})
		}
		t.Run(tt.name, func(t *testing.T) {
			got, err := LoadConfig(tt.args.ConfigPath)
			assert.Equalf(t, err != nil, tt.wantErr, "LoadConfig() error = %v, wantErr %v", err, tt.wantErr)
			assert.EqualValuesf(t, got, tt.want, "LoadConfig() = %v, want %v", got, tt.want)
		})
	}
	err := os.Chmod("samples/not_good_rights_config.json", 0644)
	if err != nil {
		fmt.Println(err)
	}
}
