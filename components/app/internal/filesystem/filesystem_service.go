package filesystem

import (
	"github.com/sommerfeld-io/source2adoc/internal"
)

// FindFilesForLanguage serves as a service layer that acts as a bridge between the `cobra`
// commands and the actual implementations.
func FindFilesForLanguage(currentDir string, lang string) ([]string, error) {
	pattern, err := internal.GetPatternForLanguage(lang)
	if err != nil {
		return nil, err
	}

	files, err := FindFilesByPattern(currentDir, pattern)
	if err != nil {
		return nil, err
	}
	return files, nil
}
