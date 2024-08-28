// This file contains the configuration for the godog test suite.
package testhelper

import (
	"context"
	"fmt"
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

// determineDockerImageToUse determines the container image to use for the tests based on the
// `CONTAINER_IMAGE` environment variable. If the variable is not set, the default image
// `local/source2adoc:dev` is used.
func determineDockerImageToUse() string {
	env := os.Getenv("CONTAINER_IMAGE")
	if env != "" {
		return env
	}

	return "local/source2adoc:dev"
}

// InitializeTestSuite is a godog hook that is called before the test suite is run.
// It prints a message to the console to inform the user about the container image that is used for the tests.
func InitializeTestSuite(sc *godog.TestSuiteContext) {
	textGrey := "\033[90m"
	textWhite := "\033[0m"

	sc.BeforeSuite(func() {
		img := determineDockerImageToUse()
		cmd := "CONTAINER_IMAGE=" + img

		fmt.Print(textGrey)
		fmt.Println("Run acceptance tests against the container image:", img)
		fmt.Println("Reproduce with:", cmd, "go test")
		fmt.Println(textWhite)
	})
}

func cleanupTargetDir() error {
	targetDir := "../../target"
	err := os.RemoveAll(targetDir)
	if err != nil {
		return fmt.Errorf("error removing target directory: %v", err)
	}
	return nil
}

// AfterScenario is a godog hook that is called after each scenario. It handles the cleanup after each scenario
// to avoid side effects between scenarios due leftovers from previous scenarios.
func AfterScenario(ctx context.Context, sc *godog.Scenario, err error) (context.Context, error) {
	if err != nil {
		return ctx, fmt.Errorf("error in scenario: %v", err)
	}

	err = cleanupTargetDir()
	if err != nil {
		return ctx, fmt.Errorf("error cleaning up target directory: %v", err)
	}
	return ctx, nil
}
