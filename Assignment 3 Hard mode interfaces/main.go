package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	var file string

	if len(os.Args) > 1 {
		file = os.Args[1]
	} else {
		fmt.Println("Please enter the name of a text file")
		os.Exit(1)
	}
	fs, err := os.Open(file)

	if err != nil {
		fmt.Println("Error:", err)
	}

	io.Copy(os.Stdout, fs)
}
