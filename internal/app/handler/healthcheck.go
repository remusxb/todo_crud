package handler

import (
	"net/http"

	fiber "github.com/gofiber/fiber/v2"

	"github.com/remusxb/todo_crud/internal/app/usecase"
	"github.com/remusxb/todo_crud/pkg/dto"
)

const (
	HealthMessageOK    = "OK"
	HealthMessageNotOK = "Database not ready"
)

type HealthCheck struct {
	useCase usecase.HealthCheck
}

func NewHealthCheckHandler(useCase usecase.HealthCheck) HealthCheck {
	return HealthCheck{
		useCase: useCase,
	}
}

func (h HealthCheck) Health(c *fiber.Ctx) error {
	resp := dto.HealthCheckOutput{Message: HealthMessageOK}
	err := h.useCase.PingDatabase(c.Context())
	if err != nil {
		resp = dto.HealthCheckOutput{
			Message: HealthMessageNotOK,
			Error:   err.Error(),
		}
	}

	return c.Status(http.StatusOK).JSON(resp)
}
