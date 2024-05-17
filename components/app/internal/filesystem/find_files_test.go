package filesystem

import (
	"strings"
	"testing"

	"github.com/sommerfeld-io/source2adoc/internal/helper"
	"github.com/stretchr/testify/assert"
)

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
			"path/to/some/bash/scripts/ansible-cli.sh",
			"path/to/some/bash/scripts/docker-stacks-cli.sh",
			"path/to/some/code/bash-with-functions.sh",
		}},
		{"Dockerfile*", []string{
			"Dockerfile",
			"path/Dockerfile",
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
			"path/to/some/ansible/desktops-main.yml",
			"path/to/some/ansible/desktops-steam.yml",
			"path/to/some/ansible/raspi-main.yml",
			"path/to/some/ansible/update-upgrade.yml",
		}},
		{"*.yaml", []string{
			"docker-compose.yaml",
		}},
	}
	for _, test := range tests {
		t.Run(test.pattern, func(t *testing.T) {
			paths, err := FindFilesByPattern(helper.TestDataDir(), test.pattern)
			paths = trimPaths(paths, helper.TestDataDir())

			assert.NoError(err, "Should not return an error")
			assert.Equal(test.expected, paths, "Should return correct files")
		})
	}
}

func TestPathWithoutCWD(t *testing.T) {
	startPath := "/path/to/start"
	path := "/path/to/start/file.txt"
	expected := "file.txt"

	result, err := pathWithoutCWD(startPath, path)

	assert.NoError(t, err, "Should not return an error")
	assert.Equal(t, expected, result, "Should return the expected path without CWD")
}
