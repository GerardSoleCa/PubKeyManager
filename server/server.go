package server

import (
	"net/http"
	"strconv"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/op/go-logging"
	"github.com/GerardSoleCa/PubKeyManager/server/handlers"
)

var glog = logging.MustGetLogger("server")

func Load() {
	glog.Debug("Loading server module...")

	router := configureRouter()
	n := setupContext(router)

	handlers.ConfigureAuthHandler(router)
	handlers.ConfigureKeysRouter(router)
	handlers.ConfigureStaticRouter(router)

	glog.Debugf("Server listening on port %d", 8080)
	glog.Fatal(http.ListenAndServe(":" + strconv.Itoa(8080), n))
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