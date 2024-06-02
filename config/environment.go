package config

import (
	"log"

	"github.com/spf13/viper"
)

type EnvironmentConfig struct {
	// Apps
	APP_URL              string `mapstructure:"APP_URL"`
	API_URL              string `mapstructure:"API_URL"`
	SERVER_REDIRECT_ADDR string `mapstructure:"SERVER_REDIRECT_ADDR"`
	// GCP
	GCP_PROJECT_ID   string `mapstructure:"GCP_PROJECT_ID"`
	GCP_ACCOUNT_TYPE string `mapstructure:"GCP_ACCOUNT_TYPE"`
	// Firbase
	FB_ADMIN_PRIVATE_KEY  string `mapstructure:"FB_ADMIN_PRIVATE_KEY"`
	FB_ADMIN_CLIENT_EMAIL string `mapstructure:"FB_ADMIN_CLIENT_EMAIL"`
	// Secrets Manager
	SECRETS_MAN_ACCOUNT_TYPE string `mapstructure:"SECRETS_MAN_ACCOUNT_TYPE"`
	SECRETS_MAN_PRIVATE_KEY  string `mapstructure:"SECRETS_MAN_PRIVATE_KEY"`
	SECRETS_MAN_EMAIL        string `mapstructure:"SECRETS_MAN_EMAIL"`
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
