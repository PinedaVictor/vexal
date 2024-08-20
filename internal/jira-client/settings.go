package jiraclient

import (
	"fmt"
	"log"
	"vx/config"
	"vx/pkg"
)

// --url 'https://your-domain.atlassian.net/rest/api/2/issue' \
// --user 'email@example.com:<api_token>' \
// --header 'Accept: application/json' \
// --header 'Content-Type: application/json' \

// Jira apps: https://api.atlassian.com/ex/jira/{cloudid}/{api}
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

// TODO: Get Project Meta data - this will give you the cloud ID to save
// curl --request GET \
//   --url https://api.atlassian.com/ex/jira/{cloud_id}/rest/api/2/project \
//   --header 'Authorization: Bearer aBCxYz654123' \
//   --header 'Accept: application/json'

// [{"expand":"description,lead,issueTypes,url,projectKeys,permissions,insight",
// "self":"https://api.atlassian.com/ex/jira/{cloudid}/rest/api/3/project/10000",
// "id":"10000",
// "key":"SCRUM",
// "name":"integration-test",
// "avatarUrls":{"48x48":"https://api.atlassian.com/ex/jira/{cloudid}/rest/api/3/universal_avatar/view/type/project/avatar/10418","24x24":"https://api.atlassian.com/ex/jira/e71575ee-5a19-4b70-812d-5ae551967ad8/rest/api/3/universal_avatar/view/type/project/avatar/10418?size=small","16x16":"https://api.atlassian.com/ex/jira/e71575ee-5a19-4b70-812d-5ae551967ad8/rest/api/3/universal_avatar/view/type/project/avatar/10418?size=xsmall","32x32":"https://api.atlassian.com/ex/jira/e71575ee-5a19-4b70-812d-5ae551967ad8/rest/api/3/universal_avatar/view/type/project/avatar/10418?size=medium"},"projectTypeKey":"software","simplified":true,"style":"next-gen","isPrivate":false,"properties":{},"entityId":"df5152fe-c1d5-4da4-b60c-8a2f2b02a579","uuid":"df5152fe-c1d5-4da4-b60c-8a2f2b02a579"}]
