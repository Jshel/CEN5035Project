package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/test" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method not supported,", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Hello!")
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful")
	name := r.FormValue("name")
	password := r.FormValue("password")

	fmt.Println("name: ", name)
	fmt.Println("password: ", password)

	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Password = %s\n", password)
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)

	http.HandleFunc("/test", testHandler)
	http.HandleFunc("/login", loginHandler)

	fmt.Println("Starting server on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Welcome to Attorney Manager! this is a basic setup in GO for the backend of the project.")
	db, err := sql.Open("sqlite3", "./names.db")
	checkErr(err)
	defer db.Close()
}
