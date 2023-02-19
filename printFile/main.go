package main

import (
	"io"
	"log"
	"os"
)

func main() {
	fileFromArgs := os.Args[1]
	content, err := os.Open(fileFromArgs)
	if err != nil {
		log.Fatal(err)
	}

	io.Copy(os.Stdout, content)
}
