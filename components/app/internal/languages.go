package internal

import (
	"fmt"
	"sort"
)

var patterns = map[string]string{
	"bash":        "*.sh",
	"Dockerfile":  "Dockerfile*",
	"Makefile":    "Makefile*",
	"ruby":        "*.rb",
	"Vagrantfile": "Vagrantfile*",
	"yml":         "*.yml",
}

// FileNamePatternForLanguage takes a language as input and returns the corresponding file name
// pattern. If the language is not supported, it returns an error. Otherwise, it returns the
// corresponding pattern.
func FileNamePatternForLanguage(lang string) (string, error) {
	pattern, ok := patterns[lang]
	if !ok {
		return "", fmt.Errorf("unsupported language %s", lang)
	}

	return pattern, nil
}

// SupportedLangs returns all the supported languages to e.g. display in the command description.
func SupportedLangs() []string {
	keys := make([]string, 0, len(patterns))
	for key := range patterns {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}
