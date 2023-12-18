package logger

import (
	"io"
	"log"
	"os"
)

var Logger *log.Logger

func init() {
	// Open the log file
	logFile, err := os.OpenFile("backend.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}

	// Create a multi writer
	multiWriter := io.MultiWriter(os.Stdout, logFile)

	// Set the global logger
	Logger = log.New(multiWriter, "", log.LstdFlags)
}
