package codefiles

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDocumentationPart_ShouldGetDataFromGetterFunctions(t *testing.T) {
	assert := assert.New(t)

	part := &DocumentationPart{
		sectionType:    DOCUMENTATION_PART_HEADER,
		sectionContent: "Lorem ipsum dolor sit amet",
	}

	expectedContent := "Lorem ipsum dolor sit amet"
	actualContent := part.SectionContent()
	assert.Equal(expectedContent, actualContent, "Incorrect section content")

	expectedType := DOCUMENTATION_PART_HEADER
	actualType := part.SectionType()
	assert.Equal(expectedType, actualType, "Incorrect section type")
}
