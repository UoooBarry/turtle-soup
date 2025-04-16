package main

import (
	"log"
	"os"
	"uooobarry/soup/config"
	"uooobarry/soup/internal/handler"

	"github.com/gin-gonic/gin"
)

func setupLogging() (*os.File, error) {
	if err := os.MkdirAll("log", 0755); err != nil {
		return nil, err
	}

	logFile, err := os.OpenFile("log/app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	log.SetOutput(logFile)
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	return logFile, nil
}

func main() {
	// Initialize logging
	logFile, err := setupLogging()
	if err != nil {
		log.Fatal("Failed to initialize logging:", err)
	}
	defer logFile.Close()

	config.InitDB()

	r := gin.Default()

	handler.InitRoutes(r)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
