package catalog

//go:generate go run ../../../../gen/genoai/main.go --input ./specs/catalog.json --name catalog --package catalog --output catalog-cmds.go --tpl ./gen/tpl/gen-cmds.tpl --imports github.com/spf13/cobra
//go:generate goimports -w catalog-cmds.go

import (
	"github.com/spf13/cobra"
)

// Cmd -- used to connection to rootCmd
func Cmd() *cobra.Command {
	return catalogCmd
}

var catalogCmd = &cobra.Command{
	Use:   "catalog",
	Short: "catalog service",
}
