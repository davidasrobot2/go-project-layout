package repository

import (
	"context"
	"davidasrobot2/go-boilerplate/internal/domain"
	"davidasrobot2/go-boilerplate/pkg/constant"
	"errors"
	"log/slog"

	"gorm.io/gorm"
)

type userRepository struct {
	logger *slog.Logger
	db     *gorm.DB
}

// NewUserRepository creates a new repository for user data.
func NewUserRepository(db *gorm.DB, logger *slog.Logger) domain.UserRepository {
	return &userRepository{db: db, logger: logger}
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
			return nil, constant.ErrorMessageNotFound
		}
		return nil, err
	}
	return &user, nil
}

// Get All user
func (r *userRepository) FindAll(ctx context.Context) ([]*domain.User, error) {
	var users []*domain.User
	err := r.db.WithContext(ctx).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

// FindByID finds a user by their ID.
func (r *userRepository) FindByID(ctx context.Context, id string) (*domain.User, error) {
	var user domain.User
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, constant.ErrorMessageNotFound
		}
		r.logger.Error(err.Error())
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Update(ctx context.Context, user *domain.User) error {
	err := r.db.WithContext(ctx).Save(user).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepository) Delete(ctx context.Context, id string) error {
	var user domain.User
	err := r.db.WithContext(ctx).Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return err
	}
	return nil
}

// FindByID finds a user by their ID.
func (r *userRepository) FindByIDWithDetail(ctx context.Context, id string) (*domain.User, error) {
	var user domain.User
	err := r.db.Preload("Merchant", "Accounts").WithContext(ctx).Where("id = ?", id).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, constant.ErrorMessageNotFound
		}
		r.logger.Error(err.Error())
		return nil, err
	}
	return &user, nil
}
