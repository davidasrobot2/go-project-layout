package seeder

import (
	"davidasrobot2/go-boilerplate/internal/domain"
	"davidasrobot2/go-boilerplate/pkg/constant"
	"davidasrobot2/go-boilerplate/pkg/helper"
	"fmt"
	"log/slog"

	"gorm.io/gorm"
)

func Run(db *gorm.DB, logger *slog.Logger) error {
	logger.Info("Running seeders...")
	if err := seedAdministrators(db, logger); err != nil {
		return err
	}
	logger.Info("Seeders ran successfully.")
	return nil
}

func seedAdministrators(db *gorm.DB, logger *slog.Logger) error {
	var count int64
	db.Model(&domain.Administrator{}).Count(&count)
	if count > 0 {
		logger.Info("Administrators already seeded, skipping.")
		return nil
	}

	hashedPassword, err := helper.HashPassword("password") // Replace with a strong default password or generate one
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	administrators := []domain.Administrator{
		{
			Name:     "Super Admin",
			Email:    "admin@example.com",
			Password: hashedPassword,
			Status:   true,
			Level:    constant.AdministratorLevel1,
		},
		{
			Name:     "Moderator",
			Email:    "moderator@example.com",
			Password: hashedPassword,
			Status:   true,
			Level:    constant.AdministratorLevel2,
		},
	}

	for _, admin := range administrators {
		if err := db.Create(&admin).Error; err != nil {
			return fmt.Errorf("failed to seed administrator %s: %w", admin.Email, err)
		}
	}

	logger.Info("Administrators seeded successfully.")
	return nil

}
