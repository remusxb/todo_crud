package usecase

import (
	"context"
)

type HealthChecker interface {
	Ping(ctx context.Context) error
}

type HealthCheck struct {
	repository HealthChecker
}

func NewHealthCheckUseCase(repo HealthChecker) HealthCheck {
	return HealthCheck{
		repository: repo,
	}
}

func (h HealthCheck) PingDatabase(ctx context.Context) error {
	err := h.repository.Ping(ctx)
	if err != nil {
		return err
	}

	return nil
}
