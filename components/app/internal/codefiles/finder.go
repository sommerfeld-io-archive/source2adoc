package codefiles

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
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

func isSupportedCode(path string) bool {
	filename := filepath.Base(path)
	ext := filepath.Ext(filename)

	switch ext {
	case ".yml", ".yaml", ".sh":
		return true
	}

	switch {
	case strings.HasPrefix(filename, "Dockerfile"):
		return true
	}

	switch filename {
	case "Makefile", "Vagrantfile":
		return true
	}

	return false
}

// FindSourceCodeFiles lists all files in srcDir and all subfolders.
func (finder *CodeFileFinder) FindSourceCodeFiles() ([]*CodeFile, error) {
	var files []*CodeFile

	err := filepath.Walk(finder.srcDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && isSupportedCode(path) {
			code := NewCodeFile(path, "filename")
			files = append(files, code)
		}

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to list files: %w", err)
	}

	return files, nil
}
