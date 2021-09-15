package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Create("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(f, "int: %d, float: %f, string: %s\n", 100, 12.34, "str")
}
