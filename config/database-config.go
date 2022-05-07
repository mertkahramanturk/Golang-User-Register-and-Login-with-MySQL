package config

import (
	"UserLogin2/entity"
	"fmt"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//SetupDataBaseConnection is creating a new connection to our database
func SetupDatabaseConnection() *gorm.DB {

	ErrEnv := godotenv.Load()

	if ErrEnv != nil {
		panic("Failed to load env file")
	}

	/*
		dbUser := os.Getenv("DB_USER")
		dbPass := os.Getenv("DB_PASS")
		dbHost := os.Getenv("DB_HOST")
		dbName := os.Getenv("DB_NAME")
	*/

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to create a connection to your database")
	}

	db.AutoMigrate(&entity.Book{}, &entity.User{})

	return db
}

//CloseDatabseConnection method is closing a connection between your app and your db

func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()

	if err != nil {
		panic("Failed to close connection from database")
	}
	dbSQL.Close()
}
