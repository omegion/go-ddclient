package main

import (
	"os"

	"github.com/omegion/go-ddclient/cmd"

	"github.com/spf13/cobra"
)

// RootCommand is the main entry point of this application.
func RootCommand() *cobra.Command {
	root := &cobra.Command{
		Use:          "ddclient",
		Short:        "Dynamic DNS Client CLI application",
		Long:         "Dynamic DNS Client CLI application to keep DNS record updated.",
		SilenceUsage: true,
	}

	root.AddCommand(cmd.Version())
	root.AddCommand(cmd.Set())

	return root
}

func main() {
	if err := RootCommand().Execute(); err != nil {
		os.Exit(1)
	}
}
