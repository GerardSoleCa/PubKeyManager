package domain

import (
	"golang.org/x/crypto/bcrypt"
	"github.com/op/go-logging"
	"fmt"
)

var glog = logging.MustGetLogger("User")

type UserRepository interface {
	Store(user *User) (error)
	GetUser(user string) (User, error)
	Update(user *User) (error)
	Delete(user string) (error)
}

type User struct {
	Id       int64         `json:"-"`
	Username string        `json:"username"`
	Password string        `json:"password,omitempty"`
}

func (u *User) HashPassword() {
	password, _ := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	u.Password = string(password)
}

func (u *User) CheckPassword(password string) (bool) {
	fmt.Println("User: ", u)
	glog.Debugf("User > CheckPassowrd :: %d", len([]byte(u.Password)))
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		glog.Debugf("Error > User > CheckPassword :: %s", err.Error())
		return false
	} else {
		return true
	}
}