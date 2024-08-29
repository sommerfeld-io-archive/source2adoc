package testhelper

import (
	"os"
	"testing"

	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
	"github.com/stretchr/testify/assert"
)

// Test_ShouldUseCorrectOptions tests the Options function. Especially, the concurrency is
// important to ensure that the tests run in a single thread. If the tests run in parallel, the
// tests might interfere with each other due to conflicting file system access.
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
