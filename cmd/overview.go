/*
Copyright © 2026 Victor Pineda pinedavictor095@gmail.com
*/
package cmd

import (
	"fmt"
	"vx/internal"

	"github.com/spf13/cobra"
)

var overviewCmd = &cobra.Command{
	Use:   "overview",
	Short: "Show repo structure: top dependents, file count, edge count",
	Run: func(cmd *cobra.Command, args []string) {
		g, close, ok := loadGraph()
		if !ok {
			return
		}
		defer close()

		result, err := g.Overview(10)
		if err != nil {
			internal.UserErrFeedback(fmt.Sprintf("overview: %v", err))
			return
		}

		fmt.Printf("overview  |  %d files  |  %d edges\n", result.TotalFiles, result.TotalEdges)

		if len(result.TopDependents) == 0 {
			fmt.Println("\nno internal dependencies found")
			return
		}

		fmt.Println("\nmost depended-on:")
		for _, td := range result.TopDependents {
			fmt.Printf("  %-60s ← %d file(s)\n", td.ImportPath, td.Dependents)
		}
	},
}

func init() {
	rootCmd.AddCommand(overviewCmd)
}
