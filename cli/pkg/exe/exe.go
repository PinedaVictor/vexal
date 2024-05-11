// Package exe handles executing commands
package exe

import "os/exec"

func CMD(cmd string) {

}

// OpenURL uses he command open to open a URL
func OpenURL(url string) {
	cmd := exec.Command("open", url)
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}
