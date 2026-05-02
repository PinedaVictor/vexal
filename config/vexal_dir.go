/*
Copyright © 2026 Victor Pineda pinedavictor095@gmail.com
*/
package config

import (
	"fmt"
	"os"
	"vx/pkg/paths"
)

const vexalDir = ".vexal"

// InitVexalDir creates the .vexal directory in the current repo root.
func InitVexalDir() error {
	curDir, _ := os.Getwd()
	dir := fmt.Sprintf("%s/%s", curDir, vexalDir)
	if paths.PathExists(dir) {
		return nil
	}
	return os.MkdirAll(dir, 0755)
}
