package repository

import (
	"context"
	"davidasrobot/project-layout/internal/domain"
	"errors"

	"github.com/google/wire"
	"gorm.io/gorm"
)

// Providers is a Wire provider set that provides a new UserRepository.
var Providers = wire.NewSet(NewUserRepository)

type userRepository struct {
	db *gorm.DB
}

// NewUserRepository creates a new repository for user data.
func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &userRepository{db: db}
}

// Create creates a new user record in the database.
func (r *userRepository) Create(ctx context.Context, user *domain.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}

// FindByEmail finds a user by their email address.
func (r *userRepository) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	var user domain.User
	err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		return nil, err
	}
	return &user, nil
}
