package helper

import (
	"os"
	"path/filepath"
)

// CurrentWorkingDir returns the current working directory from where the app is started for use
// with the actual app (outside of unit tests).
func CurrentWorkingDir() string {
	currentDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return currentDir
}

// TestDataDir returns the path to the testdata directory for use in unit tests.
func TestDataDir() string {
	return filepath.Join(CurrentWorkingDir(), "../../../testdata")
}
