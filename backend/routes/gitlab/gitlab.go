package gitlab

import (
	"backend/utils"
	"fmt"

	"github.com/gin-gonic/gin"
)

func SetupRouter(route string) *gin.Engine {
	router := gin.Default()
	fmt.Println("GitLab route loaded: " + route)
	router.POST(route, handlePostRequest)
	return router
}

// handlePostRequest handles the POST request and processes the JSON data.
// It binds the JSON data to the jsonData variable, creates a JSON file,
// commits the changes to Git, and creates a merge request.
func handlePostRequest(c *gin.Context) {
	fmt.Println("Handling GitLab POST request...")

	// Bind the JSON data to the jsonData variable --
	// the JSON data is available in the request body
	var jsonData map[string]interface{}

	fmt.Println(c.Request.Body)

	// If the JSON data is invalid, return an error
	if err := c.BindJSON(&jsonData); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Create a JSON file
	filename, err := utils.CreateJSONFile(jsonData)

	// If the JSON file could not be created, return an error
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// Commit the changes to Git and create a merge request
	// if err := gitCommitAndCreateMergeRequest(filename); err != nil {
	// 	c.JSON(500, gin.H{"error": err.Error()})
	// 	return
	// }

	c.JSON(200, gin.H{"message": "File processed and merge request created", "file": filename})
}
