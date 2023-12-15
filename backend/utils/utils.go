package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"time"
)

// Creates a JSON file with the given JSON payload + returns the file name and
// any error that occurred.
func CreateJSONFile(data map[string]interface{}) (string, error) {
	// topic_<topic_name>_<YYYYMMDDHHMMSS>.json
	fmt.Println("Creating JSON file...")

	fmt.Println("Topic name: " + data["topic_name"].(string))

	// Build fname + fpath --> topic_<topic_name>_<YYYYMMDDHHMMSS>.json
	fileName := "topic_" + data["topic_name"].(string) + "_" + time.Now().Format("20060102150405") + ".json"
	filePath := "topic_configs/" + fileName

	fmt.Println("File path: " + filePath)

	// Create the directory if it doesn't exist
	if _, err := os.Stat("topic_configs"); os.IsNotExist(err) {
		err := os.Mkdir("topic_configs", 0755)
		if err != nil {
			return "", err
		}
	}

	fileData, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return "", err
	}

	err = os.WriteFile(filePath, fileData, 0644)
	return filePath, err
}

func GitCommitAndCreateMergeRequest(filename string) error {
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
