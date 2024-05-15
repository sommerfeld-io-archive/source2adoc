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

	expectedContent := "\n" +
		"== testCommand\n" +
		"Test command synopsis\n" +
		"\n" +
		"[source, bash]\n" +
		"....\n" +
		"testCommand [options]\n" +
		"....\n" +
		"\n" +
		"Test command description\n\n" +
		"* `option1`, `o`, default = default1 +\n" +
		"  Option 1 usage\n" +
		"* `option2`, default = _none_ +\n" +
		"  Option 2 usage\n"

	assert.Equal(t, expectedContent, string(fileContent))
}
