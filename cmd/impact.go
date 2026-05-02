/*
Copyright © 2026 Victor Pineda pinedavictor095@gmail.com
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"vx/internal"

	"github.com/PinedaVictor/nyx/pkg/depgraph"
	"github.com/spf13/cobra"
)

var impactCmd = &cobra.Command{
	Use:   "impact <file>",
	Short: "Show what files would be affected by changing the given file",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		file := args[0]
		g, close, ok := loadGraph()
		if !ok {
			return
		}
		defer close()

		callers, err := g.ImpactDetailed(file)
		if err != nil {
			internal.UserErrFeedback(fmt.Sprintf("impact: %v", err))
			return
		}

		if len(callers) == 0 {
			fmt.Printf("impact: %s\n\nno dependents found\n", file)
			return
		}

		fmt.Printf("impact: %s\n\n%d file(s) affected:\n", file, len(callers))
		for _, c := range callers {
			fmt.Printf("  %s:%d\n", c.File, c.Line)
		}
	},
}

// loadGraph opens the snapshot from the current repo's .vexal directory.
// Returns the graph, a close func, and whether it succeeded.
// Prints a stale warning if the snapshot is out of date but still loads.
func loadGraph() (*depgraph.Graph, func(), bool) {
	curDir, _ := os.Getwd()
	snapshotPath := filepath.Join(curDir, ".vexal", "snapshot.arrow")

	if _, err := os.Stat(snapshotPath); os.IsNotExist(err) {
		internal.UserErrFeedback("no dependency graph found — run 'vx init' first")
		return nil, nil, false
	}

	g, err := depgraph.Load(snapshotPath)
	if err != nil {
		internal.UserErrFeedback(fmt.Sprintf("load graph: %v", err))
		return nil, nil, false
	}

	if stale, err := g.Stale(curDir); err == nil && stale.IsStale {
		internal.UserFeedback(fmt.Sprintf("warning: snapshot is stale (%d file(s) changed) — run 'vx init' to rebuild", len(stale.Files)))
	}

	return g, func() { g.Close() }, true
}

func init() {
	rootCmd.AddCommand(impactCmd)
}
