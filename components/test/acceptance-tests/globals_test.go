package main

import (
	"context"
	"os"
	"testing"

	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func Test_ShouldUseDockerImage(t *testing.T) {
	assert := assert.New(t)

	t.Run("Should use docker image from environment variable", func(t *testing.T) {
		os.Setenv("CONTAINER_IMAGE", "custom/image:tag")

		expected := "custom/image:tag"
		actual := determineDockerImageToUse()
		assert.Equal(expected, actual, "Should use docker image from environment variable")
	})

	t.Run("Should use default docker image", func(t *testing.T) {
		os.Unsetenv("CONTAINER_IMAGE")

		expected := "local/source2adoc:dev"
		actual := determineDockerImageToUse()
		assert.Equal(expected, actual, "Should use default docker image")
	})
}

// Test_ShouldUseCorrectOptions tests the Options function. Especially, the concurrency is
// important to ensure that the tests run in a single thread. If the tests run in parallel, the
// container setup and teardown and the state of the container under test might interfere with
// each other.
func Test_ShouldUseCorrectOptions(t *testing.T) {
	assert := assert.New(t)

	featureFile := "basic.feature"

	expected := &godog.Options{
		Format:      "pretty",
		Paths:       []string{TestSpecsDir + "/" + featureFile},
		Output:      colors.Colored(os.Stdout),
		Concurrency: 1,
		TestingT:    t,
	}
	actual := Options(t, featureFile)

	assert.Equal(expected, actual, "Options should be equal")
}

func Test_NewContainerUnderTest(t *testing.T) {
	assert := assert.New(t)

	t.Run("Should use docker image from environment variable", func(t *testing.T) {
		os.Setenv("CONTAINER_IMAGE", "custom/image:tag")
		expectedImage := "custom/image:tag"
		cut := NewContainerUnderTest()

		assert.Equal(expectedImage, cut.image, "Should set container image from environment variable")
		assert.Equal([]string{}, cut.cmd, "Should not set command")
		assert.Equal(context.Background(), cut.ctx, "Should not set context")
	})

	t.Run("Should use default docker image", func(t *testing.T) {
		os.Unsetenv("CONTAINER_IMAGE")
		expectedImage := "local/source2adoc:dev"
		cut = NewContainerUnderTest()

		assert.Equal(expectedImage, cut.image, "Should set default container image")
		assert.Equal([]string{}, cut.cmd, "Should not set command")
		assert.Equal(context.Background(), cut.ctx, "Should not set context")
	})

}

func Test_ShouldAppendCommand(t *testing.T) {
	assert := assert.New(t)

	c := NewContainerUnderTest()
	c.AppendCommand("command1", "command2")

	expected := []string{"command1", "command2"}
	actual := c.cmd

	assert.Equal(expected, actual, "Commands should be appended correctly")
}

func Test_ShouldCreateContainerRequest(t *testing.T) {
	assert := assert.New(t)

	c := NewContainerUnderTest()
	c.image = "custom/image:tag"
	c.AppendCommand("command1", "command2")
	c.CreateContainerRequest()

	expected := testcontainers.ContainerRequest{
		Image:      "custom/image:tag",
		Cmd:        []string{"command1", "command2"},
		WaitingFor: wait.ForExit(),
	}
	actual := c.req

	assert.Equal(expected, actual, "Container request should be created correctly")
}
