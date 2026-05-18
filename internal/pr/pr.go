package pr

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"vx/internal"
	"vx/pkg"
	"vx/pkg/exe"
	"vx/pkg/gutils"
	"vx/pkg/paths"

	"github.com/google/go-github/v62/github"
)

var hypeMessages = []string{
	"You're killin it!",
	"Is that Spider-Man?!",
	"Look at you go!",
	"Sheesh!",
	"Absolute unit!",
	"Built different!",
	"Main character energy!",
	"We love to see it!",
	"Different breed!",
}

func AutoPr(branch string, verbatimNotes string) {
	initGithubClient()
	workingBranch := gutils.GetWorkingBranch()
	if branch == workingBranch {
		fmt.Printf("\nYou are already on '%s'.\n", branch)
		fmt.Println("Create a new branch before opening a pull request.")
		fmt.Println("Example:")
		fmt.Println("	git checkout -b your-branch-name")
		os.Exit(0)
		return
	}
	commitTotals := gutils.CalcNumCommit(branch, workingBranch)
	ct := strconv.Itoa(commitTotals)
	if commitTotals == 0 {
		fmt.Println("No commits detected in this branch.")
		fmt.Println("Make sure changes are pushed to your remote branch before running this command.")
		os.Exit(0)
		return
	}
	hype := hypeMessages[rand.Intn(len(hypeMessages))]
	fmt.Printf("%s %s total changes.\n", hype, ct)
	if !gutils.BranchExistsOnRemote(workingBranch) {
		fmt.Printf("\nBranch '%s' has not been pushed to remote.\n", workingBranch)
		fmt.Println("Please push your branch before opening a pull request:")
		fmt.Printf("  git push -u origin %s\n", workingBranch)
		os.Exit(0)
	}
	owner, repo, _ := gutils.GetRepo()
	logs := gutils.GetGitLogs(workingBranch, ct)
	hasTpl, tpl := hasPRTemplate()
	var prBody string

	internal.StartSpinner("Generating PR ")
	if hasTpl {
		prBody = pkg.GenerateReponse(fmt.Sprintf("Use the following commit messages and PR template %s to summaraize development, use bullet points as well. Each commit log is sperated by a | %s", tpl, logs))
	} else {
		prBody = pkg.GenerateReponse(fmt.Sprintf("Use the following commit messages to summaraize development, use bullet points as well. Each commit log is sperated by a | %s", logs))
	}
	internal.StopSpinner("PR complete!")

	if verbatimNotes != "" {
		prBody = verbatimNotes + "\n\n" + prBody
	}
	prBody = prBody + "\n\n---\n*Generated with [vexal](https://www.vexal.io)*"

	maintainerCanModify := false
	pullReq := &github.NewPullRequest{
		Title:               &workingBranch,
		Head:                github.String(workingBranch),
		Base:                github.String(branch),
		Body:                &prBody,
		MaintainerCanModify: &maintainerCanModify,
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
