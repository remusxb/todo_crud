package routes

import (
	"net/http"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"

	"github.com/remusxb/todo_crud/internal/app/handler"
)

type Router struct {
	Server        *fiber.App
	HealthHandler handler.HealthCheck
	PromHandler   http.Handler
}

func (r Router) RegisterHTTPRoutes() {
	// register routes
	r.Server.Get("/health", r.HealthHandler.Health)
	r.Server.Get("/metrics", adaptor.HTTPHandler(r.PromHandler))
}
