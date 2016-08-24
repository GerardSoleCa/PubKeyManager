package usecases

import (
	"errors"
	"github.com/GerardSoleCa/PubKeyManager/domain"
	"github.com/op/go-logging"
)

// ErrNotAuthenticated. Response for non authenticated user
var ErrNotAuthenticated = errors.New("User could not be authenticated")

var glog = logging.MustGetLogger("UserUseCases")

// UserInteractor struct
type UserInteractor struct {
	UserRepository domain.UserRepository
}

// AddUser. Function contained on UserInteractor
func (interactor UserInteractor) AddUser(user *domain.User) error {
	count, err := interactor.UserRepository.Count()
	if err != nil {
		return err
	}
	if count >= 1 {
		return errors.New("Only one user is able to be in the system")
	}
	if len(user.Username) < 3 {
		return errors.New("Username must have at least 3 characters")
	}
	if len(user.Password) < 4 {
		return errors.New("Password must have at least 5 characters")
	}
	user.HashPassword()
	return interactor.UserRepository.Store(user)
}

// AuthenticateUser. Function contained on UserInteractor
func (interactor UserInteractor) AuthenticateUser(username, password string) (domain.User, error) {
	glog.Debugf("UserUseCase > AuthenticateUser :: %s, %s", username, password)
	u, err := interactor.UserRepository.GetUser(username)
	if err != nil {
		glog.Debugf("Error > UserUseCase > AuthenticateUser [%s] :: %s", username, err)
		return u, err
	}
	if u.CheckPassword(password) {
		return u, nil
	}
	return domain.User{}, ErrNotAuthenticated
}
