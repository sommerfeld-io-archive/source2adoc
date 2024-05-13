package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search for source code files that match a given pattern",
	Long:  `This command allows you to search for all source files that match a given pattern in the current directory.`,
	Run: func(cmd *cobra.Command, args []string) {
		lang, _ := cmd.Flags().GetString("lang")
		if !IsValidLanguage(lang) {
			HandleInvalidLang(lang)
			return
		}
		fmt.Println("Search source files for language:", lang)
		// Add service call here
	},
}

func init() {
	AddLangFlag(searchCmd)
	AddAntoraDirFlag(searchCmd)
	rootCmd.AddCommand(searchCmd)
}
