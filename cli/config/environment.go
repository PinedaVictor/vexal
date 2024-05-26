package config

import (
	"log"

	"github.com/spf13/viper"
)

type EnvironmentConfig struct {
	API_URL               string `mapstructure:"API_URL"`
	GCP_PROJECT_ID        string `mapstructure:"GCP_PROJECT_ID"`
	SERVER_REDIRECT_ADDR  string `mapstructure:"SERVER_REDIRECT_ADDR"`
	FB_ADMIN_PRIVATE_KEY  string `mapstructure:"FB_ADMIN_PRIVATE_KEY"`
	FB_ADMIN_ACCOUNT_TYPE string `mapstructure:"FB_ADMIN_ACCOUNT_TYPE"`
	FB_ADMIN_CLIENT_EMAIL string `mapstructure:"FB_ADMIN_CLIENT_EMAIL"`
}

func LoadEnvironment() (config EnvironmentConfig, err error) {
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		log.Println(err)
		return
	}

	err = viper.Unmarshal(&config)
	return
}
