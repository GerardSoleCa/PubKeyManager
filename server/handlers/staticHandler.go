package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
)

func ConfigureStaticRouter(router *mux.Router) {
	//staticRouter := mux.NewRouter().PathPrefix("/").Subrouter().StrictSlash(false)
	//staticRouter.Handle("/", http.FileServer(http.Dir("./static")))
	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./public/"))))
}
