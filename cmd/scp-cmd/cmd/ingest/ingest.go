package ingest

//go:generate go run ../../../../gen/genoai/main.go --input ./specs/ingest.json --name ingest --package ingest --output ingest-cmds.go --tpl ./gen/tpl/gen-cmds.tpl --imports github.com/spf13/cobra
//go:generate goimports -w ingest-cmds.go

import (
	"github.com/spf13/cobra"
)

// Cmd -- used to connection to rootCmd
func Cmd() *cobra.Command {
	return ingestCmd
}

var ingestCmd = &cobra.Command{
	Use:   "ingest",
	Short: "ingest service",
}
