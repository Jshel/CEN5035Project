package contract

import (
	"encoding/json"
	"fmt"

	"net/http"

	"github.com/jinzhu/gorm"

	// load the sqlite driver
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Contract struct {
	ContractID   string `json:"contract_ID"`
	ContractType string `json:"contract_type"`

	DateCreated     string `json:"date_created"`
	TerminationDate string `json:"termination_date"`

	ValidSigniture bool `json:"valid_signiture"`

	PaymentType string  `json:"payment_type"`
	AmountPaid  float32 `json:"amount_paid"`
	AmountOwed  float32 `json:"amount_owed"`

	AttorneyName string `json:"attorney_name"`
	AttorneyID   string `json:"attorney_ID"`

	ClientName string `json:"client_name"`
	ClientID   string `json:"client_ID"`
}

var db *gorm.DB

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
		DateCreated:     "3/2/2022",
		TerminationDate: "3/2/2023",
		ValidSigniture:  true,
		PaymentType:     "cash",
		AmountPaid:      0.0,
		AmountOwed:      100.0,
		AttorneyName:    "Bob",
		AttorneyID:      "00000001",
		ClientName:      "Alice",
		ClientID:        "00000002"}
	db.FirstOrCreate(&contract)

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
