//go:build integration

package integration

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/remusxb/todo_crud/internal/app/handler"
	"github.com/remusxb/todo_crud/pkg/dto"
	"github.com/stretchr/testify/assert"
)

func TestConnect(t *testing.T) {
	t.Run("getHealth200", getHealth200)
}

func getHealth200(t *testing.T) {
	tcs := []struct {
		name       string
		wantStatus int
		wantBody   dto.HealthCheckOutput
		wantErr    error
	}{
		{
			name:       "status serving",
			wantStatus: http.StatusOK,
			wantBody:   dto.HealthCheckOutput{Message: handler.HealthMessageOK},
			wantErr:    nil,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodGet, "/health", nil)
			resp, err := ht.app.Test(request, -1) // -1 disables the response timeout

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
