/*
Copyright © 2025 Victor Pineda pinedavictor095@gmail.com
*/
package cmd

import (
	nyx "github.com/PinedaVictor/nyx"
	"github.com/spf13/cobra"
)

// betaCmd represents the beta command
var betaCmd = &cobra.Command{
	Use:   "beta",
	Short: "",
	Run: func(cmd *cobra.Command, args []string) {
		nyx.Test()
		nyx.RunTESTDiagnostics()

	},
}

func init() {
	rootCmd.AddCommand(betaCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// betaCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// betaCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
