package manpage

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/sommerfeld-io/source2adoc/internal"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
	"gopkg.in/yaml.v3"
)

var manpageFilename = "manpage.adoc"
var dockerRun = `docker run --rm -v "$(pwd):$(pwd)" -w "$(pwd)" sommerfeldio/source2adoc:` + internal.Version()

type CommandDocs struct {
	Name        string       `yaml:"name"`
	Synopsis    string       `yaml:"synopsis"`
	Description string       `yaml:"description"`
	Usage       string       `yaml:"usage"`
	Options     []OptionDocs `yaml:"options"`
	SeeAlso     []string     `yaml:"see_also"`
}

type OptionDocs struct {
	Name         string `yaml:"name"`
	DefaultValue string `yaml:"default_value"`
	Usage        string `yaml:"usage"`
	Shorthand    string `yaml:"shorthand"`
}

func GenerateManpage(rootCmd *cobra.Command) {
	err := doc.GenYamlTree(rootCmd, internal.CurrentWorkingDir())
	if err != nil {
		log.Fatal(err)
	}

	files, err := filepath.Glob("source2adoc*.yaml")
	if err != nil {
		log.Fatal(err)
	}

	initManpageAdoc(manpageFilename)

	for _, file := range files {
		commandDocs := parseYaml(file)
		appendCommandDocsToAdoc(commandDocs, manpageFilename)
		os.Remove(file)
	}
}

func parseYaml(file string) CommandDocs {
	yamlFile, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	var commandDocs CommandDocs
	err = yaml.Unmarshal(yamlFile, &commandDocs)
	if err != nil {
		log.Fatal(err)
	}

	return commandDocs
}

func initManpageAdoc(manpageFile string) {
	file, err := os.Create(manpageFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	appendStringToFile("= Manpage\n\n", manpageFile)
}

func appendCommandDocsToAdoc(commandDocs CommandDocs, manpageFile string) {
	file, err := os.OpenFile(manpageFile, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	appendStringToFile("\n== "+commandDocs.Name+"\n", manpageFile)
	appendStringToFile(commandDocs.Synopsis+"\n\n", manpageFile)

	appendStringToFile("[source, bash]\n", manpageFile)
	appendStringToFile("....\n", manpageFile)
	usage := commandDocs.Usage
	usage = strings.ReplaceAll(usage, "source2adoc", dockerRun)
	if usage == "" {
		usage = dockerRun
	}
	appendStringToFile(usage+"\n", manpageFile)
	appendStringToFile("....\n\n", manpageFile)

	appendStringToFile(commandDocs.Description+"\n\n", manpageFile)

	for _, option := range commandDocs.Options {
		defaultValue := option.DefaultValue
		if defaultValue == "" {
			defaultValue = "_none_"
		}

		if option.Shorthand == "" {
			appendStringToFile("* `"+option.Name+"`, default = "+defaultValue+" +\n", manpageFile)
		} else {
			appendStringToFile("* `"+option.Name+"`, `"+option.Shorthand+"`, default = "+defaultValue+" +\n", manpageFile)

		}

		appendStringToFile("  "+option.Usage+"\n", manpageFile)
	}
}

func appendStringToFile(content string, filename string) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		log.Fatal(err)
	}
}
