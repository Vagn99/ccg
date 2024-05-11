package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

func processDirectory(path string, recursive bool) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	var result strings.Builder
	for _, file := range files {
		filePath := filepath.Join(path, file.Name())
		if file.Name() == "ccg" {
			continue
		}
		if file.IsDir() {
			if recursive {
				processDirectory(filePath, recursive)
			}
		} else {
			fileContents, err := ioutil.ReadFile(filePath)
			if err == nil {
				result.WriteString(string(fileContents) + "\n")
			}
		}
	}

	// Save or handle the concatenated result
	fmt.Println(result.String())

}
