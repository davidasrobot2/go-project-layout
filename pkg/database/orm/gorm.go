package orm

import (
	"davidasrobot2/go-boilerplate/config"

	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Providers is a Wire provider set that provides a new GORM database connection.
var Providers = wire.NewSet(NewGormDB)

// NewGormDB creates a new GORM database instance for PostgreSQL.
func NewGormDB(cfg *config.Config) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(cfg.Database.DSN()), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
