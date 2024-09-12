/*
Copyright © 2024 Victor Pineda pinedavictor095@gmail.com
*/
package cmd

import (
	"fmt"
	"vx/config"
	"vx/internal"
	"vx/internal/authenticate"
	jiraclient "vx/internal/jira-client"

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

var manual bool

// initJira appends to Jira configuration keys to repository config
var initJira = &cobra.Command{
	Use:   "init",
	Short: "Initialize Jira connection for your repository",
	Run: func(cmd *cobra.Command, args []string) {
		if manual {
			config.InitJira()
		} else {
			authenticate.InitJiraWithAuth()
		}
		// TODO: Error Handling: 2024/08/15 23:53:49 Error reading repo config: Config File ".vx.yaml" Not Found in "[/Users/victorpineda/repos/vexal-technologies/vexal]"
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		internal.PostFeedback("Configuration updated successfully. ✅")
	},
}

var pushAll = &cobra.Command{
	Use:   "push-all",
	Short: "Push all comments to Jira - TODO, FIXME comments",
	Run: func(cmd *cobra.Command, args []string) {
		internal.UserFeedback("Pushing all comments")
		jiraclient.PushAllComments()
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		internal.PostFeedback("Pushed All Comments. ✅")
	},
}

var pushTodos = &cobra.Command{
	Use:   "push-t",
	Short: "Push TODO comments to Jira - by defualt this will also delete TODO comments in your codebase",
	Run: func(cmd *cobra.Command, args []string) {
		internal.UserFeedback("Pushing all TODO comments")
		jiraclient.Push("todo")
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		internal.PostFeedback("Pushed TOOD Comments. ✅")
	},
}

var pushFixme = &cobra.Command{
	Use:   "push-f",
	Short: "Push FIXME comments to Jira - by defualt this will also delete FIXME comments in your codebase",
	Run: func(cmd *cobra.Command, args []string) {
		internal.UserFeedback("Pushing all FIXME comments")
		jiraclient.Push("fixme")
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		internal.PostFeedback("Pushed FIXME Comments. ✅")
	},
}

var listPrjIssueTypes = &cobra.Command{
	Use:   "lt",
	Short: "List issue types for your project",
	Run: func(cmd *cobra.Command, args []string) {
		jiraclient.ListIssueTypes()
	},
}

func init() {
	// Add jira cmd to root
	rootCmd.AddCommand(jiraCmd)
	// Add jira subcommands
	jiraCmd.AddCommand(jiraLogin)
	jiraCmd.AddCommand(initJira)
	jiraCmd.AddCommand(pushTodos)
	jiraCmd.AddCommand(pushFixme)
	jiraCmd.AddCommand(pushAll)
	jiraCmd.AddCommand(listPrjIssueTypes)

	// Here you will define your flags and configuration settings.
	initJira.Flags().BoolVarP(&manual, "manual", "m", false, "Enable manual config setup")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// jiraCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// jiraCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
