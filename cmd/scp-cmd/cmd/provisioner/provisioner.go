package provisioner

//go:generate go run ../../../../gen/genoai/main.go --input ./specs/provisioner.json --name provisioner --package provisioner --output provisioner-cmds.go --tpl ./gen/tpl/gen-cmds.tpl --imports github.com/spf13/cobra
//go:generate goimports -w provisioner-cmds.go

import (
	"github.com/spf13/cobra"
)

// Cmd -- used to connection to rootCmd
func Cmd() *cobra.Command {
	return provisionerCmd
}

var provisionerCmd = &cobra.Command{
	Use:   "provisioner",
	Short: "provisioner service",
}
