package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(writer http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(writer, "URL.Path = %q\n", r.URL.Path)
}

func counter(writer http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(writer, "Count %d\n", count)
	mu.Unlock()
}
