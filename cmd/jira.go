/*
Copyright © 2024 Victor Pineda pinedavictor095@gmail.com
*/
package cmd

import (
	"fmt"
	"vx/config"
	"vx/internal"
	"vx/internal/authenticate"

	"github.com/spf13/cobra"
)

// jiraCmd represents the features integratd into the Jira Platform https://www.atlassian.com/software/jira
var jiraCmd = &cobra.Command{
	Use:   "jira",
	Short: "Jira utils",
}

// jiraLogin initiates Jira OAuth2.0 login flow
var jiraLogin = &cobra.Command{
	Use:   "login",
	Short: "Login into your Jira board",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Running Jira login")
		authenticate.JiraLogin()
	},
}

// initJira appends to Jira configuration keys to repository config
var initJira = &cobra.Command{
	Use:   "init",
	Short: "Initialize Jira connection for your epository",
	Run: func(cmd *cobra.Command, args []string) {
		config.InitJira()
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		internal.PostFeedback("Configuration updated successfully. ✅")
	},
}

func init() {
	rootCmd.AddCommand(jiraCmd)
	jiraCmd.AddCommand(jiraLogin)
	jiraCmd.AddCommand(initJira)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// jiraCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// jiraCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
