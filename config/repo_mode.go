package config

import (
	"fmt"
	"log"
	"os"
	"vx/pkg/paths"

	"github.com/spf13/viper"
)

type RepoMode struct {
	repo       string `mapstructure:"repo"`
	openai_key string `mapstructure:"openai_key"`
	github_key string `mapstructure:"github_key"`
}

var (
	repoMode = viper.New()
	vxCfg    = ".vx"
)

func configExists() bool {
	curDir, _ := os.Getwd()
	cfgDir := fmt.Sprintf("%s/%s", curDir, vxCfg)
	return paths.PathExists(cfgDir)
}

func InitRepoMode() {
	if configExists() {
		fmt.Println("Repository config already exists.")
		return
	}
	curDir, _ := os.Getwd()
	cfgFile := fmt.Sprintf("%s/%s", curDir, vxCfg)
	repoMode.SetConfigFile(cfgFile)
	repoMode.SetConfigType("json")
	dfltCfg := RepoMode{
		repo:       "",
		openai_key: "",
		github_key: "",
	}
	repoMode.Set("repo", dfltCfg.repo)
	repoMode.Set("openai_key", dfltCfg.openai_key)
	repoMode.Set("github_key", dfltCfg.github_key)
	err := repoMode.WriteConfig()
	if err != nil {
		log.Println("error initiating repo config")
	}
	addToGitIgnore(fmt.Sprintf("%s/.gitignore", curDir))
}

func addToGitIgnore(dir string) {
	paths.AppendToFile(dir, "# vx config - vexal.io \n")
	paths.AppendToFile(dir, ".vx \n")
}
