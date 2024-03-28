package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"vx/tools/dirs"

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
	configFile = "config.json"
	configDir  = ".vx"
	pathFile   = "%s/%s"
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
		log.Println("Config Path does not exits - creating one")
		createConfigPath()
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

// LoadConfig will print sid configuration to the screen
func LoadConfig() {
	c := viper.AllSettings()
	b, err := json.MarshalIndent(c, " ", " ")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(b))
}

// createConfigPath creates the .vx config path in the users $HOME directory
func createConfigPath() {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Println(dirs.HomeDirErrMsg)
	}
	sysConfigDir := filepath.Join(home, configDir)
	wrtCfgDir()
	wrtErr := wrtCfg(sysConfigDir, configFile, defaultConfig)
	if wrtErr != nil {
		log.Println(wrtErr)
	}
	dirExists := checkConfigFile()
	if !dirExists {
		log.Println("unable to create config directory", sysConfigDir)
	}
	log.Println("config directory created successfully at", sysConfigDir)
}

// wrtCfgDir creates the $HOME/.vx directory
func wrtCfgDir() {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Println(dirs.HomeDirErrMsg)
	}
	sidCfgDir := filepath.Join(home, configDir)
	osErr := os.Mkdir(sidCfgDir, os.ModePerm)
	if osErr != nil {
		log.Println("error: error creating sind config directory", osErr)
	}
}

func wrtCfg(path, file string, code Config) error {
	aim := fmt.Sprintf(pathFile, path, file)
	newFile, errFile := os.Create(aim)
	if errFile != nil {
		log.Println(errFile)
		return fmt.Errorf("\n Error creating file %w", errFile)
	}
	defer newFile.Close()
	str, err := json.Marshal(code)
	if err != nil {
		log.Println("Error marchalling config json")
	}
	s := string(str)
	defer newFile.WriteString(strings.ToLower(s))
	return nil
}

func checkConfigFile() bool {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Println(dirs.HomeDirErrMsg)
	}
	sidDir := fmt.Sprintf(pathFile, home, configDir)
	if _, err := os.Stat(sidDir); os.IsNotExist(err) {
		return false
	}
	return true
}
