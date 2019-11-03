package action

//go:generate go run ../../../../gen/genoai/main.go --input ./specs/action.json --name action --package action --output action-cmds.go --tpl ./gen/tpl/gen-cmds.tpl --imports github.com/spf13/cobra
//go:generate goimports -w action-cmds.go

import (
	"github.com/spf13/cobra"
)

// Cmd -- used to connection to rootCmd
func Cmd() *cobra.Command {
	return actionCmd
}

var actionCmd = &cobra.Command{
	Use:   "action",
	Short: "action service",
}
