package usecases

import (
	"github.com/GerardSoleCa/PubKeyManager/domain"
	"errors"
	"github.com/op/go-logging"
)

var ErrNotAuthenticated = errors.New("User could not be authenticated")

var glog = logging.MustGetLogger("UserUseCases")

type UserInteractor struct {
	UserRepository domain.UserRepository
}

func (interactor UserInteractor) AddUser(user *domain.User) (error) {
	user.HashPassword()
	return interactor.UserRepository.Store(user)
}

func (interactor UserInteractor) AuthenticateUser(username, password string) (domain.User, error) {
	glog.Debugf("UserUseCase > AuthenticateUser :: %s, %s", username, password)
	u, err := interactor.UserRepository.GetUser(username)
	if err != nil {
		glog.Debugf("Error > UserUseCase > AuthenticateUser [%s] :: %s", username, err)
		return u, err
	}
	if u.CheckPassword(password) {
		return u, nil
	} else {
		return domain.User{}, ErrNotAuthenticated
	}
}

