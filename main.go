package main

import (
	"flag"
	"fmt"
	"os"
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

	if *recursive {
		fmt.Println("Recursive")
	} else {
		fmt.Println("Not recursive")
	}

	fmt.Println("Path:", rootPath)

	processDirectory(rootPath, *recursive)
}
