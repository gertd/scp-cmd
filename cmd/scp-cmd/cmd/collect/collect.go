package collect

//go:generate go run ../../../../gen/genoai/main.go --input ./specs/collect.json --name collect --package collect --output collect-cmds.go --tpl ./gen/tpl/gen-cmds.tpl --imports github.com/spf13/cobra
//go:generate goimports -w collect-cmds.go

import (
	"github.com/spf13/cobra"
)

// Cmd -- used to connection to rootCmd
func Cmd() *cobra.Command {
	return collectCmd
}

var collectCmd = &cobra.Command{
	Use:   "collect",
	Short: "collect service",
}
