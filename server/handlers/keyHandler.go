package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/GerardSoleCa/PubKeyManager/domain"
	"github.com/GerardSoleCa/PubKeyManager/server/utils"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/op/go-logging"
	"net/http"
	"strconv"
)

var glog = logging.MustGetLogger("keys")

// KeyInteractor interface
type KeyInteractor interface {
	AddKey(key *domain.Key) error
	GetKeys() []domain.Key
	GetUserKeys(user string) []domain.Key
	DeleteKey(id int64) error
}

// KeyServiceHandler struct
type KeyServiceHandler struct {
	utils.HttpUtils
	KeyInteractor KeyInteractor
	Session       *sessions.CookieStore
}

// ConfigureKeysRouter configures router routes for key management
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
	secureRouter.Path("/").HandlerFunc(handler.postKey).Methods("POST")
	secureRouter.Path("/").HandlerFunc(handler.getKeys).Methods("GET")
	secureRouter.Path("/{id}").HandlerFunc(handler.deleteKey).Methods("DELETE")
}

func (handler KeyServiceHandler) getKeys(rw http.ResponseWriter, q *http.Request) {
	glog.Info("ConfigureKeysRouter -> getKeys")
	keys := handler.KeyInteractor.GetKeys()
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
	}
	glog.Infof("Retrieving %d keys", len(keys))

	rw.WriteHeader(200)
	for _, k := range keys {
		fmt.Fprintln(rw, k.Key)
	}
}

func (handler KeyServiceHandler) postKey(rw http.ResponseWriter, q *http.Request) {
	glog.Info("ConfigureKeysRouter -> putKeys")
	key := &domain.Key{}
	if handler.ParseBody(q.Body, key) != nil {
		handler.BadRequest(rw)
		return
	}
	if err := handler.KeyInteractor.AddKey(key); err != nil {
		handler.ErrorResponse(rw, &utils.ApiError{Code: 500, Err: err.Error()})
		return
	}
	handler.CreatedWithBody(rw, key)
}

func (handler KeyServiceHandler) deleteKey(rw http.ResponseWriter, q *http.Request) {
	glog.Info("ConfigureKeysRouter -> putKeys")
	idPathParam := mux.Vars(q)["id"]
	if idPathParam == "" {
		handler.BadRequest(rw)
		return
	}
	id, err := strconv.ParseInt(idPathParam, 10, 64)
	if err != nil {
		handler.BadRequest(rw)
		return
	}
	if err = handler.KeyInteractor.DeleteKey(id); err != nil {
		handler.InternalServerError(rw)
		return
	}
	handler.NoContent(rw)
}

// AuthMiddleware function contained on KeyServiceHandler
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
