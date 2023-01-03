package main

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
)

// NewRouter returns a new router with the routes and handlers configured
func NewRouter(db *sql.DB) *mux.Router {
	r := mux.NewRouter()

	// Define a route for the root path and a handler function
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/index.html")
	})

	// Define a route for the /users path and a handler function
	r.HandleFunc("/users", GetUsersHandler(db)).Methods("GET")

	// Serve static files from the "static" directory
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	return r
}
