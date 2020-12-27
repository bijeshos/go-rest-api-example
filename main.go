package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	//create a new router
	r := mux.NewRouter()

	//specify endpoints
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/profile", ProfileHandler)
	r.HandleFunc("/settings", SettingsHandler)
	http.Handle("/", r)

	//start and listen to requests
	http.ListenAndServe(":8080", r)

}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Home: %v\n", vars["category"])
}

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Profile: %v\n", vars["category"])
}

func SettingsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Settings: %v\n", vars["category"])
}
