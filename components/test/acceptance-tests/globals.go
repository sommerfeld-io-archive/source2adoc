// This file contains global variables and godog hooks that are used in all acceptance test files.
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/cucumber/godog"
	"github.com/docker/docker/api/types"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

// TestSpecsDir is the directory where the feature files are located.
const TestSpecsDir = "specs"

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

// ContainerUnderTest represents the system under test.
//
// * image is the container image to use for the tests
// * cmd is the command to run in the container
// * context is the context to use for the container
// * req is the container request that configures the docker container
// * containerInstance represents the actual docker container
// * containerState is the state of the container (dedicated field to preserve state after the container is terminated)
type ContainerUnderTest struct {
	image             string
	cmd               []string
	ctx               context.Context
	req               testcontainers.ContainerRequest
	containerInstance testcontainers.Container
	containerState    *types.ContainerState
}

func determineDockerImageToUse() string {
	env := os.Getenv("CONTAINER_IMAGE")
	if env != "" {
		return env
	}

	return "local/source2adoc:dev"
}

// NewContainerUnderTest returns a container request for the system under test
//
// It sets the container image to the value of the `CONTAINER_IMAGE` environment variable if it is
// set. Otherwise, it sets it to the default value `local/source2adoc:dev` which means the test
// cases are run against a local container image. If the local image is not present, the test cases
// will fail. So you need to build  the image before running the tests.
//
// The reason behind this is that a pipeline can test against different versions of the app
// by setting the CONTAINER_IMAGE environment variable while the developer can easily run the
// tests locally through `go test` or the IDE.
//
// To set a different container image, use `CONTAINER_IMAGE=sommerfeldio/source2adoc:rc go test`.
func NewContainerUnderTest() ContainerUnderTest {
	return ContainerUnderTest{
		image: determineDockerImageToUse(),
		cmd:   []string{},
		ctx:   context.Background(),
	}
}

// Context returns the context to use for the container.
func (c *ContainerUnderTest) Context() context.Context {
	return c.ctx
}

// ContainerImage returns the docker image to use for the container that should be tested.
func (c *ContainerUnderTest) ContainerImage() string {
	return c.image
}

// Command returns the command to run in the container.
func (c *ContainerUnderTest) Command() []string {
	return c.cmd
}

// AppendCommand appends a command to the command to run in the container.
func (c *ContainerUnderTest) AppendCommand(cmd ...string) {
	c.cmd = append(c.cmd, cmd...)
}

// CreateContainerRequest creates a container request that configures the docker container.
func (c *ContainerUnderTest) CreateContainerRequest() {
	c.req = testcontainers.ContainerRequest{
		Image:      c.ContainerImage(),
		Cmd:        c.Command(),
		WaitingFor: wait.ForExit(),
	}
}

// MountVolume mounts a volume to the container request. The volumePath is mounted as source folder
// on the host and as target folder in the container.
func (c *ContainerUnderTest) MountVolume(volumePath string) {
	c.req.Mounts = testcontainers.ContainerMounts{
		{
			Source: testcontainers.GenericVolumeMountSource{
				Name: "test-volume",
			},
			Target: testcontainers.ContainerMountTarget(volumePath),
		},
	}
}

func (c *ContainerUnderTest) Run() error {
	instance, err := testcontainers.GenericContainer(c.Context(), testcontainers.GenericContainerRequest{
		ContainerRequest: c.req,
		Started:          true,
	})
	if err != nil {
		return fmt.Errorf("failed to start container: %v", err)
	}
	defer instance.Terminate(c.Context())

	c.containerInstance = instance
	c.containerState, err = c.containerInstance.State(c.Context())
	if err != nil {
		return fmt.Errorf("failed to get container state: %v", err)
	}

	return nil
}

// ExitCode returns the exit code of the container after the container terminated.
func (c *ContainerUnderTest) ExitCode() (int, error) {
	return c.containerState.ExitCode, nil
}
