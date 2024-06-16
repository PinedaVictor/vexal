package vxrequest

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"vx/config"
	"vx/pkg"
)

type VxResponse struct {
	Data string `json:"data"`
}

func VxGet(endpoint string, headers map[string]string) (bool, error) {
	env, _ := config.LoadEnvironment()
	user, _ := config.LoadAuth()
	auth := map[string]string{"authorization": user.Token, "Content-Type": "application/json"}
	combinedHeaders := make(map[string]string)
	// Copy auth headers into the combined map
	for key, value := range auth {
		combinedHeaders[key] = value
	}
	// Copy headers from the parameters into the combined map
	for key, value := range headers {
		combinedHeaders[key] = value
	}
	// fmt.Println(combinedHeaders)
	apiURL := fmt.Sprintf("%s/api/%s", env.APP_URL, endpoint)
	resp, err := pkg.GetRequest(apiURL, combinedHeaders)
	if err != nil {
		log.Fatal("error:", err)
	}
	// if resp.StatusCode != http.StatusOK {
	// 	return false, errors.New("unexpected status code: " + resp.Status)
	// }
	requestBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return false, errors.New("unable to read in respoonse body")

	}
	var data VxResponse
	if err := json.Unmarshal(requestBody, &data); err != nil {
		log.Println("Error unmarshalling data:", err)
	}
	fmt.Println("req:", data)
	return true, nil
}
