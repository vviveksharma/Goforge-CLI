package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "goforge",
	Short: "Forge production-ready Go applications",
	Long: `goforge is a CLI tool to generate production-ready Go projects 
with best practices, security, and observability built-in.`,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
}
