package handler

import (
	"net/http"

	fiber "github.com/gofiber/fiber/v2"

	"github.com/remusxb/todo_crud/pkg/dto"
)

const (
	Serving = "Serving"
)

type HealthCheck struct {
	// dependencies go here
}

func (h HealthCheck) Health(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(dto.HealthCheck{Status: Serving})
}
