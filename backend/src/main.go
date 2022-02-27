package main

import (
	auth "attorneyManager/_services"
	"fmt"
	"log"
	"net/http"
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

func main() {
	fileServer := http.FileServer(http.Dir("../../CEN5035-front-end/src")) // use "../../CEN5035-front-end/src" for frontend, static is just for testing
	http.Handle("/", fileServer)

	fmt.Println("Welcome to Attorney Manager! this is a basic setup in GO for the backend of the project.")

	auth.InitAuth("./database.db", false)

	http.HandleFunc("/api/login", auth.HandleLogin())
	// http.HandleFunc("/api/logout", auth.HandleLogout())
	http.HandleFunc("/api/create-account", auth.HandleRegister())

	fmt.Println("Starting server on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
