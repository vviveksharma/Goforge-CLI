package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// These variables are set by goreleaser during build
var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version information",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("goforge version %s\n", version)
		fmt.Printf("  commit: %s\n", commit)
		fmt.Printf("  built at: %s\n", date)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
