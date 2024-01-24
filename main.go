package main

import (
	"fmt"
	"os"

	"log"

	"github.com/joho/godotenv"
)

// @title BV API Service
// @version 1.0
// @description API para "PROJETO SECRETO"

// @host localhost:8080
// @BasePath /
func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	fmt.Println(os.Getenv("TEST"))

}
