package cmd

import (
	"davidasrobot2/go-boilerplate/internal/di"
	"davidasrobot2/go-boilerplate/pkg/database/seeder"
	"os"

	"github.com/spf13/cobra"
)

var seedCmd = &cobra.Command{
	Use:   "seed",
	Short: "Seed the database with initial data",
	Run: func(cmd *cobra.Command, args []string) {
		app, err := di.InitializeApp()
		if err != nil {
			app.Logger.Error("failed to initialize app for seeder", "error", err)
			os.Exit(1)
		}

		if err := seeder.Run(app.DB, app.Logger); err != nil {
			app.Logger.Error("failed to run seeder", "error", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(seedCmd)
}
