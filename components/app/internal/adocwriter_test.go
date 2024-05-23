package internal

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateAdocPath(t *testing.T) {
	assert := assert.New(t)

	codeFilePath := "path/to/file.yml"
	antoraDir := "docs"
	antoraModule := "source2adoc"

	expectedAdocPath := antoraDir + "/modules/" + antoraModule + "/pages/path/to/file-yml.adoc"

	adocPath := generateAdocPath(codeFilePath, antoraDir, antoraModule)

	assert.Equal(expectedAdocPath, adocPath)
}

func TestGenerateTitle(t *testing.T) {
	assert := assert.New(t)

	var tests = []string{
		"file.yml",
		"path/to/file.yml",
	}
	for _, test := range tests {
		t.Run(test, func(t *testing.T) {
			expectedTitle := "file.yml"
			assert.Equal(expectedTitle, generateTitle(test))
		})
	}
}

func TestWriteAdocFile(t *testing.T) {
	assert := assert.New(t)

	codeFilePath := "path/to/file.yml"
	antoraDir := "../../../docs"
	antoraModule := "source2adoc-test"

	expectedAdocPath := antoraDir + "/modules/" + antoraModule + "/pages/path/to/file-yml.adoc"

	adocPath, err := WriteAdocFile(codeFilePath, antoraDir, antoraModule)

	assert.NoError(err)
	assert.Equal(expectedAdocPath, adocPath)

	err = os.RemoveAll("../../../docs/modules/source2adoc-test")
	assert.NoError(err)
}

func TestAppendToNavPartial(t *testing.T) {
	assert := assert.New(t)

	antoraDir := "path/to/antora"
	antoraModule := "source2adoc-test"
	codeFilePath := "path/to/file.yml"

	expectedNavAdocPartialPath := antoraDir + "/modules/" + antoraModule + "/partials/nav.adoc"
	expectedXref := "* xref:" + antoraModule + ":path/to/file-yml.adoc[" + codeFilePath + "]"

	err := AppendToNavPartial(antoraDir, antoraModule, codeFilePath)

	assert.NoError(err, "AppendToNavPartial should not return an error")

	// Verify that the nav.adoc partial file exists
	_, err = os.Stat(expectedNavAdocPartialPath)
	assert.False(os.IsNotExist(err), "nav.adoc partial file should exist")

	// Verify that the link is appended to the nav.adoc partial file
	fileContent, err := os.ReadFile(expectedNavAdocPartialPath)
	assert.NoError(err, "reading nav.adoc partial file should not return an error")

	content := string(fileContent)
	content = strings.ReplaceAll(content, "\n", "")
	assert.Contains(content, expectedXref, "nav.adoc partial file should contain the link")

	// Clean up the generated files and directories
	err = os.RemoveAll(antoraDir)
	assert.NoError(err, "cleaning up the generated files should not return an error")
}

func TestWriteNavAdoc(t *testing.T) {
	assert := assert.New(t)

	antoraDir := "path/to/antora"
	antoraModule := "source2adoc-test"

	expectedNavAdocPath := antoraDir + "/modules/" + antoraModule + "/nav.adoc"
	expectedXref := "* xref:" + antoraModule + ":index.adoc[]"

	err := WriteNavAdoc(antoraDir, antoraModule)

	assert.NoError(err, "WriteNavAdoc should not return an error")

	// Verify that the nav.adoc file exists
	_, err = os.Stat(expectedNavAdocPath)
	assert.False(os.IsNotExist(err), "nav.adoc file should exist")

	// Verify that the xref is written to the nav.adoc file
	fileContent, err := os.ReadFile(expectedNavAdocPath)
	assert.NoError(err, "reading nav.adoc file should not return an error")

	content := string(fileContent)
	content = strings.ReplaceAll(content, "\n", "")
	assert.Contains(content, expectedXref, "nav.adoc file should contain the xref")

	// Clean up the generated files and directories
	err = os.RemoveAll(antoraDir)
	assert.NoError(err, "cleaning up the generated files should not return an error")
}

func TestWriteIndexAdoc(t *testing.T) {
	assert := assert.New(t)

	antoraDir := "path/to/antora"
	antoraModule := "source2adoc-test"

	expectedIndexAdocPath := antoraDir + "/modules/" + antoraModule + "/pages/index.adoc"
	expectedInclude := "include::" + antoraModule + ":partial$nav.adoc[]"

	err := WriteIndexAdoc(antoraDir, antoraModule)

	assert.NoError(err, "WriteIndexAdoc should not return an error")

	// Verify that the index.adoc file exists
	_, err = os.Stat(expectedIndexAdocPath)
	assert.False(os.IsNotExist(err), "index.adoc file should exist")

	// Verify that the include is written to the index.adoc file
	fileContent, err := os.ReadFile(expectedIndexAdocPath)
	assert.NoError(err, "reading index.adoc file should not return an error")

	content := string(fileContent)
	content = strings.ReplaceAll(content, "\n", "")
	assert.Contains(content, expectedInclude, "index.adoc file should contain the include")

	// Clean up the generated files and directories
	err = os.RemoveAll(antoraDir)
	assert.NoError(err, "cleaning up the generated files should not return an error")
}
