package test

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"vx/tools"
)

type TestPayload struct {
	Sup bool `json:"sup"`
}

func TestHTTP() (bool, error) {
	req, err := tools.GetRequest("http://localhost:3000/api")
	if err != nil {
		log.Println("error:", err)
	}

	if req.StatusCode != http.StatusOK {
		return false, errors.New("unexpected status code: " + req.Status)
	}
	log.Println(req)
	requestBody, err := io.ReadAll(req.Body)
	if err != nil {
		return false, errors.New("unable to read in respoonse body")

	}
	var requestData TestPayload
	if err := json.Unmarshal(requestBody, &requestData); err != nil {
		log.Println("Error unmarshalling data", err)
	}
	log.Println(requestData)
	return true, nil
}
