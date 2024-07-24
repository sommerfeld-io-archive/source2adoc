package codefiles

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCodeFileFinder_ShouldFindSourceCodeFiles(t *testing.T) {
	assert := assert.New(t)

	srcDir := "/workspaces/source2adoc/components/app/testdata"

	expectedFiles := []*CodeFile{
		NewCodeFile(filepath.Join(srcDir, "good/docker/Dockerfile")),
		NewCodeFile(filepath.Join(srcDir, "good/docker/Dockerfile.app")),
		NewCodeFile(filepath.Join(srcDir, "good/docker/Dockerfile.docs")),
		NewCodeFile(filepath.Join(srcDir, "good/yaml/some.yml")),
		NewCodeFile(filepath.Join(srcDir, "good/yaml/some.yaml")),
		NewCodeFile(filepath.Join(srcDir, "good/Makefile")),
		NewCodeFile(filepath.Join(srcDir, "good/Vagrantfile")),
		NewCodeFile(filepath.Join(srcDir, "good/script.sh")),
	}

	finder := NewFinder(srcDir)

	files, err := finder.FindSourceCodeFiles()
	assert.NoError(err, "Should not return an error")
	assert.Equal(len(expectedFiles), len(files), "Should return the expected number of files")

	for _, expectedFile := range expectedFiles {
		assert.Contains(files, expectedFile, "Expected file not found")
	}
}

func TestCodeFileFinder_ShouldNotFindUndesiredSourceCodeFiles(t *testing.T) {
	assert := assert.New(t)

	srcDir := "/workspaces/source2adoc/components/app/testdata"

	undesiredFiles := []*CodeFile{
		NewCodeFile(filepath.Join(srcDir, "bad/some.go")),
		NewCodeFile(filepath.Join(srcDir, "bad/some.java")),
		NewCodeFile(filepath.Join(srcDir, "bad/some.kotlin")),
	}

	finder := NewFinder(srcDir)

	files, err := finder.FindSourceCodeFiles()
	assert.NoError(err, "Should not return an error")
	assert.NotEqual(len(undesiredFiles), len(files), "Should return lists of different sizes")

	for _, undesiredFile := range undesiredFiles {
		assert.NotContains(files, undesiredFile, "Undesired file was found but should not")
	}
}
