package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

//var db *sql.DB

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

func apiHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Backend api test")
}

func loginHandler(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			http.ServeFile(w, r, "./static/login.html")
		case "POST":
			{
				if err := r.ParseForm(); err != nil {
					fmt.Fprintf(w, "ParseForm() err: %v", err)
					return
				}

				fmt.Fprintf(w, "POST request successful\n")
				var name string = r.FormValue("name")
				var password string = r.FormValue("password")

				fmt.Println("name: ", name)
				fmt.Println("password: ", password)

				fmt.Fprintf(w, "Name = %s\n", name)
				fmt.Fprintf(w, "Password = %s\n", password)

				stmt, error := db.Prepare("INSERT INTO Customer(Name, Password) VALUES (?, ?);")
				if error != nil {
					checkErr(error)
				}

				stmt.Exec(name, password)
				defer stmt.Close()
			}
		default:
			fmt.Println(w, "only GET and POST")
		}

	}

}

func main() {
	fileServer := http.FileServer(http.Dir("./static")) // use "../../CEN5035-front-end/src" for frontend static is just for testing
	http.Handle("/", fileServer)

	fmt.Println("Welcome to Attorney Manager! this is a basic setup in GO for the backend of the project.")

	db, err := sql.Open("sqlite3", "./names.db")
	checkErr(err)

	stmt, _ := db.Prepare(`
		CREATE TABLE IF NOT EXISTS "Customer" (
			"Name"	TEXT,
			"Password" TEXT
		);
	`)
	stmt.Exec()
	defer stmt.Close()

	fmt.Println("Database created!")

	http.HandleFunc("/test", testHandler)
	http.HandleFunc("/login", loginHandler(db))
	http.HandleFunc("/api", apiHandler)

	fmt.Println("Starting server on port 8888")
	if err := http.ListenAndServe(":8888", nil); err != nil {
		log.Fatal(err)
	}

}
