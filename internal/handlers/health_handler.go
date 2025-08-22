package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "healthy",
		"service": "gitlab-tool",
		"version": "1.0.0",
		"timestamp": gin.H{
			"unix": time.Now().Unix(),
			"iso":  time.Now().Format(time.RFC3339),
		},
	})
}

func (h *HealthHandler) ReadinessCheck(c *gin.Context) {
	// Add database connectivity check here if needed
	c.JSON(http.StatusOK, gin.H{
		"status":  "ready",
		"service": "gitlab-tool",
	})
}
