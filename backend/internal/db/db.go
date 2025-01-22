package db

import (
    "database/sql"
    log "github.com/sirupsen/logrus"
    "time"
	"fasms/internal/utilities"

    _ "github.com/go-sql-driver/mysql" // MySQL driver
)
var DB *sql.DB 

func NewDB()  error {
    // Open a connection
	dsn, err := utilities.GetDatabaseUrl()
    if err != nil {
        log.Fatal("Unable to get dsn", err)
    } 
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        return  err
    }

    // Set connection pool properties
    db.SetMaxOpenConns(25)              // Maximum open connections
    db.SetMaxIdleConns(25)              // Maximum idle connections
    db.SetConnMaxLifetime(5 * time.Minute) // Connection lifetime

    // Test the connection
    if err := db.Ping(); err != nil {
        return  err
    }


    log.Info("Connected to the MySQL database successfully")
	DB = db
	return nil
}