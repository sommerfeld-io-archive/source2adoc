package internal

import (
	"os"
	"path/filepath"
	"strings"
)

// WriteAdocFile writes the relevant comments of a code file to a newly generated AsciiDoc file.
// It takes the path of the code file, the Antora directory, and the Antora module as input parameters.
// It returns the path of the generated AsciiDoc file and an error if any.
func WriteAdocFile(codeFilePath string, antoraDir string, antoraModule string) (string, error) {
	adocPath := generateAdocPath(codeFilePath, antoraDir, antoraModule)

	err := os.MkdirAll(filepath.Dir(adocPath), os.ModePerm)
	if err != nil {
		return "", err
	}

	file, err := os.Create(adocPath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	writeTitleAndMetadataToAdocFile(file, codeFilePath)

	return adocPath, nil
}

// generateAdocPath generates the path for an AsciiDoc file based on the given code file path,
// Antora directory, and Antora module.
func generateAdocPath(codeFilePath string, antoraDir string, antoraModule string) string {
	codeFilePath = strings.ReplaceAll(codeFilePath, ".", "-")
	adocPath := antoraDir + "/modules/" + antoraModule + "/pages/" + codeFilePath + ".adoc"
	return adocPath
}

// writeTitleAndMetadataToAdocFile writes the title and metadata to an AsciiDoc file.
func writeTitleAndMetadataToAdocFile(file *os.File, codeFilePath string) error {
	linesToWrite := []string{
		"= " + generateTitle(codeFilePath),
		"source2adoc <https://source2adoc.sommerfeld.io>",
		"",
		"[cols=\"1,1\"]",
		"|===",
		"|Source Code File Location |" + codeFilePath,
		"|Some More Meta |tbd ...",
		"|===",
		"",
	}

	for _, line := range linesToWrite {
		_, err := file.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}

	return nil
}

// generateTitle generates a title from the given code file path.
// It extracts the file name from the path and returns it as the title.
func generateTitle(codeFilePath string) string {
	codeFileName := strings.LastIndex(codeFilePath, "/")
	if codeFileName == -1 {
		return codeFilePath
	}
	return codeFilePath[codeFileName+1:]
}

// AppendToNavPartial appends a link to the navigation partial file in an Antora project.
// The function generates the corresponding AsciiDoc file path based on the code file path,
// and appends a link to the navigation partial file with the generated AsciiDoc file path and code file path.
// If the navigation partial file or any required directories do not exist, they will be created.
// The function returns an error if any error occurs during the file operations or directory creation.
func AppendToNavPartial(antoraDir string, antoraModule string, codeFilePath string) error {
	adocPath := generateAdocPath(codeFilePath, antoraDir, antoraModule)
	navAdocPartialPath := antoraDir + "/modules/" + antoraModule + "/partials/nav.adoc"

	err := os.MkdirAll(filepath.Dir(navAdocPartialPath), os.ModePerm)
	if err != nil {
		return err
	}

	// open or create file
	file, err := os.OpenFile(navAdocPartialPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	xrefPath := strings.Replace(adocPath, antoraDir+"/modules/"+antoraModule+"/pages/", "", 1)
	line := "* xref:" + antoraModule + ":" + xrefPath + "[" + codeFilePath + "]"
	_, err = file.WriteString(line + "\n")
	if err != nil {
		return err
	}

	return nil
}

// WriteNavAdoc writes a navigation file (nav.adoc) for a given Antora module.
// The function creates the necessary directory structure for the nav.adoc file,
// opens or creates the file, and writes a line containing a cross-reference to the
// module's index.adoc file.
// If any error occurs during the process, the function returns the error.
func WriteNavAdoc(antoraDir string, antoraModule string) error {
	navAdocPath := antoraDir + "/modules/" + antoraModule + "/nav.adoc"

	err := os.MkdirAll(filepath.Dir(navAdocPath), os.ModePerm)
	if err != nil {
		return err
	}

	// open or create file
	file, err := os.Create(navAdocPath)
	if err != nil {
		return err
	}
	defer file.Close()

	line := "* xref:" + antoraModule + ":index.adoc[]"
	_, err = file.WriteString(line + "\n")
	if err != nil {
		return err
	}

	return nil
}

func WriteIndexAdoc(antoraDir string, antoraModule string) error {
	navAdocPath := antoraDir + "/modules/" + antoraModule + "/pages/index.adoc"

	err := os.MkdirAll(filepath.Dir(navAdocPath), os.ModePerm)
	if err != nil {
		return err
	}

	// open or create file
	file, err := os.Create(navAdocPath)
	if err != nil {
		return err
	}
	defer file.Close()

	linesToWrite := []string{
		"= Source Code Docs",
		"source2adoc <https://source2adoc.sommerfeld.io>",
		"",
		"include::" + antoraModule + ":partial$nav.adoc[]",
	}

	for _, line := range linesToWrite {
		_, err := file.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}

	return nil
}
