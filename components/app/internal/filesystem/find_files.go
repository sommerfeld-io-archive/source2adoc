package filesystem

import (
	"os"
	"path/filepath"
)

// FindFilesByPattern scans the filesystem from the given startPath and returns all filenames with
// path that match the given pattern. To use this function, you can pass the starting directory
// path and the desired pattern as arguments. It will return a slice of strings containing the
// paths of the matching files, or an error if any occurred during the process.
func FindFilesByPattern(startPath string, pattern string) ([]string, error) {
	var matchingFiles []string

	err := filepath.Walk(startPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			match, err := filepath.Match(pattern, filepath.Base(path))
			if err != nil {
				return err
			}

			if match {
				pathWithoutCWD, err := pathWithoutCWD(startPath, path)
				if err != nil {
					return err
				}
				matchingFiles = append(matchingFiles, pathWithoutCWD)
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return matchingFiles, nil
}

func pathWithoutCWD(startPath string, path string) (string, error) {
	pathWithoutCWD, err := filepath.Rel(startPath, path)
	if err != nil {
		return "", err
	}
	return pathWithoutCWD, nil
}
