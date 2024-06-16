package pr

import (
	"fmt"
	"os"
	"vx/pkg"
	"vx/pkg/exe"

	"github.com/google/go-github/v62/github"
)

func AutoPr(branch string) {
	initGithubClient()
	owner, repo, _ := GetRepo()
	workingBranch := GetWorkingBranch()
	logs := GetGitLogs(workingBranch)
	prBody := pkg.GenerateReponse(fmt.Sprintf("Use the following commit messages to summaraize development, use bullet points as well. Each commit log is sperated by a | %s", logs))
	// TODO: Input and/default for base branch
	// base := "main"
	maintainerCanModify := false
	// draft := false
	// issue := 0
	pullReq := &github.NewPullRequest{
		Title:    &workingBranch,
		Head:     github.String(workingBranch),
		HeadRepo: github.String(workingBranch),
		Base:     github.String(branch),
		// Base:                &base,
		Body:                &prBody,
		MaintainerCanModify: &maintainerCanModify,
		// Draft:               &draft,
		// Issue: &issue,
	}
	pullRequest, resp, err := gitClient.PullRequests.Create(gitCtx, owner, repo, pullReq)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println("error creating pr: \n", err)
		os.Exit(0)
	}
	url := pullRequest.HTMLURL
	exe.OpenURL(*url)
}
