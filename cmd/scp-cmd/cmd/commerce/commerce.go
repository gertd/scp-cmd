package commerce

//go:generate go run ../../../../gen/genoai/main.go --input ./specs/commerce.json --name commerce --package commerce --output commerce-cmds.go --tpl ./gen/tpl/gen-cmds.tpl --imports github.com/spf13/cobra
//go:generate goimports -w commerce-cmds.go

import (
	"github.com/spf13/cobra"
)

// Cmd -- used to connection to rootCmd
func Cmd() *cobra.Command {
	return commerceCmd
}

var commerceCmd = &cobra.Command{
	Use:   "commerce",
	Short: "commerce service",
}
