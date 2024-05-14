package internal

import (
	"os"
	"path/filepath"
)

func TestDataDir() string {
	currentDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return filepath.Join(currentDir, "../../../testdata")
}

func CurrentWorkingDir() string {
	currentDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return currentDir
}
