package models

import (
	"database/sql"

	"github.com/apex/log"
	// import mysql driver
	_ "github.com/go-sql-driver/mysql"
)

var (
	db *sql.DB
)

// Connect to DB
func Connect() {
	var err error

	/*
	 * MySQL client
	 */
	db, err = sql.Open("mysql", "social_admin:Pa$$w0rd@/social_network")
	if err != nil {
		log.WithError(err).Fatal("error connecting to DB")
	}

	// Configure pool size
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(0)

	err = db.Ping()
	if err != nil {
		log.WithError(err).Fatal("error while pinging DB")
	}
}

// GetDB session
func GetDB() *sql.DB {
	return db
}
