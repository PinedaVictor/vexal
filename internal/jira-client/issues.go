package jiraclient

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"vx/config"
	"vx/pkg"
)

type JiraIssueType struct {
	Self             string             `json:"self"`
	ID               string             `json:"id"`
	Description      string             `json:"description"`
	IconURL          string             `json:"iconUrl"`
	Name             string             `json:"name"`
	UntranslatedName string             `json:"untranslatedName"`
	Subtask          bool               `json:"subtask"`
	AvatarID         int                `json:"avatarId"`
	HierarchyLevel   int                `json:"hierarchyLevel"`
	Scope            JiraIssueTypeScope `json:"scope"`
}

type JiraIssueTypeScope struct {
	Type    string `json:"type"`
	Project struct {
		ID string `json:"id"`
	} `json:"project"`
}

func GetIssueTypes() {
	baseURL := getJiraOAuthURL()
	url := fmt.Sprintf("%s/rest/api/3/issuetype", baseURL)
	hdrs := getJiraOAuthHeaders()

	resp, err := pkg.GetRequest(url, hdrs)
	if err != nil {
		log.Println("error getting issues types")
	}
	defer resp.Body.Close()
	// Check if the response status is not OK
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Received non-OK response: %v", resp.Status)
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}
	issueID, name, _ := parseIssueTypes(string(body), "story")
	fmt.Println("Issue ID:", issueID, name)
}

func parseIssueTypes(issueBodyResp string, issueTypeName string) (string, string, string) {
	// Create Parsing function
	bodyJSON := string(issueBodyResp) // Your JSON string here

	var issueTypes []JiraIssueType
	JSONErr := json.Unmarshal([]byte(bodyJSON), &issueTypes)
	if JSONErr != nil {
		log.Fatalf("Error unmarshalling JSON: %v", JSONErr)
	}

	if issueTypeName == "todo" {
		issueTypeName = "Task"
	} else if issueTypeName == "fixme" {
		issueTypeName = "Bug"
	}
	// Use the parsed data
	for _, issueType := range issueTypes {
		if strings.EqualFold(issueType.Name, issueTypeName) {
			return issueType.ID, issueType.Name, issueType.Description
		}
	}
	return "", "", ""
}

// ALG:
// 1. Get Issue Type IDs
// 2. Scrape codebase
// 3. Use Go concurrency to update Jira per issueType
// ---3.1 How will Vexal make sure NOT to create dublicates????
func CreateIssue() {
	cfg, cfgErr := config.LoadJiraAuthCfg()
	if cfgErr != nil {
		log.Println("error reading jira auth:", cfgErr)
	}
	repoCfg, _ := config.LoadRepoConfig()
	url := fmt.Sprintf("https://api.atlassian.com/ex/jira/%s/rest/api/3/issue", repoCfg.Jira_Cloud_ID)
	fmt.Println("Updating: ", url)

	OAuthToken := fmt.Sprintf("Bearer %s", cfg.JiraToken)
	// Define the headers in a map
	headers := map[string]string{
		"Accept":        "application/json",
		"Content-Type":  "application/json",
		"Authorization": OAuthToken,
	}
	// Define the payload (JSON data)
	payload := map[string]interface{}{
		"fields": map[string]interface{}{
			"summary": "Main order flow broken",
			"issuetype": map[string]string{
				// TODO: This will likely call a getIsssueId function
				"id": "10003", // Replace with the valid ID
			},
			"project": map[string]string{
				// TODO: Get project Meta data
				"key": "SCRUM", // Ensure the project key is correct
			},
			"description": map[string]interface{}{
				"type":    "doc",
				"version": 1,
				"content": []map[string]interface{}{
					{
						"type": "paragraph",
						"content": []map[string]interface{}{
							{
								"type": "text",
								"text": "THIS IS TESTING WITH GO REST API CALL.",
							},
						},
					},
				},
			},
		},
	}

	fmt.Println("PAYLOAD:", payload)
	resp, err := pkg.PostRequest(url, headers, payload)
	if err != nil {
		log.Println("error creating issue:", err)
	}
	fmt.Println(resp)
}
