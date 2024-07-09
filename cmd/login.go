/*
Copyright © 2024 Victor Pineda pinedavictor095@gmail.com
*/package cmd

import (
	"vx/internal"
	"vx/internal/authenticate"

	"github.com/spf13/cobra"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Use login to authenticate into the vexal platform",
	Long: `We use OAuth2.0 to authenticate and secure your data. Keeping your 
	tokens secure is important to us.`,
	PreRun: func(cmd *cobra.Command, args []string) {
		internal.StartSpinner("Authenticating:")
	},
	Run: func(cmd *cobra.Command, args []string) {
		authenticate.Login()
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		internal.StopSpinner("Authenticated Successfully")
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
