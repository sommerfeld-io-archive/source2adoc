package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVersion(t *testing.T) {
	expected := "latest"
	actual := Version()

	assert.Equal(t, expected, actual, "Git invalid Version")
}
