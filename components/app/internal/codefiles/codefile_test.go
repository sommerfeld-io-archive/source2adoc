package codefiles

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCodeFile_ShouldSplitPathAndFilename(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		path             string
		expectedPath     string
		expectedFilename string
	}{
		{path: "/path/to/source.sh", expectedPath: "/path/to", expectedFilename: "source.sh"},
		{path: "path/to/Dockerfile", expectedPath: "path/to", expectedFilename: "Dockerfile"},
		{path: "source.sh", expectedPath: "", expectedFilename: "source.sh"},
	}

	for _, test := range tests {
		path, filename := splitPathAndFilename(test.path)
		assert.Equal(test.expectedPath, path, "Incorrect path")
		assert.Equal(test.expectedFilename, filename, "Incorrect filename")
	}
}

func TestCodeFile_ShouldIdentifyLanguage(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		filename  string
		expected  string
		supported bool
	}{
		{filename: "config.yml", expected: LANGUAGE_YML, supported: true},
		{filename: "config.yaml", expected: LANGUAGE_YML, supported: true},
		{filename: "Dockerfile", expected: LANGUAGE_DOCKERFILE, supported: true},
		{filename: "Dockerfile.app", expected: LANGUAGE_DOCKERFILE, supported: true},
		{filename: "Dockerfile.docs", expected: LANGUAGE_DOCKERFILE, supported: true},
		{filename: "Vagrantfile.prod", expected: LANGUAGE_VAGRANT, supported: true},
		{filename: "Makefile", expected: LANGUAGE_MAKE, supported: true},
		{filename: "script.sh", expected: LANGUAGE_BASH, supported: true},
		{filename: "script.go", expected: LANGUAGE_INVALID, supported: false},
	}

	for _, test := range tests {
		lang, supported := identifyLanguage(test.filename)
		assert.Equal(test.expected, lang, "Incorrect language identification for: "+test.filename)
		assert.Equal(test.supported, supported, "Invalid supported status for: "+test.filename)
	}
}
func TestCodeFile_ShouldGetDataFromGetterFunctions(t *testing.T) {
	assert := assert.New(t)

	codeFile := &CodeFile{
		path:      "/path/to",
		name:      "source.sh",
		lang:      LANGUAGE_BASH,
		supported: true,
	}

	expectedPath := "/path/to"
	actualPath := codeFile.Path()
	assert.Equal(expectedPath, actualPath, "Incorrect path")

	expectedName := "source.sh"
	actualName := codeFile.Filename()
	assert.Equal(expectedName, actualName, "Incorrect filename")

	expectedLang := LANGUAGE_BASH
	actualLang := codeFile.Language()
	assert.Equal(expectedLang, actualLang, "Incorrect path language")

	expectedSupported := true
	actualSupported := codeFile.IsSupported()
	assert.Equal(expectedSupported, actualSupported, "Incorrect path supported status")
}

func TestCodeFile_ShouldReadFileContent(t *testing.T) {
	assert := assert.New(t)

	codeFile := &CodeFile{
		path:      filepath.Join(TEST_DATA_DIR, "good"),
		name:      "small-comment.sh",
		lang:      LANGUAGE_BASH,
		supported: true,
	}

	err := codeFile.ReadFileContent()
	assert.Nil(err, "Error reading file content")

	expectedContent := `#!/bin/bash
## Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt
## ut labore et dolore magna aliquyam erat, sed diam voluptua.

## Not part of the header comment
`
	actualContent := codeFile.Content()
	assert.Equal(expectedContent, actualContent, "Incorrect file content")
}

func TestCodeFile_ShouldParseHeaderDocsSection(t *testing.T) {
	assert := assert.New(t)

	codeFile := &CodeFile{
		path:      filepath.Join(TEST_DATA_DIR, "good"),
		name:      "small-comment.sh",
		lang:      LANGUAGE_BASH,
		supported: true,
		content: `#!/bin/bash
## Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt
# ignore me because I do not follow the comment convention. Maybe I am a typo.
## ut labore et dolore magna aliquyam erat, sed diam voluptua.

## Not part of the header comment
`,
		headerDocsSection: "",
	}

	expectedHeaderDocsSection := `Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt
ut labore et dolore magna aliquyam erat, sed diam voluptua.
`

	err := codeFile.ParseHeaderDocsSection()
	assert.Nil(err, "Error parsing header comment")

	actualHeaderDocsSection := codeFile.HeaderDocsSection()
	assert.Equal(expectedHeaderDocsSection, actualHeaderDocsSection, "Incorrect content from file header comment")
}
