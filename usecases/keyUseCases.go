package usecases

import (
	"github.com/GerardSoleCa/PubKeyManager/domain"
	"errors"
)

type KeyInteractor struct {
	KeyRepository domain.KeyRepository
}

func (interactor KeyInteractor) AddKey(key *domain.Key) error {
	if len(key.Title) == 0 {
		return errors.New("A title is required to store a ssh key")
	}
	if len(key.User) == 0 {
		return errors.New("A user is required to store a ssh key")
	}
	if len(key.Key) == 0 {
		return errors.New("A key is required")
	}
	if err := key.CalculateFingerprint(); err != nil {
		return err
	}
	return interactor.KeyRepository.Store(key)
}

func (interactor KeyInteractor) GetKeys() []domain.Key {
	return interactor.KeyRepository.GetKeys()
}

func (interactor KeyInteractor) GetUserKeys(user string) []domain.Key {
	return interactor.KeyRepository.GetUserKeys(user)
}

func (interactor KeyInteractor) DeleteKey(id int64) error {
	return interactor.KeyRepository.Delete(id)
}
