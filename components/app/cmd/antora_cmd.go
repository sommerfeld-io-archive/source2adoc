package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const antoraDescShort = "Generate artifacts and documentation pages specific to Antora modules."
const antoraDescLong = `
Generates artifacts and documentation pages, including a nav.adoc file,
specifically for Antora modules. It is designed to work with existing Antora
modules that already contain documentation.

You can use source2adoc to generate contents into an Antora module.

Example:
  source2adoc antora --module path/to/module

Example (Docker):
  docker run -v "$(pwd):$(pwd)" -w $(pwd) sommerfeldio/source2adoc:latest antora --module path/to/module
`

var moduleDir string

var antoraCmd = &cobra.Command{
	Use:   "antora",
	Short: antoraDescShort,
	Long:  antoraDescLong,

	Args: cobra.ExactArgs(0),

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Do stuff for module " + moduleDir)
	},
}

func init() {
	var moduleParam = "module"
	var moduleParamShort = "m"
	antoraCmd.Flags().StringVarP(&moduleDir, moduleParam, moduleParamShort, "", "Directory containing antora module")
	antoraCmd.MarkFlagRequired(moduleParam)

	RegisterSubCommand(antoraCmd)
}
