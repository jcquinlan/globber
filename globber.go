package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var glob *string
var root *string
var includeHiddenFiles *bool

func main() {
	glob = flag.String("glob", "*", "The glob pattern that will be used to match filenames")
	root = flag.String("root", ".", "The starting directory that will act as the root")
	includeHiddenFiles = flag.Bool("hidden-files", false, "Whether or not to include hidden directories and files")
	flag.Parse()

	digestDir(*root, 0)

}

func digestDir(path string, depth uint8) {
	fileInfo, err := ioutil.ReadDir(path)

	if err != nil {
		panic(err)
	}

	for _, item := range fileInfo {
		itemName := item.Name()
		styledPath := itemName

		if !*includeHiddenFiles && strings.HasPrefix(itemName, ".") {
			continue
		}

		// Recursively handle subdirectories
		if item.IsDir() {
			styledPath = indentPath(styledPath, depth)
			fmt.Println(styledPath)
			digestDir(buildPathName(path, item), depth+1)
		} else {
			matched, err := filepath.Match(*glob, itemName)

			if err != nil {
				panic(err)
			}

			if matched {
				styledPath = colorPath(styledPath)
			}

			styledPath = indentPath(styledPath, depth)
			fmt.Println(styledPath)
		}
	}
}

func buildPathName(path string, file os.FileInfo) string {
	return path + "/" + file.Name()
}

func indentPath(path string, depth uint8) string {
	return buildIndentation(depth) + path
}

func colorPath(path string) string {
	return path
}

func buildIndentation(depth uint8) string {
	var spaces string
	depthMagnitude := depth * 4
	for i := uint8(0); i < depthMagnitude; i++ {
		spaces = spaces + " "
	}

	if depth > 0 {
		spaces = spaces + "\u2514 "
	} else {
		spaces = spaces + "- "
	}

	return spaces
}
