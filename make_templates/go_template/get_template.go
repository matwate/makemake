package gotemplate

import (
	"fmt"
	"io"
	"os"
)

/*
 This file will contain a helper function that will return
  the contents associated to the Makefile template contained in this
  dir.
*/

func GetTemplate() (string, error) {
	fmt.Println("Getting template for go files")
	makefile, err := os.Open(
		"./make_templates/go_template/Makefile",
	) // Should modify this path to be relative to the project.
	if err != nil {
		return "", err
	}
	file, err := io.ReadAll(makefile)
	if err != nil {
		return "", err
	}
	return string(file), nil
}
