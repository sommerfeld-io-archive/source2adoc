package testhelper

import (
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
