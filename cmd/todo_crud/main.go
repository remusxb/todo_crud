//nolint:gochecknoglobals
package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	fiberRecover "github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	flag "github.com/spf13/pflag"

	"github.com/remusxb/todo_crud/internal/config"
	"github.com/remusxb/todo_crud/internal/metrics"
	"github.com/remusxb/todo_crud/internal/server"
	"github.com/remusxb/todo_crud/version"
)

var (
	shutdownSignals = []os.Signal{os.Interrupt, syscall.SIGTERM}
	logLevel        = new(slog.LevelVar) // Info by default
)

func initSlog() {
	h := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: logLevel})
	slog.SetDefault(slog.New(h))
}

func initConfig() *config.Config {
	cfg := &config.Config{}
	fs := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	config.Parse(fs, cfg, nil)

	return cfg
}

func initMiddlewares() []any {
	middlewares := make([]any, 0)
	recovery := fiberRecover.New(fiberRecover.Config{
		EnableStackTrace: false,
	})
	middlewares = append(middlewares, recovery)

	return middlewares
}

func main() {
	initSlog()
	logBuildInfo()

	cfg := initConfig()
	if cfg.Verbose {
		logLevel.Set(slog.LevelDebug)
	}

	cfg.FfConfig.Initialize()
	middlewares := initMiddlewares()
	rootCtx, cancel := context.WithCancel(context.Background())

	// init prometheus handler
	reg := prometheus.NewRegistry()
	reg.MustRegister(metrics.DefaultCollectors()...)
	promHandler := promhttp.HandlerFor(reg, promhttp.HandlerOpts{Registry: reg})

	// start server
	srv := server.InitServer(rootCtx, cfg.SrvConfig, promHandler, middlewares...)

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, shutdownSignals...)

	select {
	case s := <-interrupt:
		slog.Info(fmt.Sprintf("received interrupt signal: %s", s.String()))
	case err := <-srv.Notify():
		slog.Error(fmt.Sprintf("srv.Notify: %s", err))
	}

	cancel()

	// Shutdown
	if err := srv.Shutdown(); err != nil {
		slog.Error(fmt.Sprintf("srv.Shutdown: %s", err))
	} else {
		slog.Info("HTTP server exited properly")
	}
}

func logBuildInfo() {
	slog.Info("build information",
		slog.Group("build_info",
			"commit", version.GitCommit,
			"version", version.Version,
			"compiled_date", version.BuildDate,
			"go_version", version.GoVersion,
		),
	)
}
