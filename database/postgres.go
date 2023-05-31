package database

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitPostgresFromDotenv() (*gorm.DB, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	host := os.Getenv("PG_HOST")
	user := os.Getenv("PG_USER")
	pass := os.Getenv("PG_PASS")
	dbn := os.Getenv("PG_DB")
	port := os.Getenv("PG_PORT")
	ssl := os.Getenv("PG_SSL")
	time := os.Getenv("PG_TIMEZONE")

	setup := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		host,
		user,
		pass,
		dbn,
		port,
		ssl,
		time,
	)
	db, err := gorm.Open(postgres.Open(setup), &gorm.Config{})
	return db, nil
}
