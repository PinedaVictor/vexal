package keys

import (
	"log"
	"vx/config"
)

func KeySrc() (config.RepoMode, config.AuthConfig) {
	repoCfg, repoErr := config.LoadRepoConfig()
	if repoErr != nil {
		log.Println("Repo config does not exists")
	}
	user, usrErr := config.LoadAuth()
	if usrErr != nil {
		log.Println("Auth user invalid")
	}
	return repoCfg, user
}
