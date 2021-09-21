package main

import (
	"io"
	"log"
	"os"
)

func main() {
	oldFile, err := os.Open("old.txt")
	if err != nil {
		log.Fatal(err)
	}
	newFile, err := os.Create("new.txt")
	if err != nil {
		log.Fatal(err)
	}
	_, err = io.Copy(newFile, oldFile)
	if err != nil {
		log.Fatal(err)
	}
}
