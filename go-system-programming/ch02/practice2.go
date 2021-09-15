package main

import (
	"encoding/csv"
	"os"
)

func main() {
	w := csv.NewWriter(os.Stdout)
	records := [][]string{
		{"id", "name"},
		{"1", "Alice"},
		{"2", "Bob"},
	}
	for _, r := range records {
		w.Write(r)
	}
	w.Flush()
}
