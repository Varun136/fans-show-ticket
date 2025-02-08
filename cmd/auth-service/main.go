package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Varun136/fan-show-ticket-booking/db"
	"github.com/Varun136/fan-show-ticket-booking/internal/auth"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	log.Print("Starting Auth service on port 8080.")

	port := os.Getenv("AUTH_PORT")
	db_uri := os.Getenv("DB_URI")
	db_driver := os.Getenv("DB_DRIVER")

	db := db.InitDB(db_driver, db_uri)
	authHandler := auth.NewAuthHandler(db)

	http.HandleFunc("/login", authHandler.LoginHandler)
	http.HandleFunc("/verify-otp", authHandler.VerifyOTPHandler)

	log.Fatal(http.ListenAndServe(port, nil))
}
