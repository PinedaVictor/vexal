/*
Copyright © 2024 Victor Pineda
*/
package cmd

import (
	"fmt"
	"os"
	"vx/pkg/directories"

	"github.com/spf13/cobra"
)

// smzCmd represents the smz command
var smzCmd = &cobra.Command{
	Use: "smz [file | directory]",

	// DisableFlagsInUseLine: true,
	Short: "Use smz to summarize files or dirictories.",
	Long: `smz generates readme.md files based on a file or directory.
	You can use smz to explain code you need to update or have written.`,
	PreRun: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help() // Display help text
			os.Exit(0)
		}
		directories.DeterminePath(args[0])
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("smz called", args)
	},
}

func init() {
	rootCmd.AddCommand(smzCmd)
	// smzCmd.SetUsageTemplate("sfsdfg")
	// smzCmd.SetUsageTemplate("Usage: vx smz [file | directory]\n")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// smzCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// smzCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
