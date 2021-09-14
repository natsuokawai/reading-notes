package main

import "os"

func main() {
	file, _ := os.Create("test.txt")
	file.Write([]byte("hello from go!\n"))
	file.Close()
}
