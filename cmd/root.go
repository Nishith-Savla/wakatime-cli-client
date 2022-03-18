package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "wakacli",
	Short: "A WakaTime Cli Client",
	Long:  `A WakaTime Cli client to view your WakaTime statistics right inside your terminal.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
