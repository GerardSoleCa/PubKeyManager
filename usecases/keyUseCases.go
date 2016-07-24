package usecases

import (
	"github.com/GerardSoleCa/PubKeyManager/domain"
)

type KeyInteractor struct {
	KeyRepository domain.KeyRepository
}

func (interactor KeyInteractor) AddKey(key *domain.Key) (error) {
	return interactor.KeyRepository.Store(key)
}

func (interactor KeyInteractor) GetKeys() ([]domain.Key) {
	return interactor.KeyRepository.GetKeys()
}

func (interactor KeyInteractor) GetUserKeys(user string) ([]domain.Key) {
	return interactor.KeyRepository.GetUserKeys(user)
}

func (interactor KeyInteractor) DeleteKey(id int64) (error) {
	return interactor.KeyRepository.Delete(id)
}
