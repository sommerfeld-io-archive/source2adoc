package codefiles

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCodeFileFinder_ShouldFindSourceCodeFiles(t *testing.T) {
	assert := assert.New(t)

	expectedFiles := []*CodeFile{
		NewCodeFile(filepath.Join(TestSourceDir, "good/docker/Dockerfile")),
		NewCodeFile(filepath.Join(TestSourceDir, "good/docker/Dockerfile.app")),
		NewCodeFile(filepath.Join(TestSourceDir, "good/docker/Dockerfile.docs")),
		NewCodeFile(filepath.Join(TestSourceDir, "good/yaml/some.yml")),
		NewCodeFile(filepath.Join(TestSourceDir, "good/yaml/some.yaml")),
		NewCodeFile(filepath.Join(TestSourceDir, "good/Makefile")),
		NewCodeFile(filepath.Join(TestSourceDir, "good/Vagrantfile")),
		NewCodeFile(filepath.Join(TestSourceDir, "good/small-comment.sh")),
		NewCodeFile(filepath.Join(TestSourceDir, "good/script.sh")),
	}

	finder := NewFinder(TestSourceDir)

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
		NewCodeFile(filepath.Join(TestSourceDir, "bad/some.go")),
		NewCodeFile(filepath.Join(TestSourceDir, "bad/some.java")),
		NewCodeFile(filepath.Join(TestSourceDir, "bad/some.kotlin")),
	}

	finder := NewFinder(TestSourceDir)

	files, err := finder.FindSourceCodeFiles()
	assert.NoError(err, "Should not return an error")
	assert.NotEqual(len(undesiredFiles), len(files), "Should return lists of different sizes")

	for _, undesiredFile := range undesiredFiles {
		assert.NotContains(files, undesiredFile, "Undesired file was found but should not")
	}
}
