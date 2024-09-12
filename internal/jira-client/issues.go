package jiraclient

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"vx/config"
	"vx/internal"
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
	cfg, _ := config.LoadRepoConfig()
	prj := GetJiraPrjtMeta(cfg.Jira_Project_key)
	url := fmt.Sprintf("/rest/api/3/issuetype/project?projectId=%s", prj.ID)
	resp, err := JiraAPIGet(url)
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

func ListIssueTypes() {
	issueTypes := GetIssueTypes()
	parsedIssueTYpes := ParseIssueTypeJSON(issueTypes)
	for _, issueType := range parsedIssueTYpes {
		internal.UserFeedback("ID: " + issueType.ID + " Name: " + issueType.Name + " Description: " + issueType.Description)
	}
}

func ParseIssueTypeJSON(bodyJSON string) []JiraIssueType {
	var issueTypes []JiraIssueType
	JSONErr := json.Unmarshal([]byte(bodyJSON), &issueTypes)
	if JSONErr != nil {
		log.Fatalf("Error unmarshalling JSON: %v", JSONErr)
	}
	return issueTypes
}

func parseIssueTypes(issueTypeName string) (string, string, string) {
	bodyJSON := GetIssueTypes()
	issueTypes := ParseIssueTypeJSON(bodyJSON)

	if issueTypeName == "todo" {
		issueTypeName = "Story"
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
	// Define the payload (JSON data)
	payload := map[string]interface{}{
		"fields": map[string]interface{}{
			"summary": summary,
			"issuetype": map[string]string{
				"id": issueTypeID,
			},
			"project": map[string]string{
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
							},
						},
					},
				},
			},
		},
	}

	resp, err := JiraAPIPost("/rest/api/3/issue", payload)
	if err != nil {
		log.Println("error creating issue", err)
	}
	if resp.StatusCode == http.StatusOK {
		fmt.Println("Push to Jira Successful")
	}
}
