package contract

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

type ContractList struct {
	Contracts []Contract
}

var db *gorm.DB

func (clist *ContractList) AddContract(contract Contract) []Contract {
	clist.Contracts = append(clist.Contracts, contract)
	return clist.Contracts
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		//log.Fatal(err)
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
		ContractID:      "00000001",
		ContractType:    "agreement",
		ContractName:    "a@a.a_00000001.pdf",
		DateCreated:     "1/1/2020",
		TerminationDate: "3/2/2024",
		ValidSigniture:  true,
		PaymentType:     "credit",
		AmountPaid:      0,
		AmountOwed:      5,
		AttorneyName:    "Bob",
		AttorneyEmail:   "a@a.a",
		ClientName:      "Nick",
		ClientEmail:     "nick@yahoo.com"}
	db.FirstOrCreate(&contract)
	//fmt.Println("added test contract")

	// contract2 := Contract{
	// 	ContractID:      "00000001",
	// 	ContractType:    "example contract",
	// 	ContractName:    "example1.pdf",
	// 	DateCreated:     "3/4/2022",
	// 	TerminationDate: "3/4/2023",
	// 	ValidSigniture:  true,
	// 	PaymentType:     "credit",
	// 	AmountPaid:      50.0,
	// 	AmountOwed:      150.0,
	// 	AttorneyName:    "John",
	// 	AttorneyEmail:   "John.law@uflaw.edu",
	// 	ClientName:      "Smith",
	// 	ClientEmail:     "smith@comcast.net"}
	// db.FirstOrCreate(&contract2)

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

		// bring in database
		contract := Contract{}

		// if contract id = * get all contracts for the user
		if contractID == "*" {
			contracts := []Contract{}
			clist := ContractList{contracts}
			rows, err := db.Model(&contract).Rows()
			if err != nil {
				fmt.Println("error finding contracts for Attorney: ", username, " ", err)
				http.Error(w, "error finding contracts", http.StatusNotFound)
				return
			}
			defer rows.Close()

			// itterate rows
			for rows.Next() {
				var contract Contract
				db.ScanRows(rows, &contract)

				clist.AddContract(contract)
			}

		}

		// Bring the contract in
		db.Where(&Contract{ContractID: contractID, AttorneyName: username}).Find(&contract)

		// if contract id = * get all contracts for the user
		if contractID == "*" {
			contracts := []Contract{}
			clist := ContractList{contracts}
			rows, err := db.Model(&contract).Where("AttorneyName = ?", username).Rows()
			if err != nil {
				fmt.Println("error finding contracts for Attorney: ", username)
				http.Error(w, "error finding contracts", http.StatusNotFound)
				return
			}
			defer rows.Close()

			// itterate rows
			for rows.Next() {
				var contract Contract
				db.ScanRows(rows, &contract)

				clist.AddContract(contract)
			}
			json.NewEncoder(w).Encode(clist)

		} else {
			// Bring the contract in
			db.Where(&Contract{ContractID: contractID}).Find(&contract)

			//check if contract exists
			fmt.Println("Username", username, "ContractID", contractID, "Contract: ",contract.ContractID);
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
}

//return contract id
func HandleFileUpload() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Uploading file...\n")

		//parse input
		fmt.Println("parsing input")
		r.ParseMultipartForm(10 << 20)
		fmt.Println("input parsed")

		//retreive files
		file, handler, err := r.FormFile("contract")
		if err != nil {
			//fmt.Println("error parsing file")
			http.Error(w, "error parsing file", http.StatusBadRequest)
			return
		}

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
			http.Error(w, fmt.Sprintf("Could not find session"), http.StatusNotFound)
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

		//get atorney id from session
		val := session.Values["Email"]
		str := fmt.Sprintf("%v", val)
		contract.AttorneyEmail = "fakeaccount@fakeaccount.com"
		fmt.Println("email for upload", contract.AttorneyEmail)
		//get attorney name from userss
		val = session.Values["Name"]
		fmt.Println("Name ", val)
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
		fmt.Println(contract.AttorneyName)

		err = os.WriteFile("./contract_store/"+contract.ContractName, fileBytes, 0644)
		if err != nil {
			http.Error(w, "Form parse error", http.StatusBadRequest)
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

func HandleCountContracts() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// get ur query params
		var attorney_email = r.URL.Query().Get("attorney_email")
		// queery db for number of entries and add one for the contract id
		count := 0
		db.Model(&Contract{}).Where("attorney_email = ?", attorney_email).Count(&count)

		// write result to response
		fmt.Fprintln(w, count)
		http.StatusText(http.StatusOK)
	}
}
