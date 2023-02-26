package main

import (
	"scraper-service/internal/http/gin/routers"

	"github.com/gin-gonic/gin"
)

var server *gin.Engine

func init() {
	server = gin.Default()
	routers.SetupRouters(server)
}

func main() {
	server.Run(":8080")
}