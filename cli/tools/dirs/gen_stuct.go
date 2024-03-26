// Package dirs handles all directory tooling
package dirs

import (
	"fmt"
	"log"
	"os"
)

const (
	HomeDirErrMsg = "error getting home directory"
)

func getCurDir() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("error: getting current directory: %w", err)
	}
	return dir, nil
}

func genRoot(root string) error {
	curDir, _ := getCurDir()
	if curDir != "" {
		dir := fmt.Sprintf("%s/%s", curDir, root)
		err := os.Mkdir(dir, os.ModePerm)
		if err != nil {
			return fmt.Errorf("error: generating root: %w", err)
		}
	}
	return nil
}

func genPath(paths []string, srcDir string) error {
	for _, path := range paths {
		err := os.Mkdir(fmt.Sprintf("%s/%s", srcDir, path), os.ModePerm)
		if err != nil {
			return fmt.Errorf("error: generating paths: %w", err)
		}
	}
	return nil
}

// GenStruct is used to create new folder structure
func GenStruct(root, src string, dirs []string) error {
	srcDir := fmt.Sprintf("%s/%s", root, src)
	_, err1 := os.Open(root)
	if err1 != nil {
		if os.IsNotExist(err1) {
			rtErr1 := genRoot(root)
			if rtErr1 != nil {
				log.Println(rtErr1)
			}
		}
	}
	_, err := os.Open(srcDir)
	if err != nil {
		if os.IsNotExist(err) {
			rtErr2 := genRoot(srcDir)
			if rtErr2 != nil {
				log.Println(rtErr2)
			}
		}
	}
	pathErr := genPath(dirs, srcDir)
	if pathErr != nil {
		return fmt.Errorf("Error generating paths: %s %s", dirs, pathErr)
	}
	return nil
}
