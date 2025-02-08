package db

import (
	"database/sql"
	"log"
	"time"
)

func InitDB(driverName string, dbAddr string) *sql.DB {
	db, err := sql.Open(driverName, dbAddr)
	if err != nil {
		log.Fatalf("Unable to connect to the database: %v", err)
	}

	db.SetConnMaxLifetime(5 * time.Minute)
	db.SetMaxIdleConns(25)
	db.SetMaxOpenConns(25)

	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping the database: %v", err)
	}

	log.Println("Database connection initiated.")
	return db
}
