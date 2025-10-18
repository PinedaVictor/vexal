/*
Copyright © 2024 Victor Pineda pinedavictor095@gmail.com
*/
package cmd

import (
	"fmt"
	"os"
	"vx/config"
	"vx/internal"
	"vx/internal/authenticate"
	"vx/internal/pr"

	"github.com/spf13/cobra"
)

var branch = ""

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
		if !config.RepoModeActive() {
			fmt.Println("Repository not initialized. Please run 'vx init' in your project root directory to create a ./vx.yaml configuration.")
			authenticate.RequireAuth()
		}
		user := pr.GetGitUser()
		if user == nil {
			fmt.Println("Unable to retrieve your GitHub user.")
			fmt.Println("Please ensure your GitHub access token is set in your .vx.yaml configuration file.")
			fmt.Println("Learn more about GitHub access tokens:")
			fmt.Printf("https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/managing-your-personal-access-tokens\n")
			os.Exit(0)
			// TODO: Marker the use case for GCP secrets manager
			// fmt.Printf("Make sure you enable the github api and supply vexal with a github token.\n")
			// fmt.Printf("command: vx config set -k github -v [ACCESS-TOKEN]\n")
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		internal.StartSpinner("Checking repository status and preparing pull request... ")
		pr.AutoPr(branch)
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		internal.StopSpinner("PR complete!")
	},
}

func init() {
	rootCmd.AddCommand(prCmd)
	prCmd.Flags().StringVarP(&branch, "branch", "b", "main", "PR to an existing remote branch. Default is main")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// prCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// prCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
