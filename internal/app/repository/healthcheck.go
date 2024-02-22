package repository

import (
	"context"
	"database/sql"
)

type HealthCheck struct {
	db *sql.DB
}

func NewHealthCheckRepository(db *sql.DB) HealthCheck {
	return HealthCheck{
		db: db,
	}
}

func (h HealthCheck) Ping(ctx context.Context) error {
	err := h.db.PingContext(ctx)
	if err != nil {
		return err
	}

	return nil
}
