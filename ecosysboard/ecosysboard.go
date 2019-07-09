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

package main

import (
	"flag"
	"fmt"
	"github.com/kpango/glg"
	"github.com/milerius/komodo-ecosysboard/ecosysboard/config"
	"github.com/milerius/komodo-ecosysboard/ecosysboard/http"
	"github.com/milerius/komodo-ecosysboard/ecosysboard/log"
)

func main() {
	configPath := flag.String("config_path", "", "path to the configuration")
	logsPath := flag.String("logs_path", "", "path where the logs should be stored")
	flag.Parse()
	if *configPath == "" || *logsPath == "" {
		fmt.Println("config path or log path are empty, set them through the command line")
		return
	}
	infolog, errlog := log.InitLogger(*logsPath)
	defer infolog.Close()
	defer errlog.Close()
	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		glg.Fatalf("error loading configuration: %v", err)
	}
	_ = glg.Infof("Successfully parsed config: %v", *cfg)
	http.LaunchServer(cfg)
}
