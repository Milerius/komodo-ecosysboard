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
	"io/ioutil"
	"os"
)

type Config struct {
	HTTPPort int `json:"http_port"`
}

func LoadConfig(ConfigPath string) (*Config, error) {
	if _, err := os.Stat(ConfigPath); os.IsNotExist(err) {
		return nil, errors.New("configuration in the specified path doesn't exist")
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
	return cfg, nil
}