package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

func createLogFile() {
	currentTime := time.Now()
	logFilename := "logtimes.json"

	logEntry := struct {
		Timestamp  string `json:"timestamp"`
		LogNumber  int    `json:"log_number"`
	}{
		Timestamp:  currentTime.Format("2006-01-02 15:04:05"),
		LogNumber:  getLogNumber(logFilename),
	}

	logFile, err := os.OpenFile(logFilename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Error opening log file:", err)
		return
	}
	defer logFile.Close()

	encoder := json.NewEncoder(logFile)
	if err := encoder.Encode(logEntry); err != nil {
		fmt.Println("Error writing to log file:", err)
		return
	}

	fmt.Printf("Log entry added to '%s'.\n", logFilename)
}

func getLogNumber(filename string) int {
	logFile, err := os.OpenFile(filename, os.O_RDONLY, 0644)
	if err != nil {
		if os.IsNotExist(err) {
			return 1
		}
		fmt.Println("Error opening log file:", err)
		return 0
	}
	defer logFile.Close()

	var logEntries []struct {
		Timestamp  string `json:"timestamp"`
		LogNumber  int    `json:"log_number"`
	}
	decoder := json.NewDecoder(logFile)
	for {
		if err := decoder.Decode(&logEntries); err != nil {
			break
		}
	}

	return len(logEntries) + 1
}

func main() {
	createLogFile()
}
