package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
)

// ConfigureStaticRouter is used for publishing static directory to internet
func ConfigureStaticRouter(router *mux.Router) {
	//staticRouter := mux.NewRouter().PathPrefix("/").Subrouter().StrictSlash(false)
	//staticRouter.Handle("/", http.FileServer(http.Dir("./static")))
	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./public/"))))
}
