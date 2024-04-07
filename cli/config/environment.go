package config

import (
	"log"

	"github.com/spf13/viper"
)

type EnvironmentConfig struct {
	API_URL string `mapstructure:"API_URL"`
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
