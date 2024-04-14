// Package directories handles the directory generation and management
package directories

import (
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"
)

const (
	FP = "file"
	DP = "directory"
)

// HomeDir returns the systems home directory
func HomeDir() (string, error) {
	// Get the home directory path
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Println("error getting users home directory", err)
		return "", errors.New("error getting user home directory")
	}
	return homeDir, nil
}

// DeterminePath returens a string "file" or "directory" depending on the file mode. Empty string for none
func DeterminePath(input string) string {
	if isFile(input) {
		log.Println(input, "is a file")
		return FP
	}
	if isDirectory(input) {
		log.Println(input, "is a directory")
		return DP
	}
	log.Println(input, "is niether a file or a directory")
	return ""
}

// isFile determines if a path is a file
func isFile(input string) bool {
	info := getFileInfo(input)
	return info.Mode().IsRegular()
}

// isDirectory determines if a path is a directory
func isDirectory(input string) bool {
	info := getFileInfo(input)
	return info.IsDir()
}

// getFileInfo gets path info using os pkg
func getFileInfo(input string) fs.FileInfo {
	currentDir, _ := os.Getwd()
	path := fmt.Sprintf("%s/%s", currentDir, input)
	log.Println("usrs current working directory:", currentDir)
	info, infoErr := os.Stat(path)
	if infoErr != nil {
		log.Fatal("error reading path:", infoErr)
	}
	return info
}
