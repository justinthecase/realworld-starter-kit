package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Router() http.Handler {
	router := mux.NewRouter()

	handler := cors.Default().Handler(router)

	return handler
}
