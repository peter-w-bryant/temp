package main

import (
	"backend/routes/gitlab"
	"fmt"
)

func main() {
	fmt.Println("Starting server...")

	// Init GitLab create JSON in GitLab route
	router := gitlab.SetupRouter("/api/create-gitlab-json")

	// TODO: Other routes

	router.Run(":8080")
}
