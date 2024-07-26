package codefiles

// DocumentationPart represents a part of the documentation of a CodeFile. The sum
// of all DocumentationParts represents a documentation page.
type DocumentationPart struct {
	sectionType    string
	sectionContent string
}

// SectionContent returns the type of the DocumentationPart to distinguish between header
// docs, meta information and function docs, etc.
func (part *DocumentationPart) SectionType() string {
	return part.sectionType
}

// SectionContent returns the content of the DocumentationPart (= the actual docs for e.g.
// the function).
func (part *DocumentationPart) SectionContent() string {
	return part.sectionContent
}
