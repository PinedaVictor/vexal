package tools

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Config struct {
	OpenAIKey string `mapstructure:"OpenAIKey"`
	Username  string `mapstructire:"Username"`
}

var defaultConfig = Config{
	OpenAIKey: "",
	Username:  "",
}

const (
	configDir = ".vx"
	pathFile  = "%s/%s"
)

// CheckConfigPath checks if the config directory exists
func CheckConfigPath() bool {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Println("dirs.HomeDirErrMsg")
	}
	cfgDir := fmt.Sprintf(pathFile, home, configDir)
	if _, err := os.Stat(cfgDir); os.IsNotExist(err) {
		return false
	}
	return true
}

// InitConfig reads in config file and ENV variables if set.
func InitConfig() {
	pathExists := CheckConfigPath()
	if !pathExists {
		// TODO: Implment createConfigPath
		log.Println("Config Path does not exits - creating one")
		// cfgFilePath()
	}
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)
	// Construct the directory path
	cfgPath := filepath.Join(home, configDir)

	// Construct the file path
	cfgFilePath := filepath.Join(cfgPath, "config.json")

	viper.AddConfigPath(cfgPath)
	viper.SetConfigType("json")
	viper.SetConfigName("config")
	viper.SetConfigFile(cfgFilePath)

	viper.AutomaticEnv()       // read in environment variables that match
	er := viper.ReadInConfig() // Find and read the config file
	if er != nil {             // Handle errors reading the config file
		log.Println("fatal error config file: %w", err)
	}
}

// createConfigPath creates the .vx config path in the users $HOME directory
func createConfigPath() {

}
