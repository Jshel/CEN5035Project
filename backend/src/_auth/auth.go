package auth

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"

	"github.com/jinzhu/gorm"
	// load the sqlite driver
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	key   = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
)

type userRegister struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type userLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Hash     []byte `json:"hash"`
}

// keep track of the database
var db *gorm.DB

// bcrypt strength
var strength = 11

func InitAuth(sqliteFile string, debugSQL bool) {
	_db, err := gorm.Open("sqlite3", sqliteFile)
	if err != nil {
		panic(fmt.Sprintf("Could not open management database %v", err))
	}
	if debugSQL {
		db = _db.Debug()
	} else {
		db = _db
	}

	// migrate schemas
	db.AutoMigrate(&User{})

}

//HandleLogin loggs in the user attaches a session COOKIE to the reply. Returns WhoAmI info
func HandleLogin() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		// convert request to registration data
		var login userLogin
		err := json.NewDecoder(r.Body).Decode(&login)

		if err != nil {
			http.Error(w, "Difformed login request", http.StatusBadRequest)
			fmt.Println("Error: Difformed login request")
			return
		}

		// Bring the user from the database
		user := User{}
		db.Where(&User{Email: login.Email}).Find(&user)

		if user.Email != login.Email {
			http.Error(w, fmt.Sprintf("User %s does not exist", login.Email), http.StatusForbidden)
			fmt.Println("ERROR: ", login.Email, " does not exist")
			return
		}

		// check the password
		err = bcrypt.CompareHashAndPassword([]byte(user.Hash), []byte(login.Password))
		if err != nil {
			http.Error(w, fmt.Sprintf("Password for user %s is incorrect", login.Email), http.StatusForbidden)
			fmt.Println("password failure for: ", login.Email, " password: ", login.Password)
			return
		}

		// existing session: Get() always returns a session, even if empty.
		session, err := store.Get(r, "cookie-name")

		if err == nil {
			session.Values["id"] = login.Email
			session.Values["authenticated"] = true
			err = session.Save(r, w)
			fmt.Println("Login success: ", login.Email)
		}
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			http.Error(w, fmt.Sprintf("Could not setup session for %s user", login.Email), http.StatusConflict)
			return
		}

		http.StatusText(http.StatusOK)

	}
}

func HandleLogout() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		session, err := store.Get(r, "cookie-name")
		if err == nil {
			session.Values["authenticated"] = false
			session.Save(r, w)
			fmt.Println("Logout success: ", session.Values["id"])
			http.Error(w, "Successfull logout", http.StatusOK)
		} else {
			http.Error(w, "No session found", http.StatusNotFound)
		}
	}
}

func HandleRegister() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// convert request to registration data
		var registration userRegister

		err := json.NewDecoder(r.Body).Decode(&registration)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		fmt.Println("Generating password hash")
		hash, err := bcrypt.GenerateFromPassword([]byte(registration.Password), strength)

		if err != nil {
			http.Error(w, "Password hashing failed", http.StatusBadRequest)
			fmt.Println("Password Hashing failed")
			return
		}
		fmt.Println("Hashed password: ", string(hash))

		var user = User{}
		user.Name = registration.Name
		user.Username = registration.Username
		user.Hash = hash
		user.Email = registration.Email

		db.Save(&user)

		session, err := store.Get(r, "cookie-name")
		if err == nil {
			session.Values["id"] = registration.Email
			session.Values["authenticated"] = true
			err = session.Save(r, w)
			fmt.Println("Login success: ", registration.Email)
		}
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			http.Error(w, fmt.Sprintf("Could not setup session for %s user", registration.Email), http.StatusConflict)
			return
		}

		http.StatusText(http.StatusOK)
	}
}
