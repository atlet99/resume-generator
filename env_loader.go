package main

import (
	"github.com/joho/godotenv"
	"log"
)

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}
