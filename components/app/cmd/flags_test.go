package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValidLanguage(t *testing.T) {
	assert := assert.New(t)

	t.Run("should pass", func(t *testing.T) {
		assert.True(IsValidLanguage("bash"))
		assert.True(IsValidLanguage("Dockerfile"))
		assert.True(IsValidLanguage("Makefile"))
		assert.True(IsValidLanguage("ruby"))
		assert.True(IsValidLanguage("Vagrantfile"))
		assert.True(IsValidLanguage("yml"))
		assert.True(IsValidLanguage("yaml"))
	})

	t.Run("should fail", func(t *testing.T) {
		assert.False(IsValidLanguage("go"))
		assert.False(IsValidLanguage("java"))
		assert.False(IsValidLanguage("kotlin"))
	})
}
