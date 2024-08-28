package main

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cucumber/godog"
	"github.com/sommerfeld-io/source2adoc-acceptance-tests/testhelper"
)

type BasicTestState struct {
	cut       testhelper.ContainerUnderTest
	sourceDir string
	outputDir string
}

func (ts *BasicTestState) reset() {
	ts.cut = testhelper.NewContainerUnderTest()
	ts.sourceDir = ""
	ts.outputDir = ""
}

func Test_BasicFeatures(t *testing.T) {
	featureFile := "basic.feature"
	opts := testhelper.Options(t, featureFile)

	suite := godog.TestSuite{
		Name:                 featureFile,
		ScenarioInitializer:  initializeBasicScenario,
		TestSuiteInitializer: testhelper.InitializeTestSuite,
		Options:              opts,
	}

	exitcode := suite.Run()
	if exitcode != 0 {
		t.Fatal(suite.Name, "|", "non-zero status returned.", "failed to run tests.", "finished with exit code", exitcode)

	}
}

func initializeBasicScenario(sc *godog.ScenarioContext) {
	ts := &BasicTestState{}

	sc.Step(`^I use the root command of the source2adoc CLI tool to generate AsciiDoc files$`, ts.iUseTheRootCommand)
	sc.Step(`^I specify the "([^"]*)" flag$`, ts.iSpecifyTheFlag)
	sc.Step(`^I specify the "([^"]*)" flag with value "([^"]*)"$`, ts.iSpecifyTheFlagWithValue)
	sc.Step(`^I run the app$`, ts.iRunTheApp)
	sc.Step(`^I run the app with volume mount "([^"]*)"$`, ts.iRunTheAppWithVolumeMount)
	sc.Step(`^exit code should be (\d+)$`, ts.exitCodeShouldBe)
	sc.Step(`^no AsciiDoc files should be generated$`, ts.noAsciiDocFilesShouldBeGenerated)

	sc.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		ts.reset()
		return ctx, nil
	})

	sc.After(testhelper.AfterScenario)
}

func (ts *BasicTestState) iUseTheRootCommand() error {
	// The root cmd does not require a dedicated command name
	return nil
}

func (ts *BasicTestState) iSpecifyTheFlag(flag string) error {
	ts.cut.AppendCommand(flag)
	return nil
}

func (ts *BasicTestState) iSpecifyTheFlagWithValue(flag, value string) error {
	if flag == "--source-dir" {
		ts.sourceDir = value
	}
	if flag == "--output-dir" {
		ts.outputDir = value
	}

	ts.cut.AppendCommand(flag, value)
	return nil
}

func (ts *BasicTestState) iRunTheApp() error {
	ts.cut.CreateContainer()
	err := ts.cut.Run()
	if err != nil {
		return fmt.Errorf("failed to run container: %v", err)
	}
	return nil
}

func (ts *BasicTestState) iRunTheAppWithVolumeMount(pathOnHost string) error {
	ts.cut.CreateContainer()
	ts.cut.MountVolume(pathOnHost)
	err := ts.cut.Run()
	if err != nil {
		return fmt.Errorf("failed to run container: %v", err)
	}
	return nil
}

func (ts *BasicTestState) exitCodeShouldBe(expected int) error {
	code, err := ts.cut.ExitCode()
	if err != nil {
		return fmt.Errorf("failed to get exit code: %v", err)
	}
	if code != expected {
		return fmt.Errorf("expected exit code %d, got %d", expected, code)
	}
	return nil
}

func (ts *BasicTestState) noAsciiDocFilesShouldBeGenerated() error {
	if ts.outputDir == "" {
		return fmt.Errorf("output directory not set")
	}

	if _, err := os.Stat(ts.outputDir); !os.IsNotExist(err) {
		return fmt.Errorf(ts.outputDir, "directory should not exist")
	}
	return nil
}
