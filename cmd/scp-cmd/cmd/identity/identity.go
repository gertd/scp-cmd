package identity

//go:generate go run ../../../../gen/genoai/main.go --input ./specs/identity.json --name identity --package identity --output identity-cmds.go --tpl ./gen/tpl/gen-cmds.tpl --imports github.com/spf13/cobra
//go:generate gofmt -w identity-cmds.go

import (
	"github.com/spf13/cobra"
)

// Cmd -- used to connection to rootCmd
func Cmd() *cobra.Command {
	return identityCmd
}

var identityCmd = &cobra.Command{
	Use:   "identity",
	Short: "identity service",
}
