package test

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"vx/config"
	"vx/tools"
)

type TestPayload struct {
	Sup bool `json:"sup"`
}

func TestHTTP() (bool, error) {
	env, _ := config.LoadEnvironment()
	user, err := config.LoadAuth()
	if err != nil {
		fmt.Println("Error reading auth:", err)
	}
	route := fmt.Sprintf("%s/api", env.API_URL)
	req, err := tools.GetRequest(route, map[string]string{"Authorization": user.UID})
	if err != nil {
		log.Println(`
		The service you're trying to use requires and internet connection.
		Failed to connect to the Vexal server. Please ensure your internet
		connection is stable and try again.
		`)
		os.Exit(0)
		return false, err
	}

	if req.StatusCode != http.StatusOK {
		return false, errors.New("unexpected status code: " + req.Status)
	}
	requestBody, err := io.ReadAll(req.Body)
	if err != nil {
		return false, errors.New("unable to read in respoonse body")

	}
	var requestData TestPayload
	if err := json.Unmarshal(requestBody, &requestData); err != nil {
		log.Println("Error unmarshalling data:", err)
	}
	// log.Println("Connected to Vexal server:", requestData.Sup)
	return true, nil
}
