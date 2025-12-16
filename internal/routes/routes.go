package routes

import (
	file_infrastructure "go-back-comp/internal/files/infrastructure"

	"github.com/gin-gonic/gin"
)

// SetupRoutes configures all application routes
func SetupRoutes(r *gin.Engine) {
	file_infrastructure.RegisterRoutes(r)
}
