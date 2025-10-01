package main

import (
	"bufio"
	"errors"
	"fmt"
	"io/fs"
	"os"

	"github.com/artking28/lemur/models"
)

func VersionPanel() {
	content, err := os.ReadFile("VERSION")
	if err != nil {
		if !errors.Is(err, fs.ErrNotExist) {
			panic(err)
		}
		content = []byte("0.0.0")
	}
	
	fmt.Printf("Welcome to lemur\n%s\n", string(content))
	os.Exit(0)
}

func HelpPanel() {
	fmt.Println(`Usage: lemur [options] [FILE/PATH]

Options:
  -h        Show this help panel
  -v        Show program version
  -f FILE   Read and process the specified file
  -p PATH   Read and process the specified directory
  stdout    Read input from pipe (stdin)

Examples:
  lemur -v
  lemur -h
  echo "content" | lemur stdout
  lemur -f file.txt
  lemur -p /my/directory`)
}

func ReadPath(input string) {
	root, err := models.NewTreePath(input)
	if err != nil {
		Fail(err)
	}

	fmt.Printf("%v\n", root.ToString())
	os.Exit(0)
}

func ReadPipe() {
	
	var lines string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		lines += scanner.Text()
	}
	
	err := scanner.Err();
	if err != nil {
		Fail(err)
	}
	
	root, err := models.NewTreeList("stdout", lines)
	if err != nil {
		Fail(err)
	}

	fmt.Printf("%v\n", root.ToString())
	os.Exit(0)
}

func ReadFile(input string) {
	
	content, err := os.ReadFile(input)
	if err != nil {
		Fail(err)
	}
	
	root, err := models.NewTreeList(input, string(content))
	if err != nil {
		Fail(err)
	}

	fmt.Printf("%v\n", root.ToString())
	os.Exit(0)
}

func Fail(err error) {
	if err == nil {
		err = errors.New("invalid args. Please, enter 'lemur -h' to get help.")
	}
	fmt.Println(err.Error())
	os.Exit(1)
}