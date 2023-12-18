package main

import (
	"backend/routes"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Starting server...")

	router := gin.Default()

	// /api/create-topic: Create a Kafka topic using Confluent's REST Proxy
	routes.SetupCreateTopicRouter(router)

	router.Run("localhost:8080")
}
