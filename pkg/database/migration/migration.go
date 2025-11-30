package migration

import (
	"davidasrobot2/go-boilerplate/internal/domain"
	"log/slog"

	"gorm.io/gorm"
)

// Run performs auto-migration for the application's domain models.
func Run(db *gorm.DB, logger *slog.Logger) error {
	logger.Info("Running database auto-migrations...")
	err := db.AutoMigrate(
		&domain.User{},
		&domain.Administrator{},
		// Add other domain models here as they are created
	)
	return err
}

func FreshRun(db *gorm.DB, logger *slog.Logger) error {
	logger.Info("Running database fresh-migrations...")
	err := db.Migrator().DropTable(
		&domain.User{},
		&domain.Administrator{},
		// Add other domain models here as they are created
	)
	if err != nil {
		return err
	}
	err = Run(db, logger)
	return err
}
