package search

//go:generate go run ../../../../gen/genoai/main.go --input ./specs/search.json --name search --package search --output search-cmds.go --tpl ./gen/tpl/gen-cmds.tpl --imports github.com/spf13/cobra
//go:generate goimports -w search-cmds.go

import (
	"github.com/spf13/cobra"
)

// Cmd -- used to connection to rootCmd
func Cmd() *cobra.Command {
	return searchCmd
}

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "search service",
}
