/*
Copyright © 2024 Victor Pineda pinedavictor095@gmail.com
*/
package cmd

import (
	"fmt"
	"os"
	"vx/internal"
	"vx/internal/authenticate"
	"vx/internal/pr"

	"github.com/spf13/cobra"
)

// prCmd represents the pr command
var prCmd = &cobra.Command{
	Use:   "pr",
	Short: "AI generated PRs based on your commit history",
	// TODO: Add long description
	Long: ``,
	PreRun: func(cmd *cobra.Command, args []string) {
		authenticate.RequireAuth()
		user := pr.GetGitUser()
		if user == nil {
			fmt.Printf("We could not get your github user.\n")
			fmt.Printf("Make sure you enable the github api and supply vexal with a github token.\n")
			fmt.Printf("command: vx config set -k github -v [ACCESS-TOKEN]\n")
			fmt.Printf("Github documentation on access tokens\n")
			fmt.Printf("https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/managing-your-personal-access-tokens\n")
			os.Exit(0)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		internal.StartSpinner("Preparing your PR ")
		pr.AutoPr()
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		internal.StopSpinner("PR complete!")
	},
}

func init() {
	rootCmd.AddCommand(prCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// prCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// prCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
