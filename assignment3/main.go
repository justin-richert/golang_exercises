package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go [fileToOpen]")
		os.Exit(1)
	}

	fileToOpen := os.Args[1]
	fmt.Println(fileToOpen)

	file, err := os.Open(fileToOpen)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, file)
}
