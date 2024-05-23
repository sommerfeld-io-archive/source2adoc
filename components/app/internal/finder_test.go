package internal

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindCodeFilesForLanguage(t *testing.T) {
	assert := assert.New(t)

	var tests = []string{
		"yml",
		"ruby",
		"bash",
		"Dockerfile",
		"Vagrantfile",
		"Makefile",
	}
	for _, test := range tests {
		t.Run(test, func(t *testing.T) {
			expectedPattern, err := GetPatternForLanguage(test)
			assert.NoError(err, "Should not return an error")
			assert.NotEmpty(expectedPattern, "Should return a pattern")

			expectedPattern = strings.ReplaceAll(expectedPattern, "*", "")

			files, err := FindCodeFilesForLanguage(test)
			assert.NoError(err, "Should not return an error")
			assert.NotEmpty(files, "Should return files")
			for _, file := range files {
				assert.True(strings.HasSuffix(file, expectedPattern))
			}
		})
	}
}
