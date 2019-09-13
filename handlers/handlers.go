package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Router register necessary routes and returns an instance of a router.
func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/home", home).Methods("GET")
	r.HandleFunc("/status", status).Methods("GET")
	r.HandleFunc("/calculate", calculate).Methods("POST")
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./template/static/")))
	http.Handle("/", r)
	return r
}
