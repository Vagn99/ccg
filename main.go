package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	recursive := flag.Bool("r", false, "Include subdirectories recursively")
	flag.Parse()
	args := flag.Args()

	if len(args) < 1 {
		fmt.Println("Usage: ccg [-r] <rootPath>")
		os.Exit(1)
	}

	rootPath := args[0]
	var result strings.Builder
	processDirectory(rootPath, *recursive, &result)

	// Write the final result to context.txt in the root directory
	contextFilePath := filepath.Join(rootPath, "context.txt")
	err := os.WriteFile(contextFilePath, []byte(result.String()), 0644)
	if err != nil {
		fmt.Println("Error writing context file:", err)
	}
}
