package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"vx/tools/dirs"
)

const (
	authFile   = "user.json"
	configFile = "config.json"
	ConfigDir  = ".vx"
	pathFile   = "%s/%s"
)

// CheckConfigPath checks if the config directory exists
func CheckConfigPath() bool {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Println("dirs.HomeDirErrMsg")
	}
	cfgDir := fmt.Sprintf(pathFile, home, ConfigDir)
	if _, err := os.Stat(cfgDir); os.IsNotExist(err) {
		return false
	}
	return true
}

// InitConfig reads in config file and ENV variables if set.
func InitConfig() {
	pathExists := CheckConfigPath()
	if !pathExists {
		fmt.Println("Creating config")
		createConfigPath()
	}
}

// createConfigPath creates the .vx config path in the users $HOME directory
func createConfigPath() {
	wrtCfgDir()
	home, err := os.UserHomeDir()
	if err != nil {
		log.Println(dirs.HomeDirErrMsg)
	}
	sysAuthDir := filepath.Join(home, ConfigDir)
	wrtErr := wrtCfg(sysAuthDir, authFile, &AuthConfig{UID: "", Token: ""})
	if wrtErr != nil {
		log.Println(wrtErr)
	}
	sysConfigDir := filepath.Join(home, ConfigDir)
	wrtCfgFile := wrtCfg(sysConfigDir, configFile, nil)
	if wrtErr != nil {
		log.Println(wrtCfgFile)
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
	sidCfgDir := filepath.Join(home, ConfigDir)
	osErr := os.Mkdir(sidCfgDir, os.ModePerm)
	if osErr != nil {
		log.Println("error: error creating sind config directory", osErr)
	}
}

func wrtCfg(path, file string, code interface{}) error {
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
	sidDir := fmt.Sprintf(pathFile, home, ConfigDir)
	if _, err := os.Stat(sidDir); os.IsNotExist(err) {
		return false
	}
	return true
}

func CheckSupportedAPI(key string) ([]reflect.Value, bool) {
	supportedAPIs := map[string]string{
		"openai": "supported",
		"github": "supported",
	}
	keys := reflect.ValueOf(supportedAPIs).MapKeys()
	_, ok := supportedAPIs[key]
	return keys, ok
}
