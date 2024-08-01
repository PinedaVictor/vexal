package gutils

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

const (
	GitCMD    = "git"
	githubURL = "https://github.com"
)

// GetRepo returns local repositroy information: owner, repo, repo URL
func GetRepo() (string, string, string) {
	// TODO: Implement working directory switching if needed
	// curDir, _ := os.Getwd()
	// fmt.Println("Current working directory:", curDir)
	execmd := exec.Command(GitCMD, "config", "--get", "remote.origin.url")
	ouput, err := execmd.Output()
	if err != nil {
		fmt.Println("err running cmd:", err)
	}
	// Split the string on multiple delimiters
	splitFunc := func(r rune) bool {
		return r == ':' || r == '.' || r == '/'
	}
	splitStr := strings.FieldsFunc(string(ouput), splitFunc)
	owner := splitStr[len(splitStr)-3]
	repo := splitStr[len(splitStr)-2]
	return owner, repo, fmt.Sprintf("%s/%s/%s", githubURL, owner, repo)
}

func GetGitLogs(branch string, commitTotals string) string {
	execmd := exec.Command(GitCMD, "log", `--pretty=format:%s`, "-n", commitTotals, branch)
	output, err := execmd.Output()
	if err != nil {
		fmt.Println("err running cmd for logs:", err)
	}
	return strings.ReplaceAll(string(output), "\n", " | ")
}

func GetWorkingBranch() string {
	// branch --show-current
	execmd := exec.Command(GitCMD, "branch", "--show-current")
	output, err := execmd.Output()
	if err != nil {
		fmt.Println("err running cmd:", err)
		return ""
	}
	// Convert to string and trim whitespace
	branch := strings.TrimSpace(string(output))
	return branch
}

// Number of commits
// git rev-list BRANCH
func CalcNumCommit(base string, feature string) int {
	totalCommits := getNumCommits(feature) - getNumCommits(base)
	fmt.Printf("You're killing it! 🔥 Calculated %d total changes.\n", totalCommits)
	return totalCommits
}

func getNumCommits(branch string) int {
	var stderr bytes.Buffer
	execmd := exec.Command(GitCMD, "rev-list", branch)
	execmd.Stderr = &stderr
	output, err := execmd.Output()
	if err != nil {
		fmt.Println("err running cmd:", err)
		fmt.Println("Error running command:", err)
		fmt.Println("Standard error output:", stderr.String())
		return 0
	}
	// Count the number of lines in the output
	lines := bytes.Count(output, []byte{'\n'})
	fmt.Printf("\nBranch '%s' has %d commits\n", branch, lines)
	return lines
}
