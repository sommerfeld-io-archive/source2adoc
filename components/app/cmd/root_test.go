package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_RootCmd(t *testing.T) {
	assert := assert.New(t)

	cmd := rootCmd

	assert.Equal("source2adoc", cmd.Use, "Incorrect command name")

	flags := cmd.Flags()
	assert.NotNil(flags.Lookup("source-dir"), "Missing --source-dir flag")
	assert.NotNil(flags.Lookup("output-dir"), "Missing --output-dir flag")
}
