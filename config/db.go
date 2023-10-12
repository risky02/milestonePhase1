package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// Init
const (
	host     = "localhost"
	port     = "3306"
	user     = "root"
	password = ""
)

// Init DB
func InitDB() (*sql.DB, error) {
	// Set up database connection
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, "sports")
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	// Ping to database
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	// Return database connection
	return db, nil
}
