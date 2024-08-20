package jiraclient

import (
	"fmt"
	"log"
	"net/http"
	"vx/config"
	"vx/pkg"
)

var (
	// OAuth2URL: https://api.atlassian.com/ex/jira is the base URL for Jira OAuth2.0 calls.
	OAuth2URL = "https://api.atlassian.com/ex/jira"
)

// getJiraOAuthURL returns the the OAuth2.0 base URL interpolated with the cloud id fetched from LoadRepoConfig()
// Jira Apps URL: https://api.atlassian.com/ex/jira/{cloudid}
func getJiraOAuthURL() string {
	cfg, _ := config.LoadRepoConfig()
	return fmt.Sprintf("%s/%s", OAuth2URL, cfg.Jira_Cloud_ID)
}

// getJiraOAuthHeaders returns configured HTTP headers for making REST API calls to Jira.
func getJiraOAuthHeaders() map[string]string {
	cfg, _ := config.LoadJiraAuthCfg()
	OAuthToken := fmt.Sprintf("Bearer %s", cfg.JiraToken)
	// Define the headers in a map
	headers := map[string]string{
		"Accept":        "application/json",
		"Content-Type":  "application/json",
		"Authorization": OAuthToken,
	}
	return headers
}

// JiraAPIGet is a GET request pre-configureed with Jira base URL and Authorization headers
func JiraAPIGet(API string) (*http.Response, error) {
	baseURL := getJiraOAuthURL()
	url := fmt.Sprintf("%s%s", baseURL, API)
	hdrs := getJiraOAuthHeaders()

	resp, err := pkg.GetRequest(url, hdrs)
	if err != nil {
		log.Println("error getting issues types")
	}
	return resp, err
}
