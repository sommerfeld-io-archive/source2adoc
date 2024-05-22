package cmd

import (
	"fmt"

	"github.com/sommerfeld-io/source2adoc/internal/manpage"
	"github.com/spf13/cobra"
)

var manpageCmd = &cobra.Command{
	Use:   "manpage",
	Short: "Generate manpage in Asciidoc format into the current working directory",
	Long:  "This command generates a manpage in Asciidoc format into the current working directory. The `manpage.adoc` contains information about all commands.",
	Run: func(cmd *cobra.Command, args []string) {
		manpage.GenerateManpage(rootCmd)
		fmt.Println("[DONE] Manpage created in the current working directory as manpage.adoc")
	},
}

func init() {
	rootCmd.AddCommand(manpageCmd)
}
