package cmd

import (
	"fmt"

	"github.com/sommerfeld-io/source2adoc/internal"
	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate AsciiDoc from inline documentation into an Antora component",
	Long:  `This command allows you to generate AsciiDoc files from your inline documentation into an Antora component. The command scans the directory and all subdirectories.`,
	Run: func(cmd *cobra.Command, args []string) {
		lang, _ := cmd.Flags().GetString("lang")
		antoraDir, _ := cmd.Flags().GetString("antora-dir")
		antoraModule, _ := cmd.Flags().GetString("antora-module")

		if !IsValidLanguage(lang) {
			HandleInvalidLang(lang)
			return
		}

		files, err := internal.FindCodeFilesForLanguage(lang)
		if err != nil {
			fmt.Println("Error finding files:", err)
			return
		}
		for _, file := range files {

			err = internal.AppendToNavPartial(antoraDir, antoraModule, file)
			if err != nil {
				fmt.Println("Error writing nav entry:", err)
				return
			}

			adocPath, err := internal.WriteAdocFile(file, antoraDir, antoraModule)
			if err != nil {
				fmt.Println("Error writing AsciiDoc:", err)
				// continue
				return
			}

			fmt.Println("Generated AsciiDoc file:", adocPath)
		}

		internal.WriteIndexAdoc(antoraDir, antoraModule)
		internal.WriteNavAdoc(antoraDir, antoraModule)

		printAntoraYmlWarning()
	},
}

func init() {
	AddLangFlag(generateCmd)
	AddAntoraDirFlag(generateCmd)
	AddAntoraModuleNameFlag(generateCmd)
	rootCmd.AddCommand(generateCmd)
}

func printAntoraYmlWarning() {
	fmt.Println("================================================================")
	fmt.Println("  WARNING")
	fmt.Println("  Remember to add the generated module to the antora.yml file.")
	fmt.Println("================================================================")
}
