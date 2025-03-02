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

	// TODO: Refactor config cmd

	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Edit vx configuartion",
	Long: `You can use config to access and edit vx configuration.
	Configuration located at $HOME/.vx`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Running config")
		// TODO: Setup a test for this integration in prod

	},
}

var (
	key   = ""
	value = ""
)

var set = &cobra.Command{
	Use:   "set [flags] [args]",
	Short: "set key value pair for vx configuration",
	PreRun: func(cmd *cobra.Command, args []string) {
		authenticate.RequireAuth()
		apis, validKey := config.CheckSupportedAPI(key)
		if !validKey {
			fmt.Printf("Invalid Key: %s \n", key)
			fmt.Printf("Supported APIs: %s \n", apis)
			os.Exit(0)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		internal.StartSpinner("Updating config")
		if !secrets.AddSecret(key, value) {
			msg := fmt.Sprintf("Make sure you have enabled the API: %s", key)
			internal.StopSpinner(msg)
			os.Exit(0)
		}
		msg := fmt.Sprintf("%s config updated successfully", key)
		internal.StopSpinner(msg)
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.AddCommand(set)
	set.Flags().StringVarP(&key, "key", "k", "", "Define key to to be updated")
	set.Flags().StringVarP(&value, "value", "v", "", "Value of key")
	set.MarkFlagsRequiredTogether("key", "value")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
