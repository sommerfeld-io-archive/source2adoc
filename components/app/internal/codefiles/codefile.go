package codefiles

type CodeFile struct {
	lang string
	path string
	name string
}

// New creates a new CodeFile instance.
func NewCodeFile(path, name string) *CodeFile {
	return &CodeFile{
		path: path,
		name: name,
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
