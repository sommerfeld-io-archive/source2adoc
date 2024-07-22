package codefiles

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCodeFileFinder_CodeShouldBeSupported(t *testing.T) {
	assert := assert.New(t)

	var tests = []string{
		"path/to/some.sh",
		"path/to/some.yml",
		"path/to/some.yaml",
		"path/to/Dockerfile",
		"path/to/Dockerfile.app",
		"path/to/Dockerfile.docs",
		"path/to/Dockerfile.dev",
		"path/to/Dockerfile.whatever",
		"path/to/Makefile",
		"path/to/Vagrantfile",
	}
	for _, test := range tests {
		t.Run(test, func(t *testing.T) {
			assert.True(isSupportedCode(test), "Should be supported: "+test)
		})
	}
}

func TestCodeFileFinder_CodeShouldNotBeSupported(t *testing.T) {
	assert := assert.New(t)

	var tests = []string{
		"path/to/some.go",
		"path/to/some.java",
		"path/to/some.kt",
		"path/to/some.json",
	}
	for _, test := range tests {
		t.Run(test, func(t *testing.T) {
			assert.False(isSupportedCode(test), "Should NOT be supported: "+test)
		})
	}
}

func TestCodeFileFinder_ShouldFindSourceCodeFiles(t *testing.T) {
	assert := assert.New(t)

	srcDir := "/workspaces/source2adoc/components/app/testdata"

	expectedFiles := []*CodeFile{
		NewCodeFile(filepath.Join(srcDir, "good/docker/Dockerfile"), "filename"),
		NewCodeFile(filepath.Join(srcDir, "good/docker/Dockerfile.app"), "filename"),
		NewCodeFile(filepath.Join(srcDir, "good/docker/Dockerfile.docs"), "filename"),
		NewCodeFile(filepath.Join(srcDir, "good/yaml/some.yml"), "filename"),
		NewCodeFile(filepath.Join(srcDir, "good/yaml/some.yaml"), "filename"),
		NewCodeFile(filepath.Join(srcDir, "good/Makefile"), "filename"),
		NewCodeFile(filepath.Join(srcDir, "good/Vagrantfile"), "filename"),
		NewCodeFile(filepath.Join(srcDir, "good/script.sh"), "filename"),
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
		NewCodeFile(filepath.Join(srcDir, "bad/some.go"), "filename"),
		NewCodeFile(filepath.Join(srcDir, "bad/some.java"), "filename"),
		NewCodeFile(filepath.Join(srcDir, "bad/some.kotlin"), "filename"),
	}

	finder := NewFinder(srcDir)

	files, err := finder.FindSourceCodeFiles()
	assert.NoError(err, "Should not return an error")
	assert.NotEqual(len(undesiredFiles), len(files), "Should return lists of different sizes")

	for _, undesiredFile := range undesiredFiles {
		assert.NotContains(files, undesiredFile, "Undesired file was found but should not")
	}
}
