package filesystem

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func testDataDir() string {
	currentDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return filepath.Join(currentDir, "../../../testdata")
}

func trimPaths(paths []string, testDataDir string) []string {
	for i := range paths {
		paths[i] = strings.TrimPrefix(paths[i], testDataDir+"/")
	}
	return paths
}

func TestShouldFindFilesByPattern(t *testing.T) {
	assert := assert.New(t)

	var tests = []struct {
		pattern  string
		expected []string
	}{
		{"*.sh", []string{
			"bash.sh",
			"path/to/some/code/bash-with-functions.sh",
		}},
		{"Dockerfile*", []string{
			"Dockerfile",
			"path/to/some/more/code/Dockerfile",
			"path/to/some/more/code/Dockerfile.test",
		}},
		{"Makefile*", []string{
			"Makefile",
			"path/to/some/more/code/Makefile",
		}},
		{"*.rb", []string{
			"some/inspec/profile/controls/basic.rb",
		}},
		{"Vagrantfile*", []string{
			"Vagrantfile",
			"Vagrantfile.master",
			"Vagrantfile.slave",
			"path/to/some/more/code/Vagrantfile",
		}},
		{"*.yml", []string{
			"docker-compose.yml",
		}},
	}
	for _, test := range tests {
		t.Run(test.pattern, func(t *testing.T) {
			paths, err := FindFilesByPattern(testDataDir(), test.pattern)
			paths = trimPaths(paths, testDataDir())

			assert.NoError(err, "Should not return an error")
			assert.Equal(test.expected, paths, "Should return correct files")
		})
	}
}
