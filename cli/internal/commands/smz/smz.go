// Package smz handles all the internal logic of the smz command
package smz

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"vx/config"
	"vx/pkg/paths"
	"vx/tools"

	"github.com/spf13/viper"
)

var pathSmz = ""

func SMZ(t string, path string, entity string) {
	pathSmz = path
	if t == paths.FP {
		smzFile(entity)
	}
	if t == paths.DP {
		smzDirectory(entity)
	}
}

func smzFile(entity string) {
	// TODO: Some of this can be abstracted out
	c := paths.GetContent(pathSmz)
	cfg := viper.GetString("openaikey")
	env, _ := config.LoadEnvironment()
	route := fmt.Sprintf("%s/api/ai/smz", env.API_URL)
	resp, err := tools.PostRequest(route,
		map[string]string{"openai": cfg},
		map[string]interface{}{"content": c, "entity": entity})
	if err != nil {
		log.Fatal("error making smz requests")
	}
	defer resp.Body.Close()
	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("error reading response body:", err)
	}

	formattedContent := strings.ReplaceAll(string(body), "\n", "\n\n")
	// TODO: File writing and formatting
	currentDir, _ := os.Getwd()
	file, err := os.Create(fmt.Sprintf("%s/%s", currentDir, "readme.md"))
	if err != nil {
		log.Fatal("error creating file:", err)
	}
	defer file.Close()

	_, err = file.WriteString(formattedContent)
	if err != nil {
		log.Fatal("error writing to file:", err)
	}
}

// TODO:
func smzDirectory(entity string) {
	log.Println("Under construction: Directory:", entity)
}

func appendToFile(path string) {
	// Open the file in append mode. Create the file if it doesn't exist.
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Data to append to the file
	data := "This is some new data to append."

	// Add a newline character to the data
	dataWithNewline := "\n" + data

	// Write the data to the file
	if _, err := file.WriteString(dataWithNewline); err != nil {
		log.Fatal(err)
	}

	// Output success message
	log.Println("Data appended to file successfully.")
}
