package main

import (
	"bufio"
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

func readIgnoreFile(path string) ([]string, error) {
	ignoreFilePath := filepath.Join(path, ".ccgignore")
	file, err := os.Open(ignoreFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			// If the file does not exist, return an empty list (no ignores)
			return nil, nil
		}
		// For other types of errors, return the error to be handled upstream
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" && !strings.HasPrefix(line, "#") { // Ignore empty lines and comments
			lines = append(lines, line)
		}
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
