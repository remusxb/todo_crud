//go:build integration

package integration

import (
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	err := setup()
	if err != nil {
		log.Fatal(err.Error())
	}

	code := m.Run()

	os.Exit(code)
}

func setup() error {
	// set env vars if needed; e.g. : os.Setenv("ENV_KEY", "value")
	// or other dependencies

	return nil
}
