package contract

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/sessions"
	"github.com/jinzhu/gorm"

	// load the sqlite driver
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Hash     []byte `json:"hash"`
}

type Contract struct {
	ContractID   string `json:"contract_ID"`
	ContractType string `json:"contract_type"`
	ContractName string `json:"contract_path"`

	DateCreated     string `json:"date_created"`
	TerminationDate string `json:"termination_date"`

	ValidSigniture bool `json:"valid_signiture"`

	PaymentType string  `json:"payment_type"`
	AmountPaid  float64 `json:"amount_paid"`
	AmountOwed  float64 `json:"amount_owed"`

	AttorneyName  string `json:"attorney_name"`
	AttorneyEmail string `json:"attorney_email"`

	ClientName  string `json:"client_name"`
	ClientEmail string `json:"client_email"`
}

//used to decode a json request form data
type ContractInit struct {
	ContractType    string  `json:"contract_type"`
	TerminationDate string  `json:"termination_date"`
	PaymentType     string  `json:"payment_type"`
	AmountPaid      float64 `json:"amount_paid"`
	AmountOwed      float64 `json:"amount_owed"`
	ClientEmail     string  `json:"client_email"`
}

var db *gorm.DB

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
}

func InitContractDB(sqliteFile string, debugSQL bool) {
	_db, err := gorm.Open("sqlite3", sqliteFile)
	if err != nil {
		panic(fmt.Sprintf("Could not open management database %v", err))
	}
	if debugSQL {
		db = _db.Debug()
	} else {
		db = _db
	}
	// test entry
	contract := Contract{
		ContractID:      "00000000",
		ContractType:    "example contract",
		ContractName:    "example0.pdf",
		DateCreated:     "3/2/2022",
		TerminationDate: "3/2/2023",
		ValidSigniture:  true,
		PaymentType:     "cash",
		AmountPaid:      0.0,
		AmountOwed:      100.0,
		AttorneyName:    "Bob",
		AttorneyEmail:   "Bob.law@gmail.com",
		ClientName:      "Alice",
		ClientEmail:     "alice@yahoo.com"}
	db.FirstOrCreate(&contract)

	contract2 := Contract{
		ContractID:      "00000001",
		ContractType:    "example contract",
		ContractName:    "example1.pdf",
		DateCreated:     "3/4/2022",
		TerminationDate: "3/4/2023",
		ValidSigniture:  true,
		PaymentType:     "credit",
		AmountPaid:      50.0,
		AmountOwed:      150.0,
		AttorneyName:    "John",
		AttorneyEmail:   "John.law@uflaw.edu",
		ClientName:      "Smith",
		ClientEmail:     "smith@comcast.net"}
	db.FirstOrCreate(&contract2)

	db.AutoMigrate(&Contract{})
}

func HandleGetContract() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// get url params as of now the search will just search for contract id from a given attorney
		// the search for which attorney should be done with a cookie or session not by url query

		// print values
		values := r.URL.Query()
		for k, v := range values {
			fmt.Println(k, " => ", v)
		}

		var username = r.URL.Query().Get("username")
		var contractID = r.URL.Query().Get("contractID")

		// Bring the contract in
		contract := Contract{}
		db.Where(&Contract{ContractID: contractID, AttorneyName: username}).Find(&contract)

		//check if contract exists
		if contract.ContractID != contractID {
			http.Error(w, fmt.Sprintf("Contract ID: %s does not exist", contractID), http.StatusForbidden)
			fmt.Println("ERROR: ", contractID, " does not exist for attorney ", username)
			return
		} else {
			// contract exists
			json.NewEncoder(w).Encode(contract)
		}
	}
}

func HandleFileUpload() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
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
		var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))
		session, err := store.Get(r, "session-login")
		if err != nil {
			//no session found
			fmt.Println("could not get session")
			//likley not logged in redirect to login page.
			http.Redirect(w, r, "http://localhost/4200/login", http.StatusBadRequest)
			checkErr(err)
		}

		//create entry to contract db
		//user input querys
		var contract = Contract{}

		//get user input from query params
		contract.ContractType = r.URL.Query().Get("contract_type")
		contract.TerminationDate = r.URL.Query().Get("termination_date")
		contract.PaymentType = r.URL.Query().Get("payment_type")
		s, err := strconv.ParseFloat(r.URL.Query().Get("ammount_paid"), 64)
		checkErr(err)
		contract.AmountPaid = s
		s, err = strconv.ParseFloat(r.URL.Query().Get("ammount_owed"), 64)
		checkErr(err)
		contract.AmountOwed = s
		contract.ClientEmail = r.URL.Query().Get("client_email")
		contract.ClientName = r.URL.Query().Get("client_name")

		// //decode form data
		// var ContractInit ContractInit
		// err := json.NewDecoder(r.)

		//get atorney id from session
		val := session.Values["Email"]
		str := fmt.Sprintf("%v", val)
		contract.AttorneyEmail = str

		//get attorney name from users
		val = session.Values["Name"]
		str = fmt.Sprintf("%v", val)
		contract.AttorneyName = str

		//queery db for number of entries and add one for the contract id
		count := 0
		db.Model(&Contract{}).Where("attorney_email = ?", contract.AttorneyEmail).Count(&count)
		count++
		hexID := fmt.Sprintf("%0*x", 8, count)
		contract.ContractID = hexID
		contract.DateCreated = time.Now().Format(time.RFC3339)

		//other values
		contract.ValidSigniture = true

		contract.ContractName = contract.AttorneyEmail + "_" + contract.ContractID + ".pdf"

		//save to db
		db.Create(&contract)

		err = os.WriteFile("./contract_store/"+contract.ContractName, fileBytes, 0644)
		if err != nil {
			fmt.Println("Error getting file from form:")
			fmt.Println(err)
			return
		}

		//success
		fmt.Fprintln(w, "Success uploading file!")
	}
}

func HandleFileDownload() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// print values
		values := r.URL.Query()
		for k, v := range values {
			fmt.Println(k, " => ", v)
		}

		// get the url params
		var attorney_email = r.URL.Query().Get("attorney_email")
		var contract_id = r.URL.Query().Get("contract_id")

		// Bring the contract in
		contract := Contract{}
		db.Where(&Contract{ContractID: contract_id, AttorneyEmail: attorney_email}).Find(&contract)

		//check if contract exists
		if contract.ContractID != contract_id {
			http.Error(w, fmt.Sprintf("Contract ID: %s does not exist", contract_id), http.StatusNotFound)
			fmt.Println("ERROR: ", contract_id, " does not exist for attorney with emaili:", attorney_email)
			return
		} else {
			// contract exists get the contract and write it to the response
			fileBytes, err := ioutil.ReadFile("./contract_store/" + contract.ContractName)
			if err != nil {
				checkErr(err)
			}
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/octet-stream")
			w.Write(fileBytes)
			return
		}
	}
}
