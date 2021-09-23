package main

import (
	"crypto/rand"
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.Create("rand.txt")
	if err != nil {
		log.Fatal(err)
	}
	r := rand.Reader
	io.CopyN(f, r, 1024)
}
