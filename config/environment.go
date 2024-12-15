package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type EnvironmentConfig struct {
	// Environments
	ENVIRONMENT string `mapstructure:"ENVIRONMENT"`
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
	// Jira Software
	JIRA_OAUTH_CLIENT_ID string `mapstructure:"JIRA_OAUTH_CLIENT_ID"`
	JIRA_STATE           string `mapstructure:"JIRA_STATE"`
}

// Variables to be set at build time using -ldflags
var (
	ENVIRONMENT              string
	APP_URL                  string
	SERVER_REDIRECT_ADDR     string
	GCP_PROJECT_ID           string
	GCP_ACCOUNT_TYPE         string
	FB_ADMIN_PRIVATE_KEY     string
	FB_ADMIN_CLIENT_EMAIL    string
	SECRETS_MAN_ACCOUNT_TYPE string
	SECRETS_MAN_PRIVATE_KEY  string
	SECRETS_MAN_EMAIL        string
	JIRA_OAUTH_CLIENT_ID     string
	JIRA_STATE               string
)

func LoadEnvironment() (config EnvironmentConfig, err error) {
	if ENVIRONMENT != "production" {
		viper.AddConfigPath(".")
		viper.SetConfigName(".env.development")
		viper.SetConfigType("env")
		viper.AutomaticEnv()

		err = viper.ReadInConfig()
		if err != nil {
			log.Println(err)
			return
		}

		err = viper.Unmarshal(&config)
		if err != nil {
			fmt.Println("error: reading in environment config")
		}
		return config, nil
	}

	// Override with build-time variables if they are set
	if APP_URL != "" {

		config.APP_URL = APP_URL
	}
	if SERVER_REDIRECT_ADDR != "" {
		config.SERVER_REDIRECT_ADDR = SERVER_REDIRECT_ADDR
	}
	if GCP_PROJECT_ID != "" {
		config.GCP_PROJECT_ID = GCP_PROJECT_ID
	}
	if GCP_ACCOUNT_TYPE != "" {
		config.GCP_ACCOUNT_TYPE = GCP_ACCOUNT_TYPE
	}
	if FB_ADMIN_PRIVATE_KEY != "" {
		config.FB_ADMIN_PRIVATE_KEY = FB_ADMIN_PRIVATE_KEY
	}
	if FB_ADMIN_CLIENT_EMAIL != "" {
		config.FB_ADMIN_CLIENT_EMAIL = FB_ADMIN_CLIENT_EMAIL
	}
	if SECRETS_MAN_ACCOUNT_TYPE != "" {
		config.SECRETS_MAN_ACCOUNT_TYPE = SECRETS_MAN_ACCOUNT_TYPE
	}
	if SECRETS_MAN_PRIVATE_KEY != "" {
		config.SECRETS_MAN_PRIVATE_KEY = SECRETS_MAN_PRIVATE_KEY
	}
	if SECRETS_MAN_EMAIL != "" {
		config.SECRETS_MAN_EMAIL = SECRETS_MAN_EMAIL
	}
	if JIRA_OAUTH_CLIENT_ID != "" {
		config.JIRA_OAUTH_CLIENT_ID = JIRA_OAUTH_CLIENT_ID
	}
	if JIRA_STATE != "" {
		config.JIRA_STATE = JIRA_STATE
	}
	return config, nil
}
