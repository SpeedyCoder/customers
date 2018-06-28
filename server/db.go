package server

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
)

type dbConfig struct {
	host     string
	port     string
	user     string
	dbname   string
	password string
}

func (config dbConfig) String() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		config.host,
		config.port,
		config.user,
		config.dbname,
		config.password,
	)
}

func getDBConfig() dbConfig {
	return dbConfig{
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASSWORD"),
	}
}

func getStaticDBConfig() dbConfig {
	return dbConfig{
		"localhost",
		"25432",
		"root",
		"customers",
		"root",
	}
}

// MigrateDB performs automagration.
func MigrateDB(db *gorm.DB) {
	db.AutoMigrate(&Customer{})
}

// GetDB returns a GORM database object.
func GetDB() *gorm.DB {
	config := getStaticDBConfig()
	db, err := gorm.Open("postgres", config.String())
	if err != nil {
		msg := fmt.Sprintf("Failed to connect database. Reason: %s", err)
		panic(msg)
	}
	return db
}
