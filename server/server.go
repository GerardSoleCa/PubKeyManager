package server

import (
	"net/http"
	"strconv"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/op/go-logging"
	"github.com/GerardSoleCa/PubKeyManager/server/handlers"
	"github.com/GerardSoleCa/PubKeyManager/usecases"

	"github.com/gorilla/sessions"
	"github.com/gorilla/securecookie"
)

var glog = logging.MustGetLogger("server")

type Server struct {
	KeyInteractor  *usecases.KeyInteractor
	UserInteractor *usecases.UserInteractor
}

//Load starts and bootstraps the server
func (s Server) Start() {
	glog.Debug("Loading server module...")

	router := configureRouter()
	n := setupContext(router)

	store := sessions.NewCookieStore(securecookie.GenerateRandomKey(256))

	handlers.ConfigureAuthHandler(router, s.UserInteractor, store)
	handlers.ConfigureKeysRouter(router, s.KeyInteractor, store)
	handlers.ConfigureStaticRouter(router)

	glog.Debugf("Server listening on port %d", 8080)
	glog.Fatal(http.ListenAndServe(":" + strconv.Itoa(8080), context.ClearHandler(n)))
}

func configureRouter() (router *mux.Router) {
	router = mux.NewRouter().StrictSlash(false)
	router.KeepContext = true
	router.StrictSlash(true)
	return router
}

func setupContext(router *mux.Router) (*negroni.Negroni) {
	n := negroni.Classic()
	n.UseHandler(router)
	context.ClearHandler(n)
	return n
}