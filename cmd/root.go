package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "kiwi-server",
	Short: "Kiwi Server",
	Long: `Kiwi Server is an open-source application that provides an API to fetch data from the PostgreSQL integration of ChirpStack.
		> documentation & support: coming soon...
		> source & copyright information: coming soon...`,
	RunE: run,
}

// Execute executes the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		// log.Fatal(err)
	}
}
