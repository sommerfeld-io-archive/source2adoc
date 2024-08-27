package main

import (
	"context"
	"fmt"
	"testing"

	"github.com/cucumber/godog"
	"github.com/docker/docker/api/types"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

var containerCmd []string
var containerState *types.ContainerState

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

func initializeBasicScenario(sc *godog.ScenarioContext) {
	sc.Step(`^I am using the root command of the source2adoc CLI tool to generate AsciiDoc files$`, iAmUsingTheRootCommand)
	sc.Step(`^I specify "([^"]*)" using the --output-dir flag$`, iSpecifyTheOutputDir)
	sc.Step(`^I specify "([^"]*)" using the --source-dir flag$`, iSpecifyTheSourceDir)
	sc.Step(`^I specify the "([^"]*)" flag$`, iSpecifyTheFlag)
	sc.Step(`^I run the app$`, iRunTheApp)
	sc.Step(`^exit code should be (\d+)$`, exitCodeShouldBe)

	containerCmd = []string{}
	containerState = nil
}

func iAmUsingTheRootCommand() error {
	// root cmd does not require a dedicated command name
	containerCmd = append(containerCmd, "")
	return nil
}

func iSpecifyTheOutputDir(dir string) error {
	containerCmd = append(containerCmd, "--output-dir", dir)
	return nil
}

func iSpecifyTheSourceDir(dir string) error {
	containerCmd = append(containerCmd, "--source-dir", dir)
	return nil
}

func iSpecifyTheFlag(flag string) error {
	containerCmd = append(containerCmd, flag)
	return nil
}

func iRunTheApp() error {
	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		Image:      "sommerfeldio/source2adoc:rc",
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
