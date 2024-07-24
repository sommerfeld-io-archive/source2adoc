package codefiles

import (
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
		{filename: "config.yml", expected: LanguageYML, supported: true},
		{filename: "config.yaml", expected: LanguageYML, supported: true},
		{filename: "Dockerfile", expected: LanguageDockerfile, supported: true},
		{filename: "Dockerfile.app", expected: LanguageDockerfile, supported: true},
		{filename: "Dockerfile.docs", expected: LanguageDockerfile, supported: true},
		{filename: "Vagrantfile.prod", expected: LanguageVagrantfile, supported: true},
		{filename: "Makefile", expected: LanguageMakefile, supported: true},
		{filename: "script.sh", expected: LanguageShellScript, supported: true},
		{filename: "script.go", expected: LanguageInvalid, supported: false},
	}

	for _, test := range tests {
		lang, supported := identifyLanguage(test.filename)
		assert.Equal(test.expected, lang, "Incorrect language identification for: "+test.filename)
		assert.Equal(test.supported, supported, "Invalid supported status for: "+test.filename)
	}
}
func TestCodeFile_Path(t *testing.T) {
	codeFile := &CodeFile{
		path:      "/path/to",
		name:      "source.sh",
		lang:      LanguageShellScript,
		supported: true,
	}

	expectedPath := "/path/to"
	actualPath := codeFile.Path()
	assert.Equal(t, expectedPath, actualPath, "Incorrect path")

	expectedName := "source.sh"
	actualName := codeFile.Filename()
	assert.Equal(t, expectedName, actualName, "Incorrect filename")

	expectedLang := LanguageShellScript
	actualLang := codeFile.Language()
	assert.Equal(t, expectedLang, actualLang, "Incorrect path language")

	expectedSupported := true
	actualSupported := codeFile.IsSupported()
	assert.Equal(t, expectedSupported, actualSupported, "Incorrect path supported status")
}
