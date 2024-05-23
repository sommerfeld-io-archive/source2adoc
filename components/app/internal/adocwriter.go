package internal

import (
	"os"
	"path/filepath"
	"strings"
)

// WriteAdoc writes the relevant comments of a code file to a newly generated AsciiDoc file.
// It takes the path of the code file, the Antora directory, and the Antora module as input parameters.
// It returns the path of the generated AsciiDoc file and an error if any.
func WriteAdoc(codeFilePath string, antoraDir string, antoraModule string) (string, error) {
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
