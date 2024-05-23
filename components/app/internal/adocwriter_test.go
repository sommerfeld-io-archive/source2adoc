package internal

import (
	"os"
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
