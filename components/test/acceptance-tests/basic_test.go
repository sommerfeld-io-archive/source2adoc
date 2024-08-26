package main

import (
	"testing"

	"github.com/cucumber/godog"
)

func Test_BasicFeatures(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: initializeBasicScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{TestSpecsDir + "/basic.feature"},
			TestingT: t,
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}

func initializeBasicScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^AsciiDoc files should be generated for all source code files$`, asciiDocFilesShouldBeGeneratedForAllSourceCodeFiles)
	ctx.Step(`^I am using the root command of the source2adoc CLI tool to generate AsciiDoc files$`, iAmUsingTheRootCommand)
	ctx.Step(`^I run the app$`, iRunTheApp)
	ctx.Step(`^I specify "([^"]*)" using the --output-dir flag$`, iSpecifyTheOutputDir)
	ctx.Step(`^I specify "([^"]*)" using the --source-dir flag$`, iSpecifyTheSourceDir)
}

func iAmUsingTheRootCommand() error {
	// TODO prepare docker run command with volumes and images etc.
	return godog.ErrPending
}

func iRunTheApp() error {
	// TODO run the container with all docker flags and app flags
	return godog.ErrPending
}

func iSpecifyTheOutputDir(dir string) error {
	// TODO prepare --output-dir flag
	return godog.ErrPending
}

func iSpecifyTheSourceDir(dir string) error {
	// TODO prepare --source-dir flag
	return godog.ErrPending
}

func asciiDocFilesShouldBeGeneratedForAllSourceCodeFiles() error {
	// TODO check if all source code files have been converted to AsciiDoc files
	// TODO this is where the real validation should happen
	// TODO think about assert library to use
	return godog.ErrPending
}
