package filesystem

import (
	"testing"

	"github.com/sommerfeld-io/source2adoc/internal"
	"github.com/sommerfeld-io/source2adoc/internal/helper"
	"github.com/stretchr/testify/assert"
)

func TestFindFilesForLanguage(t *testing.T) {
	assert := assert.New(t)

	languages := internal.SupportedLangs()

	for _, lang := range languages {
		t.Run(lang, func(t *testing.T) {
			files, err := FindFilesForLanguage(helper.TestDataDir(), lang)
			assert.NotNil(files, "Should return files for language = "+lang)
			assert.NoError(err, "Should not return an error")
		})
	}
}
