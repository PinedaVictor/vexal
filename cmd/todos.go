/*
Copyright © 2024 Victor Pineda pinedavictor095@gmail.com
*/
package cmd

import (
	"vx/internal"
	"vx/internal/scraper"

	"github.com/spf13/cobra"
)

var todosCmd = &cobra.Command{
	Use:   "todos",
	Short: `Find all "TODO:" comments in your codebase.`,
	PreRun: func(cmd *cobra.Command, args []string) {
		internal.PreFeedback(`Finding all "TODO:" comments`)
	},
	Run: func(cmd *cobra.Command, args []string) {
		scraper.ScrapeTodos()
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		internal.PostFeedback("All done. ✅")
	},
}

var clearTodos = &cobra.Command{
	Use:   "clear",
	Short: "",
	PreRun: func(cmd *cobra.Command, args []string) {
		internal.PreFeedback(`Clearing all TODO comments`)
	},
	Run: func(cmd *cobra.Command, args []string) {
		scraper.ClearTodos()
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		internal.PostFeedback("All done. ✅")
	},
}

func init() {
	rootCmd.AddCommand(todosCmd)
	todosCmd.AddCommand(clearTodos)
}
