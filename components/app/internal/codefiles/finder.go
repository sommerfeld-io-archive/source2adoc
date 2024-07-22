package codefiles

import (
	"fmt"
	"os"
	"path/filepath"
)

// CodeFileFinder is responsible for finding code files in a given directory.
type CodeFileFinder struct {
	srcDir string
}

// NewFinder creates a new CodeFileFinder instance.
func NewFinder(srcDir string) *CodeFileFinder {
	return &CodeFileFinder{
		srcDir: srcDir,
	}
}

// FindSourceCodeFiles lists all files in srcDir and all subfolders.
// TODO scan for *.sh, *.yml, *.yaml, Dockerfile*, Vagrantfile, Makefile
func (f *CodeFileFinder) FindSourceCodeFiles() ([]*CodeFile, error) {
	var files []*CodeFile

	err := filepath.Walk(f.srcDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			code := NewCodeFile(path, "filename") // TODO determine language etc. (autodetect filename from path)
			files = append(files, code)
		}

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to list files: %w", err)
	}

	return files, nil
}
