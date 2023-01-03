package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Connect to the database
	db, err := sql.Open("mysql", "user:password@/database")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	// Initialize the router
	r := NewRouter(db)

	// Use the os package to read the PORT environment variable
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Start the server
	http.ListenAndServe(":"+port, r)
}
