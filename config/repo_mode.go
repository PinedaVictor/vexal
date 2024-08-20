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
	Openai_key string `mapstructure:"openai_key"`
	// Github
	Github_key string `mapstructure:"github_key"`
	Repo       string `mapstructure:"repo"`
	RepoURL    string `mapstructure:"repo_URL"`
	// Jira Software
	Jira_Name     string `mapstructure:"jira_name"`
	Jira_Email    string `mapstructure:"jira_email"`
	Jira_URL      string `mapstructure:"jira_url"`
	Jira_Cloud_ID string `mapstructure:"jira_cloud_id"`
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
	// Load the existing configuration
	config, err := LoadRepoConfig()
	if err != nil {
		return fmt.Errorf("error loading config: %w", err)
	}

	config.Jira_Name = "jira_name"
	config.Jira_URL = "your_jira_url"
	config.Jira_Cloud_ID = "your_jira_cloud_id"
	config.Jira_Email = "your_jira_email"

	// Set the updated Jira settings back to Viper
	repoMode.Set("jira_url", config.Jira_URL)
	repoMode.Set("jira_cloud_id", config.Jira_Cloud_ID)
	repoMode.Set("jira_email", config.Jira_Email)
	repoMode.Set("jira_name", config.Jira_Name)

	// Write the updated config back to the file
	if err := repoMode.WriteConfig(); err != nil {
		return fmt.Errorf("error writing config file: %w", err)
	}

	return nil
}

func UpdateJiraRepoCfg(name string, url string, id string) error {
	config, err := LoadRepoConfig()
	if err != nil {
		return fmt.Errorf("error loading config: %w", err)
	}
	config.Jira_Name = name
	config.Jira_URL = url
	config.Jira_Cloud_ID = id

	// Set the updated Jira settings back to Viper
	repoMode.Set("jira_url", config.Jira_URL)
	repoMode.Set("jira_cloud_id", config.Jira_Cloud_ID)
	repoMode.Set("jira_name", config.Jira_Name)
	if err := repoMode.WriteConfig(); err != nil {
		return fmt.Errorf("error writing config file: %w", err)
	}

	return nil
}
