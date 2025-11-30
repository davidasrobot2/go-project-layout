package handler

import (
	"davidasrobot2/go-boilerplate/internal/domain"
	"davidasrobot2/go-boilerplate/pkg/response"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
)

// AdminHandlerSet is a Wire provider set that provides a new AdminHandler.
var AdminHandlerSet = wire.NewSet(NewAdminHandler)

// AdminHandler handles HTTP requests for users.
type AdminHandler struct {
	adminUsecase domain.AdministratorUsecase
	validate     *validator.Validate
}

// NewAdminHandler creates a new AdminHandler.
func NewAdminHandler(administratorUsecase domain.AdministratorUsecase, validate *validator.Validate) *AdminHandler {
	return &AdminHandler{adminUsecase: administratorUsecase, validate: validate}
}

func (h *AdminHandler) SignIn(c *fiber.Ctx) error {
	var req domain.AdministratorLoginForm
	if err := c.BodyParser(&req); err != nil {
		return err
	}

	if err := h.validate.Struct(req); err != nil {
		return err
	}

	token, err := h.adminUsecase.SignIn(c.Context(), req.Email, req.Password)
	if err != nil {
		return err
	}

	return response.HandleSuccess(c, token)
}
