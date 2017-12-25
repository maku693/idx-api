package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var addr = flag.String("addr", ":8080", "the address of the server")

// Log returns a handler that wraps the given handler with request logging
// feature.
func Log(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf(`"%s %s %s"`, r.Method, r.RequestURI, r.Proto)
		h.ServeHTTP(w, r)
	})
}

func main() {
	flag.Parse()

	mux := http.NewServeMux()
	mux.HandleFunc("/hello", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintln(w, "Hello, golang")
	})
	log.Fatal(http.ListenAndServe(*addr, Log(mux)))
}
