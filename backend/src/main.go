package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// For this case, we will always pipe "Hello World" into the response writer
	fmt.Fprintf(w, "Hello World!")
}

func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/hello", handler).Methods("GET")

	staticFileDirectory := http.Dir("../../CEN5035-front-end/src/")

	staticFileHandler := http.StripPrefix("/index", http.FileServer(staticFileDirectory))

	r.PathPrefix("/index").Handler(staticFileHandler).Methods("GET")
	return r
}

func main() {

	r := newRouter()

	r.HandleFunc("/hello", handler).Methods("GET")

	http.ListenAndServe(":8080", r)
}
