package main

import (
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index).Methods("GET")
	router.HandleFunc("/me", Me).Methods("GET")
	router.HandleFunc("/users/{username}", GetUser).Methods("GET")
	router.HandleFunc("/projects", GetProjects).Methods("GET")
	router.HandleFunc("/projects/{projectId}", GetProject).Methods("GET")
	return router
}
