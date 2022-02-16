package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	username string
	password string
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func testHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		fmt.Println("test")
		fmt.Fprintf(w, "test")
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
			http.ServeFile(w, r, "./static/login.html") // use "../../CEN5035-front-end/src" for frontend, static is just for testing
		case "POST":
			{
				// if err := r.ParseForm(); err != nil {
				// 	fmt.Fprintf(w, "ParseForm() err: %v", err)
				// 	return
				// }

				fmt.Fprintf(w, "POST request successful\n")

				//read the json
				var data User
				var decoder = json.NewDecoder(r.Body)

				err := decoder.Decode(&data)
				if err != nil {
					log.Fatal("Error when opening file: ", err)
				}

				var name = data.username
				var password = data.password

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
	fileServer := http.FileServer(http.Dir("./static")) // use "../../CEN5035-front-end/src" for frontend, static is just for testing
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

	http.HandleFunc("/api/test", testHandler)
	http.HandleFunc("/api/login", loginHandler(db))
	//http.HandleFunc("/api", apiHandler)

	fmt.Println("Starting server on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
