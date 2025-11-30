package cmd

import (
	"davidasrobot2/go-boilerplate/internal/di"
	"os"

	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Starts the HTTP server",
	Run: func(cmd *cobra.Command, args []string) {
		app, err := di.InitializeApp()
		if err != nil {
			app.Logger.Error("failed to initialize app", "error", err)
			os.Exit(1)
		}

		if err := app.Start(); err != nil {
			app.Logger.Error("failed to start server", "error", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
