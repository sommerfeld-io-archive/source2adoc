package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_AntoraCmd(t *testing.T) {
	assert := assert.New(t)

	cmd := antoraCmd

	assert.Equal("antora", cmd.Use, "Incorrect command name")

	flags := cmd.Flags()
	assert.NotNil(flags.Lookup("module"), "Missing --module flag")
}
