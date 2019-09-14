package cmd

import (
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/spf13/cobra"

	"github.com/gertd/scp-cmd/cmd/scp-cmd/cmd/action"
	"github.com/gertd/scp-cmd/cmd/scp-cmd/cmd/appreg"
	"github.com/gertd/scp-cmd/cmd/scp-cmd/cmd/catalog"
	"github.com/gertd/scp-cmd/cmd/scp-cmd/cmd/collect"
	"github.com/gertd/scp-cmd/cmd/scp-cmd/cmd/config"
	"github.com/gertd/scp-cmd/cmd/scp-cmd/cmd/forwarders"
	"github.com/gertd/scp-cmd/cmd/scp-cmd/cmd/identity"
	"github.com/gertd/scp-cmd/cmd/scp-cmd/cmd/ingest"
	"github.com/gertd/scp-cmd/cmd/scp-cmd/cmd/login"
	"github.com/gertd/scp-cmd/cmd/scp-cmd/cmd/ml"
	"github.com/gertd/scp-cmd/cmd/scp-cmd/cmd/provisioner"
	"github.com/gertd/scp-cmd/cmd/scp-cmd/cmd/search"
	"github.com/gertd/scp-cmd/cmd/scp-cmd/cmd/streams"
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
