package config

import (
	"log"

	"github.com/spf13/viper"
)

type EnvironmentConfig struct {
	API_URL              string `mapstructure:"API_URL"`
	SERVER_REDIRECT_ADDR string `mapstructure:"SERVER_REDIRECT_ADDR"`
}

func LoadEnvironment() (config EnvironmentConfig, err error) {
	viper.AddConfigPath(".")
	viper.SetConfigName("app.env")
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
