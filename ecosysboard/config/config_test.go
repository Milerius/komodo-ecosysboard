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
		{
			name: "Good path",
			args: args{"samples/good_config.json"},
			want: &Config{
				HTTPPort:       8080,
				GitBotClientId: "d1d63db76668da1141cd",
				GitBotSecretId: "74afa0fe4b377bdfa1d41ddc33147c0042a952b1",
				Coins: []struct {
					Coin          string   `json:"coin"`
					CoinPaprikaID string   `json:"coin_paprika_id"`
					GitReposList  []string `json:"git"`
				}{
					{"kmd", "kmd-komodo", []string{"KomodoPlatform", "https://github.com/pbca26/komodo-omni-explorer.git",
						"https://github.com/jl777/komodo.git",
						"https://github.com/atomiclabs/hyperdex.git",
						"https://github.com/jl777/chips3.git",
						"https://github.com/jl777/chipsln.git",
						"https://github.com/sg777/bet.git",
						"https://github.com/jl777/coins.git",
						"https://github.com/DeckerSU/komodo-explorers-install.git",
						"https://github.com/DeckerSU/KomodoOcean.git",
						"https://github.com/blackjok3rtt/ScaleTestV2.git",
						"https://github.com/patchkez/scaletest_containers.git",
						"https://github.com/smk762/txscl_vis.git",
						"https://github.com/Meshbits/TxBlaster.git",
						"https://github.com/webworker01/knomp.git",
						"https://github.com/pbca26/agama-mobile.git",
						"https://github.com/pbca26/agama-web.git"}},
					{"k64", "k64-komodore64", []string{"KomodoPlatform"}},
					{"vrsc", "vrsc-verus-coin", []string{"KomodoPlatform"}},
					{"rick", "test coin", []string{"KomodoPlatform"}},
					{"revs", "revs", []string{"KomodoPlatform"}},
					{"supernet", "unity-supernet", []string{"KomodoPlatform"}},
				},
				GitReposDirectory: "/tmp/repos",
			},
			wantErr: false,
		},
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
