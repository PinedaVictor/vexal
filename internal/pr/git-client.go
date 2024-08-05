package pr

import (
	"context"
	"fmt"
	"log"
	"vx/config"
	"vx/internal/keys"
	"vx/internal/secrets"

	"github.com/google/go-github/v62/github"
)

var (
	gitCtx    context.Context
	gitClient *github.Client
)

func initGithubClient() {
	// TODO: We need a control scructure on where to get our keys
	rCfg, _ := keys.KeySrc()
	gitCtx = context.Background()
	gitClient = github.NewClient(nil).WithAuthToken(rCfg.Github_key)
}

// TODO: implement control logic
func initWithAuth() string {
	user, err := config.LoadAuth()
	if err != nil {
		log.Println("There is no auth user")
	}
	fmt.Println("Loading user with empty keys:", user)
	repoCfg, _ := config.LoadRepoConfig()
	fmt.Println("REPO CONFIG:", repoCfg)
	secretName := fmt.Sprintf("%s_github", user.UID)
	gitToken := secrets.GetSecretVersion(secretName)
	return gitToken
}

func GetGitUser() *github.User {
	initGithubClient()
	user, resp, err := gitClient.Users.Get(gitCtx, "")
	defer resp.Body.Close()
	if err != nil {
		fmt.Printf("\nerror: %v\n", err)
		return nil
	}
	return user
}
