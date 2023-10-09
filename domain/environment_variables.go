package domain

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type EnvironmentVariables struct {
	PORT               int
	NUMBER_OF_SECONDS  int
	NUMBER_OF_REQUESTS int
	DB_HOST            string
	DB_PORT            int
	DB_USER            string
	DB_PASSWORD        string
	DB_NAME            string
	TEST_DB_NAME       string
}

func NewEnvironmentVariables() *EnvironmentVariables {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	PORT, _ := strconv.Atoi(os.Getenv("PORT"))
	NUMBER_OF_SECONDS, _ := strconv.Atoi(os.Getenv("NUMBER_OF_SECONDS"))
	NUMBER_OF_REQUESTS, _ := strconv.Atoi(os.Getenv("NUMBER_OF_REQUESTS"))
	DB_HOST := os.Getenv("DB_HOST")
	DB_PORT, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	DB_USER := os.Getenv("DB_USER")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_NAME := os.Getenv("DB_NAME")
	TEST_DB_NAME := os.Getenv("TEST_DB_NAME")

	return &EnvironmentVariables{
		PORT:               PORT,
		NUMBER_OF_SECONDS:  NUMBER_OF_SECONDS,
		NUMBER_OF_REQUESTS: NUMBER_OF_REQUESTS,
		DB_HOST:            DB_HOST,
		DB_PORT:            DB_PORT,
		DB_USER:            DB_USER,
		DB_PASSWORD:        DB_PASSWORD,
		DB_NAME:            DB_NAME,
		TEST_DB_NAME:       TEST_DB_NAME,
	}

}
