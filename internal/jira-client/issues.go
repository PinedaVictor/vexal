package jiraclient

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
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
	parseIssueTypes(string(body))
}

func parseIssueTypes(issueBodyResp string) {
	// Create Parsing function
	bodyJSON := string(issueBodyResp) // Your JSON string here

	var issueTypes []JiraIssueType
	JSONErr := json.Unmarshal([]byte(bodyJSON), &issueTypes)
	if JSONErr != nil {
		log.Fatalf("Error unmarshalling JSON: %v", JSONErr)
	}

	// Use the parsed data
	for _, issueType := range issueTypes {
		fmt.Printf("ID: %s, Name: %s, Description: %s\n", issueType.ID, issueType.Name, issueType.Description)
	}
	// TODO: Parse Data types
}
