package usecases

import (
	"errors"
	"github.com/GerardSoleCa/PubKeyManager/domain"
)

// KeyInteractor. Holds a KeyRepository instance
type KeyInteractor struct {
	KeyRepository domain.KeyRepository
}

// AddKey. Function contained on KeyInteractor
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

// GetKeys. Function contained on KeyInteractor
func (interactor KeyInteractor) GetKeys() []domain.Key {
	return interactor.KeyRepository.GetKeys()
}

// GetUserKeys. Function contained on KeyInteractor
func (interactor KeyInteractor) GetUserKeys(user string) []domain.Key {
	return interactor.KeyRepository.GetUserKeys(user)
}

// DeleteKey. Function contained on KeyInteractor
func (interactor KeyInteractor) DeleteKey(id int64) error {
	return interactor.KeyRepository.Delete(id)
}
