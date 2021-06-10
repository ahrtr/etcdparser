package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
)

var (
	dataDir string
	showDetails bool
	rawFormat bool
)

func CreateRootCommand() *cobra.Command {
	var rootCmd  = &cobra.Command {
		Use:   "ep",
		Short: "Etcd parser",
		Long: "Etcd parser is used to parse etcd's data, including WAL and snapshot",
	}

	rootCmd.PersistentFlags().StringVarP(&dataDir, "data-dir", "d", "", "Etcd data directory")
	rootCmd.MarkPersistentFlagRequired("data-dir")
	rootCmd.PersistentFlags().BoolVarP(&showDetails, "show-details", "s", false, "Whether to show the details: entries or snapshot data")
	rootCmd.PersistentFlags().BoolVarP(&rawFormat, "raw", "r", false, "Whether to print the data in raw format")

	rootCmd.AddCommand(createWALCommand())
	rootCmd.AddCommand(createSnapCommand())

	return rootCmd
}

func checkDataDir() error {
	if _, err := os.Stat(dataDir); os.IsNotExist(err) {
		return fmt.Errorf("the data directory %q doesn't exist", dataDir)
	}

	return nil
}

func silenceUsage(f func(cmd *cobra.Command, args []string) error) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		cmd.SilenceUsage = true
		return f(cmd, args)
	}
}