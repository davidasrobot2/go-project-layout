package usecase

import (
	"context"
	"davidasrobot/project-layout/internal/domain"
	"davidasrobot/project-layout/internal/helper"
	"davidasrobot/project-layout/pkg/auth"
	"davidasrobot/project-layout/pkg/constant"

	"github.com/google/wire"
)

// Providers is a Wire provider set that provides a new UserUsecase.
var Providers = wire.NewSet(NewUserUsecase)

type userUsecase struct {
	userRepo     domain.UserRepository
	jwtGenerator *auth.JWTGenerator
}

// NewUserUsecase creates a new use case for user logic, satisfying the domain.UserUsecase interface.
func NewUserUsecase(userRepo domain.UserRepository, jwtGenerator *auth.JWTGenerator) domain.UserUsecase {
	return &userUsecase{
		userRepo:     userRepo,
		jwtGenerator: jwtGenerator,
	}
}

// Create handles the business logic for creating a new user.
func (uc *userUsecase) Create(ctx context.Context, name, email, password string) (*domain.User, error) {
	hashedPassword, err := helper.HashPassword(password)
	if err != nil {
		return nil, err
	}

	user := &domain.User{
		Name:     name,
		Email:    email,
		Password: hashedPassword,
	}

	if err := uc.userRepo.Create(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

// Login handles the business logic for user authentication.
func (uc *userUsecase) Login(ctx context.Context, email, password string) (string, error) {
	user, err := uc.userRepo.FindByEmail(ctx, email)
	if err != nil {
		return "", constant.ErrNotFound // Return a known error
	}

	if !helper.CheckPasswordHash(password, user.Password) {
		return "", constant.ErrNotFound // Treat as not found for security
	}

	// Generate JWT
	token, err := uc.jwtGenerator.GenerateToken(user.ID)
	return token, err
}
