package config

import (
	"fmt"
	"strings"
)

var supportedAPIsMap = map[string]struct{}{
	"github": {},
	"openai": {},
}

func APISupported(API string) bool {
	enableAPI := strings.ToLower(API)
	_, supported := supportedAPIsMap[enableAPI]
	return supported
}

func PrintSupportedAPIs() {
	fmt.Println("Supported APIs:")
	for api := range supportedAPIsMap {
		fmt.Println("- " + api)
	}
}
