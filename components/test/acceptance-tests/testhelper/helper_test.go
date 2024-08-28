package testhelper

import (
	"os"
	"testing"

	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
	"github.com/stretchr/testify/assert"
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
		Paths:       []string{"specs/" + featureFile},
		Output:      colors.Colored(os.Stdout),
		Concurrency: 1,
		TestingT:    t,
	}
	actual := Options(t, featureFile)

	assert.Equal(expected, actual, "Options should be equal")
}

func Test_ShouldCleanupTargetDir(t *testing.T) {
	assert := assert.New(t)

	t.Run("Should remove target directory", func(t *testing.T) {
		err := cleanupTargetDir()
		assert.NoError(err, "Should not return an error")

		_, err = os.Stat("../../../target")
		assert.True(os.IsNotExist(err), "Target directory should not exist")
	})
}
