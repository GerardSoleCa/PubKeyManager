package handlers

import (
	"bytes"
	"encoding/json"
	"github.com/GerardSoleCa/PubKeyManager/domain"
	"github.com/GerardSoleCa/PubKeyManager/server/utils"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"net/http"
)

// UserInteractor interface
type UserInteractor interface {
	AddUser(user *domain.User) error
	AuthenticateUser(username, password string) (domain.User, error)
}

// UserServiceHandler struct
type UserServiceHandler struct {
	utils.HttpUtils
	UserInteractor UserInteractor
	Session        *sessions.CookieStore
}

// ConfigureAuthHandler configures router routes for authentication
func ConfigureAuthHandler(router *mux.Router, interactor UserInteractor, session *sessions.CookieStore) {
	handler := UserServiceHandler{}
	handler.UserInteractor = interactor
	handler.Session = session

	router.Path("/auth/login").HandlerFunc(handler.login).Methods("POST")
	router.Path("/auth/register").HandlerFunc(handler.register).Methods("POST")
}

func (handler UserServiceHandler) login(rw http.ResponseWriter, q *http.Request) {
	glog.Debugf("Request > UserHandler > Login")
	user := domain.User{}
	if handler.ParseBody(q.Body, &user) != nil {
		handler.BadRequest(rw)
		return
	}
	user, err := handler.UserInteractor.AuthenticateUser(user.Username, user.Password)
	if err != nil {
		glog.Debugf("Error > UserHandler > Login :: %s", err.Error())
		handler.Unauthorized(rw)
		return
	}
	user.Password = ""
	session, _ := handler.Session.Get(q, "authenticated")
	session.Values["id"] = user.Id
	session.Values["name"] = user.Username
	session.Save(q, rw)

	var response bytes.Buffer
	encoder := json.NewEncoder(&response)
	encoder.Encode(user)
	handler.Ok(rw, response.String())
	glog.Debugf("End Request > UserHandler > Login :: %s", response.String())
}

func (handler UserServiceHandler) register(rw http.ResponseWriter, q *http.Request) {
	user := &domain.User{}
	if handler.ParseBody(q.Body, user) != nil {
		handler.BadRequest(rw)
		return
	}
	if err := handler.UserInteractor.AddUser(user); err != nil {
		handler.ErrorResponse(rw, &utils.ApiError{Code: 500, Err: err.Error()})
		return
	}
	user.Password = ""
	handler.CreatedWithBody(rw, user)
}
