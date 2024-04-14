// Package paths handles path logic for stat
package paths

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
func DeterminePath(input string) (string, string) {
	if ok, path := isFile(input); ok {
		log.Println(input, "is a file")
		return FP, path
	}
	if ok, path := isDirectory(input); ok {
		log.Println(input, "is a directory")
		return DP, path
	}
	log.Println(input, "is niether a file or a directory")
	return "", ""
}

// isFile determines if a path is a file
func isFile(input string) (bool, string) {
	info, p := getFileInfo(input)
	return info.Mode().IsRegular(), p
}

// isDirectory determines if a path is a directory
func isDirectory(input string) (bool, string) {
	info, p := getFileInfo(input)
	return info.IsDir(), p
}

// getFileInfo gets path info using os pkg
func getFileInfo(input string) (fs.FileInfo, string) {
	currentDir, _ := os.Getwd()
	path := fmt.Sprintf("%s/%s", currentDir, input)
	info, infoErr := os.Stat(path)
	if infoErr != nil {
		log.Fatal("error reading path:", infoErr)
	}
	return info, path
}
