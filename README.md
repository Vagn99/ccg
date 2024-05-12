# Code Context Generator (ccg)

## Overview

`ccg` is a tool developed in Go, designed to consolidate all the code files within a specified directory (and optionally its subdirectories) into a single text file. This tool is particularly useful for preparing a codebase for analysis or processing by AI systems, as it formats the entire codebase in a readily accessible manner.

## Features

- **Single File Output**: Generates a single `context.txt` that contains the contents of all files processed, along with their paths.
- **Recursive Option**: Includes a recursive option `-r` to process all subdirectories.
- **Ignore Functionality**: Supports a `.ccgignore` file which works similarly to `.gitignore`, allowing specific files and directories to be excluded from processing.

## Installation

To use `ccg`, clone this repository and build the application using Go:

```bash
git clone [repository-url]
cd [repository-name]
go build -o ccg

# Add the binary to your PATH
mv ccg /usr/local/bin
```

## Usage
```bash
# Process the current directory, excluding subdirectories:
ccg .


# Process the current directory and all subdirectories:
ccg -r .


# Process a specific directory and all its subdirectories:
ccg -r /path/to/directory
```


## Configuration
Create a `.ccgignore` file in the root directory to exclude specific files and directories from processing. The format is the same as `.gitignore`:

```.gitignore
# Example of contents in .ccgignore
.ccgignore
ccg
context.txt
/.git
/.idea
/.gitignore
README.md
```

## Contributing
Contributions are welcome! Please fork the repository and submit a pull request with your changes.