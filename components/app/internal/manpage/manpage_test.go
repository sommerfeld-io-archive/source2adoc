package manpage

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAppendCommandDocsToAdoc(t *testing.T) {
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
				DefaultValue: "",
				Usage:        "Option 2 usage",
				Shorthand:    "",
			},
		},
	}

	appendCommandDocsToAdoc(sampleCommandDocs, tempFile.Name())

	fileContent, err := os.ReadFile(tempFile.Name())
	if err != nil {
		t.Fatal(err)
	}

	expectedContent := `
== testCommand
Test command synopsis

[source, bash]
....
testCommand [options]
....

Test command description

* option1, o, default = default1 +
  Option 1 usage
* option2, default = _none_ +
  Option 2 usage
`

	assert.Equal(t, expectedContent, string(fileContent))
}
