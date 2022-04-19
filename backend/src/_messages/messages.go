package messages

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/jinzhu/gorm"

	// load the sqlite driver
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Message struct {
	Sender   string `json:"sender"`
	Receiver string `json:"receiver"`
	Message  string `json:"message"`
	Time     string `json:"time"`
}

type MessageList struct {
	Messages []Message
}

// keep track of the database
var db *gorm.DB

func (mlist *MessageList) AddMessage(message Message) []Message {
	mlist.Messages = append(mlist.Messages, message)
	return mlist.Messages
}

func InitMessageDB(sqliteFile string, debugSQL bool) {
	_db, err := gorm.Open("sqlite3", sqliteFile)
	if err != nil {
		panic(fmt.Sprintf("Could not open management database %v", err))
	}
	if debugSQL {
		db = _db.Debug()
	} else {
		db = _db
	}

	db.AutoMigrate(&Message{})
}

func HandleSendMessage() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// convert request to message data
		var message Message
		err := json.NewDecoder(r.Body).Decode(&message)

		if err != nil {
			fmt.Fprintln(w, err)
			http.Error(w, "Difformed message", http.StatusBadRequest)
			fmt.Println("Error: Difformed message request")
			return
		}
		fmt.Println(message)
		fmt.Fprintln(w, "message saved!")
		db.Save(&message)
		http.StatusText(http.StatusOK)
	}
}

func HandleGetMessage() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// mesage list and count
		messages := []Message{}
		mlist := MessageList{messages}

		// get url params
		var sender = r.URL.Query().Get("sender")
		var receiver = r.URL.Query().Get("receiver")
		var n = r.URL.Query().Get("n")

		N, err := strconv.Atoi(n)

		if err != nil {
			http.Error(w, "n is not an int", http.StatusBadRequest)
			fmt.Println("n is not an int")
			return
		}

		// get the message database
		message := Message{}

		// if sender or reciever contains a * return all messages that go with the wildcard and desired user
		// contract and messages num. post with email
		if sender == "*" {
			rows, err := db.Model(&message).Where("receiver = ?", receiver).Rows()
			if err != nil {
				fmt.Println("error finding messages for receiver: ", receiver)
				http.Error(w, "error finding messages", http.StatusNotFound)
				return
			}
			defer rows.Close()

			// itterate rows
			for rows.Next() {
				if N != 0 {
					var message Message
					db.ScanRows(rows, &message)

					mlist.AddMessage(message)
					N--
				} else {
					break
				}
			}
			// encode list of messages
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(mlist)
		} else if receiver == "*" {
			rows, err := db.Model(&message).Where("sender = ?", sender).Rows()
			if err != nil {
				fmt.Println("error finding messages for sender: ", sender)
				http.Error(w, "error finding messages", http.StatusNotFound)
				return
			}
			defer rows.Close()

			// itterate rows
			for rows.Next() {
				if N != 0 {
					var message Message
					db.ScanRows(rows, &message)

					mlist.AddMessage(message)
					N--
				} else {
					break
				}
			}
			// encode list of messages
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(mlist)
		} else {
			// get the messages with these users
			rows, err := db.Model(&message).Where("sender = ?", sender).Where("receiver = ?", receiver).Rows()
			if err != nil {
				fmt.Println("error finding messages for sender: ", sender, " and receiver: ", receiver)
				http.Error(w, "error finding messages", http.StatusNotFound)
				return
			}
			defer rows.Close()

			// itterate rows
			for rows.Next() {
				if N != 0 {
					var message Message
					db.ScanRows(rows, &message)

					mlist.AddMessage(message)
					N--
				} else {
					break
				}
			}

			// encode list of messages
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(mlist)
		}
	}
}

func HandleCountMessages() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// get ur query params
		var attorney_email = r.URL.Query().Get("attorney_email")
		// queery db for number of entries and add one for the contract id
		count := 0
		db.Model(&Message{}).Where("receiver = ?", attorney_email).Count(&count)

		// write result to response
		fmt.Fprintln(w, count)
		http.StatusText(http.StatusOK)
	}
}
