// Package paths handles path logic content
package paths

import (
	"fmt"
	"os"
)

func GetContent(path string) string {
	// Open the file
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return ""
	}
	defer file.Close()

	// Get the file stat to determine its size
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("Error getting file info:", err)
		return ""
	}

	// Create a byte slice with the file size
	content := make([]byte, fileInfo.Size())

	// Read the file content into the byte slice
	_, err = file.Read(content)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return ""
	}

	// Print file content
	fmt.Println(string(content))
	return string(content)
}
