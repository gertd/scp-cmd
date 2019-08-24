package login

import (
	impl "gitlab.com/d5s/scp-cmd/pkg/login"
	"github.com/spf13/cobra"
)

// Cmd -- used to connection to rootCmd
func Cmd() *cobra.Command {
	return loginCmd
}

// loginCmd represents the catalog command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "login service",
	RunE:  impl.Login,
}

func init() {

	loginCmd.Flags().StringP("uid", "u", "", "user name")
	loginCmd.Flags().StringP("pwd", "p", "", "password")
}
