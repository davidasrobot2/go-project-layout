package handler

import (
	"davidasrobot2/go-boilerplate/internal/domain"
	"davidasrobot2/go-boilerplate/pkg/constant"
	"davidasrobot2/go-boilerplate/pkg/response"
	"log/slog"

	// "davidasrobot2/go-boilerplate/pkg/response" // No longer needed here

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
)

// UserHandlerSet is a Wire provider set that provides a new UserHandler.
var UserHandlerSet = wire.NewSet(NewUserHandler)

// UserHandler handles HTTP requests for users.
type UserHandler struct {
	logger      *slog.Logger
	userUsecase domain.UserUsecase
	validate    *validator.Validate
}

// NewUserHandler creates a new UserHandler.
func NewUserHandler(logger *slog.Logger, userUsecase domain.UserUsecase, validate *validator.Validate) *UserHandler {
	return &UserHandler{
		userUsecase: userUsecase,
		validate:    validate,
		logger:      logger,
	}
}

// GetAll handles getting all users.
func (h *UserHandler) GetAll(c *fiber.Ctx) error {
	users, err := h.userUsecase.FindAll(c.Context())
	if err != nil {
		return err
	}

	return response.HandleSuccess(c, users)
}

// Create handles the creation of a new user.
func (h *UserHandler) Create(c *fiber.Ctx) error {
	var req domain.UserCreateForm
	if err := c.BodyParser(&req); err != nil {
		return err
	}

	// Validate the request struct
	if err := h.validate.Struct(req); err != nil {
		return err
	}

	user, err := h.userUsecase.Create(c.Context(), req)
	if err != nil {
		return err // The middleware will catch this and return a 500 error
	}

	return response.HandleSuccess(c, user)
}

// FindByID handle get user by id
func (h *UserHandler) FindByID(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := h.userUsecase.FindByID(c.Context(), id)
	if err != nil {
		return err
	}

	return response.HandleSuccess(c, user)
}

// Update by ID
func (h *UserHandler) Update(c *fiber.Ctx) error {
	id := c.Params("id")
	var req domain.UserUpdateForm
	if err := c.BodyParser(&req); err != nil {
		return err
	}
	if err := h.validate.Struct(req); err != nil {
		return err
	}

	user, err := h.userUsecase.Update(c.Context(), req, id)
	if err != nil {
		return err
	}

	return response.HandleSuccess(c, user)
}

func (h *UserHandler) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	err := h.userUsecase.Delete(c.Context(), id)
	if err != nil {
		return err
	}

	return response.HandleSuccess(c, nil)
}

func (h *UserHandler) SignIn(c *fiber.Ctx) error {
	var req domain.UserSignInForm
	if err := c.BodyParser(&req); err != nil {
		return err
	}

	if err := h.validate.Struct(req); err != nil {
		return err
	}

	token, err := h.userUsecase.SignIn(c.Context(), req.Email, req.Password)
	if err != nil {
		return err
	}

	return response.HandleSuccess(c, token)
}

func (h *UserHandler) ActivateUser(c *fiber.Ctx) error {
	// change the user status to true and change password to new password
	var req domain.UpdatePasswordForm
	if err := c.BodyParser(&req); err != nil {
		return err
	}

	if err := h.validate.Struct(req); err != nil {
		return err
	}
	// get user ID from context local
	id, ok := c.Locals("user_id").(string)
	if !ok {
		return constant.ErrorMessageInvalidToken
	}
	h.logger.Info("Activating user from token", "id", id)

	user, err := h.userUsecase.ActivateUser(c.Context(), id, req.Password)
	if err != nil {
		return err
	}

	return response.HandleSuccess(c, user)
}

func (h *UserHandler) GetCurrentUser(c *fiber.Ctx) error {
	// get user ID from context local
	id, ok := c.Locals("user_id").(string)
	if !ok {
		return constant.ErrorMessageInvalidToken
	}
	h.logger.Info("Getting current user from token", "id", id)

	user, err := h.userUsecase.FindByID(c.Context(), id)
	if err != nil {
		return err
	}

	return response.HandleSuccess(c, user)
}
