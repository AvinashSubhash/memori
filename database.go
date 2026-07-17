package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDatabase() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error Loading Environment Variables")
	}

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("PASSWORD")
	db_name := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, pass, db_name)
	fmt.Println(" Connecting ", dsn)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		fmt.Println("Unable to Connect to Database")
	}

	DB = db
	response := db.Ping()
	fmt.Println(response)
	fmt.Println("Database connection established")
}
