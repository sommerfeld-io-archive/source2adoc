package testhelper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ShouldFindSourceCodeFiles(t *testing.T) {
	assert := assert.New(t)
	expectedFilesCount := 9

	files, err := FindSourceCodeFiles(TestDataPath)
	assert.NoError(err, "error should be nil but was", err)
	assert.NotNil(files, "files should not be nil")
	assert.Len(files, expectedFilesCount, "files count should be", expectedFilesCount)
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
