package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

func TagRouter() http.Handler {
	r := mux.NewRouter()

	r.Path("/tag").Methods("GET").HandlerFunc(getTags)

	return r
}

func getTags(w http.ResponseWriter, r *http.Request) {

}
