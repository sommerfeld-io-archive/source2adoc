package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldGetGetPatternForLanguage(t *testing.T) {
	assert := assert.New(t)

	var tests = []struct {
		lang     string
		expected string
	}{
		{"bash", "*.sh"},
		{"Dockerfile", "Dockerfile"},
		{"Makefile", "Makefile"},
		{"ruby", "*.rb"},
		{"Vagrantfile", "Vagrantfile"},
		{"yml", "*.yml"},
		{"yaml", "*.yaml"},
	}
	for _, test := range tests {
		t.Run(test.lang, func(t *testing.T) {
			pattern, err := GetPatternForLanguage(test.lang)

			assert.NoError(err, "Should not return an error")
			assert.NotEmpty(pattern, "Should return a pattern")
			assert.Equal(test.expected, pattern, "Should return correct file pattern")
		})
	}
}

func TestShouldGetErrorForInvalidLanguage(t *testing.T) {
	assert := assert.New(t)

	var tests = []struct {
		lang string
	}{
		{"java"},
		{"kotlin"},
	}
	for _, test := range tests {
		t.Run(test.lang, func(t *testing.T) {
			pattern, err := GetPatternForLanguage(test.lang)

			assert.Error(err, "Should return an error")
			assert.Empty(pattern, "Should not return a pattern")
		})
	}
}
func TestSupportedLangs(t *testing.T) {
	assert := assert.New(t)

	expected := []string{"Dockerfile", "Makefile", "Vagrantfile", "bash", "ruby", "yml", "yaml"}
	langs := SupportedLangs()

	assert.ElementsMatch(expected, langs, "Should return supported languages")
}
