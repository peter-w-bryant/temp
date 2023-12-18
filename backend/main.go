package main

import (
	"backend/logger"
	"backend/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	log.SetOutput(logger.Logger.Writer())      // Set the global logger to use the custom logger
	gin.DefaultWriter = logger.Logger.Writer() // Use the custom logger for gin
	log.Println("Starting GIN server...")

	router := gin.Default()

	// /api/create-topic: Create a Kafka topic using Confluent's REST Proxy
	routes.SetupCreateTopicRouter(router)

	router.Run("localhost:8080")
}
