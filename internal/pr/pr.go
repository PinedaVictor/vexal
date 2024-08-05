package pr

import (
	"fmt"
	"os"
	"strconv"
	"vx/pkg"
	"vx/pkg/exe"
	"vx/pkg/gutils"
	"vx/pkg/paths"

	"github.com/google/go-github/v62/github"
)

// github package: "github.com/google/go-github/v62/github"

func AutoPr(branch string) {
	initGithubClient()
	workingBranch := gutils.GetWorkingBranch()
	commitTotals := gutils.CalcNumCommit(branch, workingBranch)
	ct := strconv.Itoa(commitTotals)
	owner, repo, _ := gutils.GetRepo()
	logs := gutils.GetGitLogs(workingBranch, ct)
	hasTpl, tpl := hasPRTemplate()
	var prBody string

	if hasTpl {
		prBody = pkg.GenerateReponse(fmt.Sprintf("Use the following commit messages and PR template %s to summaraize development, use bullet points as well. Each commit log is sperated by a | %s", tpl, logs))
	} else {
		prBody = pkg.GenerateReponse(fmt.Sprintf("Use the following commit messages to summaraize development, use bullet points as well. Each commit log is sperated by a | %s", logs))
	}

	maintainerCanModify := false
	// draft := false
	// issue := 0
	pullReq := &github.NewPullRequest{
		Title:               &workingBranch,
		Head:                github.String(workingBranch),
		HeadRepo:            github.String(workingBranch),
		Base:                github.String(branch),
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

func hasPRTemplate() (bool, string) {
	dir, _ := os.Getwd()
	templateDir := fmt.Sprintf("%s/.github/pull_request_template.md", dir)
	tplExists := paths.PathExists(templateDir)
	if tplExists {
		templateContent := paths.GetContent(templateDir)
		return true, templateContent
	}
	return false, ""
}
