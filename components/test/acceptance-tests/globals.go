// This file contains global variables and godog hooks that are used in all acceptance test files.
package main

import (
	"fmt"
	"os"

	"github.com/cucumber/godog"
)

// TestSpecsDir is the directory where the feature files are located.
const TestSpecsDir = "specs"

// ContainerImage is the container image to use for the tests. This var is used in every test
// that runs its tests against a container.
var ContainerImage string

// InitializeTestSuite is a godog hook that runs before the test suite starts. The function
// should be used for all acceptance tests. It sets the ContainerImage to the value of the
// `CONTAINER_IMAGE` environment variable if it is set. Otherwise, it sets it to the default
// value `local/source2adoc:dev` which means the test cases are run against a local container
// image. If the local image is not present, the test cases will fail. So you need to build
// the image before running the tests.
//
// The reason behind this is that a pipeline can test against different versions of the app
// by setting the CONTAINER_IMAGE environment variable while the developer can easily run the
// tests locally through `go test` or the IDE.
//
// To set a different container image, use `CONTAINER_IMAGE=sommerfeldio/source2adoc:rc go test`.
func InitializeTestSuite(sc *godog.TestSuiteContext) {
	textGrey := "\033[90m"
	textWhite := "\033[0m"

	envVarName := "CONTAINER_IMAGE"
	defaultContainerImage := "local/source2adoc:dev"

	sc.BeforeSuite(func() {
		fmt.Println(textGrey)

		if os.Getenv(envVarName) != "" {
			ContainerImage = os.Getenv(envVarName)

			fmt.Println("Running acceptance tests against container image specified in environment variable", envVarName)
			fmt.Println("Using", ContainerImage)
		} else {
			ContainerImage = defaultContainerImage

			fmt.Println("No container image specified through environment variable.")
			fmt.Println("Defaulting to", ContainerImage)
		}

		fmt.Println(textWhite)
	})
}
