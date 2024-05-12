package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func processDirectory(path string, recursive bool) {
	files, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	var result strings.Builder
	for _, file := range files {
		filePath := filepath.Join(path, file.Name())
		if ignoreFile(file.Name()) {
			continue
		}
		if file.IsDir() {
			if recursive {
				processDirectory(filePath, recursive)
			}
		} else {
			//Include filepath and name in the result and a divider line
			result.WriteString(filePath)
			result.WriteString(strings.Repeat("-", len(filePath)) + "\n")
			fileContents, err := os.ReadFile(filePath)
			if err == nil {
				result.WriteString(string(fileContents) + "\n")
			}
		}
	}

	// Print the result to context.txt
	contextFilePath := filepath.Join(path, "context.txt")
	err = os.WriteFile(contextFilePath, []byte(result.String()), 0644)
	if err != nil {
		fmt.Println("Error writing context file:", err)
	}

}

func ignoreFile(fileName string) bool {
	//List of excluded filenames
	ignoreList := []string{"gcc", "context.txt"}
	for _, ignore := range ignoreList {
		if fileName == ignore {
			return true
		}
	}
	return false
}
