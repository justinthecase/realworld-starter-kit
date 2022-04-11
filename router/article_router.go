package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

func ArticleRouter() http.Handler {
	r := mux.NewRouter()

	r.Path("/articles").Methods("GET").HandlerFunc(listArticle)
	r.Path("/articles").Methods("POST").HandlerFunc(createArticle)
	r.Path("/articles/feed").Methods("GET").HandlerFunc(feedArticle)
	r.Path("/articles/{slug}").Methods("GET").HandlerFunc(getArticle)
	r.Path("/articles/{slug}").Methods("PUT").HandlerFunc(updateArticle)
	r.Path("/articles/{slug}").Methods("DELETE").HandlerFunc(deleteArticle)
	r.Path("/articles/{slug}/comments").Methods("POST").HandlerFunc(addComments)
	r.Path("/articles/{slug}/comments").Methods("GET").HandlerFunc(getComments)
	r.Path("/articles/{slug}/comments/{id:[0-9]+}").Methods("DELETE").HandlerFunc(deleteComments)
	r.Path("/articles/{slug}/favorite").Methods("POST").HandlerFunc(addFavoriteArticle)
	r.Path("/articles/{slug}/favorite").Methods("DELETE").HandlerFunc(deleteFavoriteArticle)

	return r
}

func listArticle(w http.ResponseWriter, r *http.Request) {

}

func createArticle(w http.ResponseWriter, r *http.Request) {

}

func feedArticle(w http.ResponseWriter, r *http.Request) {

}

func getArticle(w http.ResponseWriter, r *http.Request) {

}

func updateArticle(w http.ResponseWriter, r *http.Request) {

}

func deleteArticle(w http.ResponseWriter, r *http.Request) {

}

func addComments(w http.ResponseWriter, r *http.Request) {

}

func getComments(w http.ResponseWriter, r *http.Request) {

}

func deleteComments(w http.ResponseWriter, r *http.Request) {

}

func addFavoriteArticle(w http.ResponseWriter, r *http.Request) {

}

func deleteFavoriteArticle(w http.ResponseWriter, r *http.Request) {

}
