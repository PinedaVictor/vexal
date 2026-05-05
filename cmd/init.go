/*
Copyright © 2024 Victor Pineda pinedavictor095@gmail.com
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"vx/config"
	"vx/internal"

	"github.com/PinedaVictor/nyx/pkg/depgraph"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize repository utilities. (Only needed if you plan on using github and OpenAI)",
	Long:  `Run vx init in the root of your project to initialize repository utilities.`,
	Run: func(cmd *cobra.Command, args []string) {
		config.InitRepoMode()
		n := buildDepGraph()
		if n > 0 {
			if err := config.InitClaudeMD(n); err != nil {
				internal.UserErrFeedback(fmt.Sprintf("claude.md: %v", err))
			}
		}
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		// TODO: Revisit at end of workflow
		internal.PostFeedback("Directory .vexal success")
		internal.PostFeedback("File ./vx.yaml succesa")
		internal.PostFeedback("NOTE: You can safely delete the ./vx.yaml file if vx context is configured.")
	},
}

func buildDepGraph() int {
	curDir, _ := os.Getwd()
	snapshotPath := filepath.Join(curDir, ".vexal", "snapshot.arrow")

	internal.StartSpinner("Building dependency graph ")
	n, err := depgraph.Snapshot(curDir, snapshotPath)
	if err != nil {
		internal.StopSpinner("")
		internal.UserErrFeedback(fmt.Sprintf("dependency graph: %v", err))
		return 0
	}
	internal.StopSpinner(fmt.Sprintf("Dependency graph ready — %d edges indexed", n))
	return n
}

func init() {
	rootCmd.AddCommand(initCmd)
}
