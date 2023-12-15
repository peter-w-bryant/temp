package gitlab

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"time"

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
	filename, err := createJSONFile(jsonData)

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

func createJSONFile(data map[string]interface{}) (string, error) {
	fileName := "data-" + time.Now().Format("20060102150405") + ".json"
	fileData, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return "", err
	}

	err = os.WriteFile(fileName, fileData, 0644)
	return fileName, err
}

func gitCommitAndCreateMergeRequest(filename string) error {
	// Git add
	cmd := exec.Command("git", "add", filename)
	if err := cmd.Run(); err != nil {
		return err
	}

	// Git commit
	commitMessage := "Add " + filename
	cmd = exec.Command("git", "commit", "-m", commitMessage)
	if err := cmd.Run(); err != nil {
		return err
	}

	// Git push and create a merge request
	// Replace the following with your GitLab token and project specifics
	gitLabToken := "your_gitlab_token"
	projectID := "your_project_id"
	sourceBranch := "your_branch"
	targetBranch := "main"
	cmd = exec.Command("curl", "-X", "POST", "-H", "PRIVATE-TOKEN: "+gitLabToken,
		"https://gitlab.com/api/v4/projects/"+projectID+"/merge_requests",
		"-d", "source_branch="+sourceBranch,
		"-d", "target_branch="+targetBranch,
		"-d", "title=Merge "+filename)

	return cmd.Run()
}
