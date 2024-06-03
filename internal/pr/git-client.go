package pr

import (
	"context"
	"fmt"
	"vx/config"
	"vx/internal/secrets"

	"github.com/google/go-github/v62/github"
)

var (
	gitCtx    context.Context
	gitClient *github.Client
)

func initGithubClient() {
	user, _ := config.LoadAuth()
	secretName := fmt.Sprintf("%s_github", user.UID)
	gitToken := secrets.GetSecretVersion(secretName)
	// github client
	gitCtx = context.Background()
	gitClient = github.NewClient(nil).WithAuthToken(gitToken)
}

func GetGitUser() *github.User {
	initGithubClient()
	user, resp, err := gitClient.Users.Get(gitCtx, "")
	defer resp.Body.Close()
	if err != nil {
		fmt.Printf("\nerror: %v\n", err)
		return nil
	}
	fmt.Println("user:", user)
	return user
}
