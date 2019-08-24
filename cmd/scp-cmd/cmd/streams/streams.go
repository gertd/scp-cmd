package streams

//go:generate scpgen gen-cmd --name streams --package streams --output streams-gen.go | gofmt

import (
	"github.com/spf13/cobra"
)

// Cmd -- used to connection to rootCmd
func Cmd() *cobra.Command {
	return streamsCmd
}

// streamsCmd represents the search command
var streamsCmd = &cobra.Command{
	Use:   "streams",
	Short: "streams service",
}
