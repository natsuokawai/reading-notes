package main

import "os"

func main() {
	os.Stdout.Write([]byte("write to stdout from go\n"))
}
