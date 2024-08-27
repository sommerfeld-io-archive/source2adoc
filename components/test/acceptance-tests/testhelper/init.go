// This file contains the configuration for the godog test suite.
package testhelper

import (
	"fmt"
	"os"

	"github.com/cucumber/godog"
)

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
		fmt.Print(textGrey)
		fmt.Println("Run acceptance tests against the container image:", determineDockerImageToUse())
		fmt.Println(textWhite)
	})
}
