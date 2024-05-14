package filesystem

import (
	"github.com/sommerfeld-io/source2adoc/internal"
)

func FindFilesForLanguage(currentDir string, lang string) ([]string, error) {

	pattern, err := internal.FileNamePatternForLanguage(lang)
	if err != nil {
		return nil, err
	}

	files, err := FindFilesByPattern(currentDir, pattern)
	if err != nil {
		return nil, err
	}
	return files, nil
}
