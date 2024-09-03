package main

import (
	"os"
	"testing"

	"github.com/sommerfeld-io/source2adoc/internal/codefiles"
)

func cleanup() error {
	return os.RemoveAll(codefiles.TestOutputDir)
}

// TestMain is the entry point for the test suite and allows you to set up code that runs before
// and after your tests.
func TestMain(m *testing.M) {
	err := cleanup()
	if err != nil {
		panic(err)
	}

	exitCode := m.Run()
	os.Exit(exitCode)
}
