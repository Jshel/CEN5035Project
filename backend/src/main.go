package main

import (
	auth "attorneyManager/_auth"
	contract "attorneyManager/_contract"
	"fmt"
	"io/ioutil"
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

func uploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Uploading file...\n")

	//parse input
	r.ParseMultipartForm(10 << 20)

	//retreive files
	file, handler, err := r.FormFile("contract")
	if err != nil {
		fmt.Println("Error getting file from form:")
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Println("Uploaded file name: ", handler.Filename)
	fmt.Println("File size: ", handler.Size)
	fmt.Println("MIME headder: ", handler.Header)

	//write to a temp file
	tempFile, err := ioutil.TempFile("temp_Contracts", "upload-*.pdf")
	if err != nil {
		fmt.Println("Error creating a temp file for the contract:")
		fmt.Println(err)
		return
	}
	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Error creating a temp file for the contract:")
		fmt.Println(err)
		return
	}

	tempFile.Write(fileBytes)

	//success
	fmt.Fprintln(w, "Success uploading file!")
}

func setUpRoutes() {
	// request handlers
	http.HandleFunc("/api/login", auth.HandleLogin())
	// http.HandleFunc("/api/logout", auth.HandleLogout())
	http.HandleFunc("/api/create-account", auth.HandleRegister())
	http.HandleFunc("/api/get-contract", contract.HandleGetContract())
	http.HandleFunc("/api/upload", uploadFile)
}

func databaseInit() {
	// database inits
	auth.InitAuth("./user_database.db", false)
	contract.InitContractDB("./contract_database.db", false)
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
