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
	Run: func(cmd *cobra.Command, args []string) {
		internal.StartSpinner("Scanning for TODO comments ")
		scraper.ScrapeTodos()
		internal.StopSpinner("Results saved to todos.md")
	},
}

var clearTodos = &cobra.Command{
	Use:   "clear",
	Short: "Remove all TODO comments from your codebase.",
	Run: func(cmd *cobra.Command, args []string) {
		internal.StartSpinner("Clearing TODO comments ")
		scraper.ClearTodos()
		internal.StopSpinner("All TODO comments removed")
	},
}

func init() {
	rootCmd.AddCommand(todosCmd)
	todosCmd.AddCommand(clearTodos)
}
