package main

import (
	"compress/gzip"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Encoding", "gzip")
	w.Header().Set("Content-Type", "application/json")
	source := map[string]string{
		"Hello": "World",
	}
	zw := gzip.NewWriter(w)
	sourceJSON, err := json.Marshal(source)
	if err != nil {
		log.Fatal(err)
	}
	mw := io.MultiWriter(zw, os.Stdout)
	io.WriteString(mw, string(sourceJSON))
	io.WriteString(os.Stdout, "\n")
	zw.Flush()
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
