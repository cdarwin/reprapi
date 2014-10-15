package main

import (
	"flag"
	"github.com/gorilla/mux"
	"net/http"
)

// Default handler for root endpoint
func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write(doMarshall("Welcome to the reprepro REST API - reprapi"))
}

// Handler for /package/ endpoint
func packageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pkg := vars["pkg"]
	w.Write(marshallPackage(pkg))
}

// Handler for /distro/ endpoint
func distroHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	dist := vars["dist"]
	w.Write(marshallDistro(dist))
}

// Handler for /create/ endpoint
func createHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	dist := vars["dist"]
	w.Write(marshalCreate(dist))
}

func main() {
	// Parse optional command line arguments
	flag.Parse()
	// Create a new router using the gorilla/mux library
	r := mux.NewRouter()
	// Set a prefix for all subsequent endpoints
	s := r.PathPrefix("/reprapi/v2").Subrouter()
	// Handle some routes
	s.HandleFunc("/", homeHandler)
	s.HandleFunc("/package/{pkg}/", packageHandler)
	s.HandleFunc("/distro/{dist}/", distroHandler)
	s.HandleFunc("/create/{dist}/", createHandler)
	http.Handle("/", s)
	serve()
}
