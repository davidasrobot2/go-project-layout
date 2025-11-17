package handler

import (
	"davidasrobot/project-layout/internal/domain"
	"davidasrobot/project-layout/pkg/response"

	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
)

// AdminHandlerSet is a Wire provider set that provides a new AdminHandler.
var AdminHandlerSet = wire.NewSet(NewAdminHandler)

// AdminHandler handles HTTP requests for users.
type AdminHandler struct {
	userUsecase domain.UserUsecase
}

// NewAdminHandler creates a new AdminHandler.
func NewAdminHandler(userUsecase domain.UserUsecase) *AdminHandler {
	return &AdminHandler{userUsecase: userUsecase}
}

// CreateUserRequest defines the request body for creating a user.
type CreateAdminRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

// Create handles the creation of a new user.
func (h *AdminHandler) Create(c *fiber.Ctx) error {
	var req CreateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return err
	}

	// Note: Add validation logic here using a validator library.

	user, err := h.userUsecase.Create(c.Context(), req.Name, req.Email, req.Password)
	if err != nil {
		// In a real app, you'd check for specific errors, like duplicate email
		return err
	}

	return response.Success(c, user)
}
