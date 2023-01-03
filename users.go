package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

// User represents a user in the database
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

// GetUsersHandler returns a handler function that retrieves all users from the database and returns them as JSON
func GetUsersHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Retrieve all users from the database
		rows, err := db.Query("SELECT id, username, password, email FROM users")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		// Convert the rows to a slice of User structs
		var users []User
		for rows.Next() {
			var u User
			err := rows.Scan(&u.ID, &u.Username, &u.Password, &u.Email)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			users = append(users, u)
		}

		// Return the users as JSON
		err = json.NewEncoder(w).Encode(users)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
