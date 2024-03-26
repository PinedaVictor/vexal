/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"vx/config"

	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "View and edit vx configuartion",
	Long: `You can use config to access and edit vx configuration.
	Configuration located at $HOME/.vx`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("config called")
	},
}

var view = &cobra.Command{
	Use:   "view",
	Short: "view config.json from $HOME/.vx",
	Run: func(cmd *cobra.Command, args []string) {
		config.LoadConfig()
	},
}

// var (
// 	key   = ""
// 	value = ""
// )

// var set = &cobra.Command{
// 	Use:   "set [flags] [args]",
// 	Short: "set key value pairs for sid configuration",
// 	Run: func(cmd *cobra.Command, args []string) {
// 		validKey := tools.KeyCheck(key)
// 		if validKey {
// 			viper.Set(key, value)
// 			viper.WriteConfig()
// 			log.Println("Config updated successfully")
// 		}
// 	},
// }

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.AddCommand((view))

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
