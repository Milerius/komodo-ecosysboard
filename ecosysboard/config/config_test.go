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
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadConfig(t *testing.T) {
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
		{"Good path", args{"samples/good_config.json"}, &Config{8080}, false},
		{"Not enough writes", args{"samples/not_good_rights_config.json"}, nil, true},
		{"Not real json", args{"samples/not_real_json.json"}, nil, true},
		{"Non Complete path", args{"config/config"}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LoadConfig(tt.args.ConfigPath)
			assert.Equalf(t, err != nil, tt.wantErr, "LoadConfig() error = %v, wantErr %v", err, tt.wantErr)
			assert.EqualValuesf(t, got, tt.want, "LoadConfig() = %v, want %v", got, tt.want)
		})
	}
}
