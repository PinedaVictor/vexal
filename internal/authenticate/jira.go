package authenticate

import (
	"fmt"
	"net/url"
	"vx/config"
	"vx/pkg/exe"
)

func JiraLogin() {
	env, _ := config.LoadEnvironment()
	clientID := env.JIRA_OAUTH_CLIENT_ID // Replace with your actual client ID
	redirectURI := "https://api.vexal.io/jira/callback"
	state := env.JIRA_STATE // Replace with a real state value
	scope := "read:jira-work write:jira-work manage:jira-project read:jira-user manage:jira-configuration read:me report:personal-data"

	// URL encode the parameters
	encodedRedirectURI := url.QueryEscape(redirectURI)
	encodedState := url.QueryEscape(state)
	encodedScope := url.QueryEscape(scope)

	// Construct the final URL
	authorizationURL := fmt.Sprintf(
		"https://auth.atlassian.com/authorize?audience=api.atlassian.com&client_id=%s&redirect_uri=%s&state=%s&scope=%s&response_type=code&prompt=consent",
		clientID, encodedRedirectURI, encodedState, encodedScope)

	fmt.Println("THE URL:", authorizationURL)
	exe.OpenURL(authorizationURL)
}

// Step to Authorize - which will be handle be the server
// curl --request POST \
//   --url 'https://auth.atlassian.com/oauth/token' \
//   --header 'Content-Type: application/json' \
//   --data '{"grant_type": "authorization_code","client_id": "CLIENT_ID","client_secret": "YOUR_CLIENT_SECRET","code": "ATOAjrE3vulLNTrG1XEjWb-ioPZKMFB8yA6ognBBA_7ie5gg2xW70TxESY_9vsHT39w8040F258F","redirect_uri": "https://api.vexal.io/callback"}'
