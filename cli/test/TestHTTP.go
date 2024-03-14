package test

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"vx/tools"
)

var client = tools.HTTPClient

type TestPayload struct {
	Sup bool `json:"sup"`
}

func TestHTTP() (bool, error) {
	req, err := client.Get("http://localhost:3000/api")
	if err != nil {
		log.Println("error:", err)
	}
	resp, respError := client.Do(req.Request)
	if respError != nil {
		log.Println("Error with server respoonse:", respError)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return false, errors.New("unexpected status code: " + resp.Status)
	}
	log.Println(resp)
	requestBody, err := io.ReadAll(resp.Body)
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
