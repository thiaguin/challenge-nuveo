package repositories

import (
	"backend/utils"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// NewDB func
func NewDB() *gorm.DB {
	port := utils.GetEnvVariable("DB_PORT")
	dbName := utils.GetEnvVariable("DB_NAME")
	localhost := utils.GetEnvVariable("DB_HOST")
	password := utils.GetEnvVariable("DB_PASSWORD")
	username := utils.GetEnvVariable("DB_USERNAME")

	dbConfig := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		localhost,
		username,
		password,
		dbName,
		port,
	)

	db, err := gorm.Open(postgres.Open(dbConfig), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	return db
}
