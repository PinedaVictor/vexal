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

func GetGitLogs() string {
	// branch
	// git rev-parse --abbrev-ref HEAD
	// commits to 15
	// git log --pretty=format:"%s" -n 15 origin/main
	// commits to branch
	// git log --pretty=format:"%s" origin/main
	// execmd := exec.Command(GitCMD, "log", `--pretty=format:%s | %ci`, "origin/main")
	execmd := exec.Command(GitCMD, "log", `--pretty=format:%s`, "origin/main")
	output, err := execmd.Output()
	if err != nil {
		fmt.Println("err running cmd:", err)
	}
	return strings.ReplaceAll(string(output), "\n", " | ")
}

func GetWorkingBranch() string {
	// branch --show-current
	execmd := exec.Command(GitCMD, "branch", "--show-current")
	ouput, err := execmd.Output()
	if err != nil {
		fmt.Println("err running cmd:", err)
	}
	return string(ouput)
}
