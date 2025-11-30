package cmd

import (
	"davidasrobot2/go-boilerplate/internal/di"
	"davidasrobot2/go-boilerplate/pkg/database/migration"
	"os"

	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Run database migrations",
	Run: func(cmd *cobra.Command, args []string) {
		app, err := di.InitializeApp()
		if err != nil {
			app.Logger.Error("failed to initialize app for migration", "error", err)
			os.Exit(1)
		}

		if err := migration.Run(app.DB, app.Logger); err != nil {
			app.Logger.Error("failed to run migrations", "error", err)
			os.Exit(1)
		}
		app.Logger.Info("Migrations completed successfully.")
	},
}

var freshMigrateCmd = &cobra.Command{
	Use:   "fresh-migrate",
	Short: "Run database fresh-migrations",
	Run: func(cmd *cobra.Command, args []string) {
		app, err := di.InitializeApp()
		if err != nil {
			app.Logger.Error("failed to initialize app for migration", "error", err)
			os.Exit(1)
		}

		if err := migration.FreshRun(app.DB, app.Logger); err != nil {
			app.Logger.Error("failed to run fresh migrations", "error", err)
			os.Exit(1)
		}
		app.Logger.Info("Migrations completed successfully.")
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd, freshMigrateCmd)
}
