package main

import (
	"fmt"
	"io/fs"
	"maps"
	"os"
	"path/filepath"
	"slices"
)

/*
  What we'll be doing
  Cli command that will scan the files in the current directory
  and generate a makefile according to the file extensions
*/

// Available File extensions
var extensions = []string{".py", ".cpp", ".go"}

func main() {
	// Get the files in the current directory
	files, err := getFiles()
	if err != nil {
		_ = fmt.Errorf("error while getting files: %v", err)
	}
	// Make a set of the file extensions found in the project
	extSet := getExtensions(files)
	fmt.Println(extSet)

	// Get a list of the keys of the map
	exts := slices.Collect(maps.Keys(extSet))
	// Generate makefiles for the extensions
	err = generate_makefiles(exts)
	if err != nil {
		_ = fmt.Errorf("error while generating makefiles: %v", err)
	}
	fmt.Println("Makefiles generated successfully")
}

func getExtensions(files []fs.DirEntry) map[string]bool {
	extSet := make(map[string]bool)
	for _, file := range files {
		ext := filepath.Ext(file.Name())
		// Remove not supported extensions
		if slices.Contains(extensions, ext) {
			extSet[ext] = true
		}
	}
	return extSet
}

func getFiles() ([]os.DirEntry, error) {
	files, err := os.ReadDir(".")
	if err != nil {
		return nil, err
	}
	return files, nil
}
