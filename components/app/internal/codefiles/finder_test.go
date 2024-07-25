package codefiles

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCodeFileFinder_ShouldFindSourceCodeFiles(t *testing.T) {
	assert := assert.New(t)

	expectedFiles := []*CodeFile{
		NewCodeFile(filepath.Join(TestDataDir, "good/docker/Dockerfile")),
		NewCodeFile(filepath.Join(TestDataDir, "good/docker/Dockerfile.app")),
		NewCodeFile(filepath.Join(TestDataDir, "good/docker/Dockerfile.docs")),
		NewCodeFile(filepath.Join(TestDataDir, "good/yaml/some.yml")),
		NewCodeFile(filepath.Join(TestDataDir, "good/yaml/some.yaml")),
		NewCodeFile(filepath.Join(TestDataDir, "good/Makefile")),
		NewCodeFile(filepath.Join(TestDataDir, "good/Vagrantfile")),
		NewCodeFile(filepath.Join(TestDataDir, "good/small-comment.sh")),
		NewCodeFile(filepath.Join(TestDataDir, "good/script.sh")),
	}

	finder := NewFinder(TestDataDir)

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
		NewCodeFile(filepath.Join(TestDataDir, "bad/some.go")),
		NewCodeFile(filepath.Join(TestDataDir, "bad/some.java")),
		NewCodeFile(filepath.Join(TestDataDir, "bad/some.kotlin")),
	}

	finder := NewFinder(TestDataDir)

	files, err := finder.FindSourceCodeFiles()
	assert.NoError(err, "Should not return an error")
	assert.NotEqual(len(undesiredFiles), len(files), "Should return lists of different sizes")

	for _, undesiredFile := range undesiredFiles {
		assert.NotContains(files, undesiredFile, "Undesired file was found but should not")
	}
}
