package usecase_test

import (
	"context"
	"davidasrobot2/go-boilerplate/config"
	"davidasrobot2/go-boilerplate/internal/domain"
	"davidasrobot2/go-boilerplate/internal/domain/mocks"
	"davidasrobot2/go-boilerplate/internal/usecase"
	"davidasrobot2/go-boilerplate/pkg/auth"
	"davidasrobot2/go-boilerplate/pkg/constant"
	"davidasrobot2/go-boilerplate/pkg/helper"
	"errors"
	"io"
	"log/slog"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUserUsecase_Create(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))

	t.Run("Success", func(t *testing.T) {
		mockRepo := mocks.NewUserRepository(t)
		uc := usecase.NewUserUsecase(logger, mockRepo, nil) // jwtGenerator not needed

		createForm := domain.UserCreateForm{
			Name:     "John Doe",
			Email:    "john.doe@example.com",
			Password: "password123",
			Phone:    "+15555555555",
			Address:  "123 Main St",
		}

		// We expect the Create method to be called.
		// We use mock.AnythingOfType because the ID and timestamps are generated inside the usecase.
		mockRepo.On("Create", mock.Anything, mock.AnythingOfType("*domain.User")).Return(nil).Once()

		user, err := uc.Create(context.Background(), createForm)

		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, createForm.Name, user.Name)
		assert.Equal(t, createForm.Email, user.Email)
		// Verify the password was hashed
		assert.True(t, helper.CheckPasswordHash(createForm.Password, user.Password))
	})

	t.Run("Repository Error", func(t *testing.T) {
		mockRepo := mocks.NewUserRepository(t)
		uc := usecase.NewUserUsecase(logger, mockRepo, nil)

		createForm := domain.UserCreateForm{
			Name:     "John Doe",
			Email:    "john.doe@example.com",
			Password: "password123",
		}

		expectedError := errors.New("database error")
		mockRepo.On("Create", mock.Anything, mock.AnythingOfType("*domain.User")).Return(expectedError).Once()

		user, err := uc.Create(context.Background(), createForm)

		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
		assert.Nil(t, user)
	})
}

func TestUserUsecase_SignIn(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	// For SignIn, we need a real JWTGenerator
	cfg := &config.Config{JWT: config.JWTConfig{Secret: "test-secret"}}
	jwtGenerator := auth.NewJWTGenerator(cfg)

	t.Run("Success", func(t *testing.T) {
		mockRepo := mocks.NewUserRepository(t)
		uc := usecase.NewUserUsecase(logger, mockRepo, jwtGenerator)

		email := "jane.doe@example.com"
		password := "strongpassword"
		hashedPassword, _ := helper.HashPassword(password)

		mockUser := &domain.User{
			BaseDomain: domain.BaseDomain{ID: uuid.New()},
			Email:      email,
			Password:   hashedPassword,
		}

		mockRepo.On("FindByEmail", mock.Anything, email).Return(mockUser, nil).Once()

		token, err := uc.SignIn(context.Background(), email, password)

		assert.NoError(t, err)
		assert.NotNil(t, token)
		assert.NotEmpty(t, token.Token)
	})

	t.Run("User Not Found", func(t *testing.T) {
		mockRepo := mocks.NewUserRepository(t)
		uc := usecase.NewUserUsecase(logger, mockRepo, jwtGenerator)

		email := "notfound@example.com"
		password := "password"

		// Mock repository returns a "not found" error
		mockRepo.On("FindByEmail", mock.Anything, email).Return(nil, constant.ErrorMessageNotFound).Once()

		token, err := uc.SignIn(context.Background(), email, password)

		assert.Error(t, err)
		assert.Nil(t, token)
		// The usecase should return a generic "invalid credential" error for security
		assert.Equal(t, constant.ErrorMessageInvalidCredential, err)
	})

	t.Run("Incorrect Password", func(t *testing.T) {
		mockRepo := mocks.NewUserRepository(t)
		uc := usecase.NewUserUsecase(logger, mockRepo, jwtGenerator)

		email := "jane.doe@example.com"
		correctPassword := "strongpassword"
		wrongPassword := "wrongpassword"
		hashedPassword, _ := helper.HashPassword(correctPassword)

		mockUser := &domain.User{
			BaseDomain: domain.BaseDomain{ID: uuid.New()},
			Email:      email,
			Password:   hashedPassword,
		}

		mockRepo.On("FindByEmail", mock.Anything, email).Return(mockUser, nil).Once()

		token, err := uc.SignIn(context.Background(), email, wrongPassword)

		assert.Error(t, err)
		assert.Nil(t, token)
		assert.Equal(t, constant.ErrorMessageInvalidCredential, err)
	})
}
