package pr

import (
	"fmt"
	"vx/pkg"
	"vx/pkg/exe"

	"github.com/google/go-github/v62/github"
)

func AutoPr() {
	initGithubClient()
	owner, repo, _ := GetRepo()
	fmt.Println(owner, repo)
	workingBranch := GetWorkingBranch()
	fmt.Println("working branch:", workingBranch)
	logs := GetGitLogs()
	prBody := pkg.GenerateReponse(fmt.Sprintf("Use the following commit messages to summaraize development, use bullet points as well. Each commit log is sperated by a | %s", logs))
	// TODO: Input and/default for base branch
	base := "main"
	maintainerCanModify := false
	draft := false
	// // issue := 0
	pullReq := &github.NewPullRequest{
		Title:               &workingBranch,
		Head:                &workingBranch,
		HeadRepo:            &workingBranch,
		Base:                &base,
		Body:                &prBody,
		MaintainerCanModify: &maintainerCanModify,
		Draft:               &draft,
		// Issue:               &issue,
	}
	fmt.Println("Pull request data:", prBody)
	fmt.Println("Pull request data:", pullReq)
	// pullRequest, resp, err := gitClient.PullRequests.Create(gitCtx, "ServerGalaxy", "origins", pullReq)
	pullRequest, resp, err := gitClient.PullRequests.Create(gitCtx, owner, repo, pullReq)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println("error creating pr", err)
	}
	url := pullRequest.HTMLURL
	exe.OpenURL(*url)
}
