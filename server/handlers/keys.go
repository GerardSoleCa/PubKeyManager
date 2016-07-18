package handlers

import (
	"github.com/gorilla/mux"
	"github.com/codegangsta/negroni"
	"net/http"
	"github.com/GerardSoleCa/PubKeyManager/database"
	"fmt"
	"encoding/json"
	"github.com/GerardSoleCa/PubKeyManager/server/responses"
	"github.com/GerardSoleCa/PubKeyManager/utils"
)

func ConfigureKeysRouter(router *mux.Router){
	keysRouter := mux.NewRouter().PathPrefix("/keys").Subrouter().StrictSlash(false)
	router.PathPrefix("/keys").Handler(negroni.New(
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
	keys := database.GetUserKeys(user)
	rw.WriteHeader(200)
	for _, value := range keys {
		fmt.Fprintln(rw, value.Key)
	}
}

func putKeys(rw http.ResponseWriter, q *http.Request) {
	key := &database.Key{}
	decoder := json.NewDecoder(q.Body)
	err := decoder.Decode(key)
	if err != nil {
		responses.BadRequest(rw)
		return
	}
	key.User = mux.Vars(q)["user"]
	key.Fingerprint = utils.KeyFingerprint([]byte(key.Key))
	if err := database.AddUserKey(key); err != nil {
		responses.ErrorResponse(rw, &responses.ApiError{Code: 500, Err: err.Error()})
	} else {
		responses.Created(rw)
	}
}