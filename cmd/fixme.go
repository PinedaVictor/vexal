/*
Copyright © 2024 Victor Pineda pinedavictor095@gmail.com
*/
package cmd

import (
	"vx/internal"
	"vx/internal/scraper"

	"github.com/spf13/cobra"
)

var fixmeCmd = &cobra.Command{
	Use:   "fixme",
	Short: `Find all "FIXME:" comments in your codebase.`,
	Run: func(cmd *cobra.Command, args []string) {
		internal.StartSpinner("Scanning for FIXME comments ")
		scraper.ScrapeFixMe()
		internal.StopSpinner("Results saved to fixme.md")
	},
}

func init() {
	rootCmd.AddCommand(fixmeCmd)
}
