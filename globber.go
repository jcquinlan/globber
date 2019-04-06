package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

var glob *string
var root *string

func main() {
	glob = flag.String("glob", "*", "The glob pattern that will be used to match filenames")
	root = flag.String("root", ".", "The starting directory that will act as the root")
	flag.Parse()

	digestDir(*root, 0)

}

func digestDir(path string, depth uint8) {
	fileInfo, err := ioutil.ReadDir(path)

	if err != nil {
		panic(err)
	}

	for _, item := range fileInfo {
		// Recursively handle subdirectories
		if item.IsDir() {
			digestDir(buildPathName(path, item), depth+1)

		} else {
			matched, err := filepath.Match(*glob, item.Name())

			if err != nil {
				panic(err)
			}

			if matched {
				adornPath(buildPathName(path, item), depth)
			}
		}
	}
}

func buildPathName(path string, file os.FileInfo) string {
	return path + "/" + file.Name()
}

func adornPath(path string, depth uint8) {
	indentation := buildIndentation(depth)
	fmt.Println(indentation + path)
}

func buildIndentation(depth uint8) string {
	var spaces string
	for i := uint8(0); i < depth*4; i++ {
		spaces = spaces + " "
	}

	return spaces
}
