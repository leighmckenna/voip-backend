package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func TelnyxWebhook(c *gin.Context) {
	// Handle incoming Telnyx Webhook

	log.Println("Telnyx Webhook recieved")

	c.JSON(http.StatusOK, gin.H{})
}

func HealthCheck(c *gin.Context) {
	// Health check endpoint

	log.Println("Health check endpoint hit")

	c.JSON(http.StatusOK, gin.H{})
}
