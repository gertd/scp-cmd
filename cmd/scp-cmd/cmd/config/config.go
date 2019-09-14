package config

import (
	"github.com/spf13/cobra"
	impl "github.com/gertd/scp-cmd/pkg/config"
)

// Cmd -- used to connection to rootCmd
func Cmd() *cobra.Command {
	return configCmd
}

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "config service",
}

// getConfigCmd --
var getConfigCmd = &cobra.Command{
	Use:   "get-config",
	Short: "",
	RunE:  impl.GetConfig,
}

// setConfigCmd --
var setConfigCmd = &cobra.Command{
	Use:   "set-config",
	Short: "",
	RunE:  impl.SetConfig,
}

// removeConfigCmd --
var removeConfigCmd = &cobra.Command{
	Use:   "remove-config",
	Short: "",
	RunE:  impl.RemoveConfig,
}

// listConfigCmd --
var listConfigCmd = &cobra.Command{
	Use:   "list-config",
	Short: "",
	RunE:  impl.ListConfig,
}

func init() {
	configCmd.AddCommand(getConfigCmd)
	configCmd.AddCommand(setConfigCmd)
	configCmd.AddCommand(removeConfigCmd)
	configCmd.AddCommand(listConfigCmd)
}
