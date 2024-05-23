package cmd

import (
	"fmt"
	"strings"

	"github.com/sommerfeld-io/source2adoc/internal"
	"github.com/spf13/cobra"
)

// AddLangFlag adds a language flag to the given command.
// The language flag is used to specify the language of the source files.
// It is a required flag.
func AddLangFlag(cmd *cobra.Command) {
	langs := strings.Join(internal.SupportedLangs(), ", ")
	cmd.Flags().String("lang", "", fmt.Sprintf("Source code language (required). Allowed languages are: %s.", langs))
	cmd.MarkFlagRequired("lang")
}

// AddAntoraDirFlag adds the "antora-dir" flag to the given command.
// This flag allows specifying the directory for Antora, which is optional.
// The default value for the flag is "docs".
func AddAntoraDirFlag(cmd *cobra.Command) {
	cmd.Flags().String("antora-dir", "docs", "Directory for Antora (optional)")
}

// AddAntoraDirFlag adds the "antora-module" flag to the given command.
// This flag allows specifying the target module name (= directory) for the adoc files, which is optional.
// The default value for the flag is "source2adoc".
func AddAntoraModuleNameFlag(cmd *cobra.Command) {
	cmd.Flags().String("antora-module", "source2adoc", "Antora target module for the generated adoc files (optional)")
}

// IsValidLanguage checks if the given language is valid.
//
// Parameters:
// - lang: The language to be checked.
//
// Returns:
// - true if the language is valid, false otherwise.
func IsValidLanguage(lang string) bool {
	for _, b := range internal.SupportedLangs() {
		if b == lang {
			return true
		}
	}
	return false
}

// HandleInvalidLang prints an error message indicating that the provided language is invalid.
func HandleInvalidLang(lang string) {
	langs := strings.Join(internal.SupportedLangs(), ", ")
	errorMsg := fmt.Sprintf("Error: Invalid language. Allowed languages are: %s.", langs)
	fmt.Println(errorMsg)
}
