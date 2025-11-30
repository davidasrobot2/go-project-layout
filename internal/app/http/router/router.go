package router

import (
	"davidasrobot2/go-boilerplate/config"
	"davidasrobot2/go-boilerplate/internal/app/http/handler"
	"davidasrobot2/go-boilerplate/internal/app/http/middleware"

	"github.com/gofiber/fiber/v2"
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
func NewRouter(
	cfg *config.Config,
	app *fiber.App,
	userHandler *handler.UserHandler,
	adminHandler *handler.AdminHandler) *Router {
	return &Router{
		Cfg:          cfg,
		App:          app,
		UserHandler:  userHandler,
		AdminHandler: adminHandler,
	}
}

// RegisterRoutes registers all application routes.
func (r *Router) RegisterRoutes() {
	// Group routes under /api
	api := r.App.Group("/api")

	v1 := api.Group("/v1")

	dashboard := v1.Group("/dashboard")

	dashboard.Post("/signin", r.AdminHandler.SignIn)

	dashboardUsers := dashboard.Group("/users").Use(middleware.Protected(r.Cfg))
	dashboardUsers.Get("/", r.UserHandler.GetAll)
	dashboardUsers.Post("/", r.UserHandler.Create)
	dashboardUsers.Get("/:id", r.UserHandler.FindByID)
	dashboardUsers.Post("/delete/:id", r.UserHandler.Delete)
	dashboardUsers.Post("/update/:id", r.UserHandler.Update)

	// Admin routes
	// api.Post("/admin", r.AdminHandler.Create)
}
