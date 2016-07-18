package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
)

func ConfigureStaticRouter(router *mux.Router){
	router.Handle("/", http.FileServer(http.Dir("./static")))
}
