package main

import (
	"log"
	"net/http"

	"github.com/justinthecase/real-world-go-mux/router"
)

func main() {
	r := router.Router()

	log.Fatal(http.ListenAndServe(":8080", r))
}
