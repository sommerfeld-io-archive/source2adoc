package internal

import (
	"os"
	"path/filepath"
)

// FindCodeFilesForLanguage returns all files in the current directory and all subdirectories that
// match the given language.
func FindCodeFilesForLanguage(lang string) ([]string, error) {
	pattern, err := GetPatternForLanguage(lang)
	if err != nil {
		return nil, err
	}

	var files []string
	err = filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		matched, err := filepath.Match(pattern, filepath.Base(path))
		if err != nil {
			return err
		}
		if matched {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return files, nil
}
