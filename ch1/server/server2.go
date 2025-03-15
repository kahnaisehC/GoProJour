package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu struct {
	sync.Mutex
	count int
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	mu.count++
	mu.Unlock()
	fmt.Fprintf(w, "request to %q", r.URL)
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "%v", mu.count)
	mu.Unlock()
}
