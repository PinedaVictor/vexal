package pr

import (
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

func GetGitLogs(branch string) string {
	execmd := exec.Command(GitCMD, "log", `--pretty=format:%s`, "-n", "20", branch)
	output, err := execmd.Output()
	if err != nil {
		fmt.Println("err running cmd:", err)
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
