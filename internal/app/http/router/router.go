package router

import (
	"davidasrobot/project-layout/config"
	"davidasrobot/project-layout/internal/app/http/handler"
	"davidasrobot/project-layout/internal/app/http/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/wire"
)

// RouterSet is a Wire provider set that provides a new Router.
var RouterSet = wire.NewSet(NewRouter)

// Router holds the handlers and registers routes.
type Router struct {
	Cfg          *config.Config
	App          *fiber.App
	UserHandler  *handler.UserHandler
	AdminHandler *handler.AdminHandler
}

// NewRouter creates a new Router.
func NewRouter(cfg *config.Config, app *fiber.App, userHandler *handler.UserHandler, adminHandler *handler.AdminHandler) *Router {
	return &Router{Cfg: cfg, App: app, UserHandler: userHandler, AdminHandler: adminHandler}
}

// RegisterRoutes registers all application routes.
func (r *Router) RegisterRoutes() {
	// Group routes under /api
	api := r.App.Group("/api")

	// Authentication routes
	api.Post("/users", r.UserHandler.Create)
	api.Post("/login", r.UserHandler.Login)

	// Protected route example
	api.Get("/me", middleware.Protected(r.Cfg), func(c *fiber.Ctx) error {
		user := c.Locals("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		return c.JSON(fiber.Map{"message": "Welcome!", "user_id": claims["sub"]})
	})

	// Admin routes
	api.Post("/admin", r.AdminHandler.Create)
}
