package identity

//go:generate scpgen gen-cmd --name identity --package identity --output identity-gen.go | gofmt

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
