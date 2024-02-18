//go:build unit

package unit

import (
	"testing"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"

	"github.com/remusxb/todo_crud/internal/app/handler"
)

func TestHTTPHealth_Health(t *testing.T) {
	app := fiber.New()
	c := app.AcquireCtx(&fasthttp.RequestCtx{})

	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: handler.Serving,
			args: args{
				c: c,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := handler.HealthCheck{}
			if err := h.Health(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("Health() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
