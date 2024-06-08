package pr

import (
	"fmt"
	"os/exec"
	"strings"
)

const GitCMD = "git"

func GetRepo() (string, string) {
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
	// fmt.Println("Owner:", splitStr[len(splitStr)-3])
	// fmt.Println("Repo:", splitStr[len(splitStr)-2])
	return splitStr[len(splitStr)-3], splitStr[len(splitStr)-2]

}
