package cmd

import (
	"github.com/spf13/cobra"
)

const rootDescShort = "Streamline the process of generating documentation from inline comments within source code files."
const rootDescLong = `
The primary objective of source2adoc is to facilitate the creation of
comprehensive and well-structured documentation directly from code comments.
By leveraging the familiar syntax of inline comments in a style similar to
JavaDoc, developers can annotate their code, ensuring that insights and
explanations are captured and preserved in the generated AsciiDoc files.

The app supports multiple source code languages. The common ground is, that
these languages mark their comments through the hash-symbol (#).
`

var rootCmd = &cobra.Command{
	Use:   "source2adoc",
	Short: rootDescShort,
	Long:  rootDescLong,

	Args: cobra.ExactArgs(0),
}

func init() {
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
