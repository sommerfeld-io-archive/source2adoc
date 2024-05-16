package metadata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVersion(t *testing.T) {
	assert := assert.New(t)

	actual := Version()
	assert.NotEmpty(actual)
	assert.NotEqual("UNSPECIFIED", actual)
}
