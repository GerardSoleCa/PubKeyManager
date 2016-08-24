package server

import (
	"net/http"
	"strconv"

	"github.com/GerardSoleCa/PubKeyManager/server/handlers"
	"github.com/GerardSoleCa/PubKeyManager/usecases"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/op/go-logging"

	"github.com/GerardSoleCa/PubKeyManager/infrastructure"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
)

var glog = logging.MustGetLogger("server")

// Server struct
type Server struct {
	KeyInteractor  *usecases.KeyInteractor
	UserInteractor *usecases.UserInteractor
	Configuration  *infrastructure.Configuration
}

//Start loads, starts and bootstraps the server
func (s Server) Start() {
	glog.Debug("Loading server module...")

	router := configureRouter()
	n := setupContext(router)

	store := sessions.NewCookieStore(securecookie.GenerateRandomKey(256))

	handlers.ConfigureAuthHandler(router, s.UserInteractor, store)
	handlers.ConfigureKeysRouter(router, s.KeyInteractor, store)
	handlers.ConfigureStaticRouter(router)

	glog.Debugf("Server listening on port %d", s.Configuration.Port)
	glog.Fatal(http.ListenAndServe(":"+strconv.Itoa(s.Configuration.Port), context.ClearHandler(n)))
}

func configureRouter() (router *mux.Router) {
	router = mux.NewRouter().StrictSlash(false)
	router.KeepContext = true
	router.StrictSlash(true)
	return router
}

func setupContext(router *mux.Router) *negroni.Negroni {
	n := negroni.Classic()
	n.UseHandler(router)
	context.ClearHandler(n)
	return n
}
