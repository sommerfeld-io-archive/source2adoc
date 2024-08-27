package main

import (
	"context"
	"fmt"
	"testing"

	"github.com/cucumber/godog"
	"github.com/sommerfeld-io/source2adoc-acceptance-tests/testhelper"
)

var cut testhelper.ContainerUnderTest

func Test_BasicFeatures(t *testing.T) {
	opts := testhelper.Options(t, "basic.feature")

	suite := godog.TestSuite{
		ScenarioInitializer:  initializeBasicScenario,
		TestSuiteInitializer: testhelper.InitializeTestSuite,
		Options:              opts,
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}

func initializeBasicScenario(sc *godog.ScenarioContext) {
	sc.Step(`^I am using the root command of the source2adoc CLI tool to generate AsciiDoc files$`, iAmUsingTheRootCommand)
	sc.Step(`^I specify the "([^"]*)" flag$`, iSpecifyTheFlag)
	sc.Step(`^I specify the "([^"]*)" flag with value "([^"]*)"$`, iSpecifyTheFlagWithValue)
	sc.Step(`^I run the app$`, iRunTheApp)
	sc.Step(`^I run the app with volume mount "([^"]*)"$`, iRunTheAppWithVolumeMount)
	sc.Step(`^exit code should be (\d+)$`, exitCodeShouldBe)

	sc.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		cut = testhelper.NewContainerUnderTest()
		return ctx, nil
	})
}

func iAmUsingTheRootCommand() error {
	// The root cmd does not require a dedicated command name
	cut.AppendCommand("")
	return nil
}

func iSpecifyTheFlag(flag string) error {
	cut.AppendCommand(flag)
	return nil
}

func iSpecifyTheFlagWithValue(flag, value string) error {
	cut.AppendCommand(flag, value)
	return nil
}

func iRunTheApp() error {
	cut.CreateContainerRequest()
	err := cut.Run()
	if err != nil {
		return fmt.Errorf("failed to run container: %v", err)
	}

	return nil
}

func iRunTheAppWithVolumeMount(volumePath string) error {
	cut.CreateContainerRequest()
	cut.MountVolume(volumePath)
	err := cut.Run()
	if err != nil {
		return fmt.Errorf("failed to run container: %v", err)
	}

	return nil
}

func exitCodeShouldBe(expected int) error {
	code, err := cut.ExitCode()
	if err != nil {
		return fmt.Errorf("failed to get exit code: %v", err)
	}
	if code != 0 {
		return fmt.Errorf("expected exit code %d, got %d", expected, code)
	}
	return nil
}
