package codefiles

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCodeFileFinder_ShouldFindSourceCodeFiles(t *testing.T) {
	assert := assert.New(t)

	expectedFiles := []*CodeFile{
		NewCodeFile(filepath.Join(TEST_DATA_DIR, "good/docker/Dockerfile")),
		NewCodeFile(filepath.Join(TEST_DATA_DIR, "good/docker/Dockerfile.app")),
		NewCodeFile(filepath.Join(TEST_DATA_DIR, "good/docker/Dockerfile.docs")),
		NewCodeFile(filepath.Join(TEST_DATA_DIR, "good/yaml/some.yml")),
		NewCodeFile(filepath.Join(TEST_DATA_DIR, "good/yaml/some.yaml")),
		NewCodeFile(filepath.Join(TEST_DATA_DIR, "good/Makefile")),
		NewCodeFile(filepath.Join(TEST_DATA_DIR, "good/Vagrantfile")),
		NewCodeFile(filepath.Join(TEST_DATA_DIR, "good/small-comment.sh")),
		NewCodeFile(filepath.Join(TEST_DATA_DIR, "good/script.sh")),
	}

	finder := NewFinder(TEST_DATA_DIR)

	files, err := finder.FindSourceCodeFiles()
	assert.NoError(err, "Should not return an error")
	assert.Equal(len(expectedFiles), len(files), "Should return the expected number of files")

	for _, expectedFile := range expectedFiles {
		assert.Contains(files, expectedFile, "Expected file not found")
	}
}

func TestCodeFileFinder_ShouldNotFindUndesiredSourceCodeFiles(t *testing.T) {
	assert := assert.New(t)

	undesiredFiles := []*CodeFile{
		NewCodeFile(filepath.Join(TEST_DATA_DIR, "bad/some.go")),
		NewCodeFile(filepath.Join(TEST_DATA_DIR, "bad/some.java")),
		NewCodeFile(filepath.Join(TEST_DATA_DIR, "bad/some.kotlin")),
	}

	finder := NewFinder(TEST_DATA_DIR)

	files, err := finder.FindSourceCodeFiles()
	assert.NoError(err, "Should not return an error")
	assert.NotEqual(len(undesiredFiles), len(files), "Should return lists of different sizes")

	for _, undesiredFile := range undesiredFiles {
		assert.NotContains(files, undesiredFile, "Undesired file was found but should not")
	}
}
