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
	"yaml":        "*.yaml",
	"yml":         "*.yml",
}

func FileNamePatternForLanguage(lang string) (string, error) {
	pattern, ok := patterns[lang]
	if !ok {
		return "", fmt.Errorf("unsupported language %s", lang)
	}

	return pattern, nil
}

func SupportedLangs() []string {
	keys := make([]string, 0, len(patterns))
	for key := range patterns {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}
