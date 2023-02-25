package main

import (
	"scraper-service/infra/http/gin/handlers"

	"github.com/gin-gonic/gin"
)

var server *gin.Engine

func init() {
	server = gin.Default()
	handlers.SetupRouters(server)
}

func main() {
	server.Run(":8080")
}