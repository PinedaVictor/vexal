/*
Copyright © 2025 Victor Pineda pinedavictor095@gmail.com
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// surgeCmd represents the surge command
var surgeCmd = &cobra.Command{
	Use:   "surge",
	Short: "Server generation",
	Long:  `Server generation for different libraries and frameworks.`,
}

func init() {
	rootCmd.AddCommand(surgeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// surgeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// surgeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
