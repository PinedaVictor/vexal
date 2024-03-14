package tools

// import (
// 	"io"
// 	"net/http"
// )

// // Client is an interface for making HTTP requests
// type Client interface {
// 	Get(url string) (*http.Response, error)
// 	Post(url, contentType string, body io.Reader) (*http.Response, error)
// 	// Add other HTTP methods as needed
// }

// // DefaultClient implements the Client interface using net/http
// type DefaultClient struct{}

// func (c *DefaultClient) Get(url string) (*http.Response, error) {
// 	return http.Get(url)
// }

// func (c *DefaultClient) Post(url, contentType string, body io.Reader) (*http.Response, error) {
// 	return http.Post(url, contentType, body)
// }

// // package httpclient is a wrapper for http.Client
// package httpclient

import (
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

// NewClient creates a new instance of Client with default settings.
func NewClient() *Client {
	return &Client{
		client: &http.Client{},
	}
}

// Do sends an HTTP request and returns an HTTP response.
func (c *Client) Do(req *http.Request) (*http.Response, error) {
	// If needed add custom logic or pre-processing here
	return c.client.Do(req)
}

func (c *Client) Get(url string) (*http.Response, error) {
	return http.Get(url)
}

func (c *Client) Post(url, contentType string, body io.Reader) (*http.Response, error) {
	return http.Post(url, contentType, body)
}

// GetAPIURL returns the URL for the API server
func GetAPIURL() string {
	viper.AutomaticEnv()
	apiURL := viper.GetString("API_URL")
	if apiURL == "" {
		return "API_URL is not defined"
	}
	return apiURL
}
