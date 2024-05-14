package internal

import "fmt"

func FileNamePatternForLanguage(lang string) (string, error) {
	patterns := map[string]string{
		"bash":        "*.sh",
		"Dockerfile":  "Dockerfile*",
		"Makefile":    "Makefile*",
		"ruby":        "*.rb",
		"Vagrantfile": "Vagrantfile*",
		"yaml":        "*.yaml",
		"yml":         "*.yml",
	}

	pattern, ok := patterns[lang]
	if !ok {
		return "", fmt.Errorf("unsupported language %s", lang)
	}

	return pattern, nil
}
