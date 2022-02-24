package auth

import (
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/jinzhu/gorm"
	// load the sqlite driver
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var cookie = "TC_Audit"

type userRegister struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type userLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// keep track of the database
var db *gorm.DB

// variable to keep track of the base URL
var baseURL string

// bcrypt strength
var strength = 11

func initAuth(sqliteFile string, _baseURL string, debugSQL bool) {
	_db, err := gorm.Open("sqlite3", sqliteFile)
	if err != nil {
		panic(fmt.Sprintf("Could not open management database %v", err))
	}
	if debugSQL {
		db = _db.Debug()
	} else {
		db = _db
	}
	baseURL = _baseURL

	// migrate schemas
	db.AutoMigrate(&User{})

}

// HandleLogin loggs in the user attaches a session COOKIE to the reply. Returns WhoAmI info
func HandleLogin(w http.ResponseWriter, r *http.Request) {
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
	db.Where(&User{Username: login.Username}).Find(&user)

	if user.Username != login.Username {
		http.Error(w, fmt.Sprintf("User %s does not exist", login.Username), http.StatusForbidden)
		fmt.Println("ERROR: %s does not exist", login.Username)
		return
	}

	// check the password
	err = bcrypt.CompareHashAndPassword([]byte(user.Hash), []byte(login.Password))
	if err != nil {
		http.Error(w, fmt.Sprintf("Password for user %s is incorrect", login.Username), http.StatusForbidden)
		return
	}

	// setup the session and tell user that everything is fine
	session, err := cookieStore.New(r, cookie)
	if err == nil {
		session.Values["id"] = login.Username
		err = session.Save(r, w)
	}
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		http.Error(w, fmt.Sprintf("Could not setup session for %s user", login.Username), http.StatusConflict)
		return
	}

	util.EncodeJSONResponse(w, user)
}

func HandleLogout(w http.ResponseWriter, r *http.Request) {
	session, err := cookieStore.Get(r, cookie)
	if err == nil {
		// delete the cookie
		session.Options.MaxAge = -1
		session.Save(r, w)

		http.Error(w, "Successfull logout", http.StatusOK)
	} else {
		http.Error(w, "No session found", http.StatusNotFound)
	}
}

func HandleRegister(w http.ResponseWriter, r *http.Request) {
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
	fmt.Println("Hashed password: %s", string(hash))

	// Count admins. If none, make this user an admin
	var count int
	db.Model(&User{}).Where("role = ?", "admin").Count(&count)

	var user = User{}
	db.FirstOrCreate(&user, User{Email: registration.Email})
	user.Name = registration.Name
	user.Hash = string(hash)
	if count == 0 {
		user.Role = "admin"
	} else {
		user.Role = "user"
	}

	db.Save(&user)

	// Redirect to main page
	http.Error(w, "Registration successfull", http.StatusOK)
	fmt.Println("Registration successfull")
}
