package cmd

import (
	"github.com/spf13/cobra"
)

func newRootCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "api-template",
		Short: "A template for APIs",
	}
}

// Execute initializes and executes the CLI.
func Execute() error {
	rootCmd := newRootCmd()
	rootCmd.AddCommand(newServeCmd())
	rootCmd.AddCommand(newVersionCmd())
	return rootCmd.Execute()
}
