package appreg

//go:generate go run ../../../../gen/genoai/main.go --input ./specs/app-registry.json --name app-registry --package appreg --output appreg-cmds.go --tpl ./gen/tpl/gen-cmds.tpl --imports github.com/spf13/cobra
//go:generate goimports -w appreg-cmds.go

import (
	"github.com/spf13/cobra"
)

// Cmd -- used to connection to rootCmd
func Cmd() *cobra.Command {
	return appregCmd
}

var appregCmd = &cobra.Command{
	Use:   "appreg",
	Short: "application registry service",
}
