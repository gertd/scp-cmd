package forwarders

//go:generate go run ../../../../gen/genoai/main.go --input ./specs/forwarders.json --name forwarders --package forwarders --output forwarders-cmds.go --tpl ./gen/tpl/gen-cmds.tpl --imports github.com/spf13/cobra
//go:generate gofmt -w forwarders-cmds.go

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
