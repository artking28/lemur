# Lemur

A simple program to display a directory tree structure or parse a list of paths.

## Installation

1. Ensure Go is installed.
2. Run `make install` to build and install the binary to your system's bin directory.

## Usage

```bash
# Show version info
lemur
lemur -v

# Show help panel
lemur -h

# Read paths from stdin
cat paths.txt | lemur stdout

# Read paths from a text file
lemur -f paths.txt

# Read a filesystem directory
lemur -p <directory_path>
