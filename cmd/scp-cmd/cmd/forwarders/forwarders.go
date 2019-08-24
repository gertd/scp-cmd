package forwarders

//go:generate scpgen gen-cmd --name forwarders --package forwarders --output forwarders-gen.go | gofmt

import (
	"github.com/spf13/cobra"
)

// Cmd -- used to connection to rootCmd
func Cmd() *cobra.Command {
	return forwardersCmd
}

var forwardersCmd = &cobra.Command{
	Use:   "forwarders",
	Short: "forwarders service",
}
