package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

// addr indicates the address used by the server.
var addr = flag.String("addr", ":8080", "the address of the server")

// Log returns a handler that wraps the given handler with request logging
// feature.
func Log(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf(`"%s %s %s"`, r.Method, r.RequestURI, r.Proto)
		h.ServeHTTP(w, r)
	})
}

// Error replies to the request with the specified status code.
func Error(w http.ResponseWriter, status int) {
	text := http.StatusText(status)
	http.Error(w, text, status)
}

// ErrorHandler returns a handler replies to the request with the specified
// status code.
func ErrorHandler(status int) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { Error(w, status) })
}

func Events(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		CreateEvent(w, r)
	default:
		Error(w, http.StatusMethodNotAllowed)
	}
}

func CreateEvent(w http.ResponseWriter, r *http.Request) {
	// Generate random 32 characters
	s := ""
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 4; i++ {
		s += strconv.FormatUint(rand.Uint64(), 36)[0:8]
	}
	fmt.Fprintln(w, s)
}

func Event(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		GetEvent(w, r)
	case "DELETE":
		DestroyEvent(w, r)
	default:
		Error(w, http.StatusMethodNotAllowed)
	}
}

func GetEvent(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, r.URL.Path)
}

func DestroyEvent(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, r.URL.Path)
}

func main() {
	flag.Parse()

	mux := http.NewServeMux()
	mux.HandleFunc("/events", Events)
	mux.Handle("/events/", http.StripPrefix("/events/", http.HandlerFunc(Event)))

	log.Fatal(http.ListenAndServe(*addr, Log(mux)))
}
