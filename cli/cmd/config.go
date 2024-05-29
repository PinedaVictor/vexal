/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"vx/config"
	"vx/internal/commands"
	"vx/internal/secrets"

	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "View and edit vx configuartion",
	Long: `You can use config to access and edit vx configuration.
	Configuration located at $HOME/.vx`,
}

// TODO: This will likely be deleted
// var view = &cobra.Command{
// 	Use:   "view",
// 	Short: "view config.json from $HOME/.vx",
// 	Run: func(cmd *cobra.Command, args []string) {
// 		// config.LoadConfig()
// 	},
// }

var (
	key   = ""
	value = ""
)

var set = &cobra.Command{
	Use:   "set [flags] [args]",
	Short: "set key value pair for vx configuration",
	PreRun: func(cmd *cobra.Command, args []string) {
		apis, validKey := config.CheckSupportedAPI(key)
		if !validKey {
			fmt.Printf("Invalid Key: %s \n", key)
			fmt.Printf("Supported APIs: %s \n", apis)
			os.Exit(0)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		commands.StartSpinner("Updating config")
		secrets.AddSecret(key, value)
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		msg := fmt.Sprintf("%s config updated successfully", key)
		commands.StopSpinner(msg)
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
	// TODO: view cmd will likely not be needed
	// configCmd.AddCommand(view)
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
