package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
)

// Determines how exaggerated the indentation in the fs visualization should be
var depthMultiplier uint8 = 4

var glob *string
var root *string
var includeHiddenFiles *bool
var maxDepth *uint64
var maxFilesToScan *uint64
var success, warning *color.Color
var globMatches []string

var fileCountLimitReached bool = false
var scannedItemsCount uint64 = 0
var matchedItemsCount uint64 = 0

func init() {
	success = color.New(color.FgHiGreen)
	warning = color.New(color.FgHiYellow)
}

func main() {
	glob = flag.String("glob", "*", "The glob pattern that will be used to match filenames")
	root = flag.String("root", ".", "The starting directory that will act as the root")
	includeHiddenFiles = flag.Bool("hidden-files", false, "Whether or not to include hidden directories and files")
	maxDepth = flag.Uint64("max-depth", 20, "How many nested directories should be traversed before stopping")
	maxFilesToScan = flag.Uint64("max-files-to-scan", 200, "How many files to scan before stopping (useful to avoid scanning every file from root or something)")
	flag.Parse()

	absPath, err := filepath.Abs(*root)
	goToRoot(absPath) // Go to correct working directory

	globMatches, err = filepath.Glob(*glob)
	if err != nil {
		panic(err)
	}

	globMatches, err = mapRelativeToAbsolutePaths(globMatches)
	if err != nil {
		panic(err)
	}

	digestDir(absPath, 0)
	if fileCountLimitReached {
		warning.Printf(fileCountOverflowMessage, scannedItemsCount)
	}
	success.Printf("%d files matched out of %d files scanned.\n", matchedItemsCount, scannedItemsCount)
}

func goToRoot(absPath string) {
	err := os.Chdir(absPath)
	if err != nil {
		panic(err)
	}
}

func digestDir(path string, depth uint8) {
	// As long as we haven't reached the max depth, we can continue traversing directories
	if uint64(depth) < *maxDepth {
		processDirEntries(path, depth)
	}
}

func processDirEntries(path string, depth uint8) {
	fileInfo, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}

	for _, item := range fileInfo {
		if scannedItemsCount >= *maxFilesToScan {
			fileCountLimitReached = true
			continue
		}

		itemName := item.Name()

		// Check to see if we should ignore hidden items
		if !*includeHiddenFiles && strings.HasPrefix(itemName, ".") {
			continue
		}

		styledName := formatName(itemName, depth)
		absItemPath := buildNextPathName(path, item)

		// Recursively handle subdirectories
		if item.IsDir() {
			fmt.Println(styledName)
			digestDir(absItemPath, depth+1)
		} else {
			scannedItemsCount++
			matched := contains(globMatches, absItemPath)

			if matched {
				success.Println(styledName)
				matchedItemsCount++
			} else {
				fmt.Println(styledName)
			}
		}
	}
}

func buildNextPathName(path string, file os.FileInfo) string {
	return filepath.Join(path, file.Name())
}

func formatName(name string, depth uint8) string {
	return buildIndentation(depth) + setPathGlyph(depth) + name
}

func buildIndentation(depth uint8) string {
	var spaces string
	depthMagnitude := depth * depthMultiplier
	for i := uint8(0); i < depthMagnitude; i++ {
		spaces = spaces + " "
	}

	return spaces
}

func setPathGlyph(depth uint8) string {
	if depth > 0 {
		return "\u2514 " // ‚îî
	} else {
		return "- "
	}
}
