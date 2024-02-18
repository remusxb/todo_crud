package routes

import (
	"net/http"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"

	"github.com/remusxb/todo_crud/internal/app/handler"
)

func RegisterHTTPRoutes(server *fiber.App, promHandler http.Handler) {
	// init services
	healthService := handler.HealthCheck{}

	// register routes
	server.Get("/health", healthService.Health)
	server.Get("/metrics", adaptor.HTTPHandler(promHandler))
}
