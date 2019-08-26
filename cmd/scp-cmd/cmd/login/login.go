package login

import (
	"github.com/spf13/cobra"
	impl "gitlab.com/d5s/scp-cmd/pkg/login"
)

// Cmd -- used to connection to rootCmd
func Cmd() *cobra.Command {
	return loginCmd
}

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "login service",
	RunE:  impl.Login,
}

func init() {

	loginCmd.Flags().StringP("uid", "u", "", "user name")
	loginCmd.Flags().StringP("pwd", "p", "", "password")
}
