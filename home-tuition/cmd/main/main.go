package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rebontadeb/home-tuition/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.StudentRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":9191", r))
}
