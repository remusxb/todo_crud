package server

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/bytedance/sonic"
	fiber "github.com/gofiber/fiber/v2"

	"github.com/remusxb/todo_crud/internal/app/routes"
)

// Server -.
type Server struct {
	HTTP            *fiber.App
	ctx             context.Context
	notify          chan error
	addr            string
	shutdownTimeout time.Duration
}

func New(ctx context.Context, cfg Config) *Server {
	app := fiber.New(fiber.Config{
		Prefork:      false,
		AppName:      "todo_crud",
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
		JSONEncoder:  sonic.Marshal,
		JSONDecoder:  sonic.Unmarshal,
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			// Status code defaults to 500
			code := fiber.StatusInternalServerError

			// Retrieve the custom status code if it's a *fiber.Error
			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}

			// log the error
			slog.Error(
				err.Error(),
				"request_info", slog.GroupValue(
					slog.String("protocol", c.Protocol()),
					slog.String("method", c.Method()),
					slog.String("path", c.Path()),
					slog.Int("status_code", code),
				),
			)
			// Set Content-Type: text/plain; charset=utf-8
			c.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)

			// Return status code with error message
			return c.Status(code).JSON(err)
		},
	})

	s := Server{
		HTTP:            app,
		ctx:             ctx,
		notify:          make(chan error, 1),
		addr:            fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		shutdownTimeout: cfg.ShutdownTimeout,
	}

	return &s
}

func (s *Server) Init(router routes.Router, middlewares ...any) {
	// middlewares
	for _, middleware := range middlewares {
		s.HTTP.Use(middleware)
	}

	router.RegisterHTTPRoutes()
	s.start()
}

func (s *Server) start() {
	go func() {
		s.notify <- s.HTTP.Listen(s.addr)
		close(s.notify)
	}()
}

// Notify -.
func (s *Server) Notify() <-chan error {
	return s.notify
}

// Shutdown -.
func (s *Server) Shutdown() error {
	<-s.ctx.Done()
	slog.Info("stopping HTTP server")

	return s.HTTP.ShutdownWithTimeout(s.shutdownTimeout)
}
