package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var lintCmd = &cobra.Command{
	Use:   "lint",
	Short: "Lint source code files for compliance with source2adoc rules",
	Long:  `This command allows you to lint your source files and check if the comments comply with the rules of the source2adoc app.`,
	Run: func(cmd *cobra.Command, args []string) {
		lang, _ := cmd.Flags().GetString("lang")
		if !IsValidLanguage(lang) {
			HandleInvalidLang(lang)
			return
		}
		fmt.Println("Linting source files for language:", lang)
		// Add service call here
	},
}

func init() {
	AddLangFlag(lintCmd)
	AddAntoraDirFlag(lintCmd)
	rootCmd.AddCommand(lintCmd)
}
