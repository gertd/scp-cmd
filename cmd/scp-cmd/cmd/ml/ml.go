package ml

//go:generate go run ../../../../gen/genoai/main.go --input ./specs/ml.json --name ml --package ml --output ml-cmds.go --tpl ./gen/tpl/gen-cmds.tpl --imports github.com/spf13/cobra
//go:generate gofmt -w ml-cmds.go

import (
	"github.com/spf13/cobra"
)

// Cmd -- used to connection to rootCmd
func Cmd() *cobra.Command {
	return mlCmd
}

var mlCmd = &cobra.Command{
	Use:   "ml",
	Short: "ml service",
}
