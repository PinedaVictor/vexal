/*
Copyright © 2025 Victor Pineda pinedavictor095@gmail.com
*/
package cmd

import (
	"fmt"
	"vx/internal"

	pluto "github.com/PinedaVictor/pluto"
	"github.com/spf13/cobra"
)

// surgeCmd represents the surge command
var surgeCmd = &cobra.Command{
	Use:   "surge",
	Short: "Server generation",
	Long:  `Server generation for different libraries and frameworks.`,
}

var name = ""
var exjs = &cobra.Command{
	Use:   "ejs",
	Short: "Generate ExpressJS server",
	Long: `
	This command generates an ExpressJS server with: 
	Dependencies: express, ts-node, and nodemon
	Dev Dependencies: typescript, @types/express, prettier, gts
	Configurations: Docker, Prettier, ESLint, and both development and build scripts.
	`,
	PreRun: func(cmd *cobra.Command, args []string) {
		var msg = fmt.Sprintf("Spawning server: %s", name)
		internal.StartSpinner(msg)
	},
	Run: func(cmd *cobra.Command, args []string) {
		pluto.Ejs(name)

	},
	PostRun: func(cmd *cobra.Command, args []string) {
		internal.StopSpinner(internal.DftlDoneMsg)
	},
}

var basal = &cobra.Command{
	Use:   "basal",
	Short: "Generate Python FastAPI server",
	Long: `
	This command generates a Python FastAPI server with:
	Dependencies: fastapi, uvicorn, python-dotenv, gunicorn, flake8
	`,
	PreRun: func(cmd *cobra.Command, args []string) {
		var msg = fmt.Sprintf("Spawning server: %s", name)
		internal.StartSpinner(msg)
	},
	Run: func(cmd *cobra.Command, args []string) {
		pluto.Basal(name)

	},
	PostRun: func(cmd *cobra.Command, args []string) {
		internal.StopSpinner(internal.DftlDoneMsg)
	},
}

func init() {
	rootCmd.AddCommand(surgeCmd)
	surgeCmd.AddCommand(exjs)
	exjs.Flags().StringVarP(&name, "name", "n", "server", "Name of server you want to generate")
	surgeCmd.AddCommand(basal)
	basal.Flags().StringVarP(&name, "name", "n", "server", "Name of server you want to generate")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// surgeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// surgeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
