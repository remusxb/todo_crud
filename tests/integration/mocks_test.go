//go:build integration

package integration

import (
	"context"
	"errors"
)

type healthRepoMock struct {
	PingFunc func(ctx context.Context) error

	expectError error
}

func (rm *healthRepoMock) Ping(ctx context.Context) error {
	if rm.PingFunc == nil {
		return errors.New("repoMock.PingFunc: method is nil but repoMock.Ping was just called")
	}

	if rm.expectError != nil {
		return rm.expectError
	}

	return rm.Ping(ctx)
}
