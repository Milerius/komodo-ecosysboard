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
	"github.com/KomodoPlatform/komodo-ecosysboard/ecosysboard/config"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

func TestCGetGithubOrgsRepos(t *testing.T) {
	dir, _ := os.Getwd()
	parent := filepath.Dir(dir)
	_, err := config.LoadConfig(parent + "/config/samples/good_config.json")
	assert.Nil(t, err, "should be nil ")
	res := CGetGithubOrgsRepos("KomodoPlatform")
	assert.False(t, cmp.Equal(res, GithubOrgRepos{}), "should not be true")
	assert.NotZero(t, len(res), "Should not be empty")
}
