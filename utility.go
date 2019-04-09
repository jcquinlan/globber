package main

import (
	"path/filepath"
)

func contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}

	return false
}

func mapRelativeToAbsolutePaths(paths []string) ([]string, error) {
	numPaths := len(paths)
	absPaths := make([]string, numPaths, numPaths)

	for i := 0; i < numPaths; i++ {
		absPath, err := filepath.Abs(paths[i])
		if err != nil {
			return nil, err
		}

		absPaths = append(absPaths, absPath)
	}

	return absPaths, nil
}
