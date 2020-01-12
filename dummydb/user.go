package dummydb

import (
	"log"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

const userActive int32 = 1

// User holds our customer data
type User struct {
	ID        int64  `json:"id"`
	UserName  string `json:"name"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Phone     string `json:"phone"`
	Status    int32  `json:"userstatus"`
}

var userID int64
var userTBL []*User

// NewUser adds a new user to the db, returns the ID
func NewUser(username, password, email string, firstname, lastname string, phone string) int64 {
	userID++
	id := userID
	u := new(User)
	u.ID = id
	u.UserName = username
	u.FirstName = firstname
	u.LastName = lastname
	u.Email = email
	u.Password = encodePassword(password)
	u.Status = userActive
	userTBL = append(userTBL, u)
	return id
}

func encodePassword(password string) string {
	encoded, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
	}
	return string(encoded)
}

// UserByUserName returns pointer to User with specified username
func UserByUserName(username string) *User {
	for _, user := range userTBL {
		if strings.ToLower(user.UserName) == strings.ToLower(username) {
			return user
		}
	}
	return nil
}
