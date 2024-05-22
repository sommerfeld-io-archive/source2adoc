package cmd

import (
	"github.com/spf13/cobra"
)

var lintCmd = &cobra.Command{
	Use:   "lint",
	Short: "Lint source code files for compliance with source2adoc rules",
	Long:  `This command allows you to lint your source files and check if the comments comply with the rules of the source2adoc app. The command scans the directory and all subdirectories.`,
	Run: func(cmd *cobra.Command, args []string) {
		lang, _ := cmd.Flags().GetString("lang")
		if !IsValidLanguage(lang) {
			HandleInvalidLang(lang)
			return
		}
		// Add service call here
		// Print some success / failure message
	},
}

func init() {
	AddLangFlag(lintCmd)
	AddAntoraDirFlag(lintCmd)
	rootCmd.AddCommand(lintCmd)
}
