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

	log.Println("Viper reading config")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("We are getting passed the error")

	err = viper.Unmarshal(&config)
	return
}
