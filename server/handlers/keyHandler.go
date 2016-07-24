package handlers

import (
	"github.com/gorilla/mux"
	"github.com/codegangsta/negroni"
	"net/http"
	"github.com/GerardSoleCa/PubKeyManager/domain"
	"github.com/op/go-logging"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/GerardSoleCa/PubKeyManager/server/utils"
	"github.com/gorilla/sessions"
)

var glog = logging.MustGetLogger("keys")

type KeyInteractor interface {
	AddKey(key *domain.Key) (error)
	GetKeys() ([]domain.Key)
	GetUserKeys(user string) ([]domain.Key)
	DeleteKey(id int64) (error)
}

type KeyServiceHandler struct {
	utils.HttpUtils
	KeyInteractor KeyInteractor
	Session       *sessions.CookieStore
}

func ConfigureKeysRouter(router *mux.Router, interactor KeyInteractor, session *sessions.CookieStore) {
	glog.Info("ConfigureKeysRouter")

	handler := KeyServiceHandler{}
	handler.KeyInteractor = interactor
	handler.Session = session

	router.Path("/keys/{user}").HandlerFunc(handler.getUserKeys).Methods("GET")
	secureRouter := mux.NewRouter().PathPrefix("/keys").Subrouter().StrictSlash(false)
	router.PathPrefix("/keys").Handler(negroni.New(
		negroni.HandlerFunc(handler.AuthMiddleware),
		negroni.Wrap(secureRouter),
	))
	secureRouter.Path("/{user}").HandlerFunc(handler.putKeys).Methods("PUT")
	secureRouter.Path("/").HandlerFunc(handler.getKeys).Methods("GET")
}

func (handler KeyServiceHandler) getKeys(rw http.ResponseWriter, q *http.Request) {
	glog.Info("ConfigureKeysRouter -> getKeys")
	keys := handler.KeyInteractor.GetKeys()
	if len(keys) == 0 {
		handler.BadRequest(rw)
		return
	} else {
		glog.Infof("Retrieving %d keys", len(keys))
	}
	var response bytes.Buffer
	encoder := json.NewEncoder(&response)
	encoder.Encode(keys)
	handler.Ok(rw, response.String())
}

func (handler KeyServiceHandler) getUserKeys(rw http.ResponseWriter, q *http.Request) {
	glog.Info("ConfigureKeysRouter -> getUserKeys")
	user := mux.Vars(q)["user"]
	if user == "" {
		handler.BadRequest(rw)
		return
	}
	keys := handler.KeyInteractor.GetUserKeys(user)
	if len(keys) == 0 {
		handler.BadRequest(rw)
		return
	} else {
		glog.Infof("Retrieving %d keys", len(keys))
	}

	rw.WriteHeader(200)
	for _, k := range keys {
		fmt.Fprintln(rw, k.Key)
	}
}

func (handler KeyServiceHandler) putKeys(rw http.ResponseWriter, q *http.Request) {
	glog.Info("ConfigureKeysRouter -> putKeys")
	userQuery := mux.Vars(q)["user"]
	if userQuery == "" {
		handler.BadRequest(rw)
		return
	}
	key := &domain.Key{}
	if handler.ParseBody(q.Body, key) != nil {
		handler.BadRequest(rw)
		return
	}
	key.User = userQuery
	key.CalculateFingerprint()

	if err := handler.KeyInteractor.AddKey(key); err != nil {
		handler.ErrorResponse(rw, &utils.ApiError{Code: 500, Err: err.Error()})
		return
	}
	handler.CreatedWithBody(rw, key)
}

func (handler KeyServiceHandler) AuthMiddleware(rw http.ResponseWriter, q *http.Request, next http.HandlerFunc) {
	glog.Infof("AuthMiddleware")
	session, _ := handler.Session.Get(q, "authenticated")
	if _, ok := session.Values["id"]; !ok {
		handler.Unauthorized(rw)
		return
	}
	if _, ok := session.Values["name"]; !ok {
		handler.Unauthorized(rw)
		return
	}
	next(rw, q)
}