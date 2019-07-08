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

package log

import (
	"github.com/kpango/glg"
	"os"
)

func InitLogger(logsPath string) (*os.File, *os.File) {
	mode := int(0777)
	infolog := glg.FileWriter(logsPath+"/komodo_ecosysboard.info.log", os.FileMode(mode))
	errlog := glg.FileWriter(logsPath+"/komodo_ecosysboard.error.log", os.FileMode(mode))
	glg.Get().SetMode(glg.BOTH).AddLevelWriter(glg.INFO, infolog).AddLevelWriter(glg.ERR, errlog).AddLevelWriter(glg.LOG, infolog).AddLevelWriter(glg.FATAL, errlog)
	return infolog, errlog
}
