package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func processDirectory(path string, recursive bool, result *strings.Builder) {
	ignoreList, err := readIgnoreFile(path)
	if err != nil {
		fmt.Println("Error reading .ccgignore:", err)
		// Optional: decide whether to continue or exit on this error
	}

	files, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	for _, file := range files {
		if ignoreFile(file.Name(), ignoreList) {
			continue
		}
		filePath := filepath.Join(path, file.Name())
		if file.IsDir() {
			if recursive {
				processDirectory(filePath, recursive, result)
			}
		} else {
			result.WriteString(filePath + "\n" + strings.Repeat("-", len(filePath)) + "\n")
			fileContents, err := os.ReadFile(filePath)
			if err == nil {
				result.WriteString(string(fileContents) + "\n")
			}
		}
	}
}
