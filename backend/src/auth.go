package rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	//"github.com/tera-insights/tiCrypt-logview/backend/reports"
	//"github.com/tera-insights/tiCrypt-logview/backend/util"
	"golang.org/x/crypto/bcrypt"

	"github.com/jinzhu/gorm"

	// load the sqlite driver
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var cookie = "TC_Audit"

type userRegister struct {
	Email      string `json:"email"`
	Name       string `json:"name"`
	Password   string `json:"password"`
	Invitation string `json:"invitation"`
}

type userLogin struct {
	Email    string `json:"email"`
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
	db.AutoMigrate(&UserInvite{})
	db.AutoMigrate(&QueryFavorite{})
	db.AutoMigrate(&QueryExec{})
	db.AutoMigrate(&ReportExec{})
	db.AutoMigrate(&AccessToken{})
	db.AutoMigrate(&reports.ReportInfo{})
	db.AutoMigrate(&reports.QueryInfo{})
	db.Model(&reports.QueryInfo{}).AddForeignKey("report_info_id", "report_infos(id)", "CASCADE", "CASCADE")

	// We can now initialize the reports
	loadReports()
}

// GetUser returns the user from the session info
func GetUser(r *http.Request) (*User, error) {
	user := User{}
	session, err := cookieStore.Get(r, cookie)

	if err != nil {
		return nil, err
	}

	if email, ok := session.Values["id"].(string); ok {
		db.Where(&User{Email: email}).Find(&user)

		if user.Email == email {
			return &user, nil
		}
		return nil, fmt.Errorf("No session found for user %s", email)
	}
	return nil, fmt.Errorf("No session found")
}

// GetAdmin ensures the user exists and it is an admin
func GetAdmin(r *http.Request) (*User, error) {
	user, err := GetUser(r)
	if err == nil && user.Role != "admin" {
		return nil, fmt.Errorf("User %s is not an admin", user.Email)
	}
	return user, err
}

// GetUserEmail recoveres user email from session, nil if error
// and fills in the reply. This works even for token based sessions
func GetUserEmail(w http.ResponseWriter, r *http.Request) string {
	// Try first to find a token
	token := r.Header.Get("x-access-token")
	if token != "" {
		if verifyToken(token) {
			return "token:" + hashToken(token) // we have a session
		}

		http.Error(w, fmt.Errorf("Invalid or expired token %s", token).Error(), http.StatusForbidden)
		return ""
	}

	user, err := GetUser(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return ""
	}
	return user.Email
}

// NoSession replies with standard answer if no session
// Usage: if NoSession(w,r) return
func NoSession(w http.ResponseWriter, r *http.Request) bool {
	// First, determine if this a token-based request
	token := r.Header.Get("x-access-token")
	if token != "" {
		if verifyToken(token) {
			return false // we have a session
		}

		http.Error(w, fmt.Errorf("Invalid or expired token %s", token).Error(), http.StatusForbidden)
		return true
	}

	_, err := GetUser(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return true
	}
	return false
}

// NoAdmin replies with standard answer if user not an admin or no session
func NoAdmin(w http.ResponseWriter, r *http.Request) bool {
	_, err := GetAdmin(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return true
	}
	return false
}

// HandleRegister registers a new user based on an invitation
func HandleRegister(w http.ResponseWriter, r *http.Request) {
	// convert request to registration data
	var registration userRegister
	err := json.NewDecoder(r.Body).Decode(&registration)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// look up the registration
	var invitation UserInvite
	db.Where("invitation = ?", registration.Invitation).Find(&invitation)

	if invitation.Created == 0 {
		http.Error(w, "Invitation not found. Possibly already used.", http.StatusForbidden)
		return
	}
	// figure out if the invitation is expired
	var expired = isExpired(invitation.Expires)
	hash, err := bcrypt.GenerateFromPassword([]byte(registration.Password), strength)

	if err != nil {
		http.Error(w, "Password hashing failed", http.StatusBadRequest)
		return
	}

	if registration.Email == invitation.Email && !expired {
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

		// Delete registration
		db.Delete(&invitation)

		// Redirect to main page
		http.Error(w, "Registration successfull", http.StatusOK)
	} else {
		if expired {
			http.Error(w, "Expired invitation", http.StatusForbidden)
		} else {
			http.Error(w, fmt.Sprintf("User in registration %s and invitation %s do not match", registration.Email, invitation.Email), http.StatusForbidden)
		}
	}
}

// HandleWhoAmI returns information about logged in user
func HandleWhoAmI(w http.ResponseWriter, r *http.Request) {
	user, err := GetUser(r)
	if err == nil {
		util.EncodeJSONResponse(w, user)
	} else {
		util.EncodeJSONResponse(w, map[string]string{
			"error": "Not logged in",
		})
	}
}

// HandleLogin loggs in the user attaches a session COOKIE to the reply. Returns WhoAmI info
func HandleLogin(w http.ResponseWriter, r *http.Request) {
	// convert request to registration data
	var login userLogin
	err := json.NewDecoder(r.Body).Decode(&login)

	if err != nil {
		http.Error(w, "Difformed login request", http.StatusBadRequest)
		return
	}

	// Bring the user from the database
	user := User{}
	db.Where(&User{Email: login.Email}).Find(&user)

	if user.Email != login.Email {
		http.Error(w, fmt.Sprintf("User %s does not exist", login.Email), http.StatusForbidden)
		return
	}

	// check the password
	err = bcrypt.CompareHashAndPassword([]byte(user.Hash), []byte(login.Password))
	if err != nil {
		http.Error(w, fmt.Sprintf("Password for user %s is incorrect", login.Email), http.StatusForbidden)
		return
	}

	// setup the session and tell user that everything is fine
	session, err := cookieStore.New(r, cookie)
	if err == nil {
		session.Values["id"] = login.Email
		err = session.Save(r, w)
	}
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		http.Error(w, fmt.Sprintf("Could not setup session for %s user", login.Email), http.StatusConflict)
		return
	}

	util.EncodeJSONResponse(w, user)
}

// HandleLogout closess current session
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
