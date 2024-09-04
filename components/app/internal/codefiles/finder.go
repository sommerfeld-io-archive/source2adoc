package codefiles

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// CodeFileFinder is responsible for finding code files in a given directory.
type CodeFileFinder struct {
	srcDir  string
	exclude []string
}

// NewFinder creates a new CodeFileFinder instance.
func NewFinder(srcDir string) *CodeFileFinder {
	return &CodeFileFinder{
		srcDir:  srcDir,
		exclude: []string{},
	}
}

// SetExcludes sets the list of files and/or folders to exclude when generating documentation.
func (finder *CodeFileFinder) SetExcludes(excludes []string) {
	finder.exclude = excludes
}

// FindSourceCodeFiles lists all files in srcDir and all subfolders. It returns a list of supported code files.
// All paths from the exclude (with and without filename) are not part of the result.
func (finder *CodeFileFinder) FindSourceCodeFiles() ([]*CodeFile, error) {
	var files []*CodeFile

	err := filepath.Walk(finder.srcDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("failed to walk the filesystem: %v", err)
		}

		for _, exclude := range finder.exclude {
			if strings.Contains(path, exclude) {
				return nil
			}
		}

		if info.IsDir() {
			return nil
		}

		code := NewCodeFile(path)
		if code.IsSupportedLanguage() {
			files = append(files, code)
		}

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to list files: %w", err)
	}

	return files, nil
}
