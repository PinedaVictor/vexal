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
		smzDirectory()
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

	log.Println("the resp data:", string(body))
	formattedContent := strings.ReplaceAll(string(body), "\n", "\n\n")
	// TODO: File writing and formatting
	file, err := os.Create("readme.md")
	if err != nil {
		log.Fatal("error creating file:", err)
	}
	defer file.Close()

	_, err = file.WriteString(formattedContent)
	if err != nil {
		log.Fatal("error writing to file:", err)
	}

}

func smzDirectory() {}
