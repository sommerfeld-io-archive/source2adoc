package main

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
	"github.com/docker/docker/api/types"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

var containerCmd []string
var containerState *types.ContainerState
var containerImage string

func Test_BasicFeatures(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: initializeBasicScenario,
		Options: &godog.Options{
			Format:      "pretty",
			Paths:       []string{TestSpecsDir + "/basic.feature"},
			Output:      colors.Colored(os.Stdout),
			Concurrency: 1,
			TestingT:    t,
		},
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
	sc.Step(`^exit code should be (\d+)$`, exitCodeShouldBe)

	sc.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		containerImage = "sommerfeldio/source2adoc:rc"
		containerCmd = []string{}
		containerState = nil
		return ctx, nil
	})
}

func iAmUsingTheRootCommand() error {
	// root cmd does not require a dedicated command name
	containerCmd = append(containerCmd, "")
	return nil
}

func iSpecifyTheFlag(flag string) error {
	containerCmd = append(containerCmd, flag)
	return nil
}

func iSpecifyTheFlagWithValue(flag, value string) error {
	containerCmd = append(containerCmd, flag, value)
	return nil
}

func iRunTheApp() error {
	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		Image:      containerImage,
		Cmd:        containerCmd,
		WaitingFor: wait.ForExit(),
	}

	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		return fmt.Errorf("Failed to start container: %v", err)
	}
	defer container.Terminate(ctx)

	containerState, err = container.State(ctx)
	if err != nil {
		return fmt.Errorf("Failed to get container state: %v", err)
	}

	return nil
}

func exitCodeShouldBe(expected int) error {
	if containerState.ExitCode != 0 {
		return fmt.Errorf("Expected exit code %d, got %d", expected, containerState.ExitCode)
	}
	return nil
}
