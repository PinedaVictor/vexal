package config

import (
	"fmt"
	"log"
	"os"
	"vx/pkg/gutils"
	"vx/pkg/paths"

	"github.com/spf13/viper"
)

type RepoMode struct {
	Repo       string `mapstructure:"repo"`
	RepoURL    string `mapstructure:"repoURL"`
	Openai_key string `mapstructure:"openai_key"`
	Github_key string `mapstructure:"github_key"`
}

var (
	repoMode = viper.New()
	vxCfg    = ".vx"
)

func RepoModeActive() bool {
	curDir, _ := os.Getwd()
	cfgDir := fmt.Sprintf("%s/%s", curDir, vxCfg)
	return paths.PathExists(cfgDir)
}

func InitRepoMode() {
	if RepoModeActive() {
		fmt.Println("Repository config already exists.")
		return
	}
	curDir, _ := os.Getwd()
	cfgFile := fmt.Sprintf("%s/%s", curDir, vxCfg)
	repoMode.SetConfigFile(cfgFile)
	repoMode.SetConfigType("yaml")
	_, repoName, gitRepoURL := gutils.GetRepo()

	dfltCfg := RepoMode{
		Repo:       repoName,
		RepoURL:    gitRepoURL,
		Openai_key: "",
		Github_key: "",
	}
	repoMode.Set("repo", dfltCfg.Repo)
	repoMode.Set("repoURL", dfltCfg.RepoURL)
	repoMode.Set("openai_key", dfltCfg.Openai_key)
	repoMode.Set("github_key", dfltCfg.Github_key)
	err := repoMode.WriteConfig()
	if err != nil {
		log.Println("error initiating repo config")
	}

	addToGitIgnore(fmt.Sprintf("%s/.gitignore", curDir))
}

func addToGitIgnore(dir string) {
	paths.AppendToFile(dir, "# vexal.io vx config \n")
	paths.AppendToFile(dir, ".vx \n")
}

func LoadRepoConfig() (config RepoMode, err error) {
	curDir, _ := os.Getwd()
	repoMode.AddConfigPath(curDir)
	repoMode.SetConfigName(vxCfg) // name of config file (without extension)
	repoMode.SetConfigType("yaml")
	repoMode.AutomaticEnv()

	err = repoMode.ReadInConfig()
	if err != nil {
		log.Println("Error reading repo config:", err)
		return
	}

	err = repoMode.Unmarshal(&config)
	if err != nil {
		log.Printf("Unable to decode into struct: %s\n", err)
		return config, err
	}
	return config, nil
}
