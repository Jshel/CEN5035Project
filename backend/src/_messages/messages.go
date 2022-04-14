package messages

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/jinzhu/gorm"

	// load the sqlite driver
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Message struct{
	Sender	string `json:"sender"`
	Receiver	string `json:"receiver"`
	Message	string `json:"message"`
	Time 	int64 `json:"time"`
}

type MessageList struct {
	Messages []Message
}

func (mlist *MessageList) AddMessage(message Message) []Message {
    mlist.Message = append(mlist.Message, message)
    return mlist.Message
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
	return func(w http.ResponseWriter, r *http.Request)
		// convert request to message data
		var Message message
		err := json.NewDecoder(r.Body).Decode(&Message)

		if err != nil {
			http.Error(w, "Difformed message", http.StatusBadRequest)
			fmt.Println("Error: Difformed message request")
			return
		}

		db.Save(&message)
		http.StatusText(http.StatusOK)
}

func HandleGetMessage() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request)
		// mesage list and count
		messages := []Message{}
		mlist := MessageList{messages}

		// get url params
		var sender = r.URL.Query().Get("sender")
		var receiver = r.URL.Query().Get("receiver")
		var n = r.URL.Query().Get("n")

		// get the message database
		message := Message{}

		// get the messages with these users
		rows, err := db.Model(&Message{}).Where("sender = ?", sender).Where("receiver = ?", receiver).Rows()
		defer rows.Close()
		
		// itterate rows
		for rows.Next() {
			if n != 0 {
				var message Message
				db.ScanRows(row, &message)

				mlist.AddMessage(message)
				n--
			} else {
				break
			}
		}

		// encode list of messages 
		json.NewEncoder(w).Encode(mlist)
	
}
