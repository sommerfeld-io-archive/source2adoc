package testhelper

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ShouldFindSourceCodeFiles(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		expectedCount int
		excludes      []string
	}{
		{
			expectedCount: 9,
			excludes:      []string{},
		},
		{
			expectedCount: 8,
			excludes:      []string{filepath.Join(TestDataPath, "script.sh")},
		},
		{
			expectedCount: 7,
			excludes:      []string{filepath.Join(TestDataPath, "yaml")},
		},
		{
			expectedCount: 7,
			excludes: []string{
				filepath.Join(TestDataPath, "Makefile"),
				filepath.Join(TestDataPath, "Vagrantfile"),
			},
		},
	}

	for _, test := range tests {
		files, err := FindSourceCodeFiles(TestDataPath, test.excludes)

		assert.NoError(err, "error should be nil but was", err)
		assert.NotNil(files, "files should not be nil")
		assert.Len(files, test.expectedCount, "files count should be", test.expectedCount)
	}
}

func Test_ShouldTranslateFilename(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		filename     string
		expectedName string
	}{
		{
			filename:     "example.sh",
			expectedName: "example-sh.adoc",
		},
		{
			filename:     "file.name.with.dots",
			expectedName: "file-name-with-dots.adoc",
		},
		{
			filename:     "UPPERCASE",
			expectedName: "uppercase.adoc",
		},
		{
			filename:     "DOckerfile.app",
			expectedName: "dockerfile-app.adoc",
		},
	}

	for _, test := range tests {
		actualName := TranslateFilename(test.filename)
		assert.Equal(test.expectedName, actualName, "filename translation is incorrect")
	}
}
func Test_ShouldFindInString(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		needle   string
		haystack string
		expected error
	}{
		{
			needle:   "needle",
			haystack: "This is a haystack",
			expected: fmt.Errorf("needle needle was not found in haystack"),
		},
		{
			needle:   "This is",
			haystack: "This is a haystack",
			expected: nil,
		},
		{
			needle:   "123",
			haystack: "456789",
			expected: fmt.Errorf("needle 123 was not found in haystack"),
		},
	}

	for _, test := range tests {
		err := findInString(test.needle, test.haystack)
		assert.Equal(test.expected, err, "expected error does not match")
	}
}
