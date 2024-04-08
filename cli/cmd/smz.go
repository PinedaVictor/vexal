/*
Copyright © 2024 Victor Pineda
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// smzCmd represents the smz command
var smzCmd = &cobra.Command{
	Use:   "smz",
	Short: "Use smz to summarize files or dirictories.",
	Long: `smz generates readme.md files based on a file or directory.
	You can use this to explain code you need to update or have written.`,
	PreRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("In the PreRun stage")
		fmt.Println(args)
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("smz called", args)
	},
}

func init() {
	rootCmd.AddCommand(smzCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// smzCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// smzCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
