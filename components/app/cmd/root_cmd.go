package cmd

import (
	"strings"

	"github.com/sommerfeld-io/source2adoc/internal"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "source2adoc",
	Version: internal.Version(),
	Short:   "Generate AsciiDoc from inline documentation",
	Long:    strings.Trim(`Convert inline documentation into AsciiDoc files, tailored for seamless integration with Antora.`, " "),

	Args: cobra.ExactArgs(0),
}

func init() {
	rootCmd.CompletionOptions.HiddenDefaultCmd = true
}

// Execute acts as the entrypoint for the command line interface.
func Execute() error {
	return rootCmd.Execute()
}
