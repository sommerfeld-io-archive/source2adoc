package metadata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVersion(t *testing.T) {
	expected := "dev"
	actual := Version()

	assert.Equal(t, expected, actual, "Git invalid Version")
}
