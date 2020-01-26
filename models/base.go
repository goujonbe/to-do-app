package models

import (
	"fmt"
	"os"

	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func init() {
	host := os.Getenv("HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("USERNAME")
	dbName := os.Getenv("DB")

	dbUri := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable", host, port, user, dbName)
	fmt.Println(dbUri)

	conn, err := gorm.Open("postgres", dbUri)
	if err != nil {
		fmt.Print(err)
	}

	db = conn
	db.AutoMigrate(&Task{})
}

func GetDB() *gorm.DB {
	return db
}
