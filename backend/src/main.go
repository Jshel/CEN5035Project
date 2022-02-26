package main

import (
	auth "attorneyManager/_services"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func loginHandler(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method == "POST" {

			//read the json
			var data User
			var decoder = json.NewDecoder(r.Body)

			err := decoder.Decode(&data)
			if err != nil {
				log.Fatal("Error when opening file: ", err)
			}

			var name = data.Username
			var password = data.Password

			//fmt.Fprintf(w, "POST request successful\n")

			fmt.Println("name: ", name)
			fmt.Println("password: ", password)

			// fmt.Fprintf(w, "Name = %s\n", name)
			// fmt.Fprintf(w, "Password = %s\n", password)

			stmt, error := db.Prepare("INSERT INTO Customer(Name, Password) VALUES (?, ?);")
			if error != nil {
				checkErr(error)
			}

			stmt.Exec(name, password)
			defer stmt.Close()
		}

	}

}

func main() {
	fileServer := http.FileServer(http.Dir("../../CEN5035-front-end/src")) // use "../../CEN5035-front-end/src" for frontend, static is just for testing
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

	auth.InitAuth("./database.db", false)

	fmt.Println("Database created!")

	http.HandleFunc("/api/login", auth.HandleLogin())
	// http.HandleFunc("/api/logout", auth.HandleLogout())
	// http.HandleFunc("/api/register", auth.HandleRegister())

	fmt.Println("Starting server on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
