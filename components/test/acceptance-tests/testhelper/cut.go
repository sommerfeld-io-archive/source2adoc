// This file contains the `ContainerUnderTest` struct and its methods. The struct is used to configure and run the
// container under test.
package testhelper

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

// ContainerUnderTest represents the system under test.
//
// * image is the container image to use for the tests
// * cmd is the command to run in the container
// * context is the context to use for the container
// * req is the container request that configures the docker container
// * containerInstance represents the actual docker container
// * containerState is the state of the container (dedicated field to preserve state after the container is terminated)
// * volumes is a list of volumes to mount to the container as bind mounts (see https://docs.docker.com/engine/storage/bind-mounts)
type ContainerUnderTest struct {
	image             string
	cmd               []string
	ctx               context.Context
	req               testcontainers.ContainerRequest
	containerInstance testcontainers.Container
	containerState    *types.ContainerState
	volumes           []string
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
		image:   determineDockerImageToUse(),
		cmd:     []string{},
		ctx:     context.Background(),
		volumes: []string{},
	}
}

// AppendCommand appends a command to the command to run in the container.
func (c *ContainerUnderTest) AppendCommand(cmd ...string) {
	c.cmd = append(c.cmd, cmd...)
}

// CreateContainer configures and creates a container but does not start it. All config must be done
// here. To add volumes, use the `MountVolume` method.
func (c *ContainerUnderTest) CreateContainer() {
	c.req = testcontainers.ContainerRequest{
		Image:      c.image,
		Cmd:        c.cmd,
		WaitingFor: wait.ForExit(),
		HostConfigModifier: func(hostConfig *container.HostConfig) {
			hostConfig.Binds = c.volumes
		},
	}
}

// MountVolume add a configuration to mount a volume to the container request as bind mounts
// (see https://docs.docker.com/engine/storage/bind-mounts). The pathOnHost is mounted as source
// folder on the host and as target folder in the container keeping the paths the same.
//
// Call this method before or after `CreateContainer` but before `Run`.
func (c *ContainerUnderTest) MountVolume(pathOnHost string) {
	mount := pathOnHost + ":" + pathOnHost
	c.volumes = append(c.volumes, mount)
}

func (c *ContainerUnderTest) Run() error {
	instance, err := testcontainers.GenericContainer(c.ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: c.req,
		Started:          true,
	})
	if err != nil {
		return fmt.Errorf("failed to start container: %v", err)
	}
	defer instance.Terminate(c.ctx)

	c.containerInstance = instance
	c.containerState, err = c.containerInstance.State(c.ctx)
	if err != nil {
		return fmt.Errorf("failed to get container state: %v", err)
	}

	return nil
}

// ExitCode returns the exit code of the container after the container terminated.
func (c *ContainerUnderTest) ExitCode() (int, error) {
	return c.containerState.ExitCode, nil
}
