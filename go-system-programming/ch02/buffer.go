package main

import (
	"bytes"
	"fmt"
)

func main() {
	var buffer bytes.Buffer
	buffer.Write([]byte("byte test\n"))
	fmt.Println(buffer.String())
}
