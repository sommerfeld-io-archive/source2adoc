package codefiles

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test_ShouldFindSourceCodeFiles tests the case where the finder should return all
// supported code files that are found in the source directory.
func Test_ShouldFindSourceCodeFiles(t *testing.T) {
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
	assert.Empty(finder.exclude, "Exclude list should be empty")

	files, err := finder.FindSourceCodeFiles()
	assert.NoError(err, "Should not return an error")
	assert.Equal(len(expectedFiles), len(files), "Should return the expected number of files")

	for _, expectedFile := range expectedFiles {
		assert.Contains(files, expectedFile, "Expected file not found")
	}
}

// Test_ShouldNotFindUndesiredSourceCodeFiles tests the case where the finder should
// not return files that are unsupported code files.
func Test_ShouldNotFindUndesiredSourceCodeFiles(t *testing.T) {
	assert := assert.New(t)

	undesiredFiles := []*CodeFile{
		NewCodeFile(filepath.Join(TestSourceDir, "bad/some.go")),
		NewCodeFile(filepath.Join(TestSourceDir, "bad/some.java")),
		NewCodeFile(filepath.Join(TestSourceDir, "bad/some.kotlin")),
	}

	finder := NewFinder(TestSourceDir)
	assert.Empty(finder.exclude, "Exclude list should be empty")

	files, err := finder.FindSourceCodeFiles()
	assert.NoError(err, "Should not return an error")
	assert.NotEqual(len(undesiredFiles), len(files), "Should return lists of different sizes")

	for _, undesiredFile := range undesiredFiles {
		assert.NotContains(files, undesiredFile, "Undesired file was found but should not")
	}
}

// Test_ShouldFindSourceCodeFilesWithoutExcludes tests the case where the exclude
// list is not empty. The expected result is that the finder should not return the files that are
// in the exclude list but should still return all other files that should be found.
func Test_ShouldFindSourceCodeFilesWithoutExcludes(t *testing.T) {
	assert := assert.New(t)

	exclude := []string{
		filepath.Join(TestSourceDir, "good/docker/Dockerfile.docs"),
		filepath.Join(TestSourceDir, "good/yaml"),
		filepath.Join(TestSourceDir, "good/yaml/"),
		filepath.Join(TestSourceDir, "good/Makefile"),
	}

	expectedFiles := []*CodeFile{
		NewCodeFile(filepath.Join(TestSourceDir, "good/docker/Dockerfile")),
		NewCodeFile(filepath.Join(TestSourceDir, "good/docker/Dockerfile.app")),
		NewCodeFile(filepath.Join(TestSourceDir, "good/Vagrantfile")),
		NewCodeFile(filepath.Join(TestSourceDir, "good/small-comment.sh")),
		NewCodeFile(filepath.Join(TestSourceDir, "good/script.sh")),
	}

	finder := NewFinder(TestSourceDir)
	finder.SetExcludes(exclude)
	assert.NotEmpty(finder.exclude, "Exclude list should not be empty")

	files, err := finder.FindSourceCodeFiles()
	assert.NoError(err, "Should not return an error")
	assert.Equal(len(expectedFiles), len(files), "Should return the expected number of files")

	for _, expectedFile := range expectedFiles {
		assert.Contains(files, expectedFile, "Expected file not found")
	}
}
