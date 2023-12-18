package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func ExtractTopicNameAndActionName(gitlabData map[string]interface{}) (string, string) {
	// Get action name ["object_attributes"]["action"]
	actionName := gitlabData["object_attributes"].(map[string]interface{})["action"].(string)
	// Get topic name
	sourceBranch := gitlabData["object_attributes"].(map[string]interface{})["source_branch"].(string)
	// Split source branch by - to get topic name
	topicName := strings.Split(sourceBranch, "-")[1]
	return topicName, actionName
}

func GetClusterID(url string) (string, error) {
	clusterEndpoint := url + "/v3/clusters"
	fmt.Printf("Cluster endpoint: %s\n", clusterEndpoint)

	// Build new GET request object
	req, err := http.NewRequest("GET", clusterEndpoint, nil)
	if err != nil {
		return "", err
	}
	// Set headers
	req.Header.Set("Content-Type", "application/json")

	// Make GET request to REST Proxy for cluster ID
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("Error sending request to REST Proxy: %s\n", err.Error())
		return "", err
	}
	defer resp.Body.Close()
	// Extract JSON data from the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %s\n", err.Error())
		return "", err
	}
	var responseMap map[string]interface{}
	err = json.Unmarshal(body, &responseMap)
	if err != nil {
		fmt.Printf("Error parsing JSON: %s\n", err.Error())
		return "", err
	}

	// Extract the cluster ID from the response
	clusterID := responseMap["data"].([]interface{})[0].(map[string]interface{})["cluster_id"].(string)
	return clusterID, nil
}

func CreateTopic(url string, clusterID string, topicSpec map[string]interface{}) error {
	topicEndpoint := url + "/v3/clusters/" + clusterID + "/topics"
	// fmt.Printf("Topic endpoint: %s\n", topicEndpoint)

	// Build new POST request object
	req, err := http.NewRequest("POST", topicEndpoint, nil)
	genericErrorHandler(err)
	// Set headers
	req.Header.Set("Content-Type", "application/json")
	// Get "topic_name" from topicSpec
	topicName := topicSpec["topic_name"].(string)
	req.Body = io.NopCloser(strings.NewReader(fmt.Sprintf(`{"topic_name": "%s"}`, topicName)))

	// Make POST request to REST Proxy to create topic
	resp, err := http.DefaultClient.Do(req)
	genericErrorHandler(err)
	defer resp.Body.Close()

	// Extract data from the response body
	body, err := io.ReadAll(resp.Body)
	genericErrorHandler(err)
	// Store the JSON data into a map
	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	genericErrorHandler(err)

	// Check if the topic was created successfully
	_, cluster_id_exists := data["cluster_id"]
	// If the cluster ID is not in the response, then assume that the topic was not created successfully
	if !cluster_id_exists {
		fmt.Println("Topic was not created successfully")
		if _, error_code := data["error_code"]; error_code {
			fmt.Println("Topic already exists")
			return fmt.Errorf("Topic already exists")
		} else {
			fmt.Println("Topic was not created successfully")
			return fmt.Errorf("Topic was not created successfully")
		}
	}
	fmt.Println("Topic created successfully")
	return nil
}

func genericErrorHandler(err error) error {
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return err
	}
	return nil
}
