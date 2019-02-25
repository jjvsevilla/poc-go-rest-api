package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/me", Me)
	router.HandleFunc("/users/{username}", GetUser)
	router.HandleFunc("/projects", GetProjects)
	router.HandleFunc("/projects/{projectId}", GetProject)

	server := http.ListenAndServe(":8080", router)

	log.Fatal(server)
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home")
}

func Me(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Me")
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	username := params["username"]

	fmt.Fprintf(w, "You want to fetch information from user "+username)
}

func GetProjects(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Projects")
}

func GetProject(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	projectId := params["projectId"]

	fmt.Fprintf(w, "You want to fetch information from projectId "+projectId)
}
