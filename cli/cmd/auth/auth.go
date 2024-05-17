package auth

import (
	"vx/internal/authenticate"
	"vx/internal/commands"

	"github.com/spf13/cobra"
)

// loginCmd represents the login command
var AuthCmd = &cobra.Command{
	Use:   "auth",
	Short: "Authenticate into the vexal platform",
	Long:  ``,
}

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Use login to authenticate into the vexal platform",
	Long: `We use OAuth2.0 to authenticate and secure your data. Keeping your 
	tokens secure is important to us.`,
	PreRun: func(cmd *cobra.Command, args []string) {
		commands.StartSpinner("Authenticating:")
	},
	Run: func(cmd *cobra.Command, args []string) {
		authenticate.Login()
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		commands.StopSpinner("Authenticated Successfully")
	},
}

func init() {
	AuthCmd.AddCommand(loginCmd)
}
