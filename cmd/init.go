/*
Copyright © 2024 Victor Pineda pinedavictor095@gmail.com
*/
package cmd

import (
	"vx/config"
	"vx/internal"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize repository utilities. (Only needed if you plan on using github and OpenAI)",
	Long:  `Run vx init in the root of your project to initialize repository utilities.`,
	Run: func(cmd *cobra.Command, args []string) {
		config.InitRepoMode()
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		internal.PostFeedback("Configuration initialized successfully. ./vx.yaml ✅")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
