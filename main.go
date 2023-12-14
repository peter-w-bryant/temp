package main

import (
	"encoding/json"
	"os"
	"os/exec"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/post", handlePostRequest)

	router.Run(":8080")
}

func handlePostRequest(c *gin.Context) {
	var jsonData map[string]interface{}

	if err := c.BindJSON(&jsonData); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	filename, err := createJSONFile(jsonData)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if err := gitCommitAndCreateMergeRequest(filename); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

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
