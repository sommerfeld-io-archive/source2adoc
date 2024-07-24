package codefiles

import "strings"

// CodeFile represents a source code file in the file system.
type CodeFile struct {
	path      string
	name      string
	lang      string
	supported bool
}

const (
	LanguageYML         = "yml"
	LanguageDockerfile  = "Dockerfile"
	LanguageVagrantfile = "Vagrantfile"
	LanguageMakefile    = "Makefile"
	LanguageShellScript = "sh"
	LanguageInvalid     = "invalid"
)

// SupportedCodeFilenames maps supported file extensions to their corresponding languages.
var SupportedCodeFilenames = map[string]string{
	".yml":        LanguageYML,
	".yaml":       LanguageYML,
	"Dockerfile":  LanguageDockerfile,
	"Vagrantfile": LanguageVagrantfile,
	"Makefile":    LanguageMakefile,
	".sh":         LanguageShellScript,
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

// Split the path and filename
// If no "/" is found, return the entire path as the filename
func splitPathAndFilename(path string) (string, string) {
	lastIndex := strings.LastIndex(path, "/")
	if lastIndex == -1 {
		return "", path
	}
	return path[:lastIndex], path[lastIndex+1:]
}

// Identify the language of the file based on the filename or extension
// Return the language and a boolean indicating if the language is supported
func identifyLanguage(filename string) (string, bool) {
	for key, value := range SupportedCodeFilenames {
		if strings.HasSuffix(filename, key) || strings.HasPrefix(filename, value) {
			return value, true
		}
	}
	return LanguageInvalid, false
}
