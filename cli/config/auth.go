package config

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type AuthConfig struct {
	UID   string `mapstructure:"uid"`
	Token string `mapstructure:"token"`
}

func SetUserCfg(uid string, token string) {
	home, err := os.UserHomeDir()
	authCfgPath := fmt.Sprintf("%s/%s/user.json", home, CfgDir)
	cobra.CheckErr(err)
	viper.SetConfigFile(authCfgPath)
	config := AuthConfig{
		UID:   uid,
		Token: token,
	}
	viper.Set("uid", config.UID)
	viper.Set("token", config.Token)
	err = viper.WriteConfigAs(authCfgPath)
	if err != nil {
		log.Println("error saving user data to your system")
	}
}

func LoadAuth() (config AuthConfig, err error) {
	home, _ := os.UserHomeDir()
	authCfgPath := fmt.Sprintf("%s/%s", home, CfgDir)
	viper.AddConfigPath(authCfgPath)
	viper.SetConfigName("user") // name of config file (without extension)
	viper.SetConfigType("json")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		log.Println(err)
		return
	}

	err = viper.Unmarshal(&config)
	return
}
