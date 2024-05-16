package helper

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCurrentWorkingDir(t *testing.T) {
	expected, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	dir := CurrentWorkingDir()

	assert.Equal(t, expected, dir, "Should return the correct current working directory")
}
