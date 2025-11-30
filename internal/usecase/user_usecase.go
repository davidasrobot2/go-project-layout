package usecase

import (
	"context"
	"davidasrobot2/go-boilerplate/internal/domain"
	"davidasrobot2/go-boilerplate/pkg/auth"
	"davidasrobot2/go-boilerplate/pkg/constant"
	"davidasrobot2/go-boilerplate/pkg/helper"
	"log/slog"
)

type userUsecase struct {
	logger       *slog.Logger
	userRepo     domain.UserRepository
	jwtGenerator *auth.JWTGenerator
}

// NewUserUsecase creates a new use case for user logic, satisfying the domain.UserUsecase interface.
func NewUserUsecase(logger *slog.Logger, userRepo domain.UserRepository, jwtGenerator *auth.JWTGenerator) domain.UserUsecase {
	return &userUsecase{
		logger:       logger,
		userRepo:     userRepo,
		jwtGenerator: jwtGenerator,
	}
}

// Create handles the business logic for creating a new user.
func (uc *userUsecase) Create(ctx context.Context, req domain.UserCreateForm) (*domain.User, error) {

	hashedPassword, err := helper.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user := domain.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
		Phone:    req.Phone,
		Address:  req.Address,
	}

	if err := uc.userRepo.Create(ctx, &user); err != nil {
		return nil, err
	}

	return &user, nil
}

// Get All users
func (uc *userUsecase) FindAll(ctx context.Context) ([]*domain.User, error) {
	users, err := uc.userRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}

// Find By ID
func (uc *userUsecase) FindByID(ctx context.Context, id string) (*domain.User, error) {
	user, err := uc.userRepo.FindByID(ctx, id)
	if err != nil {
		if err == constant.ErrorMessageNotFound {
			return nil, constant.ErrorMessageNotFound
		}
		return nil, constant.ErrorMessageInternalServerError
	}

	return user, nil
}

// update
func (uc *userUsecase) Update(ctx context.Context, userreq domain.UserUpdateForm, id string) (*domain.User, error) {
	// find by id
	user, err := uc.userRepo.FindByID(ctx, id)
	if err != nil {
		return nil, constant.ErrorMessageNotFound
	}

	user.Name = userreq.Name
	user.Email = userreq.Email
	user.Phone = userreq.Phone
	user.Address = userreq.Address

	// update
	err = uc.userRepo.Update(ctx, user)
	if err != nil {
		return nil, constant.ErrorMessageInternalServerError
	}

	return user, nil
}

// delete
func (uc *userUsecase) Delete(ctx context.Context, id string) error {
	err := uc.userRepo.Delete(ctx, id)
	if err != nil {
		return constant.ErrorMessageNotFound
	}

	return nil
}

// Login handles the business logic for user authentication.
func (uc *userUsecase) SignIn(ctx context.Context, email, password string) (*domain.AuthToken, error) {
	user, err := uc.userRepo.FindByEmail(ctx, email)
	if err != nil {
		return nil, constant.ErrorMessageInvalidCredential // Return a known error
	}

	if !helper.CheckPasswordHash(password, user.Password) {
		return nil, constant.ErrorMessageInvalidCredential // Treat as not found for security
	}

	// Generate JWT
	token, err := uc.jwtGenerator.GenerateToken(user.ID.String())
	return token, err
}

func (uc *userUsecase) ActivateUser(ctx context.Context, id string, password string) (*domain.User, error) {
	// find by id
	user, err := uc.userRepo.FindByID(ctx, id)
	if err != nil {
		return nil, constant.ErrorMessageNotFound
	}
	user.Status = true

	// hash user new password
	hashedPassword, err := helper.HashPassword(password)
	if err != nil {
		return nil, err
	}
	user.Password = hashedPassword

	// update
	err = uc.userRepo.Update(ctx, user)
	if err != nil {
		return nil, constant.ErrorMessageInternalServerError
	}
	return user, nil
}

func (uc *userUsecase) GetCurrentUser(ctx context.Context, id string) (*domain.User, error) {
	user, err := uc.userRepo.FindByIDWithDetail(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
