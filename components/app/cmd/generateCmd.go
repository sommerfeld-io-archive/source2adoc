package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate AsciiDoc from inline documentation into an Antora component",
	Long:  `This command allows you to generate AsciiDoc files from your inline documentation into an Antora component.`,
	Run: func(cmd *cobra.Command, args []string) {
		lang, _ := cmd.Flags().GetString("lang")
		if !IsValidLanguage(lang) {
			HandleInvalidLang(lang)
			return
		}
		fmt.Println("Generating AsciiDoc for language:", lang)
		// Add service call here
	},
}

func init() {
	AddLangFlag(generateCmd)
	AddAntoraDirFlag(generateCmd)
	rootCmd.AddCommand(generateCmd)
}
