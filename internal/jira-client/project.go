package jiraclient

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type AvatarUrls struct {
	X48x48 string `json:"48x48"`
	X24x24 string `json:"24x24"`
	X16x16 string `json:"16x16"`
	X32x32 string `json:"32x32"`
}

type JiraProject struct {
	Expand         string                 `json:"expand"`
	Self           string                 `json:"self"`
	ID             string                 `json:"id"`
	Key            string                 `json:"key"`
	Name           string                 `json:"name"`
	AvatarUrls     AvatarUrls             `json:"avatarUrls"`
	ProjectTypeKey string                 `json:"projectTypeKey"`
	Simplified     bool                   `json:"simplified"`
	Style          string                 `json:"style"`
	IsPrivate      bool                   `json:"isPrivate"`
	Properties     map[string]interface{} `json:"properties"`
	EntityID       string                 `json:"entityId"`
	UUID           string                 `json:"uuid"`
}

// TODO: Update - this functionality is depreicated in v3
func GetAllProjects() []JiraProject {
	resp, err := JiraAPIGet("/rest/api/3/project")
	if err != nil {
		log.Println("error gettting Jira project meta data:", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Received non-OK response: %v", resp.Status)
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	var projects []JiraProject
	err = json.Unmarshal(body, &projects)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	for _, project := range projects {
		fmt.Printf("Project Name: %s, ID: %s, Key: %s\n", project.Name, project.ID, project.Key)
	}
	return projects
}

func GetJiraPrjtMeta(projectKey string) JiraProject {
	url := fmt.Sprintf("/rest/api/3/project/%s", projectKey)
	resp, err := JiraAPIGet(url)
	if err != nil {
		log.Println("error gettting Jira project meta data:", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Received non-OK response: %v", resp.Status)
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	var project JiraProject
	err = json.Unmarshal(body, &project)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	return project
}
