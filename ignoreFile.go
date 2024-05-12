package main

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
)

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
