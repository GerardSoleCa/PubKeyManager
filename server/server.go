package server


import (
	"net/http"
	"strconv"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/op/go-logging"
	"github.com/rs/cors"
	"github.com/GerardSoleCa/PubKeyManager/server/handlers"
)

var glog = logging.MustGetLogger("server")

func Load(){
	glog.Debug("Loading server module...")

	c := setupCors()
	router := configureRouter()

	handlers.ConfigureStaticRouter(router)
	handlers.ConfigureKeysRouter(router)

	n := negroni.Classic()
	n.Use(c)
	n.UseHandler(router)
	// Clear context route --> Context is a data holder to share data across stacked middlewares
	context.ClearHandler(n)

	glog.Debugf("Server listening on port %d", 8080)
	glog.Fatal(http.ListenAndServe(":" + strconv.Itoa(8080), n))
}

func setupCors() (c *cors.Cors) {
	c = cors.New(cors.Options{
		AllowedOrigins:   []string{"https://pre-decipher.aaaida.com", "http://localhost:8100"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"X-Requested-With", "Content-Type", "Cache-Control", "token", "Authorization", "Accept"},
		Debug:            false,
	})
	return c;
}

func configureRouter() (router *mux.Router) {
	router = mux.NewRouter().StrictSlash(false)
	router.KeepContext = true
	router.StrictSlash(true)
	//router.NotFoundHandler = func(rw http.ResponseWriter, q *http.Request) {
	//	handlers.WriteErrorResponse(rw, "Not found", 404)
	//}
	return router
}