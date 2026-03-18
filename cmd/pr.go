/*
Copyright © 2024 Victor Pineda pinedavictor095@gmail.com
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"vx/internal"
	"vx/internal/pr"

	"github.com/spf13/cobra"
)

var branch = ""
var verbatimNotes bool

// prCmd represents the pr command
var prCmd = &cobra.Command{
	Use:   "pr",
	Short: "AI-assisted generated PRs based on your commit history",
	Long: `
	Generate AI-assisted pull requests based on your commit history. 
	Vexal analyzes your Git commits—not your code—to summarize changes and prepare a 
	ready-to-review PR. It does not inspect, interpret, or modify your source code in any way.
	`,
	PreRun: func(cmd *cobra.Command, args []string) {
		user := pr.GetGitUser()
		if user == nil {
			fmt.Println("Unable to retrieve your GitHub user.")
			fmt.Println("Run 'vx init' to configure this repo, or 'vx context add <name>' to set up a global context.")
			os.Exit(0)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		notes := ""
		if verbatimNotes {
			fmt.Println("Enter verbatim notes (leave a blank line when done):")
			scanner := bufio.NewScanner(os.Stdin)
			var lines []string
			for scanner.Scan() {
				line := scanner.Text()
				if line == "" {
					break
				}
				lines = append(lines, line)
			}
			notes = strings.Join(lines, "\n")
		}
		internal.StartSpinner("Checking repository status and preparing pull request... ")
		pr.AutoPr(branch, notes)
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		internal.StopSpinner("PR complete!")
	},
}

func init() {
	rootCmd.AddCommand(prCmd)
	prCmd.Flags().StringVarP(&branch, "branch", "b", "main", "PR to an existing remote branch. Default is main")
	prCmd.Flags().BoolVarP(&verbatimNotes, "verbatim-notes", "n", false, "Prompt for verbatim notes to prepend to the PR body")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// prCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// prCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
