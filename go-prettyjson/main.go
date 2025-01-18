package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Println("Error: Please provide a file path")
	}

	filepath := flag.Arg(0)

	content, err := os.ReadFile(filepath)

	if err != nil {
		fmt.Printf("Error in reading the file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("content %v", string(content))

}
