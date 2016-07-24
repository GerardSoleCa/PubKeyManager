package main

import (
	"github.com/GerardSoleCa/PubKeyManager/infrastructure"
	"github.com/GerardSoleCa/PubKeyManager/interfaces"
	"github.com/GerardSoleCa/PubKeyManager/server"
	"github.com/GerardSoleCa/PubKeyManager/usecases"
)

func main() {
	dbHandler := infrastructure.NewSqliteHandler("PubKeyManager.db")

	keyInteractor := new(usecases.KeyInteractor)
	keyInteractor.KeyRepository = interfaces.NewDbKeyRepository(dbHandler)

	userInteractor := new(usecases.UserInteractor)
	userInteractor.UserRepository = interfaces.NewDbUserRepository(dbHandler)

	server := &server.Server{}
	server.KeyInteractor = keyInteractor
	server.UserInteractor = userInteractor
	server.Start()
}
