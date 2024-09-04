package cmd

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func Test_RootCmd(t *testing.T) {
	assert := assert.New(t)

	cmd := rootCmd

	assert.Equal("source2adoc", cmd.Use, "Incorrect command name")

	flags := cmd.Flags()
	assert.NotNil(flags.Lookup("source-dir"), "Missing --source-dir flag")
	assert.NotNil(flags.Lookup("output-dir"), "Missing --output-dir flag")
	assert.NotNil(flags.Lookup("exclude"), "Missing --exclude flag")
}

func Test_ShouldGetExcludes(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		expected []string
	}{
		{expected: []string{}},
		{expected: []string{"file1"}},
		{expected: []string{"file1", "file2"}},
	}

	for _, test := range tests {
		cmd := &cobra.Command{}
		cmd.Flags().StringSlice("exclude", test.expected, "Exclude files and/or folders when generating documentation")

		excludes, err := getExcludes(cmd)
		assert.Nil(err, "Error should be nil")

		assert.Equal(test.expected, excludes, "Incorrect excludes")
	}
}
