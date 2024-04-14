// Package smz handles all the internal logic of the smz command
package smz

import (
	"log"
	"vx/pkg/paths"
)

var pathSmz = ""

func SMZ(t string, path string) {
	pathSmz = path
	if t == paths.FP {
		smzFile()
	}
	if t == paths.DP {
		smzDirectory()
	}
}

func smzFile() {
	log.Println("---- smzFile -----")
	log.Println("path:", pathSmz)
	c := paths.GetContent(pathSmz)
	log.Println("File content:", c)
	// TODO: Call server
}

func smzDirectory() {}
