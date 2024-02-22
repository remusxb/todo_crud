//go:build integration

package integration

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/stretchr/testify/assert"

	"github.com/remusxb/todo_crud/internal/app/handler"
	"github.com/remusxb/todo_crud/internal/app/routes"
	"github.com/remusxb/todo_crud/internal/app/usecase"
	"github.com/remusxb/todo_crud/internal/metrics"
	"github.com/remusxb/todo_crud/pkg/dto"
)

func TestConnect(t *testing.T) {
	app := fiber.New()
	repoMock := &healthRepoMock{}

	// init use cases
	healthUseCase := usecase.NewHealthCheckUseCase(repoMock)

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

	tcs := []struct {
		name       string
		wantStatus int
		wantBody   dto.HealthCheckOutput
		wantErr    error
		PingFunc   func(ctx context.Context) error
	}{
		{
			name:       "status serving",
			wantStatus: http.StatusOK,
			wantBody:   dto.HealthCheckOutput{Message: handler.HealthMessageOK},
			wantErr:    nil,
			PingFunc: func(ctx context.Context) error {
				return nil
			},
		},
		{
			name:       "database not ready",
			wantStatus: http.StatusOK,
			wantBody: dto.HealthCheckOutput{
				Message: handler.HealthMessageNotOK,
				Error:   "failed to ping",
			},
			wantErr: nil,
			PingFunc: func(ctx context.Context) error {
				return errors.New("failed to ping")
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			repoMock.PingFunc = tc.PingFunc
			request := httptest.NewRequest(http.MethodGet, "/health", nil)
			resp, err := app.Test(request, 1)

			// Check for errors in the request
			assert.Equal(t, tc.wantErr, err)
			assert.Equal(t, tc.wantStatus, resp.StatusCode)

			// Read the response body into a byte slice
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Fatalf("Error reading response body: %v", err)
			}

			// Parse the response JSON
			var responseBody dto.HealthCheckOutput
			if err := json.Unmarshal(body, &responseBody); err != nil {
				t.Fatalf("Error unmarshaling JSON response: %v", err)
			}

			assert.EqualValues(t, tc.wantBody, responseBody)
		})
	}
}
