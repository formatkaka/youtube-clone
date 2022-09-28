package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var client *sql.DB

type DBConfig struct {
	DBUserName string
	DBPassword string
	DBHost     string
	DBName     string
	DBPort     string
}

func Init() {
	config := DBConfig{
		DBUserName: "postgres",
		DBPassword: "postgres",
		DBHost:     "host.docker.internal",
		DBPort:     "5431",
		DBName:     "youtube",
	}

	url := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.DBHost, config.DBPort, config.DBUserName, config.DBPassword, config.DBName)
	var err error
	client, err = sql.Open("postgres", url)

	if err != nil {
		// unable to connect to DB
		// critical error
		log.Fatal(err)
	}

	fmt.Println("DB connection successful")
}

func GetClient() *sql.DB {
	return client
}
