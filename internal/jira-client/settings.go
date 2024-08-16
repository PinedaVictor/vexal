package jiraclient

import (
	"fmt"
	"vx/config"
)

// --url 'https://your-domain.atlassian.net/rest/api/2/issue' \
// --user 'email@example.com:<api_token>' \
// --header 'Accept: application/json' \
// --header 'Content-Type: application/json' \
func CreateIssue() {
	// cfg, cfgErr := config.LoadJiraAuthCfg()
	// if cfgErr != nil {
	// 	log.Println("error reading jira auth:", cfgErr)
	// }
	repoCfg, _ := config.LoadRepoConfig()
	// fmt.Println("Auth config:", cfg)
	fmt.Println("Repo config:", repoCfg)
	fmt.Println("Jira Repo config:", repoCfg)
	url := fmt.Sprintf("%s/rest/api/2/issue", repoCfg.Jira_URL)
	fmt.Println("Updating: ", url)

	// OAuthToken := fmt.Sprintf("Bearer %s", cfg.JiraToken)
	// // Define the headers in a map
	// headers := map[string]string{
	// 	"Accept":        "application/json",
	// 	"Content-Type":  "application/json",
	// 	"Authorization": OAuthToken,
	// }
	// // Define the payload (JSON data)
	// payload := map[string]interface{}{
	// 	"fields": map[string]interface{}{
	// 		"summary": "Main order flow broken",
	// 		"issuetype": map[string]string{
	// 			"id": "10000", // Replace with the actual ID for your issue type
	// 		},
	// 		"project": map[string]string{
	// 			"id": "SCRUM", // Replace with the actual ID for your project
	// 		},
	// 		"description": "Order entry fails when selecting supplier.",
	// 	},
	// }
	// resp, err := pkg.PostRequest(url, headers, payload)
	// if err != nil {
	// 	log.Println("error creating issue:", err)
	// }
	// fmt.Println(resp)
}
