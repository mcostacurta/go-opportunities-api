package handler

import (
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @summary HealthCheck
// @Description Endpoint to check if the server is alive
// @Accept json
// @Produce json
// @Success 200
// @Router /health [get]
func HealthCheck(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "health",
	})
}
