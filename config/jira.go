package config

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

// JiraAuthCfg handles the OAuth2.0
type JiraAuthCfg struct {
	JiraToken string `mapstructure:"token"`
	JiraUID   string `mapstructure:"uid"`
}

// jiraCfg manages ./vx/jira.json config only used for OAuth2.0
var jiraCfg = viper.New()

// TODO: Theres a pattern forming here - explore abstraction
// Only tricky thing woudl be the viper instance
// Could be done at run time?
func SetJiraCfg(userId string, token string) {
	home, _ := os.UserHomeDir()
	jiraCfgPath := fmt.Sprintf("%s/%s/jira.json", home, ConfigDir)
	jiraCfg.SetConfigFile(jiraCfgPath)
	config := JiraAuthCfg{
		JiraToken: token,
	}
	jiraCfg.Set("token", config.JiraToken)
	err := jiraCfg.WriteConfig()
	if err != nil {
		log.Println("error saving user data to your system")
	}
	fmt.Println("All done")
}

func LoadJiraAuthCfg() (cfg JiraAuthCfg, err error) {
	home, _ := os.UserHomeDir()
	jiraCfgPath := fmt.Sprintf("%s/%s", home, ConfigDir)
	jiraCfg.AddConfigPath(jiraCfgPath)
	jiraCfg.SetConfigName("jira") // name of config file (without extension)
	jiraCfg.SetConfigType("json")
	jiraCfg.AutomaticEnv()

	err = jiraCfg.ReadInConfig()
	if err != nil {
		log.Println(err)
		return
	}

	err = jiraCfg.Unmarshal(&cfg)
	if err != nil {
		log.Println(err)
		return
	}
	return cfg, nil
}
