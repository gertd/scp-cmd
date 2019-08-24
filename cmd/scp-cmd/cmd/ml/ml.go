package ml

//go:generate scpgen gen-cmd --name ml --package ml --output ml-gen.go | gofmt

import (
	"github.com/spf13/cobra"
)

// Cmd -- used to connection to rootCmd
func Cmd() *cobra.Command {
	return mlCmd
}

var mlCmd = &cobra.Command{
	Use:   "ml",
	Short: "ml service",
}
