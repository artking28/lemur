package main

import (
	"fmt"
	"os"
	"github.com/artking28/lemur/models"
)

func main() {
	if len(os.Args) != 2 {
		panic("invalid args. Please, enter a folder path")
	}

	input := os.Args[1]
	root, err := models.NewTree(input)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", root.ToString())
}
