//go:build integration

package integration

import (
	"log"
	"os"
	"testing"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/remusxb/todo_crud/internal/app/routes"
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
	// set env vars; e.g. : os.Setenv("ENV_KEY", "value")

	app := fiber.New()

	// init prometheus handler
	reg := prometheus.NewRegistry()
	reg.MustRegister(metrics.DefaultCollectors()...)
	promHandler := promhttp.HandlerFor(reg, promhttp.HandlerOpts{Registry: reg})

	routes.RegisterHTTPRoutes(app, promHandler)
	ht.app = app

	return nil
}
