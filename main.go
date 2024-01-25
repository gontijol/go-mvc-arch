package main

import (
	"bv-api/src/controller/routes"
	"log"

	"github.com/gin-gonic/gin"
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

	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
