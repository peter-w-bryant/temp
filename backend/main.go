package main

import (
	"backend/logger"
	"backend/routes"
	"fmt"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Starting server...")
	gin.DefaultWriter = io.MultiWriter(logger.Logger.Writer(), os.Stdout)

	router := gin.Default()

	// /api/create-topic: Create a Kafka topic using Confluent's REST Proxy
	routes.SetupCreateTopicRouter(router)

	router.Run("localhost:8080")
}
