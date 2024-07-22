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
		NewCodeFile(filepath.Join(srcDir, "docker/Dockerfile"), "filename"),
		NewCodeFile(filepath.Join(srcDir, "docker/Dockerfile.app"), "filename"),
		NewCodeFile(filepath.Join(srcDir, "docker/Dockerfile.docs"), "filename"),
		NewCodeFile(filepath.Join(srcDir, "yaml/some.yml"), "filename"),
		NewCodeFile(filepath.Join(srcDir, "yaml/some.yaml"), "filename"),
		NewCodeFile(filepath.Join(srcDir, "Makefile"), "filename"),
		NewCodeFile(filepath.Join(srcDir, "Vagrantfile"), "filename"),
	}

	finder := NewFinder(srcDir)

	files, err := finder.FindSourceCodeFiles()
	assert.NoError(err, "Should not return an error")
	assert.Equal(len(expectedFiles), len(files), "Should return the expected number of files")

	for _, expectedFile := range expectedFiles {
		assert.Contains(files, expectedFile, "Expected file not found")
	}
}
