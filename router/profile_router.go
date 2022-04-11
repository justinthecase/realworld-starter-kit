package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

func ProfileRouter() http.Handler {
	r := mux.NewRouter()

	r.Path("/profiles/{username}").Methods("GET").HandlerFunc(getProfile)
	r.Path("/profiles/{username}/follow").Methods("POST").HandlerFunc(followProfile)
	r.Path("/profiles/{username}/follow").Methods("DELETE").HandlerFunc(unfollowProfile)

	return r
}

func getProfile(w http.ResponseWriter, r *http.Request) {

}

func followProfile(w http.ResponseWriter, r *http.Request) {

}

func unfollowProfile(w http.ResponseWriter, r *http.Request) {

}
