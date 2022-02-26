package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

// keep track of the database
var db *gorm.DB

// variable to keep track of the base URL
var baseURL string

// bcrypt strength
var strength = 11

type User struct {
	ID       uint   `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func initAuth(sqliteFile string, _baseURL string, debugSQL bool) {
	_db, err := gorm.Open("sqlite3", sqliteFile)
	if err != nil {
		panic(fmt.Sprintf("Could not open management database %v", err))
	}
	if debugSQL {
		db = _db.Debug()
	} else {
		db = _db
	}
	baseURL = _baseURL

	// migrate schemas
	db.AutoMigrate(&User{})
}

func HandleLogin() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// convert request to registration data
		var login userLogin
		err := json.NewDecoder(r.Body).Decode(&login)

		if err != nil {
			http.Error(w, "Difformed login request", http.StatusBadRequest)
			fmt.Println("Error: Difformed login request")
			return
		}

		// Bring the user from the database
		user := User{}
		db.Where(&User{Username: login.Username}).Find(&user)

		if user.Username != login.Username {
			http.Error(w, fmt.Sprintf("User %s does not exist", login.Username), http.StatusForbidden)
			fmt.Println("ERROR: %s does not exist", login.Username)
			return
		}

		// check the password
		err = bcrypt.CompareHashAndPassword([]byte(user.Hash), []byte(login.Password))
		if err != nil {
			http.Error(w, fmt.Sprintf("Password for user %s is incorrect", login.Username), http.StatusForbidden)
			return
		}

		// setup the session and tell user that everything is fine
		var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

		// existing session: Get() always returns a session, even if empty.
		session, err := store.Get(r, "session-name    ")
		if err == nil {
			session.Values["id"] = login.Username
			err = session.Save(r, w)
		}
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			http.Error(w, fmt.Sprintf("Could not setup session for %s user", login.Username), http.StatusConflict)
			return
		}

		json.NewEncoder(w).Encode(user)
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

			fmt.Fprintf(w, "POST request successful\n")

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

	}

}

func main() {
	fileServer := http.FileServer(http.Dir("../../CEN5035-front-end/src")) // use "../../CEN5035-front-end/src" for frontend, static is just for testing
	http.Handle("/", fileServer)

	fmt.Println("Welcome to Attorney Manager! this is a basic setup in GO for the backend of the project.")

	db, err := sql.Open("sqlite3", "./names.db")
	checkErr(err)

	initAuth("sample", "http://localhost:8080/", true)

	stmt, _ := db.Prepare(`
		CREATE TABLE IF NOT EXISTS "Customer" (
			"Name"	TEXT,
			"Password" TEXT
		);
	`)
	stmt.Exec()
	defer stmt.Close()

	fmt.Println("Database created!")

	http.HandleFunc("/api/login", loginHandler(db))

	fmt.Println("Starting server on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
