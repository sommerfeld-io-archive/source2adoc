package cmd

import (
	"fmt"

	"github.com/sommerfeld-io/source2adoc/internal/filesystem"
	"github.com/sommerfeld-io/source2adoc/internal/helper"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List for source code files that match a given pattern",
	Long:  `This command allows you to list all source files that match a given pattern in the current directory and all subdirectories.`,
	Run: func(cmd *cobra.Command, args []string) {
		lang, _ := cmd.Flags().GetString("lang")
		if !IsValidLanguage(lang) {
			HandleInvalidLang(lang)
			return
		}
		fmt.Println("Search source files for language:", lang)

		files, err := filesystem.FindFilesForLanguage(helper.CurrentWorkingDir(), lang)
		if err != nil {
			fmt.Println("Error finding files:", err)
			return
		}
		for _, file := range files {
			fmt.Println(file)
		}
	},
}

func init() {
	AddLangFlag(listCmd)
	rootCmd.AddCommand(listCmd)
}
