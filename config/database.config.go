package config

import (
	"fmt"
	"os"
	"stock_management/entity"

	// "github.com/joho/godotenv"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//SetupDatabaseConnection is creating a new connection to our database postgres
func SetupDatabaseConnection() *gorm.DB {
	// errEnv := godotenv.Load()
	// if errEnv != nil {
	// 	panic("Failed to load env file. Make sure .env file is exists!")
	// }

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=require TimeZone=Asia/Shanghai", dbHost, dbUser, dbPass, dbPort, dbName)
	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		panic("Failed to create a connection to database")
	}

	db.AutoMigrate(&entity.User{}, &entity.Product{}, &entity.Cart{}, &entity.History{}, &entity.Purchase{})
	db.LogMode(true)
	println("Database connected !")
	return db
}

//CloseDatabaseConnection method is closing a connection between your app and your db
func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL := db.DB()
	
	dbSQL.Close()
}
