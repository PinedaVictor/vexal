/*
Copyright © 2024 Victor Pineda pinedavictor095@gmail.com
*/
package cmd

import (
	"vx/internal"
	"vx/internal/scraper"

	"github.com/spf13/cobra"
)

// fixmeCmd represents the fixme command
var fixmeCmd = &cobra.Command{
	Use:   "fixme",
	Short: `Find all "FIXME:" comments in your codebase.`,
	PreRun: func(cmd *cobra.Command, args []string) {
		internal.PreFeedback(`Finding all "FIXME:" comments`)
	},
	Run: func(cmd *cobra.Command, args []string) {
		scraper.ScrapeFixMe()
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		internal.PostFeedback("All done! fixme.md ✅")
	},
}

func init() {
	rootCmd.AddCommand(fixmeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fixmeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fixmeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
