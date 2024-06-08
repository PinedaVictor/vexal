package pr

import (
	"context"
	"fmt"
	"vx/config"
	"vx/internal/secrets"

	"github.com/google/go-github/v62/github"
)

const githubURL = "https://github.com"

var (
	gitCtx    context.Context
	gitClient *github.Client
)

func initGithubClient() {
	user, _ := config.LoadAuth()
	secretName := fmt.Sprintf("%s_github", user.UID)
	gitToken := secrets.GetSecretVersion(secretName)
	// github
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

func AutoPr() {
	initGithubClient()
	// TODO: ALG:
	// 1. git config --get remote.origin.url
	owner, repo := GetRepo()
	// 2. construct HTML URL
	url := fmt.Sprintf("%s/%s/%s", githubURL, owner, repo)
	fmt.Println("github repo url:", url)
	// 3. Set owner and repo name

	// TODO: Get pull request data
	// title := "ticket number"
	// head := "dev"
	// headRepo := "dev"
	// base := "main"
	// body := "This is a test was generated with vx pr"
	// maintainerCanModify := false
	// draft := false
	// // issue := 0
	// pullReq := &github.NewPullRequest{
	// 	Title:               &title,
	// 	Head:                &head,
	// 	HeadRepo:            &headRepo,
	// 	Base:                &base,
	// 	Body:                &body,
	// 	MaintainerCanModify: &maintainerCanModify,
	// 	Draft:               &draft,
	// 	// Issue:               &issue,
	// }
	// // https://github.com/ServerGalaxy/origins
	// pullRequest, resp, err := gitClient.PullRequests.Create(gitCtx, "ServerGalaxy", "origins", pullReq)
	// defer resp.Body.Close()
	// if err != nil {
	// 	fmt.Println("error creating pr", err)
	// }
	// fmt.Println("RESP")
	// fmt.Println(resp)
	// fmt.Println("-----------")
	// fmt.Println("PULL REQUEST")
	// fmt.Println(pullRequest)
	// url := pullRequest.HTMLURL
	// exe.OpenURL(*url)

}
