package main

import (
	auth "attorneyManager/_auth"
	contract "attorneyManager/_contract"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/jinzhu/gorm"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var db *gorm.DB

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
}

func uploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Uploading file...\n")

	//parse input
	r.ParseMultipartForm(10 << 20)

	//retreive files
	file, handler, err := r.FormFile("contract")
	checkErr(err)

	defer file.Close()
	//print file info to output
	fmt.Println("Uploaded file name: ", handler.Filename)
	fmt.Println("File size: ", handler.Size)
	fmt.Println("MIME headder: ", handler.Header)

	//read file to byte arrray
	fileBytes, err := ioutil.ReadAll(file)
	checkErr(err)

	//get the session
	// var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))
	// session, err := store.Get(r, "session-login")
	if err != nil {
		//no session found
		fmt.Println("could not get session")
		//likley not logged in redirect to login page.
		http.Redirect(w, r, "http://localhost/4200/login", http.StatusBadRequest)
		checkErr(err)
	}

	filename := "./contract_store/file.pdf"
	err = os.WriteFile(filename, fileBytes, 0644)
	if err != nil {
		fmt.Println("Error getting file from form:")
		fmt.Println(err)
		return
	}

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
