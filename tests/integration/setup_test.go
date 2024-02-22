//go:build integration

package integration

import (
	"context"
	"log"
	"os"
	"testing"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/remusxb/todo_crud/internal/app/handler"
	"github.com/remusxb/todo_crud/internal/app/routes"
	"github.com/remusxb/todo_crud/internal/app/usecase"
	"github.com/remusxb/todo_crud/internal/metrics"
)

type healthTests struct {
	app *fiber.App
}

var ht healthTests

func TestMain(m *testing.M) {
	err := setup()
	if err != nil {
		log.Fatal(err.Error())
	}

	code := m.Run()

	os.Exit(code)
}

func setup() error {
	// set env vars if needed; e.g. : os.Setenv("ENV_KEY", "value")

	app := fiber.New()

	// init use cases
	healthUseCase := usecase.NewHealthCheckUseCase(&healthRepoMock{
		PingFunc: func(ctx context.Context) error {
			return nil
		},
	})

	// init handlers
	healthHandler := handler.NewHealthCheckHandler(healthUseCase)
	reg := prometheus.NewRegistry()
	reg.MustRegister(metrics.DefaultCollectors()...)
	promHandler := promhttp.HandlerFor(reg, promhttp.HandlerOpts{Registry: reg})

	router := routes.Router{
		Server:        app,
		HealthHandler: healthHandler,
		PromHandler:   promHandler,
	}
	router.RegisterHTTPRoutes()
	ht.app = app

	return nil
}
