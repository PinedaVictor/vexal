// package tools
package tools

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/spf13/viper"
)

// Client wraps http.Client and provides additional functionality
type Client struct {
	client *http.Client
}

// HTTPClient serves as a way to make API requests
var HTTPClient = NewClient()

// Do sends an HTTP request and returns an HTTP response.
func (c *Client) Do(req *http.Request) (*http.Response, error) {
	// If needed add custom logic or pre-processing here
	return c.client.Do(req)
}

// NewClient creates a new instance of Client with default settings.
func NewClient() *Client {
	return &Client{
		client: &http.Client{},
	}
}

// GetRequest performs an HTTP GET request and returns the response and any error.
func GetRequest(url string) (*http.Response, error) {
	req, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	resp, err := HTTPClient.Do(req.Request)
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()
	return resp, nil
}

// TODO: Test PostRquest
// PostRequest performs an HTTP POST request with the specified payload and returns the response and any error.
func PostRequest(url string, payload interface{}) (*http.Response, error) {
	var reqBody io.Reader
	if payload != nil {
		jsonData, err := json.Marshal(payload)
		if err != nil {
			return nil, errors.New("error marshalling payload: " + err.Error())
		}
		reqBody = bytes.NewBuffer(jsonData)
	}

	req, err := http.NewRequest(http.MethodPost, url, reqBody)
	if err != nil {
		return nil, err
	}

	resp, err := HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetAPIURL returns the URL for the API server
func GetAPIURL() (string, error) {
	viper.AutomaticEnv()
	apiURL := viper.GetString("API_URL")
	if apiURL == "" {
		return "", errors.New("api_url is not defined")
	}
	return apiURL, nil
}
