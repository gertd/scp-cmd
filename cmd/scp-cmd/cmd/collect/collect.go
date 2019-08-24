package collect

//go:generate scpgen gen-cmd --name collect --package collect --output collect-gen.go | gofmt

import (
	"github.com/spf13/cobra"
)

// Cmd -- used to connection to rootCmd
func Cmd() *cobra.Command {
	return collectCmd
}

// collectCmd represents the catalog command
var collectCmd = &cobra.Command{
	Use:   "collect",
	Short: "collect service",
}
