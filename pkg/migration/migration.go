package migration

import (
	"davidasrobot/project-layout/internal/domain"
	"log/slog"

	"gorm.io/gorm"
)

// Run performs auto-migration for the application's domain models.
func Run(db *gorm.DB, logger *slog.Logger) error {
	logger.Info("Running database auto-migrations...")
	err := db.AutoMigrate(
		&domain.User{},
		// Add other domain models here as they are created
	)
	return err
}
