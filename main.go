package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var addr = flag.String("addr", ":8080", "the address of the server")

func main() {
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[%s] %s", r.Method, r.URL)
		fmt.Fprintf(w, "Hello, golang")
	})
	log.Fatal(http.ListenAndServe(*addr, nil))
}
