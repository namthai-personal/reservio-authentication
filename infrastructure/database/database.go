package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database struct {
	*gorm.DB
}
type GormLogger struct{}

var DB *gorm.DB

func OpenDbConnection() *gorm.DB {

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			LogLevel: logger.Info,
			Colorful: true,
		},
	)

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")

	dbUri := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable ", host, port, username, dbName, password)
	db, err := gorm.Open(postgres.Open(dbUri), &gorm.Config{Logger: newLogger})

	if err != nil {
		panic(err)
	}
	DB = db
	return DB
}

func GetDb() *gorm.DB {
	return DB
}
