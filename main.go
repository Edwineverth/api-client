package main

import (
	"log"
	"os"

	"api-client/dao"
	"api-client/db"
	"api-client/handler"
	"api-client/repository"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func getEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return fallback
	}
	return value
}

func main() {
	loadEnv()

	mongoURI := getEnv("MONGO_URI", "mongodb://localhost:27017")
	mongoDBName := getEnv("MONGO_DB", "clientDB")
	serverPort := getEnv("SERVER_PORT", "8080")

	// Initialize MongoDB
	mongoDB := &db.MongoDB{}
	mongoDB.Connect(mongoURI, mongoDBName)

	// Initialize DAO, Repository, and Handler
	clientDAO := &dao.ClientDAO{Collection: mongoDB.Database.Collection("client")}
	clientRepo := &repository.ClientRepository{ClientDAO: clientDAO}
	clientHandler := &handler.ClientHandler{ClientRepo: clientRepo}

	// Set up Gin router
	router := gin.Default()
	router.GET("/clients", clientHandler.GetClients)
	router.GET("/clients/:id", clientHandler.GetClientByID)
	router.POST("/clients", clientHandler.CreateClient)
	router.PUT("/clients/:id", clientHandler.UpdateClient)
	router.DELETE("/clients/:id", clientHandler.DeleteClient)

	router.Run(":" + serverPort)
}
