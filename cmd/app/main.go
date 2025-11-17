package main

import (
	"davidasrobot/project-layout/internal/di"
	"davidasrobot/project-layout/pkg/migration"
	"os"
)

func main() {
	// Initialize the application using dependency injection
	app, err := di.InitializeApp()
	if err != nil {
		// Use a basic logger for startup errors
		// log.Fatalf("failed to initialize app: %v", err)
		// For now, just exit. We will use the logger soon.
		os.Exit(1)
	}

	// Check database connection
	sqlDB, err := app.DB.DB()
	if err != nil {
		app.Logger.Error("failed to get underlying sql.DB", "error", err)
		os.Exit(1)
	}
	if err := sqlDB.Ping(); err != nil {
		app.Logger.Error("failed to ping database", "error", err)
		os.Exit(1)
	}
	app.Logger.Info("Database connection successful.")

	// Auto-migrate database schema
	if err := migration.Run(app.DB, app.Logger); err != nil {
		app.Logger.Error("failed to migrate database", "error", err)
		os.Exit(1)
	}

	// convenience way to Register routes
	// app.RegisterRoutes()

	// Start the server
	if err := app.Start(); err != nil {
		app.Logger.Error("failed to start server", "error", err)
	}
}
