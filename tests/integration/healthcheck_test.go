//go:build integration

package integration

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/remusxb/todo_crud/internal/app/handler"
	"github.com/remusxb/todo_crud/pkg/dto"
)

func TestConnect(t *testing.T) {
	t.Run("getHealth200", getHealth200)
}

func getHealth200(t *testing.T) {
	tcs := []struct {
		name       string
		wantStatus int
		wantBody   dto.HealthCheck
		wantErr    bool
	}{
		{
			name:       "status serving",
			wantStatus: http.StatusOK,
			wantBody:   dto.HealthCheck{Status: handler.Serving},
			wantErr:    false,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodGet, "/health", nil)
			resp, err := ht.app.Test(request, -1) // -1 disables the response timeout

			// Check for errors in the request
			if (err != nil) != tc.wantErr {
				t.Fatalf("Expected no error, but got %v", err)
			}

			// Check the HTTP status code
			if resp.StatusCode != tc.wantStatus {
				t.Fatalf("Expected status code %d, but got %d", tc.wantStatus, resp.StatusCode)
			}

			// Read the response body into a byte slice
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Fatalf("Error reading response body: %v", err)
			}

			// Parse the response JSON
			var responseBody dto.HealthCheck
			if err := json.Unmarshal(body, &responseBody); err != nil {
				t.Fatalf("Error unmarshaling JSON response: %v", err)
			}

			if !reflect.DeepEqual(responseBody, tc.wantBody) {
				t.Fatalf("Response body different than expected. Expected: %+v; got: %+v", tc.wantBody, responseBody)
			}
		})
	}
}
