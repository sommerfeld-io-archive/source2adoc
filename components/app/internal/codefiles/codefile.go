package codefiles

import (
	"os"
	"strings"
)

// SupportedCodeFilenames maps supported file extensions to their corresponding languages.
var SupportedCodeFilenames = map[string]string{
	".yml":        LANGUAGE_YML,
	".yaml":       LANGUAGE_YML,
	"Dockerfile":  LANGUAGE_DOCKERFILE,
	"Vagrantfile": LANGUAGE_VAGRANT,
	"Makefile":    LANGUAGE_MAKE,
	".sh":         LANGUAGE_BASH,
}

// CodeFile represents a source code file in the file system.
type CodeFile struct {
	path               string
	name               string
	lang               string
	supported          bool
	fileContent        string
	documentationParts []DocumentationPart
}

// New creates a new CodeFile instance.
func NewCodeFile(fullPath string) *CodeFile {
	path, name := splitPathAndFilename(fullPath)
	lang, supported := identifyLanguage(name)

	return &CodeFile{
		path:      path,
		name:      name,
		lang:      lang,
		supported: supported,
	}
}

// Path returns the path of the CodeFile.
func (cf *CodeFile) Path() string {
	return cf.path
}

// Filename returns the name of the CodeFile.
func (cf *CodeFile) Filename() string {
	return cf.name
}

// Language returns the language of the CodeFile.
func (cf *CodeFile) Language() string {
	return cf.lang
}

// IsSupported returns true if the CodeFile is supported.
func (cf *CodeFile) IsSupported() bool {
	return cf.supported
}

// FileContent returns the content of the CodeFile.
func (cf *CodeFile) FileContent() string {
	return cf.fileContent
}

// ReadFileContent reads the content of the CodeFile from the file system.
func (cf *CodeFile) ReadFileContent() error {
	content, err := os.ReadFile(cf.path + "/" + cf.name)
	if err != nil {
		return err
	}
	cf.fileContent = string(content)
	return nil
}

// Parse parses the CodeFile and extracts the documentation parts.
func (cf *CodeFile) Parse() error {
	// TODO err := cf.parseMetadata()
	err := cf.parseHeaderDocs()
	return err
}

// parsedDocumentation returns the parsed documentation of the CodeFile. The parsed documentation
// is ready to be written to a file.
func (cf *CodeFile) parsedDocumentation() string {
	parsedDocs := ""
	for _, part := range cf.documentationParts {
		parsedDocs += part.sectionContent
	}
	return parsedDocs
}

// parseHeaderDocs finds all relevant comments (marked with `##`) at the beginning of the file
// and stores them in the CodeFile.
//
// See "Rules for the header documentation" in `docs/modules/ROOT/pages/index.adoc`.
func (cf *CodeFile) parseHeaderDocs() error {
	headerDocs := ""
	lines := strings.Split(cf.fileContent, "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "##") {
			trimmedLine := strings.TrimPrefix(line, "##")
			trimmedLine = strings.TrimSpace(trimmedLine)
			headerDocs += trimmedLine + "\n"
		} else if line == "" {
			break
		}
	}

	part := DocumentationPart{
		sectionType:    DOCUMENTATION_PART_HEADER,
		sectionContent: headerDocs,
	}
	cf.documentationParts = append(cf.documentationParts, part)

	return nil
}

// Split the path and filename
// If no "/" is found, return the entire path as the filename
// TODO turn into a method of CodeFile
func splitPathAndFilename(path string) (string, string) {
	lastIndex := strings.LastIndex(path, "/")
	if lastIndex == -1 {
		return "", path
	}
	return path[:lastIndex], path[lastIndex+1:]
}

// Identify the language of the file based on the filename or extension
// Return the language and a boolean indicating if the language is supported
// TODO turn into a method of CodeFile
func identifyLanguage(filename string) (string, bool) {
	for key, value := range SupportedCodeFilenames {
		if strings.HasSuffix(filename, key) || strings.HasPrefix(filename, value) {
			return value, true
		}
	}
	return LANGUAGE_INVALID, false
}
