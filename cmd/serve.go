package cmd

import (
	"api-template/api"

	"github.com/spf13/cobra"
)

func newServeCmd() *cobra.Command {
	var configFile string

	command := &cobra.Command{
		Use:   "serve",
		Short: "Start the API",
		Run: func(cmd *cobra.Command, args []string) {
			app := api.New(configFile)
			app.Start()
		},
	}

	command.Flags().
		StringVarP(&configFile, "config", "c", "config.yaml", "config file (default is config.yaml)")

	return command
}
