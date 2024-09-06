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

func GetIssueTypes() string {
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

	return string(body)
}

func parseIssueTypes(issueTypeName string) (string, string, string) {
	bodyJSON := GetIssueTypes()

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

func CreateIssue(issueTypeID string, projctKey string, summary string, description string) {
	repoCfg, _ := config.LoadRepoConfig()
	url := fmt.Sprintf("%s/%s/rest/api/3/issue", OAuth2URL, repoCfg.Jira_Cloud_ID)
	headers := getJiraOAuthHeaders()
	fmt.Println("Updating: ", url)

	// Define the payload (JSON data)
	payload := map[string]interface{}{
		"fields": map[string]interface{}{
			"summary": summary,
			// "summary": "Main order flow broken",
			"issuetype": map[string]string{
				// TODO: This will likely call a getIsssueId function
				// "id": "10003", // Replace with the valid ID
				"id": issueTypeID,
			},
			"project": map[string]string{
				// TODO: Get project Meta data
				// "key": "SCRUM", // Ensure the project key is correct
				"key": projctKey, // Ensure the project key is correct
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
								"text": description,
								// "text": "THIS IS TESTING WITH GO REST API CALL.",
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
