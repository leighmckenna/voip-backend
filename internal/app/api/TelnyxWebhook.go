package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func TelnyxWebhook(c *gin.Context) {
	// Handle incoming Telnyx Webhook

	c.JSON(http.StatusOK, gin.H{})
}
