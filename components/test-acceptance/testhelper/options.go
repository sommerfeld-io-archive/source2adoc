package testhelper

import (
	"os"
	"testing"

	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
)

// Options returns the godog options for the test suite. The options are configured centrally to ensure consistency
// across all tests.
//
// Concurrently is limited to 1 because the file system is shared between the tests and the container a does not
// support concurrent writes (files from different tests might be written at the same time and interfere with each other).
func Options(t *testing.T, featureFile string) *godog.Options {
	return &godog.Options{
		Format:      "pretty",
		Paths:       []string{"specs/" + featureFile},
		Output:      colors.Colored(os.Stdout),
		Concurrency: 1,
		TestingT:    t,
	}
}
