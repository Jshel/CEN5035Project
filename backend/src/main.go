package main

import (
	auth "attorneyManager/_auth"
	contract "attorneyManager/_contract"
	messages "attorneyManager/_messages"
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type M map[string]interface{}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
}

func setUpRoutes() {
	// request handlers
	http.HandleFunc("/api/login", auth.HandleLogin())
	http.HandleFunc("/api/logout", auth.HandleLogout())
	http.HandleFunc("/api/create-account", auth.HandleRegister())
	http.HandleFunc("/api/get-contract", contract.HandleGetContract())
	http.HandleFunc("/api/upload", contract.HandleFileUpload())
	http.HandleFunc("/api/download", contract.HandleFileDownload())
	http.HandleFunc("/api/getuser", auth.GetUserEmail())
	http.HandleFunc("/api/send-message", messages.HandleSendMessage())
	http.HandleFunc("/api/get-message", messages.HandleGetMessage())
}

func databaseInit() {
	// database inits
	auth.InitAuth("./user_database.db", false)
	contract.InitContractDB("./contract_database.db", false)
	messages.InitMessageDB("./message_database.db", false)
	//session init
	gob.Register(&M{})
}

func main() {
	fileServer := http.FileServer(http.Dir("../../CEN5035-front-end/src"))
	http.Handle("/", fileServer)

	fmt.Println("Welcome to Attorney Manager! this is a basic setup in GO for the backend of the project.")

	databaseInit()
	setUpRoutes()

	fmt.Println("Starting server on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
