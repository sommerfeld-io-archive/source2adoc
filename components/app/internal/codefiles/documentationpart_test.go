package codefiles

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ShouldGetDataFromSectionGetters(t *testing.T) {
	assert := assert.New(t)

	part := &DocumentationPart{
		sectionType:    DocumentationPartHeader,
		sectionContent: "Lorem ipsum dolor sit amet",
	}

	expectedContent := "Lorem ipsum dolor sit amet"
	actualContent := part.SectionContent()
	assert.Equal(expectedContent, actualContent, "Incorrect section content")

	expectedType := DocumentationPartHeader
	actualType := part.SectionType()
	assert.Equal(expectedType, actualType, "Incorrect section type")
}
