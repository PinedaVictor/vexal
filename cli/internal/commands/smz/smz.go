// Package smz handles all the internal logic of the smz command
package smz

import (
	"fmt"
	"log"
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
	c := paths.GetContent(pathSmz)
	cfg := viper.GetString("openaikey")
	env, _ := config.LoadEnvironment()
	route := fmt.Sprintf("%s/api/ai/smz", env.API_URL)
	// TODO: Call server
	resp, err := tools.PostRequest(route,
		map[string]string{"openai": cfg},
		map[string]interface{}{"content": c, "entity": entity})
	if err != nil {
		log.Fatal("error making smz requests")
	}
	defer resp.Body.Close()

	log.Println("the resp:", resp.Status)

}

func smzDirectory() {}
