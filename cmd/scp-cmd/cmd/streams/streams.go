package streams

//go:generate go run ../../../../gen/genoai/main.go --input ./specs/streams.json --name streams --package streams --output streams-cmds.go --tpl ./gen/tpl/gen-cmds.tpl --imports github.com/spf13/cobra
//go:generate goimports -w streams-cmds.go

import (
	"github.com/spf13/cobra"
)

// Cmd -- used to connection to rootCmd
func Cmd() *cobra.Command {
	return streamsCmd
}

var streamsCmd = &cobra.Command{
	Use:   "streams",
	Short: "streams service",
}
