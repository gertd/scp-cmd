package action

//go:generate scpgen gen-cmd --name action --package action --output action-gen.go | gofmt

import (
	"github.com/spf13/cobra"
)

// Cmd -- used to connection to rootCmd
func Cmd() *cobra.Command {
	return actionCmd
}

// catalogCmd represents the catalog command
var actionCmd = &cobra.Command{
	Use:   "action",
	Short: "action service",
}
