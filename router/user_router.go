package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

func UserRouter() http.Handler {
	r := mux.NewRouter()

	r.Path("/user").Methods("POST", "PUT").HandlerFunc(updateUser)
	r.Path("/user/login").Methods("POST").HandlerFunc(loginUser)
	r.Path("/user").Methods("Get").HandlerFunc(getUser)

	return r
}

func updateUser(w http.ResponseWriter, r *http.Request) {

}

func loginUser(w http.ResponseWriter, r *http.Request) {

}

func getUser(w http.ResponseWriter, r *http.Request) {

}
