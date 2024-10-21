package main

import (
	"fmt"
	"os"
	"slices"

	gotemplate "github.com/matwate/makemake/make_templates/go_template"
)

var extensions_to_generate = map[string]func() (string, error){
	".go": gotemplate.GetTemplate,
}

func generate_makefiles(exts []string) error {
	// Generate makefiles for the given extensions
	fmt.Println("Generating makefiles for the following extensions:")
	for _, ext := range exts {
		if !slices.Contains(extensions, ext) {
			return fmt.Errorf("extension %s not supported", ext)
		}
		template, ok := extensions_to_generate[ext]
		if !ok {
			return fmt.Errorf("template for extension %s not found", ext)
		}
		content, err := template()
		if err != nil {
			return fmt.Errorf("error while generating makefile for extension %s: %v", ext, err)
		}
		fmt.Println(content)
		err = write_to_file(fmt.Sprintf("Makefile"), content)
	}
	return nil
}

func write_to_file(s, content string) error {
	err := os.WriteFile(s, []byte(content), 0644)
	if err != nil {
		return fmt.Errorf("error while writing to file %s: %v", s, err)
	}
	return nil
}
