package manpage

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAppendCommandDocsToAdoc(t *testing.T) {
	// Create a temporary file for testing
	tempFile, err := os.Create("test_manpage.adoc")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tempFile.Name())

	sampleCommandDocs := CommandDocs{
		Name:        "testCommand",
		Synopsis:    "Test command synopsis",
		Description: "Test command description",
		Usage:       "testCommand [options]",
		Options: []OptionDocs{
			{
				Name:         "option1",
				DefaultValue: "default1",
				Usage:        "Option 1 usage",
				Shorthand:    "o",
			},
			{
				Name:         "option2",
				DefaultValue: "default2",
				Usage:        "Option 2 usage",
				Shorthand:    "t",
			},
		},
	}

	appendCommandDocsToAdoc(sampleCommandDocs, tempFile.Name())

	fileContent, err := os.ReadFile(tempFile.Name())
	if err != nil {
		t.Fatal(err)
	}

	expectedContent := `== testCommand
Test command synopsis

Test command description
[source, bash]
....
testCommand [options]
....

[options="header"]
|===
|Flag Name |Shorthand |Desc |Default Value
|option1
|o
|Option 1 usage
|default1
|option2
|t
|Option 2 usage
|default2
|===
`

	assert.Equal(t, expectedContent, string(fileContent))
}
