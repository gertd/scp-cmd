package appreg

//go:generate scpgen gen-cmd --name app-registry --package appreg --output appreg-gen.go | gofmt

import (
	"github.com/spf13/cobra"
)

// Cmd -- used to connection to rootCmd
func Cmd() *cobra.Command {
	return appregCmd
}

// catalogCmd represents the catalog command
var appregCmd = &cobra.Command{
	Use:   "appreg",
	Short: "application registry service",
}
