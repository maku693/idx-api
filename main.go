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
		fmt.Fprint(w, "Hello, golang")
	})

	log.Fatal(http.ListenAndServe(*addr, nil))
}
