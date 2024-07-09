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
	"vx/internal/secrets"

	"github.com/spf13/cobra"
)

// enableCmd represents the enable command
var enableCmd = &cobra.Command{
	Use:   "enable",
	Short: "Enable supported API integrations by vx",
	Long: `
	vx allows users to integrate and enable third-party APIs,
	enhancing the capabilities of their applications. This tool supports various popular APIs,
	enabling easy configuration and management to streamline workflows and automate processes`,
	PreRun: func(cmd *cobra.Command, args []string) {
		authenticate.RequireAuth()
		if !config.APISupported(args[0]) {
			fmt.Printf("API: %s is not supported \n", args[0])
			config.PrintSupportedAPIs()
			os.Exit(0)
		}
		internal.StartSpinner("Enabling API")
	},
	Run: func(cmd *cobra.Command, args []string) {
		if secrets.CreateSecret(args[0]) {
			msg := fmt.Sprintf("API Enabled run: vx config set -k %s -v [YOUR API KEY]", args[0])
			internal.StopSpinner(msg)
		}
		internal.StopSpinner("Failed to enable API")
	},
}

func init() {
	rootCmd.AddCommand(enableCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// enableCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// enableCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
