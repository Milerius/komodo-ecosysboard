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
	"bytes"
	"github.com/KomodoPlatform/komodo-ecosysboard/ecosysboard/config"
	"github.com/KomodoPlatform/komodo-ecosysboard/ecosysboard/http"
	"github.com/KomodoPlatform/komodo-ecosysboard/ecosysboard/utils"
	"github.com/kpango/glg"
	"net/url"
	"os"
	"os/exec"
	"path"
	"strings"
	"sync"
	"sync/atomic"
)

func SetupGitRepos() uint64 {
	var nbReposCloned uint64 = 0
	var wg sync.WaitGroup
	if !utils.IsPathExist(config.GConfig.GitReposDirectory) {
		_ = os.Mkdir(config.GConfig.GitReposDirectory, os.ModePerm)
	}

	if !path.IsAbs(config.GConfig.GitReposDirectory) {
		currentDirectory, _ := os.Getwd()
		config.GConfig.GitReposDirectory = currentDirectory + "/" + config.GConfig.GitReposDirectory
	}
	_ = glg.Debugf("Repos path: %s", config.GConfig.GitReposDirectory)
	for _, currentCoin := range config.GConfig.Coins {
		for _, currentRepo := range currentCoin.GitReposList {
			_, err := url.ParseRequestURI(currentRepo)
			if err != nil {
				if utils.IsPathExist(config.GConfig.GitReposDirectory + "/" + currentRepo) {
					continue
				}
				res := http.CGetGithubOrgsRepos(currentRepo)
				if len(res) == 0 {
					continue
				} else {
					for _, currentRepoOrg := range res {
						gitCloneInternal(currentRepoOrg.CloneURL, &wg, &nbReposCloned)
					}
				}
			} else {
				gitCloneInternal(currentRepo, &wg, &nbReposCloned)
			}
		}
	}
	wg.Wait()
	_, _ = PullGitRepos()
	return atomic.LoadUint64(&nbReposCloned)
}

func gitCloneInternal(currentRepo string, wg *sync.WaitGroup, nbReposCloned *uint64) {
	_ = glg.Debugf("Full path git check: %s", config.GConfig.GitReposDirectory+"/"+currentRepo)
	wg.Add(1)
	go gitClone(currentRepo, nbReposCloned, wg)
}

func gitClone(urlRepo string, nbReposCloned *uint64, wg *sync.WaitGroup) {
	ss := strings.Split(urlRepo, "/")
	s := ss[len(ss)-1]
	if len(ss) < 2 {
		wg.Done()
		return
	}
	username := ss[len(ss)-2]
	s = strings.Trim(s, ".git")
	finalUrlPath := config.GConfig.GitReposDirectory + "/" + username + "/" + s
	if !utils.IsPathExist(finalUrlPath) {
		var out bytes.Buffer
		var stderr bytes.Buffer
		cmd := exec.Command("git", "clone", urlRepo, finalUrlPath)
		cmd.Stdout = &out
		cmd.Stderr = &stderr
		err := cmd.Run()
		if err == nil {
			atomic.AddUint64(nbReposCloned, 1)
		} else {
			_ = glg.Errorf("Error occured during execution: %v", err)
			_ = glg.Errorf("Error occured during execution: %s", out.String())
			_ = glg.Errorf("Error occured during execution: %s", stderr.String())
		}
	} else {
		_ = glg.Infof("Repository %s already cloned", urlRepo)
	}
	wg.Done()
}

func PullGitRepos() (uint64, error) {
	var nbReposPulled uint64 = 0
	var wg sync.WaitGroup
	dirs, err := utils.IOReadDir(config.GConfig.GitReposDirectory, config.GConfig.GitReposDirectory)
	if err != nil || len(dirs) == 0 {
		_ = glg.Errorf("Cannot get directory from path: %s", config.GConfig.GitReposDirectory)
		return nbReposPulled, err
	}
	for _, currentUsername := range dirs {
		dirs, err = utils.IOReadDir(currentUsername, currentUsername)
		if err == nil {
			for _, currentRepo := range dirs {
				wg.Add(1)
				go gitPullInternal(currentRepo, &wg, &nbReposPulled)
			}
		}
	}
	wg.Wait()
	return atomic.LoadUint64(&nbReposPulled), nil
}

func gitPullInternal(currentRepo string, wg *sync.WaitGroup, nbReposPulled *uint64) {
	cmd := `git -C ` + currentRepo + ` branch | grep \* | cut -d ' ' -f2`
	out, err := exec.Command("sh", "-c", cmd).Output()
	_ = glg.Debugf("output cmd %s: %s", cmd, string(out))
	if err == nil {
		cmd = `git -C ` + currentRepo + ` pull origin ` + string(out)
		out, err = exec.Command("sh", "-c", cmd).Output()
		if err == nil {
			atomic.AddUint64(nbReposPulled, 1)
		}
		_ = glg.Debugf("output cmd %s: %s", cmd, string(out))
	}
	wg.Done()
}

func FetchGitStats() {

}

/*func LaunchGitStatsServices() {
	SetupGitRepos()
	FetchGitStats()
	gitStatsTicker := time.NewTicker(5 * time.Minute)
	for {
		select {
		case <-gitStatsTicker.C:
			FetchGitStats()
		}
	}
}*/
