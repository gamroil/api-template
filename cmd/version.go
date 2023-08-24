package cmd

import (
	"runtime/debug"

	"go.uber.org/zap"

	"github.com/spf13/cobra"
)

func newVersionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the API binary version",
		Run: func(cmd *cobra.Command, args []string) {
			if info, ok := debug.ReadBuildInfo(); ok {
				for _, setting := range info.Settings {
					if setting.Key == "vcs.revision" {
						logger := zap.Must(zap.NewProduction())
						defer func() {
							_ = logger.Sync()
						}()

						logger.Info(setting.Value)
					}
				}
			}
		},
	}
}
