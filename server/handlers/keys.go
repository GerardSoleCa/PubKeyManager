package handlers

import (
	"github.com/gorilla/mux"
	"github.com/codegangsta/negroni"
	"net/http"
	"github.com/GerardSoleCa/PubKeyManager/database"
	"fmt"
	"github.com/GerardSoleCa/PubKeyManager/server/responses"
	"github.com/GerardSoleCa/PubKeyManager/utils"
)

func ConfigureKeysRouter(router *mux.Router) {
	keysRouter := mux.NewRouter().PathPrefix("/keys").Subrouter().StrictSlash(false)
	router.PathPrefix("/keys").Handler(negroni.New(
		negroni.HandlerFunc(TokenExistsMiddleware),
		negroni.Wrap(keysRouter),
	))
	setupPaths(keysRouter)
}

func setupPaths(router *mux.Router) {
	router.Path("/{user}").HandlerFunc(getKeys).Methods("GET")
	router.Path("/{user}").HandlerFunc(putKeys).Methods("PUT")
	//keysRouter.Path("/{user}/{fingerprint}").HandlerFunc().Methods("DELETE")
}

func getKeys(rw http.ResponseWriter, q *http.Request) {
	user := mux.Vars(q)["user"]
	keys := database.GetKeyByUser(user)
	rw.WriteHeader(200)
	for _, value := range keys {
		fmt.Fprintln(rw, value.Key)
	}
}

func putKeys(rw http.ResponseWriter, q *http.Request) {
	key := &database.Key{}
	if utils.ParseBody(q.Body, key) != nil {
		responses.BadRequest(rw)
		return
	}
	key.User = mux.Vars(q)["user"]
	if err := key.Save(); err != nil {
		responses.ErrorResponse(rw, &responses.ApiError{Code: 500, Err: err.Error()})
	} else {
		responses.CreatedWithBody(rw, key)
	}
}