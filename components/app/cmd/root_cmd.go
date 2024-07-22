package cmd

import (
	"fmt"
	"log"

	"github.com/sommerfeld-io/source2adoc/internal/codefiles"
	"github.com/spf13/cobra"
)

const rootDescShort = "Streamline the process of generating documentation from inline comments within source code files."
const rootDescLong = `
Facilitate the creation of comprehensive and well-structured documentation
directly from code comments. The app supports multiple source code languages.
The common ground is, that these languages mark their comments through the
hash-symbol (#).

For more information, visit the project's documentation:
  https://source2adoc.sommerfeld.io

Quick Start:
  The root command source2adoc [flags] scans the --source-dir for code files
  and starts the conversion process. The output is written to --output-dir.

Example:
  source2adoc --source-dir ./src --output-dir ./docs
`

var sourceDir string
var outputDir string

var rootCmd = &cobra.Command{
	Use:   "source2adoc",
	Short: rootDescShort,
	Long:  rootDescLong,

	Args: cobra.ExactArgs(0),

	Run: func(cmd *cobra.Command, args []string) {
		sourceCodeFiles, err := codefiles.NewFinder(sourceDir).FindSourceCodeFiles()
		if err != nil {
			log.Fatal(err)
		}

		for _, file := range sourceCodeFiles {
			fmt.Println(file)
		}
	},
}

func init() {
	var sourceParam = "source-dir"
	var sourceParamShort = "s"
	rootCmd.Flags().StringVarP(&sourceDir, sourceParam, sourceParamShort, "", "Directory containing the source code files")
	rootCmd.MarkFlagRequired(sourceParam)

	var outputParam = "output-dir"
	var outputParamShort = "o"
	rootCmd.Flags().StringVarP(&outputDir, outputParam, outputParamShort, "", "Directory to write the generated documentation to")
	rootCmd.MarkFlagRequired(outputParam)

	rootCmd.CompletionOptions.HiddenDefaultCmd = true
}

// Execute acts as the entrypoint for the command line interface.
func Execute() error {
	return rootCmd.Execute()
}

// RegisterSubCommand adds a subcommand to the root command.
func RegisterSubCommand(cmd *cobra.Command) {
	rootCmd.AddCommand(cmd)
}
