package contract

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"

	// load the sqlite driver
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type ContractDraft struct {
	UserID          string `json:"userID"`
	AttorneyList    string `json:"attorney_list"`
	ClientList      string `json:"client_list"`
	ContractTitle   string `json:"contract_title"`
	Date            string `json:"date"`
	TerminationDate string `json:"termination_date"`
	PaymentType     string `json:"payment_type"`
	OtherNotes      string `json:"notes"`
}

func InitContractDraftDB(sqliteFile string, debugSQL bool) {
	_db, err := gorm.Open("sqlite3", sqliteFile)
	if err != nil {
		panic(fmt.Sprintf("Could not open management database %v", err))
	}
	if debugSQL {
		db = _db.Debug()
	} else {
		db = _db
	}

	db.AutoMigrate(&ContractDraft{})
}

func HandleGetContractDraft() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		var contract ContractDraft
		err := json.NewDecoder(r.Body).Decode(&contract)

		if err != nil {
			http.Error(w, "Failed to send Contract Request", http.StatusBadRequest)
			fmt.Println("Error in sendinf Contract form")
			return
		}

		var newcontract = ContractDraft{}

		newcontract.UserID = contract.UserID
		newcontract.AttorneyList = contract.AttorneyList
		newcontract.ClientList = contract.ClientList
		newcontract.ContractTitle = contract.ContractTitle
		newcontract.Date = contract.Date
		newcontract.TerminationDate = contract.TerminationDate
		newcontract.PaymentType = contract.PaymentType
		newcontract.OtherNotes = contract.OtherNotes

		db.Save(&newcontract)
		http.StatusText(http.StatusOK)
	}
}
