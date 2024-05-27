package config

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

type AuthConfig struct {
	UID   string `mapstructure:"uid"`
	Token string `mapstructure:"token"`
}

var authCfg = viper.New()

func SetUserCfg(uid string, token string) {
	home, _ := os.UserHomeDir()
	authCfgPath := fmt.Sprintf("%s/%s/user.json", home, ConfigDir)
	// viper.AddConfigPath(authCfgPath)
	// cobra.CheckErr(err)
	authCfg.SetConfigFile(authCfgPath)
	config := AuthConfig{
		UID:   uid,
		Token: token,
	}

	authCfg.Set("uid", config.UID)
	authCfg.Set("token", config.Token)

	err := authCfg.WriteConfig()
	if err != nil {
		log.Println("error saving user data to your system")
	}
}

func LoadAuth() (config AuthConfig, err error) {
	home, _ := os.UserHomeDir()
	authCfgPath := fmt.Sprintf("%s/%s", home, ConfigDir)
	authCfg.AddConfigPath(authCfgPath)
	authCfg.SetConfigName("user") // name of config file (without extension)
	authCfg.SetConfigType("json")
	authCfg.AutomaticEnv()

	err = authCfg.ReadInConfig()
	if err != nil {
		log.Println(err)
		return
	}

	err = authCfg.Unmarshal(&config)
	return
}
