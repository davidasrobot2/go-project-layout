package handler

import (
	"davidasrobot/project-layout/internal/domain"
	"davidasrobot/project-layout/pkg/response"

	// "davidasrobot/project-layout/pkg/response" // No longer needed here

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
)

// UserHandlerSet is a Wire provider set that provides a new UserHandler.
var UserHandlerSet = wire.NewSet(NewUserHandler)

// UserHandler handles HTTP requests for users.
type UserHandler struct {
	userUsecase domain.UserUsecase
	validate    *validator.Validate
}

// NewUserHandler creates a new UserHandler.
func NewUserHandler(userUsecase domain.UserUsecase, validate *validator.Validate) *UserHandler {
	return &UserHandler{
		userUsecase: userUsecase,
		validate:    validate,
	}
}

// CreateUserRequest defines the request body for creating a user.
type CreateUserRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

// Create handles the creation of a new user.
func (h *UserHandler) Create(c *fiber.Ctx) error {
	var req CreateUserRequest
	if err := c.BodyParser(&req); err != nil {
		// Let the middleware handle this generic error
		return err
	}

	// Validate the request struct
	if err := h.validate.Struct(req); err != nil {
		return err // The middleware will detect validator.ValidationErrors and format it
	}

	user, err := h.userUsecase.Create(c.Context(), req.Name, req.Email, req.Password)
	if err != nil {
		return err // The middleware will catch this and return a 500 error
	}

	return response.Success(c, user)
}

// LoginRequest defines the request body for logging in.
type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// Login handles user authentication and returns a JWT.
func (h *UserHandler) Login(c *fiber.Ctx) error {
	var req LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return err
	}

	if err := h.validate.Struct(req); err != nil {
		return err
	}

	token, err := h.userUsecase.Login(c.Context(), req.Email, req.Password)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"token": token})
}
