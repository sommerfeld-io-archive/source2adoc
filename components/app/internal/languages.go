package internal

import (
	"fmt"
)

type LanguagePatternMapping struct {
	Language string
	Pattern  string
}

var languageMappings = []LanguagePatternMapping{
	{Language: "bash", Pattern: "*.sh"},
	{Language: "Dockerfile", Pattern: "Dockerfile*"},
	{Language: "Makefile", Pattern: "Makefile*"},
	{Language: "ruby", Pattern: "*.rb"},
	{Language: "Vagrantfile", Pattern: "Vagrantfile*"},
	{Language: "yml", Pattern: "*.yml"},
}

// GetPatternForLanguage returns a single Pattern from languageMappings for a given language.
// If the language is not found in the mappings, it returns an empty string.
func GetPatternForLanguage(lang string) (string, error) {
	for _, mapping := range languageMappings {
		if mapping.Language == lang {
			return mapping.Pattern, nil
		}
	}
	return "", fmt.Errorf("unsupported language %s", lang)
}

// SupportedLangs returns all supported languages to e.g. display in the command description.
func SupportedLangs() []string {
	languages := make([]string, len(languageMappings))
	for i, mapping := range languageMappings {
		languages[i] = mapping.Language
	}
	return languages
}
