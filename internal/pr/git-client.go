package pr

import (
	"context"
	"fmt"
	"vx/config"
	"vx/internal/secrets"

	"github.com/google/go-github/v62/github"
)

// const githubURL = "https://github.com"

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

// func AutoPr() {
// 	initGithubClient()
// 	// TODO: ALG:
// 	owner, repo, _ := GetRepo()
// 	workingBranch := GetWorkingBranch()
// 	fmt.Println("Working branch:", workingBranch)

// 	// TODO: Get pull request commit messages - This will be it's own function where
// 	logs := GetGitLogs()
// 	prBody := pkg.GenerateReponse(fmt.Sprintf("Use the following commit messages to summaraize development, use bullet points as well. Each commit log is sperated by a | %s", logs))
// 	fmt.Println(prBody)
// 	// TODO: Input and/default for base branch
// 	base := "main"
// 	maintainerCanModify := false
// 	draft := false
// 	// // issue := 0
// 	pullReq := &github.NewPullRequest{
// 		Title:               &workingBranch,
// 		Head:                &workingBranch,
// 		HeadRepo:            &workingBranch,
// 		Base:                &base,
// 		Body:                &prBody,
// 		MaintainerCanModify: &maintainerCanModify,
// 		Draft:               &draft,
// 		// Issue:               &issue,
// 	}
// 	// pullRequest, resp, err := gitClient.PullRequests.Create(gitCtx, "ServerGalaxy", "origins", pullReq)
// 	pullRequest, resp, err := gitClient.PullRequests.Create(gitCtx, owner, repo, pullReq)
// 	fmt.Println(pullRequest, resp, err)

// 	defer resp.Body.Close()
// 	if err != nil {
// 		fmt.Println("error creating pr", err)
// 	}
// 	fmt.Println("RESP")
// 	fmt.Println(resp)
// 	fmt.Println("-----------")
// 	fmt.Println("PULL REQUEST")
// 	fmt.Println(pullRequest)
// 	url := pullRequest.HTMLURL
// 	exe.OpenURL(*url)
// }
