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

package services

import (
	"os"
	"sync"
	"sync/atomic"
	"testing"

	"github.com/KomodoPlatform/komodo-ecosysboard/ecosysboard/config"
	"github.com/stretchr/testify/assert"
)

func TestSetupGitRepos(t *testing.T) {
	_, err := config.LoadConfig("../config/samples/good_config.json")
	assert.Nil(t, err, "should be nil")
	config.GConfig.Coins[2].GitReposList = []string{"NonExistentPlatform"}
	assert.NotZero(t, SetupGitRepos(), "should clone more than 0 repository")
	assert.Zero(t, SetupGitRepos(), "should clone 0 repository")
	assert.Nil(t, os.RemoveAll(config.GConfig.GitReposDirectory), "should be nil")
	config.GConfig.GitReposDirectory = "vendor"
	assert.NotZero(t, SetupGitRepos(), "should clone more than 0 repository")
}

func Test_gitClone(t *testing.T) {
	_, err := config.LoadConfig("../config/samples/good_config.json")
	assert.Nil(t, err, "should be nil")
	var nbReposCloned uint64 = 0
	var wg sync.WaitGroup
	wg.Add(3)
	go gitClone("non existent repository", &nbReposCloned, &wg)
	go gitClone("https://github.com/Milerius/nonexistentrepository", &nbReposCloned, &wg)
	go gitClone("https://github.com/Milerius/komodo-playground", &nbReposCloned, &wg)
	wg.Wait()
	assert.Equal(t, uint64(1), atomic.LoadUint64(&nbReposCloned), "should be one")
	assert.Nil(t, os.RemoveAll(config.GConfig.GitReposDirectory), "should be nil")
}

func TestFetchGitStats(t *testing.T) {
	FetchGitStats()
}

func TestPullGitRepos(t *testing.T) {
	_, err := config.LoadConfig("../config/samples/good_config.json")
	assert.Nil(t, err, "should be nil")
	config.GConfig.GitReposDirectory = "/nonexistent"
	nbRepos, err := PullGitRepos()
	assert.Zero(t, nbRepos, "should pull 0 repository")
	assert.NotNil(t, err, "should not be nill")
}
