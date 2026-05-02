/*
Copyright © 2026 Victor Pineda pinedavictor095@gmail.com
*/
package cmd

import (
	"fmt"
	"vx/internal"

	"github.com/spf13/cobra"
)

var depsCmd = &cobra.Command{
	Use:   "deps <file>",
	Short: "Show what a file imports and what depends on it",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		file := args[0]
		g, close, ok := loadGraph()
		if !ok {
			return
		}
		defer close()

		out, err := g.Describe(file)
		if err != nil {
			internal.UserErrFeedback(fmt.Sprintf("deps: %v", err))
			return
		}
		fmt.Print(out)
	},
}

func init() {
	rootCmd.AddCommand(depsCmd)
}
