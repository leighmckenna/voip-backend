package main

import (
	"internal/api"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/recieve_messages", api.TelnyxWebhook)

	router.Run("localhost:8080")
}
