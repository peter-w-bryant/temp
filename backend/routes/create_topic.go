package routes

import (
	"backend/utils"
	"encoding/json"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func SetupCreateTopicRouter(router *gin.Engine) *gin.Engine {
	router.POST("/api/create-topic", handleCreateTopicPostRequest)
	return router
}

func handleCreateTopicPostRequest(c *gin.Context) {
	// Access the data passed in the request body as JSON
	var gitlabPayload map[string]interface{}
	c.BindJSON(&gitlabPayload)

	// Extract the topic name, and the action from the payload
	topicName, actionName := utils.ExtractTopicNameAndActionName(gitlabPayload)

	fmt.Printf("Topic name: %s\n", topicName)
	fmt.Printf("Action name: %s\n", actionName)

	// If the action is not "approved", ignore the request
	if actionName != "approved" {
		fmt.Println("Action is not 'approved', ignoring...")
		c.JSON(200, gin.H{
			"message": "Action is not 'approved', ignoring...",
		})
		return
	}

	// Read the topic spec JSON file for the topic
	filePath := "../topic_specs/" + topicName + ".json"
	topicSpecJSON, err := os.ReadFile(filePath)
	genericErrorHandler(500, err, c)

	// Store the JSON data into a map
	var topicMap map[string]interface{}
	err = json.Unmarshal(topicSpecJSON, &topicMap)
	genericErrorHandler(500, err, c)

	// Get the cluster ID from the REST Proxy
	url := topicMap["url"].(string)
	clusterID, err := utils.GetClusterID(url)
	genericErrorHandler(500, err, c)
	fmt.Println("Cluster ID:", clusterID)

	// Create the topic
	err = utils.CreateTopic(url, clusterID, topicMap)
	genericErrorHandler(500, err, c)

	c.JSON(200, gin.H{
		"message": "Topic created successfully",
	})

}

func genericErrorHandler(statusCode int, err error, c *gin.Context) {
	if err != nil {
		c.JSON(statusCode, gin.H{
			"error": err.Error(),
		})
		return
	}
}
