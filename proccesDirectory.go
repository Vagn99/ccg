package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func processDirectory(path string, recursive bool) {
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

	var result strings.Builder
	for _, file := range files {
		if ignoreFile(file.Name(), ignoreList) {
			continue
		}
		filePath := filepath.Join(path, file.Name())
		if file.IsDir() {
			if recursive {
				processDirectory(filePath, recursive)
			}
		} else {
			result.WriteString(filePath + "\n" + strings.Repeat("-", len(filePath)) + "\n")
			fileContents, err := os.ReadFile(filePath)
			if err == nil {
				result.WriteString(string(fileContents) + "\n")
			}
		}
	}

	contextFilePath := filepath.Join(path, "context.txt")
	err = os.WriteFile(contextFilePath, []byte(result.String()), 0644)
	if err != nil {
		fmt.Println("Error writing context file:", err)
	}
}

func readIgnoreFile(path string) ([]string, error) {
	ignoreFilePath := filepath.Join(path, ".ccgignore")
	file, err := os.Open(ignoreFilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func ignoreFile(fileName string, ignoreList []string) bool {
	for _, ignore := range ignoreList {
		if fileName == ignore {
			return true
		}
	}
	return false
}
