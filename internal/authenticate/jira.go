package authenticate

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/url"
	"vx/config"
	"vx/pkg"
	"vx/pkg/exe"
)

type JiraResource struct {
	ID        string   `json:"id"`
	URL       string   `json:"url"`
	Name      string   `json:"name"`
	Scopes    []string `json:"scopes"`
	AvatarURL string   `json:"avatarUrl"`
}

func InitJiraWithAuth() {
	JiraLogin()
	fmt.Println("All Done Logging into Jira")
	resources, _ := getJiraAccessibleResources()
	fmt.Println("Resoruces:", resources)
	config.UpdateJiraRepoCfg(resources.Name, resources.URL, resources.ID)
}

func JiraLogin() {
	env, _ := config.LoadEnvironment()
	clientID := env.JIRA_OAUTH_CLIENT_ID
	redirectURI := "https://api.vexal.io/jira/callback"
	state := env.JIRA_STATE
	scope := "read:jira-work write:jira-work manage:jira-project read:jira-user manage:jira-configuration read:me report:personal-data"

	// URL encode the parameters
	encodedRedirectURI := url.QueryEscape(redirectURI)
	encodedState := url.QueryEscape(state)
	encodedScope := url.QueryEscape(scope)

	// Construct the final URL
	authorizationURL := fmt.Sprintf(
		"https://auth.atlassian.com/authorize?audience=api.atlassian.com&client_id=%s&redirect_uri=%s&state=%s&scope=%s&response_type=code&prompt=consent",
		clientID, encodedRedirectURI, encodedState, encodedScope)

	exe.OpenURL(authorizationURL)
	RunAuthServer("jira")
}

func getJiraAccessibleResources() (JiraResource, error) {
	jiraAuthCfg, _ := config.LoadJiraAuthCfg()
	tokenHeader := fmt.Sprintf("Bearer %s", jiraAuthCfg.JiraToken)
	headers := map[string]string{
		"Authorization": tokenHeader,
		"Accept":        "application/json",
	}

	resp, err := pkg.GetRequest("https://api.atlassian.com/oauth/token/accessible-resources", headers)
	if err != nil {
		log.Println("error getting jira accessible resources")
	}
	defer resp.Body.Close()
	// Read the body of the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("error getting Jira Accessible Resources:", err)
		return JiraResource{}, err
	}

	// Unmarshal the JSON into a slice of JiraResource
	var resources []JiraResource
	err = json.Unmarshal(body, &resources)
	if err != nil {
		log.Println("error getting Jira Accessible Resources:", err)
		return JiraResource{}, err
	}

	return resources[0], nil
}
