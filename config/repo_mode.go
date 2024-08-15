package config

import (
	"fmt"
	"log"
	"os"
	"vx/pkg/gutils"
	"vx/pkg/paths"

	"github.com/spf13/viper"
)

type JiraRepoCfg struct {
	Jira_URL      string `mapstructure:"jira_url"`
	Jira_Cloud_ID string `mapstructure:"jira_cloud_id"`
	Jira_Email    string `mapstructure:"jira_email"`
}

type RepoMode struct {
	Repo         string `mapstructure:"repo"`
	RepoURL      string `mapstructure:"repo_URL"`
	Openai_key   string `mapstructure:"openai_key"`
	Github_key   string `mapstructure:"github_key"`
	JiraSettings JiraRepoCfg
}

var (
	repoMode = viper.New()
	vxCfg    = ".vx.yaml"
)

func RepoModeActive() bool {
	curDir, _ := os.Getwd()
	cfgDir := fmt.Sprintf("%s/%s", curDir, vxCfg)
	return paths.PathExists(cfgDir)
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
	repoMode.Set("repo_url", dfltCfg.RepoURL)
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
	paths.AppendToFile(dir, ".vx* \n")
}

func InitJira() error {
	LoadRepoConfig()

	repoCfg := RepoMode{
		JiraSettings: JiraRepoCfg{
			Jira_URL:      "",
			Jira_Cloud_ID: "",
			Jira_Email:    "",
		},
	}
	repoMode.Set("jira_url", repoCfg.JiraSettings.Jira_URL)
	repoMode.Set("jira_email", repoCfg.JiraSettings.Jira_Email)
	repoMode.Set("jira_cloud_id", repoCfg.JiraSettings.Jira_Cloud_ID)

	// Write the updated config back to the file
	if err := repoMode.WriteConfig(); err != nil {
		return fmt.Errorf("error writing config file: %w", err)
	}

	return nil
}
