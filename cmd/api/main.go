package main

import (
	"uooobarry/soup/config"
	"uooobarry/soup/internal/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitDB()

	r := gin.Default()
	handler.InitRoutes(r)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
