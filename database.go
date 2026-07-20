package main

import (
	"fmt"
	"memori/server/models"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// var DB *sql.DB
var DB *gorm.DB

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

	//
	//dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	//

	//db, err := sql.Open("postgres", dsn)
	if err != nil {
		fmt.Println("Unable to Connect to Database")
	}

	DB = db
	err = db.AutoMigrate(&models.Topic{})
	if err != nil {
		fmt.Println("Unable to Migrate Database")
	}
	//response := db.Ping()
	//fmt.Println(response)
	fmt.Println("Database connection established")
}
