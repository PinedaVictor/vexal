/*
Copyright © 2024 Victor Pineda pinedavictor095@gmail.com
*/
package cmd

import (
	"os"
	"regexp"
	"strings"
	"vx/cmd/auth"
	"vx/config"
	"vx/internal/authenticate"
	"vx/test"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "vx",
	Short: "AI Powered tools for Developers",
	Long:  ``,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	config.InitConfig()
	config.WriteLicense()
	test.TestHTTP()
	rootCmd.AddCommand(auth.AuthCmd)
	// Color output setup
	rootCmd.SetOutput(color.Output)
	cobra.AddTemplateFunc("StyleHeading", color.New(color.FgBlue, color.Underline, color.Bold).SprintFunc())
	usageTemplate := rootCmd.UsageTemplate()
	usageTemplate = strings.NewReplacer(
		`Usage:`, `{{StyleHeading "Usage:"}}`,
		`Aliases:`, `{{StyleHeading "Aliases:"}}`,
		`Available Commands:`, `{{StyleHeading "Commands:"}}`,
		`Global Flags:`, `{{StyleHeading "Global Flags:"}}`,
		// The following one steps on "Global Flags:"
		`Flags:`, `{{StyleHeading "Flags:"}}`,
	).Replace(usageTemplate)
	re := regexp.MustCompile(`(?m)^Flags:\s*$`)
	usageTemplate = re.ReplaceAllLiteralString(usageTemplate, `{{StyleHeading "Flags:"}}`)
	usageTemplate = usageTemplate + "Auth Status: " + authenticate.RootAuthStatus() + "\n"
	rootCmd.SetUsageTemplate(usageTemplate)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
