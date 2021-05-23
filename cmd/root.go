package cmd

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

//nolint:gochecknoglobals // for root cmd.
var rootCmd = &cobra.Command{
	Use:          "ddclient",
	Short:        "Dynamic DNS Client CLI application",
	Long:         "Dynamic DNS Client CLI application to keep DNS record updated.",
	SilenceUsage: true,
}

func setPersistentFlags() {
	rootCmd.PersistentFlags().String("logLevel", "info", "Set the logging level. One of: debug|info|warn|error")
}

func initConfig() {
	logLevel, _ := rootCmd.Flags().GetString("logLevel")

	level, err := log.ParseLevel(logLevel)
	if err != nil {
		log.Fatalf("Lethal damage: %s\n\n", err)
	}

	log.SetLevel(level)
	log.SetFormatter(&log.TextFormatter{
		TimestampFormat: "02-01-2006 15:04:05",
		FullTimestamp:   true,
	})
}

// Execute is entry point for commands.
func Execute() {
	cobra.OnInitialize(initConfig)

	setPersistentFlags()

	rootCmd.AddCommand(Version())
	rootCmd.AddCommand(Set())

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
