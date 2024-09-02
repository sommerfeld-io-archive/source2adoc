package testhelper

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// FindSourceCodeFiles finds source code files in the given directory and its subdirectories that
// are supported by the app.
func FindSourceCodeFiles(sourceDir string) ([]string, error) {
	var codeFiles []string

	err := filepath.Walk(sourceDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		fileName := info.Name()
		if strings.HasPrefix(fileName, "Dockerfile") ||
			strings.HasSuffix(fileName, ".yml") ||
			strings.HasSuffix(fileName, ".yaml") ||
			strings.HasPrefix(fileName, "Makefile") ||
			strings.HasPrefix(fileName, "Vagrantfile") ||
			strings.HasSuffix(fileName, ".sh") {
			codeFiles = append(codeFiles, path)
		}

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to find source code files: %v", err)
	}

	return codeFiles, nil
}

// TranslateFilename translates the given filename to a valid AsciiDoc filename.
//
// @see components/app/internal/codefiles/codefile.go::documentationFileName()
func TranslateFilename(filename string) string {
	name := strings.ReplaceAll(filename, ".", "-")
	name = strings.ToLower(name)
	return name + ".adoc"
}
