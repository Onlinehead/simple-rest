package tests

import (
	"os"
	"testing"
)

func SkipHTTPTest(t *testing.T) {
	if os.Getenv("SkipHTTPTest") != "" {
		t.Skip("Skipping testing in CI environment")
	}
}

func SkipPostgresTest(t *testing.T) {
	if os.Getenv("SkipPostgresTest") != "" {
		t.Skip("Skipping testing in CI environment")
	}
}

