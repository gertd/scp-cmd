package cmd

import (
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/spf13/cobra"

	"gitlab.com/d5s/scp-cmd/cmd/scp-cmd/cmd/action"
	"gitlab.com/d5s/scp-cmd/cmd/scp-cmd/cmd/appreg"
	"gitlab.com/d5s/scp-cmd/cmd/scp-cmd/cmd/catalog"
	"gitlab.com/d5s/scp-cmd/cmd/scp-cmd/cmd/collect"
	"gitlab.com/d5s/scp-cmd/cmd/scp-cmd/cmd/config"
	"gitlab.com/d5s/scp-cmd/cmd/scp-cmd/cmd/forwarders"
	"gitlab.com/d5s/scp-cmd/cmd/scp-cmd/cmd/identity"
	"gitlab.com/d5s/scp-cmd/cmd/scp-cmd/cmd/ingest"
	"gitlab.com/d5s/scp-cmd/cmd/scp-cmd/cmd/login"
	"gitlab.com/d5s/scp-cmd/cmd/scp-cmd/cmd/ml"
	"gitlab.com/d5s/scp-cmd/cmd/scp-cmd/cmd/provisioner"
	"gitlab.com/d5s/scp-cmd/cmd/scp-cmd/cmd/search"
	"gitlab.com/d5s/scp-cmd/cmd/scp-cmd/cmd/streams"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     appName,
	Short:   appDescr,
	Version: versionInfo(),
}

// Execute -- Adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {

	rootCmd.PersistentFlags().String("tenant", "", "tenant identifier")

	rootCmd.AddCommand(action.Cmd())
	rootCmd.AddCommand(appreg.Cmd())
	rootCmd.AddCommand(catalog.Cmd())
	rootCmd.AddCommand(collect.Cmd())
	rootCmd.AddCommand(config.Cmd())
	rootCmd.AddCommand(forwarders.Cmd())
	rootCmd.AddCommand(identity.Cmd())
	rootCmd.AddCommand(ingest.Cmd())
	rootCmd.AddCommand(login.Cmd())
	rootCmd.AddCommand(ml.Cmd())
	rootCmd.AddCommand(provisioner.Cmd())
	rootCmd.AddCommand(search.Cmd())
	rootCmd.AddCommand(streams.Cmd())
}

func versionInfo() string {

	versionString := version
	if versionString == "" {
		versionString = "0.0.0"
	}

	commitString := commit
	if commitString == "" {
		commitString = "develop"
	}

	buildString := date
	if buildString == "" {
		buildString = time.Now().Format(time.RFC3339)
	}

	return fmt.Sprintf("%s #%s [%s.%s.%s]",
		versionString,
		commitString,
		runtime.GOOS,
		runtime.GOARCH,
		buildString,
	)
}
