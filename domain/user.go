package domain

import (
	"fmt"
	"github.com/op/go-logging"
	"golang.org/x/crypto/bcrypt"
)

var glog = logging.MustGetLogger("User")

// UserRepository interface
type UserRepository interface {
	Store(user *User) error
	GetUser(user string) (User, error)
	Update(user *User) error
	Delete(user string) error
	Count() (int, error)
}

// User struct
type User struct {
	Id       int64  `json:"-"`
	Username string `json:"username"`
	Password string `json:"password,omitempty"`
}

// HashPassword function contained on User
func (u *User) HashPassword() {
	password, _ := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	u.Password = string(password)
}

// CheckPassword function contained on User
func (u *User) CheckPassword(password string) bool {
	fmt.Println("User: ", u)
	glog.Debugf("User > CheckPassowrd :: %d", len([]byte(u.Password)))
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		glog.Debugf("Error > User > CheckPassword :: %s", err.Error())
		return false
	}
	return true
}
